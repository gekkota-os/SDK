nsISystemAdapterAppColorLed interface Reference
===============================================

Public Attributes
-----------------

-   const PRUint16 LEFT\_COLOR\_LED

<!-- -->

-   const PRUint16 RIGHT\_COLOR\_LED

<!-- -->

-   attribute boolean enabled

<!-- -->

-   attribute PRUint32 color

<!-- -->

-   readonly attribute PRUint16 placement

Detailed Description
--------------------

The nsISystemAdapterAppColorLed interface allows to manage color led. Here is an example

    <!doctype html>
    <html>
    <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

    <style>

      #button1{
       width:100px;
       height:100px;
       background-color:red;
       color:black;
      } 
      #button2{
       width:100px;
       height:100px;
       background-color:green;
       color:black;
      } 
      #button3{
       width:100px;
       height:100px;
       background-color:blue;
       color:black;
      } 
      #button4{
       width:100px;
       height:100px;
       background-color:#E54F04;
       color:black;
      } 
      #button5{
       width:100px;
       height:100px;
       background-color:#09E5DE;
       color:black;
      } 
     #konsole{
       width:100%;
       height:100%;
       background-color:white;
       color:black;
      }
      
    </style>

    <script type="text/javascript">

    // LED parameters
    var adaptersList;
    var Ci;
    var ledcolor;


    function init ()
    {
        var konsole = document.getElementById("konsole");
      
      
     // Init    
     netscape.security.PrivilegeManager.enablePrivilege("UniversalXPConnect");
     Ci = Components.interfaces;
     adaptersList = systemManager.getAdaptersByIId(Ci.nsISystemAdapterAppColorLed);
     ledcolor=0x77FF00; // RGB hex  
    }


    function log(value) {
     konsole.innerHTML += "<li>"+value+"</li>";
     console.log(value);
     
     //console.log(userAgent);
     
     console.log(this.navigator.userAgent);
        console.log(this.navigator.appName);
    };

    function error(value) {
     var now = new Date();
     
     konsole.innerHTML += "<li>"+ now + " : " + value+"</li>";
     console.error(value);
    };


    function action()
    {  
      netscape.security.PrivilegeManager.enablePrivilege("UniversalXPConnect");
      
     // Set Color
    if (adaptersList.length > 0)
       {
        var appColorLed = adaptersList.queryElementAt(Ci.nsISystemAdapterAppColorLed.LEFT_COLOR_LED,Ci.nsISystemAdapterAppColorLed);
        appColorLed.enabled = true;
        appColorLed.color = ledcolor;
      
        appColorLed = adaptersList.queryElementAt(Ci.nsISystemAdapterAppColorLed.RIGHT_COLOR_LED,Ci.nsISystemAdapterAppColorLed);
        appColorLed.enabled = true;
        appColorLed.color = ledcolor;
       }  


    }

    function action1()
    { 
     ledcolor=0XFF0000;
     action();
    }
    function action2()
    { 
     ledcolor=0X00FF00;
     action();
    }
    function action3()
    { 
     ledcolor=0X0000FF;
     action();
    }
    function action4()
    { 
     ledcolor=0XE54F04;
     action();
    }
    function action5()
    { 
     ledcolor=0X09E5DE;
     action();
    }

    </script>
    </head>

    <body onload="init()">
    <h1>LED</h1>

    <input id="button1" type="button" onclick="action1()" ></input>
    <input id="button2" type="button" onclick="action2()" ></input>
    <input id="button3" type="button" onclick="action3()" ></input>
    <input id="button4" type="button" onclick="action4()" ></input>
    <input id="button5" type="button" onclick="action5()" ></input>

    <br><br><br>


    </body>

Member Data Documentation
-------------------------

### const PRUint16 nsISystemAdapterAppColorLed::LEFT\_COLOR\_LED

placement left

### const PRUint16 nsISystemAdapterAppColorLed::RIGHT\_COLOR\_LED

placement right

### attribute boolean nsISystemAdapterAppColorLed::enabled

enable or disable the color led

### attribute PRUint32 nsISystemAdapterAppColorLed::color

color RGB (XXRRGGBB) of the color led

### readonly attribute PRUint16 nsISystemAdapterAppColorLed::placement

placement of the color led (LEFT\_COLOR\_LED or RIGHT\_COLOR\_LED)
