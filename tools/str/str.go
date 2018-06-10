package str

import (
	"fmt"
	"strconv"
	"time"
)

func parseBool(str string) (value bool, err error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "y", "ON", "on", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "n", "OFF", "off", "Off":
		return false, nil
	}
	return false, fmt.Errorf("parsing \"%s\": invalid syntax", str)
}

// Bool returns bool type value.
func Bool(str string) (bool, error) {
	return parseBool(str)
}

// Float64 returns float64 type value.
func Float64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Int returns int type value.
func Int(str string) (int, error) {
	return strconv.Atoi(str)
}

// Int64 returns int64 type value.
func Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// Uint returns uint type valued.
func Uint(str string) (uint, error) {
	u, e := strconv.ParseUint(str, 10, 64)
	return uint(u), e
}

// Uint64 returns uint64 type value.
func Uint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// Duration returns time.Duration type value.
func Duration(str string) (time.Duration, error) {
	return time.ParseDuration(str)
}

// TimeFormat parses with given format and returns time.Time type value.
func TimeFormat(str, format string) (time.Time, error) {
	return time.Parse(format, str)
}

// Time parses with RFC3339 format and returns time.Time type value.
func Time(str string) (time.Time, error) {
	return TimeFormat(str, time.RFC3339)
}

// MustString returns default value if key value is empty.
func MustString(str, defaultVal string) string {
	val := str
	if len(val) == 0 {
		return defaultVal
	}
	return val
}

// MustBool always returns value without error,
// it returns false if error occurs.
func MustBool(str string, defaultVal bool) bool {
	val, err := Bool(str)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustFloat64 always returns value without error,
// it returns 0.0 if error occurs.
func MustFloat64(str string, defaultVal float64) float64 {
	val, err := Float64(str)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustInt always returns value without error,
// it returns 0 if error occurs.
func MustInt(str string, defaultVal int) int {
	val, err := Int(str)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustInt64 always returns value without error,
// it returns 0 if error occurs.
func MustInt64(str string, defaultVal int64) int64 {
	val, err := Int64(str)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustUint always returns value without error,
// it returns 0 if error occurs.
func MustUint(str string, defaultVal uint) uint {
	val, err := Uint(str)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustUint64 always returns value without error,
// it returns 0 if error occurs.
func MustUint64(str string, defaultVal uint64) uint64 {
	val, err := Uint64(str)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustDuration always returns value without error,
// it returns zero value if error occurs.
func MustDuration(str string, defaultVal time.Duration) time.Duration {
	val, err := Duration(str)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustTimeFormat always parses with given format and returns value without error,
// it returns zero value if error occurs.
func MustTimeFormat(str, format string, defaultVal time.Time) time.Time {
	val, err := TimeFormat(str, format)
	if err != nil {
		return defaultVal
	}
	return val
}

// MustTime always parses with RFC3339 format and returns value without error,
// it returns zero value if error occurs.
func MustTime(str string, defaultVal time.Time) time.Time {
	return MustTimeFormat(str, time.RFC3339, defaultVal)
}
