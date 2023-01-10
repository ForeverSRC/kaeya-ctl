package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/ForeverSRC/kaeya-ctl/pkg/domain"
)

const (
	httpSchema = "http://"
)

type KaeyaClient interface {
	Getter
	Setter
}

type Getter interface {
	Get(ctx context.Context, key string) (domain.KV, error)
}

type Setter interface {
	Set(ctx context.Context, kv domain.KV) error
}

type DefaultKaeyaClient struct {
	*clientopts
}

type clientopts struct {
	httpClient *http.Client
	addr       string
}

type Option func(opts *clientopts)

func WithHttpClient(c *http.Client) Option {
	return func(opts *clientopts) {
		opts.httpClient = c
	}
}

func NewDefaultKaeyaClient(addr string, options ...Option) *DefaultKaeyaClient {
	opts := &clientopts{
		addr: addr,
	}

	for _, op := range options {
		op(opts)
	}

	if opts.httpClient == nil {
		opts.httpClient = http.DefaultClient
	}

	return &DefaultKaeyaClient{
		opts,
	}
}

func (d *DefaultKaeyaClient) Get(ctx context.Context, key string) (domain.KV, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, httpSchema+path.Join(d.addr, "kv", key), nil)
	if err != nil {
		return domain.KV{}, err
	}

	body, err := d.send(req)
	if err != nil {
		return domain.KV{}, err
	}

	var resp KVDataResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return domain.KV{}, err
	}

	if resp.Code != CodeSuccess {
		return domain.KV{}, fmt.Errorf("server code=%d, msg=%s", resp.Code, resp.Message)
	}

	return domain.KV{
		Key:   resp.Data.Key,
		Value: resp.Data.Value,
	}, nil

}

func (d *DefaultKaeyaClient) Set(ctx context.Context, kv domain.KV) error {
	data := KVData{
		Key:   kv.Key,
		Value: kv.Value,
	}

	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, httpSchema+path.Join(d.addr, "kv"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	body, err := d.send(req)
	if err != nil {
		return err
	}

	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	if resp.Code != CodeSuccess {
		return fmt.Errorf("server code=%d, msg=%s", resp.Code, resp.Message)
	}

	return nil

}

func (d *DefaultKaeyaClient) send(req *http.Request) ([]byte, error) {
	resp, err := d.httpClient.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("code=%d, msg=%s", resp.StatusCode, string(b))
	}

	return b, nil

}
