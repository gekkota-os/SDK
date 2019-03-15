# Human Interface Device Specific Input

HID devices generate events which describes the user interaction with the device. In order for the application to know about specific sources that generates these events, they can be filtered with the HTML5 property **location** from the key event.
The **location** property is a number representing the location of the key on the keyboard or other input device.

The following constants allow to filter different Innes devices inputs:

### const unsigned short nsIDOMKeyEvent::DOM\_KEY\_LOCATION\_RC\_CEC

Constant to apply on location property to filter HID inputs from remote controls of type Consumer Electronic Control Passthrough. Here is an example:

````javascript
window.top.document.addEventListener("keydown", function (event) {
    var isCecInput;
    // If this condition is true, the event input comes from a CEC remote control
    if (event.location & event.DOM_KEY_LOCATION_RC_CEC) {
        isCecInput = true;
    } else {
        isCecInput = false;
    }
});
````

### const unsigned short nsIDOMKeyEvent::DOM\_KEY\_LOCATION\_SLATE

Constant to apply on location property to filter HID inputs from Slate devices. Once you ensure that the input is from a Slate you can retrieve its Id from 1 to 1000 by filtering the least significant bits of the location value. Here is an example:

````javascript
const SLATE_ID_MASK = 0X0FFF; // mask to find the id of a Slate, maximum value is 1000.
window.top.document.addEventListener("keydown", function (event) {
    var slateId;
    // If this condition is true, the event input comes from a Slate device
    if (event.location & event.DOM_KEY_LOCATION_SLATE) {
        // This filtering gives you the Id of the Slate which generated the input
        slateId = event.location & SLATE_ID_MASK;
    }
});
````

