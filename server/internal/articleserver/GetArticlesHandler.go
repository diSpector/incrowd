package articleserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/diSpector/incrowd.git/internal/models/domain"
	"github.com/diSpector/incrowd.git/internal/validators"
)

func (s ArticleServer) GetArticlesHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctxDb, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()

		queryParams := r.URL.Query()
		pageSize := queryParams.Get("pageSize")
		pageNumber := queryParams.Get("pageNumber")

		log.Printf("serve: /articles, pageSize = %s, pageNumber = %s\n", pageSize, pageNumber)

		response, err := s.prepareGetArticlesHandlerResponse(ctxDb, pageSize, pageNumber)
		if err != nil {
			log.Println("err during response preparing:", err)
		}

		switch response.Status {
		case STATUS_SUCCESS:
			w.WriteHeader(http.StatusOK)
		case STATUS_ERROR:
			w.WriteHeader(http.StatusInternalServerError)
		case STATUS_FAIL:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(response)
	}
}

func (s ArticleServer) validateGetArticlesHandlerParams(pageSize, pageNumber string) bool {
	if pageSize != `` && !validators.ValidatePositiveInt(pageSize) {
		return false
	}

	if pageNumber != `` && !validators.ValidateNonNegativeInt(pageNumber) {
		return false
	}

	return true
}

func (s ArticleServer) prepareGetArticlesHandlerResponse(ctx context.Context, pageSize, pageNumber string) (domain.ResponseMulti, error) {
	var response domain.ResponseMulti
	var errResponse error

	if !s.validateGetArticlesHandlerParams(pageSize, pageNumber) {
		response.Status = STATUS_FAIL
		response.Message = "incorrect parameters"
	} else {
		var pageSizeInt = int64(DEFAULT_PAGESIZE)
		var pageNumberInt = int64(DEFAULT_PAGENUMBER)
		if pageSize != `` {
			pageSizeInt, _ = strconv.ParseInt(pageSize, 10, 64)
		}
		if pageNumber != `` {
			pageNumberInt, _ = strconv.ParseInt(pageNumber, 10, 64)
		}

		articles, err := s.storage.GetArticles(ctx, pageSizeInt, pageNumberInt*pageSizeInt)
		if err != nil {
			response.Status = STATUS_ERROR
			response.Message = "err get articles from storage"
			errResponse = err
		} else {
			response.Status = STATUS_SUCCESS

			s.prepareArticlesForOutput(articles)
			response.Data.Articles = articles

			count, err := s.storage.GetArticlesCount(ctx)
			if err != nil {
				// won't break the flow here, because it's service information
				log.Println("err count articles:", err)
			}
			response.Metadata.PageItems = new(int64)
			*response.Metadata.PageItems = int64(len(articles))
			response.Metadata.PageNumber = &pageNumberInt
			response.Metadata.PageSize = &pageSizeInt
			response.Metadata.TotalItems = &count
			totalPages := count / pageSizeInt
			if count%pageSizeInt != 0 {
				totalPages++
			}
			response.Metadata.TotalPages = &totalPages
		}
	}

	response.Metadata.CreatedAt = time.Now()

	return response, errResponse
}
