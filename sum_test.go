package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should return 3 when input 1 and 2", func(t *testing.T) {
		//Arange
		want := 3

		//Act
		got := sum(1, 2)

		//Assert
		if got != want {
			t.Errorf("sum(1,2) = %d; want 3", got)
		}
	})
	t.Run("should return 1 when input 1 and 0", func(t *testing.T) {
		want := 1

		got := sum(1, 0)

		if got != want {
			t.Errorf("sum(1,0) = %d; want 1", got)
		}
	})
	t.Run("should return -2 when input -1 and -1", func(t *testing.T) {
		want := -2

		got := sum(-1, -1)

		if got != want {
			t.Errorf("sum(-1,-1) = %d; want -2", got)
		}
	})
}
