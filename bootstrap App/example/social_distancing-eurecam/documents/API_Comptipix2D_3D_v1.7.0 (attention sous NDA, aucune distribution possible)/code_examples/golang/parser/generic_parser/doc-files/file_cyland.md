# Principle

Generation of a file with same format as *cyland*
Precision time set to 30 min

## File name

Filename : *YYYYMMDD.csv*, exemple: **20160515.csv** for 2016 May 15

1. 2016 *year*
2. 05 *month*
3. 15 *day*

Extension is **csv**

## File content

### File header

There is no header

### File content (data)

Agregation time is **30 min**, field are :

1. *date DDMMYYYY*
2. *time HH:MM*
3. *error code 1=OK*
4. *entrances*
5. *exits*

- Field Separator is 1 semicolon ';'
- New line char is **LF** → TODO:To verify
- Counting resolution (time slot data) is 30 min

### Example 1

```csv
...
25092017;08:30;1;0;0
25092017;09:00;1;0;0
25092017;09:30;1;0;0
25092017;10:00;1;1;0
25092017;10:30;1;3;1
25092017;11:00;1;2;1
25092017;11:30;1;1;3
25092017;12:00;1;6;2
25092017;12:30;1;0;4
25092017;13:00;1;5;1
25092017;13:30;1;2;5
25092017;14:00;1;4;2
25092017;14:30;1;3;3
25092017;15:00;1;1;2
...
```
