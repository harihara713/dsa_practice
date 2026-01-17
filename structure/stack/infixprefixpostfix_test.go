package stack

import "testing"

func TestInfixToPostfix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "lower-case",
			expression: "(a+b)*(c+d)",
			expected:   "ab+cd+*",
		},
		{
			name:       "upper-case",
			expression: "(A+B)*C-D+F",
			expected:   "AB+C*D-F+",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post := InfixToPostfix(tt.expression)

			if post != tt.expected {
				t.Errorf("Expected %s, Got %s\n", tt.expected, post)
			}
		})
	}
}

func TestInfixToPrefix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "lower-case",
			expression: "(a+b)*(c+d)",
			expected:   "*+ab+cd",
		},
		{
			name:       "upper-case",
			expression: "(A+B)*C-D+F",
			expected:   "+-*+ABCDF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InfixToPrefix(tt.expression)

			if got != tt.expected {
				t.Errorf("Expected %s, Got %s\n", tt.expected, got)
			}
		})
	}
}

func TestReverseExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "lower-case",
			expression: "(a+b)*(c+d)",
			expected:   "(d+c)*(b+a)",
		},
		{
			name:       "upper-case",
			expression: "(A+B)*C-D+F",
			expected:   "F+D-C*(B+A)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseExpression(tt.expression)

			if got != tt.expected {
				t.Errorf("Expected %s, Got %s\n", tt.expected, got)
			}
		})
	}
}

func TestPostfixToInfix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "lower-case",
			expression: "ab+cd+*",
			expected:   "((a+b)*(c+d))",
		},
		{
			name:       "upper-case",
			expression: "AB+C*D-F+",
			expected:   "((((A+B)*C)-D)+F)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post := PostfixToInfix(tt.expression)

			if post != tt.expected {
				t.Errorf("Expected %s, Got %s\n", tt.expected, post)
			}
		})
	}
}

func TestPrefixToInfix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "lower-case",
			expression: "*+ab+cd",
			expected:   "((a+b)*(c+d))",
		},
		{
			name:       "upper-case",
			expression: "+-*+ABCDF",
			expected:   "((((A+B)*C)-D)+F)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrefixToInfix(tt.expression)

			if got != tt.expected {
				t.Errorf("Expected %s, Got %s\n", tt.expected, got)
			}
		})
	}
}

func TestPostfixToPrefix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "lower-case",
			expression: "ab+cd+*",
			expected:   "*+ab+cd",
		},
		{
			name:       "upper-case",
			expression: "AB+C*D-F+",
			expected:   "+-*+ABCDF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post := PostfixToPrefix(tt.expression)

			if post != tt.expected {
				t.Errorf("Expected %s, Got %s\n", tt.expected, post)
			}
		})
	}
}

func TestPrefixToPostfix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "lower-case",
			expression: "*+ab+cd",
			expected:   "ab+cd+*",
		},
		{
			name:       "upper-case",
			expression: "+-*+ABCDF",
			expected:   "AB+C*D-F+",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrefixToPostfix(tt.expression)

			if got != tt.expected {
				t.Errorf("Expected %s, Got %s\n", tt.expected, got)
			}
		})
	}
}
