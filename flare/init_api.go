package flare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

type Flare struct {
	c *cloudflare.API
}

func InitApi(ctx context.Context, token string) (*Flare, error) {
	api, err := cloudflare.NewWithAPIToken(token)
	if err != nil {
		return nil, err
	}

	flare := &Flare{
		c: api,
	}

	return flare, nil
}
