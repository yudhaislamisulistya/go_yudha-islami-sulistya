package calculates

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
	if Add(1, 1) != 2 {
		t.Error("1 + 1 did not equal 2")
	}
}

func TestSub(t *testing.T) {
	if Sub(2, 1) != 1 {
		t.Error("2 - 1 did not equal 1")
	}
	if Sub(1, 1) != 0 {
		t.Error("1 - 1 did not equal 0")
	}
}

func TestMul(t *testing.T) {
	if Mul(2, 2) != 4 {
		t.Error("2 * 2 did not equal 4")
	}
	if Mul(2, 1) != 2 {
		t.Error("2 * 1 did not equal 2")
	}
}

func TestDiv(t *testing.T) {
	if Div(4, 2) != 2 {
		t.Error("4 / 2 did not equal 2")
	}
	if Div(2, 2) != 1 {
		t.Error("2 / 2 did not equal 1")
	}
}
