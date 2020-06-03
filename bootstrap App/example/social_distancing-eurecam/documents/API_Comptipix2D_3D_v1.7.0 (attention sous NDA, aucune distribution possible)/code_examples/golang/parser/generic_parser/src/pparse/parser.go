package pparse

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"pconf"
	"pformat"
	"putil"
	"strconv"
	"strings"
	"time"
)

// SiteChainHeader : Header parse status (because Comptipix header is 2 lines : first the header, 2nd must be "fichier de comptage v2" or header is not valid)
type SiteChainHeader int32

// SiteChainHeader : site header parsing state machine
const (
	SITE_CHAIN_IDLE SiteChainHeader = iota
	SITE_CHAIN_IN_PORGRESS
	SITE_CHAIN_DONE
)

// DataType : inputFile type
type DataType int

// DataType : inputFile type
const (
	DataCounting DataType = iota
	DataOccupancy
)

// FileInfo all file infos
type FileInfo struct {
	FileType         DataType               `json:"file_type"`        // file type
	Site             string                 `json:"site"`             // site name
	Chain            string                 `json:"chain"`            // chain name
	Site_after       string                 `json:"site_after"`       // site part after Dst_chain_site_after separator
	Chain_after      string                 `json:"chain_after"`      // chain part after Dst_chain_site_after separator
	Date_time_format pformat.DateTimeFormat `json:"date_time_format"` // date and time format type see \ref DateTimeFormat
	Channel_name     map[int]string         `json:"channel_name"`     // channel name
	Enabled_channel  []int                  `json:"enabled_channel"`  // channel id (id is position) that has name != ""
	Access_channel   []int                  `json:"access_channel"`   // channel id (id is position) that are access and taking into account for summing Entry and Exits
	All_channel      []int                  `json:"all_channel"`      // channel id (id is position), all
}

// CountingData : info of 1 line of 1 channel
type CountingData struct {
	Entry              int `json:"entry"`               // entry  --> entry sum for occupancy
	Exit               int `json:"exit"`                // exit	 --> exit sum for occupancy
	Occupancy          int `json:"occupancy"`           // occupancy
	CorrectionPositive int `json:"correction_positive"` // occupancy positive correction
	CorrectionNegative int `json:"correction_negative"` // occupancy negative correction
}

// FileParsed represent a parsed file
type FileParsed struct {
	FileInfo
	FileData           map[int64][]CountingData // map of slice counting info -> index = unix timestamp -> position in slice = channel id
	FileTimestampSteps []int64                  // All timestamp step list
	FileTimestamp      int64                    // File timestamp (starting at 00:00:00)
}

// TimestampedCountingData timestamped counting data for all channel
type TimestampedCountingData struct {
	Timestamp   int64          // unix timestamp
	ChannelData []CountingData // Data for all used channel
}

// OccupancyData : 1 line in occupancy file
type OccupancyData struct {
	EntrySum int `json:"entry_sum"` // entry sum
	ExitSum  int `json:"exit_sum"`  // exit sum

}

// Private Functions
// --------------------------

