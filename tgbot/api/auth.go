package api

import (
	"context"
	"net/http"
	"strings"
)

func (q *QBit) Login(ctx context.Context) (*http.Response, error) {
	req := Request{
		method:   "POST",
		endpoint: "auth/login",
		Type:     "application/x-www-form-urlencoded",
		payload:  strings.NewReader("username=" + q.username + "&password=" + q.password),
	}

	return q.Do(ctx, req)
}
