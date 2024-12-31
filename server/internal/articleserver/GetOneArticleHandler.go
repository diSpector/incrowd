package articleserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/diSpector/incrowd.git/internal/models/domain"
	"github.com/gorilla/mux"
)

func (s ArticleServer) GetOneArticleHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctxDb, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()

		vars := mux.Vars(r)
		articleId := vars["id"]

		log.Printf("serve: /articles/{id}, id = %s\n", articleId)

		response, err := s.prepareGetOneArticleHandlerResponse(ctxDb, articleId)
		if err != nil {
			log.Println("err during response preparing:", err)
		}

		switch response.Status {
		case STATUS_SUCCESS:
			w.WriteHeader(http.StatusOK)
		case STATUS_FAIL:
			w.WriteHeader(http.StatusNotFound)
		case STATUS_ERROR:
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(response)
	}
}

func (s ArticleServer) prepareGetOneArticleHandlerResponse(ctx context.Context, articleId string) (domain.ResponseOne, error) {
	var response domain.ResponseOne

	article, err := s.storage.GetArticleById(ctx, articleId)
	if err != nil {
		response.Status = STATUS_ERROR
		response.Message = "err get article by id from storage"
	} else {
		if article == nil {
			response.Status = STATUS_FAIL
			response.Message = "article not found"
		} else {
			response.Status = STATUS_SUCCESS
			s.prepareOneArticleForOutput(article)
		}

		response.Data.Article = article
	}

	response.Metadata.CreatedAt = time.Now()

	return response, err
}
