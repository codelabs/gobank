package applns

import (
	"github.com/codelabs/gobank/lib/v1/adt"
	"strings"
)

type Operator string

const (
	plus             Operator = "+"
	minus            Operator = "-"
	multiplication   Operator = "*"
	division         Operator = "/"
	exponent         Operator = "^"
	openParenthesis  Operator = "("
	closeParenthesis Operator = ")"
	pound            Operator = "#"
)

// operatorPrecedence returns the precedence value for an Operator.
func operatorPrecedence(op string) int {

	var precedence int

	switch Operator(op) {
	case pound:
		precedence = 1
	case plus, minus:
		precedence = 2
	case multiplication, division:
		precedence = 3
	case openParenthesis, closeParenthesis, exponent:
		precedence = 4
	}

	return precedence
}

// isOperator returns true if selected symbol is one of the Operator, otherwise false.
func isOperator(symbol string) bool {
	switch Operator(symbol) {
	case plus, minus, multiplication, division, exponent, openParenthesis, closeParenthesis, pound:
		return true
	}
	return false
}

// ConvertToPostFix converts the infix expression to postfix expression and returns it.
func ConvertToPostFix(infix string) string {

	postfix := make([]string, 0)
	input := strings.Split(infix, "")
	stack := adt.NewStack[string]()

	for _, symbol := range input {

		// Push to stack if symbol is (
		if symbol == string(openParenthesis) {
			stack.Push(symbol)
			continue
		}

		// If symbol is ), pop from stack until we encounter (.
		// Append popped symbol to postfix string, ignoring both ( and ).
		if symbol == string(closeParenthesis) {
			for {
				if stack.IsEmpty() {
					break
				}
				operator, _ := stack.Pop()
				if operator == string(openParenthesis) {
					break
				}
				postfix = append(postfix, operator)
			}
			continue
		}

		// If symbol is not operator, just append to postfix string
		if !isOperator(symbol) {
			postfix = append(postfix, symbol)
			continue
		}

		// If stack is empty, just push to stack
		if stack.IsEmpty() {
			stack.Push(symbol)
			continue
		}

		// If top of stack has higher precedence than current symbol,
		// pop from stack and append to postfix string.
		// Repeat this until top of stack has lesser precedence or till stack becomes empty.
		for {

			if stack.IsEmpty() {
				break
			}

			top := stack.Top()

			if top != "(" && operatorPrecedence(top) > operatorPrecedence(symbol) {
				operator, _ := stack.Pop()
				postfix = append(postfix, operator)
			} else {
				break
			}
		}

		// if top of stack has lesser precedence than symbol, push to stack
		stack.Push(symbol)
	}

	// When all the symbols from input are processed, whatever left in
	// the stack pop it and append to postfix string
	for !stack.IsEmpty() {
		operator, _ := stack.Pop()
		postfix = append(postfix, operator)
	}

	return strings.Join(postfix, "")
}
