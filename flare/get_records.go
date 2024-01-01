package flare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func (f *Flare) GetRecords(ctx context.Context, zoneId string) (list []cloudflare.DNSRecord, err error) {
	rc := cloudflare.ResourceContainer{
		Identifier: zoneId,
	}

	params := cloudflare.ListDNSRecordsParams{
		Name:    "jsb.by",
		Comment: "byfly",
	}

	list, _, err = f.c.ListDNSRecords(ctx, &rc, params)
	if err != nil {
		return nil, err

	}

	return list, nil
}
