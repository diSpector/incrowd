package ecbpoll

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/diSpector/incrowd.git/internal/cache/article"
	"github.com/diSpector/incrowd.git/internal/storage"
)

const (
	ECB_SOURCE           = `ecb`
	VERSION              = 16
	DEFAULT_PAGESIZE     = 10
	PAGESIZE_MAX_LIMIT   = 20
	DEFAULT_MAX_ARTICLES = 100
	MAX_ARTICLES_LIMIT   = 1000
)

type EcbPoll struct {
	url         string
	maxArticles int
	pageSize    int
	period      time.Duration
	name        string
	storage     storage.PollStorage
	cache       article.Cache
}

func New(url string, maxArticles, pageSize int, period time.Duration, name string, storage storage.PollStorage, cache article.Cache) *EcbPoll {
	if maxArticles <= 0 || maxArticles > MAX_ARTICLES_LIMIT {
		log.Println(`maxArticles set to:`, DEFAULT_MAX_ARTICLES)
		maxArticles = DEFAULT_MAX_ARTICLES
	}

	if pageSize <= 0 || pageSize > PAGESIZE_MAX_LIMIT {
		log.Println(`pageSize set to:`, DEFAULT_PAGESIZE)
		pageSize = DEFAULT_PAGESIZE
	}

	if pageSize > maxArticles {
		log.Println(`pageSize set to:`, maxArticles)
		pageSize = maxArticles
	}

	return &EcbPoll{
		url:         url,
		maxArticles: maxArticles,
		pageSize:    pageSize,
		period:      period,
		name:        name,
		storage:     storage,
		cache:       cache,
	}
}

func (s EcbPoll) Poll(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ctxCache, cancelCache := context.WithTimeout(ctx, 2*time.Minute)
	defer cancelCache()
	err := s.prepareCache(ctxCache, 2*s.maxArticles, ECB_SOURCE)
	if err != nil {
		// if we got an error from cache, we can continue
		log.Printf("err caching articles from db: %s", err)
	}

	for {
		// here we use a strategy for get n = ecb_api.max (from config) amount of articles
		// from api divided by pages with size = ecb_api.pagesize

		numApiCalls := s.maxArticles / s.pageSize

		for i := 0; i < numApiCalls; i++ {
			log.Printf("polling api page = %d, page size = %d\n", i, s.pageSize)
			// polling api endpoint
			ctxPoll, cancelPoll := context.WithTimeout(ctx, 2*time.Minute)
			articles, err := s.pollApi(ctxPoll, i, s.pageSize)
			if err != nil {
				log.Printf("err get articles from %s api: %s", ECB_SOURCE, err)
			} else {
				ctxDb, cancelDb := context.WithTimeout(ctx, 2*time.Minute)
				log.Printf("processing %d articles from api to db\n", len(articles.Content))
				err = s.storeArticles(ctxDb, articles.Content)
				if err != nil {
					log.Printf("err save articles to storage: %s", err)
				}
				cancelDb()
			}
			cancelPoll()

			// exit loop if error from api or from db is not nil, or
			// page number equals to the last api page
			if err != nil || i+1 >= articles.PageInfo.NumPages {
				break
			}
		}

		select {
		case <-ctx.Done():
			log.Printf(`poll for %s stopped`, ECB_SOURCE)
			return
		case <-time.After(s.period):
		}
	}
}
