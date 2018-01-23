package com

import (
	"testing"
)

func TestFXRate_Convert(t *testing.T) {
	m := GetRate()
	if m == nil {
		t.Error("汇率接口请求失败")
	}
	if m.Convert("EUR", "CNY", 1) < 7 {
		t.Error("汇率转换不正确")
	}
}
