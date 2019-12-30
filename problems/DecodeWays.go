package problems

import (
	"strconv"
)

func NumDecodings(s string) int {
	return numDecodings(s)
}

func numDecodings(s string) int {
	if len(s) == 0 || string(string(s)[0:1]) == "0"{
		return 0
	} else if len(s) == 1 {
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0
		} else {
			if num > 0 && num < 10 {
				return 1
			} else {
				return 0
			}
		}
	} else if len(s) == 2 {
		num2, err := strconv.Atoi(s)
		if err != nil {
			return 0
		} else { //说明字符串完全可以转换成正常的数字
			numfirst, _:= strconv.Atoi(string(string(s)[0:1]))
			numSecond, _:= strconv.Atoi(string(string(s)[1:2]))
			if num2 > 0 { //两个数字需要检验
				if numSecond > 0 && numfirst > 0{
					if num2 > 26 {
						return 1
					}
					return 2
				} else if (numfirst == 0 && numSecond > 0) || (numfirst > 0 && numSecond == 0 && numfirst < 3) {
					return 1
				} else {
					return 0
				}
			}
			return 0
		}
	} else  if len(s) > 2 {// 字符串会很长，不能直接转string
		num2, _ := strconv.Atoi(string(string(s)[0:2]))
		result1 := numDecodings(string(string(s)[1:]))
		result2 := 0
		if num2 < 27 && num2 > 0 { //两个数字需要检验
			result2 = numDecodings(string(string(s)[2:]))
		}
		return result1 + result2
	}
	return 0
}


