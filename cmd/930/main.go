package main

func main() {
}

func numSubarraysWithSum(nums []int, goal int) int {
	countMap := map[int]int{}
	countMap[0] = 1
	count := 0
	sum := 0

	for _, val := range nums {
		sum += val
		remain := sum - goal
		if num, ok := countMap[remain]; ok {
			count += num
		}
		countMap[sum]++
	}

	return count
}
