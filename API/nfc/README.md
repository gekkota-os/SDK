nsISystemAdapterNfc interface Reference
=======================================

    #include <nsISystemAdapterNfc.idl>

Public Attributes
-----------------

-   const unsigned short MODULATION\_ISO14443A

    *type ISO 14443 A*

<!-- -->

-   const unsigned short MODULATION\_JEWEL

    *type Jewel*

<!-- -->

-   const unsigned short MODULATION\_ISO14443B

    *type ISO 14443 B*

<!-- -->

-   const unsigned short MODULATION\_ISO14443BI

    *type pre ISO 14443 B or ISO 14443 B'*

<!-- -->

-   const unsigned short MODULATION\_ISO14443B2SR

    *type ISO 14443 2B ST SRx*

<!-- -->

-   const unsigned short MODULATION\_ISO14443B2CT

    *type ISO 14443 2B ASK CTx*

<!-- -->

-   const unsigned short MODULATION\_FELICA

    *type Felica*

<!-- -->

-   const unsigned short MODULATION\_DEP

    *protocol ISO DEP*

<!-- -->

-   const unsigned short BAUD\_106

    *baud rate 106 kbit/s*

<!-- -->

-   const unsigned short BAUD\_212

    *baud rate 212 kbit/s*

<!-- -->

-   const unsigned short BAUD\_424

    *baud rate 424 kbit/s*

<!-- -->

-   const unsigned short BAUD\_847

    *baud rate 847 kbit/s*

<!-- -->

-   readonly attribute nsIArray supportedModulations

    *array of bags*

-   void configurePolling ( in PRUint32 aModulationLength, in unsigned short aModulation, in PRUint32 aBaudRateLength, in unsigned short aBaudRate)

    *Configure type of polling.*

<!-- -->

-   void configureMultipleModulations ( in nsIPropertyBag aModulations, in PRUint32 aModulationLength)

    *Specify type(s) of polling.*

<!-- -->

-   boolean isSupported ( in unsigned short aModulation, in unsigned short aBaudRate)

    *Tell if the current device supports the modulation and its baud rate.*

<!-- -->

-   nsIEnumerator scanTargets ( )

    *Detect synchronously the targets in the near field.*

<!-- -->

-   unsigned long startPoll ( in nsISystemAdapterNfcObserver aObserver)

    *Start polling to detect some targets.*

<!-- -->

-   void stopPoll ( in unsigned long aObserverId)

    *Stop polling.*

Detailed Description
--------------------

