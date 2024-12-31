package articleserver

import (
	"github.com/diSpector/incrowd.git/internal/models/domain"
	"github.com/diSpector/incrowd.git/internal/storage"
)

const (
	STATUS_SUCCESS = `success` // 1XX/2XX/3XX
	STATUS_FAIL    = `fail`    // 4XX
	STATUS_ERROR   = `error`   // 5XX

	DEFAULT_PAGESIZE   = 10
	DEFAULT_PAGENUMBER = 0
)

type ArticleServer struct {
	storage storage.ServerStorage
}

func New(storage storage.ServerStorage) *ArticleServer {
	return &ArticleServer{storage: storage}
}

func (s ArticleServer) prepareArticlesForOutput(articles []domain.Article) {
	for i := range articles {
		// we "hide" only service information about source now,
		// but we could apply more complex logic here as well
		s.prepareOneArticleForOutput(&articles[i])
	}
}

func (s ArticleServer) prepareOneArticleForOutput(article *domain.Article) {
	if article != nil {
		article.Source = nil
	}
}
