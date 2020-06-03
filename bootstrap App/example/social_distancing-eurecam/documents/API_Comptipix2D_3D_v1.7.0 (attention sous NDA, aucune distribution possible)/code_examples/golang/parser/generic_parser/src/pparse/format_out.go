package pparse

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"pconf"
	"pformat"
	"putil"
	"time"
)

// OutPutJson_1
type OutPutJson_1 struct {
	Site                string     `json:"site" xml:"site"`                               // site name
	Chain               string     `json:"chain" xml:"chain"`                             // chain name
	Type                string     `json:"type" xml:"type"`                               // will be "EX":Entries/exits, "E":entries, "X":exits
	Resolution          int64      `json:"resolution" xml:"resolution"`                   // File resolution in seconds
	Begin_end           []string   `json:"begin_end" xml:"begin_end"`                     // Array containing 2 value [begin,end] : a begin and a end date in ISO format "YYYY-MM-DD HH:MM:SS"
	Opening_hour        [][]string `json:"opening_hour" xml:"opening_hour"`               // Array containing tuple value Opening,Closing hour "HH:MM:SS"
	Data_channel        []string   `json:"data_channel" xml:"data_channel"`               // All channels
	Data_channel_access []int      `json:"data_channel_access" xml:"data_channel_access"` // All channels position in data_channel that are access channels
	Data_time           []string   `json:"data_time" xml:"data_time"`                     // All time for data in ISO : "YYYY-MM-DD HH:MM:SS"
	Data                [][]int    `json:"data" xml:"data"`                               // All by time: [channels[E1,X1,E2,X2,...],[E1,X1,E2,X2,...]]
}

// isTimestampInOpening return true if timestamp is inside opening hour, or if ther isn't any opening hour
func isTimestampInOpening(timestampFile int64, timestampStep int64, conf *pconf.Config) bool {
	timestamp := timestampStep - timestampFile

	if len(conf.Out_opening_hour) > 0 {
		for k := range conf.Out_opening_hour {
			if (timestamp >= int64(conf.Out_opening_hour[k].Open.Seconds())) && (timestamp <= int64(conf.Out_opening_hour[k].Close.Seconds())) {
				return true
			} //else {
			// 	fmt.Printf("t:%d not in %d , %d\n", timestamp, int64(conf.Out_opening_hour[k].Open.Seconds()), int64(conf.Out_opening_hour[k].Close.Seconds()))
			// }
		}
	} else {
		return true
	}

	return false
}

// getFormatedOpeningHour return human readable opening hour
func getFormatedOpeningHour(conf *pconf.Config) [][]string {
	opening := make([][]string, len(conf.Out_opening_hour))

	if len(conf.Out_opening_hour) > 0 {
		for k := range conf.Out_opening_hour {
			opening[k] = []string{putil.FormatDuration(conf.Out_opening_hour[k].Open), putil.FormatDuration(conf.Out_opening_hour[k].Close)}
			// opening = append(opening, toAppend)
		}
	} else {
		opening = append(opening, []string{"00:00:00", "23:59:59"})
	}

	return opening
}

// Add method to file parsed
// --------------------------

