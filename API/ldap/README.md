nsILdapService interface Reference
==================================

Public Attributes
-----------------

-   const PRUint16 CONFIG\_OK

<!-- -->

-   const PRUint16 CONFIG\_ERROR\_SERVER\_UNAVAILABLE

<!-- -->

-   const PRUint16 CONFIG\_ERROR\_SERVER\_AUTH\_FAILED

<!-- -->

-   const PRUint16 CONFIG\_ERROR\_SEARCH\_BASE\_DN

<!-- -->

-   const PRUint16 STATE\_NONE

<!-- -->

-   const PRUint16 STATE\_STARTED

<!-- -->

-   const PRUint16 STATE\_LDAP\_URL\_ERROR

<!-- -->

-   const PRUint16 STATE\_LDAP\_INIT\_ERROR

<!-- -->

-   const PRUint16 STATE\_LDAP\_INITIALIZED

<!-- -->

-   const PRUint16 STATE\_LDAP\_PASSWORD\_ERROR

<!-- -->

-   const PRUint16 STATE\_AUTHENTICATION\_FAILED

<!-- -->

-   const PRUint16 STATE\_SERVER\_UNAVAILABLE

<!-- -->

-   const PRUint16 STATE\_AUTHENTICATION\_SUCCEEDED

<!-- -->

-   const PRUint16 STATE\_SEARCH\_FAILED

<!-- -->

-   const PRUint16 STATE\_SUCCESS\_NO\_RESULTS

<!-- -->

-   const PRUint16 STATE\_SUCCESS

<!-- -->

-   const PRUint16 STATE\_INTERNAL\_ERROR

<!-- -->

-   readonly attribute nsILdapContext context

-   bool checkAuthentication ( in nsILdapContext aContext, in AUTF8String aUserDN, in AUTF8String aPassword)

<!-- -->

-   PRUint16 checkConfig ( in AUTF8String aURL, in AUTF8String baseDN, in bool searchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted)

<!-- -->

-   nsILdapCheckPollListener createCheckListener ( )

<!-- -->

-   nsILdapSearchPollListener createSearchListener ( )

<!-- -->

-   void asyncCheckConfig ( in AUTF8String aURL, in AUTF8String baseDN, in bool searchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted, in nsILdapCheckListener aListener)

<!-- -->

-   void search ( in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in PRUint32 aAttrsLength, in string aAttrs, in PRUint32 aOffsetIndex, in PRUint32 aCountLimit, out PRUint32 aTotalCount, out PRUint32 aResCount, out nsIPropertyBag2 aRes)

<!-- -->

-   void asyncSearch ( in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in PRUint32 aAttrsLength, in string aAttrs, in PRUint32 aOffsetIndex, in PRUint32 aCountLimit, in nsILdapSearchListener aListener)

Detailed Description
--------------------

For using the nsILdapService interface, we call properties of Javascript ldapService object. Here is an example that uses this API :

    <html>
    <head>
    <title></title></head>
    <body bgcolor="white">
     <script type="text/javascript;version=1.8">
      const Ci=Components.interfaces
       function appendLog(msg){
        var e = window.document.getElementById("123");
       e.innerHTML =  e.innerHTML + " <br> " + msg;
       }
      var elt = document.createElement("div");
      elt.setAttribute("id", "123");
      document.body.insertBefore(elt, null);
      
      appendLog("start...");
      try{ 
       // edit values here
       var url = "ldap://192.168.1.51";
       var baseDN =  "dc=exchange2007,dc=innes,dc=pro";
       var searchUserDN = "cn=Administrator,cn=Users,dc=exchange2007,dc=innes,dc=pro";
       var searchIsAuth = true;
       var passwordEncrypted = false;
       var password = "abcdefgh";  
       var passwordEnc = "";
       var searchFilter = "(&(objectClass=user)(mail=*))";
       var properties = ["mail", "sAMAccountName"];
       var limitSearchCount = 20;
       
       // test creation of context
       var context = new LdapContext(url, baseDN, 
        searchIsAuth, searchUserDN, password, passwordEncrypted);
       
       /** TEST SYNCHRONOUS **/
       appendLog("TEST SYNC :");
       
       // check ldap configuration
       var checkRes = ldapService.checkConfig(url,
           baseDN,
           true,
           searchUserDN,
           password,
           false);
           
       appendLog("checkConfig result = " + checkRes);
       
       // search
       var totalCount = new Object();
       var resCount = new Object();
       var r = ldapService.search(
         context, 
         "", 
         searchFilter, 
         properties.length,
         properties,
         0,
         limitSearchCount, 
         totalCount,
         resCount);

       appendLog("search result : ");
       r.forEach( function(elt){
         try{
          for(var j=0; j<properties.length; j++){
           appendLog(elt[properties[j]]);
          }
         }catch(e){
          appendLog("onSearchResult exception = " + e);
         }
        }
       )
       
       /** TEST ASYNCHRONOUS */
       appendLog("TEST ASYNC :");
       
       // callback for asyncCheckConfig
       function onStateChanged(res){
        appendLog("onStateChanged res : "  + res);
       }
       
       // callback for asyncSearch
       function onSearchResult(status, resCount, r){
        appendLog("onSearchResult : status = " + status);
        for(var i=0; i<resCount; i++){
         try{
          for(var j=0; j<properties.length; j++){
           appendLog(r[i][properties[j]]);
          }
         }catch(e){
          appendLog("onSearchResult exception : " + e);
         }
        }
       };
     
       ldapService.asyncCheckConfig(url,
         baseDN,
         true,
         searchUserDN,
         password,
         false,
         onStateChanged);
         
       ldapService.asyncSearch(null, 
         "", 
         searchFilter, 
         properties.length,
         properties,
         0,
         limitSearchCount, 
         onSearchResult);

       
      }catch(e){
       appendLog("Exception = " + e);
      }
      function clearDisplay(){
       var e = window.document.getElementById("123");
       e.innerHTML =  "";
      }
     </script>
    </body>
    </html>

