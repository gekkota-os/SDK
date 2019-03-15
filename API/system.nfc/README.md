nsISystemAdapterNfc interface Reference
=======================================

Public Attributes
-----------------

-   const unsigned short MODULATION\_ISO14443A

<!-- -->

-   const unsigned short MODULATION\_JEWEL

<!-- -->

-   const unsigned short MODULATION\_ISO14443B

<!-- -->

-   const unsigned short MODULATION\_ISO14443BI

<!-- -->

-   const unsigned short MODULATION\_ISO14443B2SR

<!-- -->

-   const unsigned short MODULATION\_ISO14443B2CT

<!-- -->

-   const unsigned short MODULATION\_FELICA

<!-- -->

-   const unsigned short MODULATION\_DEP

<!-- -->

-   const unsigned short BAUD\_106

<!-- -->

-   const unsigned short BAUD\_212

<!-- -->

-   const unsigned short BAUD\_424

<!-- -->

-   const unsigned short BAUD\_847

<!-- -->

-   readonly attribute nsIArray supportedModulations

-   void configurePolling ( in unsigned long aModulationLength, in unsigned short aModulation, in unsigned long aBaudRateLength, in unsigned short aBaudRate)

<!-- -->

-   void configureMultipleModulations ( in nsIPropertyBag aModulations, in unsigned long aModulationLength)

<!-- -->

-   boolean isSupported ( in unsigned short aModulation, in unsigned short aBaudRate)

<!-- -->

-   nsIEnumerator scanTargets ( )

<!-- -->

-   unsigned long startPoll ( in nsISystemAdapterNfcObserver aObserver)

<!-- -->

-   void stopPoll ( in unsigned long aObserverId)

Detailed Description
--------------------

The nsISystemAdapterNfc interface is the point of entry for using NFC. HTML example using this API on SMT platform [here.](example1.html)

Member Data Documentation
-------------------------

### const unsigned short nsISystemAdapterNfc::MODULATION\_ISO14443A

Type ISO 14443 A.

### const unsigned short nsISystemAdapterNfc::MODULATION\_JEWEL

Type Jewel.

### const unsigned short nsISystemAdapterNfc::MODULATION\_ISO14443B

Type ISO 14443 B.

### const unsigned short nsISystemAdapterNfc::MODULATION\_ISO14443BI

Type pre ISO 14443 B or ISO 14443 B'.

### const unsigned short nsISystemAdapterNfc::MODULATION\_ISO14443B2SR

Type ISO 14443 2B ST SRx.

### const unsigned short nsISystemAdapterNfc::MODULATION\_ISO14443B2CT

Type ISO 14443 2B ASK CTx.

### const unsigned short nsISystemAdapterNfc::MODULATION\_FELICA

Type Felica.

### const unsigned short nsISystemAdapterNfc::MODULATION\_DEP

Protocol ISO DEP.

### const unsigned short nsISystemAdapterNfc::BAUD\_106

Baud rate 106 kbit/s.

### const unsigned short nsISystemAdapterNfc::BAUD\_212

Baud rate 212 kbit/s.

### const unsigned short nsISystemAdapterNfc::BAUD\_424

Baud rate 424 kbit/s.

### const unsigned short nsISystemAdapterNfc::BAUD\_847

Baud rate 847 kbit/s.

### readonly attribute nsIArray nsISystemAdapterNfc::supportedModulations

The array could be:

     [
     { modulationType:1, baudRates:[ 1,2 ]},
     { modulationType:2, baudRates: [ 3 ]}
    ] 

void nsISystemAdapterNfc::configurePolling (in unsigned long aModulationLength, \[array, size\_is(aModulationLength)\] in unsigned short aModulation, in unsigned long aBaudRateLength, \[array, size\_is(aBaudRateLength)\] in unsigned short aBaudRate)
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Configure type of polling. Must be called before startPoll. If configurePolling() is not called, polling use all supported modulations (slower).

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aModulationLength</td>
<td align="left"><p>Size of array of modulation types</p></td>
</tr>
<tr class="even">
<td align="left">aModulation</td>
<td align="left"><p>Array of modulation type (see constants MODULATION_*)</p></td>
</tr>
<tr class="odd">
<td align="left">aBaudRateLength</td>
<td align="left"><p>Size of array of baud rates</p></td>
</tr>
<tr class="even">
<td align="left">aBaudRate</td>
<td align="left"><p>Array of baud rate for each type of modulation (see constants BAUD_*)</p></td>
</tr>
</tbody>
</table>

**Returns:.**

Void.

Deprecated

Use configureMultipleModulations instead.

void nsISystemAdapterNfc::configureMultipleModulations (\[array, size\_is(aModulationLength)\] in nsIPropertyBag aModulations, in unsigned long aModulationLength)
------------------------------------------------------------------------------------------------------------------------------------------------------------------

Specify type(s) of polling. For each modulation desired, specify one or more baud rates. Must be called before startPoll. If configureMultipleModulations() is not called or if empty array, polling use all supported modulations (slower).

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aModulations</td>
<td align="left"><p>Array of property bags</p>
<pre><code>[
 { modulationType: 1, baudRates: [ 1,2 ]},
 { modulationType: 2, baudRates: [ 3 ]}
]</code></pre></td>
</tr>
<tr class="even">
<td align="left">aModulationLength</td>
<td align="left"><p>Size of array of property bags</p></td>
</tr>
</tbody>
</table>

**Returns:.**

