nsISystemAdapterSerial interface Reference
==========================================

-   const unsigned long SERIAL\_CAP\_DIR\_IN

<!-- -->

-   const unsigned long SERIAL\_CAP\_DIR\_OUT

<!-- -->

-   const unsigned long SERIAL\_CAP\_HALF\_DUPLEX

<!-- -->

-   const unsigned long SERIAL\_CAP\_FULL\_DUPLEX

<!-- -->

-   const unsigned long SERIAL\_CAP\_PHY\_TTL

<!-- -->

-   const unsigned long SERIAL\_CAP\_PHY\_RS232

<!-- -->

-   const unsigned long SERIAL\_CAP\_PHY\_RS422\_485

<!-- -->

-   const unsigned long SERIAL\_CAP\_HARD\_FLOW\_CTL

-   const unsigned long DIRECTION\_IN

<!-- -->

-   const unsigned long DIRECTION\_OUT

    *Serial port can be writing.*

-   \*const unsigned long BAUD50

<!-- -->

-   const unsigned long BAUD75

<!-- -->

-   const unsigned long BAUD110

<!-- -->

-   const unsigned long BAUD134

<!-- -->

-   const unsigned long BAUD150

<!-- -->

-   const unsigned long BAUD200

<!-- -->

-   const unsigned long BAUD300

<!-- -->

-   const unsigned long BAUD600

<!-- -->

-   const unsigned long BAUD1200

<!-- -->

-   const unsigned long BAUD1800

<!-- -->

-   const unsigned long BAUD2400

<!-- -->

-   const unsigned long BAUD4800

<!-- -->

-   const unsigned long BAUD9600

<!-- -->

-   const unsigned long BAUD14400

<!-- -->

-   const unsigned long BAUD19200

<!-- -->

-   const unsigned long BAUD38400

<!-- -->

-   const unsigned long BAUD36000

<!-- -->

-   const unsigned long BAUD56000

<!-- -->

-   const unsigned long BAUD57600

<!-- -->

-   const unsigned long BAUD76800

<!-- -->

-   const unsigned long BAUD115200

<!-- -->

-   const unsigned long BAUD128000

<!-- -->

-   const unsigned long BAUD230400

<!-- -->

-   const unsigned long BAUD256000

<!-- -->

-   const unsigned long BAUD460800

-   const unsigned long PARITY\_NONE

<!-- -->

-   const unsigned long PARITY\_ODD

<!-- -->

-   const unsigned long PARITY\_EVEN

<!-- -->

-   const unsigned long PARITY\_MARK

<!-- -->

-   const unsigned long PARITY\_SPACE

-   const unsigned long STOPBIT\_1

<!-- -->

-   const unsigned long STOPBIT\_1\_5

    *1.5 stopbit*

<!-- -->

-   const unsigned long STOPBIT\_2

    *2 stopbit*

-   const unsigned long FLOWCONTROL\_OFF

<!-- -->

-   const unsigned long FLOWCONTROL\_HARDWARE

    *Hardware handshaking (RTS/CTS)*

<!-- -->

-   const unsigned long FLOWCONTROL\_SOFTWARE

    *Software handshaking (XON/XOFF)*

<!-- -->

-   const unsigned long FLOWCONTROL\_WINDOWED

    *Hardware windowed.*

-   attribute boolean BREAK

    *Set ot get BREAK signal value.*

<!-- -->

-   readonly attribute boolean DCD

    *Get DCD signal value.*

<!-- -->

-   readonly attribute boolean CTS

    *Get CTS signal value.*

<!-- -->

-   readonly attribute boolean DSR

    *Get DSR signal value.*

<!-- -->

-   attribute boolean DTR

    *SetDTR and GetDTR Data Terminal Ready.*

<!-- -->

-   readonly attribute boolean RING

    *GetRing.*

<!-- -->

-   attribute boolean RTS

    *SetRTS and GetRTS request to send.*

<!-- -->

-   attribute boolean window

    *set and get window state for RS485/RS422 writing*

<!-- -->

-   readonly attribute boolean isTQEmpty

    *True if transmiting queue is empty, false otherwise.*

<!-- -->

-   void sendBreak ( in unsigned long aPeriod)

<!-- -->

-   void drainTQ ( )

    *Wait until all data in transmiting queue has been transmitted.*

<!-- -->

-   void clearRQ ( )

    *clear revieve queue*

<!-- -->

-   void clearTQ ( )

    *clear transmit queue*

<!-- -->

-   void addListener ( in nsISystemSerialListener aListener)

<!-- -->

-   void removeListener ( in nsISystemSerialListener aListener)

-   const unsigned long RECIEVE\_MODE\_SYNC

<!-- -->

-   const unsigned long RECIEVE\_MODE\_ASYNC

<!-- -->

-   attribute long recieveMode

Public Attributes
-----------------

-   readonly attribute unsigned long capabilities

<!-- -->

-   attribute boolean recieveIsBlocking

<!-- -->

-   attribute long recieveThreshold

<!-- -->

-   attribute long recieveTimeout

<!-- -->

-   readonly attribute nsIInputStream inputStream

<!-- -->

-   readonly attribute nsIOutputStream outputStream

    *Object for write some characters to to the serial port.*

