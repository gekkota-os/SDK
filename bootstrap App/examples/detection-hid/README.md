# HID detection App

This App displays two images "yellow.jpg" and "example.svg" : 

- "example.svg" when an HID event occurs (keyboard, mouse, screen);
- "yellow.jpg" after a delay of inactivity; 

At startup of app, the two images are displayed.  

The script which manage the HID event callbacks is *xpfDetectionHid.js*. It is based on *Idle API*.

To XPF know what to do or to play when an HID event occurs, or when XPF enters Inactivity state, we have to implement the following
callbacks :  

- *XpfDetectionHid.onStateidleBegin*;
- *XpfDetectionHid.onStateidleEnd*;
- *XpfDetectionHid.onStateactiveBegin*;
- *XpfDetectionHid.onStateactiveEnd*.

## init

This function can be called explicitely in order to define the delay (in seconds) before the XPF enters Inactive state.
If it is not called, then the default delay is 5 seconds.

```
XpfDetectionHid.init(1);
```

*This code sets the delay to 1 second.*

## onStateidleBegin

This callback occurs at the time when the XPF enters Inactive state.

```
XpfDetectionHid.onStateidleBegin = function onStateidleBegin() {
	document.getElementById("id_yellow").beginElement();
}
```

*This code plays the image "yellow.jpg" when the XPF enters Inactive state.*

## onStateidleEnd

This callback occurs at the time when the XPF leaves Inactive state.

```
XpfDetectionHid.onStateidleEnd = function onStateidleEnd() {
	document.getElementById("id_yellow").endElement();
}
```

*This code stops playing the image "yellow.jpg" when the XPF leaves Inactive state.*

## onStateactiveBegin

This callback occurs at the time when the XPF enters Active state.

```
XpfDetectionHid.onStateactiveBegin = function onStateactiveBegin() {
	document.getElementById("id_example").beginElement();
}
```

*This code plays the image "example.svg" when the XPF enters Active state.*

## onStateactiveEnd

This callback occurs at the time when the XPF leaves Active state.

```
XpfDetectionHid.onStateactiveEnd = function onStateactiveEnd() {
	document.getElementById("id_example").endElement();
}
```

*This code stops playing the image "example.svg" when the XPF leaves Active state.*