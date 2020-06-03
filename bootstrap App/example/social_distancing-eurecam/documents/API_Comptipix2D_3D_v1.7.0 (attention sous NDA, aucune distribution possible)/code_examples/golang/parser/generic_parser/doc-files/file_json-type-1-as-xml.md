# Principle

Generation of a xml file containing same data structure as json type 1
See json type 1 for explanation

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

See json type 1 for explanation

### Example 1

*A full file with precision set to 28800s (8h), 5 channels, for the whole day 2017-12-12* :

```xml
<?xml version="1.0" encoding="UTF-8"?>
<OutPutJson_1>
   <site>bat9</site>
   <chain>EURECAM</chain>
   <type>EX</type>
   <resolution>28800</resolution>
   <begin_end>2017-12-13 00:00:00</begin_end>
   <begin_end>2017-12-13 00:00:00</begin_end>
   <data_channel>V3BlancLABO</data_channel>
   <data_channel>V3NoirEntréeEU</data_channel>
   <data_channel>iCpxEntréeEUR</data_channel>
   <data_channel>CPXSPOTArche</data_channel>
   <data_channel>TinyCountArche</data_channel>
   <data_channel_access>2</data_channel_access>
   <data_time>2017-12-13 00:00:00</data_time>
   <data_time>2017-12-13 08:00:00</data_time>
   <data_time>2017-12-13 16:00:00</data_time>
   <data>1</data>
   <data>8</data>
   <data>1</data>
   <data>1</data>
   <data>3</data>
   <data>0</data>
   <data>0</data>
   <data>5</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>72</data>
   <data>81</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>24</data>
   <data>37</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
   <data>0</data>
</OutPutJson_1>
```
