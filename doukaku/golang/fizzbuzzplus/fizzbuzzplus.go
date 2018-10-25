package main

import (
	"fmt"
	"strconv"
)

// 初期データの件数
var count int = 100

//Write a program that prints the numbers from 1 to 100, but...
//
//numbers that are exact multiples of 3,
//or that contain 3, must print a string containing "Fizz"
//  For example 9 -> "...Fizz..."
//  For example 31 -> "...Fizz..."
//
//numbers that are exact multiples of 5,
//or that contain 5, must print a string containing "Buzz"
//  For example 10 -> "...Buzz..."
//  For example 51 -> "...Buzz..."

func main() {
	fmt.Println("--------")
	slice := inputData()
	for _, i := range slice {
		v := fizzBuzzPlus(i)
		fmt.Printf("%v -> %v\n", i, v)
	}
	fmt.Println("--------")
}

// データ作成
func inputData() []int {
	var slice []int
	for i := 1; i <= count; i++ {
		slice = append(slice, i)
	}
	return slice
}

// 剰余判定
func checkMod(a, n int) bool {
	if a%n == 0 {
		return true
	}
	return false
}

// 同一文字判定
func sameStr(src, dist string) bool {
	if src == dist {
		return true
	}
	return false
}

// 対象の判定
func check(i, target int) bool {
	// 剰余判定
	if checkMod(i, target) {
		return true
	}

	// 文字列を一旦rune配列へ変換
	s := strconv.Itoa(i)
	rs := []rune(s)

	// 文字中でtargetを含んでいるか確認
	tv := strconv.Itoa(target)
	flag := false
	for _, runeVal := range rs {
		v := fmt.Sprintf("%c", runeVal)
		flag = flag || sameStr(v, tv)
	}
	return flag
}

// FizzBuzzPlus判定結果
func fizzBuzzPlus(i int) string {

	f3 := check(i, 3)
	f5 := check(i, 5)

	// 戻り値の設定
	v := strconv.Itoa(i)
	if f3 || f5 {
		fizz, buzz := "", ""
		if f3 == true {
			fizz = "Fizz"
		}
		if f5 == true {
			buzz = "Buzz"
		}
		v = fmt.Sprintf("%v%v", fizz, buzz)
	}
	return v
}
