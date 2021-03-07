package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 470. 用 Rand7() 实现 Rand10()
// 已有方法rand7可生成 1 到 7 范围内的均匀随机整数，试写一个方法rand10生成 1 到 10 范围内的均匀随机整数。
//不要使用系统的Math.random()方法

// rand10 期望时间复杂度O(1),最坏情况O(∞)
func rand10() int {
	var num int
	for {
		// (randX() - 1) * Y + randY()，可以生成等概率[1, XY]随机数
		num = (rand7() - 1) * 7 + rand7()
		// 在40之内，直接返回
		if num <= 40 {
			return num % 10 + 1
		}
		// 剩余的41-49之间，利用随机数再操作一次
		num = (num - 40) * 7 + rand7()
		// 在60之内，直接返回
		if num <= 60 {
			return num % 10 + 1
		}
		// 剩余61-63之间，利用随机数再操作一次
		num = (num - 60) * 7 + rand7()
		// 在20之内，直接返回
		if num <= 20 {
			return num % 10 + 1
		}
		// 仅剩余21，舍弃
	}
}

// 已有的 rand7
func rand7() int {
	return rand.Intn(7) + 1
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	inp := 1
	arr := make([]int, inp)
	for i := 0; i < inp; i++ {
		arr[i] = rand10()
	}
	fmt.Println(arr)
}
