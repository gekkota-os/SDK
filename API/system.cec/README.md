nsISystemAdapterCec interface Reference
=======================================

Public Attributes
-----------------

-   readonly attribute bool isSinkConnected

<!-- -->

-   readonly attribute bool isSourceConnected

<!-- -->

-   attribute bool enablePassThrough

<!-- -->

-   readonly attribute bool isActiveSource

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_TV

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_RECORDING\_DEVICE

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_RESERVED

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_TUNER

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_PLAYBACK\_DEVICE

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_AUDIO\_SYSTEM

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_PURE\_CEC\_SWITCH

<!-- -->

-   const short CEC\_DEVICE\_TYPE\_VIDEO\_PROCESSOR

<!-- -->

-   const short CECDEVICE\_UNKNOWN

<!-- -->

-   const short CECDEVICE\_TV

<!-- -->

-   const short CECDEVICE\_RECORDINGDEVICE1

<!-- -->

-   const short CECDEVICE\_RECORDINGDEVICE2

<!-- -->

-   const short CECDEVICE\_TUNER1

<!-- -->

-   const short CECDEVICE\_PLAYBACKDEVICE1

<!-- -->

-   const short CECDEVICE\_AUDIOSYSTEM

<!-- -->

-   const short CECDEVICE\_TUNER2

<!-- -->

-   const short CECDEVICE\_TUNER3

<!-- -->

-   const short CECDEVICE\_PLAYBACKDEVICE2

<!-- -->

-   const short CECDEVICE\_RECORDINGDEVICE3

<!-- -->

-   const short CECDEVICE\_TUNER4

<!-- -->

-   const short CECDEVICE\_PLAYBACKDEVICE3

<!-- -->

-   const short CECDEVICE\_RESERVED1

<!-- -->

-   const short CECDEVICE\_RESERVED2

<!-- -->

-   const short CECDEVICE\_FREEUSE

<!-- -->

-   const short CECDEVICE\_UNREGISTERED

<!-- -->

-   const short CECDEVICE\_BROADCAST

-   void getDevicesTypes ( out unsigned long aLen, out uint8\_t aDeviceTypes)

<!-- -->

-   void getActiveDevices ( out unsigned long aLen, out uint8\_t aLogicAddr)

<!-- -->

-   boolean isActiveDevice ( in uint8\_t aLogicAddr)

<!-- -->

-   boolean isActiveDeviceType ( in uint8\_t aDeviceType)

<!-- -->

-   uint16\_t getPhysicalAddress ( )

<!-- -->

-   void getLogicalAddresses ( out uint8\_t aPrimaryDevAddr, out unsigned long aLen, out uint8\_t aLogicAddr)

<!-- -->

-   void getLogicalAddress ( in uint8\_t aDeviceType, out uint8\_t aLogicAddr)

<!-- -->

-   void sendMsg ( in uint8\_t aSource, in uint8\_t aDest, in ACString aByteMessage, in bool aExpectedReply, in uint8\_t aReplyOpcodeFilter, in nsIGktEventHandler aSendMsgResultEventHandler)

<!-- -->

-   void addObserver ( in nsIGktEventHandler aObserver)

<!-- -->

-   void removeObserver ( in nsIGktEventHandler aObserver)

Detailed Description
--------------------

The XPCOM interface nsISystemAdapterCec is the entry point of the cec adapter.

Member Data Documentation
-------------------------

### readonly attribute bool nsISystemAdapterCec::isSinkConnected

A boolean which says if a sink is currently connected.

### readonly attribute bool nsISystemAdapterCec::isSourceConnected

A boolean which says if a source is currently connected.

### attribute bool nsISystemAdapterCec::enablePassThrough

A boolean which enables or diables the Remote Control Pass Through feature.

### readonly attribute bool nsISystemAdapterCec::isActiveSource

A boolean which says if the local device(s) is(are) the active source in the network.

void nsISystemAdapterCec::getDevicesTypes (\[optional\] out unsigned long aLen, \[array, size\_is(aLen), retval\] out uint8\_t aDeviceTypes)
--------------------------------------------------------------------------------------------------------------------------------------------

Gets the type(s) of device(s) controlled by the adapter.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLen</td>
<td align="left"><p>receives the number of devices types.</p></td>
</tr>
<tr class="even">
<td align="left">aDeviceType</td>
<td align="left"><p>receives an array of devices types.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a boolean (success/failure).

