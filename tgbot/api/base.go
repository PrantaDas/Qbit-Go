package api

import (
	"context"
	"net/http"
)

// Construct a QBitClient
func New(username string, password string, url string) QbitClient {
	return &QBit{
		base:     url + "/api/v2/",
		username: username,
		password: password,
	}
}

func (q *QBit) Do(ctx context.Context, r Request) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, r.method, q.base+r.endpoint, r.payload)
	if err != nil {
		return nil, err
	}

	if r.method != "GET" {
		req.Header.Add("Content-Type", r.Type)
	}
	req.Header.Add("Referer", q.base)
	req.Header.Add("Cookie", q.cookie)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer res.Body.Close()

	if len(res.Cookies()) > 0 {
		q.cookie = res.Cookies()[0].String()
	}
	return res, nil
}
