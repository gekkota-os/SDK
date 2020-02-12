# GEKKOTA Device Status

## Introduction to Status file

This document describes the content of the Status file.
The format of the Status file is XML.

This file is sent from Gekkota to the status server.
The file is put on the following URI :
*\[STATUS\_URI\]/status.\[ID\].xml*, where *\[STATUS\_URI\]* is the server URI (including the path) and *\[ID\]* is the device identifier. This identifier can have following values:

* the device mac address (lowercase string with hyphens),
* the hostname,
* the upnp-like uuid (eg. 00000000-0000-0000-0001-00e04b3b3e9a).

The type of identifier is defined in the following preference:
*innes.app.device-id-type* where the value is an enum :

* 0 : mac address,
* 1 : hostname,
* 2 : uuid.

This file allows to (at a specific date) :

* Describe the device configuration (device),
* Describe the device status (status).

````mermaid
graph TD
    B[status file]
    B-->C(device)
    B-->D(status);
````

## Example

````xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<device-status xmlns="ns.innes.device-status">
    <device>
        <id-type>MAC</id-type>
        <mac>00-1c-e6-02-25-7f</mac>
        <hostname>00912-01394</hostname>
        <uuid>05b001ea-0000-0000-0000-001ce602257f</uuid>
        <modelName>dmb400</modelName>
        <modelNumber>4.12.10</modelNumber>
        <serialNumber>00910-00490</serialNumber>
        <middleware>gekkota-4</middleware>
        <field1></field1>
        <field2></field2>
        <field3></field3>
        <field4></field4>
        <field5></field5>
        <ip-addresses>
            <ip-address>
                <if-type>LAN</if-type>
                <origin>dhcp</origin>
                <value>192.168.1.124/17</value>
            </ip-address>
            <ip-address>
                <if-type>LAN</if-type>
                <origin>auto</origin>
                <value>fc00::21c:e6ff:fe02:257f/64</value>
            </ip-address>
        </ip-addresses>
        <addons/>
    </device>
    <status>
        <date>2019-07-04T14:03:23.397783+02:00</date>
        <launcher>
            <power-manager level="MAX"/>
            <start>2019-06-21T09:10:10Z</start>
            <manifest-metadata xmlns:cms="ns.innes.custom">
                <cms:app-name>custom</cms:app-name>
                <cms:app-version>1.10.10</cms:app-version>
            </manifest-metadata>
            <state>PLAY</state>
        </launcher>
        <storage>
            <total unit="byte">15695020032</total>
            <used unit="byte">44486656</used>
        </storage>
        <display-outputs>
            <display-output>
                <connector-id>hdmi_2</connector-id>
                <connector-label>HDMI1 OUT</connector-label>
                <connected>true</connected>
                <screens>
                    <screen>
                        <manufacturer-id>SAM</manufacturer-id>
                        <product-name>SAMSUNG</product-name>
                        <product-code>A7D</product-code>
                        <serial-number>00000001</serial-number>
                        <power-mode>ON</power-mode>
                        <input-source/>
                        <is-active-source/>
                        <cec-version>1.4</cec-version>
                        <menu-language/>
                        <manufacturer-OUI>0000F0</manufacturer-OUI>
                    </screen>
                </screens>
            </display-output>
        </display-outputs>
        <setup/>
    </status>
</device-status>
````

## Element <*device*>

This element contains the device configuration information :

* <*id-type*>;
* <*mac*>;
* <*hostname*>;
* <*uuid*>;
* <*modelName*>;
* <*modelNumber*>;
* <*serialNumber*>;
* <*field[1-5]*>;
* <*ip-adresses*>;
* <*addons*>.

### Element <*id-type*>

This element is required.
It defines the type of identifier which can have the following values:

* "MAC" : mac address of the first network interface;
* "Hostname" : device name;
* "uuid" : Uuid identifier of the device ;

### Element <*mac*>

This element is required.
It defines the mac address of the device.
We can retrieve it through *nsISystemGeneralSettings::mac* attribute.

### Element <*hostname*>

This element is required.
It defines the device name.
We can retrieve it through *nsISystemGeneralSettings::hostname* attribute.

### Element <*uuid*>

This element is required.
It defines the device identifier (UPNP).
We can retrieve it through *nsISystemGeneralSettings::uuid* attribute.

### Element <*modelName*>

This element is required.
It defines the model name of the platform.
We can retrieve it through *nsISystemGeneralSettings::platform* attribute.

### Element <*modelNumber*>

This element is required.
It defines the model number of the platform.
We can retrieve it through *nsISystemGeneralSettings::version* attribute.

### Element <*serialNumber*>

This element is required.
It defines the serial number of the platform.
We can retrieve it through *nsISystemGeneralSettings::psn* attribute.

### Element <*field[1..5]*>

We can retrieve the values of the fields through *nsISystemGeneralSettings::field[1..5]* attributes.
These attributes are defined in the user preferences *innes.player.device-info.field[1..5]*.
The fields can be empty.

### Element <*ip-addresses*>

This element is required.
It contains the list of the IP addresses of the device.

### Element <*ip-address*>

This element is required by the <*ip-addresses*> element.
It defines an IP address of the device.
It must contain the following elements:

* <*if-type*>;
* <*origin*>;
* <*value*>.

#### Element <*if-type*>

This element is required by <*ip-address*>.
It defines the type of the network interface :

* "LAN" for LAN network interface,
* "WLAN" for WLAN network interface.

