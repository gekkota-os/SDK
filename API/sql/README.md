nsIGktSqlConnection interface Reference
=======================================

    #include <nsIGktSqlConnection.idl>

Public Attributes
-----------------

-   readonly attribute AString serverVersion

    *A string holding the name and/or version info of the database.*

<!-- -->

-   readonly attribute AString errorMessage

    *The most recent error message.*

<!-- -->

-   readonly attribute long lastID

    *The ID of the most recently added record.*

-   void init ( in AString aType, in AString aHost, in long aPort, in AString aDatabase, in AString aUsername, in AString aPassword)

    *Set up the connection.*

<!-- -->

-   nsIGktSqlResult executeQuery ( in AString aQuery)

    *Execute an SQL query synchronously and return the query result.*

<!-- -->

-   long executeUpdate ( in AString aUpdate)

    *Execute an SQL update synchronously and return the number of updated rows.*

<!-- -->

-   nsIGktSqlRequest asyncExecuteQuery ( in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)

    *Execute an SQL query asynchronously and return a request.*

<!-- -->

-   nsIGktSqlRequest asyncExecuteUpdate ( in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)

    *Execute an SQL query for update asynchronously and return a request.*

<!-- -->

-   void beginTransaction ( )

    *Begin a transaction.*

<!-- -->

-   void commitTransaction ( )

    *Commit the current transaction.*

<!-- -->

-   void rollbackTransaction ( )

    *Rollback (cancel) the current transaction.*

<!-- -->

-   nsIGktSqlResult getPrimaryKeys ( in AString aSchema, in AString aTable)

    *Return the primary keys of a given table.*

Detailed Description
--------------------

The nsIGktSqlConnection interface provides few methods and attributes to initialize the connection to database and execute some requests. This interface can not be instantiated. See nsIGktSqlConnectionODBC.

Definition at line 29 of file nsIGktSqlConnection.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlConnection.idl

Member Data Documentation
-------------------------

void nsIGktSqlConnection::init (in AString aType, in AString aHost, in long aPort, in AString aDatabase, in AString aUsername, in AString aPassword)
----------------------------------------------------------------------------------------------------------------------------------------------------

Set up the connection.
This is called by the SQL service. There is no need to call this method directly.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aType</td>
<td align="left"><p>the type of database : sqloledb, MySQLProv</p></td>
</tr>
<tr class="even">
<td align="left">aHost</td>
<td align="left"><p>the host name.</p></td>
</tr>
<tr class="odd">
<td align="left">aPort</td>
<td align="left"><p>the port at which the host is listening.</p></td>
</tr>
<tr class="even">
<td align="left">aDatabase</td>
<td align="left"><p>the real database name to connect to.</p></td>
</tr>
<tr class="odd">
<td align="left">aUsername</td>
<td align="left"><p>the username to connect as.</p></td>
</tr>
<tr class="even">
<td align="left">aPassword</td>
<td align="left"><p>the password to use in authentification phase.</p></td>
</tr>
</tbody>
</table>

nsIGktSqlResult nsIGktSqlConnection::executeQuery (in AString aQuery)
---------------------------------------------------------------------

Execute an SQL query synchronously and return the query result.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aQuery</td>
<td align="left"><p>the SQL string of the query to execute</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the result of the query

    <html>
    <head>
    <title></title>
    </head>
    <body bgcolor="white">
    <script language="javascript">
     var mysql_host = "192.168.1.30";
     var mysql_database = "grrDBTest";
     var mysql_user = "root";
     var mysql_password = ""; //to set
     var type_database = "MySQLProv";
     var request_valid_mysql = "select * from grr_utilisateurs";

        try{
            var conn_mysql = new GktSqlConnectionODBC();
            conn_mysql.init(type_database, mysql_host , 3306, mysql_database, mysql_user, mysql_password);
            var result_mysql = conn_mysql.executeQuery(request_valid_mysql);
            var enum_mysql = result_mysql.enumerate();
            enum_mysql.beforeFirst();
      var resNext = true;
            while(resNext){
       resNext = enum_mysql.next();
                for(var  i= 0; i< result_mysql.columnCount; i++){
                    if(result_mysql.getColumnType(i) == Components.interfaces.nsIGktSqlResult.TYPE_STRING){
                        document.write("<br>data found : " + enum_mysql.getString(i)) ;
                    }
                }
            }
        }catch(e){
            document.write("<br>MySQL Exception : " + e);
        }
    </script>
    </body>
    </html>

