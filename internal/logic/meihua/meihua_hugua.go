package meihua

import (
	"context"
	v1 "shu/api/shu/v1"
)

func (m *Meihua) HuGua(ctx context.Context, req *v1.MeihuaHuReq) (res *v1.MeihuaHuRes, err error) {
	return &v1.MeihuaHuRes{
		ShangHuGuaName: req.GuaName,
		ShangHuGuaYao:  req.Gua,
		XiaHuGuaName:   req.GuaName,
		XiaHuGuaYao:    req.Gua,
	}, nil
}
