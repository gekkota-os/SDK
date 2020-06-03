package pformat

import (
	"fmt"
	"time"
)

// DateFormatType to use
type DateFormatType int
const (														// number : example for Mon Jan 2 2006 (1136210645)
	DATE_NO_CHANGE					DateFormatType = iota // 0 : --> depends on OutFormat
	DATE_NOTHING											// 1 : No date used
	DATE_YYYYMMDD											// 2 : 20060102
	DATE_YYYYDDMM											// 3 : 20060201
	DATE_DDMMYYYY											// 4 : 02012006
	DATE_MMDDYYYY											// 5 : 01022006
	DATE_DDYYYYMM											// 6 : 02200601
	DATE_MMYYYYDD											// 7 : 01200602
	DATE_YYMMDD												// 8 : 060201
	DATE_DDMMYY												// 9 : 020106
	DATE_MMDDYY												// 10 : 010206
	DATE_DDYYMM												// 11 : 020601
	DATE_MMYYDD												// 12 : 010602
	DATE_MMDD												// 13 : 0102
	DATE_DDMM												// 14 : 0201
	DATE_DD													// 15 : 02
	DATE_D													// 16 : 2
	DATE_TIMESTAMP_DATE									// 17 : timestamp for Mon Jan 2 2006 at 00:00:00 = 1136156400
	DATE_TIMESTAMP_FULL									// 18 : timestamp for Mon Jan 2 2006 at 15:04:05 = 1136210645
)
var dateFormatTypeTxt = [...]string { // (example for Mon Jan 2 2006 at GMT-0700)
	"No change input, or the defined type for this output format (Default)",// 0
	"No date (date is not used)",						// 1
	"YYYYMMDD (20060102)",								// 2
	"YYYYDDMM (20060201)",								// 3
	"DDMMYYYY (02012006)",								// 4
	"MMDDYYYY (01022006)",								// 5
	"DDYYYYMM (02200601)",								// 6
	"MMYYYYDD (01200602)",								// 7
	"YYMMDD (060201)",									// 8
	"DDMMYY (020106)",									// 9
	"MMDDYY (010206)",									// 10
	"DDYYMM (020601)",									// 11
	"MMYYDD (010602)",									// 12
	"MMDD (0102)",											// 13
	"DDMM (0201)",											// 14
	"DD (02)",												// 15
	"D (2)",													// 16
	"Unix timestamp at 00:00:00 (for Mon Jan 2 2006 at 00:00:00 = 1136156400)",// 17
	"Unix timestamp (for Mon Jan 2 2006 at 15:04:05 = 1136210645)",// 18
	// "ANSIC (Mon Jan _2 15:04:05 2006)",				// 19
	// "UnixDate (Mon Jan _2 15:04:05 MST 2006)",	// 20
	// "RubyDate (Mon Jan 02 15:04:05 -0700 2006)",	// 21
	// "RFC822 (02 Jan 06 15:04 MST)",					// 22
	// "RFC822Z (02 Jan 06 15:04 -0700)",				// 23
	// "RFC850 (Monday, 02-Jan-06 15:04:05 MST)",	// 24
	// "RFC1123 (Mon, 02 Jan 2006 15:04:05 MST)",	// 25
	// "RFC1123Z (Mon, 02 Jan 2006 15:04:05 -0700)",// 26
	// "RFC3339 (2006-01-02T15:04:05Z07:00)",			// 27
	// "Kitchen (3:04PM)",									// 28
	// "Stamp (Jan _2 15:04:05)",							// 29
}

// TimeFormatType to use
type TimeFormatType int
const (														// number : example for 15:04:05
	TIME_NO_CHANGE					TimeFormatType = iota // 0 : --> depends on OutFormat
	TIME_NOTHING											// 1 : No time used
	TIME_HHMMSS												// 2 : 150405
	TIME_HHMM												// 3 : 1504
	TIME_HH													// 4 : 15
	TIME_H													// 5 : 15
	TIME_SSMMHH												// 6 : 050415
	TIME_MMSSHH												// 7 : 040515
	TIME_MMHHSS												// 8 : 041505
	TIME_MMHH												// 9 : 0415
	TIME_TIMESTAMP_DAY									// 10 : 54245 (=15⋅60⋅60+4⋅60+5)
)
var timeFormatTypeTxt = [...]string { // (example for Mon Jan 2 2006 at 15:04:05)
	"No change input, or the defined type for this output format (Default)",// 0
	"No time (time is not used)",						// 1
	"HHMMSS (150405)",									// 2
	"HHMM (1504)",											// 3
	"HH (15)",												// 4
	"H (15)",												// 5
	"SSMMHH (050415)",									// 6
	"MMSSHH (040515)",									// 7
	"MMHHSS (041505)",									// 8
	"MMHH (0515)",											// 9
	"Number of second since 00:00:00 ( 54245 (=15⋅60⋅60+4⋅60+5) )",// 10
}

