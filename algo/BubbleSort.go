package algo

func Bubble(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		//1,2,3,4,5,6
		for j := 0; j < length-1-i; j++ {
			//0,1,2,3,5,6
			if array[j] > array[j+1] {
				a := array[j]
				array[j] = array[j+1]
				array[j+1] = a
			}
		}
	}
	return array

}
