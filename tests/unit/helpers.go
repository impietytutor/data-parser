package data_parser

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetHTTPResponse retrieves the HTTP response from the given URL
func GetHTTPResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err!= nil {
		return nil, err
	}
	return resp, nil
}

// GetHTTPResponseWithClient retrieves the HTTP response from the given URL using a custom HTTP client
func GetHTTPResponseWithClient(url string, client *http.Client) (*http.Response, error) {
	resp, err := client.Get(url)
	if err!= nil {
		return nil, err
	}
	return resp, nil
}

// GetHTTPResponseWithTimeout retrieves the HTTP response from the given URL with a custom timeout
func GetHTTPResponseWithTimeout(url string, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err!= nil {
		return nil, err
	}
	return resp, nil
}

// StringSliceContains checks if a string slice contains a specific string
func StringSliceContains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// StringSliceContainsAny checks if a string slice contains any of the strings in another slice
func StringSliceContainsAny(slice []string, other []string) bool {
	for _, str := range other {
		if StringSliceContains(slice, str) {
			return true
		}
	}
	return false
}

// ParseInt parses a string to an integer using the given base
func ParseInt(str string, base int) (int, error) {
	i, err := strconv.ParseInt(str, base, 64)
	if err!= nil {
		return 0, err
	}
	return int(i), nil
}

// ParseFloat64 parses a string to a float64
func ParseFloat64(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err!= nil {
		return 0, err
	}
	return f, nil
}

// TrimStrings trims the given strings to a specified length
func TrimStrings(strings []string, length int) []string {
	trimmed := make([]string, 0, len(strings))
	for _, str := range strings {
		if len(str) > length {
			trimmed = append(trimmed, str[:length])
		} else {
			trimmed = append(trimmed, str)
		}
	}
	return trimmed
}

// ReadAll reads the entire contents of the given reader
func ReadAll(reader io.Reader) ([]byte, error) {
	buf := make([]byte, 1024)
	var data []byte
	for {
		n, err := reader.Read(buf)
		data = append(data, buf[:n]...)
		if err!= nil {
			if err!= io.EOF {
				return nil, err
			}
			break
		}
	}
	return data, nil
}

// ValidateEmail checks if the given string is a valid email address
func ValidateEmail(email string) bool {
	if!strings.Contains(email, "@") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts)!= 2 {
		return false
	}
	return true
}

// ValidatePhone checks if the given string is a valid phone number
func ValidatePhone(phone string) bool {
	if len(phone)!= 10 ||!strings.Contains(phone, "-") {
		return false
	}
	return true
}

// GetTimeFromTimezone retrieves the current time in the given timezone
func GetTimeFromTimezone(tz string) (time.Time, error) {
	loc, err := time.LoadLocation(tz)
	if err!= nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}

// GetTimeFromTimeString retrieves the time from the given string in the specified format
func GetTimeFromTimeString(str string, layout string) (time.Time, error) {
	t, err := time.Parse(layout, str)
	if err!= nil {
		return time.Time{}, err
	}
	return t, nil
}

// GetDateFromTime retrieves the date from the given time
func GetDateFromTime(t time.Time) time.Time {
	return t.Truncate(time.Hour * 24)
}

// GetDayFromTime retrieves the day of the week from the given time
func GetDayFromTime(t time.Time) string {
	return t.Format("Monday")
}

// GetMonthFromTime retrieves the month from the given time
func GetMonthFromTime(t time.Time) string {
	return t.Format("January")
}

// GetYearFromTime retrieves the year from the given time
func GetYearFromTime(t time.Time) string {
	return t.Format("2006")
}

// GetHourFromTime retrieves the hour from the given time
func GetHourFromTime(t time.Time) string {
	return t.Format("15")
}

// GetMinuteFromTime retrieves the minute from the given time
func GetMinuteFromTime(t time.Time) string {
	return t.Format("04")
}

// GetSecondFromTime retrieves the second from the given time
func GetSecondFromTime(t time.Time) string {
	return t.Format("05")
}

// IsStringInSlice checks if a string is in a given slice
func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

// IsStringInSliceAny checks if a string is in any of the given slices
func IsStringInSliceAny(a string, list...[]string) bool {
	for _, l := range list {
		if IsStringInSlice(a, l) {
			return true
		}
	}
	return false
}

// FormatTime formats the given time in the specified format
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// IsEmptyString checks if the given string is empty
func IsEmptyString(str string) bool {
	return str == ""
}

// IsNil checks if the given value is nil
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	return false
}

// GetErrorType checks the type of the given error
func GetErrorType(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, io.EOF) {
		return io.ErrUnexpectedEOF
	}
	return err
}

// PrintError prints the given error with a message
func PrintError(err error) {
	if err!= nil {
		fmt.Println(err.Error())
	}
}

// PrintErrorWithMessage prints the given error with a message
func PrintErrorWithMessage(err error, message string) {
	if err!= nil {
		fmt.Printf("%s: %s\n", message, err.Error())
	}
}

// PrintErrorWithCode prints the given error with a message and code
func PrintErrorWithCode(err error, code int, message string) {
	if err!= nil {
		fmt.Printf("%s (%d): %s\n", message, code, err.Error())
	}
}