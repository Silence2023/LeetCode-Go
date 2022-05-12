package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// 遍历数组，得到与target的差值，使用map保存对应的value和index位置
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		if _, ok := m[value]; ok {
			return []int{m[value], i}
		}
		m[nums[i]] = i
	}
	return nil
}
