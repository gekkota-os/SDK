nsIMyComponent interface Reference
==================================

-   void captureObject ( in nsIDOMWindow aWindow, in AUTF8String aOuputPath, in AUTF8String aMimetype, in unsigned long aRotationAngle)

Detailed Description
--------------------

The nsIMyComponent interface can only be used on SMH300.

void nsIMyComponent::captureObject (in nsIDOMWindow aWindow, in AUTF8String aOuputPath, in AUTF8String aMimetype, in unsigned long aRotationAngle)
--------------------------------------------------------------------------------------------------------------------------------------------------

Function

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aWindow</td>
<td align="left"><p>Global DOM object Window</p></td>
</tr>
<tr class="even">
<td align="left">aOutputPath</td>
<td align="left"><p>Output path</p></td>
</tr>
<tr class="odd">
<td align="left">aMimetype</td>
<td align="left"><p>Mime type of the output picture (supported: &quot;image/g4&quot;, &quot;image/jpeg&quot;). If null, choose default Mime type according to platform (supports)</p></td>
</tr>
<tr class="even">
<td align="left">aRotationAngle</td>
<td align="left"><p>Specifies a rotation for the capture between 0-359° (not supported yet)</p></td>
</tr>
</tbody>
</table>

The function is used like below in Javascript.

    var myComponent = new MyComponent();
    var path = "/tmp/myComponent/img.ppk";
    myComponent.captureObject(window, path, "image/g4", 0);

Complete HTML example [here.](example1.html)
