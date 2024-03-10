package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	mapped := map[int]bool{}
	result := []int{}

	for _, val := range nums1 {
		mapped[val] = true
	}

	for _, val := range nums2 {
		if exists := mapped[val]; exists {
			result = append(result, val)
			mapped[val] = false
		}
	}

	return result
}