// Get Comptipix data name file + Comptipix content
// NOTE: Comptipix is the same as input but the result will have some modifications :
// - file precision is set by user
// - duplicate timestamp line are summed
// - duplicate headers are removed
// - can keep only access channel or only enabled channel, removing all the others
func (file_parsed *FileParsed) getOutComptipixCount(conf *pconf.Config) (cpx_name string, cpx_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	// cpx_date_format := DateTimeFormat{ // Default Date-Time Comptipix is "02/01/2006,15:04:05"
	// 	Date_format_type: DATE_DDMMYYYY,
	// 	Date_format_sep: "/",
	// 	Date_time_sep: ",",
	// 	Date_time_order: DT_ORDER_DATE_TIME,
	// 	Time_format_type: TIME_HHMMSS,
	// 	Time_format_sep: ":",
	// }

	// Content

	// Re build header
	c_i := 1                                                // Channel id
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line
	cpx_content = ""
	if !conf.Out_strip_header {
		cpx_content += fmt.Sprintf("%s,%s%s", file_parsed.Site, file_parsed.Chain, eol) // Site,Chain
		cpx_content += fmt.Sprintf("fichier de comptage v2%s", eol)
		date_hour_header := "Date,Heure"
		for c := range file_parsed.FileInfo.All_channel {
			is_wanted_header, _ := putil.SliceIntContains(channel_to_check, c+1) // Channel start at 1, so use k+1 to look for a channel
			if is_wanted_header {
				cpx_content += fmt.Sprintf("%d,", c_i) // Channel id
				chan_name, chan_name_ok := file_parsed.FileInfo.Channel_name[c+1]
				if chan_name_ok {
					cpx_content += fmt.Sprintf("%s,", chan_name) // Channel name
				} else {
					cpx_content += ","
				}
				is_chan_access, _ := putil.SliceIntContains(file_parsed.FileInfo.Access_channel, c+1)
				if is_chan_access {
					cpx_content += fmt.Sprintf("acces%s", eol) // Channel access
				} else {
					cpx_content += fmt.Sprintf("passage%s", eol) // Channel crossing
				}

				// Complete date_hour_header
				date_hour_header += fmt.Sprintf(",E%d,S%d", c_i, c_i)

				c_i++
			}
		}
		cpx_content += fmt.Sprintf("%s%s", date_hour_header, eol)
	}

	// Re build content
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			// cpx_content += fmt.Sprintf("%s,%s", date_line.Format("02/01/2006"), date_line.Format("15:04:05"))
			cpx_content += pformat.GetTimeFormated(&date_line, &file_parsed.FileInfo.Date_time_format, &conf.Date_time_format)
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					cpx_content += fmt.Sprintf(",%d,%d", time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit)
				}
			}
			cpx_content += fmt.Sprintf("%s", eol)
		}
	}

	// File name
	// '-> Comptipix counting name is YYYYMMDD.csv
	cpx_name = fmt.Sprintf("%s.csv", time.Unix(file_parsed.FileTimestamp, 0).Format("20060102"))

	return cpx_name, cpx_content
}

// Get Comptipix occupancy data name file + Comptipix occupancy content from an Comptipix counting file
func (file_parsed *FileParsed) getOutComptipixOcc(conf *pconf.Config) (cpx_name string, cpx_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	// cpx_date_format := DateTimeFormat{ // Default Date-Time Comptipix is "02/01/2006,15:04:05"
	// 	Date_format_type: DATE_DDMMYYYY,
	// 	Date_format_sep: "/",
	// 	Date_time_sep: ",",
	// 	Date_time_order: DT_ORDER_DATE_TIME,
	// 	Time_format_type: TIME_HHMMSS,
	// 	Time_format_sep: ":",
	// }

	// Content
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line
	cpx_content = ""

	// Header
	// Occupancy header is just :
	// "site,chain"
	// fichier de presence v1
	// Date,Heure,E,S,P,C+,C-
	cpx_content += fmt.Sprintf("%s,%s%s", file_parsed.Site, file_parsed.Chain, eol) // Site,Chain
	cpx_content += fmt.Sprintf("fichier de presence v1%s", eol)
	cpx_content += fmt.Sprintf("Date,Heure,E,S,P,C+,C-%s", eol)

	// Re build content
	occ_total := 0
	line_e, line_x, tmp_correction_e, tmp_correction_x := 0, 0, 0, 0
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			// cpx_content += fmt.Sprintf("%s,%s", date_line.Format("02/01/2006"), date_line.Format("15:04:05"))
			cpx_content += pformat.GetTimeFormated(&date_line, &file_parsed.FileInfo.Date_time_format, &conf.Date_time_format)

			if 0 == conf.In_format {
				// We rebuild occ from E/X
				line_e, line_x = 0, 0 // We rebuild occ from E/X
				for k := range time_data[j].ChannelData {
					is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
					if is_wanted {
						line_e += time_data[j].ChannelData[k].Entry
						line_x += time_data[j].ChannelData[k].Exit
					} else {
						fmt.Printf("pouette %d\n", k)
					}
				}
				line_occ := line_e - line_x
				occ_total += line_occ

				// Clip occupancy
				tmp_correction_e, tmp_correction_x = 0, 0
				if occ_total < 0 {
					tmp_correction_e = 0 - occ_total
					occ_total = 0 // Don't allow negative occupancy
				}
			} else {
				// Rewrite occ from occupancy data
				if len(time_data[j].ChannelData) > 0 {
					//fmt.Printf("%s: Chan(%d)(%d): %+v\n", pformat.GetTimeFormated(&date_line, &file_parsed.FileInfo.Date_time_format, &conf.Date_time_format), j, 0, time_data[j].ChannelData)
					line_e = time_data[j].ChannelData[0].Entry
					line_x = time_data[j].ChannelData[0].Exit
					occ_total = time_data[j].ChannelData[0].Occupancy
					tmp_correction_e = time_data[j].ChannelData[0].CorrectionPositive
					tmp_correction_x = time_data[j].ChannelData[0].CorrectionNegative
				} else {
					fmt.Printf("%s: Nothing!!!\n", pformat.GetTimeFormated(&date_line, &file_parsed.FileInfo.Date_time_format, &conf.Date_time_format))
				}
			}

			cpx_content += fmt.Sprintf(",%d,%d,%d,%d,%d%s", line_e, line_x, occ_total, tmp_correction_e, tmp_correction_x, eol) // Write E correction if occupancy has been clipped
		}
	}

	// File name
	// '-> Comptipix occupancy name is YYYYMMDD_presence.csv
	cpx_name = fmt.Sprintf("%s_presence.csv", time.Unix(file_parsed.FileTimestamp, 0).Format("20060102"))

	return cpx_name, cpx_content
}

