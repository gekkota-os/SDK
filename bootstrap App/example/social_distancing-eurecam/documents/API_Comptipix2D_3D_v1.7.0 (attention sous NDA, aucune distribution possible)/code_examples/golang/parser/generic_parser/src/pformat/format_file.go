package pformat

import (
	"fmt"
)

// OutFormat format program can generate
type OutFormat int

const (
	OUTFORMAT_COMPTIPIX_COUNT OutFormat = iota // 0 : comptipix (counting) -> change resolution and sum double timestamp line
	OUTFORMAT_COMPTIPIX_OCC                    // 1 : comptipix (occupancy) -> rebuild occupancy file from a counting file
	OUTFORMAT_SLCP_SENSOR                      // 2 : SLCP sensor
	OUTFORMAT_SLCP_AREA                        // 3 : SLCP area
	OUTFORMAT_ZUL                              // 4 : ZUL
	OUTFORMAT_SUM_EX                           // 5 : SUM entries,exits
	OUTFORMAT_SUM_E                            // 6 : SUM entries
	OUTFORMAT_SUM_X                            // 7 : SUM exits
	OUTFORMAT_CYLAND                           // 8 : Cyland
	OUTFORMAT_WIN                              // 9 : Winflow
	OUTFORMAT_JSON_1                           // 10 : JSON type 1
	OUTFORMAT_JSON_AS_XML_1                    // 11 : JSON type 1 as XML (a copy of JSON 1)
	OUTFORMAT_XML_1                            // 12 : XML type 1 (a more xmlish encoder)
)

var outFormatTxt = [...]string{
	"Comptipix counting -> change resolution and sum double timestamp lines", // 0
	"Comptipix occupancy -> rebuild occupancy file from a counting file",     // 1
	"SLCP Sensor",        // 2
	"SLCP Area",          // 3
	"Zul",                // 4
	"Sum: entries,exits", // 5
	"Sum: entries",       // 6
	"Sum: exits",         // 6
	"Cyland",             // 7
	"Win",                // 8
	"JSON type 1",        // 9
	"JSON type 1 as xml", // 10
	"XML type 1",         // 11
}

// InFormat format program can parse
type InFormat int

const (
	INFORMAT_COMPTIPIX_COUNT InFormat = iota // 0 : comptipix (counting) -> change resolution and sum double timestamp line
	INFORMAT_COMPTIPIX_OCC                   // 1 : comptipix (occupancy) -> change resolution and remove double timestamp line
)

var inFormatTxt = [...]string{
	"Comptipix counting -> change resolution and sum double timestamp lines",    // 0
	"Comptipix occupancy -> change resolution and remove double timestamp line", // 1
}

// Output to use
type OutChannel int

const (
	OUTCHAN_ALL             OutChannel = iota // 0 : All channel in output
	OUTCHAN_ACCESS                            // 1 : Only access channel in output
	OUTCHAN_CROSSING                          // 2 : Only crossing channel in output
	OUTCHAN_SETTED                            // 3 : Only channel with setted name in output
	OUTCHAN_SETTED_ACCESS                     // 4 : Only channel access with setted name in output
	OUTCHAN_SETTED_CROSSING                   // 5 : Only channel crossing with setted name in output
	OUTCHAN_SELECTION                         // 6 : Only channel selected by out_channel_select in output
)

var outChannelTxt = [...]string{
	"All channel",                                     // 0
	"Only access channel",                             // 1
	"Only crossing channel",                           // 2
	"Only channel with setted name",                   // 3
	"Only channel access with setted name",            // 4
	"Only channel crossing with setted name",          // 5
	"Only channel selected by out_channel_select_src", // 6
}

// FileEndLine to use
type FileEOL int

const (
	EOL_LF      FileEOL = iota // 0 : LF
	EOL_CR                     // 1 : CR
	EOL_CRLF                   // 2 : CRLF
	EOL_LFCR                   // 3 : LFCR ... because we never know what client will ask
	NO_END_LINE                // 4 : No end line
)

var fileEOLTxt = [...]string{
	"LF (default)", // 0
	"CR",           // 1
	"CRLF",         // 2
	"LFCR",         // 3
	"No end line",  // 4
}

// Functions
// --------------------------

// PrettyPrint All OutChannel
func PrintAllFileEOL() string {
	str := ""
	for i := 0; i < len(fileEOLTxt); i++ {
		str += fmt.Sprintf("\t%d : %s\n", i, fileEOLTxt[i])
	}
	return str
}

// Pretty print date format conf
func PrintFileEOL(out_file_end_line int) string {
	str := "??"
	if (out_file_end_line >= 0) && (out_file_end_line < len(fileEOLTxt)) {
		str = fileEOLTxt[out_file_end_line]
	}
	return str
}

// Get configured end of line char
func GetConfEndOfLine(out_file_end_line int) string {
	line_end := "\n"
	switch FileEOL(out_file_end_line) {
	case EOL_CR: // 1
		line_end = "\r"
	case EOL_CRLF: // 2
		line_end = "\r\n"
	case EOL_LFCR: // 3
		line_end = "\n\r"
	case NO_END_LINE: // 4
		line_end = ""
	default: // 0 (default: LF)
		line_end = "\n"
	}
	return line_end
}

// PrintOutFormat Pretty Print OutFormat
func PrintOutFormat(out_format int) string {
	str := "??"
	if (out_format >= 0) && (out_format < len(outFormatTxt)) {
		str = outFormatTxt[out_format]
	}
	return str
}

// PrintAllOutFormat Pretty Print All OutFormat
func PrintAllOutFormat() string {
	str := ""
	for i := 0; i < len(outFormatTxt); i++ {
		str += fmt.Sprintf("\t%d : %s\n", i, outFormatTxt[i])
	}
	return str
}

// PrintInFormat Pretty Print InFormat
func PrintInFormat(in_format int) string {
	str := "??"
	if (in_format >= 0) && (in_format < len(inFormatTxt)) {
		str = inFormatTxt[in_format]
	}
	return str
}

// PrintAllInFormat Pretty Print All InFormat
func PrintAllInFormat() string {
	str := ""
	for i := 0; i < len(inFormatTxt); i++ {
		str += fmt.Sprintf("\t%d : %s\n", i, inFormatTxt[i])
	}
	return str
}

// PrintOutChannel Pretty Print OutChannel
func PrintOutChannel(out_channel int, out_channel_select_src string) string {
	str := "Unknow -> will use All channel"
	if (out_channel >= 0) && (out_channel < len(outChannelTxt)) {
		str = outChannelTxt[out_channel]
	}
	if OUTCHAN_SELECTION == OutChannel(out_channel) {
		str = fmt.Sprintf("Use selected channel = %s", out_channel_select_src)
	}
	return str
}

// PrettyPrint All OutChannel
func PrintAllOutChannel() string {
	str := ""
	for i := 0; i < len(outChannelTxt); i++ {
		str += fmt.Sprintf("\t%d : %s\n", i, outChannelTxt[i])
	}
	return str
}
