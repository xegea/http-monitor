package httprequest

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15"

type Response struct {
	Url            string
	StatusCode     int
	ResponseTimeMs int
	Body           []byte
}

func Get(url string) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Get: failed creating request: %w", err)
	}

	req.Header = http.Header{
		"User-Agent":      {userAgent},
		"Accept":          {"*/*"},
		"Host":            {req.URL.Host},
		"Accept-Encoding": {"gzip, deflate, br"},
		"Connection":      {"keep-alive"},
	}

	var start time.Time
	var duration time.Duration

	var httpClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	start = time.Now()
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Get: failed executing request. Error: %w ", err)
	}
	duration = time.Since(start)

	body, err := decodeGzip(res)
	if err != nil {
		body = []byte("")
	}

	r := Response{
		url,
		res.StatusCode,
		int(duration.Abs().Milliseconds()),
		body,
	}

	return &r, nil
}

func decodeGzip(res *http.Response) ([]byte, error) {

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Get: failed reading body. Error: %w ", err)
	}

	reader := bytes.NewReader([]byte(b))
	gzreader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, fmt.Errorf("Get: failed creating gzip reader. Error: %w ", err)
	}

	output, err := io.ReadAll(gzreader)
	if err != nil {
		return nil, fmt.Errorf("Get: failed decoding gzip. Error: %w ", err)
	}

	return output, nil
}