// Get SLCP sensor data name file + content from an Comptipix counting file
func (file_parsed *FileParsed) getOutSLCPSensor(conf *pconf.Config) (slc_name string, slc_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	slc_date_format := pformat.DateTimeFormat{ // Default Date-Time slcp is "10:05:00"
		Date_format_type: pformat.DATE_NOTHING,
		Date_format_sep:  "",
		Date_time_sep:    "",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_HHMMSS,
		Time_format_sep:  ":",
	}

	// Content
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line
	slc_content = ""

	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			slc_date_str := pformat.GetTimeFormated(&date_line, &slc_date_format, &conf.Date_time_format)
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					// fmt.Printf("%d -> %d:%d = %d / %d (%d)\n",time_data[j].Timestamp, j,k, time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit, channel_to_check[k])//DEBUG
					slc_content += fmt.Sprintf("%s,%d,1,%d,%d%s", slc_date_str, k+1, time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit, eol)
				}
			}
		}
	}

	// File name
	// '-> SLCP sensor name is YYYY/MMDD.csv
	slc_name = fmt.Sprintf("%s.csv", time.Unix(file_parsed.FileTimestamp, 0).Format("2006/0102"))

	return slc_name, slc_content
}

// Get SLCP Area data name file + content from an Comptipix counting file
func (file_parsed *FileParsed) getOutSLCPArea(conf *pconf.Config) (slc_name string, slc_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	slc_date_format := pformat.DateTimeFormat{ // Default Date-Time slcp is "10:05:00"
		Date_format_type: pformat.DATE_NOTHING,
		Date_format_sep:  "",
		Date_time_sep:    "",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_HHMMSS,
		Time_format_sep:  ":",
	}

	// Content
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line
	slc_content = ""

	occ_total := 0
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			slc_date_str := pformat.GetTimeFormated(&date_line, &slc_date_format, &conf.Date_time_format)
			line_e, line_x := 0, 0 // We rebuild occ from E/X
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					line_e += time_data[j].ChannelData[k].Entry
					line_x += time_data[j].ChannelData[k].Exit
				}
			}
			line_occ := line_e - line_x
			occ_total += line_occ

			// Clip occupancy
			if occ_total < 0 {
				occ_total = 0 // Don't allow negative occupancy
			}
			slc_content += fmt.Sprintf("%s,1,1,%d%s", slc_date_str, occ_total, eol) // Write E correction if occupancy has been clipped
		}
	}

	// File name
	// '-> SLCP sensor name is YYYY/MMDD.csv
	slc_name = fmt.Sprintf("%s.csv", time.Unix(file_parsed.FileTimestamp, 0).Format("2006/0102"))

	fmt.Printf("Saving SLCP area as %s\n", slc_name) //DEBUG

	return slc_name, slc_content
}

