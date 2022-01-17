package main

import (
	"fmt"
	"strconv"
)

/*
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
示例 1：
输入：num1 = "11", num2 = "123"
输出："134"
示例 2：
输入：num1 = "456", num2 = "77"
输出："533"
示例 3：
输入：num1 = "0", num2 = "0"
输出："0"
提示：
    1 <= num1.length, num2.length <= 104
    num1 和num2 都只包含数字 0-9
*/
func addStrings(num1 string, num2 string) string {
	carry := 0//进位
	ans := ""
	for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || carry != 0; i, j = i - 1, j - 1 {
		var x, y int   //x,y分别代表 本次计算相加的两部分
		if i >= 0 {
			 x = int(num1[i] - 48)
			//通过debug  num1[i] type是byte,值是57  48是untyped constant值为48
			//         conversion: num1[i]-48           得到的是byte类型，值为9
			//                        byte-untyped constant-->byte
			//int({byte}9)  将byte9转为int 9
			//{byte}'9'-->{int}9   byte是uint8别名 uint8转int是一定可行的

		}
		if j >= 0 {
			y = int(num2[j] - '0') // '0'是untyped constant

		}
		result := x + y + carry
		ans = strconv.Itoa(result%10) + ans//字符串拼接
		carry = result / 10
	}
	return ans
}
func main(){
	s1:=addStrings("999","222")
	fmt.Println(s1)
}