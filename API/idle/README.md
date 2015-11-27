nsIIdleObserver interface Reference
===================================

    #include <nsIIdleObserver.idl>

Public Attributes
-----------------

-   readonly attribute unsigned long time

    *Time is in seconds and is read only when idle observers are added and removed.*

-   void onidle ( )

    *Callback is called when the device is in idle mode after duration defined with attribute time.*

<!-- -->

-   void onactive ( )

    *Callback is called as soon as the device receives some keyboard/mouse event.*

This interface defines the observer to add or remove from Idle service.

Member Data Documentation
-------------------------

void nsIIdleObserver::onidle ()
-------------------------------

Callback is called when the device is in idle mode after duration defined with attribute time.
void nsIIdleObserver::onactive ()
---------------------------------

Callback is called as soon as the device receives some keyboard/mouse event.
nsIIdleService interface Reference
==================================

    #include <nsIIdleService.idl>

Public Attributes
-----------------

-   readonly attribute unsigned long idleTime

    *The amount of time in milliseconds that has passed since the last user activity.*

-   void addIdleObserver ( in nsIObserver observer, in unsigned long time)

    *Add an observer to be notified when the user idles for some period of time, and when they get back from that.*

<!-- -->

-   void removeIdleObserver ( in nsIObserver observer, in unsigned long time)

    *Remove an observer registered with addIdleObserver.*

This interface lets you monitor how long the user has been 'idle', i.e.

Detailed Description
--------------------

not used their mouse or keyboard. You can get the idle time directly, but in most cases you will want to register an observer for a predefined interval. The observer will get an 'idle' notification when the user is idle for that interval (or longer), and receive a 'back' notification when the user starts using their computer again.

Here is an example that uses this API :

    <html>
    <head>
    <title></title>
    </head>
    <body bgcolor="white">
     <INPUT type="button" value="Start Idle" onClick=startIdle()>
     <INPUT type="button" value="Stop Idle" onClick=stopIdle()>
     <script language="javascript">
     var elt = document.createElement("div");
     elt.setAttribute("id", "123");
     document.body.insertBefore(elt, null);
     function appendLog(msg){
      var e = window.document.getElementById("123");
      e.innerHTML =  e.innerHTML + " <br> " + msg;
     }
     var obs = {
      time: 3, 
      onidle: function () {
         appendLog("onidle");
      },
      onactive: function () {
        appendLog("onactive");
      }
     }
     function startIdle(){
      try{
       navigator.addIdleObserver(obs);
       appendLog("startIdle ok");
      }catch(e){
       appendLog("startIdle exception = " + e);
      } 
     }
     function stopIdle(){
      try{
       navigator.removeIdleObserver(obs);
       appendLog("stopIdle ok");
      }catch(e){
       appendLog("stopIdle exception = " + e);
      } 
     }
     </script>
    </body>
    </html>

Definition at line 55 of file nsIIdleService.idl

The Documentation for this struct was generated from the following file:

-   nsIIdleService.idl

Member Data Documentation
-------------------------

### readonly attribute unsigned long nsIIdleService::idleTime

If we do not have a valid idle time to report, 0 is returned (this can happen if the user never interacted with the browser at all, and if we are also unable to poll for idle time manually).

Definition at line 64 of file nsIIdleService.idl

The Documentation for this struct was generated from the following file:

-   nsIIdleService.idl

void nsIIdleService::addIdleObserver (in nsIObserver observer, in unsigned long time)
-------------------------------------------------------------------------------------

Add an observer to be notified when the user idles for some period of time, and when they get back from that.
**.**

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">observer</td>
<td align="left"><p>the observer to be notified</p></td>
</tr>
<tr class="even">
<td align="left">time</td>
<td align="left"><p>the amount of time in seconds the user should be idle before the observer should be notified.</p></td>
</tr>
</tbody>
</table>

**Note:.**

The subject of the notification the observer will get is always the nsIIdleService itself. When the user goes idle, the observer topic is "idle" and when they get back, the observer topic is "back". The data param for the notification contains the current user idle time.

You can add the same observer twice.

Most implementations need to poll the OS for idle info themselves, meaning your notifications could arrive with a delay up to the length of the polling interval in that implementation. Current implementations use a delay of 5 seconds.

void nsIIdleService::removeIdleObserver (in nsIObserver observer, in unsigned long time)
----------------------------------------------------------------------------------------

Remove an observer registered with addIdleObserver.
**.**

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">observer</td>
<td align="left"><p>the observer that needs to be removed.</p></td>
</tr>
<tr class="even">
<td align="left">time</td>
<td align="left"><p>the amount of time they were listening for.</p></td>
</tr>
</tbody>
</table>

**Note:.**

Removing an observer will remove it once, for the idle time you specify. If you have added an observer multiple times, you will need to remove it just as many times.
