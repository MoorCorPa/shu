package shu

import (
	"context"

	v1 "shu/api/shu/v1"
)

func (c *ControllerV1) MeihuaHu(ctx context.Context, req *v1.MeihuaHuReq) (res *v1.MeihuaHuRes, err error) {
	return c.meihua.HuGua(ctx, req)
}
