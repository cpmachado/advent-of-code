package day02

import "testing"

var mapSafe = map[bool]string{
	false: "unsafe",
	true:  "safe",
}

func TestReportIsSafe(t *testing.T) {
	tests := []struct {
		report Report
		want   bool
	}{
		{report: Report{7, 6, 4, 2, 1}, want: true},
		{report: Report{1, 2, 7, 8, 9}, want: false},
		{report: Report{9, 7, 6, 2, 1}, want: false},
		{report: Report{1, 3, 2, 4, 5}, want: false},
		{report: Report{8, 6, 4, 4, 1}, want: false},
		{report: Report{1, 3, 6, 7, 9}, want: true},
	}

	for _, test := range tests {
		report := test.report

		got := report.IsSafe()
		want := test.want

		if got != want {
			t.Fatalf("Expected %s report(%v), but got %s", mapSafe[want], report, mapSafe[got])
		}
	}
}

func TestCountSafeReports(t *testing.T) {
	reports := []Report{
		Report{7, 6, 4, 2, 1},
		Report{1, 2, 7, 8, 9},
		Report{9, 7, 6, 2, 1},
		Report{1, 3, 2, 4, 5},
		Report{8, 6, 4, 4, 1},
		Report{1, 3, 6, 7, 9},
	}

	got := CountSafeReports(reports)
	want := 2
	if got != want {
		t.Fatalf("Expected '%d' safe reports, but got '%d'", want, got)
	}
}
