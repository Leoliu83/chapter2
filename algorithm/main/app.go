package main

import (
	"log"
	"math"
)

func main() {
	// var s stack.Stack
	// s.Init()
	// s.Distory()

	// i := 2500.00

	/*
		i: 预审额度
		a: 最高信用卡授信额度
		b: 申请额度
	*/
	f := func(a float64, b float64, i float64) {
		t1 := math.Ceil(b / i)
		// min := b / a
		// max := (i + b) / a
		min := (i * t1) / a
		max := (1 + t1) * i / a

		log.Println("===================================")
		log.Printf("申请额度: %f", b)
		log.Printf("最高信用卡授信额度: %f", a)
		log.Printf("预审额度: %f", i)
		log.Printf("系数: %f ≤ x ≤ %f ", min, max)
		// log.Println(a*min/i, math.Floor(a*min/i))
		// log.Println(a*max/i, math.Floor(a*max/i))
		log.Printf("验证(min): math.Floor(最高信用卡授信额度 × 系数(min) ÷ 预审额度) × 预审额度-申请额度 = %f", math.Floor(a*min/i)*i-b)
		log.Printf("验证(max): math.Floor(最高信用卡授信额度 × 系数(max) ÷ 预审额度) × 预审额度-申请额度 = %f", math.Floor(a*max/i)*i-b)
		// FLOOR(t.wx_credit_credit_card_highest_amount*t.index/2500)*2500-t.a_carloan_amount >= 0
		log.Println("===================================")
	}

	f(30000.00, 100000.00, 2333.00)
	f(30000.00, 80000.00, 2333.00)
	f(20000.00, 30000.00, 2333.00)

	// f(30000.00, 100000.00, 0.00)
	// f(30000.00, 80000.00, 0.00)
	// f(20000.00, 30000.00, 0.00)

}
