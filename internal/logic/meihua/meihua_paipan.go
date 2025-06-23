package meihua

import (
	"context"
	"fmt"
	v1 "shu/api/shu/v1"
	"shu/utility"
)

func (m *Meihua) Paipan(ctx context.Context, req *v1.MeihuaReq) (res *v1.MeihuaRes, err error) {
	// 创建梅花易数起卦工具实例
	qigua := utility.NewMeihuaQigua()

	var result *utility.QiguaResult

	// 根据起卦类型选择相应的起卦方式
	switch req.Type {
	case "time":
		result, err = qigua.Qigua(req.Time)
		if err != nil {
			return nil, fmt.Errorf("时间起卦失败: %v", err)
		}

	case "number":
		// 双数起卦
		result, err = qigua.QiguaByNumber(req.ShangShu, req.XiaShu, req.Time)
		if err != nil {
			return nil, fmt.Errorf("双数起卦失败: %v", err)
		}

	case "manual":
		// 手动排卦
		if req.Manual == "" {
			return nil, fmt.Errorf("手动排卦需要提供卦序列")
		}
		result, err = qigua.QiguaByManual(req.Manual, req.DongYao, req.Time)
		if err != nil {
			return nil, fmt.Errorf("手动排卦失败: %v", err)
		}

	default:
		return nil, fmt.Errorf("不支持的起卦类型: %s", req.Type)
	}

	// 转换为API响应格式
	return &v1.MeihuaRes{
		QiguaResult: result,
	}, nil
}
