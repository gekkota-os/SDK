nsISystemAdapterAppColorLed interface Reference
===============================================

Public Attributes
-----------------

-   const unsigned short LEFT\_COLOR\_LED

<!-- -->

-   const unsigned short RIGHT\_COLOR\_LED

<!-- -->

-   attribute boolean enabled

<!-- -->

-   attribute unsigned long color

<!-- -->

-   readonly attribute unsigned short placement

Detailed Description
--------------------

The nsISystemAdapterAppColorLed interface allows to manage color led. HTML example using this API [here.](example1.html)

**Build note**: You need to execute the **build.cmd** file to generate the boostrap app. Otherwise there will be a mismatch between the html file name and the one the manifest tries to launch. Find more information in *SDK-G4/bootstrap App/* documentation.

Member Data Documentation
-------------------------

### const unsigned short nsISystemAdapterAppColorLed::LEFT\_COLOR\_LED

Placement of color leds: left.

### const unsigned short nsISystemAdapterAppColorLed::RIGHT\_COLOR\_LED

Placement of color leds: right.

### attribute boolean nsISystemAdapterAppColorLed::enabled

Enable or disable the color led.

### attribute unsigned long nsISystemAdapterAppColorLed::color

Color RGB (XXRRGGBB) of the color led.

### readonly attribute unsigned short nsISystemAdapterAppColorLed::placement

Placement of the color led (LEFT\_COLOR\_LED or RIGHT\_COLOR\_LED).
