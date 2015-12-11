nsIMessaging interface Reference
================================

Public Attributes
-----------------

-   readonly attribute nsIEmailManager email

Detailed Description
--------------------

The nsIMessaging interface is the point of entry for the email manager

Here is an example

    function XpfLogger(name){
     this._logger = log4Service.getLogger(name);
    }

    XpfLogger.prototype = {
     _logger: undefined,

     _getLocationInfoFromException: function(e, level){
      if(!e.stack){
       return null;
      }
      var lines = e.stack.split("\n");
      var reg = new RegExp("^(.*)@(.*):(.*)", "g");

      var line = lines[level];
      var file = line.replace(reg, "$2");
      var func = line.replace(reg, "$1");
      var line = line.replace(reg, "$3");
      return log4Service.createLocationInfo(file, func, line);
     },
     _getLocationInfo: function(){
      try{
       throw new Error();
      }catch(e){
       return this._getLocationInfoFromException(e, 2);
      }
     },

     isDebugEnabled: function(){
      return this._logger.isDebugEnabled();
     },
     debug: function(msg){
      if(this.isDebugEnabled()){
       this._logger.debug(msg, this._getLocationInfo());
      }
     },
     debugEx: function(e){
      if(this.isDebugEnabled()){
       this._logger.debug(e, this._getLocationInfoFromException(e, 0));
      }
     },
     isWarnEnabled: function(){
      return this._logger.isWarnEnabled();
     },
     warn: function(msg){
      if(this.isWarnEnabled()){
       this._logger.warn(msg, this._getLocationInfo());
      }
     },
     warnEx: function(e){
      if(this.isWarnEnabled()){
       this._logger.warn(e, this._getLocationInfoFromException(e, 0));
      }
     },
     isErrorEnabled: function(){
      return this._logger.isErrorEnabled();
     },
     error: function(msg){
      if(this.isErrorEnabled()){
       this._logger.error(msg, this._getLocationInfo());
      }
     },
     errorEx: function(e){
      if(this.isErrorEnabled()){
       this._logger.error(e, this._getLocationInfoFromException(e, 0));
      }
     },
     isFatalEnabled: function(){
      return this._logger.isFatalEnabled();
     },
     fatal: function(msg){
      if(this.isFatalEnabled()){
       this._logger.fatal(msg, this._getLocationInfo());
      }
     },
     fatalEx: function(e){
      if(this.isFatalEnabled()){
       this._logger.fatal(e, this._getLocationInfoFromException(e, 0));
      }
     }
    };

    //var logger = new XpfLogger("jstest");
    var logger = log4Service.getLogger("jstest");

    function log(val)
    { 
     logger.debugAS("###### JSTEST ###### " + val, null);
    }

    function browseObject(obj)
    {
     var strObj = String(obj);
     log("obj (" + strObj + ") =>");
     for (var i in obj) 
     {
      log("obj." + i +" = "+ obj[i]);
     }
    }

    function addTest(name)
    {
     var root = document.getElementById("root");
     
     // add new test button
     var newButton = document.createElement("button");
     newButton.innerHTML = name ;
     var start = "log('start " + name + "...'); ";
     var end = "log('end " + name + ".'); "; 
     newButton.setAttribute("OnClick", start + name + "(); " + end);
     root.appendChild(document.createElement("br"));
     root.appendChild(document.createElement("br"));
     root.appendChild(newButton);
     log("The button " + name  + " is added !");
    }
    /* ------------------------------- test functions -------------------------------------------------------  */
    function testEmailRecepients()
    {
     
     var recepList = new EmailRecipientsList();
     
     // avec retval + optional 
     recepList.addRecipient(EmailRecipientsType.TO, "michel.bonjour@test.fr");
     recepList.addRecipient(EmailRecipientsType.TO, "michel1@mailinator.com");
     recepList.addRecipient(EmailRecipientsType.CC, "michel2@mailinator.com");
     recepList.addRecipient(EmailRecipientsType.BCC, "michel3@mailinator.com");
     
     var recipients_to = recepList.getRecipients(EmailRecipientsType.TO);
     var recipients_cc = recepList.getRecipients(EmailRecipientsType.CC);
     var recipients_bcc = recepList.getRecipients(EmailRecipientsType.BCC);
     
     log("TO = " + recipients_to);
     log("CC = " + recipients_cc);
     log("BCC = " + recipients_bcc);
    }

    function testEmailAttachment()
    {
     var test = new EmailAttachment("marianne.jpg");
     var fp = test.filePath;
     var mime = test.mimeType;
     log("fp = " + fp);
     log("mime = " + mime);
    }

    function testEmailBody()
    { 
     var txt = "Salut,\r\nMIME définit des mécanismes pour l'envoi d'autre sortes d'informations dont des textes dans des langages autres que l'anglais utilisant des codages de caractères autres que l'ASCII et des données binaires comme des fichiers contenant des images, des sons, des films ou des programmes informatiques. MIME est également un composant fondamental des protocoles de communications comme HTTP, qui requièrent l'envoi de données dans le même contexte que l'envoi de courriels, même si ces données ne sont pas des courriels. L'intégration ou l'extraction des données au format MIME est généralement automatiquement effectuée par le client de messagerie ou par le serveur de messagerie électronique quand le courriel est envoyé ou reçu.\r\n\r\nBye!\r\n";
     var body = new EmailBody(txt, EmailBodyType.PLAIN_TEXT);
     browseObject(body);
    }

    function testSmtpAccount()
    {
     /* SmtpAccount (serverUri, username, password, authorizeUncrytptedAuth, sender) */
     var test = new SmtpAccount("smtp://smtp.gmail.com:465", "michel.toto.bzh@gmail.com", "motdepasse", true, "Michel Toto <michel.toto.bzh@gmail.com>");
     browseObject(test); 
    }

    function testEmailMessage()
    {
     var subject = "[test nsIEmailMessage from messaging idl in JS] hé hè €";
     var priority = EmailPriorityType.HIGH; //high
     var smtpAccount = new SmtpAccount("smtp://smtp.gmail.com:465", "michel.toto.bzh@gmail.com", "motdepasse", true, "Michel Toto <michel.toto.bzh@gmail.com>"); 
     var txt = "<html>   <head>      <meta http-equiv='content-type' content: text/html; charset=UTF-8>   </head><body><font size='2'>Hello !!<br></font><ol><li><font size='2'>This is HTML text, héhéhéhé !</font></li><li><fontsize='2'>Hèèè</font></li><li><font size='2'>€€€</font></li><li><fontsize='2'><b>blablalba</b><br></font></li></ol><font size='2'>bye!<br></font></body></html>";
     var body = new EmailBody(txt, EmailBodyType.HTML);
     
     var att1 = new EmailAttachment("file:///C:/Users/test/Downloads/marianne.jpg");
     var att2 = new EmailAttachment("file:///C:/Users/test/Downloads/toto.txt");
     var att3 = new EmailAttachment("file:///C:/Users/test/Downloads/portecle.zip"); 

     var attachments = new Array(att1, att2); /* array of nsIEmailAttachment */ //  
     var hasMDN = false;
     
     var message = new EmailMessage(subject, priority, smtpAccount, body, new Array(), hasMDN);
     message.addAttachment(att3);
     
     var tab = message.getAttachments();
     var tab = message.tab;
     log("tab.length = " + tab.length);
     for (var i in tab) 
     {
      browseObject(tab[i]);
     }
     
    }

    var emailNumber = 1;

    function SuccessListener(subject)
    {
     this._subject=subject;
    }
    SuccessListener.prototype = {
       handleEvent : function(aIntVal)
       {
      log("The email with subject: " + this._subject + " is successfully sent !");
       },
    }

    function ProgressListener(subject)
    {
     this._subject=subject;
    }
    ProgressListener.prototype = {
       handleEvent : function(aIntVal)
       {
      log("Email subject: " + this._subject + "  => sending progress =  " + aIntVal + " %");
       },
    }

    //globals
    var smtpAccount1 = new SmtpAccount("smtps://smtp.gmail.com:465", "michel.toto.bzh@gmail.com", "motdepasse", true, "michel.toto.bzh@gmail.com");
    var smtpAccount2 = new SmtpAccount("smtps://ssl0.ovh.net:465", "mbonjour%test.fr", "mbonjour%test.fr", true, "michel.bonjour@test.fr");
    var smtpAccount3 = new SmtpAccount("smtp://mailtrap.io:2525", "bidon-8530caec86973507", "f2ae174706e29c59", true, "me@fromdomain.com");

    function testEmailManager()
     {
      var subject = "=€ é[test nsIEmailMessage from messaging idl in JS] n° " + emailNumber;
     emailNumber = emailNumber + 1;
     var priority = EmailPriorityType.HIGH; //high
     var smtpAccount = smtpAccount1;
     
     //var txt = "Money=€\r\nSalut,\r\nMIME définit des mécanismes pour l'envoi d'autre sortes d'informations dont des textes dans des langages autres que l'anglais utilisant des codages de caractères autres que l'ASCII et des données binaires comme des fichiers contenant des images, des sons, des films ou des programmes informatiques. MIME est également un composant fondamental des protocoles de communications comme HTTP, qui requièrent l'envoi de données dans le même contexte que l'envoi de courriels, même si ces données ne sont pas des courriels. L'intégration ou l'extraction des données au format MIME est généralement automatiquement effectuée par le client de messagerie ou par le serveur de messagerie électronique quand le courriel est envoyé ou reçu.\r\n\r\nBye!\r\n";
     //var txt = "<html>   <head>      <meta http-equiv='content-type' content: text/html; charset=UTF-8>   </head><body><font size='2'>Hello !!<br></font><ol><li><font size='2'>héhéhéhé !</font></li><li><fontsize='2'>Hèèè</font></li><li><font size='2'>€€€</font></li><li><fontsize='2'><b>blablalba</b><br></font></li></ol><font size='2'>bye!<br></font></body></html>";
     
     var inlineId_1 = "cid:test@01";
     
     var txt = '<div dir="ltr"><img src="' + inlineId_1 + '"<br></div>'
     var body = new EmailBody(txt, EmailBodyType.HTML); 
     
     var att1 = new EmailAttachment("marianne.jpg"); // relative path OK
     var att2 = new EmailAttachment("file:///C:/Users/test/Downloads/toto.txt");
     var att3 = new EmailAttachment("file:///C:/Users/test/Downloads/portecle.zip");
     
     var attachments = new Array(att1); /* array of nsIEmailAttachment */ 
     
     body.addInlineAttachment(inlineId_1, "file:///C:/Users/test/Downloads/innes.png");
     
     var hasMDN = false;
     
     var message = new EmailMessage(subject, priority, smtpAccount, body, attachments, hasMDN);
      
     var recepList = new EmailRecipientsList(); 
     recepList.addRecipient(EmailRecipientsType.TO, "michel.bonjour@test.fr");
     recepList.addRecipient(EmailRecipientsType.TO, "michel.toto.bzh@gmail.com");
     recepList.addRecipient(EmailRecipientsType.CC, "michel2@mailinator.com");
     recepList.addRecipient(EmailRecipientsType.BCC, "michel3@mailinator.com");
     
     var emailRequest = new EmailRequest();
     var mySuccessListener = new SuccessListener(subject);
     var myProgressListener = new ProgressListener(subject);
     
     log("Start sending email with subject: " + subject); 
     emailRequest = navigator.messaging.email.send(recepList, message);

     emailRequest.addEmailEventListener(EmailEventType.SUCCESS, mySuccessListener);
     emailRequest.addEmailEventListener(EmailEventType.PROGRESS, myProgressListener);
     
     }
       
     
     function testShowSmtpAccounts()
     {  
      log("email.accounts.length = " + navigator.messaging.email.accounts.length);
     for (var i in navigator.messaging.email.accounts) 
     {
      browseObject(navigator.messaging.email.accounts[i]);
     }
     }
     
     function testCancelProtocol()
     {
     // TODO
     email.cancel();
     }
     /* ------------------------------------------------------------------------------------------------------  */

    function init()
    {  
     log("init js"); 
     addTest("testEmailRecepients");
     addTest("testEmailAttachment");
     addTest("testEmailBody");
     addTest("testSmtpAccount");
     addTest("testEmailMessage");
     addTest("testEmailManager");
     addTest("testShowSmtpAccounts");
     addTest("testCancelProtocol");
    }



