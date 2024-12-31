package ecbpoll

import (
	"context"
	"fmt"
	"log"
)

func (s EcbPoll) prepareCache(ctx context.Context, n int, sourceSystem string) error {
	lastNIdsMod, err := s.storage.GetLastNIdsModTimeBySource(ctx, n, sourceSystem)
	if err != nil {
		return fmt.Errorf("err get last n articles: %s", err)
	} else {
		// caching last n articles (SourceId, date of last modification)
		// cache key: SourceId, cache value: timestampLastMod_uuid
		for i := range lastNIdsMod {
			err = s.cache.Set(ctx, lastNIdsMod[i].Source.SourceId, fmt.Sprintf("%d:%s", lastNIdsMod[i].LastModified.Date.UnixMilli(), lastNIdsMod[i].Id))
			if err != nil {
				log.Printf("err set to cache: %s", err)
			}
		}
	}

	return nil
}
