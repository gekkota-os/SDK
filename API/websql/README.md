nsIGktSqlConnection interface Reference
=======================================

Public Attributes
-----------------

-   readonly attribute AString serverVersion

<!-- -->

-   readonly attribute AString errorMessage

<!-- -->

-   readonly attribute long lastID

-   void init ( in AString aType, in AString aHost, in long aPort, in AString aDatabase, in AString aUsername, in AString aPassword)

<!-- -->

-   nsIGktSqlResult executeQuery ( in AString aQuery)

<!-- -->

-   long executeUpdate ( in AString aUpdate)

<!-- -->

-   nsIGktSqlRequest asyncExecuteQuery ( in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)

<!-- -->

-   nsIGktSqlRequest asyncExecuteUpdate ( in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)

<!-- -->

-   void beginTransaction ( )

<!-- -->

-   void commitTransaction ( )

<!-- -->

-   void rollbackTransaction ( )

<!-- -->

-   nsIGktSqlResult getPrimaryKeys ( in AString aSchema, in AString aTable)

Detailed Description
--------------------

The nsIGktSqlConnection interface provides few methods and attributes to initialize the connection to database and execute some requests. This interface cannot be instantiated. See nsIGktSqlConnectionODBC. HTML example using this API [here.](example1.html)

**Build note**: You need to execute the **build.cmd** file to generate the boostrap app. Otherwise there will be a mismatch between the html file name and the one the manifest tries to launch. Find more information in *SDK-G4/bootstrap App/* documentation.

Member Data Documentation
-------------------------

### readonly attribute AString nsIGktSqlConnection::serverVersion

A string holding the name and/or version info of the database.

### readonly attribute AString nsIGktSqlConnection::errorMessage

The most recent error message.

### readonly attribute long nsIGktSqlConnection::lastID

The ID of the most recently added record.

void nsIGktSqlConnection::init (in AString aType, in AString aHost, in long aPort, in AString aDatabase, in AString aUsername, in AString aPassword)
----------------------------------------------------------------------------------------------------------------------------------------------------

Set up the connection. This is called by the SQL service. There is no need to call this method directly.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aType</td>
<td align="left"><p>The type of database: sqloledb, MySQLProv</p></td>
</tr>
<tr class="even">
<td align="left">aHost</td>
<td align="left"><p>The host name</p></td>
</tr>
<tr class="odd">
<td align="left">aPort</td>
<td align="left"><p>The port at which the host is listening</p></td>
</tr>
<tr class="even">
<td align="left">aDatabase</td>
<td align="left"><p>The real database name to connect to</p></td>
</tr>
<tr class="odd">
<td align="left">aUsername</td>
<td align="left"><p>The username to connect as</p></td>
</tr>
<tr class="even">
<td align="left">aPassword</td>
<td align="left"><p>The password to use in authentication phase</p></td>
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
<td align="left"><p>The SQL string of the query to execute</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The result of the query.

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
<td align="left"><p>The update to execute</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The result of the query.

nsIGktSqlRequest nsIGktSqlConnection::asyncExecuteQuery (in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)
-------------------------------------------------------------------------------------------------------------------------------------------

Execute an SQL query asynchronously and return a request. An observer may be used to track when the query has completed.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aQuery</td>
<td align="left"><p>The SQL string of the query to execute</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>Extra argument that will be passed to the observer</p></td>
</tr>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>Observer that will be notified when the query is done</p></td>
</tr>
</tbody>
</table>

**Returns:.**

A request object.

nsIGktSqlRequest nsIGktSqlConnection::asyncExecuteUpdate (in AString aQuery, in nsISupports aContext, in nsIGktSqlRequestObserver aObserver)
--------------------------------------------------------------------------------------------------------------------------------------------

Execute an SQL query for update asynchronously and return a request. An observer may be used to track when the query has completed.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aQuery</td>
<td align="left"><p>The SQL string of the query to execute</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>Extra argument that will be passed to the observer</p></td>
</tr>
<tr class="odd">
<td align="left">aObserver</td>
<td align="left"><p>Observer that will be notified when the query is done</p></td>
</tr>
</tbody>
</table>

**Returns:.**

A request object.

void nsIGktSqlConnection::beginTransaction ()
---------------------------------------------

Begin a transaction. Updates made during the transaction will not be made permanent until it is committed using commitTransaction.

void nsIGktSqlConnection::commitTransaction ()
----------------------------------------------

Commit the current transaction

void nsIGktSqlConnection::rollbackTransaction ()
------------------------------------------------

Rollback (cancel) the current transaction

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
<td align="left"><p>The schema</p></td>
</tr>
<tr class="even">
<td align="left">aTable</td>
<td align="left"><p>The table name of the keys to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The result which holds the keys.

nsIGktSqlConnectionODBC interface Reference
===========================================

Detailed Description
--------------------

The nsIGktSqlConnectionODBC interface inherits from nsIGktSqlConnection and implements a part of its methods.

In javascript, instantiating this object with the following code:

    new GktSqlConnectionODBC ()

It is the entry point to connect to an ODBC-compatible database and execute some requests.

nsIGktSqlRequest interface Reference
====================================

Public Attributes
-----------------

-   readonly attribute AString errorMessage

<!-- -->

-   readonly attribute nsIGktSqlResult result

<!-- -->

-   readonly attribute long affectedRows

<!-- -->

-   readonly attribute long lastID

<!-- -->

-   readonly attribute AString query

<!-- -->

-   readonly attribute nsISupports ctxt

<!-- -->

-   readonly attribute nsIGktSqlRequestObserver observer

<!-- -->

-   readonly attribute long status

<!-- -->

-   const long STATUS\_NONE

<!-- -->

-   const long STATUS\_EXECUTED

<!-- -->

-   const long STATUS\_COMPLETE

<!-- -->

-   const long STATUS\_ERROR

<!-- -->

-   const long STATUS\_CANCELLED

-   void cancel ( )

Detailed Description
--------------------

The nsIGktSqlRequest interface implements requests used during an asynchronous SQL query or update operation.

Member Data Documentation
-------------------------

### readonly attribute AString nsIGktSqlRequest::errorMessage

The most recent error message.

### readonly attribute nsIGktSqlResult nsIGktSqlRequest::result

The result of the operation.

### readonly attribute long nsIGktSqlRequest::affectedRows

The number of rows that were affected during an update.

### readonly attribute long nsIGktSqlRequest::lastID

The ID of the most recently added record.

### readonly attribute AString nsIGktSqlRequest::query

The SQL query.

### readonly attribute nsISupports nsIGktSqlRequest::ctxt

The context passed to the connection's asyncExecuteQuery or asyncExecuteUpdate method.

### readonly attribute nsIGktSqlRequestObserver nsIGktSqlRequest::observer

The observer that listens for when the request is finished. Methods of the observer should be called by the request.

### readonly attribute long nsIGktSqlRequest::status

The status of the request.

### const long nsIGktSqlRequest::STATUS\_NONE

Status none.

### const long nsIGktSqlRequest::STATUS\_EXECUTED

Status executed.

### const long nsIGktSqlRequest::STATUS\_COMPLETE

Status complete.

### const long nsIGktSqlRequest::STATUS\_ERROR

Status error.

### const long nsIGktSqlRequest::STATUS\_CANCELLED

Status cancelled.

void nsIGktSqlRequest::cancel ()
--------------------------------

Cancels the operation.

nsIGktSqlRequestObserver interface Reference
============================================

-   void onStartRequest ( in nsIGktSqlRequest aRequest, in nsISupports aContext)

<!-- -->

-   void onStopRequest ( in nsIGktSqlRequest aRequest, in nsISupports aContext)

Detailed Description
--------------------

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
<td align="left"><p>The request that has started</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>A context that was supplied in the query/update call</p></td>
</tr>
</tbody>
</table>

void nsIGktSqlRequestObserver::onStopRequest (in nsIGktSqlRequest aRequest, in nsISupports aContext)
----------------------------------------------------------------------------------------------------

This method will be called when the request has finished. This function will be called in both success and error cases.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRequest</td>
<td align="left"><p>The request that has started</p></td>
</tr>
<tr class="even">
<td align="left">aContext</td>
<td align="left"><p>A context that was supplied in the query/update call</p></td>
</tr>
</tbody>
</table>

nsIGktSqlResult interface Reference
===================================

Public Attributes
-----------------

-   attribute boolean displayNullAsText

<!-- -->

-   readonly attribute nsIGktSqlConnection connection

<!-- -->

-   readonly attribute AString query

<!-- -->

-   readonly attribute AString tableName

<!-- -->

-   readonly attribute long rowCount

<!-- -->

-   readonly attribute long columnCount

<!-- -->

-   const long TYPE\_STRING

<!-- -->

-   const long TYPE\_INT

<!-- -->

-   const long TYPE\_FLOAT

<!-- -->

-   const long TYPE\_DECIMAL

<!-- -->

-   const long TYPE\_DATE

<!-- -->

-   const long TYPE\_TIME

<!-- -->

-   const long TYPE\_DATETIME

<!-- -->

-   const long TYPE\_BOOL

-   AString getColumnName ( in long aColumnIndex)

<!-- -->

-   long getColumnIndex ( in AString aColumnName)

<!-- -->

-   long getColumnType ( in long aColumnIndex)

<!-- -->

-   AString getColumnTypeAsString ( in long aColumnIndex)

<!-- -->

-   long getColumnDisplaySize ( in long aColumnIndex)

<!-- -->

-   nsIGktSqlResultEnumerator enumerate ( )

<!-- -->

-   void reload ( )

Detailed Description
--------------------

The nsIGktSqlResult interface manage results of an SQL query. Use the enumerate method to retrieve each row.

Member Data Documentation
-------------------------

### attribute boolean nsIGktSqlResult::displayNullAsText

By default, this value is false.

### readonly attribute nsIGktSqlConnection nsIGktSqlResult::connection

The connection used to execute the query.

### readonly attribute AString nsIGktSqlResult::query

The SQL query.

### readonly attribute AString nsIGktSqlResult::tableName

The table that was used in the query. If more than one table was used, only the first is returned.

### readonly attribute long nsIGktSqlResult::rowCount

The number of rows in the result

### readonly attribute long nsIGktSqlResult::columnCount

The number of columns in the result

### const long nsIGktSqlResult::TYPE\_STRING

column type constants used by |getColumnType|. Column type string, used by |getColumnType|.

### const long nsIGktSqlResult::TYPE\_INT

Column type integer.

### const long nsIGktSqlResult::TYPE\_FLOAT

Column type float.

### const long nsIGktSqlResult::TYPE\_DECIMAL

Column type decimal.

### const long nsIGktSqlResult::TYPE\_DATE

Column type date.

### const long nsIGktSqlResult::TYPE\_TIME

Column type time.

### const long nsIGktSqlResult::TYPE\_DATETIME

Column type datetime.

### const long nsIGktSqlResult::TYPE\_BOOL

Column type boolean.

AString nsIGktSqlResult::getColumnName (in long aColumnIndex)
-------------------------------------------------------------

Retrieves the name of a column given its index. Indices start at zero.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The index of the column to return</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The column name.

long nsIGktSqlResult::getColumnIndex (in AString aColumnName)
-------------------------------------------------------------

Retrieves the index of a column given its name. If the column does not exist, -1 is returned.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnName</td>
<td align="left"><p>The column name to return</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The column index.

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
<td align="left"><p>The index of the column to return the type of</p></td>
</tr>
</tbody>
</table>

**Returns:.**

Te column type.

AString nsIGktSqlResult::getColumnTypeAsString (in long aColumnIndex)
---------------------------------------------------------------------

Returns the type of the data in a given column as a string. This is an alternative to using the constants and will return either string, int, float, decimal, date, time, datetime or bool.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The index of the column to return the type of</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The column type.

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
<td align="left"><p>The index of the column to return the size of</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The column size.

nsIGktSqlResultEnumerator nsIGktSqlResult::enumerate ()
-------------------------------------------------------

Returns an enumerator to enumerator over the returned rows.

**Returns:.**

The row enumerator.

void nsIGktSqlResult::reload ()
-------------------------------

Re-executes the query.

nsIGktSqlResultEnumerator interface Reference
=============================================

Public Attributes
-----------------

-   readonly attribute AString errorMessage

<!-- -->

-   readonly attribute AString currentCondition

-   boolean next ( )

<!-- -->

-   boolean previous ( )

<!-- -->

-   void beforeFirst ( )

<!-- -->

-   void first ( )

<!-- -->

-   void last ( )

<!-- -->

-   void relative ( in long aRowIndex)

<!-- -->

-   void absolute ( in long aRowIndex)

<!-- -->

-   boolean isNull ( in long aColumnIndex)

<!-- -->

-   nsIVariant getVariant ( in long aColumnIndex)

<!-- -->

-   AString getString ( in long aColumnIndex)

<!-- -->

-   long getInt ( in long aColumnIndex)

<!-- -->

-   float getFloat ( in long aColumnIndex)

<!-- -->

-   float getDecimal ( in long aColumnIndex)

<!-- -->

-   long long getDate ( in long aColumnIndex)

<!-- -->

-   boolean getBool ( in long aColumnIndex)

Detailed Description
--------------------

The nsIGktSqlResultEnumerator interface is used to get the results from an SQL query. The enumerator uses a row pointer which can be adjusted with the next and previous methods. Other methods operate only on the row selected by the pointer.

The row pointer starts just before the first row, so you should always call the next method once before attempting to read row data.

Member Data Documentation
-------------------------

### readonly attribute AString nsIGktSqlResultEnumerator::errorMessage

The most recent error message.

### readonly attribute AString nsIGktSqlResultEnumerator::currentCondition

Holds the SQL condition clause.

boolean nsIGktSqlResultEnumerator::next ()
------------------------------------------

Moves the row pointer to the next row in the results. Returns true if there is a next row and false if there are no more rows.

**Returns:.**

False if there are no more rows.

boolean nsIGktSqlResultEnumerator::previous ()
----------------------------------------------

Moves the row pointer to the previous row in the results. Returns true if there is a previous row.

**Returns:.**

False if there are no previous rows.

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

Moves the row pointer by a number relative to the current row. An error occurs if this causes the row pointer to extend past the last row. This method may also be used to move the row pointer back by using a negative value.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRowIndex</td>
<td align="left"><p>The number of rows to move by</p></td>
</tr>
</tbody>
</table>

void nsIGktSqlResultEnumerator::absolute (in long aRowIndex)
------------------------------------------------------------

Moves the row pointer to a specific row. An error occurs if the index is after the last row.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aRowIndex</td>
<td align="left"><p>The index of the row to move to</p></td>
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
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

True if the value is null.

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
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The value as a variant.

AString nsIGktSqlResultEnumerator::getString (in long aColumnIndex)
-------------------------------------------------------------------

Returns the value at the specified column in the current row as a string. An error occurs if the value is not a string type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The string value.

long nsIGktSqlResultEnumerator::getInt (in long aColumnIndex)
-------------------------------------------------------------

Returns the value at the specified column in the current row as an integer. An error occurs if the value is not an integer type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The integer value.

float nsIGktSqlResultEnumerator::getFloat (in long aColumnIndex)
----------------------------------------------------------------

Returns the value at the specified column in the current row as a float. An error occurs if the value is not a float type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The float value.

float nsIGktSqlResultEnumerator::getDecimal (in long aColumnIndex)
------------------------------------------------------------------

Returns the value at the specified column in the current row as a decimal. An error occurs if the value is not a decimal type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The decimal value.

long long nsIGktSqlResultEnumerator::getDate (in long aColumnIndex)
-------------------------------------------------------------------

Returns the value at the specified column in the current row as a date. An error occurs if the value is not a date type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The date value.

boolean nsIGktSqlResultEnumerator::getBool (in long aColumnIndex)
-----------------------------------------------------------------

Returns the value at the specified column in the current row as a boolean. An error occurs if the value is not a boolean type.

<table>
<caption>Parameters</caption>
<colgroup>
<col width="20%" />
<col width="80%" />
</colgroup>
<tbody>
<tr class="odd">
<td align="left">aColumnIndex</td>
<td align="left"><p>The column to retrieve</p></td>
</tr>
</tbody>
</table>

**Returns:.**

The boolean value.

nsIGktSqlResultODBC interface Reference
=======================================

Detailed Description
--------------------

The nsIGktSqlResultODBC interface manage the results of an SQL query for ODBC databases. Implements some nsIGtkSqlResult methods and attributes.