Member Data Documentation
-------------------------

### readonly attribute nsIEmailManager nsIMessaging::email

Email read-only

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

The nsIEmailManager interface allows to send en email.

Member Data Documentation
-------------------------

### readonly attribute jsval nsIEmailManager::accounts

array of nsISmtpAccount

nsIEmailRequest nsIEmailManager::send (in nsIEmailRecipientsList aReciptientsList, in nsIEmailMessage aMessage)
---------------------------------------------------------------------------------------------------------------

Send message

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aReciptientsList</td>
<td align="left"><p>list of recipients</p></td>
</tr>
<tr class="even">
<td align="left">aMessage</td>
<td align="left"><p>message</p></td>
</tr>
</tbody>
</table>

void nsIEmailManager::cancel ()
-------------------------------

Cancel

nsIEmailMessage interface Reference
===================================

Public Attributes
-----------------

-   attribute DOMString subject

<!-- -->

-   attribute PRUint32 priority

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

-   void getAttachments ( out PRUint32 aLength, out nsIEmailAttachment aAttachments)

<!-- -->

-   void buildData ( in nsIEmailRecipientsList aRecipientsList)

Detailed Description
--------------------

The nsIEmailMessage interface allows to create an email message

Member Data Documentation
-------------------------

### attribute DOMString nsIEmailMessage::subject

