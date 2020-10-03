package main

import "fmt"

// 给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。

// 注意：1 ≤ k ≤ n ≤ 109。

// 示例 :

// 输入:
// n: 13   k: 2

// 输出:
// 10

// 解释:
// 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findKthNumber(n int, k int) int {
	return 0
}

func getCount(n, i int) int {
	lo := i
	hi := i + 1

}

func main() {
	n := 10
	k := 3
	ans := findKthNumber(n, k)
	fmt.Printf("ans=%d\n", ans)
}
