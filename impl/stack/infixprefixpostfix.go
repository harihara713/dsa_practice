package stack

import (
	"fmt"
	"unicode"
)

func InfixToPostfix(exp string) string {
	ans := ""
	s := NewStack[rune]()

	for _, r := range exp {
		// if the rune is the operand then add it to the ans
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			ans += string(r)
			continue
		}

		if r == '(' {
			s.Push(r)
			continue
		}

		if r == ')' {
			// pop out all the rune till it finds a opening bracket
			for !s.Empty() {
				top, _ := s.Peek()
				if top == '(' {
					break
				}
				val, _ := s.Pop()
				ans += string(val)
			}
			// pop (
			s.Pop()
			continue
		}

		// the operators
		// if the stack is not empty and priority of rune r is less than the top elements priority then pop and add to answer till
		// we don't find a top whose priority is less than rune r or stack is empty
		for !s.Empty() {
			top, _ := s.Peek()
			if priority(top) < priority(r) {
				break
			}
			val, _ := s.Pop()
			ans += string(val)
		}
		// push to stack
		s.Push(r)

	}

	// if stack has elements
	for !s.Empty() {
		val, _ := s.Pop()
		ans += string(val)
	}

	return ans
}

func InfixToPrefix(exp string) string {
	/*
		1. reverse and change ( -> to ->) and vice-versa
		2. transform to postfix with a tweak in condition (pop out of stack when priority of current rune strictly less than priority
			of top element of stack unless rune r == '^')
		3. reverse the answer
	*/
	ans := ""
	st := NewStack[rune]()

	exp = reverseExpression(exp)

	for _, r := range exp {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			ans += string(r)
			continue
		}

		if r == '(' {
			st.Push(r)
			continue
		}

		if r == ')' {
			// pop out all the rune till it finds a opening bracket
			for !st.Empty() {
				top, _ := st.Peek()
				if top == '(' {
					break
				}
				val, _ := st.Pop()
				ans += string(val)
			}
			// pop (
			st.Pop()
			continue
		}
		// the operators
		// if the stack is not empty and priority of rune r is less than the top elements priority then pop and add to answer till
		// we don't find a top whose priority is less than rune r or stack is empty
		if r == '^' {
			for !st.Empty() {
				top, _ := st.Peek()
				if priority(r) > priority(top) {
					break
				}
				val, _ := st.Pop()
				ans += string(val)
			}
		} else {
			for !st.Empty() {
				top, _ := st.Peek()
				if priority(r) >= priority(top) {
					break
				}
				val, _ := st.Pop()
				ans += string(val)
			}
		}
		// push to stack
		st.Push(r)

	}

	// if stack has elements
	for !st.Empty() {
		val, _ := st.Pop()
		ans += string(val)
	}

	ans = reverseExpression(ans)

	return ans
}

func PostfixToInfix(exp string) string {
	/*
		1. when encounter an operand push it into stack
		2. if we encounter an operator then take out two elements from stack then put the operator in between them and wrap it with parenthesis and
			push back to the stack
	*/
	st := NewStack[string]()

	for _, r := range exp {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			st.Push(string(r))
		} else {
			// take out two element
			first, _ := st.Pop()
			second, _ := st.Pop()
			tf := fmt.Sprintf("(%s%s%s)", second, string(r), first)
			st.Push(tf)
		}
	}

	val, _ := st.Pop()
	return val
}

func PrefixToInfix(exp string) string {
	st := NewStack[string]()
	// start from last, reverse order
	for i := len(exp) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(exp[i])) || unicode.IsDigit(rune(exp[i])) {
			st.Push(string(rune(exp[i])))
		} else {
			first, _ := st.Pop()
			second, _ := st.Pop()
			tf := fmt.Sprintf("(%s%s%s)", first, string(rune(exp[i])), second)
			st.Push(tf)
		}
	}

	val, _ := st.Pop()
	return val
}

func PostfixToPrefix(exp string) string {
	st := NewStack[string]()
	for _, r := range exp {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			st.Push(string(r))
		} else {
			// take out two
			first, _ := st.Pop()
			second, _ := st.Pop()
			tf := fmt.Sprintf("%s%s%s", string(r), second, first)
			st.Push(tf)
		}
	}

	val, _ := st.Pop()
	return val
}

func PrefixToPostfix(exp string) string {
	st := NewStack[string]()
	// start from last, reverse order
	for i := len(exp) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(exp[i])) || unicode.IsDigit(rune(exp[i])) {
			st.Push(string(rune(exp[i])))
		} else {
			first, _ := st.Pop()
			second, _ := st.Pop()
			tf := fmt.Sprintf("%s%s%s", first, second, string(rune(exp[i])))
			st.Push(tf)
		}
	}

	val, _ := st.Pop()
	return val
}

func priority(operator rune) int {
	switch operator {
	case '^':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	default:
		return -1
	}
}

func reverseExpression(exp string) string {
	runes := []rune(exp)
	i, j := 0, len(runes)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	// change (->to->) and vice-versa
	for i = 0; i < len(runes); i++ {
		switch runes[i] {
		case ')':
			runes[i] = '('
		case '(':
			runes[i] = ')'
		}
	}

	return string(runes)
}