subject

### attribute PRUint32 nsIEmailMessage::priority

priority

### attribute nsISmtpAccount nsIEmailMessage::smtpAccount

SMTP account

### attribute nsIEmailBody nsIEmailMessage::body

body

### attribute boolean nsIEmailMessage::hasMDN

has Message Disposition Notifications

### readonly attribute AUTF8String nsIEmailMessage::data

data read-only

void nsIEmailMessage::addAttachment (in nsIEmailAttachment aAttachment)
-----------------------------------------------------------------------

Add email attachment

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAttachment</td>
<td align="left"><p>the email attachment to add</p></td>
</tr>
</tbody>
</table>

void nsIEmailMessage::removeAttachment (in nsIEmailAttachment aAttachment)
--------------------------------------------------------------------------

Remove email attachment

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aAttachment</td>
<td align="left"><p>the email attachment to remove</p></td>
</tr>
</tbody>
</table>

void nsIEmailMessage::getAttachments (\[optional\] out PRUint32 aLength, \[array, size\_is(aLength), retval, optional\] out nsIEmailAttachment aAttachments)
------------------------------------------------------------------------------------------------------------------------------------------------------------

Get email attachment

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>array length</p></td>
</tr>
<tr class="even">
<td align="left">aAttachments</td>
<td align="left"><p>array of attachments</p></td>
</tr>
</tbody>
</table>

