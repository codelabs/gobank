package applns

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestOperatorPrecedence(t *testing.T) {

	t.Run("operator #", func(t *testing.T) {
		assert.Equal(t, 1, operatorPrecedence("#"))
	})
	t.Run("operator +", func(t *testing.T) {
		assert.Equal(t, 2, operatorPrecedence("+"))
	})
	t.Run("operator -", func(t *testing.T) {
		assert.Equal(t, 2, operatorPrecedence("-"))
	})
	t.Run("operator *", func(t *testing.T) {
		assert.Equal(t, 3, operatorPrecedence("*"))
	})
	t.Run("operator /", func(t *testing.T) {
		assert.Equal(t, 3, operatorPrecedence("/"))
	})
	t.Run("operator (", func(t *testing.T) {
		assert.Equal(t, 4, operatorPrecedence("("))
	})
	t.Run("operator )", func(t *testing.T) {
		assert.Equal(t, 4, operatorPrecedence(")"))
	})
	t.Run("operator ^", func(t *testing.T) {
		assert.Equal(t, 4, operatorPrecedence("^"))
	})
}

func TestIsOperator(t *testing.T) {
	t.Run("operator +", func(t *testing.T) {
		assert.Equal(t, true, isOperator("+"))
	})
	t.Run("operator -", func(t *testing.T) {
		assert.Equal(t, true, isOperator("-"))
	})
	t.Run("operator *", func(t *testing.T) {
		assert.Equal(t, true, isOperator("*"))
	})
	t.Run("operator /", func(t *testing.T) {
		assert.Equal(t, true, isOperator("/"))
	})
	t.Run("operator ^", func(t *testing.T) {
		assert.Equal(t, true, isOperator("^"))
	})
	t.Run("operator (", func(t *testing.T) {
		assert.Equal(t, true, isOperator("("))
	})
	t.Run("operator )", func(t *testing.T) {
		assert.Equal(t, true, isOperator(")"))
	})
	t.Run("operator #", func(t *testing.T) {
		assert.Equal(t, true, isOperator("#"))
	})
	t.Run("operator %", func(t *testing.T) {
		assert.Equal(t, false, isOperator("%"))
	})
}

func TestConvertToPostFix(t *testing.T) {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	t.Run("infix expression - (B+C)", func(t *testing.T) {
		postfix := ConvertToPostFix("(B+C)")
		assert.Equal(t, "BC+", postfix)
	})

	t.Run("infix expression - A*B+C", func(t *testing.T) {
		postfix := ConvertToPostFix("A*B+C")
		assert.Equal(t, "AB*C+", postfix)
	})

	t.Run("infix expression - A+B*C", func(t *testing.T) {
		postfix := ConvertToPostFix("A+B*C")
		assert.Equal(t, "ABC*+", postfix)
	})

	t.Run("infix expression - A*(B+C)", func(t *testing.T) {
		postfix := ConvertToPostFix("A*(B+C)")
		assert.Equal(t, "ABC+*", postfix)
	})

	t.Run("infix expression - A*B^C+D", func(t *testing.T) {
		postfix := ConvertToPostFix("A*B^C+D")
		assert.Equal(t, "ABC^*D+", postfix)
	})

	t.Run("infix expression - A*(B+C*D)+E", func(t *testing.T) {
		postfix := ConvertToPostFix("A*(B+C*D)+E")
		assert.Equal(t, "ABCD*+*E+", postfix)
	})

	t.Run("infix expression - A*(B+C/(D-E))", func(t *testing.T) {
		postfix := ConvertToPostFix("A*(B+C/(D-E))")
		assert.Equal(t, "ABCDE-/+*", postfix)
	})
}
