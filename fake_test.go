/*
 * Copyright (c) 2017, 奶爸<1@5.nu>
 * All rights reserved.
 */

package com

import (
	"testing"
	"regexp"
)

func TestFakeUserAgent(t *testing.T) {
	if len(FakeUserAgent()) == 0 {
		t.Error("FakeUserAgent 获取UA不能为空")
	}
}

func TestFakeIP(t *testing.T) {
	pass, err := regexp.Match(`^((?:(?:25[0-5]|2[0-4]\d|(?:1\d{2}|[1-9]?\d))\.){3}(?:25[0-5]|2[0-4]\d|(?:1\d{2}|[1-9]?\d)))$`, []byte(FakeIP()))
	if !pass || err != nil {
		t.Error("FakeIP IP生成错误", err)
	}
}
