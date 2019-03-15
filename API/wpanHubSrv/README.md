nsIWpanHubSrv interface Reference
=================================

Public Attributes
-----------------

-   const unsigned long BASEURI\_SLATE\_PICTURE

-   void getFileInfo ( in unsigned long aBaseUri, in ACString aFilePath, out boolean aIsFileSync, out ACString aLastUpdateTime)

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