// Get Zul data name file + zul content
func (file_parsed *FileParsed) getOutZul(conf *pconf.Config) (zul_name string, zul_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	zul_date_format := pformat.DateTimeFormat{ // Default Date-Time Zul is "02/01/2006 \t 15:04"
		Date_format_type: pformat.DATE_DDMMYYYY,
		Date_format_sep:  "/",
		Date_time_sep:    "\t",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_HHMM,
		Time_format_sep:  ":",
	}

	// Content
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line
	if !conf.Out_strip_header {
		zul_content = fmt.Sprintf("%d doors%s", len(channel_to_check), eol) // Zul file identify channel as doors
	}

	last_date_line := time.Unix(file_parsed.FileTimestamp, 0).UTC()
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			zul_door := 1
			zul_date_str := pformat.GetTimeFormated(&date_line, &zul_date_format, &conf.Date_time_format)
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					// fmt.Printf("%d -> %d:%d = %d / %d (%d)\n",time_data[j].Timestamp, j,k, time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit, channel_to_check[k])//DEBUG

					// zul_content += fmt.Sprintf("%d\t%s\t%s\t%04d\t%04d\t1\n", zul_door, date_line.Format("02/01/2006"), date_line.Format("15:04"), time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit)
					zul_content += fmt.Sprintf("%d\t%s\t%04d\t%04d\t1%s", zul_door, zul_date_str, time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit, eol)
					zul_door++
				}
			}
			last_date_line = date_line
		}
	}

	// File name
	// '-> zul name is MMDDHHMM.zul
	zul_name = fmt.Sprintf("%s.zul", last_date_line.Format("01021504"))

	return zul_name, zul_content
}

// Get Sum data name file + sum content
func (file_parsed *FileParsed) getOutSum(conf *pconf.Config) (sum_name string, sum_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	sum_date_format := pformat.DateTimeFormat{ // Default Date-Time Sum is "02/01/2006,15:04:05"
		Date_format_type: pformat.DATE_DDMMYYYY,
		Date_format_sep:  "/",
		Date_time_sep:    ",",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_HHMMSS,
		Time_format_sep:  ":",
	}

	// Content
	sum_content = "" // Sum has no header
	line_e, line_x := 0, 0
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					if (pformat.OutFormat(conf.Out_format) == pformat.OUTFORMAT_SUM_EX) || (pformat.OutFormat(conf.Out_format) == pformat.OUTFORMAT_SUM_E) {
						line_e += time_data[j].ChannelData[k].Entry
					}
					if (pformat.OutFormat(conf.Out_format) == pformat.OUTFORMAT_SUM_EX) || (pformat.OutFormat(conf.Out_format) == pformat.OUTFORMAT_SUM_X) {
						line_x += time_data[j].ChannelData[k].Exit
					}
				}
			}

			// Write content

			// Write date if file has a precision < 1 day
			if conf.Out_resolution != 86400 {
				// sum_content += fmt.Sprintf("%s,%s", date_line.Format("02/01/2006"), date_line.Format("15:04:05"))
				sum_content += fmt.Sprintf("%s", pformat.GetTimeFormated(&date_line, &sum_date_format, &conf.Date_time_format))
			}

			// Write sum
			if pformat.OutFormat(conf.Out_format) == pformat.OUTFORMAT_SUM_EX {
				sum_content += fmt.Sprintf("%d,%d%s", line_e, line_x, eol)
			} else if pformat.OutFormat(conf.Out_format) == pformat.OUTFORMAT_SUM_E {
				sum_content += fmt.Sprintf("%d%s", line_e, eol)
			} else if pformat.OutFormat(conf.Out_format) == pformat.OUTFORMAT_SUM_X {
				sum_content += fmt.Sprintf("%d%s", line_x, eol)
			}
		}
	}

	// File name
	// '-> sum name is YYYYMMDD.csv
	sum_name = fmt.Sprintf("%s.csv", time.Unix(file_parsed.FileTimestamp, 0).Format("20060102"))

	return sum_name, sum_content
}

