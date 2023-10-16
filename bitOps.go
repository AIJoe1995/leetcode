package main

import "fmt"

// 位运算
// 取出数字的某个二进制位 操作是  n >> i & 1
//ans |= 1 << i // 把ans的第i位置为1
// 取出为1的最低位 mask = xor & (-xor); ??
// [a, b, b, b, c, c, c] 只有一个元素出现一次 其他元素都出现了n次(如3次) 在考虑二进制位时，每个位%n的结果就是唯一一个元素该二进制位的数字

// 进阶: 需要考虑溢出问题 正负数的二进制位表示 在不同语言中稍有差别 底层实现

func getUniqueElementOutOfSameRepeatedList(slice []int32, repeat int32) int32 {
	//bitarr := make([]int, 0, 32) // 考虑32位
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, n := range slice {
			total += n >> i & 1
		}
		if total%repeat > 0 {
			ans |= 1 << i // 把ans的第i位置为1
		}
	}
	return ans

}

// 如果有两个元素只出现了一次 其他元素都出现了两次 找出这两个元素
//# // 全部元素异或后的结果 a^b
//# 需要从a^b的结果中拆分出a和b, 考虑二进制位， a和b相同的二进制位在a^b中表现为0，a和b不同的二进制位在a^b中为1
//# 找到a^b为1的最低二进制位，在这个二进制位上a和b一个是0 一个是1 而所有元素可以分成在这个二进制位上是1的元素和在这个二进制位上是0的元素， 那么在这两组元素组内做异或就会得到a和b

func singleNumber(nums []int) []int {
	a_xo_b := 0
	for _, n := range nums {
		a_xo_b ^= n
	}
	b := a_xo_b & (-a_xo_b)
	type1, type2 := 0, 0
	for _, n := range nums {
		if n&b > 0 {
			type1 ^= n
		} else {
			type2 ^= n
		}

	}
	return []int{type1, type2}

}

func main() {
	slice := []int32{1, 2, 2, 2, 3, 3, 3}
	repeat := int32(3)
	ans := getUniqueElementOutOfSameRepeatedList(slice, repeat)
	fmt.Println(ans)
}
