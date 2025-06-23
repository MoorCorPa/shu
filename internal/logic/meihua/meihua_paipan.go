package meihua

import (
	"context"
	v1 "shu/api/shu/v1"
	"shu/utility"
)

func (m *Meihua) Paipan(ctx context.Context, req *v1.MeihuaReq) (res *v1.MeihuaRes, err error) {
	// 创建梅花易数起卦工具实例
	qigua := utility.NewMeihuaQigua()

	// 使用时间起卦
	result, err := qigua.Qigua(req.Time)
	if err != nil {
		return nil, err
	}

	// 转换为API响应格式
	return &v1.MeihuaRes{
		QiguaResult: result,
	}, nil
}
