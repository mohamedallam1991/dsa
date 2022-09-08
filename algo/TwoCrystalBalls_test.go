package algo

import (
	"log"
	"testing"
)

// test("two crystal balls", function () {
//     let idx = Math.floor(Math.random() * 10000);
//     const data = new Array(10000).fill(false);

//     for (let i = idx; i < 10000; ++i) {
//         data[i] = true;
//     }

//	    expect(two_crystal_balls(data)).toEqual(idx);
//	    expect(two_crystal_balls(new Array(821).fill(false))).toEqual(-1);
//	});
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
