package flare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func (f *Flare) UpdateDnsRecord(ctx context.Context, zoneId, recordId, content string) (err error) {
	rc := cloudflare.ResourceContainer{
		Identifier: zoneId,
	}
	upParams := cloudflare.UpdateDNSRecordParams{
		ID:      recordId,
		Content: content,
	}
	_, err = f.c.UpdateDNSRecord(ctx, &rc, upParams)

	return err
}
