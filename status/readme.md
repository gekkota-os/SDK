#GEKKOTA:: Device Status 
http://www.innes.pro/

#Introduction to Status file 

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

~~~mermaid
graph TD
    B[status file]
    B-->C(device)
    B-->D(status);
~~~
    

## Example    
    <?xml version="1.0" encoding="UTF-8"?>
    <device-status version="1.0" xmlns="ns.innes.gekkota">
    <device>
        <!— identificator used in status file name. Can be MAC,HOSTNAME,UUID -->
        <id-type>MAC</id-type>
        <!— mac address used as id in lower case and with “-” -->
        <mac>ff-ff-ff-ff-ff-ff</mac>
        <hostname>sma200_fbar</hostname>
        <uuid>00000000-0000-0000-0001-00e04b3b3e9a</uuid>
        <ip-addresses>
            <ip-address>
                <!--Can be LAN, WLAN, ...-->
                <if-type>LAN</if-type>
                <!--Can be auto, dhcp or static-->
                <origin>auto</origin>
                <value>fc00::d141:cb74:2caa:a6a2</value>
            </ip-address>
            <ip-address>
                <if-type>LAN</if-type>
                <origin>auto</origin>
                <value>fc00::f832:bf5f:4520:5b1b</value>
            </ip-address>
            <ip-address>
                <if-type>LAN</if-type>
                <origin>dhcp</origin>
                <value>192.168.1.4</value>
            </ip-address>
            <ip-address>
                <if-type>LAN</if-type>
                <if-type>auto</if-type>
                <value>fc01::1</value>
            </ip-address>
            <ip-address>
                <if-type>LAN</if-type>
                <origin>static</origin>
                <value>192.168.156.1</value>
            </ip-address>
        </ip-addresses>
        <storage>
            <total-size>104857600</total-size>
            <used-size>63963136</used-size>
        </storage>
        <firmware>
            <version>3.10.39</version>
            <extensions>
                <extension>
                    <id>ns.innes.playzilla.extensions.example</id>
                    <label>Label d’exemple</label>
                    <version>1.10.10</version>
                </extension>
            </extension>
        </firmware>
    </device>
    <status>
        <!-- device timezone required! ISO 8601 -->
        <date>2013-10-16T14:41:50.783+02:00</date>
        <launcher>
            <!-- metadata extracted from the manifest -->
            <manifest-metadata xmlns:pzpm="ns.innes.playzilla.manifest" xmlns:ex="ns.innes.example">
                <pzpm:publish-id>e5fdaa88-7eb3-4e20-8602-1d94030d2bcc</pzpm:publish-id>
                <pzpm:publish-size>654681689</pzpm:publish-size>
                <pzpm:publish-generator>Screen Composer</pzpm:publish-generator>
                <pzpm:publish-date>2013-10-16T14:41:50.783Z</pzpm:publish-date>
                <ex:playout-id>e5fdaa88-7eb3-4e20-8602-1d94030d2bcd</ex:playout-id>
                <!-- other metadata -->
            </manifest-metadata>
            <!-- UTC timezone -->
            <start>2013-10-16T12:41:50.783Z</start>
            <state>PLAY</state>
            <!--<state>STOP</state>-->
            <!--<state>NO_CONTENT</state>-->
            <!--<state>SYSTEM</state>-->
            <!--<state>MIRE</state>-->
            <!--<state>FAILSOFT_CLEANUP</state>-->
            <!--<state>FAILSOFT</state>-->
            <messages>
                <message>
                    <!--severity can be INFO, WARN, ERROR-->
                    <severity>INFO</severity>
                    <!--subject can be SYSTEM, NETWORK, FILESYSTEM OR PLAYOUT-->
                    <subject>SYSTEM</subject>
                    <!--description is free and in english-->
                    <description></description>
                    <!--date in UTC timezone-->
                    <date>2013-10-16T12:41:50.783Z </date>
                </message>
            </messages>
        </launcher>
        <screens>
            <screen>
                <state>on</state>
            </screen>
            <screen>
                <state>off</state>
            </screen>
        </screens>
        <commands>
            <command>
                <urn>urn:innes:owl:display-cmd:1#standby</urn>
                <arguments>
                    <argument>
                        <value>true</value>
                        <type>xsd:boolean</type>
                    </argument>
                </arguments>
            </command>
        </commands>
        <downloader>
            <!-- metadata extracted from the manifest -->
            <manifest-metadata xmlns:pzpm="ns.innes.playzilla.manifest" xmlns:ex="ns.innes.example">
                <pzpm:publish-id>e5fdaa88-7eb4-5653-a456-165984135795</pzpm:publish-id>
                <pzpm:publish-size>654681689</pzpm:publish-size>
                <pzpm:publish-generator>Screen Composer</pzpm:publish-generator>
                <pzpm:publish-date>2013-10-17T14:41:50.783Z</pzpm:publish-date>
                <ex:playout-id>e5fdaa88-7eb3-4e20-8602-1d94030d2bcd</ex:playout-id>
            </manifest-metadata>
            <!-- UTC timezone -->
            <start>2013-10-16T14:41:50.783Z</start>
            <!-- double between 0 and 1 -->
            <progress>0.8</progress>
            <!-- octet/s -->
            <bitrate>1111</bitrate>
            <!--state is an enum and can be : MANIFEST, INTERNAL_RESOURCES, EXTERNAL_RESOURCES, TERMINATED-->
            <state>INTERNAL_RESOURCES</state>
            <messages>
                <message>
                    <!--severity can be INFO, WARN, ERROR-->
                    <severity>ERROR</severity>
                    <!--subject can be SYSTEM, NETWORK, FILESYSTEM OR PLAYOUT-->
                    <subject>NETWORK</subject>
                    <!--description is free and in english-->
                    <description>HTTP 404 on .medias/monImage.jpg</description>
                    <!--date in UTC timezone-->
                    <date>2013-10-16T12:41:50.783Z </date>
                </message>
            </messages>
        </downloader>
    </status>
	</device-status>



## Element <*device*>

This element contains the device configuration information : 

- <*id-type*>;
- <*mac*>;
- <*hostname*>;
- <*uuid*>;
- <*modelName*>;
- <*modelNumber*>;
- <*serialNumber*>;
- <*field[1-5]*>;
- <*ip-adresses*>;
- <*addons*>.

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

- "LAN" for LAN network interface;
-  "WLAN" for WLAN network interface.

#### Element <*origin*>

This element is required by <*ip-address*>.
It defines the origin of the IP address : 

- "auto" for an automatic IPv6 address ;
- "dhcp" for an IP address assigned by DHCP;
- "static" for a static IP address.

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

- <*date*>;
- <*storage*>;
- <*launcher*>;
- <*setup*>.

It may contain the <*downloader*> element.

### Element <*date*>
This element is required.  
It defines the current UTC date expressed in ISO-8601 format followed by the time difference of the local area (eg.  "2015-11-17T09:32:27.402+01:00"). 

### Element <*storage*>
This element is required.  
It describes the data storage space of the platform.  
It must contain the following elements :

- <*total-size*>;
- <*used-size*>;

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

- <*configuration*>;
- <*installer*>.

#### Element <*configuration*>
This element is required.
It describes a configuration extension installed on the platform.  
It must contain the following elements :

- <*version*>;
- <*metadatas*>.

#### Element <*installer*>
This element is required.  
it describes a software installation extension installed on the platform.  
It must contain the following elements :  

- <*version*>;
- <*metadatas*>;

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


### Element <*screens*>

This element contains informations about the screens connected to the device.

### Element <*commands*>

This element contains informations about the commands run from the device.
(eg. standby screen)