package fetch_api

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int
	Err        error
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	client := &http.Client{
		Timeout: timeout,
	}

	type result struct {
		index int
		resp  *APIResponse
	}
	results := make(chan result, len(urls))

	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				results <- result{idx, &APIResponse{URL: u, Err: err}}
				return
			}

			httpResp, err := client.Do(req)
			if err != nil {
				results <- result{idx, &APIResponse{URL: u, Err: err}}
				return
			}
			defer httpResp.Body.Close()

			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				results <- result{idx, &APIResponse{
					URL:        u,
					StatusCode: httpResp.StatusCode,
					Err:        err,
				}}
				return
			}

			results <- result{idx, &APIResponse{
				URL:        u,
				Data:       string(body),
				StatusCode: httpResp.StatusCode,
				Err:        nil,
			}}
		}(i, url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	responses := make([]*APIResponse, len(urls))
	for res := range results {
		responses[res.index] = res.resp
	}

	return responses
}
