package IsomorphicStrings

func IsIsomorphic(s string, t string) bool {
	dictionary := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := dictionary[s[i]]; !ok {
			for _, value := range dictionary {
				if value == t[i] {
					return false
				}
			}
			dictionary[s[i]] = t[i]
		} else {
			if dictionary[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}
