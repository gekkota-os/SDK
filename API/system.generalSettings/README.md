nsISystemGeneralSettings interface Reference
============================================

Public Attributes
-----------------

-   attribute ACString hostname

<!-- -->

-   attribute boolean showCursor

<!-- -->

-   readonly attribute ACString psn

<!-- -->

-   readonly attribute nsISystemMACStructure mac

<!-- -->

-   readonly attribute ACString uuid

<!-- -->

-   readonly attribute ACString platform

<!-- -->

-   readonly attribute ACString version

<!-- -->

-   readonly attribute nsISystemHDCPStructure hdcp

<!-- -->

-   attribute AString field1

<!-- -->

-   attribute AString field2

<!-- -->

-   attribute AString field3

<!-- -->

-   attribute AString field4

<!-- -->

-   attribute AString field5

Detailed Description
--------------------

The nsISystemGeneralSettings interface is the point of entry to get information about the hardware platform running Gekkota. HTML example using this API [here.](example1.html)

**Build note**: You need to execute the **build.cmd** file to generate the boostrap app. Otherwise there will be a mismatch between the html file name and the one the manifest tries to launch. Find more information in *SDK-G4/bootstrap App/* documentation.

Member Data Documentation
-------------------------

### attribute ACString nsISystemGeneralSettings::hostname

Hardware platform hostname.

### attribute boolean nsISystemGeneralSettings::showCursor

Visibility of the cursor.

### readonly attribute ACString nsISystemGeneralSettings::psn

Hardware platform PSN.

### readonly attribute nsISystemMACStructure nsISystemGeneralSettings::mac

Hardware platform MAC address.

### readonly attribute ACString nsISystemGeneralSettings::uuid

Hardware platform UUID.

### readonly attribute ACString nsISystemGeneralSettings::platform

Hardware platform name.

### readonly attribute ACString nsISystemGeneralSettings::version

Hardware platform version.

### readonly attribute nsISystemHDCPStructure nsISystemGeneralSettings::hdcp

Hardware platform HDCP information structure.

### attribute AString nsISystemGeneralSettings::field1

Hardware platform field1.

### attribute AString nsISystemGeneralSettings::field2

Hardware platform field2.

### attribute AString nsISystemGeneralSettings::field3

Hardware platform field3.

### attribute AString nsISystemGeneralSettings::field4

Hardware platform field4.

### attribute AString nsISystemGeneralSettings::field5

Hardware platform field5.

nsISystemHDCPStructure interface Reference
==========================================

Public Attributes
-----------------

-   readonly attribute boolean supported

<!-- -->

-   readonly attribute ACString version

<!-- -->

-   readonly attribute boolean keyStatus

Member Data Documentation
-------------------------

### readonly attribute boolean nsISystemHDCPStructure::supported

Processor supports the encryption.

### readonly attribute ACString nsISystemHDCPStructure::version

HDCP version: "", "1.4" or "2.2".

### readonly attribute boolean nsISystemHDCPStructure::keyStatus

HDCP keys loaded.