// DateTimeFormatOrder to use
type DateTimeFormatOrder int
const (
	DT_ORDER_DATE_TIME			DateTimeFormatOrder = iota // 0 : 2006-01-02 15:04:05
	DT_ORDER_TIME_DATE									// 1 : 15:04:05 2006-01-02
)

// Specify a date format
// Ex : for an ISO date '2018-02-30' -> Date_format_type=1 & Date_format_sep="-"
type DateTimeFormat struct {
	Date_format_type			DateFormatType			`json:"date_format_type"`
	Date_format_sep			string					`json:"date_format_sep"`
	Date_time_sep				string					`json:"date_time_sep"`
	Date_time_order			DateTimeFormatOrder	`json:"date_time_order"`
	Time_format_type			TimeFormatType			`json:"time_format_type"`
	Time_format_sep			string					`json:"time_format_sep"`
}


// Functions
// --------------------------

// PrettyPrint OutResolution
func PrintResolution(out_resolution int64) (string) {
	return fmt.Sprintf("%s", time.Duration(out_resolution)*time.Second)
}

// Pretty print date format conf
func PrintDateFormatType(out_date_format_type int) (string) {
	str := "??"
	if (out_date_format_type >= 0) && (out_date_format_type < len(dateFormatTypeTxt)) {
		str = dateFormatTypeTxt[out_date_format_type]
	}
	return str
}

// Pretty Print All DateFormatType
func PrintAllDateFormatType() (string) {
	str := ""
	for i := 0; i<len(dateFormatTypeTxt); i++ {
		str += fmt.Sprintf("\t%d : %s\n", i, dateFormatTypeTxt[i])
	}
	return str
}

// Pretty print time format conf
func PrintTimeFormatType(out_time_format_type int) (string) {
	str := "??"
	if (out_time_format_type >= 0) && (out_time_format_type < len(timeFormatTypeTxt)) {
		str = timeFormatTypeTxt[out_time_format_type]
	}
	return str
}

// Pretty Print All TimeFormatType
func PrintAllTimeFormatType() (string) {
	str := ""
	for i := 0; i<len(timeFormatTypeTxt); i++ {
		str += fmt.Sprintf("\t%d : %s\n", i, timeFormatTypeTxt[i])
	}
	return str
}



