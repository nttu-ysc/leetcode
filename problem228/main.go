package problem228

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	tmp := []int{nums[0]}
	ans := make([]string, 0)

	for i := 1; i < len(nums); i++ {
		if len(tmp) > 0 {
			if nums[i]-tmp[len(tmp)-1] == 1 {
				tmp = append(tmp, nums[i])
				continue
			}

			if len(tmp) == 1 {
				ans = append(ans, strconv.Itoa(tmp[0]))
			} else {
				ans = append(ans, strconv.Itoa(tmp[0])+"->"+strconv.Itoa(tmp[len(tmp)-1]))
			}
			tmp = []int{}
		}

		tmp = append(tmp, nums[i])
	}

	if len(tmp) > 1 {
		ans = append(ans, strconv.Itoa(tmp[0])+"->"+strconv.Itoa(tmp[len(tmp)-1]))
	}

	if len(tmp) == 1 {
		ans = append(ans, strconv.Itoa(tmp[0]))
	}

	return ans
}
