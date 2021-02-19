package main

import (
	"fmt"
	"github.com/ser163/WordBot/generate"
)

func main() {
	//  1-19随机字符模式
	fmt.Println("1-19 random character lowercase mode:")
	for i := 0; i < 5; i++ {
		w, err := generate.GenRandomWorld(0, "none")
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(w.Chance, w.Word)
	}

	fmt.Println("8 lengths, blends and modes:")
	for i := 0; i < 5; i++ {
		w, err := generate.GenRandomWorld(8, "mix")
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(w.Chance, w.Word)
	}
	fmt.Println("5 lengths, capital letters:")
	for i := 0; i < 5; i++ {
		w, err := generate.GenRandomWorld(5, "title")
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(w.Chance, w.Word)
	}
	fmt.Println("Simplify the model：")
	// 简化模式,获取小写
	w1, err := generate.GenRandomNone(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("none:",w1.Chance, w1.Word)
	// 简化模式，获取小写
	w2, err := generate.GenRandomLower(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("lower:",w2.Chance, w2.Word)

	// 简化模式，大小写混合
	w3, err := generate.GenRandomMix(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("mix:",w3.Chance, w3.Word)

	// 简化模式，获取首字母大写
	w4, err := generate.GenRandomTitle(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("title:",w4.Chance, w4.Word)

	// 简化模式，获取首字母大写
	w5, err := generate.GenRandomUpper(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("upper",w5.Chance, w5.Word)
}
