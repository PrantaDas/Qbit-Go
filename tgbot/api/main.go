package api

import (
	"context"
	"net/http"
	"strings"
)

type QBit struct {
	base     string
	username string
	password string
	cookie   string
}

type Request struct {
	method   string
	endpoint string
	Type     string
	payload  *strings.Reader
}

type QbitClient interface {
	// Auth
	Login(context.Context) (*http.Response, error)

	// Torrent
	Add(context.Context, string) (*http.Response, error)
}
