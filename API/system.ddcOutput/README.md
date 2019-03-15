nsISystemAdapterDdcOutput interface Reference
=============================================

Public Attributes
-----------------

-   readonly attribute nsISystemMonitorEdid monitorEdid

<!-- -->

-   readonly attribute nsISystemMonitorDdcci monitorDdcci

<!-- -->

-   readonly attribute bool isMonitorConnected

-   void getVcpCode ( in octet aVcpCode, in nsIGktEventDdcHandler aEventHandler)

<!-- -->

-   void setVcpCode ( in octet aVcpCode, in unsigned short aVcpValue, in nsIGktEventDdcHandler aEventHandler)

<!-- -->

-   void addObserver ( in nsIGktEventDdcHandler aObserver)

<!-- -->

-   void removeObserver ( in nsIGktEventDdcHandler aObserver)

Detailed Description
--------------------

The XPCOM interface nsISystemAdapterDdcOutput is the entry point of the ddc-output adapter. It contains the methods to send and receive VCP codes to a monitor, be informed about the connection state, get the EDID tables for further parsing and the MCCS capabilities.

Member Data Documentation
-------------------------

### readonly attribute nsISystemMonitorEdid nsISystemAdapterDdcOutput::monitorEdid

An object containing EDID information (see nsISystemMonitorEdid).

### readonly attribute nsISystemMonitorDdcci nsISystemAdapterDdcOutput::monitorDdcci

An object containing MCCS capabilities information (see nsISystemMonitorDdcci).

### readonly attribute bool nsISystemAdapterDdcOutput::isMonitorConnected

A boolean which says if a monitor is currently connected.

void nsISystemAdapterDdcOutput::getVcpCode (in octet aVcpCode, in nsIGktEventDdcHandler aEventHandler)
------------------------------------------------------------------------------------------------------

Gets the value of a VCP code from the monitor. This method sends a request to the monitor. A listener allows to catch the reply, composed of the current value and the maximum value.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aVcpCode</td>
<td align="left"><p>The VCP code to get</p></td>
</tr>
<tr class="even">
<td align="left">aEventHandler</td>
<td align="left"><p>The event handler for the reply to the request The event returned in method nsIGktEventDdcHandler::handleEvent inherit of nsIGktEventGetVcpCode interface.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

None.

void nsISystemAdapterDdcOutput::setVcpCode (in octet aVcpCode, in unsigned short aVcpValue, in nsIGktEventDdcHandler aEventHandler)
-----------------------------------------------------------------------------------------------------------------------------------

Sets the value of a VCP code to the monitor. This method sends a request to the monitor. A listener allows to catch the success of the setting.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aVcpCode</td>
<td align="left"><p>The VCP code to set</p></td>
</tr>
<tr class="even">
<td align="left">aVcpValue</td>
<td align="left"><p>The value of the VCP code to set</p></td>
</tr>
<tr class="odd">
<td align="left">aEventHandler</td>
<td align="left"><p>The event handler for the reply to the request. The event returned in method nsIGktEventDdcHandler::handleEvent inherit of nsIGktEventSetVcpCode interface.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

None.

void nsISystemAdapterDdcOutput::addObserver (in nsIGktEventDdcHandler aObserver)
--------------------------------------------------------------------------------

void nsISystemAdapterDdcOutput::removeObserver (in nsIGktEventDdcHandler aObserver)
-----------------------------------------------------------------------------------

nsISystemMonitorDdcci interface Reference
=========================================

Public Attributes
-----------------

-   const short POWER\_MODE\_NONE

<!-- -->

-   const short POWER\_MODE\_ON

<!-- -->

-   const short POWER\_MODE\_STANDBY

<!-- -->

-   const short POWER\_MODE\_SUSPEND

<!-- -->

-   const short POWER\_MODE\_OFF

<!-- -->

-   const short POWER\_MODE\_OFF\_POWER\_BUTTON

<!-- -->

-   const short INPUT\_NONE

<!-- -->

-   const short INPUT\_RGB1

<!-- -->

-   const short INPUT\_RGB2

<!-- -->

