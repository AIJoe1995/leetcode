package main

//https://leetcode.cn/problems/movement-of-robots/solutions/2478642/qing-xi-xiang-xi-zheng-ming-wei-shi-yao-8kl3t/?envType=daily-question&envId=2023-10-10
import "sort"

// 求数组中两两元素之间距离的和
// 需要先排序 排序统计求和 时间复杂度nlog(n)

func sumOfDistanceEachTwoElemV2(nums []int) int {
	// 基于所有元素之间的距离公式1
	// 推导过程
	// 第一个元素到后n-1个元素的距离和 可以写作: （di表示i-1元素到i的距离, 即元素分割的第i段距离 n个元素共有n-1段）
	// d1 + (d1 + d2) + (d1 + d2 + d3)  + ... + (d1 + d2 +...+ dn-1) = (n-1)d1 + (n-2)d2 + (n-3)d3 + ...+ dn-1
	// 第二个元素到后n-2个元素的距离和 可以写作：
	// d2 + (d2 + d3) + ... + (d2 + d3 + dn-1) = (n-2)d2 +(n-3)d3 + dn-1
	// 依次类推...
	// 把所有距离加和
	// sum = (n-1)d1 + 2*(n-2)d2 + 2*(n-3)d3 + ... + (n-1)*dn-1
	// sum = sum(i * (n-i) * di)
	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		res += (nums[i] - nums[i-1]) * i * (len(nums) - i)
	}
	return res

}

func sumOfDistanceEachTwoElem(nums []int) int {
	// 基于所有元素之间的距离公式2
	// 公式1  sum = sum(i * (n-i) * di) 是从每一个位置计算到之后位置的所有元素的距离和来推导的
	// 公式2 基于前缀和的形式  从每个位置计算当前元素到之前的元素的距离和
	// 当前第i个元素位置ai 当前元素到之前所有元素的距离和为 (ai - ai-1) + (ai - ai-2) + (ai -a1) = i * ai - (a1 + a2 +..an-1)
	// 逐个遍历元素位置 累计元素位置的前缀和

	ans := 0
	sort.Ints(nums)

	sum := 0 // 前缀和 循环中 表示位置i处的前缀和
	for i, x := range nums {
		ans = ans + i*x - sum
		sum += x
	}
	return ans
}
