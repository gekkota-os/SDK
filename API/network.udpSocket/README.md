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

In Javascript, instantiating this object with the following code :

    new GktUDPSocket ({option_object})

The option object allows the following properties (corresponding to the attributes listed below): localPort, localAddress, remoteAddress, remotePort, \* loopback.

    var socket = new UDPSocket({"localPort":1900,
        "remoteAddress":"192.168.0.23", "remotePort":1800});
     socket.onmessage= function (UDPMessageEvent) {
          console.log("Remote address: " + UDPMessageEvent.remoteAddress +
              " Remote port: " + UDPMessageEvent.remotePort +
              " Received data" + UDPMessageEvent.data);
          };
     socket.send("message");

Find an HTML example on how to use UDP Socket [here.](unicast_example1.html)

Find an HTML example on how to use UDP multicast [here.](multicast_example1.html)

Member Data Documentation
-------------------------

### readonly attribute AUTF8String nsIGktUDPSocket::localAddress

The IPv4/6 address of the interface, e.g. Wi-Fi or 3G, that the UDPSocket object is bound to. Can be set by the options argument in the constructor. If not set the user agent binds the socket to the IPv4/6 address of the default local interface.

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
<td align="left"><p>If true, the UDP socket will only respond to connections on the local loopback interface. Otherwise, it will accept connections from any interface. To specify a particular network interface, use initWithAddress.</p></td>
</tr>
</tbody>
</table>

void nsIGktUDPSocket::joinMulticast (in AUTF8String aMulticastAddress, \[optional\] in AUTF8String aIface)
----------------------------------------------------------------------------------------------------------

Join the multicast group specified by the given address. You are then able to receive future datagrams addressed to the group. To join a multicast group, the local address of the socket must be set to a anycast address ("0.0.0.0" on IpV4, "::" on Ipv6), otherwise an error will be returned.

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
<td align="left"><p>The local address of the interface on which to join the group. If this is not specified, the OS may join the group on all interfaces or only the primary interface</p></td>
</tr>
</tbody>
</table>

void nsIGktUDPSocket::leaveMulticast (in AUTF8String aMulticastAddress, \[optional\] in AUTF8String aIface)
-----------------------------------------------------------------------------------------------------------

Leave the multicast group specified by the given address. You will no longer receive future datagrams addressed to the group.

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
<td align="left"><p>The local address of the interface on which to join the group. If this is not specified, the OS may join the group on all interfaces or only the primary interface</p></td>
</tr>
</tbody>
</table>

void nsIGktUDPSocket::close ()
------------------------------

Closes the UDP socket. A closed UDP socket cannot be used any more.

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
<td align="left"><p>The array data to write</p></td>
</tr>
<tr class="even">
<td align="left">dataLength</td>
<td align="left"><p>The array data length to write</p></td>
</tr>
</tbody>
</table>

**Returns:.**

True if the data the request has been completed, false otherwise.

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
<td align="left"><p>The data to write witch can be a string or a array of unsigned integer</p></td>
</tr>
</tbody>
</table>

**Returns:.**

True if the data the request has been completed, false otherwise.

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

The error name.

### readonly attribute AUTF8String nsIUDPErrorEvent::message

The error message.

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

Type of event.

### readonly attribute nsIGktUDPSocket nsIUDPEvent::target

The target of event (the UDP socket).

nsIUDPEventHandler interface Reference
======================================

-   void handleEvent ( in nsIUDPEvent aEvent)

Detailed Description
--------------------

The nsIUDPEventHandler interface allows to manage UDP events.

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

The nsIUDPMessageEvent interface represents events related to received UDP data.

Member Data Documentation
-------------------------

### readonly attribute AUTF8String nsIUDPMessageEvent::remoteAddress

The address of the remote machine.

### readonly attribute unsigned short nsIUDPMessageEvent::remotePort

The port of the remote machine.

### readonly attribute ACString nsIUDPMessageEvent::dataString

Custom data associated with this event.

### readonly attribute jsval nsIUDPMessageEvent::data

Jsval data.
