nsISystemAPBAVCmd interface Reference
=====================================

Public Attributes
-----------------

-   readonly attribute nsIPropertyBag2 properties

<!-- -->

-   attribute boolean standby

<!-- -->

-   const unsigned long POWER\_MODE\_OFF

<!-- -->

-   const unsigned long POWER\_MODE\_ON

<!-- -->

-   const unsigned long POWER\_MODE\_SUSPEND

<!-- -->

-   const unsigned long POWER\_MODE\_STANDBY

<!-- -->

-   const unsigned long POWER\_MODE\_OFF\_POWER\_BUTTON

<!-- -->

-   attribute unsigned long powerMode

<!-- -->

-   attribute bool mute

<!-- -->

-   const short VIDEO\_INPUT\_NONE

<!-- -->

-   const short VIDEO\_INPUT\_RGB1

<!-- -->

-   const short VIDEO\_INPUT\_RGB2

<!-- -->

-   const short VIDEO\_INPUT\_TMDS1

<!-- -->

-   const short VIDEO\_INPUT\_TMDS2

<!-- -->

-   const short VIDEO\_INPUT\_COMPOSITE1

<!-- -->

-   const short VIDEO\_INPUT\_COMPOSITE2

<!-- -->

-   const short VIDEO\_INPUT\_SVIDEO1

<!-- -->

-   const short VIDEO\_INPUT\_SVIDEO2

<!-- -->

-   const short VIDEO\_INPUT\_TUNER1

<!-- -->

-   const short VIDEO\_INPUT\_TUNER2

<!-- -->

-   const short VIDEO\_INPUT\_TUNER3

<!-- -->

-   const short VIDEO\_INPUT\_COMPONENT1

<!-- -->

-   const short VIDEO\_INPUT\_COMPONENT2

<!-- -->

-   const short VIDEO\_INPUT\_COMPONENT3

<!-- -->

-   const short VIDEO\_INPUT\_DISPLAY\_PORT1

<!-- -->

-   const short VIDEO\_INPUT\_DISPLAY\_PORT2

<!-- -->

-   const short VIDEO\_INPUT\_TMDS3

<!-- -->

-   const short VIDEO\_INPUT\_TMDS4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI5

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI6

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI7

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI8

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI\_CEC\_ACTIVE\_SOURCE

<!-- -->

-   const unsigned long VIDEO\_INPUT\_HDMI\_CEC\_INACTIVE\_SOURCE

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VGA1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VGA2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VGA3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VGA4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DVI1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DVI2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DVI3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DVI4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_RGB3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_RGB4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_COMPONENT4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VIDEO1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VIDEO2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VIDEO3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_VIDEO4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_SVIDEO3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_SVIDEO4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_COMPOSITE3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_COMPOSITE4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_PC1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_PC2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_PC3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_PC4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DTV1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DTV2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DTV3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DTV4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_TV

<!-- -->

-   const unsigned long VIDEO\_INPUT\_TV1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_TV2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_SCART1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_SCART2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_TUNER4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DISPLAY\_PORT3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DISPLAY\_PORT4

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DISPLAY\_PORT5

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DISPLAY\_PORT6

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DISPLAY\_PORT7

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DISPLAY\_PORT8

<!-- -->

-   const unsigned long VIDEO\_INPUT\_OPTION1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_OPTION2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_OPTION3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DVD1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DVD2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_DVD3

<!-- -->

-   const unsigned long VIDEO\_INPUT\_OTHER1

<!-- -->

-   const unsigned long VIDEO\_INPUT\_OTHER2

<!-- -->

-   const unsigned long VIDEO\_INPUT\_OTHER3

<!-- -->

-   attribute long videoInput

<!-- -->

-   attribute unsigned short brightness

<!-- -->

-   attribute unsigned short backlight

<!-- -->

-   attribute unsigned short volume

-   void setIds ( in string aIds, in unsigned long aLength)

<!-- -->

-   void getIds ( out unsigned long aLength, out string aIds)

