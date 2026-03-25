package main

import "fmt"

//numbered sequence of specific length
func main() {
	//zeroed values
	var nums [4]int 
	fmt.Println("length of nums ", len(nums))
	fmt.Println(nums[0])
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	fmt.Println(nums)

	var names [3]string
	names[1] = "golang"
	fmt.Println(names)

	//to declare in single line

	nums1 := [4]int{0, 1, 2, 3}
	fmt.Println(nums1)

	nums3 := [...]int{10, 20, 30}
	fmt.Println("nums3 len", len(nums3))

	//2d array
	nums2 := [4][2]int{{0,1}, {2,3}, {4,5}, {6,7}}
	fmt.Println(nums2)

	//arrays used for fixed size, memory optimization, constant time access
}