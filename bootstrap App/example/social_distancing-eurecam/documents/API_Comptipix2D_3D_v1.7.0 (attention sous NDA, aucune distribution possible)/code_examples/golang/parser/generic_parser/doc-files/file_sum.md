# Principle

A sum file contains just sum (by default only *entry*, but it can be like 'entry,exit,crossing' or only exit, only crossing)  

## File name

Filename : *YYYYMMDD.csv*, exemple : **20180102.csv** for 2018 Jan 02

1. YYYY = year (4 digits), exemple : 2018
2. MM = month (2 digits), exemple : 2018
3. DD = day (2 digits), exemple : 2018

- Extension is **csv**

&rarr; **Same as Eurecam format**

## File content

- Field Separator is comma char ','
- New line char is LF

### Exemple: Only entry

*372 entry this day*
372

### Exemple: Entry and exit

*372 entry + 369 exits this day* :

```csv
372,369
```
