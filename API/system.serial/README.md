nsISystemAdapterSerial interface Reference
==========================================

Public Attributes
-----------------

-   readonly attribute nsISystemAdapterSerial registered

<!-- -->

-   readonly attribute unsigned long capabilities

<!-- -->

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

<!-- -->

-   const unsigned long DIRECTION\_IN

<!-- -->

-   const unsigned long DIRECTION\_OUT

<!-- -->

-   const unsigned long BAUD50

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

<!-- -->

-   const unsigned long PARITY\_NONE

<!-- -->

-   const unsigned long PARITY\_ODD

<!-- -->

-   const unsigned long PARITY\_EVEN

<!-- -->

-   const unsigned long PARITY\_MARK

<!-- -->

-   const unsigned long PARITY\_SPACE

<!-- -->

-   const unsigned long STOPBIT\_1

<!-- -->

-   const unsigned long STOPBIT\_1\_5

<!-- -->

-   const unsigned long STOPBIT\_2

<!-- -->

-   const unsigned long FLOWCONTROL\_OFF

<!-- -->

-   const unsigned long FLOWCONTROL\_HARDWARE

<!-- -->

-   const unsigned long FLOWCONTROL\_SOFTWARE

<!-- -->

-   const unsigned long FLOWCONTROL\_WINDOWED

<!-- -->

-   attribute boolean BREAK

<!-- -->

-   readonly attribute boolean DCD

<!-- -->

-   readonly attribute boolean CTS

<!-- -->

-   readonly attribute boolean DSR

<!-- -->

-   attribute boolean DTR

<!-- -->

-   readonly attribute boolean RING

<!-- -->

-   attribute boolean RTS

<!-- -->

-   attribute boolean window

<!-- -->

-   readonly attribute boolean isTQEmpty

<!-- -->

-   attribute boolean recieveIsBlocking

<!-- -->

-   attribute long recieveThreshold

<!-- -->

-   attribute long recieveTimeout

<!-- -->

-   const unsigned long RECIEVE\_MODE\_SYNC

<!-- -->

-   const unsigned long RECIEVE\_MODE\_ASYNC

<!-- -->

-   attribute long recieveMode

<!-- -->

-   readonly attribute nsIInputStream inputStream

<!-- -->

-   readonly attribute nsIOutputStream outputStream

-   void setConfig ( in unsigned long aDirection, in unsigned long aBaudrate, in unsigned long aCharSize, in unsigned long aParity, in unsigned long aNbStopBits, in unsigned long aFlowControl)

<!-- -->

-   void getConfig ( out unsigned long aDirection, out unsigned long aBaudrate, out unsigned long aCharSize, out unsigned long aParity, out unsigned long aNbStopBits, out unsigned long aFlowControl)

<!-- -->

-   void sendBreak ( in unsigned long aPeriod)

<!-- -->

-   void drainTQ ( )

<!-- -->

-   void clearRQ ( )

<!-- -->

-   void clearTQ ( )

<!-- -->

-   void addListener ( in nsISystemSerialListener aListener)

<!-- -->

-   void removeListener ( in nsISystemSerialListener aListener)

<!-- -->

-   void open ( )

<!-- -->

-   void close ( )

Detailed Description
--------------------

The nsISystemAdapterSerial Interface allows to manage serial port. You can find an example writing a command and asynchronously reading the result [here.](example1.html)

Member Data Documentation
-------------------------

### readonly attribute nsISystemAdapterSerial nsISystemAdapterSerial::registered

Retreive an secondary instance of the adapter that work with registrered values (for exemple preferences) otherwise live values.

### readonly attribute unsigned long nsISystemAdapterSerial::capabilities

Capabilities bits flag. It can be a union of the different flags.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_DIR\_IN

GPIO input possible.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_DIR\_OUT

GPIO output possible.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_HALF\_DUPLEX

GPIO input or output in the same time only.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_FULL\_DUPLEX

GPIO input+output in the same time possible.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_PHY\_TTL

TTL possible.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_PHY\_RS232

RS232 possible.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_PHY\_RS422\_485

RS422 (if full duplex) or RS485 (if half duplex) possible.

