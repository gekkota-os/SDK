# Principle

Generation of a xml file human friendly (ie human readable)
Precision (counting data resolution) may vary

## File name

Filename : *YYYYMMDD.xml*, exemple : **20180102.xml** for 2018 Jan 02

1. 2018 *year*
2. 01 *month*
3. 02 *day*

Extension is **xml**

## File Content

- All date are ISO : *YYYY-MM-DD HH:MM:SS*
- Encoding is UTF8
- New line is **LF**
- Counting resolution (time slot data) may vary from 10s to 6h or more (always < 24h)

### File Header and content

File contains header info (in xml) :

- **site** : site name
- **chain** : chain name
- **type** :
  - "EX" : Entries and Exits
  - "E" : Only Entries
  - "X" : Only Exits
- **resolution** : file resolution in seconds
- **date begin+end** : file begin and end date (end is the last parsed from original file but no infos is filled by 0 values)
- **opening_hour** : file opening periods
- **channels** :
  -channel name (as value)
  -channel *id* (attribut)
  -channel type : *crossing* or *access* (attribut)

And counting data grouped by resolution (*data_group* element containing *data_chan*) :

- **data_group** with resolution formated ISO (YYYY-MM-DD HH:MM:SS) as attribut
- **data_can** value is counting value, attibuts are :
  - *chan_id* channel id owning this data
  - *type* can be :
    - "E" : counting info are entries
    - "X" : counting info area exits

### Example 1

*A full file with precision set to 28800s (8h), 5 channels, for the whole day 2017-12-12, site open at 08h30 close at 18h30 and is closed between 12h30 and 14h* :

```xml
<?xml version="1.0" encoding="UTF-8"?>
<xml_1 date="2017-12-12 00:00:00">
  <site>bat9</site>
  <chain>EURECAM</chain>
  <type>EX</type>
  <resolution>28800</resolution>
  <date>
    <begin>2017-12-12 08:30:00</begin>
    <end>2017-12-12 18:30:00</end>
  </date>
  <opening_hour>
    <open id="0">08:30:00</open>
    <close id="0">12:30:00</close>
    <open id="1">14:00:00</open>
    <close id="1">18:30:00</close>
  </opening_hour>
  <channel>
    <name>
      <chan id="1" type="crossing">V3BlancLABO</chan>
      <chan id="2" type="access">V3NoirEntréeEU</chan>
      <chan id="3" type="crossing">iCpxEntréeEUR</chan>
      <chan id="4" type="crossing">CPXSPOTArche</chan>
      <chan id="5" type="crossing">TinyCountArche</chan>
    </name>
  </channel>
  <data>
    <data_group time="2017-12-12 00:00:00">
      <data_chan chan_id="1" type="E">0</data_chan>
      <data_chan chan_id="1" type="X">0</data_chan>
      <data_chan chan_id="2" type="E">0</data_chan>
      <data_chan chan_id="2" type="X">0</data_chan>
      <data_chan chan_id="3" type="E">0</data_chan>
      <data_chan chan_id="3" type="X">0</data_chan>
      <data_chan chan_id="4" type="E">0</data_chan>
      <data_chan chan_id="4" type="X">0</data_chan>
      <data_chan chan_id="5" type="E">0</data_chan>
      <data_chan chan_id="5" type="X">0</data_chan>
    </data_group>
    <data_group time="2017-12-12 08:00:00">
      <data_chan chan_id="1" type="E">0</data_chan>
      <data_chan chan_id="1" type="X">0</data_chan>
      <data_chan chan_id="2" type="E">79</data_chan>
      <data_chan chan_id="2" type="X">78</data_chan>
      <data_chan chan_id="3" type="E">0</data_chan>
      <data_chan chan_id="3" type="X">0</data_chan>
      <data_chan chan_id="4" type="E">0</data_chan>
      <data_chan chan_id="4" type="X">0</data_chan>
      <data_chan chan_id="5" type="E">0</data_chan>
      <data_chan chan_id="5" type="X">0</data_chan>
    </data_group>
    <data_group time="2017-12-12 16:00:00">
      <data_chan chan_id="1" type="E">0</data_chan>
      <data_chan chan_id="1" type="X">0</data_chan>
      <data_chan chan_id="2" type="E">27</data_chan>
      <data_chan chan_id="2" type="X">31</data_chan>
      <data_chan chan_id="3" type="E">0</data_chan>
      <data_chan chan_id="3" type="X">0</data_chan>
      <data_chan chan_id="4" type="E">0</data_chan>
      <data_chan chan_id="4" type="X">0</data_chan>
      <data_chan chan_id="5" type="E">0</data_chan>
      <data_chan chan_id="5" type="X">0</data_chan>
    </data_group>
  </data>
</xml_1>
```
