// -----------------------------------------
// Relay demo module
// -----------------------------------------

// IMPORTANT: Dependence on app_request.js

// A relay request is a JSON request to server, JSON like :
// {
// 	method: 'POST' or 'GET'
// 	host: host adress, special case to read a file on file system : set host to 'fs' and url to file path + file name
// 	url: url to get
// 	status: start at 0, and set by relay server according to comptipix answer
// 	content: content of the answer, + used on POST request to set POST content
// 	content_type: used to set content type, it's 'application/octet-stream' for POST authentication, and
// }

// NOTE: you probably don't need to use low level JSON request : here is 4 helper functions tha fit many need :
// - parse_result: 							--> use it on every request to help get req status (relay + comptipix) and automaticaly parse comptipix parameters
// - get_cpx_tkn: to get a token			--> app_relay.get_cpx_tkn(f_callback, host, user, pass);   // if omited user='reader' and pass='reader'
// - get_cpx_param: to get parameters	--> app_relay.get_cpx_param(f_callback, host, tkn, 'List&Of&comptipix&parameters&separated&by&commercial&and');
// - fs_read: to do a request to read a file on relay server file system

/*!
	\brief	Module containing message utility function

*/
var app_relay = (function()
{
	'use strict';// we are strict   (see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions_and_function_scope/Strict_mode?redirectlocale=en-US&redirectslug=JavaScript%2FReference%2FFunctions_and_function_scope%2FStrict_mode)

	// ---------------------
	// Public method
	// ---------------------

	/*!
		\brief	Helper function to check and parse a request result

		\note 	return object is like :
					{
						ok: 			// true if request to relay status == 200 and relay request to comptipix == 200, false is one or the other is not 200
						ok_relay: 	// true if request to relay status == 200
						ok_cpx: 		// true if request from relay to comptipix status == 200
						req: 			// it is the full req object returned by relay server like :
						{
							method: 'POST' or 'GET'  								// if you use app_relay.get_cpx_tkn() + app_relay.get_cpx_param(), setting post/get is automatic ('POST' for auth, 'GET' for all other)
							host: host adress	or 'fs-read' or 'fs-write'		// - host: like: '192.168.0.139' or 'http://192.168.0.139/' 	--> don't mind 'http://' relay server will add it for you
																							// - fs-read: the file to read from server
																							// - fs-write: the file to write on server
																							// - fs-write-append: the content to add to file on server
							url: url,					 								// url to get like 'CONFIG?tkn=TKN&occupancy_state' 				--> if you use app_relay.get_cpx_tkn() + app_relay.get_cpx_param(), don't need to set token
							status: 200,				 								// this is the status of request from relay to comptipix (if not 200 you have a problem ... may be 503 can happen on sdcard read)
							content: answer content, 								// this is your answer content, you have to read it only for file or token, otherwise if you use parse_result(), parameters are alredy parsed for you in 'param'
							content_type: 				 								// don't need to set this is if you use app_relay.get_cpx_tkn(), that will set 'application/octet-stream' for authentification and 'application/x-www-form-urlencoded' for all other
						}
						param: 		// parameter parsed in object + array --> like : answer[param_1_name] = [param_1_value_1, param_1_value_2]
					}

		\param 	result 	xhr object
		\return 	object parsed
	*/
	var parse_result = function(result)
	{
		var tmp = {
			ok: false,			// will be okay if relay + comptipix are ok
			ok_relay: false,	// ok if relay is ok
			ok_cpx: false,		// ok if comptipix is ok
			req: null,			// contain full request
			param: null			// contain parsed comptipix parameters
		};

		if (200 == result.status)
		{
			tmp.ok_relay = true; // Request to relay server ok
			tmp.req = JSON.parse(result.response);
			if ('' !== tmp.req.content)
				tmp.param = app_relay.read_cpx_answer(tmp.req.content);
			if (200 == tmp.req.status)
				tmp.ok_cpx = true;// Request from relay server to Comptipix ok

			if ((tmp.ok_cpx)&&(tmp.ok_relay))
				tmp.ok = true;
		}

		return tmp;
	};

	/*!
		\brief	Do a POST request to a comptipix to get a token

		\param 	old_mode			true to enable Concentrix/iComptipix compatibility mode
		\param	f_call			callback function
		\param	target_host		host target ('192.168.100')
		\param	user				user to use to connect											(DEFAULT 'reader')
		\param	pass				password to use to connect										(DEFAULT 'reader')
		\param	f_opt_1			OPTIONNAL parameter 1 passed to callback function (as parameter 2, parameter 1 is xhr result object)
		\param	f_opt_2			OPTIONNAL parameter 2 passed to callback function (as parameter 3, parameter 1 is xhr result object)
		\param	req_timeout		OPTIONNAL used to change default timeout in ms			(DEFAULT: app_request_config.timeout_ms)
		\param	f_onerror		OPTIONNAL fonction to override default onerror			(DEFAULT: app_request_config.f_onerror)
		\param	f_ontimeout		OPTIONNAL fonction to override default timeout 			(DEFAULT: app_request_config.f_ontimeout)
	*/
	var _get_cpx_tkn = function(old_mode, f_call, target_host, user, pass, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout)
	{
		// Apply default parameters
		if ('undefined' === typeof user)
			user = 'reader';
		if ('undefined' === typeof pass)
			pass = 'reader';

		var content = 'user=' + user + '&pass=' + pass;
		var content_type = 'application/octet-stream';// for post content we have to use octet-stream (because we use post only on login page)

		// build request to read file
		var tmp_r = get_req_obj(target_host, 'CONFIG?get_tkn', 'POST', content, content_type);
		if (old_mode)
			tmp_r = get_req_obj(target_host, '', 'POST', content, content_type);

		// relay request
		app_request.post('/api', JSON.stringify(tmp_r), f_call, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Do a POST request to a comptipix to get a token

		\params ... see _get_cpx_tkn()
	*/
	var get_cpx_tkn = function(f_call, target_host, user, pass, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout)
	{
		_get_cpx_tkn(false, f_call, target_host, user, pass, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Do a POST request to a comptipix to get a token

		\params ... see _get_cpx_tkn()
	*/
	var get_cpx_tkn_old = function(f_call, target_host, user, pass, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout)
	{
		_get_cpx_tkn(true, f_call, target_host, user, pass, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Do a POST request to get comptipix parameters

		\param	f_call			callback function
		\param	target_host		host target ('192.168.100')
		\param 	tkn 				token to use
		\param	req_param		request parameter (like 'occupancy_state&occupancy_config')
		\param	f_opt_1			OPTIONNAL parameter 1 passed to callback function (as parameter 2, parameter 1 is xhr result object)
		\param	f_opt_2			OPTIONNAL parameter 2 passed to callback function (as parameter 3, parameter 1 is xhr result object)
		\param	req_timeout		OPTIONNAL used to change default timeout in ms			(DEFAULT: app_request_config.timeout_ms)
		\param	f_onerror		OPTIONNAL fonction to override default onerror			(DEFAULT: app_request_config.f_onerror)
		\param	f_ontimeout		OPTIONNAL fonction to override default timeout 			(DEFAULT: app_request_config.f_ontimeout)
	*/
	var get_cpx_param = function(f_call, target_host, tkn, req_param, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout)
	{
		var req_url = 'CONFIG?tkn=' + tkn + '&' + req_param;
		if ('' == tkn) // old compatibility mode
			req_url = req_param;
		var tmp_r = get_req_obj(target_host, req_url, 'GET');

		// console.log('get_cpx_param');
		// console.log(tmp_r);

		if (null === tkn)
		{
			console.warn('token null passed to get_cpx_param('+req_param+') ABORT');
			return;
		}

		// relay request
		app_request.post('/api', JSON.stringify(tmp_r), f_call, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Do a POST request to get comptipix image (base64 encoded response)

		\param	f_call			callback function
		\param	target_host		host target ('192.168.100')
		\param 	tkn 				token to use
		\param	req_param		request image parameter (like 'sensor&full&jpeg')		(DEFAULT: 'sensor&full&jpeg')
		\param 	image_name 		image file name to get											(DEFAULT: 'sensor.jpg')
		\param 	image_type 		image MIME type to get 											(DEFAULT: 'image/jpeg')
		\param	f_opt_1			OPTIONNAL parameter 1 passed to callback function (as parameter 2, parameter 1 is xhr result object)
		\param	f_opt_2			OPTIONNAL parameter 2 passed to callback function (as parameter 3, parameter 1 is xhr result object)
		\param	req_timeout		OPTIONNAL used to change default timeout in ms			(DEFAULT: app_request_config.timeout_ms)
		\param	f_onerror		OPTIONNAL fonction to override default onerror			(DEFAULT: app_request_config.f_onerror)
		\param	f_ontimeout		OPTIONNAL fonction to override default timeout 			(DEFAULT: app_request_config.f_ontimeout)
	*/
	var get_cpx_image = function(f_call, target_host, tkn, req_param, image_name, image_type, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout)
	{
		// Apply default parameters
		if ('undefined' === typeof req_param)
			req_param = 'sensor&full&jpeg';
		if ('undefined' === typeof image_name)
			image_name = 'sensor.jpg';
		if ('undefined' === typeof image_type)
			image_type = 'image/jpeg';

		var req_url = 'IMAGE/' + image_name + '?tkn=' + tkn + '&' + req_param + '&jpeg=' + Math.random(); // Add a random to bypass browser cache (need to be somewhere inside request, so here is good)
		if ('' == tkn) // old compatibility mode
			req_url = req_param;

		var tmp_r = get_req_obj(target_host, req_url, 'GET', '', image_type);

		// console.log('get_cpx_image');
		// console.log(tmp_r);

		if (null === tkn)
		{
			console.warn('token null passed to get_cpx_image('+req_param+') ABORT');
			return;
		}

		// relay request
		app_request.post('/api', JSON.stringify(tmp_r), f_call, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Do a request to read fs file

		\param	f_call			callback function
		\param	file_name		file name to read
		\param	f_opt_1			OPTIONNAL parameter 1 passed to callback function (as parameter 2, parameter 1 is xhr result object)
		\param	f_opt_2			OPTIONNAL parameter 2 passed to callback function (as parameter 3, parameter 1 is xhr result object)
		\param	req_timeout		OPTIONNAL used to change default timeout in ms			(DEFAULT: app_request_config.timeout_ms)
		\param	f_onerror		OPTIONNAL fonction to override default onerror			(DEFAULT: app_request_config.f_onerror)
		\param	f_ontimeout		OPTIONNAL fonction to override default timeout 			(DEFAULT: app_request_config.f_ontimeout)
	*/
	var fs_read = function(f_call, file_name, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout)
	{
		// build request to read file
		var tmp_r = get_req_obj('fs-read', file_name);

		// read file request
		app_request.post('/api', JSON.stringify(tmp_r), f_call, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout);
	};

	/*!
		\brief	Process comptipix answer to get an object of parameter like answer[param_1_name] = [param_1_value_1, param_1_value_2]
					Or return answer[param_1_name] = param_1_value_1, param_1_value_2   if do_array is set to false

		\note 	This function is low level you probably don't need to use it directly (use app_relay.parse_result() that will use this one)

		\note		input parameters must be separated by '\n'
					input value will be after '='

		\param	answer 					a comptipix server answer
		\param	do_array 				true to return an answer array, false otherwise 	(DEFAULT: true)
		\param	filter_save_current	true to filter the saved + current answer 			(DEFAULT: true)
		\return	an array of pair param_name -> param_value
	*/
	var read_cpx_answer = function(answer, do_array, filter_save_current)
	{
		if ('undefined' === typeof do_array)
			do_array = true;
		if ('undefined' === typeof filter_save_current)
			filter_save_current = true;

		var tmp_obj = {};
		if (null !== answer)
		{
			var requete_reponse = '';
			if ('string' != typeof content)	// because number can't be splitted and javascript is so stupid that it throw blocking error if we try splitting number
				requete_reponse = answer.toString().split('\n');
			else
				requete_reponse = answer.split('\n');

			for (var index in requete_reponse)
			{
				// filter unwanted 'save=' + 'current='
				if ( 	(true === filter_save_current) &&
						(-1 == requete_reponse[index].indexOf('save=')) &&
						(-1 == requete_reponse[index].indexOf('current='))
					)
				{
					var tmp = requete_reponse[index].indexOf('=');
					if (tmp >= 0)
					{
						var tmp_name=requete_reponse[index].substring(0,tmp);
						var tmp_value=requete_reponse[index].substring(tmp+1);

						if ( (-1 == tmp_value.indexOf(',')) || (!do_array) )
							tmp_obj[tmp_name] = [tmp_value];
						else
							tmp_obj[tmp_name] = tmp_value.split(',');
					}
				}
			}
		}
		else
			console.warn('app_config.get_answer_obj() --> answer is null !!!');

		//DEBUG:
		//console.log('app_config.get_answer_obj -->');
		//console.table(tmp_obj);
		//console.log(tmp_obj);

		return tmp_obj;
	};

	/*!
		\brief	Get a relay request object

		\note 	This function is low level you probably don't need to use it directly (use app_relay.get_cpx_param() that will use this one)

		\param	p_host			host target ('192.168.100' or 'fs' to read a file on relay server)										(DEFAULT: "")
		\param	p_url				url to use (like 'CONFIG?tkn=myTKN&occupancy_state' or file_name to read file)						(DEFAULT: "")
		\param	p_method			method "POST" or "GET"																									(DEFAULT: "GET")
		\param	p_content		content, used on POST request to set post content, and used to receive response content			(DEFAULT: "")
		\param	p_content_type	used to set a content-type header on relay request																(DEFAULT: "application/x-www-form-urlencoded")
	*/
	var get_req_obj = function(p_host, p_url, p_method, p_content, p_content_type)
	{
		// Apply default parameters
		if ('undefined' === typeof p_host)
			p_host = '';
		if ('undefined' === typeof p_url)
			p_url = '';
		if ('undefined' === typeof p_method)
			p_method = 'GET';
		if ('undefined' === typeof p_content)
			p_content = '';
		if ('undefined' === typeof p_content_type)
			p_content_type = 'application/x-www-form-urlencoded';

		// Return request object
		return {
			method: p_method,
			host: p_host,
			url: p_url,
			status: 0,// always start status at 0
			content: p_content,
			content_type: p_content_type
		};
	};

	/*!
		\brief	Do a POST request relay

		\note 	This function is low level you probably don't need to use it directly (use app_relay.get_cpx_param())

		\param	f_call			callback function
		\param	target_host		host target ('192.168.100')
		\param	target_url		url to use (like 'CONFIG?tkn=myTKN&occupancy_state')
		\param	method			method "POST" or "GET"
		\param	content			content, used on POST request to set post content, and used to receive response content
		\param	content_type	used to set a content-type header on relay request
		\param	f_opt_1			OPTIONNAL parameter 1 passed to callback function (as parameter 2, parameter 1 is xhr result object)
		\param	f_opt_2			OPTIONNAL parameter 2 passed to callback function (as parameter 3, parameter 1 is xhr result object)
		\param	req_timeout		OPTIONNAL used to change default timeout in ms			(DEFAULT: app_request_config.timeout_ms)
		\param	f_onerror		OPTIONNAL fonction to override default onerror			(DEFAULT: app_request_config.f_onerror)
		\param	f_ontimeout		OPTIONNAL fonction to override default timeout 			(DEFAULT: app_request_config.f_ontimeout)
	*/
	var relay_req = function(f_call, target_host, target_url, method, content, content_type, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout)
	{
		// build request to read file
		var tmp_r = get_req_obj(target_host, target_url, method, content, content_type);

		// relay request
		app_request.post('/api', JSON.stringify(tmp_r), f_call, f_opt_1, f_opt_2, req_timeout, f_onerror, f_ontimeout);
	};

	// ---------------------
	// Return public method
	// ---------------------

	return {
		parse_result:		parse_result,
		get_cpx_tkn:		get_cpx_tkn,
		get_cpx_tkn_old:	get_cpx_tkn_old,
		get_cpx_param:		get_cpx_param,
		get_cpx_image:		get_cpx_image,
		fs_read:				fs_read,
		read_cpx_answer:	read_cpx_answer,	// NOTE: low level function to parse a comptipix answer   			--> not used directly in demo.js !
		get_req_obj:		get_req_obj,		// NOTE: low level function to get an object for a relay request  --> not used directly in demo.js !
		relay_req:			relay_req			// NOTE: low level function to do a relay request						--> not used directly in demo.js !
	};
})();
