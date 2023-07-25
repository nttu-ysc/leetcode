package problem852

func peakIndexInMountainArray(arr []int) int {
	var l, r = 0, len(arr)
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
