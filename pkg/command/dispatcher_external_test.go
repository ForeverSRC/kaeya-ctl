package command_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ForeverSRC/kaeya-ctl/pkg/command"
	"github.com/ForeverSRC/kaeya-ctl/pkg/domain"
)

type mockKaeyaClient struct {
	KVs map[string]domain.KV
}

func (m *mockKaeyaClient) Get(ctx context.Context, key string) (domain.KV, error) {
	res := domain.KV{Key: key}
	kv, ok := m.KVs[key]
	if !ok {
		return res, nil
	}

	return kv, nil
}

func (m *mockKaeyaClient) Set(ctx context.Context, kv domain.KV) error {
	m.KVs[kv.Key] = kv
	return nil
}

func TestDispatch(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"get aaa", "aaa -> "},
		{"set aaa 1", "success"},
		{"get aaa", "aaa -> 1"},
	}

	mc := &mockKaeyaClient{
		KVs: make(map[string]domain.KV),
	}

	dispatcher := command.NewDefaultDispatcher(mc)
	ctx := context.Background()

	for _, c := range cases {
		out, err := dispatcher.Dispatch(ctx, c.input)
		assert.NoError(t, err)
		assert.Equal(t, c.output, out)
	}
}
