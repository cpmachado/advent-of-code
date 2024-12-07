package day02

const (
	// Levels must differ by at least 1
	MinLevelDifference = 1
	// Levels must differ at most 3
	MaxLevelDifference = 3
)

func SafeReport(report []int) bool {
	if len(report) < 3 {
		return true
	}

	increasing := report[0] < report[1]
	prev := report[0]

	for i := 1; i < len(report); i++ {
		curr := report[i]
		if increasing {
			if (prev+MinLevelDifference) > curr || (prev+MaxLevelDifference) < curr {
				return false
			}
		} else {
			if (prev-MinLevelDifference) < curr || (prev-MaxLevelDifference) > curr {
				return false
			}
		}
		prev = curr
	}

	return true
}

func CountSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if SafeReport(report) {
			count++
		}
	}
	return count
}