// Get Cyland data name file + cyland content
func (file_parsed *FileParsed) getOutCyland(conf *pconf.Config) (cyl_name string, cyl_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	cyl_date_format := pformat.DateTimeFormat{ // Default Date-Time Sum is "02012006,15:04"
		Date_format_type: pformat.DATE_DDMMYYYY,
		Date_format_sep:  "",
		Date_time_sep:    ";",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_HHMM,
		Time_format_sep:  ":",
	}

	// Content
	cyl_content = ""                                        // Cyland has no header
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			line_e, line_x := 0, 0 // Cyland has no multi column support : sum all wanted channel
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					line_e += time_data[j].ChannelData[k].Entry
					line_x += time_data[j].ChannelData[k].Exit
				}
			}
			// cyl_content += fmt.Sprintf("%s;%s;1;%d;%d\n", date_line.Format("02012006"), date_line.Format("15:04"), line_e, line_x)
			cyl_content += fmt.Sprintf("%s;1;%d;%d%s", pformat.GetTimeFormated(&date_line, &cyl_date_format, &conf.Date_time_format), line_e, line_x, eol)
		}
	}

	// File name
	// '-> Cyland name is YYYYMMDD.csv
	cyl_name = fmt.Sprintf("%s.csv", time.Unix(file_parsed.FileTimestamp, 0).Format("20060102"))

	return cyl_name, cyl_content
}

// Get a json or xml type 1 data name + content
func (file_parsed *FileParsed) getOutJsonXml_1(conf *pconf.Config) (js1_name string, js1_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	js1_date_format := pformat.DateTimeFormat{ // Default Date-Time Sum is "2006-01-02 15:04:05"
		Date_format_type: pformat.DATE_YYYYMMDD,
		Date_format_sep:  "-",
		Date_time_sep:    " ",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_HHMMSS,
		Time_format_sep:  ":",
	}

	json_file := OutPutJson_1{
		Site:       file_parsed.FileInfo.Site,
		Chain:      file_parsed.FileInfo.Chain,
		Type:       "EX",
		Resolution: conf.Out_resolution,
	}

	// Get channel
	c_i := 1
	for _, c_pos := range channel_to_check {
		chan_name, chan_name_ok := file_parsed.FileInfo.Channel_name[c_pos]
		if chan_name_ok {
			json_file.Data_channel = append(json_file.Data_channel, chan_name)
			is_access, _ := putil.SliceIntContains(file_parsed.FileInfo.Access_channel, c_pos)
			if is_access {
				json_file.Data_channel_access = append(json_file.Data_channel_access, c_i)
			}
			c_i++
		}
	}

	// Fill data
	begin_timestamp := file_parsed.FileTimestamp + 86400 // Add +1 day (later we will compare timestamp used to use the smaller one)
	end_timestamp := file_parsed.FileTimestamp
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			// Data time
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			json_file.Data_time = append(json_file.Data_time, fmt.Sprintf("%s", pformat.GetTimeFormated(&date_line, &js1_date_format, &conf.Date_time_format)))

			// Update end timestamp (the last timestamp with data)
			if time_data[j].Timestamp > end_timestamp {
				end_timestamp = time_data[j].Timestamp
			}

			// Update begin timestamp (the smallest timestamp with data)
			if time_data[j].Timestamp < begin_timestamp {
				begin_timestamp = time_data[j].Timestamp
			}

			// Data channel
			var tmp_data_chan []int
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					tmp_data_chan = append(tmp_data_chan, time_data[j].ChannelData[k].Entry)
					tmp_data_chan = append(tmp_data_chan, time_data[j].ChannelData[k].Exit)
				}
			}
			json_file.Data = append(json_file.Data, tmp_data_chan)
		}
	}

	// Begin end
	// fmt.Printf("LAST End timestamp : %d = %s\n", end_timestamp, time.Unix(end_timestamp, 0).Format("2006-01-02 15:04:05"))
	json_file.Begin_end = append(json_file.Begin_end, fmt.Sprintf("%s", time.Unix(begin_timestamp, 0).UTC().Format("2006-01-02 15:04:05")))
	json_file.Begin_end = append(json_file.Begin_end, fmt.Sprintf("%s", time.Unix(end_timestamp, 0).UTC().Format("2006-01-02 15:04:05")))

	// Opening hour
	json_file.Opening_hour = getFormatedOpeningHour(conf)

	// Encode to json or xml
	js1_content = ""
	if pformat.OUTFORMAT_JSON_1 == pformat.OutFormat(conf.Out_format) {
		export_js1, export_js1_err := json.Marshal(json_file)
		if nil == export_js1_err {
			js1_content = string(export_js1[:])
		}
	} else if pformat.OUTFORMAT_JSON_AS_XML_1 == pformat.OutFormat(conf.Out_format) {
		export_xm1, export_xm1_err := xml.Marshal(json_file)
		if nil == export_xm1_err {
			js1_content = string(export_xm1[:])
		}
	} else {
		js1_content = json_file.getContentXml_1(conf)
	}

	// File name
	// '-> Json1 name is YYYYMMDD.json or YYYYMMDD.xml
	js1_ext := "json"
	if (pformat.OUTFORMAT_JSON_AS_XML_1 == pformat.OutFormat(conf.Out_format)) || (pformat.OUTFORMAT_XML_1 == pformat.OutFormat(conf.Out_format)) {
		js1_ext = "xml"
	}
	js1_name = fmt.Sprintf("%s.%s", time.Unix(file_parsed.FileTimestamp, 0).Format("20060102"), js1_ext)

	return js1_name, js1_content
}

