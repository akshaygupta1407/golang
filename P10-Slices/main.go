package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// useful methods
func main() {
	//uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums1 = make([]int, 2)
	fmt.Println(nums1)
	fmt.Println(cap(nums1))

	var nums2 = make([]int, 3)
	fmt.Println(nums2)
	fmt.Println(cap(nums2))

	var nums3 = make([]int, 3, 5)
	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	nums3 = append(nums3, 1)
	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	nums3 = append(nums3, 2)
	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	nums3 = append(nums3, 3)
	fmt.Println(nums3)
	fmt.Println(cap(nums3))

	nums4 := []int{}
	fmt.Println(nums4 == nil) //false
	fmt.Println(nums4)        //[]
	fmt.Println(len(nums4))   //0
	fmt.Println(cap(nums4))   //0

	//copy function
	nums5 := make([]int, 0, 5)
	nums5 = append(nums5, 1)
	nums6 := make([]int, len(nums5), 5)
	fmt.Println(nums5, nums6)
	copy(nums6, nums5)
	fmt.Println(nums5, nums6)

	//slice operator
	nums7 := []int{1, 2, 3}
	fmt.Println(nums7[0:2]) // [1 2]
	fmt.Println(nums7[1:])  //[2 3]

	//comparator
	nums8 := []int{1, 2}
	nums9 := []int{1, 2}
	fmt.Println(slices.Equal(nums8, nums9)) //true

	//2D slices
	nums10 := [][]int{{1, 2}, {3, 4}}
	fmt.Println(nums10) 	//[[1 2] [3 4]]
}
