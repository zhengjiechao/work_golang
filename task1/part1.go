package task1

import (
	"sort"
	"strconv"
	"strings"
)

// 1.只出现一次的数字
func SingleNumber(nums []int) int {
	len := len(nums)
	if len == 1 {
		return nums[0]
	}
	countMap := make(map[int]int, len/2+1)
	for _, v := range nums {
		countMap[v]++
	}

	for k, v := range countMap {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 2.回文数
func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	strByte := []byte(str)

	len := len(strByte)
	for i := 0; i <= len/2; i++ {
		if strByte[i] != strByte[len-i-1] {
			return false
		}
	}
	return true
}

// 3.有效的括号
func IsValid(s string) bool {
	lastStr, newStr := s, s
	for {
		newStr = strings.ReplaceAll(newStr, "()", "")
		newStr = strings.ReplaceAll(newStr, "{}", "")
		newStr = strings.ReplaceAll(newStr, "[]", "")
		if lastStr == newStr {
			break
		} else {
			lastStr = newStr
		}
	}
	return lastStr == ""
}

// 4.最长公共前缀
func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}

	str := strs[0]
	for i := 1; i < length; i++ {
		for {
			if strings.HasPrefix(strs[i], str) {
				break
			} else {
				strRune := []rune(str)
				str = string(strRune[:len(strRune)-1])
			}
		}
		if str == "" {
			return str
		}
	}
	return str
}

// 5.加一
func PlusOne(digits []int) []int {
	length := len(digits)
	digits[length-1] = digits[length-1] + 1
	for i := length - 1; i >= 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				newSilece := []int{1}
				return append(newSilece, digits...)
			} else {
				digits[i-1] = digits[i-1] + 1
			}
		}
	}
	return digits
}

// 6.删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	for i := 1; i < length; i++ {
		if nums[i-1] == nums[i] {
			nums = append(nums[0:i], nums[(i+1):length]...)
			length--
			i--
		}
	}
	return len(nums)
}

// 7.合并区间
func Merge(intervals [][]int) [][]int {
	// 先进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	length := len(intervals)
	returnSlice := make([][]int, 0)
	for i := 0; i < length; i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		for j := i + 1; j < length; j++ {
			if intervals[j][0] > end {
				break
			} else {
				i = j
				if intervals[j][1] > end {
					end = intervals[j][1]
				}
			}
		}
		itemSlice := []int{start, end}
		returnSlice = append(returnSlice, itemSlice)
	}
	return returnSlice
}

// 8.两数之和
func TwoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return make([]int, 0)
}
