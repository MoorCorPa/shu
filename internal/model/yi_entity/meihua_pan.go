package yi_entity

type MeihuaPan struct {
	Time        string   `json:"time" dc:"起卦时间"`
	SiZhu       string   `json:"si_zhu" dc:"四柱"`
	KongWang    string   `json:"kong_wang" dc:"空亡"`
	ZhuGuaName  string   `json:"zhu_gua_name" dc:"主卦名称"`
	ZhuGuaYao   []string `json:"zhu_gua_yao" dc:"主卦爻"`
	HuGuaName   string   `json:"hu_gua_name" dc:"互卦名称"`
	HuGuaYao    []string `json:"hu_gua_yao" dc:"互卦爻"`
	BianGuaName string   `json:"bian_gua_name" dc:"变卦名称"`
	BianGuaYao  []string `json:"bian_gua_yao" dc:"变卦爻"`
	CuoGuaName  string   `json:"cuo_gua_name" dc:"错卦名称"`
	CuoGuaYao   []string `json:"cuo_gua_yao" dc:"错卦爻"`
	ZongGuaName string   `json:"zong_gua_name" dc:"综卦名称"`
	ZongGuaYao  []string `json:"zong_gua_yao" dc:"综卦爻"`
}
