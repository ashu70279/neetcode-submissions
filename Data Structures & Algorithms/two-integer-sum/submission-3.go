func twoSum(nums []int, target int) []int {
    elementMap := make(map[int][]int)
	ans := []int{}
	for i:=0; i<len(nums); i++{
		 elementMap[nums[i]] = append(elementMap[nums[i]],i)
	}

    for i:=0;i<len(nums);i++{
		val,exist := elementMap[target-nums[i]]; if exist {
            ans = append(ans,i)
			for j:=0;j<len(val);j++{
				if val[j]>i{
					ans = append(ans,val[j])
					break
				}
			}
			if len(ans) == 2 {
			return ans
			} else {
				ans = []int{}
			}
		}
	}
	return []int{0,0}

}
