package leetcode

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

// 输入：s = "([)]"
// 输出：false

// s = "{[]}"
// s = "()[]{}"

// 栈
func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return false
	}
	if num := length % 2; num != 0 {
		return false
	}
	var (
		leftMap = map[rune]struct{}{
			'(': {},
			'{': {},
			'[': {},
		}
		rightMap = map[rune]rune{
			')': '(',
			'}': '{',
			']': '[',
		}
		tempArr = make([]rune, 0)
		cur     int
	)

	for _, v := range s {
		if _, exist := leftMap[v]; exist {
			tempArr = append(tempArr, v)
			cur++
		}
		if left, exist := rightMap[v]; exist {
			if len(tempArr) == 0 {
				return false
			}
			if tempArr[cur-1] != left {
				return false
			}
			cur--
			tempArr = tempArr[:cur]
		}
	}

	return len(tempArr) == 0
}
