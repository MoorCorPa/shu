package utility

import (
	"os"
	"testing"
)

func init() {
	// 设置时区为中国时区
	os.Setenv("TZ", "PRC")
}

func TestMeihuaQigua_Qigua(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试时间起卦
	timeStr := "2025-06-23 17:30:00"
	result, err := mq.Qigua(timeStr)

	if err != nil {
		t.Errorf("时间起卦失败: %v", err)
		return
	}

	// 验证结果不为空
	if result == nil {
		t.Error("起卦结果为空")
		return
	}

	// 验证时间
	if result.Time != timeStr {
		t.Errorf("时间不匹配，期望: %s, 实际: %s", timeStr, result.Time)
	}

	// 验证卦象格式（应该是6位二进制字符串）
	if len(result.ZhuGuaYao) != 6 {
		t.Errorf("主卦爻数不正确，期望: 6, 实际: %d", len(result.ZhuGuaYao))
	}

	// 验证卦名不为空
	if result.ZhuGuaName == "" {
		t.Error("主卦名为空")
	}

	// 验证四柱不为空
	if result.SiZhu == "" {
		t.Error("四柱为空")
	}

	// 验证四柱空亡不为空
	if result.KongWang == "" {
		t.Error("四柱空亡为空")
	}

	t.Logf("时间起卦结果:")
	t.Logf("时间: %s", result.Time)
	t.Logf("四柱: %s", result.SiZhu)
	t.Logf("四柱空亡: %s", result.KongWang)
	t.Logf("主卦: %s (%s)", result.ZhuGuaName, result.ZhuGuaYao)
	t.Logf("互卦: %s (%s)", result.HuGuaName, result.HuGuaYao)
	t.Logf("变卦: %s (%s)", result.BianGuaName, result.BianGuaYao)
	t.Logf("错卦: %s (%s)", result.CuoGuaName, result.CuoGuaYao)
	t.Logf("综卦: %s (%s)", result.ZongGuaName, result.ZongGuaYao)
}

func TestMeihuaQigua_QiguaByNumber(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试双数起卦（使用当前时间）
	result, err := mq.QiguaByNumber(23, 45, "")

	if err != nil {
		t.Errorf("数字起卦失败: %v", err)
		return
	}

	// 验证结果不为空
	if result == nil {
		t.Error("起卦结果为空")
		return
	}

	// 验证卦象格式
	if len(result.ZhuGuaYao) != 6 {
		t.Errorf("主卦爻数不正确，期望: 6, 实际: %d", len(result.ZhuGuaYao))
	}

	// 验证卦名不为空
	if result.ZhuGuaName == "" {
		t.Error("主卦名为空")
	}

	// 验证四柱不为空
	if result.SiZhu == "" {
		t.Error("四柱为空")
	}

	// 验证四柱空亡不为空
	if result.KongWang == "" {
		t.Error("四柱空亡为空")
	}

	t.Logf("双数起卦结果:")
	t.Logf("时间: %s", result.Time)
	t.Logf("四柱: %s", result.SiZhu)
	t.Logf("四柱空亡: %s", result.KongWang)
	t.Logf("主卦: %s (%s)", result.ZhuGuaName, result.ZhuGuaYao)
	t.Logf("互卦: %s (%s)", result.HuGuaName, result.HuGuaYao)
	t.Logf("变卦: %s (%s)", result.BianGuaName, result.BianGuaYao)
	t.Logf("错卦: %s (%s)", result.CuoGuaName, result.CuoGuaYao)
	t.Logf("综卦: %s (%s)", result.ZongGuaName, result.ZongGuaYao)
}

func TestMeihuaQigua_QiguaByManual(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试手动排卦（指定时间）
	result, err := mq.QiguaByManual("111000", 3, "2024-01-15 14:30:00")
	if err != nil {
		t.Errorf("手动排卦失败: %v", err)
		return
	}

	// 验证结果不为空
	if result == nil {
		t.Error("起卦结果为空")
		return
	}

	// 验证主卦
	if result.ZhuGuaYao != "111000" {
		t.Errorf("主卦错误，期望: 111000, 实际: %s", result.ZhuGuaYao)
	}

	// 验证变卦（第3爻变化）
	if result.BianGuaYao != "110000" {
		t.Errorf("变卦错误，期望: 110000, 实际: %s", result.BianGuaYao)
	}

	t.Logf("手动排卦结果:")
	t.Logf("时间: %s", result.Time)
	t.Logf("四柱: %s", result.SiZhu)
	t.Logf("四柱空亡: %s", result.KongWang)
	t.Logf("主卦: %s (%s)", result.ZhuGuaName, result.ZhuGuaYao)
	t.Logf("互卦: %s (%s)", result.HuGuaName, result.HuGuaYao)
	t.Logf("变卦: %s (%s)", result.BianGuaName, result.BianGuaYao)
	t.Logf("错卦: %s (%s)", result.CuoGuaName, result.CuoGuaYao)
	t.Logf("综卦: %s (%s)", result.ZongGuaName, result.ZongGuaYao)
}

