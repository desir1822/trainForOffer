package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	//从字符串大小
	if n%2 == 1 {
		return false
	}
	//对应关系，区别于惯性 v用来 表示存在情况
	pairs := map[byte]byte{
		')': '(', //41->40
		']': '[', //93->91
		'}': '{', //125->123
	}
	//https://go.dev/ref/spec#Slice_expressions
	//思路就是 先入栈 ，然后判断下一个字节 是否和入栈元素 匹配，匹配直接进行切片s[0:len-1] length equal high-low，也就是说如果是有效字符串，栈内每次最多只有一个元素，一旦匹配就会丢弃已经匹配的元素。
	//比较 的是栈内最后一个 和当前循环 是否配对，配对入栈不配对就返回false
	stack := []byte{}        //初始化byte数组
	for i := 0; i < n; i++ { //iteration
		if pairs[s[i]] > 0 { //   比较 入栈的最后一个元素 和 下一次循环遍历到的s[i]的对应字节是否一致
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { //如果栈长度0 或 栈最后一个元素 不是 s[i]想对的元素
				return false
			}
			stack = stack[:len(stack)-1] //切片 stack 0~len-2
		} else {
			stack = append(stack, s[i]) //len尾部追加s[i]
		}
	}
	return len(stack) == 0
}
func main() {
	fmt.Println(isValid(")((()))"))
	fmt.Println(isValid("[]{}()"))
}
