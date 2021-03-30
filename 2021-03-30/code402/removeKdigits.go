// 402. 移掉K位数字
package main

import (
	"fmt"
	"strings"
)

// 移除k个元素，保留最小值(不含任何前导0)
// 贪心
// Time:O(n) Space:O(n)
func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		// 通过循环，删除k个放在当前位属于最大的元素
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
		fmt.Println(k, string(stack))
	}
	// 如果前面没有删除，则删除后面部分
	stack = stack[:len(stack)-k]
	// 去除前面的0
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

// 移除k个元素，保留最大值
// 贪心
// Time:O(n) Space:O(n)
func removeKdigitsSaveMax(num string, k int) string {
	stack := []byte{}
	for i := 0; i < len(num); i++ {
		c := num[i]
		// 通过循环，删除k个放在当前位属于最小的元素
		for k > 0 && len(stack) > 0 && c > stack[len(stack) - 1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	// 全部的前面都比后面的大，则直接截取后面部分
	stack = stack[:len(stack)-k]
	num = string(stack)
	if num == "" {
		return "0"
	}
	return num
}


func main()  {
	fmt.Println(removeKdigitsSaveMax("5432155555555", 3))
}
