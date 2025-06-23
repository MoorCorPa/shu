// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package shu

import (
	"context"

	"shu/api/shu/v1"
)

type IShuV1 interface {
	Meihua(ctx context.Context, req *v1.MeihuaReq) (res *v1.MeihuaRes, err error)
	MeihuaHu(ctx context.Context, req *v1.MeihuaHuReq) (res *v1.MeihuaHuRes, err error)
}
