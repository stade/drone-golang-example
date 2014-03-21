package math

import "testing"

func TestMath1(t *testing.T) {
    if 1 + 2 != 3 {
        t.Error("Expected 3, got ", 1 + 2)
    }
}

func TestMath2(t *testing.T) {
    if 10 + 20 != 30 {
        t.Error("Expected 30, got ", 10 + 20)
    }
}
