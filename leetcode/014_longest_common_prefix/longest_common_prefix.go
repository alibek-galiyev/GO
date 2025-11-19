package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	n := len(strs)
	m := len(strs[0])
	for i := range n {
		if len(strs[i]) < m {
			m = len(strs[i])
		}
	}
	for i := range m {
		letter := strs[0][i]
		for j := range n {
			if strs[j][i] != letter {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:m]
}
