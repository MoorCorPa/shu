package utility

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/calendar"
)

// MeihuaQigua 梅花易数起卦工具类
type MeihuaQigua struct{}

// NewMeihuaQigua 创建新的梅花易数起卦实例
func NewMeihuaQigua() *MeihuaQigua {
	return &MeihuaQigua{}
}

// 八卦基本信息
var (
	// 八卦名称对应表
	guaNames = map[string]string{
		"111": "乾", "000": "坤", "001": "震", "010": "坎",
		"100": "艮", "101": "离", "011": "兑", "110": "巽",
	}

	// 完整六十四卦名称对应表
	liushisiguaNames = map[string]string{
		"111111": "乾为天", "111011": "天泽履", "111101": "天火同人", "111001": "天雷无妄",
		"111110": "天风小畜", "111010": "天水讼", "111100": "天山遁", "111000": "天地否",
		"011111": "泽天夬", "011011": "兑为泽", "011101": "泽火革", "011001": "泽雷随",
		"011110": "泽风大过", "011010": "泽水困", "011100": "泽山咸", "011000": "泽地萃",
		"101111": "火天大有", "101011": "火泽睽", "101101": "离为火", "101001": "火雷噬嗑",
		"101110": "火风鼎", "101010": "火水未济", "101100": "火山旅", "101000": "火地晋",
		"001111": "雷天大壮", "001011": "雷泽归妹", "001101": "雷火丰", "001001": "震为雷",
		"001110": "雷风恒", "001010": "雷水解", "001100": "雷山小过", "001000": "雷地豫",
		"110111": "风天小畜", "110011": "风泽中孚", "110101": "风火家人", "110001": "风雷益",
		"110110": "巽为风", "110010": "风水涣", "110100": "风山渐", "110000": "风地观",
		"010111": "水天需", "010011": "水泽节", "010101": "水火既济", "010001": "水雷屯",
		"010110": "水风井", "010010": "坎为水", "010100": "水山蹇", "010000": "水地比",
		"100111": "山天大畜", "100011": "山泽损", "100101": "山火贲", "100001": "山雷颐",
		"100110": "山风蛊", "100010": "山水蒙", "100100": "艮为山", "100000": "山地剥",
		"000111": "地天泰", "000011": "地泽临", "000101": "地火明夷", "000001": "地雷复",
		"000110": "地风升", "000010": "地水师", "000100": "地山谦", "000000": "坤为地",
	}

	// 天干
	tiangan = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	// 地支
	dizhi = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
)

// QiguaResult 起卦结果
type QiguaResult struct {
	Time        string `json:"time"`
	SiZhu       string `json:"si_zhu"`
	KongWang    string `json:"kong_wang"`
	ZhuGuaName  string `json:"zhu_gua_name"`
	ZhuGuaYao   string `json:"zhu_gua_yao"`
	HuGuaName   string `json:"hu_gua_name"`
	HuGuaYao    string `json:"hu_gua_yao"`
	BianGuaName string `json:"bian_gua_name"`
	BianGuaYao  string `json:"bian_gua_yao"`
	CuoGuaName  string `json:"cuo_gua_name"`
	CuoGuaYao   string `json:"cuo_gua_yao"`
	ZongGuaName string `json:"zong_gua_name"`
	ZongGuaYao  string `json:"zong_gua_yao"`
}

// HuGuaResult 互卦结果
type HuGuaResult struct {
	ShangHuGuaName string `json:"shang_hu_gua_name"`
	ShangHuGuaYao  string `json:"shang_hu_gua_yao"`
	XiaHuGuaName   string `json:"xia_hu_gua_name"`
	XiaHuGuaYao    string `json:"xia_hu_gua_yao"`
}