// Get date string according to DateTimeFormat
func GetTimeFormated(date_time *time.Time, format_current *DateTimeFormat, format_user *DateTimeFormat) (date_formated string) {
	format_str := ""

	// Format date
	// ---
	date_str := ""
	d_type := format_user.Date_format_type
	d_sep := format_user.Date_format_sep
	if DATE_NO_CHANGE == format_user.Date_format_type {
		d_type = format_current.Date_format_type
		d_sep = format_current.Date_format_sep
	} else {
		fmt.Printf("OVERRIDE date specified format !!")
	}
	switch DateFormatType(d_type) { // Use go string format : "2006-01-02 15:04:05"
	case DATE_NOTHING: 	// 1
		date_str = ""
	case DATE_YYYYMMDD:	// 2
		date_str = fmt.Sprintf("%s", date_time.Format("2006" + d_sep + "01" + d_sep + "02"))
	case DATE_YYYYDDMM:	// 3
		date_str = fmt.Sprintf("%s", date_time.Format("2006" + d_sep + "02" + d_sep + "01"))
	case DATE_DDMMYYYY:	// 4
		date_str = fmt.Sprintf("%s", date_time.Format("02" + d_sep + "01" + d_sep + "2006"))
	case DATE_MMDDYYYY:	// 5
		date_str = fmt.Sprintf("%s", date_time.Format("01" + d_sep + "02" + d_sep + "2006"))
	case DATE_DDYYYYMM:	// 6
		date_str = fmt.Sprintf("%s", date_time.Format("02" + d_sep + "2006" + d_sep + "01"))
	case DATE_MMYYYYDD:	// 7
		date_str = fmt.Sprintf("%s", date_time.Format("01" + d_sep + "2006" + d_sep + "02"))
	case DATE_YYMMDD:		// 8
		date_str = fmt.Sprintf("%s", date_time.Format("06" + d_sep + "01" + d_sep + "02"))
	case DATE_DDMMYY:		// 9
		date_str = fmt.Sprintf("%s", date_time.Format("02" + d_sep + "01" + d_sep + "06"))
	case DATE_MMDDYY:		// 10
		date_str = fmt.Sprintf("%s", date_time.Format("01" + d_sep + "02" + d_sep + "06"))
	case DATE_DDYYMM:		// 11
		date_str = fmt.Sprintf("%s", date_time.Format("02" + d_sep + "06" + d_sep + "01"))
	case DATE_MMYYDD:		// 12
		date_str = fmt.Sprintf("%s", date_time.Format("01" + d_sep + "06" + d_sep + "02"))
	case DATE_MMDD:		// 13
		date_str = fmt.Sprintf("%s", date_time.Format("01" + d_sep + "02"))
	case DATE_DDMM:		// 14
		date_str = fmt.Sprintf("%s", date_time.Format("02" + d_sep + "01"))
	case DATE_DD:			// 15
		date_str = fmt.Sprintf("%02d", date_time.Day())
	case DATE_D:			// 16
		date_str = fmt.Sprintf("%d", date_time.Day())
	case DATE_TIMESTAMP_DATE:	// 17
		tmp_date_at_0 := time.Date(date_time.Year(), date_time.Month(), date_time.Day(), 0, 0, 0, 0, time.UTC)
		date_str = fmt.Sprintf("%d", tmp_date_at_0.UTC().Unix())
	case DATE_TIMESTAMP_FULL:	// 18
		date_str = fmt.Sprintf("%d", date_time.UTC().Unix())
	default:
		// unknow format
		format_str = "2006" + d_sep + "01" + d_sep + "02"
		date_str = fmt.Sprintf("%s", date_time.Format(format_str))
	}

	// Format time
	// ---
	time_str := ""
	t_type := format_user.Time_format_type
	t_sep := format_user.Time_format_sep
	if TIME_NO_CHANGE == format_user.Time_format_type {
		t_type = format_current.Time_format_type
		t_sep = format_current.Time_format_sep
	}
	switch TimeFormatType(t_type) {// Use go string format : "2006-01-02 15:04:05"
		case TIME_NOTHING: 			// 1
		time_str = ""
		case TIME_HHMMSS: 			// 2
		time_str = fmt.Sprintf("%s", date_time.Format("15" + t_sep + "04" + t_sep + "05"))
		case TIME_HHMM: 				// 3
		time_str = fmt.Sprintf("%s", date_time.Format("15" + t_sep + "04"))
		case TIME_HH: 					// 4
		time_str = fmt.Sprintf("%02d", date_time.Hour())
		case TIME_H:					// 5
		time_str = fmt.Sprintf("%d", date_time.Hour())
		case TIME_SSMMHH:				// 6
		time_str = fmt.Sprintf("%s", date_time.Format("05" + t_sep + "04" + t_sep + "15"))
		case TIME_MMSSHH:				// 7
		time_str = fmt.Sprintf("%s", date_time.Format("04" + t_sep + "05" + t_sep + "15"))
		case TIME_MMHHSS:				// 8
		time_str = fmt.Sprintf("%s", date_time.Format("04" + t_sep + "15" + t_sep + "05"))
		case TIME_MMHH:				// 9
		time_str = fmt.Sprintf("%s", date_time.Format("04" + t_sep + "15"))
		case TIME_TIMESTAMP_DAY:	// 10
		time_str = fmt.Sprintf("%d", date_time.Hour()*60*60 + date_time.Minute()*60 + date_time.Second())
	}

	// Format Date-Time together
	// ---
	dt_sep := format_user.Date_time_sep
	dt_order := format_user.Date_time_order
	if DATE_NO_CHANGE == format_user.Date_format_type {
		dt_sep = format_current.Date_time_sep
		dt_order = format_current.Date_time_order
	}
	date_time_str := date_str + dt_sep + time_str
	if DT_ORDER_TIME_DATE == dt_order {
		date_time_str = time_str + dt_sep + date_str
	}

	return date_time_str
}
