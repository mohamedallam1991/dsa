package algo

func BinarySearcha(haystack []int, needle int) bool {
	var low, high, mid, val int
	low = 0
	high = len(haystack)

	for low <= high {
		mid = low + ((high - low) / 2)
		val = haystack[mid]
		if val == needle {
			return true
		} else if val < needle {
			low = mid + 1
		} else if val > needle {
			high = mid - 1
		}
	}
	return false
}

func BinarySearch(haystack []int, needle int) bool {
	var low, high, mid, val int
	low = 0
	high = len(haystack)

	for low < high {
		mid = low + ((high - low) / 2)
		val = haystack[mid]
		if val == needle {
			return true
		} else if val < needle {
			low = mid + 1
		} else if val > needle {
			high = mid
		}
	}
	return false
}