-   const short INPUT\_TMDS1

<!-- -->

-   const short INPUT\_TMDS2

<!-- -->

-   const short INPUT\_COMPOSITE1

<!-- -->

-   const short INPUT\_COMPOSITE2

<!-- -->

-   const short INPUT\_SVIDEO1

<!-- -->

-   const short INPUT\_SVIDEO2

<!-- -->

-   const short INPUT\_TUNER1

<!-- -->

-   const short INPUT\_TUNER2

<!-- -->

-   const short INPUT\_TUNER3

<!-- -->

-   const short INPUT\_COMPONENT1

<!-- -->

-   const short INPUT\_COMPONENT2

<!-- -->

-   const short INPUT\_COMPONENT3

<!-- -->

-   const short INPUT\_DISPLAY\_PORT1

<!-- -->

-   const short INPUT\_DISPLAY\_PORT2

<!-- -->

-   const short INPUT\_TMDS3

<!-- -->

-   const short INPUT\_TMDS4

<!-- -->

-   readonly attribute ACString rawMccsCapabilities

<!-- -->

-   readonly attribute ACString protocolClass

<!-- -->

-   readonly attribute ACString displayType

<!-- -->

-   readonly attribute ACString displayModel

<!-- -->

-   readonly attribute ACString mccsVersion

<!-- -->

-   readonly attribute ACString luminanceValues

<!-- -->

-   readonly attribute ACString backlightValues

<!-- -->

-   readonly attribute ACString powerModeValues

<!-- -->

-   readonly attribute ACString inputSourceValues

<!-- -->

-   readonly attribute ACString audioSpeakerVolumeValues

<!-- -->

-   readonly attribute ACString audioMuteValues

Detailed Description
--------------------

The XPCOM interface nsISystemMonitorDdcci allows to know the MCCS capabilities of a connected monitor. It is a helper to build suitable VCP codes for this monitor.

Member Data Documentation
-------------------------

### const short nsISystemMonitorDdcci::POWER\_MODE\_NONE

The possible values of power mode based on MCCS V2.2.

### const short nsISystemMonitorDdcci::POWER\_MODE\_ON

Modes based on VESA DPMS definition. See "Public attributes" To see other possible values.

### const short nsISystemMonitorDdcci::POWER\_MODE\_OFF\_POWER\_BUTTON

Power off the display - functionally equivalent to turning off power using the "power button"

### const short nsISystemMonitorDdcci::INPUT\_NONE

The possible values of input source based on MCCS V2.2 Example value of input source based on MCCS V2.2. See "Public attributes" To see all possible input values.

### readonly attribute ACString nsISystemMonitorDdcci::rawMccsCapabilities

A string holding the raw MCCS capabilties.

### readonly attribute ACString nsISystemMonitorDdcci::protocolClass

A string holding the protocol class (ex : monitor).

### readonly attribute ACString nsISystemMonitorDdcci::displayType

A string holding the display type (ex : LCD).

### readonly attribute ACString nsISystemMonitorDdcci::displayModel

A string holding the display model.

### readonly attribute ACString nsISystemMonitorDdcci::mccsVersion

A string holding the supported MCCS version of the monitor.

### readonly attribute ACString nsISystemMonitorDdcci::luminanceValues

A string that tells if the luminance VCP code is supported ("Supported") or not ("Not supported"). The VCP code accepts continuous values from 0 to 100.

### readonly attribute ACString nsISystemMonitorDdcci::backlightValues

A string that tells if the backlight VCP code is supported ("Supported") or not ("Not supported"). The VCP code accepts continuous values from 0 to 100.

### readonly attribute ACString nsISystemMonitorDdcci::powerModeValues

A string that tells the supported values of the power mode VCP code. It is a list of comma-separated values.

### readonly attribute ACString nsISystemMonitorDdcci::inputSourceValues

A string that tells the supported values of the input source VCP code. It is a list of comma-separated values.

### readonly attribute ACString nsISystemMonitorDdcci::audioSpeakerVolumeValues

A string that tells if the audio speaker volume VCP code is supported ("Supported") or not ("Not supported"). The VCP code accepts continuous values from 0 to 100.

