package validParentheses

func ft_search(s string, search uint8, i int) int {
	for i < len(s) {
		//Если снова попался открывающий символ, рекурсивно вызываю функцию
		switch s[i] {
		case '(':
			i = ft_search(s, ')', i+1)
			if i == 0 {
				return 0
			}
		case '[':
			i = ft_search(s, ']', i+1)
			if i == 0 {
				return 0
			}
		case '{':
			i = ft_search(s, '}', i+1)
			if i == 0 {
				return 0
			}
		}
		//Если искомый символ найден, возвращаю символ следующий за ним (или последний)
		if s[i] == search {
			if i == len(s)-1 {
				return i
			}
			return i + 1
		}
		//Если наткнулись на закрывающий символ отличный от искомого, то возвращаю 0
		if search == ')' && (s[i] == ']' || s[i] == '}') {
			return 0
		}
		if search == ']' && (s[i] == ')' || s[i] == '}') {
			return 0
		}
		if search == '}' && (s[i] == ']' || s[i] == ')') {
			return 0
		}
	}
	return i
}

func isValid(s string) bool {
	//Если строка пустая,то true
	if len(s) == 0 {
		return true
	}
	k := 0
	j := 0
	for i := range s {
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			k++
		}
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			j++
		}
	}
	//Если количество открывающих и закрывающих скобок различно, то false
	if j != k {
		return false
	}
	for i := range s {
		//В зависимости от открывающего символа, начинаю искать закрывающий символ
		switch s[i] {
		case '(':
			i = ft_search(s, ')', i+1)
			if i == 0 {
				return false
			}
		case '[':
			i = ft_search(s, ']', i+1)
			if i == 0 {
				return false
			}
		case '{':
			i = ft_search(s, '}', i+1)
			if i == 0 {
				return false
			}
		}
		//Если дошли до последнего символа строки, то true
		if i == len(s)-1 {
			return true
		}
	}
	return false
}
