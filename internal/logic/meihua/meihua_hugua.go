package meihua

import (
	"context"
	"fmt"
	v1 "shu/api/shu/v1"
	"shu/utility"
)

func (m *Meihua) HuGua(ctx context.Context, req *v1.MeihuaHuReq) (res *v1.MeihuaHuRes, err error) {
	qigua := utility.NewMeihuaQigua()

	var guaSequence string

	// 如果提供了卦名，先转换为卦序列
	if req.GuaName != "" {
		guaSequence, err = qigua.GetGuaSequenceByName(req.GuaName)
		if err != nil {
			return nil, err
		}
	} else if req.Gua != "" {
		// 直接使用卦序列
		guaSequence = req.Gua
	} else {
		return nil, fmt.Errorf("必须提供卦名或卦序列")
	}

	huGua, err := qigua.GetHuGua(guaSequence)
	if err != nil {
		return nil, err
	}

	return &v1.MeihuaHuRes{
		HuGuaResult: huGua,
	}, nil
}