void nsIEmailMessage::buildData (in nsIEmailRecipientsList aRecipientsList)
---------------------------------------------------------------------------

Build data from list of recipients

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientsList</td>
<td align="left"><p>the list of recipients</p></td>
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

-   void init ( in AString fileUri)

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIEmailAttachment::mimeType

mime type

### readonly attribute DOMString nsIEmailAttachment::filePath

file path

### readonly attribute DOMString nsIEmailAttachment::fileName

file name

void nsIEmailAttachment::init (in AString fileUri)
--------------------------------------------------

Initialize from an URI

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">fileUri</td>
<td align="left"><p>URI of the attachment</p></td>
</tr>
</tbody>
</table>

nsIInlineEmailAttachment interface Reference
============================================

Public Attributes
-----------------

-   readonly attribute DOMString inlineID

-   void initInline ( in AString fileUri, in AString aId)

Detailed Description
--------------------

The nsIInlineEmailAttachment interface allows to manage an inline email attachment

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIInlineEmailAttachment::inlineID

inline ID

void nsIInlineEmailAttachment::initInline (in AString fileUri, in AString aId)
------------------------------------------------------------------------------

Initialize inline

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">fileUri</td>
<td align="left"><p>file URI</p></td>
</tr>
<tr class="even">
<td align="left">aId</td>
<td align="left"><p>identifier</p></td>
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

The nsIEmailBodyType interface is an enumeration for email body type

Member Data Documentation
-------------------------

### const nsEmailBodyTypeValue nsIEmailBodyType::PLAIN\_TEXT

plain text

### const nsEmailBodyTypeValue nsIEmailBodyType::HTML

html

nsIEmailBody interface Reference
================================

