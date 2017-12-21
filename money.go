package com

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"encoding/json"
)

type Money struct {
	Base  string
	Date  string
	Rates map[string]float64
}

var moneyInstance *Money

func NewMoney() *Money {
	if moneyInstance == nil {
		req := gorequest.New()
		_, body, err := req.Get("https://api.fixer.io/latest").End()
		if err != nil {
			log.Println("NewMoney", err)
			return nil
		}
		var money Money
		errs := json.Unmarshal([]byte(body), &money)
		if errs != nil {
			log.Println("NewMoney", errs)
			return nil
		}
		moneyInstance = &money
	}
	return moneyInstance
}
func (m *Money) Convert(src string, dist string, num float64) float64 {
	rate, has := m.Rates[dist]
	rate1, has1 := m.Rates[src]
	if has && (has1 || src == m.Base) {
		if src == m.Base {
			return num * rate
		} else {
			return num / rate1 * rate
		}
	} else {
		log.Println("Money.Convert", "原始或目标汇率不支持")
	}
	return num
}
