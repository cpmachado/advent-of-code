package day02

const (
	// Levels must differ by at least 1
	MinLevelDifference = 1
	// Levels must differ at most 3
	MaxLevelDifference = 3
)

type Level int

// Checks if level is in a safe range, given the previous level and if it's
// increasing
func (l Level) InSafeRange(prev Level, increasing bool) bool {
	diff := l - prev
	if increasing {
		return MinLevelDifference <= diff && diff <= MaxLevelDifference
	} else {
		return -MinLevelDifference >= diff && diff >= -MaxLevelDifference
	}
}

type Report []Level

// Checks if a given Report is safe, considering if it uses dampener
func (r Report) IsSafe(dampener bool) bool {
	if r.IsStrictSafe() {
		return true
	}
	if !dampener {
		return false
	}
	aux := make(Report, len(r)-1)
	for i := range r {
		copy(aux, r[:i])
		copy(aux[i:], r[i+1:])
		if aux.IsStrictSafe() {
			return true
		}
	}
	return false
}

// Checks if a given Report is strictly safe
func (r Report) IsStrictSafe() bool {
	if len(r) < 3 {
		return true
	}

	increasing := r[0] < r[1]
	prev := r[0]

	for _, curr := range r[1:] {
		if !curr.InSafeRange(prev, increasing) {
			return false
		}
		prev = curr
	}

	return true
}

// Count the number of safe reports
func CountSafeReports(reports []Report, dampener bool) int {
	count := 0
	for _, report := range reports {
		if report.IsSafe(dampener) {
			count++
		}
	}
	return count
}
