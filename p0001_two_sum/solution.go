package p0001

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, num := range nums {
        if idx, found := seen[target-num]; found {
            return []int{idx, i}
        }
        seen[num] = i
    }
    return nil
}
