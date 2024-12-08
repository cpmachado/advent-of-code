package day02

import "testing"

var mapSafe = map[bool]string{
	false: "unsafe",
	true:  "safe",
}

func TestReportIsSafe(t *testing.T) {
	tests := []struct {
		report   Report
		dampener bool
		want     bool
	}{
		{report: Report{7, 6, 4, 2, 1}, dampener: false, want: true},
		{report: Report{1, 2, 7, 8, 9}, dampener: false, want: false},
		{report: Report{9, 7, 6, 2, 1}, dampener: false, want: false},
		{report: Report{1, 3, 2, 4, 5}, dampener: false, want: false},
		{report: Report{8, 6, 4, 4, 1}, dampener: false, want: false},
		{report: Report{1, 3, 6, 7, 9}, dampener: false, want: true},
		{report: Report{7, 6, 4, 2, 1}, dampener: true, want: true},
		{report: Report{1, 2, 7, 8, 9}, dampener: true, want: false},
		{report: Report{9, 7, 6, 2, 1}, dampener: true, want: false},
		{report: Report{1, 3, 2, 4, 5}, dampener: true, want: true},
		{report: Report{8, 6, 4, 4, 1}, dampener: true, want: true},
		{report: Report{1, 3, 6, 7, 9}, dampener: true, want: true},
		// skip first level as per reddit discussion
		{report: Report{1, 1, 2, 3, 4}, dampener: true, want: true},
		{report: Report{2, 5, 4, 3, 2}, dampener: true, want: true},
	}

	for _, test := range tests {
		report := test.report

		got := report.IsSafe(test.dampener)
		want := test.want

		if got != want {
			t.Fatalf(
				"Expected %s report(%v) and dampener(%v), but got %s",
				mapSafe[want], report, test.dampener, mapSafe[got])
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

	tests := []struct {
		description string
		dampener    bool
		want        int
	}{
		{description: "count safe reports without dampener", dampener: false, want: 2},
		{description: "count safe reports with dampener", dampener: true, want: 4},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := CountSafeReports(reports, test.dampener)
			want := test.want
			if got != want {
				t.Fatalf("Expected '%d' safe reports, but got '%d'", want, got)
			}
		})
	}

}
