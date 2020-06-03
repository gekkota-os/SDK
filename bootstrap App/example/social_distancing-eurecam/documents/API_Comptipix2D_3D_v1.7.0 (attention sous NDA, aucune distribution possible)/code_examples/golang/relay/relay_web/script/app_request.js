// This module is a simple layer to manage xhr
// It make also xhr easy for IE6-7 using old ActiveX object



// -----------------------------------------
// Ajax Config object
// -----------------------------------------

var app_request_config = app_request_config || {	// the "||" trick ensure that this object is a singleton

	// Default timeout timer (XHR default is 30s ... this is very loooong)
	timeout_ms: 15000,

	f_onerror: null,
	f_ontimeout: null,

	/*!
		\brief	On error default function

		\note		This error occurs when there is error on the network level, otherwise status code is sent

		\param	req	the request object
	*/
	onerror: function(req)
	{
		console.error('request_make --> req('+req.responseURL+') : ERR ! ! !');	// this happen on network error --> report it to user
		if (app_request_config.f_onerror instanceof Function)// NOTE: don't want to depend on app_utils here
			app_request_config.f_onerror(req);
	},

	/*!
		\brief	On timeout default function

		\note		This error occurs when application has no response and no info from server

		\param	req	the request object
	*/
	ontimeout: function(req)
	{
		console.error('request_make --> req('+req.responseURL+') : TIMEOUT ! ! !'); // this happen on network error --> report it to user and abort request
		if (app_request_config.f_ontimeout instanceof Function)// NOTE: don't want to depend on app_utils here
			app_request_config.f_ontimeout(req);
	},
};


// -----------------------------------------
// Ajax Functions Module
// -----------------------------------------

