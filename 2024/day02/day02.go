package day02

const (
	// Levels must differ by at least 1
	MinLevelDifference = 1
	// Levels must differ at most 3
	MaxLevelDifference = 3
)

type Level int

func (l Level) InSafeRange(prev Level, increasing bool) bool {
	diff := l - prev
	if increasing {
		return MinLevelDifference <= diff && diff <= MaxLevelDifference
	} else {
		return -MinLevelDifference >= diff && diff >= -MaxLevelDifference
	}
}

type Report []Level

func (r Report) IsSafe() bool {
	if len(r) < 3 {
		return true
	}

	increasing := r[0] < r[1]
	prev := r[0]

	for i := 1; i < len(r); i++ {
		curr := r[i]
		if !curr.InSafeRange(prev, increasing) {
			return false
		}
		prev = curr
	}

	return true
}

func CountSafeReports(reports []Report) int {
	count := 0
	for _, report := range reports {
		if report.IsSafe() {
			count++
		}
	}
	return count
}