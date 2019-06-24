nsIWpanHubSrv interface Reference
=================================

Public Attributes
-----------------

-   const unsigned long BASEURI\_SLATE\_PICTURE

-   void getFileInfo ( in unsigned long aBaseUri, in ACString aFilePath, out boolean aIsFileSync, out ACString aLastUpdateTime)

<!-- -->

-   void setButtonConfiguration ( in nsIPropertyBag aJsonArgs)

Detailed Description
--------------------

The nsIWpanHubSrv interface can only be used on SMH300 platform.

Member Data Documentation
-------------------------

### const unsigned long nsIWpanHubSrv::BASEURI\_SLATE\_PICTURE

Base URI of the file Base URI Slate picture: the last directory name before the file you want info from must be the id of the Slate. For example the SLate of id 1 must end with "1/file.ppk"

void nsIWpanHubSrv::getFileInfo (in unsigned long aBaseUri, in ACString aFilePath, out boolean aIsFileSync, out ACString aLastUpdateTime)
-----------------------------------------------------------------------------------------------------------------------------------------

Returns information for a file regarding a device managed by the hub.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aBaseUri</td>
<td align="left"><p>The base URI where the file is located</p></td>
</tr>
<tr class="even">
<td align="left">aFilePath</td>
<td align="left"><p>The file path from the specified base URI</p></td>
</tr>
<tr class="odd">
<td align="left">aIsSync</td>
<td align="left"><p>Is the file on the device synchronized with the one on the hub</p></td>
</tr>
<tr class="even">
<td align="left">aLastUploadTime</td>
<td align="left"><p>The ISO formated date of the last time the device took the specified file</p></td>
</tr>
</tbody>
</table>

Here is an example of how to use this function in Javascript :

    var wpanHubSrv = window.wpanHubSrv || {};
    if (wpanHubSrv.getFileInfo) {
       var slateId = 3;
       var sBaseUri = wpanHubSrv.BASEURI_SLATE_PICTURE;
       var sFilePath = slateId + SLATE_FILE_PATH;
       var isFileInSync = {};
       var lastUpdateTime = {};
       wpanHubSrv.getFileInfo(sBaseUri, sFilePath, isFileInSync, lastUpdateTime);
       window.logger.debug(
       " filePath: " + sFilePath +
       " isFileInSync: " + JSON.stringify(isFileInSync) +
       " lastUpdateTime: " + JSON.stringify(lastUpdateTime));
    }

void nsIWpanHubSrv::setButtonConfiguration (in nsIPropertyBag aJsonArgs)
------------------------------------------------------------------------

This a function to be used by the application to setup runtime button bar configuration for a specific device.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aJsonArgs</td>
<td align="left"><p>The JSON arguments used to set the configuration of the given device. It may be different for each type of device depending of the number of buttons.</p></td>
</tr>
<tr class="even">
<td align="left">location</td>
<td align="left"><p>The signature of the device, it specifies both the device type and index. The id of the Slate is specified using the LSB.</p></td>
</tr>
<tr class="odd">
<td align="left">buttonConfigurationN</td>
<td align="left"><p>The JSON contains N keys, one for each button, and each button has 1 key for now:</p></td>
</tr>
<tr class="even">
<td align="left">enable</td>
<td align="left"><p>Boolean value specifying if the button is enabled or disabled</p></td>
</tr>
</tbody>
</table>

The button keys are names buttonConfigurationN where N is the button number starting at one. For Slate106, there are 3 buttons: the first one is the left one, the second one is the middle one, and the third on is the right one. For example, an input from a Slate Key service type from Slate with ID 5 has the location: 0x80000005. See HID documentation for more info. Here is an example for a 3 buttons device:

    {
      location: 2147483653,
      boutonConfigurations: [{
        id: 1,
        enable: true
      }, {
        id: 2,
         enable: false
      }, {
        id: 3,
        enable: true
      }]
    }
