func groupAnagrams(strs []string) [][]string {
    ans := [][]string{}
	for i:=0;i<len(strs);i++{
		ans = append(ans,[]string{})
		ans[i] = append(ans[i],strs[i])
		for j:=i+1;j<len(strs);j++{
			if IsAnagram(strs[i],strs[j]){
				ans[i] = append(ans[i],strs[j])
				strs = append(strs[:j], strs[j+1:]...)
				j--
			}
		}
	}
	return ans
}

func IsAnagram(first string, second string) bool{
	if len(first)!=len(second){
		return false
	}

	stringMap := make(map[byte]int)

	for i:=0;i<len(first);i++{
		stringMap[first[i]]++
	}

	for i:=0;i<len(second);i++{
		stringMap[second[i]]--
		if stringMap[second[i]] < 0{
			return false
		}
	}
	return true


}
