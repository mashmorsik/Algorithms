package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := make([]int, 0, len(nums1)+len(nums2))
	arr = append(arr, nums1...)
	arr = append(arr, nums2...)

	if len(arr)%2 == 1 {
		return float64(arr[len(arr)/2])
	}

	return float64(arr[len(arr)/2-1]+arr[len(arr)/2]) / 2
}
