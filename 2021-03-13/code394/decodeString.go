// 394. 字符串解码
package main

import (
	"fmt"
	"strconv"
)

// 方法一：栈
// Time: O(n) 所有字符遍历一次
// Space:O(n) 增加list
func decodeString1(s string) string {
	list := make([]string, 0)
	var target, count string
	for _, c := range s {
		if c == '[' { // 如果遇到的是 [， 存入 target，count
			list = append(list, target, count)
			target, count = "", ""
		} else if c == ']' { // 遇到 ], 弹出 target, count
			countTmp, _ := strconv.Atoi(list[len(list)-1])
			targetTmp := list[len(list)-2]
			list = list[:len(list)-2]
			for i := 0; i < countTmp; i++ { // 重复N次
				targetTmp += target
			}
			target = targetTmp
		} else if c >= '0' && c <= '9' { // 数字，即重复的次数
			count += string(c)
		} else { // 需要重复的字符
			target += string(c)
		}
	}
	return target
}

// 方法二：递归
// 解题思想：和栈思想一致，遇到[ 往下递归，遇到 ] 则返回结果
// Time: O(n)
// Space: O(n)
func decodeString(s string) string {
	target, _ := getString(s, 0)
	return target
}

func getString(s string, pos int) (string, int) {
	var target, sub, count string
	for pos < len(s) {
		c := s[pos]
		if c == '[' { // 如果遇到的是 [， 往下递归
			sub, pos = getString(s, pos+1)
			countTmp, _ := strconv.Atoi(count)
			for i := 0; i < countTmp; i++ { // 重复N次
				target += sub
			}
			count = ""
		} else if c == ']' { // 遇到 ], 递归结束
			return target, pos
		} else if c >= '0' && c <= '9' { // 数字，即重复的次数
			count += string(c)
		} else { // 需要重复的字符
			target += string(c)
		}
		pos++
	}
	return target, pos
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	//fmt.Println(decodeString("abc3[cd]xyz"))
	//fmt.Println(decodeString("3[a2[c]]"))
}
