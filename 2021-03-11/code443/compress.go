// 443. 压缩字符串
package main

import (
	"fmt"
	"strconv"
)

// 解题思路：双指针法
// Time:O(n) 所有字符遍历一次
// Space:O(1)
func compress(chars []byte) int {
	var start, target int // 连续字符开始的位置，字符迁移到的目标位置
	for i := 0; i < len(chars); i++ {
		if len(chars)-1 > i && chars[i+1] == chars[i] {
			// 不是最后一位 && 后面的字符和当前字符一样，不做任何操作
			continue
		}
		// 最后一位，或者 字符不同，需要写入 字符 和 count
		chars[target] = chars[i]
		count := i - start + 1
		target++ // 写位置下移一位
		start = i + 1 // 下一个字符开始的位置
		if count <= 1 { // 超过1个才需要保存
			continue
		}
		arr := []byte(strconv.Itoa(count))
		for _, c := range arr {
			chars[target] = c
			target++
		}
	}
	return target
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}