Member Data Documentation
-------------------------

### const PRUint16 nsILdapService::CONFIG\_OK

configuration ok

### const PRUint16 nsILdapService::CONFIG\_ERROR\_SERVER\_UNAVAILABLE

ldap server is not reachable (url may be invalid)

### const PRUint16 nsILdapService::CONFIG\_ERROR\_SERVER\_AUTH\_FAILED

authentication failed

### const PRUint16 nsILdapService::CONFIG\_ERROR\_SEARCH\_BASE\_DN

search base DN invalid

### const PRUint16 nsILdapService::STATE\_NONE

state none (operation is not started)

### const PRUint16 nsILdapService::STATE\_STARTED

state for operation started

### const PRUint16 nsILdapService::STATE\_LDAP\_URL\_ERROR

state for invalid LDAP url

### const PRUint16 nsILdapService::STATE\_LDAP\_INIT\_ERROR

state for connection failure

### const PRUint16 nsILdapService::STATE\_LDAP\_INITIALIZED

state for connection succesfully initialized

### const PRUint16 nsILdapService::STATE\_LDAP\_PASSWORD\_ERROR

state for invalid password (bad encryption)

### const PRUint16 nsILdapService::STATE\_AUTHENTICATION\_FAILED

state for wrong login or password

### const PRUint16 nsILdapService::STATE\_SERVER\_UNAVAILABLE

state for LDAP server unreachable

### const PRUint16 nsILdapService::STATE\_AUTHENTICATION\_SUCCEEDED

state for succeeded authentication

### const PRUint16 nsILdapService::STATE\_SEARCH\_FAILED

state for searching failure

### const PRUint16 nsILdapService::STATE\_SUCCESS\_NO\_RESULTS

state for succeeded request with empty result

### const PRUint16 nsILdapService::STATE\_SUCCESS

state for succeded request with some results

### const PRUint16 nsILdapService::STATE\_INTERNAL\_ERROR

state for internal error (not LDAP problem)

### readonly attribute nsILdapContext nsILdapService::context

default ldap context

bool nsILdapService::checkAuthentication (in nsILdapContext aContext, in AUTF8String aUserDN, in AUTF8String aPassword)
-----------------------------------------------------------------------------------------------------------------------

Check LDAP server authentication

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aContext</td>
<td align="left"><p>context that define ldap information. If null, get preferences information</p></td>
</tr>
<tr class="even">
<td align="left">aUserDN</td>
<td align="left"><p>ldap User</p></td>
</tr>
<tr class="odd">
<td align="left">aPassword</td>
<td align="left"><p>ldap password</p></td>
</tr>
</tbody>
</table>

**Returns:.**

true is authentication is ok

PRUint16 nsILdapService::checkConfig (in AUTF8String aURL, in AUTF8String baseDN, in bool searchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted)
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Check if configuration is ok : this is a synchronous operation

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aURL</td>
<td align="left"><p>the server url</p></td>
</tr>
<tr class="even">
<td align="left">baseDN</td>
<td align="left"><p>the root of searches</p></td>
</tr>
<tr class="odd">
<td align="left">searchIsAuth</td>
<td align="left"><p>do we use login/password authentication, or is it anonymous ?</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserDN</td>
<td align="left"><p>user dn if searchIsAuth is true</p></td>
</tr>
<tr class="odd">
<td align="left">aPassword</td>
<td align="left"><p>user password if searchIsAuth is true (crypted by nsIWebserverCrypto)</p></td>
</tr>
<tr class="even">
<td align="left">aPasswordEncrypted</td>
<td align="left"><p>is the password encrypted ?</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a configuration error code (see constants above)