### readonly attribute ACString nsISystemMonitorDdcci::audioMuteValues

A string that tells if the audio mute VCP code is supported ("Supported") or not ("Not supported"). The VCP code accepts continuous values from 0 to 100.

nsISystemMonitorEdid interface Reference
========================================

Public Attributes
-----------------

-   readonly attribute ACString rawEedid00

<!-- -->

-   readonly attribute ACString rawEedid02

<!-- -->

-   readonly attribute ACString rawEedid10

<!-- -->

-   readonly attribute ACString rawEedid40

<!-- -->

-   readonly attribute ACString rawEedid50

<!-- -->

-   readonly attribute ACString rawEedid60

Detailed Description
--------------------

The XPCOM interface nsISystemMonitorEdid is a set of read-only EDID tables.

Member Data Documentation
-------------------------

### readonly attribute ACString nsISystemMonitorEdid::rawEedid00

A string holding the raw base EDID.

### readonly attribute ACString nsISystemMonitorEdid::rawEedid02

A string holding the raw EDID extension 02h.

### readonly attribute ACString nsISystemMonitorEdid::rawEedid10

A string holding the raw EDID extension 10h.

### readonly attribute ACString nsISystemMonitorEdid::rawEedid40

A string holding the raw EDID extension 40h.

### readonly attribute ACString nsISystemMonitorEdid::rawEedid50

A string holding the raw EDID extension 50h.

### readonly attribute ACString nsISystemMonitorEdid::rawEedid60

A string holding the raw EDID extension 60h.

nsIGktEventMonitorPlug interface Reference
==========================================

Public Attributes
-----------------

-   readonly attribute boolean isConnected

Detailed Description
--------------------

Interface XPCOM nsIGktMonitorPlugEvent. This event is returned in method nsIGktEventDdcHandler::handleEvent to return result of method nsISystemAdapterDdcOutput::setVcpCode.

Member Data Documentation
-------------------------

### readonly attribute boolean nsIGktEventMonitorPlug::isConnected

The current state of connection.

nsIGktEventSetVcpCode interface Reference
=========================================

Detailed Description
--------------------

Interface XPCOM nsIGktEventSetVcpCode. This event is returned in method nsIGktEventDdcHandler::handleEvent to return result of method nsISystemAdapterDdcOutput::setVcpCode.

nsIGktEventGetVcpCode interface Reference
=========================================

Public Attributes
-----------------

-   readonly attribute unsigned short value

<!-- -->

-   readonly attribute unsigned short maximumValue

Detailed Description
--------------------

Interface XPCOM nsIGktEventGetVcpCode. This event is returned in method nsIGktEventDdcHandler::handleEvent to return result of method nsISystemAdapterDdcOutput::getVcpCode.

Member Data Documentation
-------------------------

### readonly attribute unsigned short nsIGktEventGetVcpCode::value

The current value.

### readonly attribute unsigned short nsIGktEventGetVcpCode::maximumValue

The maximum value.

nsIGktEventDdcHandler interface Reference
=========================================

-   void handleEvent ( in nsIGktEventDdc aEvent)

Detailed Description
--------------------

Interface XPCOM nsIGktEventDdcHandler

void nsIGktEventDdcHandler::handleEvent (in nsIGktEventDdc aEvent)
------------------------------------------------------------------

nsIGktEventDdc interface Reference
==================================

Public Attributes
-----------------

-   readonly attribute AUTF8String type

<!-- -->

-   readonly attribute nsISupports target

<!-- -->

-   readonly attribute nsresult status

-   void setStatus ( in nsresult aResult)

Detailed Description
--------------------

Interface XPCOM nsIGktEventDdc.

Member Data Documentation
-------------------------

### readonly attribute AUTF8String nsIGktEventDdc::type

Type of event.

### readonly attribute nsISupports nsIGktEventDdc::target

The target of event.

### readonly attribute nsresult nsIGktEventDdc::status

Status of event.

void nsIGktEventDdc::setStatus (in nsresult aResult)
----------------------------------------------------