// GetParsedFile Get parsed data from a file content
// Parsed data are stored at the wanted resolution
func GetParsedFile(fdata io.Reader, t_date_file *time.Time, resolution int64) (file_parsed FileParsed) {
	file_info := FileInfo{Site: "", Chain: "", Channel_name: make(map[int]string)}
	file_parsed = FileParsed{FileInfo: file_info, FileData: make(map[int64][]CountingData)}

	// Round the date_file to day at 00:00:00 (it should be the case but it's better to be sure)
	date_file := time.Date(t_date_file.Year(), t_date_file.Month(), t_date_file.Day(), 0, 0, 0, 0, time.UTC)
	file_parsed.FileTimestamp = date_file.UTC().Unix()

	// Temporary store of Site chain and Site name,
	// because a Site Chain/Name line must be validated by a line "fichier de comptage v2"
	// so we store temporary read site and chain, waiting for validation
	tmp_site := ""
	tmp_chain := ""

	// Build resolution array
	steps := []int64{}
	var step_current int64 = 0
	for step_current < 86400 {
		steps = append(steps, step_current)
		step_current += resolution
	}
	file_parsed.FileTimestampSteps = steps
	// fmt.Printf("\nfile date: %s = %d, resolution: %d\n", date_file.String(), file_parsed.FileTimestamp, resolution)//DEBUG
	// fmt.Printf("%+v\n", steps)//DEBUG

	// Use scanner to read file line by line
	scanner := bufio.NewScanner(fdata)
	site_chain_header_state := SITE_CHAIN_IDLE
	file_date_type := pformat.DateTimeFormat{Date_format_type: pformat.DATE_YYYYMMDD, Date_format_sep: "-", Time_format_type: pformat.TIME_HHMMSS, Time_format_sep: ":"}
	for scanner.Scan() {
		// Read one line
		line := scanner.Text()
		line_split := strings.Split(line, ",")

		// Check site chain header preceed a "fichier de comptage v2" line
		if SITE_CHAIN_IN_PORGRESS == site_chain_header_state {
			if "fichier de comptage v2" == line || "fichier de presence v1" == line {
				site_chain_header_state = SITE_CHAIN_DONE
				file_parsed.Site = tmp_site
				file_parsed.Chain = tmp_chain

				// Set file type
				file_parsed.FileType = DataCounting
				if "fichier de presence v1" == line {
					file_parsed.FileType = DataOccupancy
				}
				continue
			} else {
				site_chain_header_state = SITE_CHAIN_IDLE
			}
		}

		// Check site/chain header
		if 1 == strings.Count(line, ",") {
			if SITE_CHAIN_DONE == site_chain_header_state {
				fmt.Printf("Double header ! (NOTE: it may be normal)\n") // We update Site/Chain header if there is another header
			}
			if len(line_split) > 0 {
				tmp_site = line_split[0]
			}
			if len(line_split) > 1 {
				tmp_chain = line_split[1]
			}
			site_chain_header_state = SITE_CHAIN_IN_PORGRESS
			continue
		}

		// Check Channel type
		if file_parsed.FileType == DataCounting {
			if (2 == strings.Count(line, ",")) && (3 == len(line_split)) {
				c_i, c_err := strconv.Atoi(line_split[0])
				if (nil == c_err) && (c_i > 0) {
					// Add/Remove channel from access channel list
					access_contains_already, access_contains_pos := putil.SliceIntContains(file_parsed.FileInfo.Access_channel, c_i)
					if ("acces" == line_split[2]) && (!access_contains_already) { // Add
						file_parsed.FileInfo.Access_channel = append(file_parsed.FileInfo.Access_channel, c_i) // Is access
					}
					if ("acces" != line_split[2]) && access_contains_already { // Remove
						file_parsed.FileInfo.Access_channel = append(file_parsed.FileInfo.Access_channel[:access_contains_pos], file_parsed.FileInfo.Access_channel[access_contains_pos+1:]...) // Remove the channel preserving the slice order see https://github.com/golang/go/wiki/SliceTricks
					}

					// Add/Remove channel from enabled channel list
					enabled_contains_already, enabled_contains_pos := putil.SliceIntContains(file_parsed.FileInfo.Enabled_channel, c_i)
					if ("" != line_split[1]) && (!enabled_contains_already) { // Add
						file_parsed.FileInfo.Enabled_channel = append(file_parsed.FileInfo.Enabled_channel, c_i) // Is enabled
					}
					if ("" == line_split[1]) && enabled_contains_already { // Remove
						file_parsed.FileInfo.Enabled_channel = append(file_parsed.FileInfo.Enabled_channel[:enabled_contains_pos], file_parsed.FileInfo.Enabled_channel[enabled_contains_pos+1:]...) // Remove the channel preserving the slice order see https://github.com/golang/go/wiki/SliceTricks
					}

					// Add/Update Channel name + all channel
					file_parsed.FileInfo.Channel_name[c_i] = line_split[1]
					file_parsed.FileInfo.All_channel = append(file_parsed.FileInfo.All_channel, c_i) // Add to all channel

					// Nothing more to do with this line
					continue
				}
			}
		}

		// Chech for line "Date,Heure,E,S"
		if strings.HasPrefix(line, "Date,Heure,") {
			continue // Nothing to do
		}

		// Check for a counting -> sum only access channel
		if strings.Count(line, ",") >= 3 {
			// Line format is like :
			// Date,Heure,E,S
			// 07/12/2017,00:00:00,0,0

			// Add channel counting to data map at good resolution
			// ---
			// Read date-time
			line_time, line_date_type, line_time_err := putil.GetDateFromComptipixFileContent(line_split[0] + "," + line_split[1])
			if nil == line_time_err {
				// Check date is the same as date_file
				if (line_time.Year() != date_file.Year()) || (line_time.Month() != date_file.Month()) || (line_time.Day() != date_file.Day()) {
					// Log this : this appen on SDcard FAT corruption -> the content of another file in 1 file
					log.Printf("ERROR: file %s, line: %s should NOT be there --> FAT may be corrupted !!\n", line_time.Format("2006-01-02"), date_file.Format("2006-01-02"))
					continue
				}

				// Get resolution position
				line_time_sec := line_time.Unix() - date_file.Unix() // Keep only second for today
				res_pos := steps[line_time_sec/resolution]
				// fmt.Printf("line: %s  ->  %d = %d (%s)\n", line, line_time_sec, res_pos, line_time.Format("2006-01-02 15:04:05")) //DEBUG

				// Parse all counting column
				counting_data := []CountingData{}
				if file_parsed.FileType == DataCounting {
					// Parse counting data
					for j := 2; j < len(line_split)-1; j += 2 {
						tmp_e, tmp_e_err := strconv.Atoi(line_split[j])
						tmp_x, tmp_x_err := strconv.Atoi(line_split[j+1])
						if (nil == tmp_e_err) && (nil == tmp_x_err) {
							counting_data = append(counting_data, CountingData{Entry: tmp_e, Exit: tmp_x})
						} else {
							// There is error inside file -> Log it!
							log.Printf("ERROR: file %s !!!\n", date_file.Format("2006-01-02"))
							if (nil != tmp_e_err) && (nil != tmp_x_err) {
								counting_data = append(counting_data, CountingData{Entry: 0, Exit: 0}) // Error on both entry + exit
							} else if nil != tmp_e_err {
								counting_data = append(counting_data, CountingData{Entry: 0, Exit: tmp_x}) // Error on entry
							} else {
								counting_data = append(counting_data, CountingData{Entry: tmp_e, Exit: 0}) // Error on exit
							}
						}
					}
				} else {
					// Parse Occupancy Data
					tmpCountingData := CountingData{}
					if len(line_split) > 4 {
						// read Entry Exit and Ocupancy
						tmpE, tmpEerr := strconv.Atoi(line_split[2])
						if nil == tmpEerr {
							tmpCountingData.Entry = tmpE
						}
						tmpX, tmpXerr := strconv.Atoi(line_split[3])
						if nil == tmpXerr {
							tmpCountingData.Exit = tmpX
						}
						tmpO, tmpOerr := strconv.Atoi(line_split[4])
						if nil == tmpOerr {
							tmpCountingData.Occupancy = tmpO
						}
						if len(line_split) > 5 {
							tmpCpositive, tmpCpositiveErr := strconv.Atoi(line_split[5])
							if nil == tmpCpositiveErr {
								tmpCountingData.CorrectionPositive = tmpCpositive
							}
							tmpCnegative, tmpCnegativeErr := strconv.Atoi(line_split[6])
							if nil == tmpCnegativeErr {
								tmpCountingData.CorrectionNegative = tmpCnegative
							}
						}
					} else {
						fmt.Printf("ERROR: occupancy file %s has %d column expect at least 5 !!!\n", date_file.Format("2006-01-02"), len(line_split))
					}
					// Add occupancy data to counting data
					// fmt.Printf("Add c data1(%d) : %+v\n", res_pos, tmpCountingData)
					counting_data = append(counting_data, tmpCountingData)
				}

				// Set it to map or add it to an existing map entrie
				map_val, map_ok := file_parsed.FileData[res_pos]
				if map_ok {
					// A timestamp already present in map
					new_counting_data := []CountingData{}
					if file_parsed.FileType == DataCounting {
						// Counting file :
						// in case of same timestamp : sum data E and X
						for k := range file_parsed.FileData[res_pos] {
							new_e := map_val[k].Entry + counting_data[k].Entry
							new_x := map_val[k].Exit + counting_data[k].Exit
							// fmt.Printf("---> Add-Add %d(%d) : E:%d+%d=%d , X:%d+%d=%d\n", res_pos, k, map_val[k].Entry, counting_data[k].Entry, new_e, map_val[k].Exit, counting_data[k].Exit, new_x)//DEBUG
							new_counting_data = append(new_counting_data, CountingData{Entry: new_e, Exit: new_x})
						}
					} else {
						// Occupancy file :
						// in case of same timestamp : keep the higest occupancy data
						if len(map_val) > 0 && len(counting_data) > 0 {
							if map_val[0].Occupancy < counting_data[0].Occupancy {
								// fmt.Printf("Update c data1(%d) : %+v\n", res_pos, counting_data[0])
								new_counting_data = append(new_counting_data, counting_data[0]) // Replace occupancy
							}
						}
					}
					// if new_counting_data[0].Occupancy > 0 {
					// 	fmt.Printf("---> Add Slice at res %d =\n", res_pos)
					// 	fmt.Printf("%+v\n", new_counting_data)
					// 	fmt.Printf("+\n")
					// 	fmt.Printf("%+v\n", file_parsed.FileData[res_pos])
					// }
					if len(new_counting_data) > 0 {
						//fmt.Printf("-> UPDATE c data1(%d) : %+v\n", res_pos, new_counting_data)
						file_parsed.FileData[res_pos] = new_counting_data
					}

					// fmt.Printf("=\n")
					// fmt.Printf("%+v\n", file_parsed.FileData[res_pos])
				} else {
					// No timestamp : set it with new data

					// fmt.Printf("Data at %s %d -> %+v\n\n", line_time, res_pos, counting_data[0])

					// fmt.Printf("-> ADD c data1(%d) : %+v\n", res_pos, counting_data)
					file_parsed.FileData[res_pos] = counting_data

					// fmt.Printf("--> Set Slice at res %d =\n", res_pos)
					// fmt.Printf("%+v\n", file_parsed.FileData[res_pos])
				}

				// Assign line date type
				file_date_type = line_date_type

			} else {
				log.Printf("date_time ERROR\n")
				//TODO : what to do on line_time err ??
			}
		}
	}

	// Assign read file input date and time type
	// so the last read line will define the date format of the whole file
	file_parsed.Date_time_format = file_date_type

	// fmt.Printf("File parsed %s :\n", date_file.Format("2006-01-02"))
	// for i := range file_parsed.FileTimestampSteps {
	// 	stepVal, stepOk := file_parsed.FileData[int64(i)]
	// 	if stepOk {
	// 		fmt.Printf("Step:%d = %+v\n", i, stepVal)
	// 	} else {
	// 		fmt.Printf("Step:%d = NOTHING!!\n", i)
	// 	}
	// }
	return file_parsed
}

