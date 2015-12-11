nsIGktUDPSocket interface Reference
===================================

Public Attributes
-----------------

-   readonly attribute AUTF8String localAddress

<!-- -->

-   readonly attribute long localPort

<!-- -->

-   readonly attribute AUTF8String remoteAddress

<!-- -->

-   readonly attribute long remotePort

<!-- -->

-   readonly attribute boolean loopback

<!-- -->

-   readonly attribute unsigned long bufferedAmount

<!-- -->

-   readonly attribute AUTF8String readyState

<!-- -->

-   attribute nsIUDPEventHandler onerror

<!-- -->

-   attribute nsIUDPEventHandler onmessage

-   void init ( in long aPort, in boolean aLoopbackOnly)

<!-- -->

-   void joinMulticast ( in AUTF8String aMulticastAddress, in AUTF8String aIface)

<!-- -->

-   void leaveMulticast ( in AUTF8String aMulticastAddress, in AUTF8String aIface)

<!-- -->

-   void close ( )

<!-- -->

-   void suspend ( )

<!-- -->

-   void resume ( )

<!-- -->

-   boolean sendArray ( in octet data, in unsigned long dataLength, in AUTF8String remoteAddress, in long remotePort)

<!-- -->

-   boolean send ( in jsval data, in AUTF8String remoteAddress, in long remotePort)

Detailed Description
--------------------

The nsIGktUDPSocket interface defines attributes and methods for UDP communication.

In javascript, instantiating this object with the following code : new GktUDPSocket ({ option\_object})

The option object allows the following properties (corresponding to the attributes listed below): localPort, localAddress, remoteAddress, remotePort, \* loopback.

    var socket = new UDPSocket({"localPort":1900, 
        "remoteAddress":"192.168.0.23", "remotePort":1800});
     socket.onmessage= function (UDPMessageEvent) {
          console.log("Remote address: " + UDPMessageEvent.remoteAddress + 
              " Remote port: " + UDPMessageEvent.remotePort +  
              " Received data" + UDPMessageEvent.data);
          };  
     socket.send("message");

