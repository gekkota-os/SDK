# HID interactivity detection App

This App displays two images "yellow.jpg" and "example.svg" : 

- "example.svg" during  HID interactivity (occurence of keyboard, mouse or screen events);
- "yellow.jpg" after a delay of inactivity (standby mode); 

At startup of app, the two images are displayed.  

The script which manage the calling of HID event callbacks is *xpfDetectionHid.js*. It is based on *Idle API*.

To XPF know what to do or to play when an HID event occurs, or when XPF enters standby mode, we have to implement the following
callbacks :  


- *XpfDetectionHid.onStateactiveBegin*;
- *XpfDetectionHid.onStateactiveEnd*.

## init

This function can be called explicitely in order to define the delay (in seconds) before the XPF enters standby mode.
If it is not called, then the default delay is 60 seconds.

```
XpfDetectionHid.init(1);
```

*This code sets the delay to 1 second.*


## onStateactiveBegin

This callback occurs at the time when the XPF enters interactive mode.

```
XpfDetectionHid.onStateactiveBegin = function onStateactiveBegin() {
	document.getElementById("id_yellow").endElement();
	document.getElementById("id_example").beginElement();
}
```

*This code plays the image "example.svg" when the XPF enters interactive mode. The image "yellow.jpg" is stopped.*

## onStateactiveEnd

This callback occurs at the time when the XPF switchs from interactive mode to standby mode.

```
XpfDetectionHid.onStateactiveEnd = function onStateactiveEnd() {
	document.getElementById("id_example").endElement();
	document.getElementById("id_yellow").beginElement();
}
```

*This code stops playing the image "example.svg" when the XPF leaves interactive mode. The image "yellow.jpg" is started.*