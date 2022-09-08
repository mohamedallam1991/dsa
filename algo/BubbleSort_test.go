package algo

import (
	"log"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBubble(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	get := BubbleSort(arr)
	want := []int{3, 4, 7, 9, 42, 69, 420}

	if cmp.Equal(want, get) == false {
		log.Fatal("some error with bubble sort")
	}

	if reflect.DeepEqual(want, get) == false {
		log.Fatal("some error with bubble sort")
	}
}
