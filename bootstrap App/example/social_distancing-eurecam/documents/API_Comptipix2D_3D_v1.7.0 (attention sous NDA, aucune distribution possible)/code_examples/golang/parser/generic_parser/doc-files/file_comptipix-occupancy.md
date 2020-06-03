# Comptipix occupancy file definition

Store occupancy timestamped data for 1 day.

## File name

Filename : *YYYYMMDD_presence.csv*, exemple : **20180102_presence.csv** for 2018 Jan 02

1. YYYY = year (4 digits), exemple : 2018
2. MM = month (2 digits), exemple : 01
3. DD = day (2 digits), exemple : 02

- Extension is **csv**

On the device SD Card, the file is inside a sub directory : *DATA/YYYY/MM/*

- DATA = data directory
- YYYY = year (4 digits)
- MM = month (2 digits)

## File content

It is a CSV file (without quote for text, name shouldn't contains comma), one file per day is generated.
File contains :

1. A header part
2. Timestamped data

File internal format :

- Field Separator is 1 comma ','
- New line char is **LF**
- Encoding is UTF8
- Counting resolution (time slot data) may vary from 10s to 6h or more (always < 24h)

### File header

Header contains :

- Site name and chain name definition (site and chain comma separated)
- Header text (**"fichier de comptage v2"**)
- Header text2 (**"Date,Heure,E,S,P,C+,C-"**)

**Note 1 :** A valid site name and chain name contains :

- 2 columns (site and chain names)
- A valid site name and chain name must preceed *fichier de comptage v2* text

**Note 3 :** header can be repeated anywhere in file :

- in case of header repetition the last header should be considered as the good header

### File content (data)

Agregation time is configurable from **10 seconds**  to **1h or more**, field are :

1. *date DD/MM/YYYY* or ISO format *YYYY-MM-DD*
2. *time HH:MM:SS*
3. *total, entries*
4. *total, exits*
5. *total, occupancy* -> Note: generic_parser will clip occupancy to never be < 0 (in this case Correction+ will be increased)
6. *timestamp, correction+* -> Note: generic_parser may write C+ > 0 if occupancy computed (E-X) is < 0
7. *timestamp, correction-* -> Note: generic_parser generate always C- = 0

**Note 1 :** A valid data line must contains :

- at least 7 columns (date, hour, entries, exits, occupancy, coorection+, correction-)
- first column must be a valid date and the date must be the same as the filename's
- second column must be a valid hour

**Note 2 :** file can contains multiple data with same timestamp :

- If many data have the same timestamp, parser should sum data.

### Example 1

```csv
bat9,Chain_name
fichier de presence v1
Date,Heure,E,S,P,C+,C-
12/12/2017,00:00:00,0,0,0,0,0
12/12/2017,00:30:00,0,0,0,0,0
12/12/2017,01:00:00,0,0,0,0,0
12/12/2017,01:30:00,0,0,0,0,0
12/12/2017,02:00:00,0,0,0,0,0
12/12/2017,02:30:00,0,0,0,0,0
12/12/2017,03:00:00,0,0,0,0,0
12/12/2017,03:30:00,0,0,0,0,0
12/12/2017,04:00:00,0,0,0,0,0
12/12/2017,04:30:00,0,0,0,0,0
12/12/2017,05:00:00,0,0,0,0,0
12/12/2017,05:30:00,0,0,0,0,0
12/12/2017,06:00:00,0,0,0,0,0
12/12/2017,06:30:00,0,0,0,0,0
12/12/2017,07:00:00,0,0,0,0,0
12/12/2017,07:30:00,0,0,0,0,0
12/12/2017,08:00:00,0,2,0,2,0
12/12/2017,08:30:00,0,3,0,1,0
12/12/2017,09:00:00,0,3,0,0,0
12/12/2017,09:30:00,7,7,0,0,0
12/12/2017,10:00:00,10,9,1,0,0
12/12/2017,10:30:00,11,10,1,0,0
12/12/2017,11:00:00,15,14,1,0,0
12/12/2017,13:30:00,16,14,2,0,0
12/12/2017,14:00:00,16,14,2,0,0
12/12/2017,14:30:00,16,14,2,0,0
12/12/2017,15:00:00,16,14,2,0,0
12/12/2017,15:30:00,16,14,2,0,0
12/12/2017,16:00:00,20,14,6,0,0
12/12/2017,16:30:00,20,18,2,0,0
12/12/2017,17:00:00,20,18,2,0,0
12/12/2017,17:30:00,25,18,7,0,0
12/12/2017,18:00:00,25,18,6,0,1
12/12/2017,18:30:00,26,18,7,0,0
12/12/2017,18:00:00,26,18,7,0,0
12/12/2017,19:00:00,26,18,7,0,0
2017-12-12,19:30:00,26,27,0,0,0
2017-12-12,20:00:00,26,27,0,0,0
2017-12-12,20:30:00,26,27,0,0,0
2017-12-12,21:00:00,26,27,0,0,0
2017-12-12,21:30:00,26,27,0,0,0
2017-12-12,22:00:00,26,27,0,0,0
2017-12-12,22:30:00,26,27,0,0,0
2017-12-12,23:00:00,26,27,0,0,0
2017-12-12,23:00:05,26,27,0,0,0
2017-12-12,23:00:10,26,27,0,0,0
2017-12-12,23:00:15,26,27,0,0,0
2017-12-12,23:00:20,26,27,0,0,0
2017-12-12,23:00:25,26,27,0,0,0
2017-12-12,23:00:30,26,27,0,0,0
2017-12-12,23:00:35,26,27,0,0,0
2017-12-12,23:00:40,26,27,0,0,0
2017-12-12,23:00:45,26,27,0,0,0
2017-12-12,23:00:50,26,27,0,0,0
2017-12-12,23:00:55,0,0,0,0,0
```

 → in this example :

- there is a data gap : 11:00:00 to 13:30:00 (maybe device was powered off, or used changed date or timeslot resolution)
- there is double line for timestamp 18:00:00 → parser should keep the data containing highest Occupancy
- date has been changed to ISO *YYYY-MM-DD* at 19:30 format → parser should manage DD/MM/YYYY and YYYY-MM-DD format
- resolution has changed to 5 min at 23:00:00

### Explanations

Timestamp definition depends the data resolution. With an 1 minute resolution, timestamp 8:00:00 sums countings from 8:00:00 to 8:00:59
With a 15 minutes writing period, timestamp 8:00:00 sums countings from 8:00:00 to 8:14:59

Number of columns depends on number of channels : columns = 2 + 2 x channel

Because files are written by small device that never re-read what's already written, so file can contains :

- multiple line for same timestamp (if hour changed by DST or user)
- header can be repeated inside file (if user change config during day)
- file prescision can change inside the file (if user change config during day)

There may be some gaps in data or multiple data with same timestamp, if :

- device was powered off
- null lines are skipped
- change in hour due to DST
- change in hour due to synchronization or manual adjustement

If many data have the same timestamp, you should keep the data containing highest Occupancy.
