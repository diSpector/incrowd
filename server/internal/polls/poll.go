package polls

import (
	"context"
	"sync"
)

type Poller interface {
	Poll(ctx context.Context, wg *sync.WaitGroup)
}