### const unsigned long nsISystemAdapterSerial::SERIAL\_CAP\_HARD\_FLOW\_CTL

Hardware flow control possible.

### const unsigned long nsISystemAdapterSerial::DIRECTION\_IN

Direction bits flag: serial port can be reading.

### const unsigned long nsISystemAdapterSerial::DIRECTION\_OUT

Direction bits flag: serial port can be writing.

### const unsigned long nsISystemAdapterSerial::BAUD50

Example of baudrate value, see "Public Attributes" for every possible value.

### const unsigned long nsISystemAdapterSerial::PARITY\_NONE

Parity scheme value: none.

### const unsigned long nsISystemAdapterSerial::PARITY\_ODD

Parity scheme value: odd.

### const unsigned long nsISystemAdapterSerial::PARITY\_EVEN

Parity scheme value: even.

### const unsigned long nsISystemAdapterSerial::PARITY\_MARK

Parity scheme value: mark.

### const unsigned long nsISystemAdapterSerial::PARITY\_SPACE

Parity scheme value: space.

### const unsigned long nsISystemAdapterSerial::STOPBIT\_1

1 stop bit.

### const unsigned long nsISystemAdapterSerial::STOPBIT\_1\_5

1.5 stop bit.

### const unsigned long nsISystemAdapterSerial::STOPBIT\_2

2 stop bits.

### const unsigned long nsISystemAdapterSerial::FLOWCONTROL\_OFF

Flow control: no handshaking.

### const unsigned long nsISystemAdapterSerial::FLOWCONTROL\_HARDWARE

Flow control: hardware handshaking (RTS/CTS).

### const unsigned long nsISystemAdapterSerial::FLOWCONTROL\_SOFTWARE

Flow control: software handshaking (XON/XOFF).

### const unsigned long nsISystemAdapterSerial::FLOWCONTROL\_WINDOWED

Flow control: hardware windowed.

### attribute boolean nsISystemAdapterSerial::BREAK

Set or get BREAK signal value.

### readonly attribute boolean nsISystemAdapterSerial::DCD

Get DCD signal value.

### readonly attribute boolean nsISystemAdapterSerial::CTS

Get CTS signal value.

### readonly attribute boolean nsISystemAdapterSerial::DSR

Get DSR signal value.

### attribute boolean nsISystemAdapterSerial::DTR

SetDTR and GetDTR Data Terminal Ready.

### readonly attribute boolean nsISystemAdapterSerial::RING

Get Ring.

### attribute boolean nsISystemAdapterSerial::RTS

SetRTS and GetRTS request to send.

### attribute boolean nsISystemAdapterSerial::window

Set and get window state for RS485/RS422 writing.

### readonly attribute boolean nsISystemAdapterSerial::isTQEmpty

True if transmitting queue is empty, false otherwise.

### attribute boolean nsISystemAdapterSerial::recieveIsBlocking

Tell if reception of characters on the serial port is in blocking mode or not. When the reading is in blocking mode, nsIInputStream::read do return only if the number of characters to read is reached (attribute recieveIsThreshold) or if the limit of time is exceeded. Otherwise it does immediately return.

### attribute long nsISystemAdapterSerial::recieveThreshold

When the reading is in blocking mode (attribute recieveIsBlocking), this attribute provides the minimal amount of characters to read before returning.

### attribute long nsISystemAdapterSerial::recieveTimeout

If the reading is in blocking mode (attribute recieveIsBlocking), this attribute defines the limit value of waiting time (in milliseconds) before the reading returns. The default value is zero which matches to infinite waiting.

### const unsigned long nsISystemAdapterSerial::RECIEVE\_MODE\_SYNC

Synchronous reception of characters. When synchronously receiving, we use nsIInputStream::read of nsIInputStream component retrieved from the inputStream attribute.

### const unsigned long nsISystemAdapterSerial::RECIEVE\_MODE\_ASYNC

Asynchronous reception of characters. In this mode, it is possible to:

-   Call nsIAsyncInputStream::asyncWait from nsIAsyncInputStream component retrieved from inputStream attribute.

