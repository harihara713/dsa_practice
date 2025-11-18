package stack

import (
	"fmt"
	"unicode"
)

func InfixToPostfix(exp string) string {
	ans := ""
	s := NewStack()

	for _, r := range exp {
		// if the rune is the operand then add it to the ans
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			ans += string(r)
		} else if r == '(' {
			s.Push(r)
		} else if r == ')' {
			// pop out all the rune till it finds a opening bracket
			for !s.Empty() && s.Peek() != '(' {
				ans += string(s.Pop().(rune))
			}
			// pop (
			s.Pop()
		} else {
			// the operators
			// if the stack is not empty and priority of rune r is less than the top elements priority then pop and add to answer till
			// we don't find a top whose priority is less than rune r or stack is empty
			for !s.Empty() && priority(r) <= priority(s.Peek().(rune)) {
				ans += string(s.Pop().(rune))
			}
			// push to stack
			s.Push(r)
		}
	}

	// if stack has elements
	for !s.Empty() {
		ans += string(s.Pop().(rune))
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
	st := NewStack()

	exp = reverseExpression(exp)

	for _, r := range exp {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			ans += string(r)
		} else if r == '(' {
			st.Push(r)
		} else if r == ')' {
			// pop out all the rune till it finds a opening bracket
			for !st.Empty() && st.Peek() != '(' {
				ans += string(st.Pop().(rune))
			}
			// pop (
			st.Pop()
		} else {
			// the operators
			// if the stack is not empty and priority of rune r is less than the top elements priority then pop and add to answer till
			// we don't find a top whose priority is less than rune r or stack is empty
			if r == '^' {
				for !st.Empty() && priority(r) <= priority(st.Peek().(rune)) {
					ans += string(st.Pop().(rune))
				}
			} else {
				for !st.Empty() && priority(r) < priority(st.Peek().(rune)) {
					ans += string(st.Pop().(rune))
				}
			}
			// push to stack
			st.Push(r)
		}
	}

	// if stack has elements
	for !st.Empty() {
		ans += string(st.Pop().(rune))
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
	st := NewStack()

	for _, r := range exp {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			st.Push(string(r))
		} else {
			// take out two element
			first := st.Pop()
			second := st.Pop()
			tf := fmt.Sprintf("(%s%s%s)", second, string(r), first)
			st.Push(tf)
		}
	}

	return st.Pop().(string)
}

func PrefixToInfix(exp string) string {
	st := NewStack()
	// start from last, reverse order
	for i := len(exp) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(exp[i])) || unicode.IsDigit(rune(exp[i])) {
			st.Push(string(rune(exp[i])))
		} else {
			first := st.Pop()
			second := st.Pop()
			tf := fmt.Sprintf("(%s%s%s)", first, string(rune(exp[i])), second)
			st.Push(tf)
		}
	}

	return st.Pop().(string)
}

func PostfixToPrefix(exp string) string {
	st := NewStack()
	for _, r := range exp {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			st.Push(string(r))
		} else {
			// take out two
			f := st.Pop()
			s := st.Pop()
			tf := fmt.Sprintf("%s%s%s", string(r), s, f)
			st.Push(tf)
		}
	}

	return st.Pop().(string)
}

func PrefixToPostfix(exp string) string {
	st := NewStack()
	// start from last, reverse order
	for i := len(exp) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(exp[i])) || unicode.IsDigit(rune(exp[i])) {
			st.Push(string(rune(exp[i])))
		} else {
			first := st.Pop()
			second := st.Pop()
			tf := fmt.Sprintf("%s%s%s", first, second, string(rune(exp[i])))
			st.Push(tf)
		}
	}

	return st.Pop().(string)
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
