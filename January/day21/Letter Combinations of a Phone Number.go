package day21

/*给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
方法一：回溯
首先使用哈希表存储每个数字对应的所有可能的字母，然后进行回溯操作。[1map数据类型]
回溯过程中维护一个字符串，表示已有的字母排列（如果未遍历完电话号码的所有数字，则已有的字母排列是不完整的）。该字符串初始为空。[2]

每次取电话号码的一位数字，从哈希表中获得该数字对应的所有可能的字母，并将其中的一个字母插入到已有的字母排列后面，[3字符串拼接上一次的字母排列]
然后继续处理电话号码的后一位数字，直到处理完电话号码中的所有数字，即得到一个完整的字母排列。[4递归同时处理下一次，循环条件满足 返回结果]
然后进行回退操作，遍历其余的字母排列。[5各自返回得到结果数组]
回溯算法用于寻找所有的可行解，如果发现一个解不可行，则会舍弃不可行的解。在这道题中，由于每个数字对应的每个字母都可能进入字母组合，因此不存在不可行的解，直接穷举所有的解即可。
*/
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "") //2维护字符串列表 ，初始为空
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination) //回退 返回所有完整的字母排列
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i])) //3拼接上一次维护的字符串列表  combination+=string(letters[i])
			//4处理下一次
		}
	}
}
