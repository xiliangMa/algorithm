package leetcode

import (
	"regexp"
	"strings"
	"testing"
)

func IsPalindrome(s string) bool {
	re := regexp.MustCompile("[^0-9A-Za-z]")
	s = re.ReplaceAllString(s, "")
	arr := strings.Split(s, "")
	i := 0
	j := len(arr) - 1
	for i <= j {
		if !strings.EqualFold(arr[i], arr[j]) {
			return false
		}
		i += 1
		j -= 1
	}

	return true
}

func Test_IsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	t.Log("测试结果：", IsPalindrome(s))
}
