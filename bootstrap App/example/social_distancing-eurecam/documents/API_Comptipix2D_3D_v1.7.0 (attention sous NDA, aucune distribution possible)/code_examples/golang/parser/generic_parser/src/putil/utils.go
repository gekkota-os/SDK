// Some utils functions

package putil

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Functions
// --------------------------

// Do a verbose print (according to config)
// and produce log (if log set to true)
func PrintAndLog(msg string, verbose bool, do_log bool) {
	if verbose {
		fmt.Printf(msg)
	}
	if do_log {
		log.Printf(msg)
	}
}

// Sanitize directory (Add the last '/' if not present)
func SanitizeDir(dir string) string {
	// Check dir ends with '/'
	if "" != dir {
		if !strings.HasSuffix(dir, "/") {
			dir += "/"
		}
	}
	return dir
}

// Get file existence + size
// Return a value < 0 for a not exist file or directory: -1 for inexistence file, -2 for permission error
// Return the size otherwise
func GetFileSize(file_full_path string) int64 {
	// Check for directory existence
	if fi, fi_err := os.Stat(file_full_path); fi_err != nil {
		if os.IsNotExist(fi_err) {
			return -1 // file does not exist
		} else {
			return -2 // other error (like permission error)
		}
	} else {
		return fi.Size() // Return file size
	}
}

// Return true + index position if found, false + anything otherwise
func SliceIntContains(s []int, e int) (bool, int) {
	for index, a := range s {
		if a == e {
			return true, index
		}
	}
	return false, -1
}

// Return true + index position if found, false + anything otherwise
func SliceStringContains(s []string, e string) (bool, int) {
	for index, a := range s {
		if a == e {
			return true, index
		}
	}
	return false, -1
}

// Return true if date check is between start and end
func InTimeSpan(start time.Time, end time.Time, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// Get string after first occurance
func GetStringAfter(str string, str_after string) (str_result string) {
	index := strings.LastIndex(str, str_after)
	if -1 == index {
		return str
	}
	runes := []rune(str)                                             // Take substring of first word with runes. This handles any kind of rune in the string.
	safe_substring := string(runes[index+len(str_after) : len(str)]) // Convert back into a string from rune slice.

	return safe_substring
}

// Get now (today set at midnight)
func GetNowMidnight() time.Time {
	t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return rounded
}

// FormatDuration format a duration to hh:mm:ss
func FormatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute

	return fmt.Sprintf("%02d:%02d:%02d", h, m, int(d.Seconds()))
}