// SumParsedFile sum a parsed file
func SumParsedFile(toSum []FileParsed, toSub []bool, fixOccupancy bool) FileParsed {
	// Invert all counting that need to be inverted (because substract)
	for i := range toSum {
		if toSub[i] {
			for k := range toSum[i].FileData {
				for l := 0; l < len(toSum[i].FileData[k]); l++ {
					toSum[i].FileData[k][l].Entry *= -1
					toSum[i].FileData[k][l].Exit *= -1
					toSum[i].FileData[k][l].Occupancy *= -1
					toSum[i].FileData[k][l].CorrectionPositive *= -1
					toSum[i].FileData[k][l].CorrectionNegative *= -1
				}
			}
		}
	}

	// Keep first element as ref
	parsed := toSum[0]

	// Check there is enough files
	if len(toSum) < 2 {
		fmt.Printf("-!! NOT enough elements in toSum %d\n", len(toSum))
		return parsed
	}

	// Remove first element from slice
	toSum = append(toSum[:1], toSum[2:]...)

	// fmt.Printf("SumParsedFile tosum=%d\n", len(toSum))
	// fmt.Printf("SumParsedFile tosub=%d -> %+v\n", len(toSub), toSub)

	for i := 1; i < len(toSum); i++ {
		// Loop through all file timestamp
		for k := 0; k < len(parsed.FileTimestampSteps); k++ {
			toSumData2, toSumOk2 := toSum[i].FileData[parsed.FileTimestampSteps[k]]
			toSumData1, toSumOk1 := parsed.FileData[parsed.FileTimestampSteps[k]]

			// fmt.Printf("Nothing to sum at %d\n", parsed.FileTimestampSteps[k])
			if toSumOk2 && toSumOk1 {
				// fmt.Printf("To sum[%d][%d] :\n", i, k)
				// fmt.Printf("%+v\n", toSumData1)
				// fmt.Printf("%+v\n", toSumData2)
				sumed := []CountingData{}
				for l := 0; l < len(toSumData1); l++ {
					tmpCountingData := CountingData{
						Entry:              toSumData1[l].Entry + toSumData2[l].Entry,
						Exit:               toSumData1[l].Exit + toSumData2[l].Exit,
						Occupancy:          toSumData1[l].Occupancy + toSumData2[l].Occupancy,
						CorrectionPositive: toSumData1[l].CorrectionPositive + toSumData2[l].CorrectionPositive,
						CorrectionNegative: toSumData1[l].CorrectionNegative + toSumData2[l].CorrectionNegative,
					}
					sumed = append(sumed, tmpCountingData)
				}
				// Write new file (sumed)
				parsed.FileData[parsed.FileTimestampSteps[k]] = sumed
			}
		}
	}

	// At the end loop through all result to fix occupancy negative
	if fixOccupancy {
		for k := range parsed.FileData {
			for l := 0; l < len(parsed.FileData[k]); l++ {
				if parsed.FileData[k][l].Occupancy < 0 {
					parsed.FileData[k][l].CorrectionPositive += -1 * parsed.FileData[k][l].Occupancy
					parsed.FileData[k][l].Occupancy = 0
				}
			}
		}
	}

	// fmt.Printf("SUMED:\n%+v\n\n", parsed)
	return parsed
}