long nsIGktSqlConnection::executeUpdate (in AString aUpdate)
------------------------------------------------------------

Execute an SQL update synchronously and return the number of updated rows.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aUpdate</td>
<td align="left"><p>the update to execute</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the result of the query

    <html>
    <head>
    <title></title>
    </head>
    <body bgcolor="white">
    <script language="javascript">
     document.write("test_transact<br><br>");
     var mysql_host = "192.168.1.30";
     var port = 3306;
     var mysql_database = "grrDBTest";
     var mysql_user = "root";
     var mysql_password = ""; // to set
     var type_database = "MySQLProv";
     var request_valid_mysql = "UPDATE grr_utilisateurs SET login='login_" + Math.random() + "' WHERE 1=1";

        try{
            var conn_mysql = new GktSqlConnectionODBC();
            conn_mysql.init(type_database, mysql_host , port, mysql_database, mysql_user, mysql_password);
      conn_mysql.beginTransaction();
            var rowAffected  = conn_mysql.executeUpdate(request_valid_mysql);
      conn_mysql.commitTransaction();
      document.write("<br>rows affected synchronously : " + rowAffected);
        }catch(e){
            document.write("<br>MySQL Exception : " + e);
        }
    </script>
    </body>
    </html>

nsIGktSqlRequest nsIGktSqlConnection::asyncExecuteQuery (in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)
-------------------------------------------------------------------------------------------------------------------------------------------

Execute an SQL query asynchronously and return a request.
An observer may be used to track when the query has completed.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aQuery</td>
<td align="left"><p>the SQL string of the query to execute</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>extra argument that will be passed to the observer</p></td>
</tr>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>observer that will be notified when the query is done</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a request object

    <html>
    <head>
    <title></title>
    </head>
    <body bgcolor="white">
    <script language="javascript">
        var elt = document.createElement("div");
        elt.setAttribute("id", "123");
        document.body.insertBefore(elt, null);
        function appendLog(msg){
      var e = window.document.getElementById("123");
      e.innerHTML =  e.innerHTML + " <br> " + msg;
     }
        var type_database = "MySQLProv";
        var host = "192.168.1.30";
        var port = 3306;
        var database = "grrDBTest";
        var user = "root";
        var password = ""; //to set
        var request = "select * from grr_utilisateurs";

        try{
            // connect to database
            var conn = new GktSqlConnectionODBC();
            conn.init(type_database, host, port, database, user, password);
            // observer
            var reqObserver = {
                    onStartRequest: function( request, context){
                        appendLog("onStartRequest");
                    },
                    onStopRequest: function(request, context){
         try{
          appendLog("onStopRequest : " + request.status) ;
          var result_mysql = request.result;
          var enum_mysql = result_mysql.enumerate();
          // go before first element
          enum_mysql.beforeFirst();
          // for each row, display only string data
          var resNext = true;
          while(resNext){
           resNext = enum_mysql.next();
           for(var  i= 0; i< result_mysql.columnCount; i++){
            if(result_mysql.getColumnType(i) == Components.interfaces.nsIGktSqlResult.TYPE_STRING 
               && enum_mysql.isNull(i) == false ){
             appendLog("data found : " + enum_mysql.getString(i));
            }
            if( ( result_mysql.getColumnType(i) == Components.interfaces.nsIGktSqlResult.TYPE_DATETIME || 
               result_mysql.getColumnType(i) == Components.interfaces.nsIGktSqlResult.TYPE_DATE || 
               result_mysql.getColumnType(i) == Components.interfaces.nsIGktSqlResult.TYPE_TIME )   
               && enum_mysql.isNull(i) == false ){
             appendLog(result_mysql.getColumnName(i) + " : " + new Date(enum_mysql.getDate(i)/1000).toString()) ;
            }
           } 
          }
         }catch(ex){
          appendLog("exception : " + ex );
         }
                    },      
            } 
            // run async.
            conn.asyncExecuteQuery(request, null, reqObserver);
        }catch(e){
            appendLog("Exception : " + e);
        }
    </script>
    </body>
    </html>