the nsISystemAdapterNfc interface is the point of entry for using NFC. Here is an example of using this interface on SMT210 device :

    <html>
    <head>
    <title></title>
    </head>
    <body bgcolor="white"><br>
     <input type="checkbox" id="useConfigurePolling" value="0"> <i> use deprecated configurePolling </i>
     <input type="checkbox" id="2_observers" value="0"> <i> use 2 observers </i>
     <input type="checkbox" id="3_observers" value="0"> <i> use 3 observers </i><br>
     <input type="checkbox" id="ISO14443A" value="0"> ISO14443A (106 kbit/s)
     <input type="checkbox" id="ISO14443B" value="0"> ISO14443B (106 kbit/s)
     <input type="checkbox" id="JEWEL" value="0"> JEWEL (106 kbit/s)
     <input type="checkbox" id="FELICA" value="0"> FELICA 
     <input type="checkbox" id="FELICA_212" value="0"> 212 kbit/s 
     <input type="checkbox" id="FELICA_424" value="0"> 424 kbit/s  <br><br>
     <INPUT type="button" value="Start poll" onClick=startPoll()>
     <INPUT type="button" value="Stop poll" onClick=stopPoll()>
     <INPUT type="button" value="Clear" onClick=clearDisplay()>
     <script type="text/javascript;version=1.8">
      const Ci = Components.interfaces;
      function appendLog(msg){
        var e = window.document.getElementById("123");
       e.innerHTML =  e.innerHTML + " <br> " + msg;
       }
      var elt = document.createElement("div");
      elt.setAttribute("id", "123");
      document.body.insertBefore(elt, null);
      
      appendLog("ready.");
      function modulationString (aMod){
       var str = "";
       switch(aMod){
        case Ci.nsISystemAdapterNfc.MODULATION_ISO14443A : 
         str = "ISO14443A";
         break;
        case Ci.nsISystemAdapterNfc.MODULATION_JEWEL  : 
         str = "JEWEL";
         break;
        case Ci.nsISystemAdapterNfc.MODULATION_ISO14443B : 
         str = "ISO14443B";
         break;
        case Ci.nsISystemAdapterNfc.MODULATION_ISO14443BI  : 
         str =  "ISO14443BI";
         break;
        case Ci.nsISystemAdapterNfc.MODULATION_ISO14443B2SR  : 
         str = "ISO14443B2SR";
         break;
        case Ci.nsISystemAdapterNfc.MODULATION_ISO14443B2CT  : 
         str = "ISO14443B2CT";
         break;
        case Ci.nsISystemAdapterNfc.MODULATION_FELICA  : 
         str = "FELICA";
         break;
        case Ci.nsISystemAdapterNfc.MODULATION_DEP  : 
         str = "DEP";
         break;
        default : break;
       }
       return str;
      }
      function baudrateString(aBr){
       var str = "";
       switch(aBr){
        case Ci.nsISystemAdapterNfc.BAUD_106  : 
         str = "106";
         break;
        case Ci.nsISystemAdapterNfc.BAUD_212 : 
         str = "212";
         break;
        case Ci.nsISystemAdapterNfc.BAUD_424 : 
         str = "424";
         break;
        case Ci.nsISystemAdapterNfc.BAUD_847 : 
         str = "847"; 
         break;
        default : break;
       }
       return str;
      };
      
      var obs1_id, nfcAdapter;
      var observerIds = [];
      var state = 0;
      var observers = [];
      var addNfcObserver = function(){
       observers.push({
        onTargetFound : function(aTarget){
         appendLog("onTargetFound");
         try{
          appendLog("id = " + aTarget.targetId + " - type : " + modulationString(aTarget.targetModulation) + " - baud rate : " + baudrateString(aTarget.targetBaudRate));
         }catch(e){
          appendLog("Exception : " + e);
         }
        },
        onTargetLost : function(aTarget){
         appendLog("onTargetLost");
         try{
          appendLog("id = " + aTarget.targetId );
         }catch(e){
          appendLog("Exception : " + e);
         }
        },
        onPollStart : function(){
         appendLog("onPollStart");
        },
        onPollStop : function(){
         appendLog("onPollStop");
         var o  = observers.indexOf(this);
         //remove current observer
         observers.splice(o, 1);
         if(observers.length == 0){
          observerIds = [];
         }
        }
       });
      }
      
      function addToModulations( m, mod, baud){
       m.push({
        modulationType: mod,
        baudRates : baud
       });
      }
      
      function startPoll(){
       try{
        if(observerIds.length > 0 ){
         appendLog("wait for the end of all observers...");
         return;
        }
        // get all NFC adapters of the current device (only one NFC adapter on SMT210)
        var adaptersList = systemManager.getAdaptersByIId(Ci.nsISystemAdapterNfc);
        var i = 0;
        for (i = 0; i < adaptersList.length; i++)
        {
         nfcadapter = adaptersList.queryElementAt(i,Ci.nsISystemAdapterNfc); 
        
         // for configurePolling
         var modulationTypes = [];
         var modulationBaudRate = [];
         // for configureMultipleModulations
         var modulations = [];
         var custom = 0;
         addNfcObserver();
         if(window.document.getElementById("2_observers").checked){
          addNfcObserver();
         }
         else if(window.document.getElementById("3_observers").checked){
          addNfcObserver();
          addNfcObserver();
         }
         if(window.document.getElementById("ISO14443A").checked){
          custom++ ;
          addToModulations(modulations, Ci.nsISystemAdapterNfc.MODULATION_ISO14443A, [Ci.nsISystemAdapterNfc.BAUD_106]);
          modulationTypes.push(Ci.nsISystemAdapterNfc.MODULATION_ISO14443A);
          if(modulationBaudRate.indexOf(Ci.nsISystemAdapterNfc.BAUD_106) == -1){
           modulationBaudRate.push(Ci.nsISystemAdapterNfc.BAUD_106);
          }
         }
         if(window.document.getElementById("ISO14443B").checked){
          custom++ ;
          addToModulations(modulations, Ci.nsISystemAdapterNfc.MODULATION_ISO14443B, [Ci.nsISystemAdapterNfc.BAUD_106]);
          modulationTypes.push(Ci.nsISystemAdapterNfc.MODULATION_ISO14443B);
          if(modulationBaudRate.indexOf(Ci.nsISystemAdapterNfc.BAUD_106) == -1){
           modulationBaudRate.push(Ci.nsISystemAdapterNfc.BAUD_106);
          }
         }
         if(window.document.getElementById("JEWEL").checked){
          custom++ 
          addToModulations(modulations, Ci.nsISystemAdapterNfc.MODULATION_JEWEL, [Ci.nsISystemAdapterNfc.BAUD_106]);            
          modulationTypes.push(Ci.nsISystemAdapterNfc.MODULATION_JEWEL);
          if(modulationBaudRate.indexOf(Ci.nsISystemAdapterNfc.BAUD_106) == -1){
           modulationBaudRate.push(Ci.nsISystemAdapterNfc.BAUD_106);
          }
         }
         if(window.document.getElementById("FELICA").checked){
          custom++ ;
          var baudRatesArray = [];
          modulationTypes.push(Ci.nsISystemAdapterNfc.MODULATION_FELICA);
          var felicaBaudrate = 0;
          if(window.document.getElementById("FELICA_212").checked){
           baudRatesArray.push(Ci.nsISystemAdapterNfc.BAUD_212);
           modulationBaudRate.push(Ci.nsISystemAdapterNfc.BAUD_212);
           felicaBaudrate ++;
          }
          if(window.document.getElementById("FELICA_424").checked){
           baudRatesArray.push(Ci.nsISystemAdapterNfc.BAUD_424);
           modulationBaudRate.push(Ci.nsISystemAdapterNfc.BAUD_424);
           felicaBaudrate ++;
          }
          if(felicaBaudrate == 0){
           appendLog("Choose a value for baud rate of FELICA modulation!");
           return;
          }
          addToModulations(modulations, Ci.nsISystemAdapterNfc.MODULATION_FELICA,baudRatesArray);
         }
         
         if(custom){
          appendLog("Start with configured modulation.");
          if(window.document.getElementById("useConfigurePolling").checked == false && 
           nfcadapter.configureMultipleModulations != undefined){
            appendLog("using configureMultipleModulations");
            nfcadapter.configureMultipleModulations(modulations, modulations.length);
          }
          else{
           appendLog("using configurePolling");
           nfcadapter.configurePolling(modulationTypes.length, modulationTypes, modulationBaudRate.length, modulationBaudRate);
          }
         }
         else{
          appendLog("Start with all supported modulations.");
         }  
         for(var i=0; i<observers.length; i++){
          var id = nfcadapter.startPoll(observers[i]);
          appendLog("startPoll id : " + id);
          observerIds.push(id);
         }
         break;
        }
       }catch(e){
        appendLog("<br>Exception : " + e);
       }
      }
      
      function stopPoll(){
       try{
        if(observerIds.length == 0 ){
         appendLog("no more observers. Please click startPoll.");
         return;
        }
        // stop polling on all observers
        for(var i=0; i<observerIds.length; i++){
         var id = observerIds[i];
         appendLog("stopPoll id : " + id);
         nfcadapter.stopPoll(id);
        }
       }catch(e){
        appendLog("<br>Exception : " + e);
       }
      }
      function clearDisplay(){
       var e = window.document.getElementById("123");
       e.innerHTML =  "";
      }
     </script>
    </body>
    </html>

