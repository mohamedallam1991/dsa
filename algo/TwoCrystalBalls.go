package algo

import (
	"math"
)

func TwoCrystalBalls(array []int, target int) int {
	step := int(math.Round(math.Sqrt(float64(len(array)))))
	rbound := len(array)
	for i := step; i < len(array); i += step {
		if array[i] > target {
			rbound = i
			break
		}
	}

	for i := rbound - step; i < rbound; i++ {
		if array[i] == target {
			return array[i]
		}
		if array[i] > target {
			break
		}
	}
	return -1
}

func JumpBinarySearch(array []int, target int) int {
	step := int(math.Round(math.Sqrt(float64(len(array)))))
	var high int
	for i := step; i < len(array); i += step {
		if array[i] > target {
			high = i
			break
		}
	}
	low := high - step
	for low < high {
		mid := low + ((high - low) / 2)
		val := array[mid]
		if val == target {
			return val
		} else if val < target {
			low = mid + 1
		} else if val > target {
			high = mid
		}
	}
	return -1
}