nsIGktSqlRequest nsIGktSqlConnection::asyncExecuteUpdate (in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)
--------------------------------------------------------------------------------------------------------------------------------------------

Execute an SQL query for update asynchronously and return a request.
An observer may be used to track when the query has completed.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aQuery</td>
<td align="left"><p>the SQL string of the query to execute</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>extra argument that will be passed to the observer</p></td>
</tr>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>observer that will be notified when the query is done</p></td>
</tr>
</tbody>
</table>

**Returns:.**

a request object

    <html>
    <head>
    <title></title>
    </head>
    <body bgcolor="white">
    <script language="javascript">
        var elt = document.createElement("div");
        elt.setAttribute("id", "123");
        document.body.insertBefore(elt, null);
     function appendLog(msg){
      var e = window.document.getElementById("123");
      e.innerHTML =  e.innerHTML + " <br> " + msg;
     }
        var type_database = "MySQLProv";
        var host = "192.168.1.30";
        var port = 3306;
        var database = "grrDBTest";
        var user =  "root";
        var password = ""; //to set
        var request = "UPDATE grr_utilisateurs SET login='Administrateur_async_" + Math.random() + "' WHERE 1=1";
        try{
            var conn = new GktSqlConnectionODBC();
            conn.init(type_database, host, port, database, user, password);
            var reqObserver = {
                    onStartRequest: function( request, context){
                       appendLog("onStartRequest");
                    },
                    onStopRequest: function(request, context){
         try{
          appendLog("onStopRequest : " + request.status);
          appendLog("Rows affected : " + request.affectedRows);
         }catch(ex){
          appendLog("exception : " + ex) ;
         }
                    },      
            } 
            conn.asyncExecuteUpdate(request, null, reqObserver);
        }catch(e){
            appendLog("Exception : " + e);
        }
    </script>
    </body>
    </html>

void nsIGktSqlConnection::beginTransaction ()
---------------------------------------------

Begin a transaction.
Updates made during the transaction will not be made permanent until it is committed using commitTransaction.

void nsIGktSqlConnection::commitTransaction ()
----------------------------------------------

Commit the current transaction.
void nsIGktSqlConnection::rollbackTransaction ()
------------------------------------------------

Rollback (cancel) the current transaction.
nsIGktSqlResult nsIGktSqlConnection::getPrimaryKeys (in AString aSchema, in AString aTable)
-------------------------------------------------------------------------------------------

Return the primary keys of a given table.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aSchema</td>
<td align="left"><p>the schema</p></td>
</tr>
<tr class="even">
<td align="left">aTable</td>
<td align="left"><p>the table name of the keys to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the result which holds the keys

nsIGktSqlConnectionODBC interface Reference
===========================================

    #include <nsIGktSqlConnectionODBC.idl>

Detailed Description
--------------------

The nsIGktSqlConnectionODBC interface inherits from nsIGktSqlConnection and implements a part of its methods.

In javascript, instantiating this object with the following code : new GktSqlConnectionODBC ()

It is the point of entry to connect to an ODBC-compatible database and execute some requests.

Definition at line 34 of file nsIGktSqlConnectionODBC.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlConnectionODBC.idl

nsIGktSqlRequest interface Reference
====================================

    #include <nsIGktSqlRequest.idl>

Public Attributes
-----------------

-   readonly attribute AString errorMessage

    *The most recent error message.*

<!-- -->

-   readonly attribute nsIGktSqlResult result

    *The result of the operation.*

<!-- -->

-   readonly attribute long affectedRows

    *The number of rows that were affected during an update.*

<!-- -->

-   readonly attribute long lastID

    *The ID of the most recently added record.*

<!-- -->

-   readonly attribute AString query

    *The SQL query.*

<!-- -->

-   readonly attribute nsISupports ctxt

    *The context passed to the connection's asyncExecuteQuery or asyncExecuteUpdate method.*

<!-- -->

-   readonly attribute nsIGktSqlRequestObserver observer

    *The observer that listens for when the request is finished.*