Definition at line 183 of file nsISystemAdapterNfc.idl

The Documentation for this struct was generated from the following file:

-   nsISystemAdapterNfc.idl

Member Data Documentation
-------------------------

### readonly attribute nsIArray nsISystemAdapterNfc::supportedModulations

the array could be

     [ 
     { modulationType:1, baudRates:[ 1,2 ]}, 
     { modulationType:2, baudRates: [ 3 ]} 
    ] 

Definition at line 255 of file nsISystemAdapterNfc.idl

The Documentation for this struct was generated from the following file:

-   nsISystemAdapterNfc.idl

void nsISystemAdapterNfc::configurePolling (in PRUint32 aModulationLength, \[array, size\_is(aModulationLength)\] in unsigned short aModulation, in PRUint32 aBaudRateLength, \[array, size\_is(aBaudRateLength)\] in unsigned short aBaudRate)
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Configure type of polling.
Must be called before startPoll if configurePolling() is not called, polling use all supported modulations (slower).

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aModulationLength</td>
<td align="left"><p>size of array of modulation types</p></td>
</tr>
<tr class="even">
<td align="left">aModulation</td>
<td align="left"><p>array of modulation type (see constants MODULATION_*)</p></td>
</tr>
<tr class="odd">
<td align="left">aBaudRateLength</td>
<td align="left"><p>size of array of baud rates</p></td>
</tr>
<tr class="even">
<td align="left">aBaudRate</td>
<td align="left"><p>array of baud rate for each type of modulation(see constants BAUD_*)</p></td>
</tr>
</tbody>
</table>