/*!
	\brief	Module containing all functions to do AJAX

*/
var app_request = (function()
{
	'use strict';// we are strict (see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions_and_function_scope/Strict_mode?redirectlocale=en-US&redirectslug=JavaScript%2FReference%2FFunctions_and_function_scope%2FStrict_mode)

	// ---------------------
	// Private methods
	// ---------------------

	// special code to manage old browser (back to IE6)
	// ---------------
	var zXml= {
		useActiveX:(typeof ActiveXObject!='undefined'),
		useDom:document.implementation&&document.implementation.createDocument,
		useXmlHttp:(typeof XMLHttpRequest!='undefined')
	};

	zXml.ARR_XMLHTTP_VERS = ['MSXML2.XmlHttp.6.0','MSXML2.XmlHttp.3.0'];
	zXml.ARR_DOM_VERS = ['MSXML2.DOMDocument.6.0','MSXML2.DOMDocument.3.0'];

	var zXmlHttp = function() {};

	zXmlHttp.createRequest = function()
	{
		if (zXml.useXmlHttp)
			return new XMLHttpRequest();	// for a recent browser we simply use a XMLHttpRequest object
		else if (zXml.useActiveX)
		{
			// for an old browser we see what we can do ... with activeX object, yuk!
			if (!zXml.XMLHTTP_VER)
			{
				for (var i=0, len=zXml.ARR_XMLHTTP_VERS.length; i<len; i++)
				{
					try
					{
						new ActiveXObject(zXml.ARR_XMLHTTP_VERS[i]);
						zXml.XMLHTTP_VER=zXml.ARR_XMLHTTP_VERS[i];
						break;
					}
					catch (oError) {}
				}
			}

			if (zXml.XMLHTTP_VER)
				return new ActiveXObject(zXml.XMLHTTP_VER);
			else
				throw new Error('Could not create XML HTTP Request');
		}
		else
			throw new Error('Your browser doesnt support an XML HTTP Request');
	};

	zXmlHttp.isSupported = function()
	{
		return zXml.useXmlHttp||zXml.useActiveX;
	};


	// AJAX syntaxic sugar
	// ---------------

	/*!
		\brief	Convenient function to make AJAX request

		\note		We avoid synchronous request because it obstruct the main thread
		\note 	If you need more than 2 parameters : use an array as param 1, and deal with it in callback ; or just use init

		\param	req_type		request type (can be 'GET' or 'POST')
		\param	req_url		request url
		\param	f_call		request callback function --> parameter 0 of callback is request object
		\param	f_param_1	request callback function parameter 1
		\param	f_param_2	request callback function parameter 2
		\param	post_data	optionnal post_data
		\param	req_timeout	used to change default timeout
		\param	f_onerror	fonction to override default onerror
		\param	f_ontimeout	fonction to override default timeout
	*/
	var _make = function(req_type, req_url, f_call, f_param_1, f_param_2, post_data, req_timeout, f_onerror, f_ontimeout)
	{
		var req = zXmlHttp.createRequest();
		req.open(req_type, req_url, true);
		if ('GET' == req_type)
			req.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');// indicate that content is URI encoded
		else
			req.setRequestHeader('Content-type', 'application/octet-stream');// for post content we have to use octet-stream (because we use post only on login page)

		req.timeout = app_request_config.timeout_ms;
		if ( (null !== req_timeout) && ('undefined' !== typeof req_timeout) )
			req.timeout = req_timeout;

		// action to do when request is done and okay
		if ('undefined' != typeof req.onload)
		{
			// Use it preferably (works better on some android browser) and is more efficient on new browser
			req.onload = function()
			{
				if (f_call instanceof Function)// NOTE: don't want to depend on app_utils here
					f_call(req, f_param_1, f_param_2);
			};
		}
		else
		{
			// Old browser compatibility path
			req.onreadystatechange = function () {
				if (4 === req.readyState)
				{
					if (f_call instanceof Function)// NOTE: don't want to depend on app_utils here
						f_call(req, f_param_1, f_param_2);
				}
			};

			// console.warn('Old bro request using onreadystatechange ! ! ! ! ');
		}

		// action to do when request is done but error
		req.onerror = function()
		{
			if (f_onerror instanceof Function)// NOTE: don't want to depend on app_utils here
				f_onerror(req, f_param_1, f_param_2);
			else
				app_request_config.onerror(req);
		};

		// action to do when request reach timeout
		req.ontimeout = function()
		{
			if (f_ontimeout instanceof Function)// NOTE: don't want to depend on app_utils here
				f_ontimeout(req, f_param_1, f_param_2);
			else
				app_request_config.ontimeout(req);
		};

		// send the request
		if ('undefined' === typeof post_data)
			post_data = null;
		req.send(post_data);
	};

	/*!
		\brief	Convenient function to make AJAX type 'GET' request

		\note		We avoid synchronous request because it obstruct the main thread

		\param	req_url		request url
		\param	f_call		request callback function --> parameter 0 of callback is responseText result
		\param	f_param_1	request callback function parameter 1
		\param	f_param_2	request callback function parameter 2
		\param	req_timeout	used to change default timeout in ms			(DEFAULT: app_request_config.timeout_ms)
		\param	f_onerror	fonction to override default onerror			(DEFAULT: app_request_config.f_onerror)
		\param	f_ontimeout	fonction to override default timeout 			(DEFAULT: app_request_config.f_ontimeout)
	*/
	var get = function(req_url, f_call, f_param_1, f_param_2, req_timeout, f_onerror, f_ontimeout)
	{
		_make('GET', req_url, f_call, f_param_1, f_param_2, null, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Convenient function to make AJAX type 'POST' request

		\note		We avoid synchronous request because it obstruct the main thread

		\param	req_url		request url
		\param	post_data	post data
		\param	f_call		request callback function --> parameter 0 of callback is responseText result
		\param	f_param_1	request callback function parameter 1
		\param	f_param_2	request callback function parameter 2
		\param	req_timeout	used to change default timeout in ms			(DEFAULT: app_request_config.timeout_ms)
		\param	f_onerror	fonction to override default onerror			(DEFAULT: app_request_config.f_onerror)
		\param	f_ontimeout	fonction to override default timeout 			(DEFAULT: app_request_config.f_ontimeout)
	*/
	var post = function(req_url, post_data, f_call, f_param_1, f_param_2, req_timeout, f_onerror, f_ontimeout)
	{
		_make('POST', req_url, f_call, f_param_1, f_param_2, post_data, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Preapare an XMLHttpRequest asynchronous

		\note		We avoid synchronous request because it obstruct the main thread

		\param	req_url					request url
		\param	req_type					request type (can be 'GET' or 'POST'), if not set 	(DEFAULT: 'GET')
		\param	default_onerror		true to add default onerror function					(DEFAULT: true)
		\param	default_ontimeout		true to add default ontimeout function					(DEFAULT: true)
	*/
	var init = function(req_url, req_type, default_onerror, default_ontimeout)
	{
		// Define request type : POST or GET (GET by default)
		var request_type = 'GET';
		if ('undefined' != typeof req_type)
			request_type = req_type;
		if ('undefined' != typeof default_onerror)
			default_onerror = true;
		if ('undefined' != typeof default_ontimeout)
			default_ontimeout = true;

		// Create the request
		var req = zXmlHttp.createRequest(); // NOTE: this is equivalent to new XMLHttpRequest();
		req.open(request_type, req_url, true); // method , url , async=true  -->  async must be always true  (synchronous request are deprecated by all browser!!)

		// Add default onerror / ontimeout function
		if (true === default_onerror)
		{
			req.onerror = function()
			{
				app_request_config.onerror(req);
			};
		}
		if (true === default_ontimeout)
		{
			req.ontimeout = function()
			{
				app_request_config.ontimeout(req);
			};
		}

		// Return the request object
		// --> and now it's up to the caller to :
		//		- add the success function with req.onload (which is simplier for attaching callback with parameter, see http://www.webdeveloper.com/forum/showthread.php?109285-Ajax-passing-an-argument-to-the-callback-function)
		//		- set req.responseType to 'arraybuffer' or whatever he want (or do not change it for a simple text response)
		//		- addEventListener to monitor progress or whaterver (see https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/Using_XMLHttpRequest#Monitoring_progress)
		//		- add other function see https://developer.mozilla.org/en-US/docs/Mozilla/Tech/XPCOM/Reference/Interface/nsIXMLHttpRequestEventTarget
		//		- do the req.send();  -->  MANDATORY !!!
		return req;
	};

	// ---------------------
	// Return public method
	// ---------------------

	return {
		get: 			get,
		post: 		post,
		init: 		init
	};
})();
