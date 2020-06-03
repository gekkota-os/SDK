# SLCP Area counting file definition

Store counting timestamped data for 1 day.

## File name

Filename : *YYYY/MMDD.csv*, exemple : **2018/0102.csv** for 2018 Jan 02

1. YYYY = year (4 digits), exemple : 2018 (this is a directory)
2. MM = month (2 digits), exemple : 01
3. DD = day (2 digits), exemple : 02

- Extension is **csv**

On the HDD, the file is inside a sub directory : *YYYY/*

## File content

It is a CSV file (without quote for text, name shouldn't contains comma), one file per day is generated.

File internal format :

- Field Separator is 1 comma ','
- New line char is **LF**
- Encoding is UTF8
- Counting resolution (time slot data) may vary from 10s to 6h or more (always < 24h)

### File header

There is no header

### File content (data)

Agregation time is configurable from **10 seconds**  to **1h or more**, field are :

1. *hour HH:MM:SS*
2. *element id*
3. *opening* -> 0: closed, 1 or more: open
4. *occupancy* (occupancy at the minute end)

**Note 1 :** A valid data line must contains :

- at least 4 columns (hour, id, opening, occupancy value)
- first column must be a valid date hour

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
