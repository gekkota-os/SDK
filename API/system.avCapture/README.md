nsISystemAdapterAVCapture interface Reference
=============================================

Public Attributes
-----------------

-   const unsigned long VIDEO\_CVBS

<!-- -->

-   const unsigned long VIDEO\_VGA

<!-- -->

-   const unsigned long VIDEO\_HDMI

<!-- -->

-   const unsigned long VIDEO\_DVI

<!-- -->

-   const unsigned long VIDEO\_SVIDEO

<!-- -->

-   const unsigned long VIDEO\_SCART

<!-- -->

-   const unsigned long AUDIO\_RCA

<!-- -->

-   const unsigned long AUDIO\_JACK

<!-- -->

-   const unsigned long AUDIO\_HDMI

<!-- -->

-   const unsigned long AUDIO\_XLR

<!-- -->

-   const unsigned long AUDIO\_OPTIC

<!-- -->

-   const unsigned long AUDIO\_SCART

<!-- -->

-   attribute unsigned long currentVideoInput

<!-- -->

-   attribute unsigned long currentAudioInput

<!-- -->

-   readonly attribute nsIInputStream inputStream

-   void init ( in AUTF8String aSysPath)

<!-- -->

-   void getAvailableAudioInputs ( out unsigned long aLength, out unsigned long aAudioInputTypes)

<!-- -->

-   void getAvailableVideoInputs ( out unsigned long aLength, out unsigned long aVideoInputTypes)

Member Data Documentation
-------------------------

### const unsigned long nsISystemAdapterAVCapture::VIDEO\_CVBS

Constants for video input types. Video input type is CVBS.

### const unsigned long nsISystemAdapterAVCapture::VIDEO\_VGA

Video input type is VGA.

### const unsigned long nsISystemAdapterAVCapture::VIDEO\_HDMI

Video input type is HDMI.

### const unsigned long nsISystemAdapterAVCapture::VIDEO\_DVI

Video input type is DVI.

### const unsigned long nsISystemAdapterAVCapture::VIDEO\_SVIDEO

Video input type is SVIDEO.

### const unsigned long nsISystemAdapterAVCapture::VIDEO\_SCART

Video input type is SCART.

### const unsigned long nsISystemAdapterAVCapture::AUDIO\_RCA

Constants for audio input types. Audio input type is RCA.

### const unsigned long nsISystemAdapterAVCapture::AUDIO\_JACK

Audio input type is JACK.

### const unsigned long nsISystemAdapterAVCapture::AUDIO\_HDMI

Audio input type is HDMI.

### const unsigned long nsISystemAdapterAVCapture::AUDIO\_XLR

Audio input type is XLR.

### const unsigned long nsISystemAdapterAVCapture::AUDIO\_OPTIC

Audio input type is OPTIC.

### const unsigned long nsISystemAdapterAVCapture::AUDIO\_SCART

Audio input type is SCART.

### attribute unsigned long nsISystemAdapterAVCapture::currentVideoInput

Select video input with its index.

### attribute unsigned long nsISystemAdapterAVCapture::currentAudioInput

Select audio input with its index.

### readonly attribute nsIInputStream nsISystemAdapterAVCapture::inputStream

Video Stream.

void nsISystemAdapterAVCapture::init (in AUTF8String aSysPath)
--------------------------------------------------------------

Initialization of the implementation class.

void nsISystemAdapterAVCapture::getAvailableAudioInputs (out unsigned long aLength, \[array, size\_is(aLength), retval\] out unsigned long aAudioInputTypes)
------------------------------------------------------------------------------------------------------------------------------------------------------------

Returns the number of audio inputs available with their type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>Number of audio entries on the capture card</p></td>
</tr>
<tr class="even">
<td align="left">aAudioInputTypes</td>
<td align="left"><p>Lists audio entries identifies by their type</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterAVCapture::getAvailableVideoInputs (out unsigned long aLength, \[array, size\_is(aLength), retval\] out unsigned long aVideoInputTypes)
------------------------------------------------------------------------------------------------------------------------------------------------------------

Returns the number of video inputs available with their type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>Number of video entries on the capture card</p></td>
</tr>
<tr class="even">
<td align="left">aVideoInputTypes</td>
<td align="left"><p>Lists video entries identifies by their type</p></td>
</tr>
</tbody>
</table>


