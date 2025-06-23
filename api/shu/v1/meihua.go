package v1

import (
	"shu/utility"

	"github.com/gogf/gf/v2/frame/g"
)

type MeihuaReq struct {
	g.Meta `path:"/meihua" tags:"Meihua" method:"get" summary:"获取梅花基础盘信息"`
	Time   string `v:"datetime" dc:"时间,格式如： 2006-01-02 12:00:00"`
}
type MeihuaRes struct {
	*utility.QiguaResult
}

type MeihuaHuReq struct {
	g.Meta  `path:"/meihua/hu" tags:"Meihua" method:"get" summary:"获取梅花上下互卦信息"`
	GuaName string `v:"length:3,4" dc:"完整卦名,如：泽风大过"`
	Gua     string `v:"required-without:GuaName|length:6" dc:"完整卦"`
}
type MeihuaHuRes struct {
	ShangHuGuaName string `json:"shang_hu_gua_name" dc:"上互卦名称"`
	ShangHuGuaYao  string `json:"shang_hu_gua_yao" dc:"上互卦爻"`
	XiaHuGuaName   string `json:"xia_hu_gua_name" dc:"下互卦名称"`
	XiaHuGuaYao    string `json:"xia_hu_gua_yao" dc:"下互卦爻"`
}
