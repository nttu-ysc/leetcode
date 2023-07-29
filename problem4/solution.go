package problem4

func findMidianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	var merge []int

	for len(nums1) > 0 || len(nums2) > 0 {
		if len(nums1) == 0 {
			n2 := nums2[0]
			nums2 = nums2[1:]
			merge = append(merge, n2)
			continue
		}

		if len(nums2) == 0 {
			n1 := nums1[0]
			nums1 = nums1[1:]
			merge = append(merge, n1)
			continue
		}

		n1 := nums1[0]
		nums1 = nums1[1:]
		n2 := nums2[0]
		nums2 = nums2[1:]

		if n1 < n2 {
			merge = append(merge, n1)
			nums2 = append([]int{n2}, nums2...)
		} else {
			merge = append(merge, n2)
			nums1 = append([]int{n1}, nums1...)
		}
	}

	mid := (m + n - 1) / 2
	if (m+n)%2 != 0 {
		return float64(merge[mid])
	}

	return float64(merge[mid]+merge[mid+1]) / 2
}