<!-- -->

-   const long STATUS\_NONE

    *Status none.*

<!-- -->

-   const long STATUS\_EXECUTED

    *Status executed.*

<!-- -->

-   const long STATUS\_COMPLETE

    *Status complete.*

<!-- -->

-   const long STATUS\_ERROR

    *Status error.*

<!-- -->

-   const long STATUS\_CANCELLED

    *Status cancelled.*

<!-- -->

-   readonly attribute long status

    *The status of the request.*

-   void cancel ( )

    *Cancels the operation.*

Detailed Description
--------------------

The nsIGktSqlRequest interface implements requests used during an asynchronous SQL query or update operation.

Definition at line 24 of file nsIGktSqlRequest.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlRequest.idl

Member Data Documentation
-------------------------

### readonly attribute nsIGktSqlRequestObserver nsIGktSqlRequest::observer

Methods of the observer should be called by the request.

Definition at line 65 of file nsIGktSqlRequest.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlRequest.idl

void nsIGktSqlRequest::cancel ()
--------------------------------

Cancels the operation.
nsIGktSqlRequestObserver interface Reference
============================================

    #include <nsIGktSqlRequestObserver.idl>

-   void onStartRequest ( in nsIGktSqlRequest aRequest, in nsISupports aContext)

    *This method will be called when the request is started.*

<!-- -->

-   void onStopRequest ( in nsIGktSqlRequest aRequest, in nsISupports aContext)

    *This method will be called when the request has finished.*

The nsIGktSqlRequestObserver interface is used to listen to asynchronous SQL query or update requests.

void nsIGktSqlRequestObserver::onStartRequest (in nsIGktSqlRequest aRequest, in nsISupports aContext)
-----------------------------------------------------------------------------------------------------

This method will be called when the request is started.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRequest</td>
<td align="left"><p>the request that has started</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>a context that was supplied in the query/update call</p></td>
</tr>
</tbody>
</table>

void nsIGktSqlRequestObserver::onStopRequest (in nsIGktSqlRequest aRequest, in nsISupports aContext)
----------------------------------------------------------------------------------------------------

This method will be called when the request has finished.
This function will be called in both success and error cases.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRequest</td>
<td align="left"><p>the request that has started</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>a context that was supplied in the query/update call</p></td>
</tr>
</tbody>
</table>

nsIGktSqlResult interface Reference
===================================

    #include <nsIGktSqlResult.idl>

Public Attributes
-----------------

-   attribute boolean displayNullAsText

    *By default, this value is false.*

<!-- -->

-   readonly attribute nsIGktSqlConnection connection

    *The connection used to execute the query.*

<!-- -->

-   readonly attribute AString query

    *The SQL query.*

<!-- -->

-   readonly attribute AString tableName

    *The table that was used in the query.*

<!-- -->

-   readonly attribute long rowCount

    *The number of rows in the result.*

<!-- -->

-   readonly attribute long columnCount

    *The number of columns in the result.*

<!-- -->

-   const long TYPE\_STRING

    *column type constants used by |getColumnType|.*

<!-- -->

-   const long TYPE\_INT

    *Type integer.*

<!-- -->

-   const long TYPE\_FLOAT

    *Type float.*

<!-- -->

-   const long TYPE\_DECIMAL

    *Type decimal.*

<!-- -->

-   const long TYPE\_DATE

    *Type date.*

<!-- -->

-   const long TYPE\_TIME

    *Type time.*

<!-- -->

-   const long TYPE\_DATETIME

    *Type datetime.*

<!-- -->

-   const long TYPE\_BOOL

    *Type boolean.*

-   AString getColumnName ( in long aColumnIndex)

    *Retrieves the name of a column given its index.*

<!-- -->

-   long getColumnIndex ( in AString aColumnName)

    *Retrieves the index of a column given its name.*

<!-- -->

-   long getColumnType ( in long aColumnIndex)

    *Returns the type of the data in a given column.*

<!-- -->

-   AString getColumnTypeAsString ( in long aColumnIndex)

    *Returns the type of the data in a given column as a string.*

<!-- -->