// Public Functions
// --------------------------

// GetConfChannelToUse Get channel to use according to config
func GetConfChannelToUse(file_parsed *FileParsed, conf *pconf.Config) []int {
	channel_to_check := make([]int, len(file_parsed.FileInfo.All_channel), cap(file_parsed.FileInfo.All_channel))
	k := 0
	k_is_access := false

	switch pformat.OutChannel(conf.Out_channel) {
	case pformat.OUTCHAN_ACCESS:
		channel_to_check = make([]int, len(file_parsed.FileInfo.Access_channel), cap(file_parsed.FileInfo.Access_channel))
		copy(channel_to_check, file_parsed.FileInfo.Access_channel)

	case pformat.OUTCHAN_CROSSING:
		channel_to_check = make([]int, 0)
		for k = range file_parsed.FileInfo.All_channel {
			k_is_access, _ = putil.SliceIntContains(file_parsed.FileInfo.Access_channel, file_parsed.FileInfo.All_channel[k])
			if !k_is_access {
				channel_to_check = append(channel_to_check, file_parsed.FileInfo.All_channel[k])
			}
		}

	case pformat.OUTCHAN_SETTED:
		channel_to_check = make([]int, len(file_parsed.FileInfo.Enabled_channel), cap(file_parsed.FileInfo.Enabled_channel))
		copy(channel_to_check, file_parsed.FileInfo.Enabled_channel)

	case pformat.OUTCHAN_SETTED_ACCESS:
		channel_to_check = make([]int, 0)
		for k = range file_parsed.FileInfo.Enabled_channel {
			k_is_access, _ = putil.SliceIntContains(file_parsed.FileInfo.Access_channel, file_parsed.FileInfo.Enabled_channel[k])
			if k_is_access {
				channel_to_check = append(channel_to_check, file_parsed.FileInfo.Enabled_channel[k])
			}
		}

	case pformat.OUTCHAN_SETTED_CROSSING:
		channel_to_check = make([]int, 0)
		for k = range file_parsed.FileInfo.Enabled_channel {
			k_is_access, _ = putil.SliceIntContains(file_parsed.FileInfo.Access_channel, file_parsed.FileInfo.Enabled_channel[k])
			if !k_is_access {
				channel_to_check = append(channel_to_check, file_parsed.FileInfo.Enabled_channel[k])
			}
		}

	case pformat.OUTCHAN_SELECTION:
		channel_to_check = make([]int, len(conf.Out_channel_select), cap(conf.Out_channel_select))
		copy(channel_to_check, conf.Out_channel_select)

	default: //OUTCHAN_ALL
		copy(channel_to_check, file_parsed.FileInfo.All_channel)
	}
	// fmt.Printf("getConfChannelToUse : \n%+v\n", channel_to_check)//DEBUG

	return channel_to_check
}