Here is an example of using UDP Socket

    <html>
    <head>
    <meta http-equiv="content-type" content="text/html; charset=UTF-8"/>
     <script type="text/javascript" language="JavaScript">
     var socket = null;
    //var localAddress="192.168.1.29";
    var remoteAddress="fc00::21c:e6ff:fe02:2";
    //var remoteAddress="fc00::7e66:9dff:fe2c:11b4";
    var localAddress="fc00::8930:33b6:9486:2013";
    //var remoteAddress="192.168.0.4";
    //var remoteAddress="fc00::4";
    //var remotePort=1800;
    //var localPort=1810;
    var remotePort=5110;
    var localPort=5110;
    function ERROR(string) 
    {
     dump("*** ERROR *** " + string + "\n");
     appendConsole("\n" + string);
    } 
    function LOG(string) 
    {
     dump("*** LOG *** " + string + "\n");
     appendConsole(string);
    } 
    function appendConsole(str)
    {
        var console=document.getElementById("console")
        if (console)
            console.value =console.value + str;
    }
    function setStr(data, offset, str)
    {
     var l= str.length;
        for (i = 0; i < l; i++)
     {
      data[offset+i]  = str.charCodeAt(i);
     }
    }
    function ArrayBufferToString(arr)
    {
     var str = '';
     for (var i = 0; i < arr.length; i++) {
      str += String.fromCharCode(arr[i]);
     }
     return str;
    }
    function testNoDefault()
    {
     try{
      LOG("\ntestNoDefault ... ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "remotePort":remotePort}); }
      socket.send("test toto\n");
     }
     catch (ex)
     {
      var msg = "" + ex;
      if (msg.indexOf("DOMException InvalidAccessError : no default remote address defined") != -1)
      {
       LOG("OK");
      }
      else
      {
       ERROR("\n    exception = " + ex);
      }
     }
        socket = null;
    }

    function testSameLocal()
    {
     try{
      LOG("\ntestSameLocal ... ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
       socket = new GktUDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      socket.send("test toto\n");
      LOG("OK");
     }
     catch (ex)
     {
      ERROR("\n    exception = " + ex);
     }
        socket = null;
    }
    function testSendAfterClose()
    {
     try{
      LOG("\ntestSendAfterClose ... ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      socket.close();
      socket.send("test toto\n");
     }
     catch (ex)
     {
      var msg = "" + ex;
      if (msg.indexOf("DOMException InvalidStateError : socket not open") != -1)
      {
       LOG("OK");
      }
      else
      {
       ERROR("\nexception = " + ex);
      }
     }
        socket = null;
    }
    function testSend()
    {
    // On the remote host, you can show the sended UPD message with command :
    // nc -ul 1800
     try{
      LOG("\ntestSend ... ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress, "localAddress":localAddress});
      }
      socket.send("test toto\n");
      var tab = [65, 51,52,53, 0xA, 0];
      socket.send(tab, "192.168.0.4", 0);
      var data = [];
      setStr(data, 0, "new Uint8Array\n");
      var uint8 = new Uint8Array(data);
      socket.send(uint8);
      socket.close();
      LOG("OK");
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
        socket = null;
    }
    function testReceive()
    {
     try{
    // On the remote host, you can send UPD message with command :
    // echo -n "hello" | nc  -u  -w 5 192.168.1.29 1810
      LOG("\ntestReceive  (send message on port '" + localPort + "')... ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort,
                            "remotePort":remotePort,
                            "remoteAddress":remoteAddress,"localAddress":localAddress});
      }
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
       socket.send("receive '" + data + "'", UDPMessageEvent.remoteAddress, UDPMessageEvent.remotePort);
      };  
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
      }
      window.setTimeout(function () {end()} , 30000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testReceiveSuspend()
    {
     try{
    // On the remote host, you can send UPD message with command :
    // echo -n "hello" | nc  -u  -w 5 192.168.1.29 1810
      LOG("\ntestReceiveSuspend  (send message on port '" + localPort + "')... ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "remotePort":remotePort, "remoteAddress":remoteAddress});
      }
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
       socket.send("receive '" + data + "'", UDPMessageEvent.remoteAddress, UDPMessageEvent.remotePort);
      };  
      socket.suspend();
      function resume()
      {
       socket.resume();
       window.setTimeout(function () {end()}, 1000);
      }
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
      }
      window.setTimeout(function () {resume()} , 20000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function init()
    {
     //testNoDefault();
     //testSameLocal();
     //testSendAfterClose();
     testSend();
     //testReceive();
     //testReceiveSuspend();
    } 
    </script>
    <body onload="setTimeout('init()', '10')" ">
       <textarea id="console" cols="80" rows="60"> </textarea>
    </body>
    </head>
    </html>

and an example of using UDP multicast

    <html>
    <head>
    <meta http-equiv="content-type" content="text/html; charset=UTF-8"/>
     <script type="text/javascript;version=1.8" language="JavaScript">
    var gLogger; 
    gLogger = log4Service.getLogger("test.upnp-multicast");
     var socket = null;
    var tabLocalAddress = [];
    var tabLocalAddressIpv6 = [];
    var localAddress=undefined;
    var localAddressIpv6=undefined;
    function ERROR(string) 
    {
     dump("*** ERROR *** " + string + "\n");
     appendConsole("\n" + string);
    } 
    function LOG(string) 
    {
     dump("*** LOG *** " + string + "\n");
     gLogger.debug (string, null);
     appendConsole(string);
    } 
    function appendConsole(str)
    {
        var console=document.getElementById("console")
        if (console)
            console.value =console.value + str;
    }
    function setStr(data, offset, str)
    {
     var l= str.length;
        for (i = 0; i < l; i++)
     {
      data[offset+i]  = str.charCodeAt(i);
     }
    }
    function getLocalAddress()
    {
     const Ci = Components.interfaces;
     let ipAdapters = systemManager.getAdaptersByIId(Ci.nsISystemAdapterIP);
     let hasIPv4 = false;
     let hasIPv6 = false;
     for(let i = 0; i < ipAdapters.length; ++i) {
      let ip = ipAdapters.queryElementAt(i, Ci.nsISystemAdapterIP);
      if (ip.status != Ci.nsISystemAdapterIP.STATUS_UP){
       continue;
      }
      let addresses = ip.unicastAddresses;
      for(let j = 0; j < addresses.length; ++j){
       let addr = addresses.queryElementAt(j, Ci.nsISystemIPAddress);
       let struct = addr.address;
       let addrStr = struct.stringAddress;
       let ipType = "IPv4";
       try{
        let structv6 = struct.QueryInterface(Ci.nsISystemIPv6Structure);
        if(addrStr.indexOf("fe80") == 0){
         continue;
        }
        ipType = "IPv6";
       }catch(ex){}
       if (ipType == "IPv4")
       {
        LOG("getLocalAddress : add Ipv4 '" + addrStr + "'\n");
                    tabLocalAddress[tabLocalAddress.length] = addrStr;
                    if (localAddress == undefined)
         localAddress = addrStr;
       }
       if (!hasIPv6 && ipType == "IPv6")
       {
        LOG("getLocalAddress : add Ipv6 '" + addrStr + "'\n");
        tabLocalAddressIpv6[tabLocalAddressIpv6.length] = addrStr;
        if (localAddressIpv6 == undefined)
         localAddressIpv6 = addrStr;
       }
      }
     }
     LOG("localAddress= " + localAddress + " localAddressIpv6=" + localAddressIpv6);
    }
    function joinMulticast(socket, multicastGroup, tabAddress)
    {
     for (let i = 0; i < tabAddress.length; i++)
     {
      socket.joinMulticast(multicastGroup, tabAddress[i]);
     }
    }
    function leaveMulticast(socket, multicastGroup, tabAddress)
    {
     for (let i = 0; i < tabAddress.length; i++)
     {
      socket.leaveMulticast(multicastGroup, tabAddress[i]);
     }
    }
    function ArrayBufferToString(arr)
    {
     var str = '';
     for (var i = 0; i < arr.length; i++) {
      str += String.fromCharCode(arr[i]);
     }
     return str;
    }
    function testMulticastIpv4Local()
    {
     var localPort=30001;
     var multicastGroup="239.255.0.1";
     var remoteAddress = multicastGroup;
     try{
      LOG("\n testMulticastIpv4Local");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "localAddress":localAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "localAddress":localAddress});
      }
      try {
            socket.joinMulticast(multicastGroup);
      }
      catch (ex)
      {
       LOG("\n    Join not autorized with local adresse : OK");
      }
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testMulticastIpv6Local()
    {
     var localPort=30001;
     var multicastGroupIpv6="ff02::2";
     try{
      LOG("\n testMulticastIpv6Local  ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort,  "localAddress":localAddressIpv6});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "localAddress":localAddressIpv6});
      }
            try {
       socket.joinMulticast(multicastGroupIpv6);
      }
      catch (ex)
      {
       LOG("\n    Join not autorized with local adresse : OK");
      }
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testMulticastIpv6AnyIpv4()
    {
     var localPort=30001;
     var multicastGroupIpv6="ff02::2";
     try{
      LOG("\n testMulticastIpv6AnyIpv4  ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort});
      }
            try {
       socket.joinMulticast(multicastGroupIpv6);
      }
      catch (ex)
      {
       LOG("\n    Join not autorized with different ip family : OK");
      }
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testMulticastIpv4Any()

    {
    /*
    * On a remote host,  you can send UPD message with command :
    * echo "test" | netcat  -v -u  -p 30001  239.255.0.1 30001
    */
     var localPort=30001;
     var multicastGroup="239.255.0.1";
     var remoteAddress = multicastGroup;
     try{
      LOG("\n testMulticastIpv4Any");
            LOG("On a remote host,  you can send UPD message with command : \n\
     echo \"test\" | netcat  -v -u  -p 30001  239.255.0.1 30001");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort});
      }
            joinMulticast(socket, multicastGroup, tabLocalAddress);
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
       socket.send("receive '" + data + "'", multicastGroup, localPort);
      };  
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
                leaveMulticast(socket, multicastGroup, tabLocalAddress);
      }
      window.setTimeout(function () {end()} , 30000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testMulticastIpv6Any()
    {
    /*
    * On a remote host,  you can send UPD message with command :
    * echo "test" |  netcat  -v -u  -p 30001  "ff02::2%eth0" 30001
    */
     var localPort=30001;
     var multicastGroupIpv6="ff02::2";
     try{
      LOG("\n testMulticastIpv6Any");
      LOG("\n On a remote host,  you can send UPD message with command :\n\
       echo \"test\" |  netcat  -v -u  -p 30001  \"ff02::2%eth0\" 30001");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort,  "localAddress":"::"});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "localAddress":"::"});
      }
            joinMulticast(socket, multicastGroupIpv6, tabLocalAddressIpv6);
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
       socket.send("receive '" + data + "'", multicastGroup, localPort);
      };  
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
                leaveMulticast(socket, multicastGroupIpv6, tabLocalAddressIpv6);
      }
      window.setTimeout(function () {end()} , 30000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testMulticastMultipleIpv4Any()
    {
    /*
    * On a remote host,  you must send UPD message with command :
    * echo "test" | netcat  -v -u  -p 30001  239.255.0.1 30001
    * On a other remote host,  you must send UPD message with command :
    * echo "test" | netcat  -v -u  -p 30001  239.255.0.2 30001
    */
     var localPort=30001;
     var multicastGroup1="239.255.0.1";
     var multicastGroup2="239.255.0.2";
     try{
      LOG("\n testMulticastMultipleIpv4Any  ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort});
      }
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
       socket.send("receive '" + data + "'", multicastGroup, localPort);
      };  
            joinMulticast(socket, multicastGroup1, tabLocalAddress);
            joinMulticast(socket, multicastGroup2, tabLocalAddress);
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
       leaveMulticast(socket, multicastGroup1, tabLocalAddress);
       leaveMulticast(socket, multicastGroup2, tabLocalAddress);
      }
      window.setTimeout(function () {end()} , 30000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testMulticastLoopbackIpv4Any()
    {
    /*
    * On a remote host,  you can send UPD message with command :
    * echo "test" | netcat  -v -u  -p 30001  239.255.0.1 30001
    */
     var localPort=30001;
     var multicastGroup="239.255.0.1";
     var remoteAddress = multicastGroup;
     try{
      LOG("\n testMulticastLoopbackIpv4Any  ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "remoteAddress":remoteAddress, "localAddress":localAddress});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort, "loopback":true});
      }
            socket.joinMulticast(multicastGroup, localAddress);
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      var loopback = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
       if (data.indexOf ("test loopback") != -1 && 
        UDPMessageEvent.remoteAddress == localAddress)
       {
        loopback = true;
        LOG("\n    Loopback OK");
       }
      };  
            socket.send("test loopback", multicastGroup, localPort);                
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
                socket.leaveMulticast(multicastGroup);
      }
      window.setTimeout(function () {end()} , 30000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }

    function testMulticastUpnpIpv4Any()
    {
     var localPort=1900;
     var multicastGroup="239.255.255.250";
     var remoteAddress = multicastGroup;
     try{
      LOG("\n testMulticastUpnpIpv4Any  ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort});
      }
            joinMulticast(socket, multicastGroup, tabLocalAddress);
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
      };  
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
                leaveMulticast(socket, multicastGroup, tabLocalAddress);
      }
      window.setTimeout(function () {end()} , 10000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function testMulticastUpnpIpv6Any()
    {
     var localPort=1900;
     var multicastGroupIpv6="ff02::c"
     var remoteAddress = multicastGroupIpv6;
     try{
      LOG("\n testMulticastUpnpIpv6Any  ");
      if (typeof UDPSocket === 'function')
      {
       socket = new UDPSocket({"localPort":localPort, "localAddress":"::"});
      }
      else
      { 
       socket = new GktUDPSocket({"localPort":localPort,"localAddress":"::"});
      }
            joinMulticast(socket, multicastGroupIpv6, tabLocalAddressIpv6);
      socket.onerror = function (UDPErrorEvent) {
       LOG ("onerror - name = " + UDPErrorEvent.name +
         " message = " + UDPErrorEvent.message);
      };
      var received = false;
      socket.onmessage = function (UDPMessageEvent) {
       var data = ArrayBufferToString (UDPMessageEvent.data);
       LOG ("\n   onmessage  - Remote address: " + UDPMessageEvent.remoteAddress + 
         " Remote port: " + UDPMessageEvent.remotePort +  
         " Received data : " + data);
       received = true;
      };  
      function end()
      {
       if (received)
       {
        LOG("\n    OK");
       }
       else
       {
        LOG("Nothing received");
       }
                leaveMulticast(socket, multicastGroupIpv6, tabLocalAddressIpv6);
      }
      window.setTimeout(function () {end()} , 30000);
      
     }
     catch (ex)
     {
      ERROR("exception = " + ex);
     }
    }
    function init()
    {
     getLocalAddress();
     //testMulticastIpv4Local();
     //testMulticastIpv6Local();
     //testMulticastIpv6AnyIpv4();
     //testMulticastIpv4Any();
     //testMulticastIpv6Any();
     //testMulticastMultipleIpv4Any();
     //testMulticastLoopbackIpv4Any();
     //testMulticastUpnpIpv4Any();
     testMulticastUpnpIpv6Any();
     //testMulticastUpnpIpv6();
    } 
    </script>
    <body onload="setTimeout('init()', '10')" ">
       <textarea id="console" cols="80" rows="60"> </textarea>
    </body>
    </head>
    </html>

Member Data Documentation
-------------------------

### readonly attribute AUTF8String nsIGktUDPSocket::localAddress

The IPv4/6 address of the interface, e.g. wifi or 3G, that the UDPSocket object is bound to. Can be set by the options argument in the constructor. If not set the user agent binds the socket to the IPv4/6 address of the default local interface.

### readonly attribute long nsIGktUDPSocket::localPort

The local port that the UDPSocket object is bound to. Can be set by the options argument in the constructor. If not set the user agent binds the socket to a random local port number.

### readonly attribute AUTF8String nsIGktUDPSocket::remoteAddress

The default remote IPv4/6 address that is used for subsequent send() calls. Null if not stated by the options argument of the constructor.

### readonly attribute long nsIGktUDPSocket::remotePort

The default remote port that is used for subsequent send() calls. Null if not stated by the options argument of the constructor

### readonly attribute boolean nsIGktUDPSocket::loopback

Only applicable for multicast. true means that sent multicast data is looped back to the sender. Can be set by the options argument in the constructor. Default is false.

### readonly attribute unsigned long nsIGktUDPSocket::bufferedAmount

This attribute contains the number of bytes which have previously been buffered by calls to the send methods of this socket.

### readonly attribute AUTF8String nsIGktUDPSocket::readyState

The state of the UDP Socket object. A UDP Socket object can be in "open" or "closed" states.

### attribute nsIUDPEventHandler nsIGktUDPSocket::onerror

The onerror event handler will be called when there is an error. The data attribute of the event passed to the onerror handler will have a description of the kind of error.

### attribute nsIUDPEventHandler nsIGktUDPSocket::onmessage

Event handler for received UDP data. The onmessage handler will be called repeatedly and asynchronously after the UDPSocket object has been created, every time a UDP datagram has been received and was read. At any time, the client may choose to pause reading and receiving onmessage callbacks, by calling the socket's suspend() method. Further invocations of onmessage will be paused until resume() is called.

void nsIGktUDPSocket::init (in long aPort, in boolean aLoopbackOnly)
--------------------------------------------------------------------

This method initializes an UDP socket.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aPort</td>
<td align="left"><p>The port of the UDP socket. Pass -1 to indicate no preference, and a port will be selected automatically.</p></td>
</tr>
<tr class="even">
<td align="left">aLoopbackOnly</td>
<td align="left"><p>If true, the UDP socket will only respond to connections on the , * local loopback interface. Otherwise, it will accept connections from any interface. To specify a particular network interface, use initWithAddress.</p></td>
</tr>
</tbody>
</table>

void nsIGktUDPSocket::joinMulticast (in AUTF8String aMulticastAddress, \[optional\] in AUTF8String aIface)
----------------------------------------------------------------------------------------------------------

Join the multicast group specified by the given adress. You are then able to receive future datagrams addressed to the group. To join a multicast group, the local adress of the socket must be setted to a anycast address ("0.0.0.0" on IpV4, "::" on Ipv6), otherwise an error will be returned.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aMulticastAddress</td>
<td align="left"><p>The multicast group address. if empty, the value of attribute remoteAddress is used</p></td>
</tr>
<tr class="even">
<td align="left">aIface</td>
<td align="left"><p>The local address of the interface on which to join the group. If this is not specified, the OS may join the group on all interfaces or only the primary interface.</p></td>
</tr>
</tbody>
</table>

void nsIGktUDPSocket::leaveMulticast (in AUTF8String aMulticastAddress, \[optional\] in AUTF8String aIface)
-----------------------------------------------------------------------------------------------------------

Leave the multicast group specified by the given adress. You will no longer receive future datagrams addressed to the group.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aMulticastAddress</td>
<td align="left"><p>The multicast group address. if empty, the value of attribute remoteAddress is used</p></td>
</tr>
<tr class="even">
<td align="left">aIface</td>
<td align="left"><p>The local address of the interface on which to join the group. If this is not specified, the OS may join the group on all interfaces or only the primary interface.</p></td>
</tr>
</tbody>
</table>

void nsIGktUDPSocket::close ()
------------------------------

Closes the UDP socket. A closed UDP socket can not be used any more.

void nsIGktUDPSocket::suspend ()
--------------------------------

Pause reading incoming UDP data and invocations of the onmessage handler until resume is called.

void nsIGktUDPSocket::resume ()
-------------------------------

Pause reading incoming UDP data and invocations of the onmessage handler until resume is called.

boolean nsIGktUDPSocket::sendArray (\[const, array, size\_is(dataLength)\]in octet data, in unsigned long dataLength, \[optional\] in AUTF8String remoteAddress, \[optional\] in long remotePort)
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Sends data on the given UDP socket to the given address and port.

If remoteAddress and remotePort arguments are not given or null the destination is the default address and port given by the UDPSocket constructor's options argument's remoteAddress and remotePort fields.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">data</td>
<td align="left"><p>The array data to write,</p></td>
</tr>
<tr class="even">
<td align="left">dataLength</td>
<td align="left"><p>The array data length to write.</p></td>
</tr>
</tbody>
</table>

**Returns:.**

true if the data the request has been completed, false otherwise.

boolean nsIGktUDPSocket::send (in jsval data, \[optional\] in AUTF8String remoteAddress, \[optional\] in long remotePort)
-------------------------------------------------------------------------------------------------------------------------

Sends data on the given UDP socket to the given address and port.

If remoteAddress and remotePort arguments are not given or null the destination is the default address and port given by the UDPSocket constructor's options argument's remoteAddress and remotePort fields.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">data</td>
<td align="left"><p>The data to write witch can be a string or a array of unsigned integer,</p></td>
</tr>
</tbody>
</table>

**Returns:.**

true if the data the request has been completed, false otherwise.

nsIUDPErrorEvent interface Reference
====================================

Public Attributes
-----------------

-   readonly attribute AUTF8String name

<!-- -->

-   readonly attribute AUTF8String message

Detailed Description
--------------------

The nsIUDPMessageEvent interface represents events related to a error.

Member Data Documentation
-------------------------

### readonly attribute AUTF8String nsIUDPErrorEvent::name

The error name

### readonly attribute AUTF8String nsIUDPErrorEvent::message

The error message

nsIUDPEvent interface Reference
===============================

Public Attributes
-----------------

-   readonly attribute AUTF8String type

<!-- -->

-   readonly attribute nsIGktUDPSocket target

Detailed Description
--------------------

The nsIUDPEvent interface is the base interface for all received events.

Member Data Documentation
-------------------------

### readonly attribute AUTF8String nsIUDPEvent::type

Type of event

### readonly attribute nsIGktUDPSocket nsIUDPEvent::target

The target of event (the UDP socket)

nsIUDPEventHandler interface Reference
======================================

-   void handleEvent ( in nsIUDPEvent aEvent)

Detailed Description
--------------------

The nsIUDPEventHandler interface allows to manage UDP events

void nsIUDPEventHandler::handleEvent (in nsIUDPEvent aEvent)
------------------------------------------------------------

nsIUDPMessageEvent interface Reference
======================================

Public Attributes
-----------------

-   readonly attribute AUTF8String remoteAddress

<!-- -->

-   readonly attribute unsigned short remotePort

<!-- -->

-   readonly attribute ACString dataString

<!-- -->

-   readonly attribute jsval data

Detailed Description
--------------------

The nsIUDPMessageEvent interface represents events related to received UDP data

Member Data Documentation
-------------------------

### readonly attribute AUTF8String nsIUDPMessageEvent::remoteAddress

The address of the remote machine

### readonly attribute unsigned short nsIUDPMessageEvent::remotePort

The port of the remote machine

### readonly attribute ACString nsIUDPMessageEvent::dataString

Custom data associated with this event.

### readonly attribute jsval nsIUDPMessageEvent::data

Jsval data
