package com

import "testing"

func TestMD5(t *testing.T) {
	if MD5("nb") != "9c59153d22d9f7cc021b17b425cc31c5" {
		t.Error("MD5 错误")
	}
}
