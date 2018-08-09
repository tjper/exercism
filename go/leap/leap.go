// Package leap provides a library to determine if a year is a leap year.
package leap

// IsLeapYear takes a year and determines if it is a leap year, if successful,
// true is returned
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
