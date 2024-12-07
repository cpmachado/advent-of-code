package day02

import "testing"

func TestSafeReport(t *testing.T) {
	t.Run("Test safe reports", func(t *testing.T) {
		safeReports := [][]int{
			[]int{7, 6, 4, 2, 1},
			[]int{1, 3, 6, 7, 9},
		}

		for _, report := range safeReports {
			if !SafeReport(report) {
				t.Fatalf("Expected '%q' safe report, but wasn't", report)
			}
		}
	})
	t.Run("Test unsafe reports", func(t *testing.T) {
		unsafeReports := [][]int{
			[]int{1, 2, 7, 8, 9},
			[]int{9, 7, 6, 2, 1},
			[]int{1, 3, 2, 4, 5},
			[]int{8, 6, 4, 4, 1},
		}
		for _, report := range unsafeReports {
			if SafeReport(report) {
				t.Fatalf("Expected '%q' unsafe report, but wasn't", report)
			}
		}
	})
}

func TestCountSafeReports(t *testing.T) {
	reports := [][]int{
		[]int{7, 6, 4, 2, 1},
		[]int{1, 2, 7, 8, 9},
		[]int{9, 7, 6, 2, 1},
		[]int{1, 3, 2, 4, 5},
		[]int{8, 6, 4, 4, 1},
		[]int{1, 3, 6, 7, 9},
	}

	got := CountSafeReports(reports)
	want := 2
	if got != want {
		t.Fatalf("Expected '%d' safe reports, but got '%d'", want, got)
	}
}