-   long getColumnDisplaySize ( in long aColumnIndex)

    *Returns the maximum number of bytes that are needed to hold a value in a particular column.*

<!-- -->

-   nsIGktSqlResultEnumerator enumerate ( )

    *Returns an enumerator to enumerator over the returned rows.*

<!-- -->

-   void reload ( )

    *Re-executes the query.*

Detailed Description
--------------------

The nsIGktSqlResult interface manage results of an SQL query. Use the enumerate method to retrieve each row.

Definition at line 23 of file nsIGktSqlResult.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlResult.idl

Member Data Documentation
-------------------------

### readonly attribute AString nsIGktSqlResult::tableName

If more than one table was used, only the first is returned.

Definition at line 48 of file nsIGktSqlResult.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlResult.idl

### const long nsIGktSqlResult::TYPE\_STRING

Type string

Definition at line 84 of file nsIGktSqlResult.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlResult.idl

AString nsIGktSqlResult::getColumnName (in long aColumnIndex)
-------------------------------------------------------------

Retrieves the name of a column given its index.
Indicies start at zero.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the index of the column to return</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the column name

long nsIGktSqlResult::getColumnIndex (in AString aColumnName)
-------------------------------------------------------------

Retrieves the index of a column given its name.
If the column does not exist, -1 is returned.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnName</td>
<td align="left"><p>the column name to return</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the column index

long nsIGktSqlResult::getColumnType (in long aColumnIndex)
----------------------------------------------------------

Returns the type of the data in a given column.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the index of the column to return the type of</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the column type

AString nsIGktSqlResult::getColumnTypeAsString (in long aColumnIndex)
---------------------------------------------------------------------

Returns the type of the data in a given column as a string.
This is used as an alternative to using the constants and will return either string, int, float, decimal, date, time, datetime or bool.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the index of the column to return the type of</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the column type

long nsIGktSqlResult::getColumnDisplaySize (in long aColumnIndex)
-----------------------------------------------------------------

Returns the maximum number of bytes that are needed to hold a value in a particular column.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the index of the column to return the size of</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the column size

nsIGktSqlResultEnumerator nsIGktSqlResult::enumerate ()
-------------------------------------------------------

Returns an enumerator to enumerator over the returned rows.
**Returns:.**

the row enumerator

void nsIGktSqlResult::reload ()
-------------------------------

Re-executes the query.
nsIGktSqlResultEnumerator interface Reference
=============================================

    #include <nsIGktSqlResultEnumerator.idl>

Public Attributes
-----------------

-   readonly attribute AString errorMessage

    *The most recent error message.*

<!-- -->

-   readonly attribute AString currentCondition

    *Holds the SQL condition clause.*

-   boolean next ( )

    *Moves the row pointer to the next row in the results.*

<!-- -->

-   boolean previous ( )

    *Moves the row pointer to the previous row in the results.*

<!-- -->

-   void beforeFirst ( )

    *Moves the row pointer to just before the first row.*

<!-- -->

-   void first ( )

    *Moves the row pointer to the first row.*

<!-- -->

-   void last ( )

    *Moves the row pointer to the last row.*

<!-- -->

-   void relative ( in long aRowIndex)

    *Moves the row pointer by a number relative to the current row.*

<!-- -->

-   void absolute ( in long aRowIndex)

    *Moves the row pointer to a specific row.*

<!-- -->

-   boolean isNull ( in long aColumnIndex)

    *Returns true if the value at the specified column in the current row is null.*

<!-- -->

-   nsIVariant getVariant ( in long aColumnIndex)

    *Returns the value at the specified column in the current row as a variant.*

<!-- -->

-   AString getString ( in long aColumnIndex)

    *Returns the value at the specified column in the current row as a string.*

<!-- -->

-   long getInt ( in long aColumnIndex)

    *Returns the value at the specified column in the current row as an integer.*

<!-- -->

-   float getFloat ( in long aColumnIndex)

    *Returns the value at the specified column in the current row as a float.*

<!-- -->

-   float getDecimal ( in long aColumnIndex)

    *Returns the value at the specified column in the current row as a decimal.*

<!-- -->

-   long long getDate ( in long aColumnIndex)

    *Returns the value at the specified column in the current row as a date.*

<!-- -->