-   void setConfig ( in unsigned long aDirection, in unsigned long aBaudrate, in unsigned long aCharSize, in unsigned long aParity, in unsigned long aNbStopBits, in unsigned long aFlowControl)

<!-- -->

-   void getConfig ( out unsigned long aDirection, out unsigned long aBaudrate, out unsigned long aCharSize, out unsigned long aParity, out unsigned long aNbStopBits, out unsigned long aFlowControl)

<!-- -->

-   void open ( )

    *Open serial port.*

<!-- -->

-   void close ( )

    *Close serial port.*

Member Data Documentation
-------------------------

### readonly attribute unsigned long nsISystemAdapterSerial::capabilities

Capabilities flags. Can be an union of the different flags

### attribute boolean nsISystemAdapterSerial::recieveIsBlocking

Tell if reception of characters on the serial port is in blocking mode or not. When the reading is in blocking mode, nsIInputStream::read do return only if the number of characters to read is reached (attribute recieveIsThreshold) or if the limit of time is exceeded. Otherwise it does immediatly return.

### attribute long nsISystemAdapterSerial::recieveThreshold

When the reading is in blocking mode (attribute recieveIsBlocking), this attribute provide the minimal amount of characters to read before returning.

### attribute long nsISystemAdapterSerial::recieveTimeout

If the reading is in blocking mode (attribute recieveIsBlocking), this attribute defines the limit value of waiting time (in milliseconds) before the reading returns. The default value is zero which matchs to infinite waiting.

### readonly attribute nsIInputStream nsISystemAdapterSerial::inputStream

Object for read the characters received by the serial port

void nsISystemAdapterSerial::setConfig (in unsigned long aDirection, in unsigned long aBaudrate, in unsigned long aCharSize, in unsigned long aParity, in unsigned long aNbStopBits, in unsigned long aFlowControl)
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Set configuration of serial port. This method must be called before open method.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aDirection</td>
<td align="left"><p>direction of serial port (DIRECTION_IN|DIRECTION_OUT)</p></td>
</tr>
<tr class="even">
<td align="left">aBaudrate</td>
<td align="left"><p>baurate of serial port (use one value of type BAUDXXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aCharSize</td>
<td align="left"><p>size of char in serial port (supported values 5,6,7,8)</p></td>
</tr>
<tr class="even">
<td align="left">aParity</td>
<td align="left"><p>parity of serial port (use one value of type PARITY_XXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aNbStopBits</td>
<td align="left"><p>number of stop bits of serial port (use one value of type STOPBIT_XXX)</p></td>
</tr>
<tr class="even">
<td align="left">aFlowControl</td>
<td align="left"><p>type of flow control (use one value of type FLOWCONTROL_XXX)</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterSerial::getConfig (out unsigned long aDirection, out unsigned long aBaudrate, out unsigned long aCharSize, out unsigned long aParity, out unsigned long aNbStopBits, out unsigned long aFlowControl)
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Get configuration of serial port.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aDirection</td>
<td align="left"><p>direction of serial port (DIRECTION_IN|DIRECTION_OUT)</p></td>
</tr>
<tr class="even">
<td align="left">aBaudrate</td>
<td align="left"><p>baurate of serial port (value of type BAUDXXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aCharSize</td>
<td align="left"><p>size of char in serial port (supported values 5,6,7,8)</p></td>
</tr>
<tr class="even">
<td align="left">aParity</td>
<td align="left"><p>parity of serial port (value of type PARITY_XXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aNbStopBits</td>
<td align="left"><p>number of stop bits of serial port (value of type STOPBIT_XXX)</p></td>
</tr>
<tr class="even">
<td align="left">aFlowControl</td>
<td align="left"><p>type of flow control (use value FLOWCONTROL_XXX)</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterSerial::open ()
------------------------------------

Open serial port.
void nsISystemAdapterSerial::close ()
-------------------------------------

Close serial port.
nsISystemSerialListener interface Reference
===========================================

-   void onCTSChanged ( in boolean aNewValue)

<!-- -->

-   void onDSRChanged ( in boolean aNewValue)

<!-- -->

-   void onRINGChanged ( in boolean aNewValue)

<!-- -->

-   void onDCDChanged ( in boolean aNewValue)

<!-- -->

-   void onDataAvailable ( in nsIInputStream aInputStream)

Detailed Description
--------------------

The nsISystemSerialListener interface is a listener for system adapter serial

void nsISystemSerialListener::onCTSChanged (in boolean aNewValue)
-----------------------------------------------------------------

Method called when CTS signal change

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>new value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onDSRChanged (in boolean aNewValue)
-----------------------------------------------------------------

Method called when DSR signal change

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>new value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onRINGChanged (in boolean aNewValue)
------------------------------------------------------------------

Method called when RING signal change

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>new value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onDCDChanged (in boolean aNewValue)
-----------------------------------------------------------------

Method called when DCD signal change

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>new value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onDataAvailable (in nsIInputStream aInputStream)
------------------------------------------------------------------------------

Method called when data available on serial port. Use nsIInputStream::available to get number of carateres available Use nsIInputStream::read to read carateres.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aInputStream</td>
<td align="left"><p>Input stream associate with serial port</p></td>
</tr>
</tbody>
</table>