**Returns:.**

Void.

Deprecated

use configureMultipleModulations instead

void nsISystemAdapterNfc::configureMultipleModulations (\[array, size\_is(aModulationLength)\] in nsIPropertyBag aModulations, in PRUint32 aModulationLength)
-------------------------------------------------------------------------------------------------------------------------------------------------------------

Specify type(s) of polling.
For each modulation desired, specify one or more baud rates. Must be called before startPoll. if configureMultipleModulations() is not called or if empty array, polling use all supported modulations (slower).

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aModulations</td>
<td align="left"><p>array of property bags</p>
<pre><code>[
 { modulationType : 1, baudRates : [ 1,2 ]}, 
 { modulationType : 2, baudRates : [ 3 ]} 
]</code></pre></td>
</tr>
<tr class="even">
<td align="left">aModulationLength</td>
<td align="left"><p>size of array of property bags</p></td>
</tr>
</tbody>
</table>

**Returns:.**

Void.

boolean nsISystemAdapterNfc::isSupported (in unsigned short aModulation, in unsigned short aBaudRate)
-----------------------------------------------------------------------------------------------------

Tell if the current device supports the modulation and its baud rate.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aModulation</td>
<td align="left"><p>type of modulation</p></td>
</tr>
<tr class="even">
<td align="left">aBaudRate</td>
<td align="left"><p>baud rate</p></td>
</tr>
</tbody>
</table>

**Returns:.**

true is the modulation is supported

nsIEnumerator nsISystemAdapterNfc::scanTargets ()
-------------------------------------------------

Detect synchronously the targets in the near field.
**Returns:.**

the list of nsISystemAdapterNfcTarget targets

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

unsigned long nsISystemAdapterNfc::startPoll (in nsISystemAdapterNfcObserver aObserver)
---------------------------------------------------------------------------------------

Start polling to detect some targets.
Polling is based on configured modulation if there (polling faster), or all supported modulations (polling is a bit slower).

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>nsISystemAdapterNfcObserver</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the ID for this observer

void nsISystemAdapterNfc::stopPoll (in unsigned long aObserverId)
-----------------------------------------------------------------

Stop polling.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserverId</td>
<td align="left"><p>the ID of the observer (resulting from startPoll)</p></td>
</tr>
</tbody>
</table>

nsISystemAdapterNfcObserver interface Reference
===============================================

    #include <nsISystemAdapterNfc.idl>

-   void onTargetFound ( in nsISystemAdapterNfcTarget aTarget)

    *Callback called when a target enter the near field of the device.*

<!-- -->

-   void onTargetLost ( in nsISystemAdapterNfcTarget aTarget)

    *Callback called when a target leave the near field of the device.*

<!-- -->

-   void onMessageRead ( in nsISystemAdapterNfcTarget aTarget, in nsIArray aRecords)

    *Callback called when a Peer send a message to the device.*

<!-- -->

-   void onPollStop ( )

    *Callback called when polling is disabled.*

<!-- -->

-   void onPollStart ( )

    *Callback called when polling is enabled.*

Detailed Description
--------------------

The nsISystemAdapterNfcObserver interface provides some callbacks for polling and detection of targets

