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

The nsISystemAPBGpioOutput interface allows to write value to Gpio. Here is an example for SMT210

    <!doctype html>
    <html>
    <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

    <style>

      #button1{
       position:absolute;
       top : 30px;
       left : 300px;
       width: 100px;
       height:80px;
       background-color:green;
       color:black;
       
      } 
      #gpio1{
       background-color:white;
       height:100px;
       font-family:arial;
       font-size:40px;
      
      
      } 
      #gpio2{
       background-color:orange;
                height:100px;
       font-family:arial;
       font-size:40px;
      
      } 
     #konsole{
       position : absolute;
       width:98%;
       top:50%;
       height:48%;
       max-height:48%;
       overflow:auto;
       background-color:black;
       color:green;
       font-family:courier;
       font-size:20px;
       border-color:green;
       border-style:solid;
       border-width:1px;
      }
      
    </style>

    <script type="text/javascript">

    var polling = 500;
    var outputIdx = 1;
    var inputIdx  = 2;
    var state1;
    var state2;

    const Ci = Components.interfaces;

    var konsole;
    var divState1;
    var divState2;

    function init ()
    {
     
        divState1 = document.getElementById("state1");
        divState2 = document.getElementById("state2");
     
     state2 = readGPIO(inputIdx);
        state1 = false;
     writeGPIO(outputIdx, state1);
     
     divState1.innerHTML = showState(state1);
     divState2.innerHTML = showState(state2);
     
     forEachGpio();
    };

    function showState(state)
    {
     if (state == true) return "state : 1"; 
     
     if (state == false) return "state : 0"; 
     
     return "state : ***";
    }

    function log(value) {

    };

    var observer = {
      onChange: function onChange(aAPBGpioInput, aOldValue, aNewValue) {
          
       state2 = aNewValue;
       divState2.innerHTML = showState(state2);
      }
     }; 
     
    function registerProfile(profile, connector) {
      log("register profile");
      profile.addObserver(observer);
      return true;
     }

    function forEachGpio() {

        var ifaceName = "gpio-input";
     var iface = Ci.nsISystemAPBGpioInput;

      log("forEachGpio " + ifaceName +  " " + iface);
      var profiles = systemManager.getApplicationProfileBindingsByProfileUri(ifaceName);
      log(profiles);
      if(profiles != null){
       log("profiles length" + profiles.length);
       for (var i = 0; i < profiles.length; ++i){
        var profile = profiles.queryElementAt(i, iface);
        var connector = getFirstConnector(profile);
        if(connector.id.indexOf("phoenix") == 0) {
         if(!registerProfile(profile, connector)) {
          return;
         }
        }
       }
      }
     } 
     
    function writeGPIO(index, value)
    {
     try{
        var profiles = systemManager.getApplicationProfileBindingsByProfileUri("gpio-output");
        var mask = 1 << (index - 1);
        if(profiles != null){

        
         for (var i = 0; i < profiles.length; ++i){
          var profile = profiles.queryElementAt(i, Components.interfaces.nsISystemAPBGpioOutput);
          var connector = getFirstConnector(profile);
          if(connector.mask == mask){
           profile.writeValue = value;
           log("writeGPIO idx" + index + " value " + value);
           return;
          }
         }
        }
      }
      catch(e)
      {
       log(e.toString());
      }
    }

    function getFirstConnector(profile) {
      var connectors = profile.adapter.connectors;
      for(var j = 0; j < connectors.length; ++j) {
       var connector = connectors.queryElementAt(j, Ci.nsISystemConnector);
       return connector;
      }
     }
     
    function readGPIO(index){

    try{
      var profiles = systemManager.getApplicationProfileBindingsByProfileUri("gpio-input");
      var mask = 1 << (index - 1);
      if(profiles != null){

       //log("looking at profiles : " + profiles.length);
       for (var i = 0; i < profiles.length; ++i){
        var profile = profiles.queryElementAt(i, Components.interfaces.nsISystemAPBGpioInput);
        var connector = getFirstConnector(profile);
        if(connector.mask == mask){
         value = profile.readValue;
         //log("readGPIO idx " + index + " value " + value);
         return value;
        }
       }
      }
      }
      catch(e)
      {
       log(e.toString());
      }
     }
     
    function action()
    {  
      state1 = !state1;
      divState1.innerHTML = showState(state1);

      writeGPIO(outputIdx, state1);
    }

    </script>
    </head>

    <body onload="init()">
    <h1></h1>
    <div id="gpio1">GPIO 1 Output
    <div id="state1"></div>
    <input id="button1" type="button" value="change state" onclick="action()" ></input>
    </div>
    <div id="gpio2">GPIO 2 Input
    <div id="state2">
    </div>
    </div>

    <!--div id=konsole> LOG
    </div-->

    <br><br><br>


    </body>

Member Data Documentation
-------------------------

### attribute boolean nsISystemAPBGpioOutput::writeValue

Value to set on the gpio

void nsISystemAPBGpioOutput::removeObserver (in nsISystemGpioOuputObserver aObserver)
-------------------------------------------------------------------------------------

Remove observer on service provider

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>observer removed</p></td>
</tr>
</tbody>
</table>

void nsISystemAPBGpioOutput::addObserver (in nsISystemGpioOuputObserver aObserver)
----------------------------------------------------------------------------------

Add observer on service provider

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>observer added</p></td>
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

The nsISystemAPBGpioInput interface allows to read the Gpio

Member Data Documentation
-------------------------

### readonly attribute boolean nsISystemAPBGpioInput::readValue

Read the value of the gpio

void nsISystemAPBGpioInput::removeObserver (in nsISystemGpioObserver aObserver)
-------------------------------------------------------------------------------

Remove observer on service provider

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>observer removed</p></td>
</tr>
</tbody>
</table>

void nsISystemAPBGpioInput::addObserver (in nsISystemGpioObserver aObserver)
----------------------------------------------------------------------------

Add observer on service provider

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>observer added</p></td>
</tr>
</tbody>
</table>

nsISystemGpioOuputObserver interface Reference
==============================================

-   void onChange ( in nsISystemAPBGpioOutput aAPBGpioOutput, in boolean aOldValue, in boolean aNewValue)

void nsISystemGpioOuputObserver::onChange (in nsISystemAPBGpioOutput aAPBGpioOutput, in boolean aOldValue, in boolean aNewValue)
--------------------------------------------------------------------------------------------------------------------------------

Callback called on changes

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAPBGpioOutput</td>
<td align="left"><p>APB Gpio Output</p></td>
</tr>
<tr class="even">
<td align="left">aOldValue</td>
<td align="left"><p>old value</p></td>
</tr>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>new value</p></td>
</tr>
</tbody>
</table>

nsISystemGpioObserver interface Reference
=========================================

-   void onChange ( in nsISystemAPBGpioInput aAPBGpioInput, in boolean aOldValue, in boolean aNewValue)

void nsISystemGpioObserver::onChange (in nsISystemAPBGpioInput aAPBGpioInput, in boolean aOldValue, in boolean aNewValue)
-------------------------------------------------------------------------------------------------------------------------

Callback : onChange

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAPBGpioInput</td>
<td align="left"><p>nsISystemAPBGpioInput</p></td>
</tr>
<tr class="even">
<td align="left">aOldValue</td>
<td align="left"><p>old value boolean</p></td>
</tr>
<tr class="odd">
<td align="left">aNewValue</td>
<td align="left"><p>new value boolean</p></td>
</tr>
</tbody>
</table>


