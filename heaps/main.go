package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

// get the length of the heap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

//Checks if i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

// Pushes the item to the heap
func (iheap *IntegerHeap) Push(heapInterface interface{}) {
	*iheap = append(*iheap, heapInterface.(int))
}

// Pops item from the front of the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

//Swaps the element of i index to j index
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	fmt.Println("the heap : ", *intHeap)
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Println("the heap with pushed value: ", *intHeap)

	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	fmt.Println()
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
		fmt.Println("the heap after pop interation: ", *intHeap)
	}
}
