package arr

import (
	"math"
	"sort"
)

// Reverse 反转
func Reverse(arr []string) (res []string) {
	pushPlanLen := len(arr)
	val := int(math.Floor(float64(pushPlanLen / 2)))
	for i := 0; i < val; i++ {
		arr[i], arr[pushPlanLen-1-i] = arr[pushPlanLen-1-i], arr[i]
	}
	return arr
}

// Sort 冒泡排序
func Sort(arr []int) (res []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	res = arr
	return
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
