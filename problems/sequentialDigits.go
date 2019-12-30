package problems

import (
	"strconv"
)

var nums = "123456789"

func SequentialDigits(low,high int) []int{
	return sequentialDigits(low,high)
}

func sequentialDigits(low int, high int) []int {
	var result []int
	lowDigits :=  getDigitsFromInteger(low)
	highDigits := getDigitsFromInteger(high)
	for i := lowDigits; i <= highDigits; i++ {
		candidateValue := getsequentialDigitsInNDigits(i)
		for i := 0; i < len(candidateValue); i++ {
			if candidateValue[i] >= low && candidateValue[i] <= high {
				result = append(result, candidateValue[i])
			}
		}
	}
	return result
}

func getDigitsFromInteger(num int) int{
	result := 0
	for num != 0 {
		result += 1
		num = num / 10
	}
	return result
}

func getsequentialDigitsInNDigits(num int) []int {
	if num > 9 {
		return nil
	}
	var result []int
	for i := 0; i <= (9 - num); i++ {
		value, _ := strconv.Atoi(nums[i:i + num])
		result = append(result, value)
	}
	return result
}