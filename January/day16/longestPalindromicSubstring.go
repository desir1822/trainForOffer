package day16

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	//for循环的目的 是让s的每个字节s[i]轮流当center位
	for i := 0; i < len(s); i++ {
		//应对两种回文字串，一种是 "baac" 偶数 一种是"bab” 奇数
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		//更新start, end的值，每次都更新 这样保证求出 最长的回文字串
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	//end+1 是因为 s[i:j]表示的是 s[i]~s[j-1]
	return s[start : end+1]
}

//中心向外扩张，保证
func expandAroundCenter(s string, left, right int) (int, int) {
	//for-clause: initStmt;conditionStmt;postStmt
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		//满足条件执行Block 并且执行postStmt;
		//不满足条件 退出for循环  不执行Block也不执行postStmt
	}
	//https://github.com/desir1822/learingGoLan/blob/main/goSpecification/12Stmt.md
	//9For statements for迭代
	//
	/*	这种写法时间超出限制:因为这样一定会执行Block，如果把if条件放在 conditionStmt，会减少Block的执行次数。上面写法不满足条件会跳过Block执行
		for ; left>=0&&right<len(s); {
			if s[left]==s[right]{
				left=left-1
				right=right+1
			}
		}
	*/

	return left + 1, right - 1 //return的是 最后一组 两两相等的 index组
}