// Get an XML 1 content from a json 1 struct
func (json_file *OutPutJson_1) getContentXml_1(conf *pconf.Config) (xml1_content string) {
	// Manual XML to manage array in xml + attr better
	// output a pretty xml (with tab and LF)
	xml1_content = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	eol := pformat.GetConfEndOfLine(conf.Out_file_end_line) // End Of Line

	// Header info
	xml1_content += fmt.Sprintf("<xml_1 date=\"%s\">%s", json_file.Begin_end[0], eol)
	xml1_content += fmt.Sprintf("\t<site>%s</site>%s", json_file.Site, eol)
	xml1_content += fmt.Sprintf("\t<chain>%s</chain>%s", json_file.Chain, eol)
	xml1_content += fmt.Sprintf("\t<type>%s</type>%s", json_file.Type, eol)
	xml1_content += fmt.Sprintf("\t<resolution>%d</resolution>%s", json_file.Resolution, eol)
	xml1_content += fmt.Sprintf("\t<date>%s", eol)
	xml1_content += fmt.Sprintf("\t\t<begin>%s</begin>%s", json_file.Begin_end[0], eol)
	xml1_content += fmt.Sprintf("\t\t<end>%s</end>%s", json_file.Begin_end[1], eol)
	xml1_content += fmt.Sprintf("\t</date>%s", eol)
	xml1_content += fmt.Sprintf("\t<opening_hour>%s", eol)
	for zz := range json_file.Opening_hour {
		xml1_content += fmt.Sprintf("\t\t<open id=\"%d\">%s</open>%s", zz, json_file.Opening_hour[zz][0], eol)
		xml1_content += fmt.Sprintf("\t\t<close id=\"%d\">%s</close>%s", zz, json_file.Opening_hour[zz][1], eol)
	}
	xml1_content += fmt.Sprintf("\t</opening_hour>%s", eol)

	// Channels names + access/crossing type
	xml1_content += fmt.Sprintf("\t<channel>%s", eol)
	xml1_content += fmt.Sprintf("\t\t<name>%s", eol)
	z := 1
	for x := range json_file.Data_channel {
		chan_type := "crossing"
		is_access, _ := putil.SliceIntContains(json_file.Data_channel_access, z)
		if is_access {
			chan_type = "access"
		}
		xml1_content += fmt.Sprintf("\t\t\t<chan id=\"%d\" type=\"%s\">%s</chan>%s", z, chan_type, json_file.Data_channel[x], eol)
		z++
	}
	xml1_content += fmt.Sprintf("\t\t</name>%s", eol)
	xml1_content += fmt.Sprintf("\t</channel>%s", eol)

	// Data by data_times slot
	xml1_content += fmt.Sprintf("\t<data>%s", eol)
	for x := range json_file.Data {
		xml1_content += fmt.Sprintf("\t\t<data_group time=\"%s\">%s", json_file.Data_time[x], eol)
		z = 1
		chan_id := 1
		chan_type := "E"
		for y := range json_file.Data[x] {
			xml1_content += fmt.Sprintf("\t\t\t<data_chan chan_id=\"%d\" type=\"%s\">%d</data_chan>%s", chan_id, chan_type, json_file.Data[x][y], eol)
			if 0 == z%2 {
				chan_type = "E"
				chan_id++
			} else {
				chan_type = "X"
			}
			z++
		}
		xml1_content += fmt.Sprintf("\t\t</data_group>%s", eol)
	}
	xml1_content += fmt.Sprintf("\t</data>%s", eol)
	xml1_content += fmt.Sprintf("</xml_1>%s", eol)

	return xml1_content
}

