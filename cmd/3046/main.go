package main

func main() {

}

func isPossibleToSplit(nums []int) bool {
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
		if frequency[num] > 2 {
			return false
		}
	}

	return true
}