Detailed Description
--------------------

Interface XPCOM nsISystemAPBAVCmd the urn of profile is : `urn:innes:owl:av-cmd:1`. For the examples, set the configuration of your platform following the instructions in the avcmd\_brightness\_configuration.js and avcmd\_power\_configuration.js files. Once it is done, you can find the examples here : [avcmd\_brightness\_example.html [avcmd\_power\_example.html.](avcmd_power_example.html)](avcmd_brightness_example.html)

**Build note**: You need to execute the **build.cmd** file to generate the boostrap app. Otherwise there will be a mismatch between the html file name and the one the manifest tries to launch. Find more information in *SDK-G4/bootstrap App/* documentation.

Member Data Documentation
-------------------------

### readonly attribute nsIPropertyBag2 nsISystemAPBAVCmd::properties

Properties of profile

### attribute boolean nsISystemAPBAVCmd::standby

Standby state for backward compatibility

### const unsigned long nsISystemAPBAVCmd::POWER\_MODE\_OFF

Powermode values

### attribute unsigned long nsISystemAPBAVCmd::powerMode

The power mode value which can be one of POWER\_MODE\_&lt;?&gt; listed above.

### attribute bool nsISystemAPBAVCmd::mute

Audio mute

### const short nsISystemAPBAVCmd::VIDEO\_INPUT\_NONE

The possible values of input source based on MCCS V2.2 ([https://milek7.pl/ddcbacklight/mccs.pdf)](https://milek7.pl/ddcbacklight/mccs.pdf)

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_HDMI1

HDMI inputs

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_VGA1

VGA inputs

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_DVI1

DVI inputs

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_RGB3

RGB inputs

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_COMPONENT4

COMPONENT input

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_VIDEO1

VIDEO input

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_SVIDEO3

SVIDEO input

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_COMPOSITE3

COMPOSITE input

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_PC1

PC input

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_DTV1

DTV input

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_TV

OTHER input

### const unsigned long nsISystemAPBAVCmd::VIDEO\_INPUT\_DISPLAY\_PORT3

DIPLAY PORT inputs

### attribute long nsISystemAPBAVCmd::videoInput

The selected video input which can be one of VIDEO\_INPUT\_&lt;?&gt; listed above.

### attribute unsigned short nsISystemAPBAVCmd::brightness

Brightness value between 0 to 100.

### attribute unsigned short nsISystemAPBAVCmd::backlight

Backlight value between 0 to 100.

### attribute unsigned short nsISystemAPBAVCmd::volume

Audio volume between 0 to 100.

void nsISystemAPBAVCmd::setIds (\[array, size\_is(aLength)\] in string aIds, \[optional\] in unsigned long aLength)
-------------------------------------------------------------------------------------------------------------------

Set the array of device identificators that need to apply the change of the attribute. By default, all devices must apply the change of the attribute (broadcast mode). This method can be used to indicate which particular device, identifed by its identificator, must apply the change. The device identificator can be of the following form : "0x&lt;HH&gt;" ("0xFE"), "&lt;number&gt;" ("255") or "&lt;character&gt;" ("\*"). The "device identificator "\*" is reserved for the broadcast mode. A empty array indicates a broadcast mode.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aIds</td>
<td align="left"><p>array of device identificators that need to apply the change of the attribute</p></td>
</tr>
<tr class="even">
<td align="left">aLength</td>
<td align="left"><p>the length of array.</p></td>
</tr>
</tbody>
</table>

void nsISystemAPBAVCmd::getIds (\[optional\] out unsigned long aLength, \[retval, array, size\_is(aLength)\] out string aIds)
-----------------------------------------------------------------------------------------------------------------------------

Get the array of device identificators that need to apply the change of the attribute.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aIds</td>
<td align="left"><p>array of device identificators that need to apply the change of the attribute</p></td>
</tr>
<tr class="even">
<td align="left">aLength</td>
<td align="left"><p>the length of array.</p></td>
</tr>
</tbody>
</table>


