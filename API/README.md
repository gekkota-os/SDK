#GEKKOTA API
## gpio
This API allows to write to or read from GPIO.  

## idle
The Idle API provides a service to detect when Gekkota is Idle state (that is no HID events on the device during a fixed time).  
It is possible to configure the delay for Idle state and implement what must happen when switching the state.  

## ldap
The LDAP Service allows to send synchronous and asynchronous search queries to an LDAP server and manage results.   

## nfc
On a SMT210 equipment, we can know when a tag enters or leaves the NFC field.  
The NFC System Adapter API provides some interfaces and callbacks to detect and get information (eg. Identifier) about the NFC tag.  

## rs232
The features implemented for rs232 are :  

- Qlite protocol

## side-leds-smt210
This API allows to manage the LEDs color on a SMT210 equipment.  

## smtp
The SMTP API provides a set of interfaces to manage messaging via the SMTP protocol.  

## sql
The SQL API allows to connect to an ODBC database then send synchronous and asynchronous queries.  
It is possible to manage the queries results.  

## udp-socket
The UDP Socket API provides some methods to send UDP datagrams to unicast or multicast addresses.   
It is possible to join and leave a multicast group.  

