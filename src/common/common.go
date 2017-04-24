package common

// DateFormat is a format string used tim time.Time.Format()
const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
