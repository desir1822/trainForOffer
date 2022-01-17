package main

import "fmt"

/*
3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without repeating characters.
Constraints:
    0 <= s.length <= 5 * 104
    s consists of English letters, digits, symbols and spaces.
*/
//delete(m, k)  // remove element m[k] from map m
/*
滑动窗口
什么是滑动窗口？
其实就是一个队列,比如例题中的 abcabcbb，进入这个队列（窗口）为 abc 满足题目要求，当再进入 a，队列变成了 abca，这时候不满足要求。所以，我们要移动这个队列！
如何移动？我们只要把队列的左边的元素移出就行了，直到满足题目要求！一直维持这样的队列，找出队列出现最长的长度时候，求出解！
abca 先进入a然后把最左边那个a去掉，同时b变成字串开始的坐标
set的value当作是证明存在的信号量，不存在为zero-value，存在的话随便赋一个值就行，区别于zero-value

首先循环决定 字串的起始index
每次循环中 需要 往 map加入元素，记录字串右边界index。
发现重复元素在本次循环不会立即添加，而是在下一次循环中 先删除 上一轮循环起始index对应的 然后本次循环中加入map，
*/
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}//参考learingGoLan/goSpec/Properties of types and values/Reprenstability
	//4lexcial element介绍rune字面量代表一个integer value，所以这两互换。stringLit没有介绍string可以代表interger value.所以不可互换。

	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		/*           rk+1<n防止 index out of range */
		for rk + 1 < n && m[s[rk+1]] == 0 {//==0代表是int zero-value 没有进入过set
			// 不断地移动右指针
			m[s[rk+1]]++   //每一个进入set的Byte对应的V==1 区别于zero-value就行
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk - i + 1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func main(){
	fmt.Println(lengthOfLongestSubstring("abbcdabbb"))
}

/*
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/
滑动窗口题目:
3. 无重复字符的最长子串
30. 串联所有单词的子串
76. 最小覆盖子串
159. 至多包含两个不同字符的最长子串
209. 长度最小的子数组
239. 滑动窗口最大值
567. 字符串的排列
632. 最小区间
727. 最小窗口子序列
*/