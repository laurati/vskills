package main

import "fmt"

// binary search recurssion
func binarySearch(array []int, target int, lowIndex int, highIndex int) int {

	if highIndex < lowIndex {
		return -1
	}

	mid := int((lowIndex + highIndex) / 2)

	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

// binary search
func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int

	for startIndex <= endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {

	sliceTest := []int{11, 2, 3, 4, 5, 6, 6, 10, 15, 17, 18, 19, 30}

	b := binarySearch(sliceTest, 5, 0, len(sliceTest)-1)
	fmt.Println("recursivo", b)

	iteB := iterBinarySearch(sliceTest, 5, 0, len(sliceTest)-1)
	fmt.Println("iterativo", iteB)

}
