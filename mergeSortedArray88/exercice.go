package mergeSortedArray88

import "fmt"

func Run() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}

	mergeSortedArray(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}

	if m == 0 {
		return
	}

	counter := 0

	for i := 0; i < n+m || m+counter < n+m; i++ {
		if nums1[i] > nums1[m+counter] {
			nums1[i], nums1[m+counter] = nums1[m+counter], nums1[i]
		}

		if nums1[m+counter] >= nums1[m+counter-1] {
			counter++
		}
	}
}
