package main

func main() {
}

func getCommon(nums1 []int, nums2 []int) int {
	mapped := map[int]bool{}
	for i := 0; i < len(nums1); i++ {
		mapped[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		if _, ok := mapped[nums2[i]]; ok {
			return nums2[i]
		}
	}

	return -1
}
