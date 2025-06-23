package shu

import (
	"context"

	v1 "shu/api/shu/v1"
)

func (c *ControllerV1) Meihua(ctx context.Context, req *v1.MeihuaReq) (res *v1.MeihuaRes, err error) {
	return c.meihua.Paipan(ctx, req)
}
