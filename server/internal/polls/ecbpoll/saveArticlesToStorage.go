package ecbpoll

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/diSpector/incrowd.git/internal/cache/article/innercache"
	"github.com/diSpector/incrowd.git/internal/models/ecb"
)

func (s EcbPoll) storeArticles(ctx context.Context, articles []ecb.Article) error {
	for i := range articles {
		var lastModifiedTimestamp int64
		var uuid string
		var isExists bool
		var err error

		cacheVal, err := s.cache.Get(ctx, strconv.Itoa(articles[i].Id))
		if err != nil {
			if !errors.Is(err, innercache.ErrNotFoundInCache) {
				log.Println(`err get from cache:`, err)
			}
		} else { // found in cache
			isExists = true
			cacheValSlice := strings.Split(cacheVal, `:`)
			if len(cacheValSlice) < 2 {
				log.Println(`cache key is incorrect:`, cacheVal)
				isExists = false
			} else {
				uuid = cacheValSlice[1]
				lastModifiedTimestamp, err = strconv.ParseInt(cacheValSlice[0], 10, 64)
				if err != nil {
					log.Println(`err parse timestamp from cache:`, err)
					isExists = false
				}
			}
		}

		domainArticle := s.convertArticle(articles[i])

		if !isExists {
			// not exists in cache, need to check in db
			articleDb, err := s.storage.GetAtricleBySource(ctx, domainArticle.Source.SourceId, domainArticle.Source.SourceSystem)
			if err != nil {
				return fmt.Errorf("err get article by source: %s", err)
			}

			if articleDb == nil {
				err = s.storage.SaveArticle(ctx, domainArticle)
				if err != nil {
					return fmt.Errorf("err save article: %s", err)
				}
			} else if articleDb.LastModified.Date.Before(domainArticle.LastModified.Date) {
				domainArticle.Id = articleDb.Id // keep the same id (uuid)
				err = s.storage.ReplaceArticleById(ctx, domainArticle.Id, domainArticle)
				if err != nil {
					return fmt.Errorf("err replace article: %s", err)
				}
			}

		} else if lastModifiedTimestamp < articles[i].LastModified {
			domainArticle.Id = uuid
			err = s.storage.ReplaceArticleById(ctx, domainArticle.Id, domainArticle)
			if err != nil {
				return fmt.Errorf("err replace article: %s", err)
			}
		}

		err = s.cache.Set(ctx, domainArticle.Source.SourceId, fmt.Sprintf("%d:%s", articles[i].LastModified, uuid))
		if err != nil {
			log.Printf("err set to cache: %s", err)
		}
	}

	return nil
}