nsILdapCheckPollListener nsILdapService::createCheckListener ()
---------------------------------------------------------------

Create an ldaplistener object for polling for checking.

**Returns:.**

a nsILdapCheckPollListener object

nsILdapSearchPollListener nsILdapService::createSearchListener ()
-----------------------------------------------------------------

Create an ldaplistener object for polling for searching.

**Returns:.**

a nsILdapSearchPollListener object

void nsILdapService::asyncCheckConfig (in AUTF8String aURL, in AUTF8String baseDN, in bool searchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted, in nsILdapCheckListener aListener)
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Check if configuration is ok : this is a synchronous operation

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aURL</td>
<td align="left"><p>the server url</p></td>
</tr>
<tr class="even">
<td align="left">baseDN</td>
<td align="left"><p>the root of searches</p></td>
</tr>
<tr class="odd">
<td align="left">searchIsAuth</td>
<td align="left"><p>do we use login/password authentication, or is it anonymous ?</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserDN</td>
<td align="left"><p>user dn if searchIsAuth is true</p></td>
</tr>
<tr class="odd">
<td align="left">aPassword</td>
<td align="left"><p>user password if searchIsAuth is true (crypted by nsIWebserverCrypto)</p></td>
</tr>
<tr class="even">
<td align="left">aPasswordEncrypted</td>
<td align="left"><p>is the password encrypted ?</p></td>
</tr>
<tr class="odd">
<td align="left">aListener</td>
<td align="left"><p>listener to detect state changes</p></td>
</tr>
</tbody>
</table>

void nsILdapService::search (in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in PRUint32 aAttrsLength, \[array, size\_is(aAttrsLength)\] in string aAttrs, in PRUint32 aOffsetIndex, in PRUint32 aCountLimit, out PRUint32 aTotalCount, out PRUint32 aResCount, \[array, size\_is(aResCount), retval\] out nsIPropertyBag2 aRes)
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Search users in ldap base

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aContext</td>
<td align="left"><p>ldap context : if null, use context attribute</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserBase</td>
<td align="left"><p>the base where to search. Example : cn=Users,dc=innes,dc=fr if searchBase is not defined in server config, cn=Users if search base is defined to dc=innes,dc=fr If search base is defined, this param can be setted to empty string</p></td>
</tr>
<tr class="odd">
<td align="left">aDnFilter</td>
<td align="left"><p>filter for the search. Example: (&amp;(objectClass=user)(mail=*innes.fr))</p></td>
</tr>
<tr class="even">
<td align="left">aAttrsLength</td>
<td align="left"><p>number of attrs</p></td>
</tr>
<tr class="odd">
<td align="left">aAttrs</td>
<td align="left"><p>the attributes to return additionnaly to dn. Example : [&quot;mail&quot;, &quot;sAMAccountName&quot;]</p></td>
</tr>
<tr class="even">
<td align="left">aOffsetIndex</td>
<td align="left"><p>the offset of the first entry to return</p></td>
</tr>
<tr class="odd">
<td align="left">aCountLimit</td>
<td align="left"><p>the maximum number of entries to return</p></td>
</tr>
<tr class="even">
<td align="left">aTotalCount</td>
<td align="left"><p>the total number of entries available for that search</p></td>
</tr>
<tr class="odd">
<td align="left">aResCount</td>
<td align="left"><p>number of bags</p></td>
</tr>
<tr class="even">
<td align="left">aRes</td>
<td align="left"><p>property bags that contain entries from LDAP response</p></td>
</tr>
</tbody>
</table>

void nsILdapService::asyncSearch (in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in PRUint32 aAttrsLength, \[array, size\_is(aAttrsLength)\] in string aAttrs, in PRUint32 aOffsetIndex, in PRUint32 aCountLimit, in nsILdapSearchListener aListener)
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Search users in ldap base

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aContext</td>
<td align="left"><p>ldap context : if null, use context attribute</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserBase</td>
<td align="left"><p>the base where to search. Example : cn=Users,dc=innes,dc=fr if searchBase is not defined in server config, cn=Users if search base is defined to dc=innes,dc=fr If search base is defined, this param can be setted to empty string</p></td>
</tr>
<tr class="odd">
<td align="left">aDnFilter</td>
<td align="left"><p>filter for the search. Example: (&amp;(objectClass=user)(mail=*innes.fr))</p></td>
</tr>
<tr class="even">
<td align="left">aAttrsLength</td>
<td align="left"><p>number of aAttrs</p></td>
</tr>
<tr class="odd">
<td align="left">aAttrs</td>
<td align="left"><p>the attributes to return additionnaly to dn. Example : [&quot;mail&quot;, &quot;sAMAccountName&quot;]</p></td>
</tr>
<tr class="even">
<td align="left">aOffsetIndex</td>
<td align="left"><p>the offset of the first entry to return</p></td>
</tr>
<tr class="odd">
<td align="left">aCountLimit</td>
<td align="left"><p>the maximum number of entries to return</p></td>
</tr>
<tr class="even">
<td align="left">aListener</td>
<td align="left"><p>listener to get results</p></td>
</tr>
</tbody>
</table>