// Get a "win" file
func (file_parsed *FileParsed) getOutWin(conf *pconf.Config) (win_name string, win_content string) {
	time_data := GetTimestampedData(file_parsed, conf)
	channel_to_check := GetConfChannelToUse(file_parsed, conf)
	win_date_format := pformat.DateTimeFormat{
		Date_format_type: pformat.DATE_TIMESTAMP_FULL, // Win format is just a timestamp
		Date_format_sep:  "",
		Date_time_sep:    "",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_NOTHING,
		Time_format_sep:  "",
	}

	// Win format is :
	// timestamp:E|X;timestamp:E|X;timestamp:E|X;timestamp:E|X;...

	// Content
	win_content = "" // Win has no header nor line end
	for j := range time_data {
		if isTimestampInOpening(file_parsed.FileTimestamp, time_data[j].Timestamp, conf) {
			date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
			line_e, line_x := 0, 0 // Win has no multicolumn support : sum all wanted channel
			for k := range time_data[j].ChannelData {
				is_wanted, _ := putil.SliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
				if is_wanted {
					line_e += time_data[j].ChannelData[k].Entry
					line_x += time_data[j].ChannelData[k].Exit
				}
			}
			win_content += fmt.Sprintf("%s:%d|%d;", pformat.GetTimeFormated(&date_line, &win_date_format, &conf.Date_time_format), line_e, line_x)
		}
	}

	// File name
	// '-> Win name is YYYYMMDD.win
	win_name = fmt.Sprintf("%s.win", time.Unix(file_parsed.FileTimestamp, 0).Format("20060102"))

	return win_name, win_content
}

// Public Functions
// --------------------------

// GetOutFile Get output file defined in config from a parsed file
func (file_parsed *FileParsed) GetOutFile(conf *pconf.Config) (out_name string, out_content string) {
	switch pformat.OutFormat(conf.Out_format) {
	case pformat.OUTFORMAT_ZUL:
		out_name, out_content = file_parsed.getOutZul(conf)
	case pformat.OUTFORMAT_COMPTIPIX_COUNT:
		out_name, out_content = file_parsed.getOutComptipixCount(conf)
	case pformat.OUTFORMAT_COMPTIPIX_OCC:
		out_name, out_content = file_parsed.getOutComptipixOcc(conf)
	case pformat.OUTFORMAT_SLCP_SENSOR:
		out_name, out_content = file_parsed.getOutSLCPSensor(conf)
	case pformat.OUTFORMAT_SLCP_AREA:
		out_name, out_content = file_parsed.getOutSLCPArea(conf)
	case pformat.OUTFORMAT_WIN:
		out_name, out_content = file_parsed.getOutWin(conf)
	case pformat.OUTFORMAT_CYLAND:
		out_name, out_content = file_parsed.getOutCyland(conf)
	case pformat.OUTFORMAT_JSON_1:
		out_name, out_content = file_parsed.getOutJsonXml_1(conf)
	case pformat.OUTFORMAT_JSON_AS_XML_1:
		out_name, out_content = file_parsed.getOutJsonXml_1(conf)
	case pformat.OUTFORMAT_XML_1:
		out_name, out_content = file_parsed.getOutJsonXml_1(conf)
	default:
		out_name, out_content = file_parsed.getOutSum(conf)
	}
	return out_name, out_content
}
