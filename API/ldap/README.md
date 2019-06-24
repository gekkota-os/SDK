nsILdapService interface Reference
==================================

Public Attributes
-----------------

-   const unsigned short CONFIG\_OK

<!-- -->

-   const unsigned short CONFIG\_ERROR\_SERVER\_UNAVAILABLE

<!-- -->

-   const unsigned short CONFIG\_ERROR\_SERVER\_AUTH\_FAILED

<!-- -->

-   const unsigned short CONFIG\_ERROR\_SEARCH\_BASE\_DN

<!-- -->

-   const unsigned short STATE\_NONE

<!-- -->

-   const unsigned short STATE\_STARTED

<!-- -->

-   const unsigned short STATE\_LDAP\_URL\_ERROR

<!-- -->

-   const unsigned short STATE\_LDAP\_INIT\_ERROR

<!-- -->

-   const unsigned short STATE\_LDAP\_INITIALIZED

<!-- -->

-   const unsigned short STATE\_LDAP\_PASSWORD\_ERROR

<!-- -->

-   const unsigned short STATE\_AUTHENTICATION\_FAILED

<!-- -->

-   const unsigned short STATE\_SERVER\_UNAVAILABLE

<!-- -->

-   const unsigned short STATE\_AUTHENTICATION\_SUCCEEDED

<!-- -->

-   const unsigned short STATE\_SEARCH\_FAILED

<!-- -->

-   const unsigned short STATE\_SUCCESS\_NO\_RESULTS

<!-- -->

-   const unsigned short STATE\_SUCCESS

<!-- -->

-   const unsigned short STATE\_INTERNAL\_ERROR

<!-- -->

-   readonly attribute nsILdapContext context

-   bool checkAuthentication ( in nsILdapContext aContext, in AUTF8String aUserDN, in AUTF8String aPassword)

<!-- -->

-   unsigned short checkConfig ( in AUTF8String aURL, in AUTF8String aBaseDN, in bool aSearchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted)

<!-- -->

-   nsILdapCheckPollListener createCheckListener ( )

<!-- -->

-   nsILdapSearchPollListener createSearchListener ( )

<!-- -->

-   void asyncCheckConfig ( in AUTF8String aURL, in AUTF8String aBaseDN, in bool aSearchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted, in nsILdapCheckListener aListener)

<!-- -->

-   void search ( in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in unsigned long aAttrsLength, in string aAttrs, in unsigned long aOffsetIndex, in unsigned long aCountLimit, out unsigned long aTotalCount, out unsigned long aResCount, out nsIPropertyBag2 aRes)

<!-- -->

-   void asyncSearch ( in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in unsigned long aAttrsLength, in string aAttrs, in unsigned long aOffsetIndex, in unsigned long aCountLimit, in nsILdapSearchListener aListener)

Detailed Description
--------------------

For using the nsILdapService interface, we call properties of Javascript ldapService object. Complete HTML example using this API [here.](example1.html)

