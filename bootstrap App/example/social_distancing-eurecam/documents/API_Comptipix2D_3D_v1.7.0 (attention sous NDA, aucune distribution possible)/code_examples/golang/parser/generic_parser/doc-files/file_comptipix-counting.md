# Comptipix counting file definition

Store counting timestamped data for 1 day.

## File name

Filename : *YYYYMMDD.csv*, exemple : **20180102.csv** for 2018 Jan 02

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
- Channels definition :
  - Channel number (start at 1)
  - Name of the channel
  - Type of the channel :
    - **"acces"** : for entrance or exit to the site
    - **"passage"** : for crossing inside a site
    - *empty* : the channel is disabled

- Column definitions, should be discarded.

**Note 1 :** A valid site name and chain name contains :

- 2 columns (site and chain names)
- A valid site name and chain name must preceed *fichier de comptage v2* text

**Note 2 :** A valid channel definition must contains :

- 3 columns (number, name and type)
- first column must be a number bigger than 0

**Note 3 :** header can be repeated anywhere in file :

- in case of header repetition the last header should be considered as the good header

### File content (data)

Agregation time is configurable from **10 seconds**  to **1h or more**, field are :

1. *date DD/MM/YYYY* or ISO format *YYYY-MM-DD*
2. *time HH:MM:SS*
3. *channel 1, entries*
4. *channel 1, exits*
5. *channel 2, entries*
6. *channel 2, exits*
7. ... etc, *there is no hard limit in file format* ...

**Note 1 :** A valid data line must contains :

- at least 4 columns (date, hour, entries, exits)
- first column must be a valid date and the date must be the same as the filename's
- second column must be a valid hour

**Note 2 :** file can contains multiple data with same timestamp :

- If many data have the same timestamp, parser should sum data.

### Example 1

```csv
Site-Name,Chain-Name
fichier de comptage v2
1,Entree,passage
2,stereo,acces
3,Testvtrine,passage
4,c1000,passage
5,openspace,passage
6,Laterale,passage
7,test cptx v3,passage
8,test affix 2,passage
Date,Heure,E1,S1,E2,S2,E3,S3,E4,S4,E5,S5,E6,S6,E7,S7,E8,S8
12/02/2018,00:00:00,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
12/02/2018,00:01:00,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
12/02/2018,00:02:00,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
12/02/2018,00:03:00,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
... blahblahblah ...
22/05/2013,18:35:00,1,0,1,0,0,0,2,0,2,0,1,0,1,0,0,0
22/05/2013,18:36:00,0,0,1,2,0,0,0,0,0,0,0,0,1,2,0,0
22/05/2013,18:37:00,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0
22/05/2013,18:38:00,1,0,1,1,0,0,1,2,1,2,0,0,1,1,0,0
22/05/2013,18:39:00,0,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0
22/05/2013,18:38:00,0,0,3,2,0,0,0,6,0,5,0,0,0,0,0,0
2013-05-22,18:40:00,0,0,1,0,0,0,0,1,0,1,0,0,1,0,0,0
2013-05-22,18:45:00,0,0,1,0,0,0,0,1,0,1,0,0,1,0,0,0
...
```

 → in this example :

- there is double line for timestamp 18:38:00 → parser should sum this data
- date has been changed to ISO *YYYY-MM-DD* at 18:40 format → parser should manage DD/MM/YYYY and YYYY-MM-DD format
- resolution has changed to 5 min at 18:39:00

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

If many data have the same timestamp, you should sum data.
