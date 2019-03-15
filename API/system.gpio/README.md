nsISystemAPBGpioOutput interface Reference
==========================================

Public Attributes
-----------------

-   attribute boolean writeValue

-   void removeObserver ( in nsISystemGpioOuputObserver aObserver)

<!-- -->

-   void addObserver ( in nsISystemGpioOuputObserver aObserver)

Detailed Description
--------------------

The nsISystemAPBGpioOutput interface allows to write value to GPIO. For SMA300 and DMB400, the jack35 GPIO is used as example: set the configuration of your platform following the instruction in the gpio\_output\_write\_configuration.js file. Once it is done, you can find an example [here.](gpio_output_write_example.html)

Member Data Documentation
-------------------------

### attribute boolean nsISystemAPBGpioOutput::writeValue

Value to set on the GPIO.

void nsISystemAPBGpioOutput::removeObserver (in nsISystemGpioOuputObserver aObserver)
-------------------------------------------------------------------------------------

Remove observer on service provider.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>Observer removed</p></td>
</tr>
</tbody>
</table>

void nsISystemAPBGpioOutput::addObserver (in nsISystemGpioOuputObserver aObserver)
----------------------------------------------------------------------------------

Add observer on service provider.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>Observer added</p></td>
</tr>
</tbody>
</table>

nsISystemAPBGpioInput interface Reference
=========================================

Public Attributes
-----------------

-   readonly attribute boolean readValue

-   void removeObserver ( in nsISystemGpioObserver aObserver)

<!-- -->

-   void addObserver ( in nsISystemGpioObserver aObserver)

Detailed Description
--------------------

The nsISystemAPBGpioInput interface allows to read the GPIO. For SMA300 and DMB400, the jack35 GPIO is used as example: set the configuration of your platform following the instruction in the gpio\_input\_observe\_configuration.js file. Once it is done, you can find an example [here.](gpio_input_observe_example.html)

Member Data Documentation
-------------------------

### readonly attribute boolean nsISystemAPBGpioInput::readValue

Read the value of the GPIO.

void nsISystemAPBGpioInput::removeObserver (in nsISystemGpioObserver aObserver)
-------------------------------------------------------------------------------

Remove observer on service provider.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>Observer removed</p></td>
</tr>
</tbody>
</table>

void nsISystemAPBGpioInput::addObserver (in nsISystemGpioObserver aObserver)
----------------------------------------------------------------------------

Add observer on service provider.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>Observer added</p></td>
</tr>
</tbody>
</table>

nsISystemGpioOuputObserver interface Reference
==============================================

-   void onChange ( in nsISystemAPBGpioOutput aAPBGpioOutput, in boolean aOldValue, in boolean aNewValue)

void nsISystemGpioOuputObserver::onChange (in nsISystemAPBGpioOutput aAPBGpioOutput, in boolean aOldValue, in boolean aNewValue)
--------------------------------------------------------------------------------------------------------------------------------

Callback called on changes.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAPBGpioOutput</td>
<td align="left"><p>APB GPIO Output</p></td>
</tr>
<tr class="even">
<td align="left">aOldValue</td>
<td align="left"><p>Old value</p></td>
</tr>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>New value</p></td>
</tr>
</tbody>
</table>

nsISystemGpioObserver interface Reference
=========================================

-   void onChange ( in nsISystemAPBGpioInput aAPBGpioInput, in boolean aOldValue, in boolean aNewValue)

void nsISystemGpioObserver::onChange (in nsISystemAPBGpioInput aAPBGpioInput, in boolean aOldValue, in boolean aNewValue)
-------------------------------------------------------------------------------------------------------------------------

Callback: onChange.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAPBGpioInput</td>
<td align="left"><p>Interface nsISystemAPBGpioInput</p></td>
</tr>
<tr class="even">
<td align="left">aOldValue</td>
<td align="left"><p>Old value boolean</p></td>
</tr>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>New value boolean</p></td>
</tr>
</tbody>
</table>