-   Call nsISystemAdapterSerial::addListener to be notified about reception of characters, and we can read characters provided from nsISystemSerialListener::onDataAvailable.

### attribute long nsISystemAdapterSerial::recieveMode

Reception mode for characters received by the serial port This attribute must be specified before the call of the open method in order to be taken in count.

### readonly attribute nsIInputStream nsISystemAdapterSerial::inputStream

Object for read the characters received by the serial port.

### readonly attribute nsIOutputStream nsISystemAdapterSerial::outputStream

Object for write some characters to the serial port.

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
<td align="left"><p>Direction of serial port (DIRECTION_IN|DIRECTION_OUT)</p></td>
</tr>
<tr class="even">
<td align="left">aBaudrate</td>
<td align="left"><p>Baurate of serial port (use one value of type BAUDXXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aCharSize</td>
<td align="left"><p>Size of char in serial port (supported values 5,6,7,8)</p></td>
</tr>
<tr class="even">
<td align="left">aParity</td>
<td align="left"><p>Parity of serial port (use one value of type PARITY_XXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aNbStopBits</td>
<td align="left"><p>Number of stop bits of serial port (use one value of type STOPBIT_XXX)</p></td>
</tr>
<tr class="even">
<td align="left">aFlowControl</td>
<td align="left"><p>Type of flow control (use one value of type FLOWCONTROL_XXX)</p></td>
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
<td align="left"><p>Direction of serial port (DIRECTION_IN|DIRECTION_OUT)</p></td>
</tr>
<tr class="even">
<td align="left">aBaudrate</td>
<td align="left"><p>Baurate of serial port (value of type BAUDXXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aCharSize</td>
<td align="left"><p>Size of char in serial port (supported values 5,6,7,8)</p></td>
</tr>
<tr class="even">
<td align="left">aParity</td>
<td align="left"><p>Parity of serial port (value of type PARITY_XXX)</p></td>
</tr>
<tr class="odd">
<td align="left">aNbStopBits</td>
<td align="left"><p>Number of stop bits of serial port (value of type STOPBIT_XXX)</p></td>
</tr>
<tr class="even">
<td align="left">aFlowControl</td>
<td align="left"><p>Type of flow control (use value FLOWCONTROL_XXX)</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterSerial::sendBreak (in unsigned long aPeriod)
-----------------------------------------------------------------

Send a break signal for a period of time in millisecond.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aPeriod</td>
<td align="left"><p>Period of time in millisecond</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterSerial::drainTQ ()
---------------------------------------

Wait until all data in transmitting queue has been transmitted.

void nsISystemAdapterSerial::clearRQ ()
---------------------------------------

Clear revieve queue.

void nsISystemAdapterSerial::clearTQ ()
---------------------------------------

Clear transmit queue.

void nsISystemAdapterSerial::addListener (in nsISystemSerialListener aListener)
-------------------------------------------------------------------------------

Add a listener called when a state of one signal change or when data become available.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aListener</td>
<td align="left"><p>The serial adapter listener</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterSerial::removeListener (in nsISystemSerialListener aListener)
----------------------------------------------------------------------------------

Remove a listener called when a state of one signal change or when data become available.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aListener</td>
<td align="left"><p>The serial adapter listener</p></td>
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

The nsISystemSerialListener interface is a listener for system adapter serial.

void nsISystemSerialListener::onCTSChanged (in boolean aNewValue)
-----------------------------------------------------------------

Method called when CTS signal change.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>New value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onDSRChanged (in boolean aNewValue)
-----------------------------------------------------------------

Method called when DSR signal change.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>New value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onRINGChanged (in boolean aNewValue)
------------------------------------------------------------------

Method called when RING signal change.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>New value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onDCDChanged (in boolean aNewValue)
-----------------------------------------------------------------

Method called when DCD signal change.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>New value of signal</p></td>
</tr>
</tbody>
</table>

void nsISystemSerialListener::onDataAvailable (in nsIInputStream aInputStream)
------------------------------------------------------------------------------

Method called when data available on serial port. Use nsIInputStream::available to get number of characters available. Use nsIInputStream::read to read characters.

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


