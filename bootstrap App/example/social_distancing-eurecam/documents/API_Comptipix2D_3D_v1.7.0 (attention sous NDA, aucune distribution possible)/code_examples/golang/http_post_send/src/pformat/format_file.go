
package pformat

import (
	"fmt"
)

// FileEndLine to use
type FileEOL int
const (
	EOL_LF						FileEOL = iota 		// 0 : LF
	EOL_CR													// 1 : CR
	EOL_CRLF													// 2 : CRLF
	EOL_LFCR													// 3 : LFCR ... because we never know what client will ask
	NO_END_LINE												// 4 : No end line
)
var fileEOLTxt = [...]string {
	"LF (default)",										// 0
	"CR",														// 1
	"CRLF",													// 2
	"LFCR",													// 3
	"No end line",											// 4
}


// Functions
// --------------------------

// PrettyPrint All OutChannel
func PrintAllFileEOL() (string) {
	str := ""
	for i := 0; i<len(fileEOLTxt); i++ {
		str += fmt.Sprintf("\t%d : %s\n", i, fileEOLTxt[i])
	}
	return str
}

// Pretty print date format conf
func PrintFileEOL(out_file_end_line int) (string) {
	str := "??"
	if (out_file_end_line >= 0) && (out_file_end_line < len(fileEOLTxt)) {
		str = fileEOLTxt[out_file_end_line]
	}
	return str
}

// Get configured end of line char
func GetConfEndOfLine(out_file_end_line int) (string) {
	line_end := "\n"
	switch FileEOL(out_file_end_line) {
		case EOL_CR: 			// 1
		line_end = "\r"
		case EOL_CRLF: 		// 2
		line_end = "\r\n"
		case EOL_LFCR: 		// 3
		line_end = "\n\r"
		case NO_END_LINE: 	// 4
		line_end = ""
		default:					// 0 (default: LF)
		line_end = "\n"
	}
	return line_end
}