// Qigua 根据时间起卦（按地支取数）
func (m *MeihuaQigua) Qigua(timeStr string) (*QiguaResult, error) {
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)

	year := int64(t.Year())
	month := int64(t.Month())
	day := int64(t.Day())
	hour := int64(t.Hour())
	minute := int64(t.Minute())
	second := int64(t.Second())

	c := calendar.BySolar(year, month, day, hour, minute, second)

	// 获取干支信息
	bytes, _ := c.ToJSON()
	var result map[string]interface{}
	json.Unmarshal(bytes, &result)

	ganzhi := result["ganzhi"].(map[string]interface{})

	// 年和时使用地支序号，月和日使用农历数值
	yearDizhi := getDizhiIndex(extractDizhi(ganzhi["year"].(string)))
	monthNumber := int(result["lunar"].(map[string]interface{})["month"].(float64))
	dayNumber := int(result["lunar"].(map[string]interface{})["day"].(float64))
	hourDizhi := getDizhiIndex(extractDizhi(ganzhi["hour"].(string)))

	// 计算卦数
	shangGuaNum := (yearDizhi + monthNumber + dayNumber) % 8
	if shangGuaNum == 0 {
		shangGuaNum = 8
	}

	xiaGuaNum := (yearDizhi + monthNumber + dayNumber + hourDizhi) % 8
	if xiaGuaNum == 0 {
		xiaGuaNum = 8
	}

	dongYaoNum := (yearDizhi + monthNumber + dayNumber + hourDizhi) % 6
	if dongYaoNum == 0 {
		dongYaoNum = 6
	}

	// 生成卦象
	shangGua := numToGua(shangGuaNum)
	xiaGua := numToGua(xiaGuaNum)
	zhuGua := shangGua + xiaGua

	// 计算各种卦象
	huGua := zhuGua[1:4] + zhuGua[2:5]
	bianGua := calculateBianGua(zhuGua, dongYaoNum)
	cuoGua := calculateCuoGua(zhuGua)
	zongGua := reverseGua(zhuGua)

	// 计算四柱和四柱空亡
	siZhu := calculateSiZhu(t)
	kongWang := calculateSiZhuKongWang(t)

	return &QiguaResult{
		Time:        timeStr,
		SiZhu:       siZhu,
		KongWang:    kongWang,
		ZhuGuaName:  getGuaName(zhuGua),
		ZhuGuaYao:   zhuGua,
		HuGuaName:   getGuaName(huGua),
		HuGuaYao:    huGua,
		BianGuaName: getGuaName(bianGua),
		BianGuaYao:  bianGua,
		CuoGuaName:  getGuaName(cuoGua),
		CuoGuaYao:   cuoGua,
		ZongGuaName: getGuaName(zongGua),
		ZongGuaYao:  zongGua,
	}, nil
}

// QiguaByNumber 双数起卦法
func (m *MeihuaQigua) QiguaByNumber(shangShu, xiaShu int, timeStr string) (*QiguaResult, error) {
	// 解析时间，如果为空则使用当前时间
	var t time.Time
	var err error
	if timeStr == "" {
		t = time.Now()
		timeStr = t.Format("2006-01-02 15:04:05")
	} else {
		t, err = time.Parse("2006-01-02 15:04:05", timeStr)
		if err != nil {
			return nil, fmt.Errorf("时间格式错误: %v", err)
		}
	}

	// 获取时辰的地支序号
	year := int64(t.Year())
	month := int64(t.Month())
	day := int64(t.Day())
	hour := int64(t.Hour())
	minute := int64(t.Minute())
	second := int64(t.Second())

	c := calendar.BySolar(year, month, day, hour, minute, second)
	bytes, _ := c.ToJSON()
	var result map[string]interface{}
	json.Unmarshal(bytes, &result)
	ganzhi := result["ganzhi"].(map[string]interface{})
	hourDizhi := getDizhiIndex(extractDizhi(ganzhi["hour"].(string)))

	// 计算卦数
	shangGuaNum := shangShu % 8
	if shangGuaNum == 0 {
		shangGuaNum = 8
	}

	xiaGuaNum := xiaShu % 8
	if xiaGuaNum == 0 {
		xiaGuaNum = 8
	}

	dongYaoNum := (shangShu + xiaShu + hourDizhi) % 6
	if dongYaoNum == 0 {
		dongYaoNum = 6
	}

	// 生成卦象
	shangGua := numToGua(shangGuaNum)
	xiaGua := numToGua(xiaGuaNum)
	zhuGua := shangGua + xiaGua

	// 计算各种卦象
	huGua := zhuGua[1:4] + zhuGua[2:5]
	bianGua := calculateBianGua(zhuGua, dongYaoNum)
	cuoGua := calculateCuoGua(zhuGua)
	zongGua := reverseGua(zhuGua)

	// 计算四柱和四柱空亡
	siZhu := calculateSiZhu(t)
	kongWang := calculateSiZhuKongWang(t)

	return &QiguaResult{
		Time:        timeStr,
		SiZhu:       siZhu,
		KongWang:    kongWang,
		ZhuGuaName:  getGuaName(zhuGua),
		ZhuGuaYao:   zhuGua,
		HuGuaName:   getGuaName(huGua),
		HuGuaYao:    huGua,
		BianGuaName: getGuaName(bianGua),
		BianGuaYao:  bianGua,
		CuoGuaName:  getGuaName(cuoGua),
		CuoGuaYao:   cuoGua,
		ZongGuaName: getGuaName(zongGua),
		ZongGuaYao:  zongGua,
	}, nil
}

