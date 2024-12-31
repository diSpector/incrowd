package storage

import (
	"context"

	"github.com/diSpector/incrowd.git/internal/models/domain"
)

type PollStorage interface {
	GetAtricleBySource(ctx context.Context, sourceId, sourceSystem string) (*domain.Article, error)
	GetLastNIdsModTimeBySource(ctx context.Context, n int, source string) ([]domain.ArticleOriginMod, error)
	SaveArticle(ctx context.Context, article domain.Article) error
	ReplaceArticleById(ctx context.Context, id string, article domain.Article) error
}

type ServerStorage interface {
	GetArticles(ctx context.Context, limit, offset int64) ([]domain.Article, error)
	GetArticleById(ctx context.Context, id string) (*domain.Article, error)
	GetArticlesCount(ctx context.Context) (int64, error)
}
