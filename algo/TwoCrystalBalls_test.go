package algo

import (
	"log"
	"testing"
)

func AddIntToArray(a []int, n int) []int {
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	return a
}

func TestTwoCrystalBalls(t *testing.T) {
	a := []int{}
	a = AddIntToArray(a, 10000)

	needle := a[767]

	if TwoCrystalBalls(a, needle) != needle {
		log.Fatal("Its not wirking")
	}
	if JumpBinarySearch(a, needle) != needle {
		log.Fatal("Its not wirking")
	}

	needle = a[767]
	a[767] = 0
	if TwoCrystalBalls(a, needle) != -1 {
		log.Fatal("Its not wirking")
	}
	if JumpBinarySearch(a, needle) != -1 {
		log.Fatal("Its not wirking")
	}

}