#### Element <*origin*>

This element is required by <*ip-address*>.
It defines the origin of the IP address :

* "auto" for an automatic IPv6 address,
* "dhcp" for an IP address assigned by DHCP,
* "static" for a static IP address.

#### Element <*value*>

This element is required by <*ip-address*>.
It defines the value of the IP address.

### Element <*addons*>

This element is required.
It contains the list of installed extensions on the device.
This list excludes all the extensions such as the "configuration" and "installer" extensions that are described by the <*configuration*> element and the <*installer*> element present in the <*setup*> element of the <*status*> element.
It may contain some <*addon*> elements.

### Element <*addon*>

It describes the installation of an extension on the platform.
It must contain the following elements :

* <*id*> ;
* <*name*> ;
* <*version*> ;

#### Element <*id*>

This element is required by the <*addon*> element.
It defines the identifier of the installed extension.

#### Element <*name*>

This element is required by the <*addon*> element.
It defines the name of the installed extension.


## Element <*status*>

This element is required.
It contains the following elements :

* <*date*>,
* <*storage*>,
* <*launcher*>,
* <*setup*>.

It may contain the <*downloader*> element.

### Element <*date*>

This element is required.
It defines the current UTC date expressed in ISO-8601 format followed by the time difference of the local area (eg.  "2015-11-17T09:32:27.402+01:00").

### Element <*storage*>

This element is required.
It describes the data storage space of the platform.
It must contain the following elements :

* <*total-size*>,
* <*used-size*>.

#### Element <*total-size*>

This element is required.

It defines the total size of the data storage space of the platform.
It must contain the *@unit* attribute set to *"byte"*.

#### Element <*used-size*>

This element is required.
It defines the size used on the data storage space of the platform.
It must contain the *@unit* attribute set to *"byte"*.

### Element <*setup*>

This element is required.

It describes the list of configuration and installation extensions installed on the platform.
It may contain the following elements :

* <*configuration*>;
* <*installer*>.

#### Element <*configuration*>

This element is required.

It describes a configuration extension installed on the platform.
It must contain the following elements :

* <*version*>;
* <*metadatas*>.

#### Element <*installer*>

This element is required.
it describes a software installation extension installed on the platform.
It must contain the following elements :

* <*version*>;
* <*metadatas*>;

#### Element <*version*>

This element is required by the <*addon*>, <*configuration*> and <*installer*> elements.
It defines the installed extension version.
If the extension is an installation or configuration extension, the version is the date of installation on the platform in the current time zone in ISO-8601 format.

#### Element <*metadatas*>

This element is required by the <*configuration*> and <*installer*> elements.
It defines the list of the configuration or installation extension.

### Element <*launcher*>

Contain information about which was launched. This includes the metadata of the current Manifest file, the play state (*PLAY*, *NO\_CONTENT*, *SYSTEM*, *MIRE*, *FAILSOFT\_CLEANUP*, *FAILSOFT*) and some messages.

For example, if a system scene is displayed because a TNT scan is currently running, then the state is set to *SYSTEM* and a message indicates  that a TNT scan is currently running (subject=*SYSTEM*, severity=*INFO*, description *"scanning TNT channels"*). The descriptions are free and in english.

The <*start*> element defines the moment of the first activation of the *bootstrap* file after a publication. The value of this element is a ISO-6801 locale date.

### Element <*downloader*>

The <*downloader*> element contains informations about which is currently downloading. This element may be empty if there are no current downloads.
It contains the metadata of the Manifest file that is currently being downloaded, the date of begin of download (UTC ISO-8601), progress (double value from 0 to 1), the speed (bytes per seconds), the download process state and the different messages (see above, free descriptions in english)

#### State *MANIFEST*

This state deals with the Manifest file downloading. No metadata are known here.

#### State *INTERNAL\_RESOURCES*

This state deals with the downloading of the resources of the manifest cache that have no *@resource* attribute.

#### State *EXTERNAL\_RESOURCES*

This state deals with the downloading of the resources of the the manifest cache that have the *@resource* attribute.

#### State *TERMINATED*

This state deals with the finalization of which was downloaded (eg. purge) before playing new content.

### Element <*display-outputs*> 

This element contains information about the display outputs of the device.

#### Element <*display-output*>

This element contains information about a display output of the device.

##### Element <*connector-id*>

This element contains the connector identifier related to the display output.

##### Element <*connector-label*>

This element contains the connector label related to the display output.

##### Element <*connected*>

This element indicates if a screen is connected to the connector of the related display output 

##### Element <*screens*>

This element contains information about the screens connected to the device.

###### Element <*screen*>

This element contains information about a screen connected to the device.

####### Element <*manufacturer-id*>

This element contains the screen manufacturer identifier.
 
####### Element <*product-name*>

This element contains the screen product name.

####### Element <*product-code*>

This element contains the screen product code.

####### Element <*serial-number*>

This element contains the screen serial number.

####### Element <*power-mode*>

This element contains the current power mode of the screen.

####### Element <*input-source*>

This element contains the current input source of the screen.

####### Element <*cec-version*>

This element contains the cec version of the screen. This element is only required if the screen supports HDMI-CEC. 

####### Element <*menu-language*>

This element contains the current menu language of the screen. 

####### Element <*manufacturer-OUI*>

This element contains the screen manufactuer OUI.

### Element <*commands*>

This element contains information about the commands run from the device.
(eg. standby screen)