**Build note**: You need to execute the **build.cmd** file to generate the boostrap app. Otherwise there will be a mismatch between the html file name and the one the manifest tries to launch. Find more information in *SDK-G4/bootstrap App/* documentation.

Member Data Documentation
-------------------------

### const unsigned short nsILdapService::CONFIG\_OK

Configuration ok.

### const unsigned short nsILdapService::CONFIG\_ERROR\_SERVER\_UNAVAILABLE

Ldap server is not reachable (URL may be invalid).

### const unsigned short nsILdapService::CONFIG\_ERROR\_SERVER\_AUTH\_FAILED

Authentication failed.

### const unsigned short nsILdapService::CONFIG\_ERROR\_SEARCH\_BASE\_DN

Search base DN invalid.

### const unsigned short nsILdapService::STATE\_NONE

State none (operation is not started).

### const unsigned short nsILdapService::STATE\_STARTED

State for operation started.

### const unsigned short nsILdapService::STATE\_LDAP\_URL\_ERROR

State for invalid LDAP URL.

### const unsigned short nsILdapService::STATE\_LDAP\_INIT\_ERROR

State for connection failure.

### const unsigned short nsILdapService::STATE\_LDAP\_INITIALIZED

State for connection successfully initialized.

### const unsigned short nsILdapService::STATE\_LDAP\_PASSWORD\_ERROR

State for invalid password (bad encryption).

### const unsigned short nsILdapService::STATE\_AUTHENTICATION\_FAILED

State for wrong login or password.

### const unsigned short nsILdapService::STATE\_SERVER\_UNAVAILABLE

State for LDAP server unreachable.

### const unsigned short nsILdapService::STATE\_AUTHENTICATION\_SUCCEEDED

State for succeeded authentication.

### const unsigned short nsILdapService::STATE\_SEARCH\_FAILED

State for searching failure.

### const unsigned short nsILdapService::STATE\_SUCCESS\_NO\_RESULTS

State for succeeded request with empty result.

### const unsigned short nsILdapService::STATE\_SUCCESS

State for succeeded request with some results.

### const unsigned short nsILdapService::STATE\_INTERNAL\_ERROR

State for internal error (not LDAP problem).

### readonly attribute nsILdapContext nsILdapService::context

Default LDAP context.

bool nsILdapService::checkAuthentication (in nsILdapContext aContext, in AUTF8String aUserDN, in AUTF8String aPassword)
-----------------------------------------------------------------------------------------------------------------------

Check LDAP server authentication.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aContext</td>
<td align="left"><p>Context that defines LDAP information. If null, get preferences information</p></td>
</tr>
<tr class="even">
<td align="left">aUserDN</td>
<td align="left"><p>Ldap User</p></td>
</tr>
<tr class="odd">
<td align="left">aPassword</td>
<td align="left"><p>Ldap password</p></td>
</tr>
</tbody>
</table>

**Returns:.**

True is authentication is ok.

unsigned short nsILdapService::checkConfig (in AUTF8String aURL, in AUTF8String aBaseDN, in bool aSearchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted)
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Checks if configuration is ok: this is a synchronous operation.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aURL</td>
<td align="left"><p>The server URL</p></td>
</tr>
<tr class="even">
<td align="left">aBaseDN</td>
<td align="left"><p>The root of searches</p></td>
</tr>
<tr class="odd">
<td align="left">aSearchIsAuth</td>
<td align="left"><p>States whether we use login/password authentication or it is anonymous</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserDN</td>
<td align="left"><p>User dn if aSearchIsAuth is true</p></td>
</tr>
<tr class="odd">
<td align="left">aPassword</td>
<td align="left"><p>User password if aSearchIsAuth is true (crypted by nsIWebserverCrypto)</p></td>
</tr>
<tr class="even">
<td align="left">aPasswordEncrypted</td>
<td align="left"><p>States whether the password is encrypted or not</p></td>
</tr>
</tbody>
</table>

**Returns:.**

A configuration error code (see constants above).

nsILdapCheckPollListener nsILdapService::createCheckListener ()
---------------------------------------------------------------

Creates a LDAP listener object for polling for checking.

**Returns:.**

A nsILdapCheckPollListener object.

nsILdapSearchPollListener nsILdapService::createSearchListener ()
-----------------------------------------------------------------

Creates a LDAP listener object for polling for searching.

**Returns:.**

A nsILdapSearchPollListener object.

void nsILdapService::asyncCheckConfig (in AUTF8String aURL, in AUTF8String aBaseDN, in bool aSearchIsAuth, in AUTF8String aSearchUserDN, in AUTF8String aPassword, in bool aPasswordEncrypted, in nsILdapCheckListener aListener)
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Checks if configuration is ok: this is a synchronous operation.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aURL</td>
<td align="left"><p>The server URL</p></td>
</tr>
<tr class="even">
<td align="left">aBaseDN</td>
<td align="left"><p>The root of searches</p></td>
</tr>
<tr class="odd">
<td align="left">aSearchIsAuth</td>
<td align="left"><p>States whether we use login/password authentication or it is anonymous</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserDN</td>
<td align="left"><p>User dn if aSearchIsAuth is true</p></td>
</tr>
<tr class="odd">
<td align="left">aPassword</td>
<td align="left"><p>User password if aSearchIsAuth is true (crypted by nsIWebserverCrypto)</p></td>
</tr>
<tr class="even">
<td align="left">aPasswordEncrypted</td>
<td align="left"><p>States whether the password is encrypted or not</p></td>
</tr>
<tr class="odd">
<td align="left">aListener</td>
<td align="left"><p>Listener to detect state changes</p></td>
</tr>
</tbody>
</table>

void nsILdapService::search (in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in unsigned long aAttrsLength, \[array, size\_is(aAttrsLength)\] in string aAttrs, in unsigned long aOffsetIndex, in unsigned long aCountLimit, out unsigned long aTotalCount, out unsigned long aResCount, \[array, size\_is(aResCount), retval\] out nsIPropertyBag2 aRes)
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Search users in LDAP base.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aContext</td>
<td align="left"><p>Ldap context: if null, use context attribute</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserBase</td>
<td align="left"><p>The base where to search. Example: cn=Users,dc=innes,dc=fr if searchBase is not defined in server config, cn=Users if search base is defined to dc=innes,dc=fr If search base is defined, this param can be set to empty string</p></td>
</tr>
<tr class="odd">
<td align="left">aDnFilter</td>
<td align="left"><p>Filter for the search. Example: (&amp;(objectClass=user)(mail=*innes.fr))</p></td>
</tr>
<tr class="even">
<td align="left">aAttrsLength</td>
<td align="left"><p>Number of attrs</p></td>
</tr>
<tr class="odd">
<td align="left">aAttrs</td>
<td align="left"><p>The attributes to return additionnaly to dn. Example: [&quot;mail&quot;, &quot;sAMAccountName&quot;]</p></td>
</tr>
<tr class="even">
<td align="left">aOffsetIndex</td>
<td align="left"><p>The offset of the first entry to return</p></td>
</tr>
<tr class="odd">
<td align="left">aCountLimit</td>
<td align="left"><p>The maximum number of entries to return</p></td>
</tr>
<tr class="even">
<td align="left">aTotalCount</td>
<td align="left"><p>The total number of entries available for that search</p></td>
</tr>
<tr class="odd">
<td align="left">aResCount</td>
<td align="left"><p>Number of bags</p></td>
</tr>
<tr class="even">
<td align="left">aRes</td>
<td align="left"><p>Property bags that contain entries from LDAP response</p></td>
</tr>
</tbody>
</table>

void nsILdapService::asyncSearch (in nsILdapContext aContext, in AUTF8String aSearchUserBase, in AUTF8String aDnFilter, in unsigned long aAttrsLength, \[array, size\_is(aAttrsLength)\] in string aAttrs, in unsigned long aOffsetIndex, in unsigned long aCountLimit, in nsILdapSearchListener aListener)
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Search users in LDAP base.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aContext</td>
<td align="left"><p>Ldap context: if null, use context attribute</p></td>
</tr>
<tr class="even">
<td align="left">aSearchUserBase</td>
<td align="left"><p>The base where to search. Example: cn=Users,dc=innes,dc=fr if searchBase is not defined in server config, cn=Users if search base is defined to dc=innes,dc=fr If search base is defined, this param can be setted to empty string</p></td>
</tr>
<tr class="odd">
<td align="left">aDnFilter</td>
<td align="left"><p>Filter for the search. Example: (&amp;(objectClass=user)(mail=*innes.fr))</p></td>
</tr>
<tr class="even">
<td align="left">aAttrsLength</td>
<td align="left"><p>Number of aAttrs</p></td>
</tr>
<tr class="odd">
<td align="left">aAttrs</td>
<td align="left"><p>The attributes to return additionnaly to dn. Example: [&quot;mail&quot;, &quot;sAMAccountName&quot;]</p></td>
</tr>
<tr class="even">
<td align="left">aOffsetIndex</td>
<td align="left"><p>The offset of the first entry to return</p></td>
</tr>
<tr class="odd">
<td align="left">aCountLimit</td>
<td align="left"><p>The maximum number of entries to return</p></td>
</tr>
<tr class="even">
<td align="left">aListener</td>
<td align="left"><p>Listener to get results</p></td>
</tr>
</tbody>
</table>

nsILdapCheckListener interface Reference
========================================

-   void onStateChanged ( in unsigned short aResult)

Detailed Description
--------------------

The nsILdapCheckListener interface provides a callback for asyncCheckConfig operation.

void nsILdapCheckListener::onStateChanged (in unsigned short aResult)
---------------------------------------------------------------------

Callback which occurs for each change of status for asyncCheckConfig operation.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aResult</td>
<td align="left"><p>Status of the asyncCheckConfig operation</p></td>
</tr>
</tbody>
</table>

nsILdapCheckPollListener interface Reference
============================================

Public Attributes
-----------------

-   readonly attribute unsigned short state

Detailed Description
--------------------

The nsILdapCheckPollListener interface provides an attribute for polling when calling asyncCheckConfig.

Member Data Documentation
-------------------------

### readonly attribute unsigned short nsILdapCheckPollListener::state

polling for asyncCheckConfig operation.

nsILdapSearchListener interface Reference
=========================================

-   void onSearchResult ( in unsigned short aResult, in unsigned long aResCount, in nsIPropertyBag2 aRes)

Detailed Description
--------------------

The nsILdapSearchListener interface provides a callback for asyncSearch operation.

void nsILdapSearchListener::onSearchResult (in unsigned short aResult, in unsigned long aResCount, \[array, size\_is(aResCount)\] in nsIPropertyBag2 aRes)
----------------------------------------------------------------------------------------------------------------------------------------------------------

Callback which occurs for each change of status for asyncSearch operation.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aResult</td>
<td align="left"><p>Status of the asyncSearch operation</p></td>
</tr>
<tr class="even">
<td align="left">aResCount</td>
<td align="left"><p>Number of entries</p></td>
</tr>
<tr class="odd">
<td align="left">aRes</td>
<td align="left"><p>Property bags which contain entries of LDAP response (fullfilled if status is STATE_SUCCESS)</p></td>
</tr>
</tbody>
</table>

nsILdapSearchPollListener interface Reference
=============================================

-   void getResults ( out unsigned short aStatus, out unsigned long aResCount, out nsIPropertyBag2 aRes)

Detailed Description
--------------------

The nsILdapSearchPollListener interface provides for polling for asyncSearch operation.

void nsILdapSearchPollListener::getResults (out unsigned short aStatus, out unsigned long aResCount, \[array, size\_is(aResCount), retval\] out nsIPropertyBag2 aRes)
---------------------------------------------------------------------------------------------------------------------------------------------------------------------

Polling for asyncSearch operation.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aStatus</td>
<td align="left"><p>Status of the asyncSearch operation (see async state constants in nsILdapService)</p></td>
</tr>
<tr class="even">
<td align="left">aResCount</td>
<td align="left"><p>Number of entries</p></td>
</tr>
<tr class="odd">
<td align="left">aRes</td>
<td align="left"><p>Property bags which contain entries of LDAP response (fullfilled if status is STATE_SUCCESS)</p></td>
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

The Ldap server url.

### attribute AUTF8String nsILdapContext::baseDN

The root of searchs.

### attribute bool nsILdapContext::searchIsAuth

True if we use login/password authentication, otherwise it is anonymous.

### attribute AUTF8String nsILdapContext::searchUserDN

User dn if searchIsAuth is true.

### attribute AUTF8String nsILdapContext::password

Password dn if searchIsAuth is true.

### attribute bool nsILdapContext::passwordEncrypted

True if the password is encrypted.