nsILdapCheckListener interface Reference
========================================

-   void onStateChanged ( in PRUint16 aResult)

Detailed Description
--------------------

The nsILdapCheckListener interface provides a callback for asyncCheckConfig operation.

void nsILdapCheckListener::onStateChanged (in PRUint16 aResult)
---------------------------------------------------------------

Callback which occurs for each change of status for asyncCheckConfig operation

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aResult</td>
<td align="left"><p>status of the asyncCheckConfig operation</p></td>
</tr>
</tbody>
</table>

nsILdapCheckPollListener interface Reference
============================================

Public Attributes
-----------------

-   readonly attribute PRUint16 state

Detailed Description
--------------------

The nsILdapCheckPollListener interface provides an attribute for polling when calling asyncCheckConfig

Member Data Documentation
-------------------------

### readonly attribute PRUint16 nsILdapCheckPollListener::state

polling for asyncCheckConfig operation

nsILdapSearchListener interface Reference
=========================================

-   void onSearchResult ( in PRUint16 aResult, in PRUint32 aResCount, in nsIPropertyBag2 aRes)

Detailed Description
--------------------

The nsILdapSearchListener interface provides a callback for asyncSearch operation

void nsILdapSearchListener::onSearchResult (in PRUint16 aResult, in PRUint32 aResCount, \[array, size\_is(aResCount)\] in nsIPropertyBag2 aRes)
-----------------------------------------------------------------------------------------------------------------------------------------------

Callback which occurs for each change of status for asyncSearch operation

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aResult</td>
<td align="left"><p>status of the asyncSearch operation</p></td>
</tr>
<tr class="even">
<td align="left">aResCount</td>
<td align="left"><p>number of entries</p></td>
</tr>
<tr class="odd">
<td align="left">aRes</td>
<td align="left"><p>property bags which contain entries of LDAP response (fullfilled if status is STATE_SUCCESS)</p></td>
</tr>
</tbody>
</table>

nsILdapSearchPollListener interface Reference
=============================================

-   void getResults ( out PRUint16 aStatus, out PRUint32 aResCount, out nsIPropertyBag2 aRes)

Detailed Description
--------------------

The nsILdapSearchPollListener interface provides for polling for asyncSearch operation

void nsILdapSearchPollListener::getResults (out PRUint16 aStatus, out PRUint32 aResCount, \[array, size\_is(aResCount), retval\] out nsIPropertyBag2 aRes)
----------------------------------------------------------------------------------------------------------------------------------------------------------

polling for asyncSearch operation

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aStatus</td>
<td align="left"><p>status of the asyncSearch operation (see async state constants in nsILdapService)</p></td>
</tr>
<tr class="even">
<td align="left">aResCount</td>
<td align="left"><p>number of entries</p></td>
</tr>
<tr class="odd">
<td align="left">aRes</td>
<td align="left"><p>property bags which contain entries of LDAP response (fullfilled if status is STATE_SUCCESS)</p></td>
</tr>
</tbody>
</table>

nsILdapContext interface Reference
==================================

Public Attributes
-----------------

-   attribute AUTF8String serverUrl

<!-- -->

-   attribute AUTF8String baseDN

<!-- -->

-   attribute bool searchIsAuth

<!-- -->

-   attribute AUTF8String searchUserDN

<!-- -->

-   attribute AUTF8String password

<!-- -->

-   attribute bool passwordEncrypted

This interface allows to instanciate a JS context object for ldapService methods.

Detailed Description
--------------------

Interface XPCOM nsILdapContext

Member Data Documentation
-------------------------

### attribute AUTF8String nsILdapContext::serverUrl

the Ldap server url

### attribute AUTF8String nsILdapContext::baseDN

the root of searchs

### attribute bool nsILdapContext::searchIsAuth

do we use login/password authentication, or is it anonymous ?

### attribute AUTF8String nsILdapContext::searchUserDN

user dn if searchIsAuth is true

### attribute AUTF8String nsILdapContext::password

password dn if searchIsAuth is true

### attribute bool nsILdapContext::passwordEncrypted

is the password encrypted ?
