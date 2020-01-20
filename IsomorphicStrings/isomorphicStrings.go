package isomorphicStrings

func isIsomorphic(s string, t string) bool {
	//Создание хеш-таблицы под символы
	dictionary := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		//Если такого символа нет в таблице и нет такого значения, как в данном символе,
		//то в хеш-таблицу записывается этот символ с данным значением. Иначе возвращается
		//false
		if _, ok := dictionary[s[i]]; !ok {
			for _, value := range dictionary {
				if value == t[i] {
					return false
				}
			}
			dictionary[s[i]] = t[i]
		} else {
			//Если такой символ уже есть в таблице, но у него другое значение, возвращается
			//false
			if dictionary[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}
