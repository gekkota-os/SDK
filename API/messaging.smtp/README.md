nsIMessaging interface Reference
================================

Public Attributes
-----------------

-   readonly attribute nsIEmailManager email

Detailed Description
--------------------

The nsIMessaging interface is the point of entry for the email manager. Complete HTML example using this API [here.](example1.html)

**Build note**: You need to execute the **build.cmd** file to generate the boostrap app. Otherwise there will be a mismatch between the html file name and the one the manifest tries to launch. Find more information in *SDK-G4/bootstrap App/* documentation.

Member Data Documentation
-------------------------

### readonly attribute nsIEmailManager nsIMessaging::email

Email read-only.

nsIEmailManager interface Reference
===================================

Public Attributes
-----------------

-   readonly attribute jsval accounts

-   nsIEmailRequest send ( in nsIEmailRecipientsList aReciptientsList, in nsIEmailMessage aMessage)

<!-- -->

-   void cancel ( )

Detailed Description
--------------------

The nsIEmailManager interface allows to send an email.

Member Data Documentation
-------------------------

### readonly attribute jsval nsIEmailManager::accounts

Array of nsISmtpAccount.

nsIEmailRequest nsIEmailManager::send (in nsIEmailRecipientsList aReciptientsList, in nsIEmailMessage aMessage)
---------------------------------------------------------------------------------------------------------------

Send message.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aReciptientsList</td>
<td align="left"><p>List of recipients</p></td>
</tr>
<tr class="even">
<td align="left">aMessage</td>
<td align="left"><p>Message to send</p></td>
</tr>
</tbody>
</table>

void nsIEmailManager::cancel ()
-------------------------------

Cancel.

nsIEmailMessage interface Reference
===================================

Public Attributes
-----------------

-   attribute DOMString subject

<!-- -->

-   attribute unsigned long priority

<!-- -->

-   attribute nsISmtpAccount smtpAccount

<!-- -->

-   attribute nsIEmailBody body

<!-- -->

-   attribute boolean hasMDN

<!-- -->

-   readonly attribute AUTF8String data

-   void addAttachment ( in nsIEmailAttachment aAttachment)

<!-- -->

-   void removeAttachment ( in nsIEmailAttachment aAttachment)

<!-- -->

-   void getAttachments ( out unsigned long aLength, out nsIEmailAttachment aAttachments)

<!-- -->

-   void buildData ( in nsIEmailRecipientsList aRecipientsList)

Detailed Description
--------------------

The nsIEmailMessage interface allows to create an email message.

Member Data Documentation
-------------------------

### attribute DOMString nsIEmailMessage::subject

Subject.

### attribute unsigned long nsIEmailMessage::priority

Priority.

### attribute nsISmtpAccount nsIEmailMessage::smtpAccount

SMTP account.

### attribute nsIEmailBody nsIEmailMessage::body

Body.

### attribute boolean nsIEmailMessage::hasMDN

Has Message Disposition Notifications.

### readonly attribute AUTF8String nsIEmailMessage::data

Data read-only.

void nsIEmailMessage::addAttachment (in nsIEmailAttachment aAttachment)
-----------------------------------------------------------------------

Add email attachment.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAttachment</td>
<td align="left"><p>The email attachment to add</p></td>
</tr>
</tbody>
</table>

void nsIEmailMessage::removeAttachment (in nsIEmailAttachment aAttachment)
--------------------------------------------------------------------------

Remove email attachment.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAttachment</td>
<td align="left"><p>The email attachment to remove</p></td>
</tr>
</tbody>
</table>

void nsIEmailMessage::getAttachments (\[optional\] out unsigned long aLength, \[array, size\_is(aLength), retval, optional\] out nsIEmailAttachment aAttachments)
-----------------------------------------------------------------------------------------------------------------------------------------------------------------

Get email attachment.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>Array length</p></td>
</tr>
<tr class="even">
<td align="left">aAttachments</td>
<td align="left"><p>Array of attachments</p></td>
</tr>
</tbody>
</table>

void nsIEmailMessage::buildData (in nsIEmailRecipientsList aRecipientsList)
---------------------------------------------------------------------------

Build data from list of recipients.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientsList</td>
<td align="left"><p>The list of recipients</p></td>
</tr>
</tbody>
</table>

nsIEmailAttachment interface Reference
======================================

Public Attributes
-----------------

-   readonly attribute DOMString mimeType

<!-- -->

-   readonly attribute DOMString filePath

<!-- -->

-   readonly attribute DOMString fileName

-   void init ( in AString aFileUri)

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIEmailAttachment::mimeType

Mime type.

### readonly attribute DOMString nsIEmailAttachment::filePath

File path.

### readonly attribute DOMString nsIEmailAttachment::fileName

File name.

void nsIEmailAttachment::init (in AString aFileUri)
---------------------------------------------------

Initialize from an URI.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aFileUri</td>
<td align="left"><p>URI of the attachment</p></td>
</tr>
</tbody>
</table>

nsIInlineEmailAttachment interface Reference
============================================

Public Attributes
-----------------

-   readonly attribute DOMString inlineID

-   void initInline ( in AString aFileUri, in AString aId)

Detailed Description
--------------------

The nsIInlineEmailAttachment interface allows to manage an inline email attachment.

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIInlineEmailAttachment::inlineID

Inline ID

void nsIInlineEmailAttachment::initInline (in AString aFileUri, in AString aId)
-------------------------------------------------------------------------------

Initialize inline.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aFileUri</td>
<td align="left"><p>File URI</p></td>
</tr>
<tr class="even">
<td align="left">aId</td>
<td align="left"><p>Identifier</p></td>
</tr>
</tbody>
</table>

nsIEmailBodyType interface Reference
====================================

Public Attributes
-----------------

-   const nsEmailBodyTypeValue PLAIN\_TEXT

<!-- -->

-   const nsEmailBodyTypeValue HTML

Detailed Description
--------------------

The nsIEmailBodyType interface is an enumeration for email body type.

Member Data Documentation
-------------------------

### const nsEmailBodyTypeValue nsIEmailBodyType::PLAIN\_TEXT

Plain text.

### const nsEmailBodyTypeValue nsIEmailBodyType::HTML

Html.

nsIEmailBody interface Reference
================================

Public Attributes
-----------------

-   readonly attribute DOMString body

<!-- -->

-   readonly attribute unsigned long bodyType

-   void addInlineAttachment ( in DOMString aID, in DOMString aFileUri)

<!-- -->

-   void removeInlineAttachment ( in DOMString aID)

<!-- -->

-   void getArrayOfInlineAttachments ( out unsigned long aLength, out nsIInlineEmailAttachment aAttachments)

Detailed Description
--------------------

The nsIEmailBody interface manages the email body.

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIEmailBody::body

Email body.

### readonly attribute unsigned long nsIEmailBody::bodyType

Type of body.

void nsIEmailBody::addInlineAttachment (in DOMString aID, in DOMString aFileUri)
--------------------------------------------------------------------------------

Add an inline attachment.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aID</td>
<td align="left"><p>String id, must be begin with &quot;cid:&quot; or &quot;mid:&quot;</p></td>
</tr>
<tr class="even">
<td align="left">aFileUri</td>
<td align="left"><p>File URI</p></td>
</tr>
</tbody>
</table>

void nsIEmailBody::removeInlineAttachment (in DOMString aID)
------------------------------------------------------------

Remove an inline attachment.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aID</td>
<td align="left"><p>String id, must be begin with &quot;cid:&quot; or &quot;mid:&quot;</p></td>
</tr>
</tbody>
</table>

void nsIEmailBody::getArrayOfInlineAttachments (\[optional\] out unsigned long aLength, \[array, size\_is(aLength), retval, optional\] out nsIInlineEmailAttachment aAttachments)
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Get array of inline attachments.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>Array length</p></td>
</tr>
<tr class="even">
<td align="left">aAttachments</td>
<td align="left"><p>Array of attachments</p></td>
</tr>
</tbody>
</table>

nsIEmailEventListener interface Reference
=========================================

-   void handleEvent ( in unsigned long aIntVal)

Detailed Description
--------------------

The nsIEmailEventListener interface provide a listener for email event.

void nsIEmailEventListener::handleEvent (in unsigned long aIntVal)
------------------------------------------------------------------

Handle event.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aIntVal</td>
<td align="left"><p>integer value</p></td>
</tr>
</tbody>
</table>

nsIEmailEventType interface Reference
=====================================

Public Attributes
-----------------

-   const nsEmailEventTypeValue SUCCESS

<!-- -->

-   const nsEmailEventTypeValue ERROR

<!-- -->

-   const nsEmailEventTypeValue PROGRESS

Member Data Documentation
-------------------------

### const nsEmailEventTypeValue nsIEmailEventType::SUCCESS

Success.

### const nsEmailEventTypeValue nsIEmailEventType::ERROR

Error.

### const nsEmailEventTypeValue nsIEmailEventType::PROGRESS

In progress.

nsIEmailPriorityType interface Reference
========================================

Public Attributes
-----------------

-   const nsPriorityTypeValue NORMAL

<!-- -->

-   const nsPriorityTypeValue HIGH

<!-- -->

-   const nsPriorityTypeValue LOW

Member Data Documentation
-------------------------

### const nsPriorityTypeValue nsIEmailPriorityType::NORMAL

Priority normal.

### const nsPriorityTypeValue nsIEmailPriorityType::HIGH

Priority high.

### const nsPriorityTypeValue nsIEmailPriorityType::LOW

Priority low.

nsIEmailRecipientsList interface Reference
==========================================

-   void addRecipient ( in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)

<!-- -->

-   void removeRecipient ( in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)

<!-- -->

-   void getAllRecipients ( out unsigned long aLength, out wstring aRecipients)

<!-- -->

-   void getRecipients ( in nsRecipientsTypeValue aRecipientType, out unsigned long aLength, out wstring aRecipients)

Detailed Description
--------------------

The nsIEmailRecipientsList interface allows to manage a list of recipients.

void nsIEmailRecipientsList::addRecipient (in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)
------------------------------------------------------------------------------------------------------------

Add a recipient.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientType</td>
<td align="left"><p>Type of recipient</p></td>
</tr>
<tr class="even">
<td align="left">aRecipient</td>
<td align="left"><p>Recipient to add</p></td>
</tr>
</tbody>
</table>

void nsIEmailRecipientsList::removeRecipient (in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)
---------------------------------------------------------------------------------------------------------------

Remove a recipient.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientType</td>
<td align="left"><p>Type of recipient</p></td>
</tr>
<tr class="even">
<td align="left">aRecipient</td>
<td align="left"><p>Recipient to remove</p></td>
</tr>
</tbody>
</table>

void nsIEmailRecipientsList::getAllRecipients (\[optional\] out unsigned long aLength, \[array, size\_is(aLength), retval, optional\] out wstring aRecipients)
--------------------------------------------------------------------------------------------------------------------------------------------------------------

Get all recipients.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>Array length</p></td>
</tr>
<tr class="even">
<td align="left">aRecipients</td>
<td align="left"><p>Array of recipients</p></td>
</tr>
</tbody>
</table>

void nsIEmailRecipientsList::getRecipients (in nsRecipientsTypeValue aRecipientType, \[optional\] out unsigned long aLength, \[array, size\_is(aLength), retval, optional\] out wstring aRecipients)
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Get some recipients.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientType</td>
<td align="left"><p>Type of recipient</p></td>
</tr>
<tr class="even">
<td align="left">aLength</td>
<td align="left"><p>Array length</p></td>
</tr>
<tr class="odd">
<td align="left">aRecipients</td>
<td align="left"><p>Array of recipients</p></td>
</tr>
</tbody>
</table>

nsIEmailRecipientsType interface Reference
==========================================

Public Attributes
-----------------

-   const nsRecipientsTypeValue TO

<!-- -->

-   const nsRecipientsTypeValue CC

<!-- -->

-   const nsRecipientsTypeValue BCC

Member Data Documentation
-------------------------

### const nsRecipientsTypeValue nsIEmailRecipientsType::TO

Value type is TO.

### const nsRecipientsTypeValue nsIEmailRecipientsType::CC

Value type is CC or Carbon Copy.

### const nsRecipientsTypeValue nsIEmailRecipientsType::BCC

Value type is BBC or Blind Carbon Copy.

nsIEmailRequest interface Reference
===================================

Public Attributes
-----------------

-   readonly attribute DOMString readyState

-   void addEmailEventListener ( in nsEmailEventTypeValue aEventType, in nsIEmailEventListener aListener)

<!-- -->

-   void handleSuccess ( )

<!-- -->

-   void handleError ( )

<!-- -->

-   void handleProgress ( in unsigned long aProgressPercentage)

Detailed Description
--------------------

The nsIEmailRequest interface is used to manage request for email event.

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIEmailRequest::readyState

Request ready state ("processing" or "done").

void nsIEmailRequest::addEmailEventListener (in nsEmailEventTypeValue aEventType, in nsIEmailEventListener aListener)
---------------------------------------------------------------------------------------------------------------------

Add email event listener.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aEventType</td>
<td align="left"><p>Type of event</p></td>
</tr>
<tr class="even">
<td align="left">aListener</td>
<td align="left"><p>Listener</p></td>
</tr>
</tbody>
</table>

void nsIEmailRequest::handleSuccess ()
--------------------------------------

Handle success.

void nsIEmailRequest::handleError ()
------------------------------------

Handle error.

void nsIEmailRequest::handleProgress (in unsigned long aProgressPercentage)
---------------------------------------------------------------------------

Retrieve handle progress.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aProgressPercentage</td>
<td align="left"><p>progress in percentage</p></td>
</tr>
</tbody>
</table>

nsISmtpAccount interface Reference
==================================

Public Attributes
-----------------

-   readonly attribute DOMString serverUri

<!-- -->

-   readonly attribute DOMString username

<!-- -->

-   readonly attribute DOMString password

<!-- -->

-   readonly attribute boolean authorizeUncrytptedAuth

<!-- -->

-   readonly attribute DOMString sender

<!-- -->

-   readonly attribute DOMString hostname

<!-- -->

-   readonly attribute boolean useSSL

<!-- -->

-   readonly attribute unsigned long port

Detailed Description
--------------------

The nsISmtpAccount interface provide information about STMP account.

Member Data Documentation
-------------------------

### readonly attribute DOMString nsISmtpAccount::serverUri

Server URI.

### readonly attribute DOMString nsISmtpAccount::username

Username.

### readonly attribute DOMString nsISmtpAccount::password

Password.

### readonly attribute boolean nsISmtpAccount::authorizeUncrytptedAuth

AuthorizeUncrytptedAuth.

### readonly attribute DOMString nsISmtpAccount::sender

Sender.

### readonly attribute DOMString nsISmtpAccount::hostname

Hostname.

### readonly attribute boolean nsISmtpAccount::useSSL

UseSSL.

### readonly attribute unsigned long nsISmtpAccount::port

Port.

nsMsgSocketType interface Reference
===================================

Public Attributes
-----------------

-   const nsMsgSocketTypeValue plain

<!-- -->

-   const nsMsgSocketTypeValue trySTARTTLS

<!-- -->

-   const nsMsgSocketTypeValue alwaysSTARTTLS

<!-- -->

-   const nsMsgSocketTypeValue SSL

Member Data Documentation
-------------------------

### const nsMsgSocketTypeValue nsMsgSocketType::plain

No SSL or STARTTLS.

### const nsMsgSocketTypeValue nsMsgSocketType::trySTARTTLS

Use TLS via STARTTLS, but only if server offers it.

Deprecated

This is vulnerable to MITM attacks.

### const nsMsgSocketTypeValue nsMsgSocketType::alwaysSTARTTLS

Insist on TLS via STARTTLS. Uses normal port.

### const nsMsgSocketTypeValue nsMsgSocketType::SSL

Connect via SSL. Needs special SSL port.

nsMsgAuthMethod interface Reference
===================================

Public Attributes
-----------------

-   const nsMsgAuthMethodValue none

<!-- -->

-   const nsMsgAuthMethodValue old

<!-- -->

-   const nsMsgAuthMethodValue passwordCleartext

<!-- -->

-   const nsMsgAuthMethodValue passwordEncrypted

<!-- -->

-   const nsMsgAuthMethodValue GSSAPI

<!-- -->

-   const nsMsgAuthMethodValue NTLM

<!-- -->

-   const nsMsgAuthMethodValue External

<!-- -->

-   const nsMsgAuthMethodValue secure

<!-- -->

-   const nsMsgAuthMethodValue anything

Detailed Description
--------------------

The nsMsgAuthMethod interface defines which authentication schemes we should try. Used by

**See also:.**

nsIMsgIncomingServer.authMethod and

nsISmtpServer.authMethod.

Member Data Documentation
-------------------------

### const nsMsgAuthMethodValue nsMsgAuthMethod::none

No login needed. E.g. IP-address-based.

### const nsMsgAuthMethodValue nsMsgAuthMethod::old

Do not use AUTH commands (e.g. AUTH=PLAIN), but the original login commands that the protocol specified (POP: "USER"/"PASS", IMAP: "login", not valid for SMTP).

### const nsMsgAuthMethodValue nsMsgAuthMethod::passwordCleartext

Password in the clear. AUTH=PLAIN/LOGIN or old-style login.

### const nsMsgAuthMethodValue nsMsgAuthMethod::passwordEncrypted

Hashed password. CRAM-MD5, DIGEST-MD5.

### const nsMsgAuthMethodValue nsMsgAuthMethod::GSSAPI

Kerberos / GSSAPI (Unix single-signon).

### const nsMsgAuthMethodValue nsMsgAuthMethod::NTLM

NTLM is a Windows single-singon scheme. Includes MSN / Passport.net, which is the same with a different name.

### const nsMsgAuthMethodValue nsMsgAuthMethod::External

Authentication External is certificat-based.

### const nsMsgAuthMethodValue nsMsgAuthMethod::secure

Encrypted password or Kerberos / GSSAPI or NTLM.

Deprecated

-   for migration only.

### const nsMsgAuthMethodValue nsMsgAuthMethod::anything

Let us pick any of the auth types supported by the server. Discouraged, because vulnerable to MITM attacks, even if server offers secure auth.