func TestMeihuaQigua_QiguaByManual_CurrentTime(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试手动排卦（使用当前时间）
	result, err := mq.QiguaByManual("101010", 1, "")
	if err != nil {
		t.Errorf("手动排卦失败: %v", err)
		return
	}

	// 验证主卦
	if result.ZhuGuaYao != "101010" {
		t.Errorf("主卦错误，期望: 101010, 实际: %s", result.ZhuGuaYao)
	}

	// 验证变卦（第1爻变化）
	if result.BianGuaYao != "101011" {
		t.Errorf("变卦错误，期望: 101011, 实际: %s", result.BianGuaYao)
	}

	t.Logf("手动排卦（当前时间）结果:")
	t.Logf("主卦: %s (%s)", result.ZhuGuaName, result.ZhuGuaYao)
	t.Logf("变卦: %s (%s)", result.BianGuaName, result.BianGuaYao)
}

func TestMeihuaQigua_QiguaByManual_InvalidInput(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试无效输入
	testCases := []struct {
		zhuGuaYao string
		dongYao   int
		timeStr   string
		desc      string
	}{
		{"11100", 3, "", "卦序列长度不足"},
		{"1110001", 3, "", "卦序列长度过长"},
		{"111002", 3, "", "卦序列包含非二进制字符"},
		{"111000", 0, "", "动爻位置小于1"},
		{"111000", 7, "", "动爻位置大于6"},
		{"111000", 3, "invalid-time", "时间格式错误"},
	}

	for _, tc := range testCases {
		_, err := mq.QiguaByManual(tc.zhuGuaYao, tc.dongYao, tc.timeStr)
		if err == nil {
			t.Errorf("%s: 期望错误但没有错误", tc.desc)
		}
	}
}

func TestGetDizhiIndex(t *testing.T) {
	tests := []struct {
		dizhi string
		want  int
	}{
		{"子", 1}, {"丑", 2}, {"寅", 3}, {"卯", 4},
		{"辰", 5}, {"巳", 6}, {"午", 7}, {"未", 8},
		{"申", 9}, {"酉", 10}, {"戌", 11}, {"亥", 12},
	}

	for _, test := range tests {
		result := getDizhiIndex(test.dizhi)
		if result != test.want {
			t.Errorf("getDizhiIndex(%s) = %d, want %d", test.dizhi, result, test.want)
		}
	}
}

func TestCalculateKongWangByOrder(t *testing.T) {
	tests := []struct {
		order int
		want  string
	}{
		{1, "戌亥"},  // 甲子旬
		{11, "申酉"}, // 甲戌旬
		{21, "午未"}, // 甲申旬
		{31, "辰巳"}, // 甲午旬
		{41, "寅卯"}, // 甲辰旬
		{51, "子丑"}, // 甲寅旬
	}

	for _, test := range tests {
		result := calculateKongWangByOrder(test.order)
		if result != test.want {
			t.Errorf("calculateKongWangByOrder(%d) = %s, want %s", test.order, result, test.want)
		}
	}
}

func TestQigua_DifferentTimes(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试不同时间的起卦结果
	times := []string{
		"2024-01-01 00:00:00",
		"2024-06-15 12:30:00",
		"2024-12-31 23:59:59",
	}

	for _, timeStr := range times {
		result, err := mq.Qigua(timeStr)
		if err != nil {
			t.Errorf("时间 %s 起卦失败: %v", timeStr, err)
			continue
		}

		// 验证基本格式
		if len(result.ZhuGuaYao) != 6 {
			t.Errorf("时间 %s 主卦格式错误: %s", timeStr, result.ZhuGuaYao)
		}

		t.Logf("时间 %s 起卦: %s (%s)", timeStr, result.ZhuGuaName, result.ZhuGuaYao)
	}
}

func TestQiguaByNumber_DifferentNumbers(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试不同数字组合（双数起卦）
	numbers := [][]int{
		{1, 1},
		{8, 8},
		{25, 30},
		{7, 14},
		{100, 200},
		{123, 456},
	}

	for _, nums := range numbers {
		result, err := mq.QiguaByNumber(nums[0], nums[1], "")
		if err != nil {
			t.Errorf("数字 %v 起卦失败: %v", nums, err)
			continue
		}

		// 验证基本格式
		if len(result.ZhuGuaYao) != 6 {
			t.Errorf("数字 %v 主卦格式错误: %s", nums, result.ZhuGuaYao)
		}

		t.Logf("数字 %v 起卦: %s (%s)", nums, result.ZhuGuaName, result.ZhuGuaYao)
	}
}

func TestQiguaByNumber_WithCustomTime(t *testing.T) {
	mq := NewMeihuaQigua()

	// 测试双数起卦配合自定义时间
	timeStr := "2024-01-15 14:30:00"
	result, err := mq.QiguaByNumber(23, 45, timeStr)
	if err != nil {
		t.Errorf("双数起卦（自定义时间）失败: %v", err)
		return
	}

	// 验证时间
	if result.Time != timeStr {
		t.Errorf("时间不匹配，期望: %s, 实际: %s", timeStr, result.Time)
	}

	// 验证基本格式
	if len(result.ZhuGuaYao) != 6 {
		t.Errorf("主卦格式错误: %s", result.ZhuGuaYao)
	}

	t.Logf("双数起卦（自定义时间）结果:")
	t.Logf("时间: %s", result.Time)
	t.Logf("主卦: %s (%s)", result.ZhuGuaName, result.ZhuGuaYao)
}
