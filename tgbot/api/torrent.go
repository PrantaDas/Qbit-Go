package api

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"strings"
)

func (q *QBit) Add(ctx context.Context, url string) (*http.Response, error) {
	var b bytes.Buffer
	form := multipart.NewWriter(&b)
	form.WriteField("urls", url)
	form.WriteField("upLimit", "0")
	form.WriteField("sequentialDownload", "true")
	form.WriteField("firstLastPiecePrio", "true")

	form.WriteField("dummy", "True")
	err := form.Close()
	if err != nil {
		return nil, err
	}

	req := Request{
		endpoint: "torrents/add",
		method:   "POST",
		Type:     form.FormDataContentType(),
		payload:  strings.NewReader(b.String()),
	}

	return q.Do(ctx, req)
}