Void.

boolean nsISystemAdapterNfc::isSupported (in unsigned short aModulation, in unsigned short aBaudRate)
-----------------------------------------------------------------------------------------------------

Tell if the current device supports the modulation and its baud rate.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aModulation</td>
<td align="left"><p>Type of modulation</p></td>
</tr>
<tr class="even">
<td align="left">aBaudRate</td>
<td align="left"><p>Baud rate</p></td>
</tr>
</tbody>
</table>

**Returns:.**

True is the modulation is supported.

nsIEnumerator nsISystemAdapterNfc::scanTargets ()
-------------------------------------------------

Detect synchronously the targets in the near field.

**Returns:.**

The list of nsISystemAdapterNfcTarget targets.

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

unsigned long nsISystemAdapterNfc::startPoll (in nsISystemAdapterNfcObserver aObserver)
---------------------------------------------------------------------------------------

Start polling to detect some targets. Polling is based on configured modulation if there (polling faster), or all supported modulations (polling is a bit slower).

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>Observer nsISystemAdapterNfcObserver</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The ID for this observer.

void nsISystemAdapterNfc::stopPoll (in unsigned long aObserverId)
-----------------------------------------------------------------

Stop polling.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserverId</td>
<td align="left"><p>The ID of the observer (resulting from startPoll)</p></td>
</tr>
</tbody>
</table>

nsISystemAdapterNfcObserver interface Reference
===============================================

-   void onTargetFound ( in nsISystemAdapterNfcTarget aTarget)

<!-- -->

-   void onTargetLost ( in nsISystemAdapterNfcTarget aTarget)

<!-- -->

-   void onMessageRead ( in nsISystemAdapterNfcTarget aTarget, in nsIArray aRecords)

<!-- -->

-   void onPollStop ( )

<!-- -->

-   void onPollStart ( )

Detailed Description
--------------------

The nsISystemAdapterNfcObserver interface provides some callbacks for polling and detection of targets.

void nsISystemAdapterNfcObserver::onTargetFound (in nsISystemAdapterNfcTarget aTarget)
--------------------------------------------------------------------------------------

Callback called when a target enters the near field of the device.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aTarget</td>
<td align="left"><p>Target which enter</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterNfcObserver::onTargetLost (in nsISystemAdapterNfcTarget aTarget)
-------------------------------------------------------------------------------------

Callback called when a target leaves the near field of the device.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aTarget</td>
<td align="left"><p>Target which leave</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterNfcObserver::onMessageRead (in nsISystemAdapterNfcTarget aTarget, in nsIArray aRecords)
------------------------------------------------------------------------------------------------------------

Callback called when a Peer send a message to the device.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aTarget</td>
<td align="left"><p>Peer</p></td>
</tr>
<tr class="even">
<td align="left">aRecords</td>
<td align="left"><p>Array of nsISystemAdapterNDEFRecord read</p></td>
</tr>
</tbody>
</table>

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

void nsISystemAdapterNfcObserver::onPollStop ()
-----------------------------------------------

Callback called when polling is disabled.

void nsISystemAdapterNfcObserver::onPollStart ()
------------------------------------------------

Callback called when polling is enabled.

nsISystemAdapterNfcTarget interface Reference
=============================================

Public Attributes
-----------------

-   const unsigned long TARGET\_TAG

<!-- -->

-   const unsigned long TARGET\_PEER

<!-- -->

-   readonly attribute AString targetId

<!-- -->

-   readonly attribute unsigned short targetModulation

<!-- -->

-   readonly attribute unsigned short targetBaudRate

<!-- -->

-   readonly attribute unsigned long targetType

<!-- -->

-   readonly attribute boolean ndefCompatible

<!-- -->

-   readonly attribute boolean isPresent

-   void writeMessage ( in unsigned long aLength, in nsISystemAdapterNDEFRecord aRecords)

Detailed Description
--------------------

The nsISystemAdapterNfcTarget interface describes the NFC targets.

Member Data Documentation
-------------------------

### const unsigned long nsISystemAdapterNfcTarget::TARGET\_TAG

Type tag.

### const unsigned long nsISystemAdapterNfcTarget::TARGET\_PEER

Type peer.

### readonly attribute AString nsISystemAdapterNfcTarget::targetId

NFC ID Identifier.

### readonly attribute unsigned short nsISystemAdapterNfcTarget::targetModulation

See MODULATION\_XXX constants.

### readonly attribute unsigned short nsISystemAdapterNfcTarget::targetBaudRate

See BAUD\_XXX constants.

### readonly attribute unsigned long nsISystemAdapterNfcTarget::targetType

Type of the target tag or peer.

### readonly attribute boolean nsISystemAdapterNfcTarget::ndefCompatible

Indicate if the target can process NDEF Message.

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

### readonly attribute boolean nsISystemAdapterNfcTarget::isPresent

Indicate if the target is still into the near field.

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

void nsISystemAdapterNfcTarget::writeMessage (in unsigned long aLength, \[array, size\_is(aLength)\] in nsISystemAdapterNDEFRecord aRecords)
--------------------------------------------------------------------------------------------------------------------------------------------

Write NDEF message to TAG or Peer.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>Size of array of NDEF records</p></td>
</tr>
<tr class="even">
<td align="left">aRecords</td>
<td align="left"><p>Array of nsISystemAdapterNDEFRecord, NDEF records to write</p></td>
</tr>
</tbody>
</table>

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>