Public Attributes
-----------------

-   readonly attribute DOMString body

<!-- -->

-   readonly attribute PRUint32 bodyType

-   void addInlineAttachment ( in DOMString aID, in DOMString aFileUri)

<!-- -->

-   void removeInlineAttachment ( in DOMString aID)

<!-- -->

-   void getArrayOfInlineAttachments ( out PRUint32 aLength, out nsIInlineEmailAttachment aAttachments)

Detailed Description
--------------------

The nsIEmailBody interface allows to manage en email body

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIEmailBody::body

body

### readonly attribute PRUint32 nsIEmailBody::bodyType

type of body

void nsIEmailBody::addInlineAttachment (in DOMString aID, in DOMString aFileUri)
--------------------------------------------------------------------------------

Add an inline attachment

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aID</td>
<td align="left"><p>string id must be begin with &quot;cid:&quot; or &quot;mid:&quot;</p></td>
</tr>
<tr class="even">
<td align="left">aFileUri</td>
<td align="left"><p>uri</p></td>
</tr>
</tbody>
</table>

void nsIEmailBody::removeInlineAttachment (in DOMString aID)
------------------------------------------------------------

Remove an inline attachment

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aID</td>
<td align="left"><p>string id must be begin with &quot;cid:&quot; or &quot;mid:&quot;</p></td>
</tr>
</tbody>
</table>

void nsIEmailBody::getArrayOfInlineAttachments (\[optional\] out PRUint32 aLength, \[array, size\_is(aLength), retval, optional\] out nsIInlineEmailAttachment aAttachments)
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Get array of inline attachments

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>array length</p></td>
</tr>
<tr class="even">
<td align="left">aAttachments</td>
<td align="left"><p>array of attachments</p></td>
</tr>
</tbody>
</table>

nsIEmailEventListener interface Reference
=========================================

-   void handleEvent ( in PRUint32 aIntVal)

Detailed Description
--------------------

The nsIEmailEventListener interface provide a listener for email event

void nsIEmailEventListener::handleEvent (in PRUint32 aIntVal)
-------------------------------------------------------------

Handle event

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

success

### const nsEmailEventTypeValue nsIEmailEventType::ERROR

error

### const nsEmailEventTypeValue nsIEmailEventType::PROGRESS

in progress

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

priority normal

### const nsPriorityTypeValue nsIEmailPriorityType::HIGH

priority high

### const nsPriorityTypeValue nsIEmailPriorityType::LOW

priority low

nsIEmailRecipientsList interface Reference
==========================================

-   void addRecipient ( in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)

<!-- -->

-   void removeRecipient ( in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)

<!-- -->

-   void getAllRecipients ( out PRUint32 aLength, out wstring aRecipients)

<!-- -->

-   void getRecipients ( in nsRecipientsTypeValue aRecipientType, out PRUint32 aLength, out wstring aRecipients)

Detailed Description
--------------------

The nsIEmailRecipientsList interface allows to manage a list of recipients

void nsIEmailRecipientsList::addRecipient (in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)
------------------------------------------------------------------------------------------------------------

Add a recipient

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientType</td>
<td align="left"><p>type of recipient</p></td>
</tr>
<tr class="even">
<td align="left">aRecipient</td>
<td align="left"><p>recipient to add</p></td>
</tr>
</tbody>
</table>

void nsIEmailRecipientsList::removeRecipient (in nsRecipientsTypeValue aRecipientType, in DOMString aRecipient)
---------------------------------------------------------------------------------------------------------------

Remove a recipient

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientType</td>
<td align="left"><p>type of recipient</p></td>
</tr>
<tr class="even">
<td align="left">aRecipient</td>
<td align="left"><p>recipient to remove</p></td>
</tr>
</tbody>
</table>

void nsIEmailRecipientsList::getAllRecipients (\[optional\] out PRUint32 aLength, \[array, size\_is(aLength), retval, optional\] out wstring aRecipients)
---------------------------------------------------------------------------------------------------------------------------------------------------------

Get all recipients

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aLength</td>
<td align="left"><p>array length</p></td>
</tr>
<tr class="even">
<td align="left">aRecipients</td>
<td align="left"><p>array of recipients</p></td>
</tr>
</tbody>
</table>