-   boolean getBool ( in long aColumnIndex)

    *Returns the value at the specified column in the current row as a boolean.*

Detailed Description
--------------------

The nsIGktSqlResultEnumerator interface is used to get the results from an SQL query. The enumerator uses a row pointer which can be adjusted with the next and previous methods. Other methods operate only on the row selected by the pointer.

The row pointer starts just before the first row, so you should always call the next method once before attempting to read row data.

Definition at line 29 of file nsIGktSqlResultEnumerator.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlResultEnumerator.idl

Member Data Documentation
-------------------------

boolean nsIGktSqlResultEnumerator::next ()
------------------------------------------

Moves the row pointer to the next row in the results.
Returns true if there is a next row and false if there are no more rows.

**Returns:.**

false if there are no more rows

boolean nsIGktSqlResultEnumerator::previous ()
----------------------------------------------

Moves the row pointer to the previous row in the results.
Returns true if there is a previous row.

**Returns:.**

false if there are no previous rows

void nsIGktSqlResultEnumerator::beforeFirst ()
----------------------------------------------

Moves the row pointer to just before the first row.
void nsIGktSqlResultEnumerator::first ()
----------------------------------------

Moves the row pointer to the first row.
void nsIGktSqlResultEnumerator::last ()
---------------------------------------

Moves the row pointer to the last row.
void nsIGktSqlResultEnumerator::relative (in long aRowIndex)
------------------------------------------------------------

Moves the row pointer by a number relative to the current row.
An error occurs if this causes the row pointer to extend past the last row. This method may also be used to move the row pointer back by using a negative value.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRowIndex</td>
<td align="left"><p>aRowIndex the number of rows to move by</p></td>
</tr>
</tbody>
</table>

void nsIGktSqlResultEnumerator::absolute (in long aRowIndex)
------------------------------------------------------------

Moves the row pointer to a specific row.
An error occurs if the index is after the last row.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRowIndex</td>
<td align="left"><p>the index of the row to move to</p></td>
</tr>
</tbody>
</table>

boolean nsIGktSqlResultEnumerator::isNull (in long aColumnIndex)
----------------------------------------------------------------

Returns true if the value at the specified column in the current row is null.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

true if the value is null

nsIVariant nsIGktSqlResultEnumerator::getVariant (in long aColumnIndex)
-----------------------------------------------------------------------

Returns the value at the specified column in the current row as a variant.
<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the value as a variant

AString nsIGktSqlResultEnumerator::getString (in long aColumnIndex)
-------------------------------------------------------------------

Returns the value at the specified column in the current row as a string.
An error occurs if the value is not a string type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the string value

long nsIGktSqlResultEnumerator::getInt (in long aColumnIndex)
-------------------------------------------------------------

Returns the value at the specified column in the current row as an integer.
An error occurs if the value is not a integer type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the integer value

float nsIGktSqlResultEnumerator::getFloat (in long aColumnIndex)
----------------------------------------------------------------

Returns the value at the specified column in the current row as a float.
An error occurs if the value is not a float type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the float value

float nsIGktSqlResultEnumerator::getDecimal (in long aColumnIndex)
------------------------------------------------------------------

Returns the value at the specified column in the current row as a decimal.
An error occurs if the value is not a decimal type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the decimal value

long long nsIGktSqlResultEnumerator::getDate (in long aColumnIndex)
-------------------------------------------------------------------

Returns the value at the specified column in the current row as a date.
An error occurs if the value is not a date type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the date value

boolean nsIGktSqlResultEnumerator::getBool (in long aColumnIndex)
-----------------------------------------------------------------

Returns the value at the specified column in the current row as a boolean.
An error occurs if the value is not a boolean type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>the column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

the boolean value

nsIGktSqlResultODBC interface Reference
=======================================

    #include <nsIGktSqlResultODBC.idl>

Detailed Description
--------------------

The nsIGktSqlResultODBC interface manage the results of an SQL query for ODBC databases. Implements some nsIGtkSqlResult methods and attributes.

Definition at line 23 of file nsIGktSqlResultODBC.idl

The Documentation for this struct was generated from the following file:

-   nsIGktSqlResultODBC.idl