// QiguaByManual 手动排卦法
func (m *MeihuaQigua) QiguaByManual(zhuGuaYao string, dongYao int, timeStr string) (*QiguaResult, error) {
	// 验证卦序列格式
	if len(zhuGuaYao) != 6 {
		return nil, fmt.Errorf("卦序列必须是6位二进制字符串")
	}
	for _, char := range zhuGuaYao {
		if char != '0' && char != '1' {
			return nil, fmt.Errorf("卦序列只能包含0和1")
		}
	}

	// 验证动爻位置
	if dongYao < 1 || dongYao > 6 {
		return nil, fmt.Errorf("动爻位置必须在1-6之间")
	}

	// 解析时间，如果为空则使用当前时间
	var t time.Time
	var err error
	if timeStr == "" {
		t = time.Now()
		timeStr = t.Format("2006-01-02 15:04:05")
	} else {
		t, err = time.Parse("2006-01-02 15:04:05", timeStr)
		if err != nil {
			return nil, fmt.Errorf("时间格式错误: %v", err)
		}
	}

	// 计算各种卦象
	huGua := zhuGuaYao[1:4] + zhuGuaYao[2:5]
	bianGua := calculateBianGua(zhuGuaYao, dongYao)
	cuoGua := calculateCuoGua(zhuGuaYao)
	zongGua := reverseGua(zhuGuaYao)

	// 计算四柱和四柱空亡
	siZhu := calculateSiZhu(t)
	kongWang := calculateSiZhuKongWang(t)

	return &QiguaResult{
		Time:        timeStr,
		SiZhu:       siZhu,
		KongWang:    kongWang,
		ZhuGuaName:  getGuaName(zhuGuaYao),
		ZhuGuaYao:   zhuGuaYao,
		HuGuaName:   getGuaName(huGua),
		HuGuaYao:    huGua,
		BianGuaName: getGuaName(bianGua),
		BianGuaYao:  bianGua,
		CuoGuaName:  getGuaName(cuoGua),
		CuoGuaYao:   cuoGua,
		ZongGuaName: getGuaName(zongGua),
		ZongGuaYao:  zongGua,
	}, nil
}

// extractDizhi 从干支字符串中提取地支
func extractDizhi(ganzhi string) string {
	runes := []rune(ganzhi)
	if len(runes) >= 2 {
		return string(runes[1]) // 地支是第二个字符
	}
	return ""
}

// getDizhiIndex 获取地支索引（1-12）
func getDizhiIndex(dizhiStr string) int {
	for i, dz := range dizhi {
		if dz == dizhiStr {
			return i + 1
		}
	}
	return 1
}

// numToGua 数字转换为八卦二进制表示
func numToGua(num int) string {
	guaMap := map[int]string{
		1: "111", 2: "011", 3: "101", 4: "001",
		5: "110", 6: "010", 7: "100", 8: "000",
	}
	return guaMap[num]
}

// getGuaName 获取卦名
func getGuaName(gua string) string {
	if len(gua) == 6 {
		if name, exists := liushisiguaNames[gua]; exists {
			return name
		}
		shangGua := gua[:3]
		xiaGua := gua[3:]
		shangName := guaNames[shangGua]
		xiaName := guaNames[xiaGua]
		return xiaName + shangName
	} else if len(gua) == 3 {
		if name, exists := guaNames[gua]; exists {
			return name
		}
	}
	return "未知"
}

// calculateBianGua 计算变卦
func calculateBianGua(zhuGua string, dongYao int) string {
	bianGua := []rune(zhuGua)
	index := 6 - dongYao
	if bianGua[index] == '1' {
		bianGua[index] = '0'
	} else {
		bianGua[index] = '1'
	}
	return string(bianGua)
}

