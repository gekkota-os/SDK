# Principle

Generation of a json file compact
Precision (counting data resolution) may vary

## File name

Filename : *YYYYMMDD.json*, exemple : **20180102.json** for 2018 Jan 02

1. 2018 *year*
2. 01 *month*
3. 02 *day*

Extension is **json**

## File Content

- All date are ISO : *YYYY-MM-DD HH:MM:SS*
- Encoding is UTF8
- New line is **LF**
- Counting resolution (time slot data) may vary from 10s to 6h or more (always < 24h)

### File Header and content

File *header* info are in json (site, chain, type, resolution, date begin+end) :
And file data is in **data** and **data_time** (contains data and data time slot)

- **site** : *(string)* site name
- **chain** : *(string)* chain name
- **type** : *(string)*
  - "EX" : Entries and Exits
  - "E" : Only Entries
  - "X" : Only Exits
- **resolution** : *(integer)* file resolution in seconds
- **date_begin_end** : *(array[string,string])* file begin and end date (end is the last parsed from original file but no infos is filled by 0 values)
- **opening_hour** : *(array[array[string,string],...])* file opening periods
- **data_channel** : *(array[string,...])*
  -channel name (position in array is channel id)
- **data_channel_access** : *(array[integer,...])*
  -channel id that are access
- **data_time** : *(array[string,...])*
  -date-time resolution
- **data** : *(array[array[int,...]])*
  -all data grouped by *data_time*, and ordered like : E1,X1,E2,X2,...

### Example 1

*A full file with precision set to 28800s (8h), 5 channels, for the whole day 2017-12-12, site open at 08h30 close at 18h30 and is closed between 12h30 and 14h* :

```json
{
   "site":"bat9",
   "chain":"EURECAM",
   "type":"EX",
   "resolution":28800,
   "begin_end":[
      "2017-12-12 08:30:00",
      "2017-12-12 18:30:00"
   ],
   "opening_hour":[
     ["08:30:00","12:30:00"],
     ["14:00:00","18:30:00"]
   ],
   "data_channel":[
      "V3BlancLABO",
      "V3NoirEntréeEU",
      "iCpxEntréeEUR",
      "CPXSPOTArche",
      "TinyCountArche"
   ],
   "data_channel_access":[
      2
   ],
   "data_time":[
      "2017-12-12 00:00:00",
      "2017-12-12 08:00:00",
      "2017-12-12 16:00:00"
   ],
   "data":[
      [1,8,1,1,3,0,0,5,0,0],
      [0,0,72,81,0,0,0,0,0,0],
      [0,0,24,37,0,0,0,0,0,0]
   ]
}
```