// GetTimestampedData Get Timestamped data from a file parsed
// Parsed data are stored at the wanted resolution
func GetTimestampedData(file_parsed *FileParsed, conf *pconf.Config) (time_data []TimestampedCountingData) {
	// fmt.Printf("FILE timestamp = %d -> %s\n", file_parsed.FileTimestamp, time.Unix(file_parsed.FileTimestamp, 0).UTC().Format("2006-01-02 15:04:05"))//DEBUG

	for i := range file_parsed.FileTimestampSteps {
		// Get map data
		map_val, map_ok := file_parsed.FileData[file_parsed.FileTimestampSteps[i]]
		if map_ok {
			// fmt.Printf("GetTimestampedData %d -> %s\n",file_parsed.FileTimestamp, time.Unix(file_parsed.FileTimestamp + file_parsed.FileTimestampSteps[i], 0).UTC().Format("2006-01-02 15:04:05"))//DEBUG
			counting_data := TimestampedCountingData{Timestamp: file_parsed.FileTimestamp + file_parsed.FileTimestampSteps[i], ChannelData: map_val}
			add_this_line := true
			if conf.Out_strip_null_line {
				add_this_line = false
				for k := range map_val {
					if (0 != map_val[k].Entry) || (0 != map_val[k].Exit) {
						add_this_line = true // There it at least one counting
						break
					}
				}
			}
			if add_this_line {
				time_data = append(time_data, counting_data)
			}
		}
	}

	return time_data
}