void nsIEmailRecipientsList::getRecipients (in nsRecipientsTypeValue aRecipientType, \[optional\] out PRUint32 aLength, \[array, size\_is(aLength), retval, optional\] out wstring aRecipients)
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Get some recipients

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRecipientType</td>
<td align="left"><p>type of recipient</p></td>
</tr>
<tr class="even">
<td align="left">aLength</td>
<td align="left"><p>array length</p></td>
</tr>
<tr class="odd">
<td align="left">aRecipients</td>
<td align="left"><p>array of recipients</p></td>
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

value TO

### const nsRecipientsTypeValue nsIEmailRecipientsType::CC

value CC

### const nsRecipientsTypeValue nsIEmailRecipientsType::BCC

value BCC

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

-   void handleProgress ( in PRUint32 aProgressPercentage)

Detailed Description
--------------------

The nsIEmailRequest interface is a request for email event

Member Data Documentation
-------------------------

### readonly attribute DOMString nsIEmailRequest::readyState

readyState

void nsIEmailRequest::addEmailEventListener (in nsEmailEventTypeValue aEventType, in nsIEmailEventListener aListener)
---------------------------------------------------------------------------------------------------------------------

Add email event listener

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aEventType</td>
<td align="left"><p>type of event</p></td>
</tr>
<tr class="even">
<td align="left">aListener</td>
<td align="left"><p>listener</p></td>
</tr>
</tbody>
</table>

void nsIEmailRequest::handleSuccess ()
--------------------------------------

Handle success

void nsIEmailRequest::handleError ()
------------------------------------

Handle error

void nsIEmailRequest::handleProgress (in PRUint32 aProgressPercentage)
----------------------------------------------------------------------

handle in progress

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

-   readonly attribute PRUint32 port

Detailed Description
--------------------

The nsISmtpAccount interface provide information about STMP account

Member Data Documentation
-------------------------

### readonly attribute DOMString nsISmtpAccount::serverUri

server URI

### readonly attribute DOMString nsISmtpAccount::username

username

### readonly attribute DOMString nsISmtpAccount::password

password

### readonly attribute boolean nsISmtpAccount::authorizeUncrytptedAuth

authorizeUncrytptedAuth

### readonly attribute DOMString nsISmtpAccount::sender

sender

### readonly attribute DOMString nsISmtpAccount::hostname

hostname

### readonly attribute boolean nsISmtpAccount::useSSL

useSSL

### readonly attribute PRUint32 nsISmtpAccount::port

port

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

No SSL or STARTTLS

### const nsMsgSocketTypeValue nsMsgSocketType::trySTARTTLS

Use TLS via STARTTLS, but only if server offers it.

Deprecated

This is vulnerable to MITM attacks

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

nsISmtpServer.authMethod

Member Data Documentation
-------------------------

### const nsMsgAuthMethodValue nsMsgAuthMethod::none

No login needed. E.g. IP-address-based.

### const nsMsgAuthMethodValue nsMsgAuthMethod::old

Do not use AUTH commands (e.g. AUTH=PLAIN), but the original login commands that the protocol specified (POP: "USER"/"PASS", IMAP: "login", not valid for SMTP)

### const nsMsgAuthMethodValue nsMsgAuthMethod::passwordCleartext

password in the clear. AUTH=PLAIN/LOGIN or old-style login.

### const nsMsgAuthMethodValue nsMsgAuthMethod::passwordEncrypted

hashed password. CRAM-MD5, DIGEST-MD5

### const nsMsgAuthMethodValue nsMsgAuthMethod::GSSAPI

Kerberos / GSSAPI (Unix single-signon)

### const nsMsgAuthMethodValue nsMsgAuthMethod::NTLM

NTLM is a Windows single-singon scheme. Includes MSN / Passport.net, which is the same with a different name.

### const nsMsgAuthMethodValue nsMsgAuthMethod::External

Auth External is cert-based authentication

### const nsMsgAuthMethodValue nsMsgAuthMethod::secure

Encrypted password or Kerberos / GSSAPI or NTLM.

Deprecated

-   for migration only.

### const nsMsgAuthMethodValue nsMsgAuthMethod::anything

Let us pick any of the auth types supported by the server. Discouraged, because vulnerable to MITM attacks, even if server offers secure auth.
