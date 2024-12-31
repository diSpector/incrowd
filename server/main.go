package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/diSpector/incrowd.git/internal/articleserver"
	"github.com/diSpector/incrowd.git/internal/cache/article/innercache"
	"github.com/diSpector/incrowd.git/internal/config"
	"github.com/diSpector/incrowd.git/internal/polls/ecbpoll"
	"github.com/diSpector/incrowd.git/internal/storage/mongodb"
	"github.com/gorilla/mux"
)

func main() {
	defConfigPath := os.Getenv(`DEFAULT_API_CONFIG_PATH`)

	// read config from -config flag or default value
	confPath := flag.String("config", defConfigPath, "application config path")
	flag.Parse()

	conf, err := config.Read(*confPath)
	if err != nil {
		log.Fatalf("err process config file: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// init storage: mongodb
	storage, err := mongodb.New(conf.Storage.User, conf.Storage.Password, conf.Storage.Host, conf.Storage.Port, conf.Storage.Database, conf.Storage.Collection)
	if err != nil {
		log.Fatalf("err connect to storage: %s", err)
	}

	// init cache: innercache
	cache := innercache.New(ctx, 1*time.Hour)

	var pollWgs sync.WaitGroup

	// init polls (now we have only one poll for ecb)
	ecbPoll := ecbpoll.New(conf.EcbApi.Url, conf.EcbApi.Max, conf.EcbApi.PageSize, conf.EcbApi.Period, conf.EcbApi.Name, storage, cache)
	pollWgs.Add(1)
	go ecbPoll.Poll(ctx, &pollWgs)

	articleServer := articleserver.New(storage)

	// init router
	r := mux.NewRouter()

	// define the endpoints
	r.HandleFunc("/articles", articleServer.GetArticlesHandler(ctx)).Methods("GET")
	r.HandleFunc("/articles/{id}", articleServer.GetOneArticleHandler(ctx)).Methods("GET")

	// prepare and run server
	srv := &http.Server{
		Addr:         conf.HttpServer.Address,
		Handler:      r,
		ReadTimeout:  conf.HttpServer.Timeout,
		WriteTimeout: conf.HttpServer.Timeout,
		IdleTimeout:  conf.HttpServer.IdleTimeout,
	}

	log.Println("run server on:", conf.HttpServer.Address)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("failed to start server:", err)
		}
	}()

	<-sigChan
	// graceful shutdown
	log.Println("shutting down application...")

	cancel()
	pollWgs.Wait()

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		log.Println("HTTP server shutdown error:", err)
	}

	storageCloseCtx, cancelStorageClose := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelStorageClose()
	err = storage.Close(storageCloseCtx)
	if err != nil {
		log.Println("storage close err:", err)
	}

	log.Println("application stopped")
}
