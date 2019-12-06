package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	currIndex := len(nums1) - 1
	n--
	m--

	for m >= 0 || n >= 0 {
		if n < 0 {
			nums1[currIndex] = nums1[m]
			m--
		} else if m < 0 {
			nums1[currIndex] = nums2[n]
			n--
		} else if nums1[m] > nums2[n] {
			nums1[currIndex] = nums1[m]
			m--
		} else {
			nums1[currIndex] = nums2[n]
			n--
		}

		currIndex--
	}

}

func main() {
	arr1 := []int{1,2,3,0,0,0}
	arr2 := []int{2,5,6}
	merge(arr1, 3, arr2, 3)
	fmt.Printf("merged %v\n", arr1)
}