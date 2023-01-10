package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(10, 7)
	if err != nil {
		t.Fatalf("expected 17 but got %v", err)
	}

	require.Equal(t, answer, int32(17))
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Subtraction(10, 7)
	if err != nil {
		t.Fatalf("expected 3 but got %v", err)
	}

	require.Equal(t, answer, int32(3))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(10, 7)
	if err != nil {
		t.Fatalf("expected 70 but got %v", err)
	}

	require.Equal(t, answer, int32(70))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(10, 7)
	if err != nil {
		t.Fatalf("expected 1 but got %v", err)
	}

	require.Equal(t, answer, int32(1))
}
