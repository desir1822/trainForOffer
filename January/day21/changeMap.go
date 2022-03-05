package day21

var phoneMap1 map[rune]string = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}
var combinations1 []string

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "") //2维护字符串列表 ，初始为空
	return combinations
}
func backtrack2(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination) //回退 返回所有完整的字母排列
	} else {
		digit := digits[index] //byte
		letters := phoneMap1[rune(digit)]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ { //letter[i]是 取字符串字节 是byte类型
			backtrack(digits, index+1, combination+string(letters[i])) //3拼接上一次维护的字符串列表  combination+=string(letters[i])
			//4处理下一次
		}
	}
}
