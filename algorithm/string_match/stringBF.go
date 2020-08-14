package stringmatch

//BFMatch bf字符串匹配算法
func BFMatch(main, pattern string) bool {
	for i := 0; i < len(main); i++ {
		for j := 0; j < len(pattern); j++ {
			if pattern[j] != main[i+j] {
				break
			}
			if j == len(pattern)-1 {
				return true
			}
		}
	}
	return false
}