Definition at line 106 of file nsISystemAdapterNfc.idl

The Documentation for this struct was generated from the following file:

-   nsISystemAdapterNfc.idl

void nsISystemAdapterNfcObserver::onTargetFound (in nsISystemAdapterNfcTarget aTarget)
--------------------------------------------------------------------------------------

Callback called when a target enter the near field of the device.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aTarget</td>
<td align="left"><p>target which enter</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterNfcObserver::onTargetLost (in nsISystemAdapterNfcTarget aTarget)
-------------------------------------------------------------------------------------

Callback called when a target leave the near field of the device.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aTarget</td>
<td align="left"><p>target which leave</p></td>
</tr>
</tbody>
</table>

void nsISystemAdapterNfcObserver::onMessageRead (in nsISystemAdapterNfcTarget aTarget, in nsIArray aRecords)
------------------------------------------------------------------------------------------------------------

Callback called when a Peer send a message to the device.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aTarget</td>
<td align="left"><p>Peer</p></td>
</tr>
<tr class="even">
<td align="left">aRecords</td>
<td align="left"><p>array of nsISystemAdapterNDEFRecord read</p></td>
</tr>
</tbody>
</table>

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

void nsISystemAdapterNfcObserver::onPollStop ()
-----------------------------------------------

Callback called when polling is disabled.
void nsISystemAdapterNfcObserver::onPollStart ()
------------------------------------------------

Callback called when polling is enabled.
nsISystemAdapterNfcTarget interface Reference
=============================================

    #include <nsISystemAdapterNfc.idl>

Public Attributes
-----------------

-   const unsigned long TARGET\_TAG

    *type tag*

<!-- -->

-   const unsigned long TARGET\_PEER

    *type peer*

<!-- -->

-   readonly attribute AString targetId

    *NFC ID Identifer.*

<!-- -->

-   readonly attribute unsigned short targetModulation

    *see MODULATION\_XXX constants*

<!-- -->

-   readonly attribute unsigned short targetBaudRate

    *see BAUD\_XXX constants*

<!-- -->

-   readonly attribute unsigned long targetType

    *Type of the target tag or peer.*

<!-- -->

-   readonly attribute boolean ndefCompatible

    *Indicate if the target can process NDEF Message.*

<!-- -->

-   readonly attribute boolean isPresent

    *Indicate if the target is still into the near field.*

-   void writeMessage ( in PRUint32 aLength, in nsISystemAdapterNDEFRecord aRecords)

    *Write NDEF message to TAG or Peer.*

Detailed Description
--------------------

The nsISystemAdapterNfcTarget interface describes the NFC targets

Definition at line 65 of file nsISystemAdapterNfc.idl

The Documentation for this struct was generated from the following file:

-   nsISystemAdapterNfc.idl

Member Data Documentation
-------------------------

### readonly attribute unsigned long nsISystemAdapterNfcTarget::targetType

Definition at line 81 of file nsISystemAdapterNfc.idl

The Documentation for this struct was generated from the following file:

-   nsISystemAdapterNfc.idl

### readonly attribute boolean nsISystemAdapterNfcTarget::ndefCompatible

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

Definition at line 86 of file nsISystemAdapterNfc.idl

The Documentation for this struct was generated from the following file:

-   nsISystemAdapterNfc.idl

### readonly attribute boolean nsISystemAdapterNfcTarget::isPresent

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>

Definition at line 91 of file nsISystemAdapterNfc.idl

The Documentation for this struct was generated from the following file:

-   nsISystemAdapterNfc.idl

void nsISystemAdapterNfcTarget::writeMessage (in PRUint32 aLength, \[array, size\_is(aLength)\] in nsISystemAdapterNDEFRecord aRecords)
---------------------------------------------------------------------------------------------------------------------------------------

Write NDEF message to TAG or Peer.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>size of array of NDEF records</p></td>
</tr>
<tr class="even">
<td align="left">aRecords</td>
<td align="left"><p>array of nsISystemAdapterNDEFRecord, NDEF records to write</p></td>
</tr>
</tbody>
</table>

<table>
<caption>Exceptions</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">NS_ERROR_NOT_IMPLEMENTED</td>
<td align="left"></td>
</tr>
</tbody>
</table>


