package calculator

import "testing"

func TestPlus(t *testing.T) {
	t.Run("should plus positive number correctly", func(t *testing.T) {
		input1 := 1
		input2 := 2
		want := 3
		got := Plus(input1, input2)
		if want != got {
			t.Fatalf("unexpected, want %d, got %d", want, got)
		}
	})
}
