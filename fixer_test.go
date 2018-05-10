package com

import (
	"testing"
)

func TestFXRate_Convert(t *testing.T) {
	m, err := NewFixer()
	if err != nil {
		t.Error("汇率接口请求失败", err)
	}
	if m.Convert("EUR", "CNY", 1) < 7 {
		t.Error("汇率转换不正确")
	}
}
