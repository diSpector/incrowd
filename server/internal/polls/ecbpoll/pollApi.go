package ecbpoll

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diSpector/incrowd.git/internal/models/ecb"
)

func (s EcbPoll) pollApi(ctx context.Context, pageNum, pageSize int) (ecb.Articles, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s.url, nil)
	if err != nil {
		return ecb.Articles{}, fmt.Errorf("failed to create request: %s", err)
	}

	q := req.URL.Query()
	q.Add("page", fmt.Sprintf("%d", pageNum))
	q.Add("pageSize", fmt.Sprintf("%d", pageSize))

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ecb.Articles{}, fmt.Errorf("failed to complete request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ecb.Articles{}, fmt.Errorf("err code from api: %d", resp.StatusCode)
	}

	var result ecb.Articles

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ecb.Articles{}, fmt.Errorf("err decode response: %s", err)
	}

	return result, nil
}
