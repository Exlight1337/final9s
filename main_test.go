package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		got := generateRandomElements(0)
		if len(got) != 0 {
			t.Errorf("expected empty slice, got len = %d", len(got))
		}
	})

	t.Run("negative size", func(t *testing.T) {
		got := generateRandomElements(-5)
		if len(got) != 0 {
			t.Errorf("expected empty slice, got len = %d", len(got))
		}
	})

	t.Run("positive size", func(t *testing.T) {
		size := 100
		got := generateRandomElements(size)
		if len(got) != size {
			t.Errorf("expected len = %d, got %d", size, len(got))
		}
		for i, v := range got {
			if v <= 0 {
				t.Errorf("element at index %d is not positive: %d", i, v)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		got := maximum([]int{})
		if got != 0 {
			t.Errorf("expected 0 for empty slice, got %d", got)
		}
	})

	t.Run("single element", func(t *testing.T) {
		got := maximum([]int{42})
		if got != 42 {
			t.Errorf("expected 42, got %d", got)
		}
	})

	t.Run("all equal", func(t *testing.T) {
		got := maximum([]int{7, 7, 7, 7})
		if got != 7 {
			t.Errorf("expected 7, got %d", got)
		}
	})

	t.Run("max at the beginning", func(t *testing.T) {
		got := maximum([]int{100, 1, 2, 3})
		if got != 100 {
			t.Errorf("expected 100, got %d", got)
		}
	})

	t.Run("max at the end", func(t *testing.T) {
		got := maximum([]int{1, 2, 3, 99})
		if got != 99 {
			t.Errorf("expected 99, got %d", got)
		}
	})

	t.Run("max in the middle", func(t *testing.T) {
		got := maximum([]int{1, 50, 3, 4})
		if got != 50 {
			t.Errorf("expected 50, got %d", got)
		}
	})

	t.Run("large values", func(t *testing.T) {
		got := maximum([]int{1, 1_000_000, 999_999})
		if got != 1_000_000 {
			t.Errorf("expected 1000000, got %d", got)
		}
	})
}