void nsISystemAdapterCec::getActiveDevices (\[optional\] out unsigned long aLen, \[array, size\_is(aLen), retval\] out uint8\_t aLogicAddr)
-------------------------------------------------------------------------------------------------------------------------------------------

The logical addresses of the devices that are active on the bus, including those handled by the adapter.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLen</td>
<td align="left"><p>receives the number of addresses.</p></td>
</tr>
<tr class="even">
<td align="left">aLogicAddr</td>
<td align="left"><p>receives the array of addresses.</p></td>
</tr>
</tbody>
</table>

boolean nsISystemAdapterCec::isActiveDevice (in uint8\_t aLogicAddr)
--------------------------------------------------------------------

Check whether a device is active on the bus.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLogicAddr</td>
<td align="left"><p>the logical address to check.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a boolean (true/false).

boolean nsISystemAdapterCec::isActiveDeviceType (in uint8\_t aDeviceType)
-------------------------------------------------------------------------

Check whether a device of the given type is active on the bus.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aDeviceType</td>
<td align="left"><p>the type to check.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a boolean (true/false).

uint16\_t nsISystemAdapterCec::getPhysicalAddress ()
----------------------------------------------------

Return the physical address of the device(s) that the adapter is controlling

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aPhysicalAddr</td>
<td align="left"><p>the corresponding physical address of the device(s)</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a boolean (success/failure).

void nsISystemAdapterCec::getLogicalAddresses (out uint8\_t aPrimaryDevAddr, \[optional\] out unsigned long aLen, \[array, size\_is(aLen), retval\] out uint8\_t aLogicAddr)
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

The list of logical addresses that the adapter is controlling.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aPrimaryDevAddr</td>
<td align="left"><p>receives the address of the primary device.</p></td>
</tr>
<tr class="even">
<td align="left">aLen</td>
<td align="left"><p>receives the number of addresses.</p></td>
</tr>
<tr class="odd">
<td align="left">aLogicAddr</td>
<td align="left"><p>receives the array of addresses.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a boolean (success/failure).

void nsISystemAdapterCec::getLogicalAddress (in uint8\_t aDeviceType, out uint8\_t aLogicAddr)
----------------------------------------------------------------------------------------------

The logical address of the given type that the adapter is controlling.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aDeviceType</td>
<td align="left"><p>the type of device whose logical address is requested.</p></td>
</tr>
<tr class="even">
<td align="left">aLogicAddr</td>
<td align="left"><p>the logical address of the device.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a boolean (success/failure).

void nsISystemAdapterCec::sendMsg (in uint8\_t aSource, in uint8\_t aDest, in ACString aByteMessage, in bool aExpectedReply, in uint8\_t aReplyOpcodeFilter, in nsIGktEventHandler aSendMsgResultEventHandler)
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Transmit a raw message.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aSource</td>
<td align="left"><p>the logical address of the source of the message</p></td>
</tr>
<tr class="even">
<td align="left">aDest</td>
<td align="left"><p>the logical address of the destination of the message</p></td>
</tr>
<tr class="odd">
<td align="left">aByteMessage</td>
<td align="left"><p>the bytes of the message, except the addresses field</p></td>
</tr>
<tr class="even">
<td align="left">aExpectedReply</td>
<td align="left"><p>a boolean, true if a reply is expected</p></td>
</tr>
<tr class="odd">
<td align="left">aReplyOpcodeFilter</td>
<td align="left"><p>the expected opcode for the reply, if any</p></td>
</tr>
<tr class="even">
<td align="left">aSendMsgResultEventHandler</td>
<td align="left"><p>the event handler for the result of the exchange, the event inherits of nsIGktEventMsg and its type has the value &quot;sendMsgResult&quot;.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a boolean (success/failure).

void nsISystemAdapterCec::addObserver (in nsIGktEventHandler aObserver)
-----------------------------------------------------------------------

Add observer to get notified of plug changes or message reception. The event returned in method nsIGktEventHandler::handleEvent inherits of nsIGktEventPlug interface for a plug notification or nsIGktEventMsg interface when receiving a message.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>the event handler to be added.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

none.

void nsISystemAdapterCec::removeObserver (in nsIGktEventHandler aObserver)
--------------------------------------------------------------------------

Remove observer.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>the event handler to be removed.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

none.