// calculateCuoGua 计算错卦
func calculateCuoGua(zhuGua string) string {
	result := ""
	for _, char := range zhuGua {
		if char == '1' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

// reverseGua 计算综卦
func reverseGua(gua string) string {
	runes := []rune(gua)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// calculateSiZhu 计算四柱
func calculateSiZhu(t time.Time) string {
	year := int64(t.Year())
	month := int64(t.Month())
	day := int64(t.Day())
	hour := int64(t.Hour())
	minute := int64(t.Minute())
	second := int64(t.Second())

	c := calendar.BySolar(year, month, day, hour, minute, second)
	bytes, err := c.ToJSON()
	if err != nil {
		return ""
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return ""
	}

	ganzhi, ok := result["ganzhi"].(map[string]interface{})
	if !ok {
		return ""
	}

	yearGanzhi := ganzhi["year"].(string)
	monthGanzhi := ganzhi["month"].(string)
	dayGanzhi := ganzhi["day"].(string)
	hourGanzhi := ganzhi["hour"].(string)

	return fmt.Sprintf("%s年 %s月 %s日 %s时", yearGanzhi, monthGanzhi, dayGanzhi, hourGanzhi)
}

// calculateSiZhuKongWang 计算四柱空亡
func calculateSiZhuKongWang(t time.Time) string {
	year := int64(t.Year())
	month := int64(t.Month())
	day := int64(t.Day())
	hour := int64(t.Hour())
	minute := int64(t.Minute())
	second := int64(t.Second())

	c := calendar.BySolar(year, month, day, hour, minute, second)
	bytes, err := c.ToJSON()
	if err != nil {
		return ""
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return ""
	}

	ganzhi, ok := result["ganzhi"].(map[string]interface{})
	if !ok {
		return ""
	}

	// 获取四柱的序号
	yearOrder := getGanzhiOrder(ganzhi["year"].(string))
	monthOrder := getGanzhiOrder(ganzhi["month"].(string))
	dayOrder := int(ganzhi["day_order"].(float64))
	hourOrder := getGanzhiOrder(ganzhi["hour"].(string))

	// 计算每柱的空亡
	yearKong := calculateKongWangByOrder(yearOrder)
	monthKong := calculateKongWangByOrder(monthOrder)
	dayKong := calculateKongWangByOrder(dayOrder)
	hourKong := calculateKongWangByOrder(hourOrder)

	return fmt.Sprintf("%s %s %s %s", yearKong, monthKong, dayKong, hourKong)
}

// getGanzhiOrder 获取干支序号
func getGanzhiOrder(ganzhi string) int {
	runes := []rune(ganzhi)
	if len(runes) != 2 {
		return 1
	}

	tg := string(runes[0])
	dz := string(runes[1])

	var tgIndex, dzIndex int
	for i, t := range tiangan {
		if t == tg {
			tgIndex = i
			break
		}
	}
	for i, d := range dizhi {
		if d == dz {
			dzIndex = i
			break
		}
	}

	// 计算60甲子序号
	for i := 0; i < 60; i++ {
		if i%10 == tgIndex && i%12 == dzIndex {
			return i + 1
		}
	}
	return 1
}

// calculateKongWangByOrder 根据序号计算空亡
func calculateKongWangByOrder(order int) string {
	xunShou := (order - 1) / 10

	var kong1, kong2 string
	switch xunShou {
	case 0:
		kong1, kong2 = "戌", "亥"
	case 1:
		kong1, kong2 = "申", "酉"
	case 2:
		kong1, kong2 = "午", "未"
	case 3:
		kong1, kong2 = "辰", "巳"
	case 4:
		kong1, kong2 = "寅", "卯"
	case 5:
		kong1, kong2 = "子", "丑"
	default:
		kong1, kong2 = "戌", "亥"
	}

	return kong1 + kong2
}

// GetHuGua 获取卦的上互卦和下互卦
func (m *MeihuaQigua) GetHuGua(guaSequence string) (*HuGuaResult, error) {
	// 验证卦序列格式
	if len(guaSequence) != 6 {
		return nil, fmt.Errorf("卦序列必须是6位二进制字符串")
	}
	for _, char := range guaSequence {
		if char != '0' && char != '1' {
			return nil, fmt.Errorf("卦序列只能包含0和1")
		}
	}

	// 计算上互卦
	// 上互1：取前五位，前三位+后三位
	front5 := guaSequence[:5]           // 前五位：10100
	shangHu1 := front5[:3] + front5[2:] // 前三位(101) + 后三位(100) = 101100

	// 上互2：取前四位，前三位+后三位
	front4 := guaSequence[:4]           // 前四位：1010
	shangHu2 := front4[:3] + front4[1:] // 前三位(101) + 后三位(010) = 101010

	// 计算下互卦
	// 下互1：取后五位，前三位+后三位
	back5 := guaSequence[1:]        // 后五位：01001
	xiaHu1 := back5[:3] + back5[2:] // 前三位(010) + 后三位(001) = 010001

	// 下互2：取后四位，前三位+后三位
	back4 := guaSequence[2:]        // 后四位：1001
	xiaHu2 := back4[:3] + back4[1:] // 前三位(100) + 后三位(001) = 100001

	// 获取卦名
	shangHu1Name := getGuaName(shangHu1)
	shangHu2Name := getGuaName(shangHu2)
	xiaHu1Name := getGuaName(xiaHu1)
	xiaHu2Name := getGuaName(xiaHu2)

	return &HuGuaResult{
		ShangHuGuaName: shangHu1Name + "," + shangHu2Name,
		ShangHuGuaYao:  shangHu1 + "," + shangHu2,
		XiaHuGuaName:   xiaHu1Name + "," + xiaHu2Name,
		XiaHuGuaYao:    xiaHu1 + "," + xiaHu2,
	}, nil
}

// GetGuaSequenceByName 根据卦名获取卦序列
func (m *MeihuaQigua) GetGuaSequenceByName(guaName string) (string, error) {
	// 遍历六十四卦名称对应表找到匹配的卦序列
	for sequence, name := range liushisiguaNames {
		if name == guaName {
			return sequence, nil
		}
	}
	return "", fmt.Errorf("未找到卦名: %s", guaName)
}
