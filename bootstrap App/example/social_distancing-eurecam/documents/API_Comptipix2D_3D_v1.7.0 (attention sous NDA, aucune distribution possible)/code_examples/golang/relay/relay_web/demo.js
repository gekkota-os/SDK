
// reference object
var app = {
	current_demo: 2,
	demo_started: false,
	user: 'reader',
	pass: 'reader',
	pass_err: false,
	file_path: '',
	old_mode: false, // true to enable old API mode : Concentrix/iComptipix/Affix mode
	hide_buttons: 0,
	hide_mouse: 0,
	hide_eurecam: 0,
	hide_hour: 0,
	tkn: null,
	host: null,
	gauge: null,
	chart: null,
	chart_file_results: [0,0],
	chart_file_contents: [0,0],
	chart_file_length: [0,0],
	occ_conf_forced:{
		forced: false,
		max: null,
		min: null,
		thres_max: null,
		thres_min: null
	},
	callback_in_progress:{
		update_gauge: 0,
		update_gauge_err: 0,
		update_gauge_server: 0,
		update_gauge_server_err: 0,
		update_chart: 0,
		update_chart_err: 0,
		update_counting: 0,
		update_counting_err: 0
	},
	intervals:{
		gauge:null,
		chart:null,
		counting:null,
		gauge_server:null,
		tkn:null,
		hour:null
	}
};

// definition of occupancy_config_0 bitmask parameters
var occ_option = {
	CONFIG_ENABLE:				0x01,
	CONFIG_CLIP_MIN:			0x02,
	CONFIG_CLIP_MAX:			0x04,
	CONFIG_THRESHOLD_MAX:	0x10,
	CONFIG_THRESHOLD_MIN:	0x20
};



// -----------------------------------------
// Relay demo module
// -----------------------------------------

/*!
	\brief	Module containing message utility function

*/
var demo = (function()
{
	// ---------------------
	// Private method
	// ---------------------

	/*!
		\brief	Destroy the current app gauge

	*/
	var destroy_gauge = function()
	{
		// Not necessary anymore ... if you use the new justgage

		if (null !== app.gauge)
		{
			console.log('DESTROY gauge');
			app.gauge.destroy();
			app.gauge = null;
		}
	};

	/*!
		\brief	Transform an old occupancy value to a new one

		\param	values_obj		object containing at least answer from this params : 'affix_min&affix_min_ok&affix_max&affix_max_ok&affix_seuil_mode&affix_seuil&affix_raz'
		\return	occupancy_config
	*/
	var old_occupancy_config_to_new = function(old_values)
	{
		var tmp = [];
		tmp[0] = occ_option.CONFIG_ENABLE; // in Concentrix/Affix/iComptipix occupancy computing is always enabled
		if ('undefined' !== typeof old_values.affix_min_ok)
		{
			if (parseInt(old_values.affix_min_ok, 10) > 0)
				tmp[0] ^= occ_option.CONFIG_CLIP_MIN;
		}
		if ('undefined' !== typeof old_values.affix_max_ok)
		{
			if (parseInt(old_values.affix_max_ok, 10) > 0)
				tmp[0] ^= occ_option.CONFIG_CLIP_MAX;
		}
		if ('undefined' !== typeof old_values.affix_seuil_mode)
		{
			if (1 == parseInt(old_values.affix_seuil_mode, 10))
				tmp[0] ^= occ_option.CONFIG_THRESHOLD_MAX;
			if (2 == parseInt(old_values.affix_seuil_mode, 10))
				tmp[0] ^= occ_option.CONFIG_THRESHOLD_MIN;
		}
		if ('undefined' !== typeof old_values.affix_min)
			tmp[1] = parseInt(old_values.affix_min[0], 10);
		if ('undefined' !== typeof old_values.affix_max)
			tmp[2] = parseInt(old_values.affix_max[0], 10);
		if ('undefined' !== typeof old_values.affix_seuil)
			tmp[3] = parseInt(old_values.affix_seuil[0], 10);
		if ('undefined' !== typeof old_values.affix_raz)
			tmp[4] = parseInt(old_values.affix_raz[0], 10);

		return tmp;
	};

	/*!
		\brief	Transform an old occupancy value to a new one

		\param	values_obj		object containing at least answer from this params : 'affix_presence&affix_cumul_e&affix_cumul_s'
		\return	occupancy_config
	*/
	var old_occupancy_state_to_new = function(old_values)
	{
		var tmp = [];
		tmp[0] = occ_option.CONFIG_ENABLE; // in Concentrix/Affix/iComptipix occupancy computing is always enabled
		if ('undefined' !== typeof old_values.affix_presence)
			tmp[1] = parseInt(old_values.affix_presence[0], 10);
		if ('undefined' !== typeof old_values.affix_cumul_e)
			tmp[2] = parseInt(old_values.affix_cumul_e[0], 10);
		if ('undefined' !== typeof old_values.affix_cumul_s)
			tmp[3] = parseInt(old_values.affix_cumul_s[0], 10);

		return tmp;
	};

	/*!
		\brief	Auto adapt user and pass input for old mode (if user/pass is reader/reader it will be transformed to user/user)

		\param	values_obj		object containing at least answer from this params : 'affix_presence&affix_cumul_e&affix_cumul_s'
		\return	occupancy_config
	*/
	var adapt_user_pass_old_mode = function()
	{
		if (document.getElementById('old_mode').checked)
		{
			if ('reader' == document.getElementById('user').value)
				document.getElementById('user').value = 'user';
			if ('reader' == document.getElementById('pass').value)
				document.getElementById('pass').value = 'user';
		}
		else
		{
			if ('user' == document.getElementById('user').value)
				document.getElementById('user').value = 'reader';
			if ('user' == document.getElementById('pass').value)
				document.getElementById('pass').value = 'reader';
		}
	};

	/*!
		\brief	Add padding to a value

		\param	value		value to add padding
		\param	pad		padding to add
		\return	value with padding
	*/
	var tostring_pad = function(value, pad)
	{
		var tmp = value.toString();
		while(tmp.length<pad)
			tmp = '0'+tmp;

		return tmp;
	};

	/*!
		\brief	Transform '2017-02-08' in '20170208'

		\param	iso_date 		date iso to transform
	*/
	var get_date_iso_no_hyphen = function(iso_date)
	{
		var re = new RegExp('-', 'g');
		return iso_date.toString().replace(re, '');
	};

	/*!
		\brief	Get iso string like '2017-02-08' from a date object

		\param	date_obj 	date object			DEFAULT: now
	*/
	var get_date_iso_str = function(date_obj)
	{
		// Apply default parameter
		if ('undefined' === typeof date_obj)
			date_obj = new Date();
		return date_obj.getFullYear().toString() + '-' + tostring_pad(date_obj.getMonth()+1, 2) + '-' +  tostring_pad(date_obj.getDate(), 2);
	};

	/*!
		\brief	Check and update callback error

		\param	callback_name 	callback_name to check
		\return true if action can continue, false otherwise
	*/
	var is_callback_ok = function(callback_name)
	{
		if (app.callback_in_progress[callback_name] > 0)
		{
			app.callback_in_progress[callback_name+'_err']++;
			if (app.callback_in_progress[callback_name+'_err'] > 5)
			{
				app.callback_in_progress[callback_name] = 0;				// Reset in progress callbacks
				app.callback_in_progress[callback_name+'_err'] = 0;	// Reset error
			}
			return false;
		}
		else
			app.callback_in_progress[callback_name+'_err'] = 0;		// Reset error (because we want consecutive error)

		return true;
	};

	/*!
		\brief	On read local file callback

		\param 	result 	xhr object
	*/
	var fs_read_callback = function(result)
	{
		// Just display file
		var tmp = app_relay.parse_result(result);
		document.getElementById('file_read').innerHTML = tmp.req.content.replace(/\n/g, '<br>').replace(/ /g, '&nbsp;'); // Print file in html
		document.getElementById('file_read_name').value = tmp.req.url; // print file name
	};

	/*!
		\brief	Get value to update a gauge

		\param 	f_call 	function callback
		\param 	host 		target host
		\param 	tkn 		token to use
	*/
	var update_gauge = function(f_call, host, tkn)
	{
		if (2 != app.current_demo)//HACK
			return;//HACK

		// Check all callback is ok
		if (!is_callback_ok('update_gauge'))
		{
			console.warn('Too many callback waiting ('+app.callback_in_progress.update_gauge+'/'+app.callback_in_progress.update_gauge_err+') for update_gauge');
			return;
		}

		// Use app tkn if undefined
		if ('undefined' === typeof tkn)
			tkn = app.tkn;
		if ('undefined' === typeof host)
			host = app.host;

		// console.log('update_gauge '+host+' '+tkn);//DEBUG

		app.callback_in_progress.update_gauge++; // increase callback in progress
		if (app.old_mode)		// Old API compatibility mode
			app_relay.get_cpx_param(f_call, host, '', 'CONFIG?affix_presence&affix_cumul_e&affix_cumul_s&affix_min&affix_min_ok&affix_max&affix_max_ok&affix_seuil_mode&affix_seuil&affix_raz'); // Concentrix/Affix/iComptipix API
		else
			app_relay.get_cpx_param(f_call, host, tkn, 'occupancy_state&occupancy_config');	// Comptipix-2D/3D API
	};

	/*!
		\brief	Update gauge

		\param 	result 	xhr object
	*/
	var update_gauge_callback = function(result)
	{
		if (2 != app.current_demo)//HACK
			return;//HACK

		app.callback_in_progress.update_gauge--; // decrease callback in progress
		var tmp = app_relay.parse_result(result);
		if (tmp.ok)
		{
			if (!app.old_mode) // Comptipix-2D/3D API
			{
				var tmp_occ_conf = tmp.param.occupancy_config;
				if (app.occ_conf_forced.forced)
					tmp_occ_conf = get_occ_conf_forced();
				app.gauge = app_gauge.plot(app.gauge, tmp.param.occupancy_state, tmp_occ_conf, 'occ_gauge');
			}
			else
			{	// Old API compatibility mode
				var new_occupancy_config = old_occupancy_config_to_new(tmp.param);
				var new_occupancy_state = old_occupancy_state_to_new(tmp.param);
				if (app.occ_conf_forced.forced)
					new_occupancy_config = get_occ_conf_forced();
				app.gauge = app_gauge.plot(app.gauge, new_occupancy_state, new_occupancy_config, 'occ_gauge');
			}
		}
		else
		{
			console.warn('ERR update_gauge_callback');
			console.log(tmp);
		}
	};

	/*!
		\brief	Update gauge from server file

		\param 	f_call 	function callback
	*/
	var update_gauge_server = function(f_call)
	{
		if (5 != app.current_demo)//HACK
			return;//HACK

		// Check all callback is ok
		if (!is_callback_ok('update_gauge_server'))
		{
			console.warn('Too many callback waiting ('+app.callback_in_progress.update_gauge_server+'/'+app.callback_in_progress.update_gauge_server_err+') for update_gauge_server');
			return;
		}

		var today = get_date_iso_no_hyphen(get_date_iso_str());
		var today_occ = today + '_presence.csv';

		app.callback_in_progress.update_gauge_server++; // increase callback in progress
		app_relay.fs_read(f_call, app.file_path + today_occ);
	};

	/*!
		\brief	Update gauge from server file callback

		\param 	result 	xhr object
	*/
	var update_gauge_server_callback = function(result)
	{
		if (5 != app.current_demo)//HACK
			return;//HACK

		app.callback_in_progress.update_gauge_server--; // decrease callback in progress
		var tmp = app_relay.parse_result(result);

		// Get last line of file
		// remove last '\n'
		if ((tmp.ok) && ('' !== tmp.req.content))
		{
			var n = tmp.req.content.lastIndexOf('\n');
			var tmp_line = tmp.req.content.slice(0, n); // remove last '\n'
			n = tmp_line.lastIndexOf('\n');	// get new last '\n' position
			tmp_line = tmp_line.slice(n+1, tmp_line.length);

			// Build an occupancy state
			var arr_line = tmp_line.split(',');
			var occ_state = [];
			occ_state[0] = '1'; // occupancy enabled
			occ_state[1] = arr_line[4]; // occupancy value
			occ_state[2] = arr_line[2]; // entry sum value
			occ_state[3] = arr_line[3]; // exits sum value

			// Build an occupancy config
			var occ_conf = ['1', '0', '0', '0', '0'];// occupancy enabled, no max no min no threshold

			if (app.occ_conf_forced.forced)
				occ_conf = get_occ_conf_forced();

			// Plot a gauge
			app.gauge = app_gauge.plot(app.gauge, occ_state, occ_conf, 'occ_gauge');
		}
		else
		{
			console.warn('ERR update_gauge_server_callback');
			console.warn(tmp);
		}
	};

	/*!
		\brief	Transform app.occ_conf_forced into occupancy config usable by app_gauge.plot()

	*/
	var get_occ_conf_forced = function()
	{
		var occ_option = app_gauge.define_option.CONFIG_ENABLE;
		var min_value = 0;
		var max_value = 0;
		var thres_value = 0;
		var reset_value = 0;
		if (null != app.occ_conf_forced.min)
		{
			occ_option ^= app_gauge.define_option.CONFIG_CLIP_MIN;
			min_value = app.occ_conf_forced.min;
		}
		if (null != app.occ_conf_forced.max)
		{
			occ_option ^= app_gauge.define_option.CONFIG_CLIP_MAX;
			max_value = app.occ_conf_forced.max;
		}
		if (null != app.occ_conf_forced.thres_min)
		{
			occ_option ^= app_gauge.define_option.CONFIG_THRESHOLD_MIN;
			thres_value = app.occ_conf_forced.thres_min;
		}
		if (null != app.occ_conf_forced.thres_max)
		{
			occ_option ^= app_gauge.define_option.CONFIG_THRESHOLD_MAX;
			thres_value = app.occ_conf_forced.thres_max;
		}

		return [occ_option, min_value, max_value, thres_value, reset_value];
	};

	/*!
		\brief	Get file to chart a counting + occupancy plot

		\param 	f_call 	function callback
		\param 	host 		target host
		\param 	tkn 		token to use
	*/
	var update_chart = function(f_call, host, tkn)
	{
		// Use app tkn if undefined
		if ('undefined' === typeof tkn)
			tkn = app.tkn;
		if ('undefined' === typeof host)
			host = app.host;

		// console.log('update_gauge '+host+' '+tkn);//DEBUG

		// Check all callback is ok
		if (!is_callback_ok('update_chart'))
		{
			console.warn('Too many callback waiting ('+app.callback_in_progress.update_chart+'/'+app.callback_in_progress.update_chart_err+') for update_chart');
			return;
		}

		// Get today file name
		var today = get_date_iso_no_hyphen(get_date_iso_str());
		var today_counting = today + '.csv';
		var today_occ = today + '_presence.csv';

		// Set app object
		app.chart_file_results = [0,0];
		// app.chart_file_contents = [0,0];

		// Get files
		app.callback_in_progress.update_chart++; // increase callback in progress
		if (!app.old_mode)
		{	// Comptipix-2D/3D API
			app_relay.get_cpx_param(f_call, host, tkn, 'sdcard_read='+today_counting);
			app_relay.get_cpx_param(f_call, host, tkn, 'sdcard_read='+today_occ);
		}
		else
		{	// Old API compatibility mode
			app_relay.get_cpx_param(f_call, host, '', 'FICHIER?lecture='+today_counting);
			app_relay.get_cpx_param(f_call, host, '', 'FICHIER?lecture='+today_occ);
		}
	};

	/*!
		\brief	Update chart callback

		\param 	result 	xhr object
	*/
	var update_chart_callback = function(result)
	{
		app.callback_in_progress.update_chart--; // decrease callback in progress
		var tmp = app_relay.parse_result(result);

		if (tmp.ok_relay)
		{
			var nb = 1;
			if (-1 == tmp.req.url.indexOf('_presence'))
				nb = 0;// Counting file

			// Set file result
			app.chart_file_results[nb] = tmp.req.status;
			if (200 == tmp.req.status)
				app.chart_file_contents[nb] = tmp.req.content;
			else
				app.chart_file_contents[nb] = 0;// no content

			// Check if we can try to plot
			if ((200 == app.chart_file_results[0])&&(0 !== app.chart_file_results[0]))
			{
				// Chart
				var tmp_files = [];
				tmp_files.push(app.chart_file_contents[0]);
				if (200 == app.chart_file_results[1])	// Add occupancy if ok
					tmp_files.push(app.chart_file_contents[1]);

				// Check file are different
				var flag_plot = false;
				if ((app.chart_file_length[0] != app.chart_file_contents[0].length) || (app.chart_file_length[1] != app.chart_file_contents[1].length))
					flag_plot = true;
				// Update length
				app.chart_file_length[0] = app.chart_file_contents[0].length;
				app.chart_file_length[1] = app.chart_file_contents[1].length;

				// New chart
				if (flag_plot)
				{
					if (null !== app.chart)
						app.chart.destroy();
					app.chart = app_chart.plot('plot', tmp_files, 'Counting');
				}
				else
				{
					// This files are already ploted ...
				}
			}
			else
			{
				// Waiting for another file ...
			}
		}
		else if (503 == tmp.req.status)
		{
			console.log('update_chart_callback -> ON 503 retry');
			// 503 = sdcard not available ... wait 3 second and retry
			window.setTimeout(function(){
				var today = get_date_iso_no_hyphen(get_date_iso_str());
				if (-1 == tmp.req.url.indexOf('_presence.csv'))
				{
					if (!app.old_mode)
						app_relay.get_cpx_param(update_chart_callback, app.host, app.tkn, 'sdcard_read='+today+'_presence.csv');	//Comptipix-2D/3D API
					else
						app_relay.get_cpx_param(update_chart_callback, app.host, '', 'FICHIER?lecture='+today+'_presence.csv');	// Old API compatibility mode
				}
				else
				{
					if (!app.old_mode)
						app_relay.get_cpx_param(update_chart_callback, app.host, app.tkn, 'sdcard_read='+today+'.csv');	//Comptipix-2D/3D API
					else
						app_relay.get_cpx_param(update_chart_callback, app.host, '', 'FICHIER?lecture='+today+'.csv');	// Old API compatibility mode
				}
			}, 3000);
		}
		else
			console.warn('ERR update_chart_callback connection to relay fail '+result.status);
	};


	/*!
		\brief	Update counting table

		\param 	f_call 	function callback
		\param 	host 		target host
		\param 	tkn 		token to use
	*/
	var update_counting = function(f_call, host, tkn)
	{
		// Use app tkn if undefined
		if ('undefined' === typeof tkn)
			tkn = app.tkn;
		if ('undefined' === typeof host)
			host = app.host;

		// Check all callback is ok
		if (!is_callback_ok('update_counting'))
		{
			console.warn('Too many callback waiting ('+app.callback_in_progress.update_counting+'/'+app.callback_in_progress.update_counting_err+') for update_counting');
			return;
		}

		// console.log('update_counting '+host+' '+tkn);//DEBUG
		app.callback_in_progress.update_counting++; // increase callback in progress
		if (!app.old_mode)	//Comptipix-2D/3D API
			app_relay.get_cpx_param(f_call, host, tkn, 'occupancy_state');
		else
			app_relay.get_cpx_param(f_call, host, '', 'CONFIG?affix_presence&affix_cumul_e&affix_cumul_s'); // OLD API compatibility mode
	};

	/*!
		\brief	Update counting

		\param 	result 	xhr object
	*/
	var update_counting_callback = function(result)
	{
		app.callback_in_progress.update_counting--; // decrease callback in progress
		var tmp = app_relay.parse_result(result);
		if (tmp.ok)
		{
			if (app.old_mode)
				tmp.param.occupancy_state = old_occupancy_state_to_new(tmp.param);
			document.getElementById('demo_4__counting_e').innerHTML = tmp.param.occupancy_state[2];
			document.getElementById('demo_4__counting_x').innerHTML = tmp.param.occupancy_state[3];
			document.getElementById('demo_4__counting_occ').innerHTML = tmp.param.occupancy_state[1];
		}
		else
		{
			console.warn('ERR update_counting_callback');
			console.warn(tmp);
		}
	};

	/*!
		\brief	Update hour display

		\param 	result 	xhr object
	*/
	var update_hour = function()
	{
		var d = new Date();
		var h = d.getHours();
		var m = d.getMinutes();
		var s = d.getSeconds();
		document.getElementById('footer_hour').innerHTML = tostring_pad(d.getHours(), 2)+':'+tostring_pad(d.getMinutes(), 2)+':'+tostring_pad(d.getSeconds(), 2);
	};

	/*!
		\brief	On get token callback

		\param 	result 	xhr object
	*/
	var get_cpx_tkn_callback = function(result)
	{
		console.log('get_cpx_tkn_callback');
		console.log(result);

		var tmp = app_relay.parse_result(result);
		if (tmp.ok && (!app.demo_started))
			start_demo(tmp.req.host, tmp.req.content);
		else
		{
			start_demo(tmp.req.host); // Still start demo : fs_read can always work
			console.warn('ERR get_cpx_tkn_callback');
			console.warn(tmp);
		}
	};

	/*!
		\brief	Update token if needed

		\param 	result 	xhr object
	*/
	var check_update_tkn = function(f_call, host, tkn)
	{
		// Use app tkn if undefined
		if ('undefined' === typeof tkn)
			tkn = app.tkn;
		if ('undefined' === typeof host)
			host = app.host;

		if ((1 == app.current_demo) || (5 == app.current_demo))
		{
			console.log('check_update_tkn() -> No token update for mode 1 or 5');
			return;
		}

		// Direct check for null token  -> tkn never initialized
		if (null === tkn)
		{
			update_tkn(app.host);
			return;
		}

		if (app.old_mode)
			app_relay.get_cpx_param(f_call, host, '', 'CONFIG?uptime');
		else
			app_relay.get_cpx_param(f_call, host, tkn, 'uptime');
	};

	/*!
		\brief	Update token if needed callback

		\param 	result 	xhr object
	*/
	var check_update_tkn_callback = function(result)
	{
		var tmp = app_relay.parse_result(result);
		if (!tmp.ok_cpx)
		{
			if ((401 == tmp.req.status)||(403 == tmp.req.status)||(418 == tmp.req.status))
				update_tkn(tmp.req.host);
		}
	};

	/*!
		\brief	Get a new token and update app.tkn

		\param 	host 	host to get token
	*/
	var update_tkn = function(host)
	{
		if ('undefined' === typeof host)
			host = app.host;

		if ((1 == app.current_demo) || (5 == app.current_demo))
		{
			console.log('update_tkn() -> No token update for mode 1 or 5');
			return;
		} else if (('' == app.host) || (null === app.host))
		{
			console.log('update_tkn() -> No host for mode 2 or 3 or 4 : try to fix it');
			set_host();
			console.log('HOST SET !');
			console.log(app.host);
		}

		console.log('YOOOOOO '+app.current_demo+' : '+app.host);
		console.log(app);

		if (app.old_mode)
			app_relay.get_cpx_tkn_old(update_tkn_callback, host, app.user, app.pass);
		else
			app_relay.get_cpx_tkn(update_tkn_callback, host, app.user, app.pass);
	};

	/*!
		\brief	Update app.tkn

		\param 	result 	xhr object
	*/
	var update_tkn_callback = function(result)
	{
		var tmp = app_relay.parse_result(result);

		// Check if demo started, and start it if it's not the case
		get_cpx_tkn_callback(result); // will check if a start app is needed

		if (tmp.ok)
			app.tkn = tmp.req.content;
		else if ((tmp.ok_relay)&&(418 == tmp.req.status))
		{
			if (!app.pass_err)
			{
				window.alert("Login/Password Error");
				app.pass_err = true;
			}
		}
	};


	// ---------------------
	// Public method
	// ---------------------

	/*!
		\brief	Start demo

	*/
	var start_demo = function(host, tkn)
	{
		console.log('START demo ('+host+', '+tkn+') !!! !!! !!! !!! !!!');
		if ('undefined' !== typeof tkn)
			app.tkn = tkn;
		if ('undefined' !== typeof host)
			app.host = host;

		// Periodically update gauge every 1 second
		if (null === app.intervals.gauge)
		{
			console.log('START gauge');
			update_gauge(update_gauge_callback);
			app.intervals.gauge = window.setInterval(function(){
				if (2 == app.current_demo)
					update_gauge(update_gauge_callback);
			}, 1000);
		}

		// Periodically plot today files every 10 second
		if (null === app.intervals.chart)
		{
			console.log('START chart');
			update_chart(update_chart_callback);
			app.intervals.chart = window.setInterval(function(){
				if (3 == app.current_demo)
					update_chart(update_chart_callback);
			}, 10000);
		}

		// Periodically update counting
		if (null === app.intervals.counting)
		{
			console.log('START counting');
			update_counting(update_counting_callback);
			app.intervals.counting = window.setInterval(function(){
				if (4 == app.current_demo)
					update_counting(update_counting_callback);
			}, 1000);
		}

		// Periodicaly get file from server
		if (null === app.intervals.gauge_server)
		{
			console.log('START gauge_server');
			update_gauge_server(update_gauge_server_callback);
			app.intervals.gauge_server = window.setInterval(function(){
				if (5 == app.current_demo)
					update_gauge_server(update_gauge_server_callback);
			}, 5000);
		}

		// Periodicaly check/update our token
		if (null === app.intervals.tkn)
		{
			console.log('START tkn');
			app.intervals.tkn = window.setInterval(function(){
				if (5 != app.current_demo)
					check_update_tkn(check_update_tkn_callback);
			}, 5000);
		}

		// Periodicaly update hour
		update_hour();
		if (null === app.intervals.hour)
		{
			console.log('START hour');
			app.intervals.hour = window.setInterval(update_hour, 1000);
		}

		app.demo_started = true;
	};

	/*!
		\brief	Show a demo and hide other

	*/
	var set_demo = function(num_show)
	{
		var num_demo = num_show;
		if (num_demo == 5) //HACK: Demo 5 and 2 share the same div
			num_show = 2;

		// Hide all demo
		for (var i=1; i<=5; i++)
		{
			var t = document.getElementById('demo_'+i);
			if (t !== null)
				t.setAttribute('hidden', ' ');
			t = document.getElementById('demo_'+i+'_btn');
			if (t !== null)
				t.removeAttribute('disabled', ' ');
		}

		// Show demo
		// console.log('SET_demo '+num_show+' , '+num_demo);
		document.getElementById('demo_'+num_show).removeAttribute('hidden');
		document.getElementById('demo_'+num_demo+'_btn').setAttribute('disabled', ' ');

		// Hack erase the current gauge
		// console.log(app.current_demo+' <-> '+num_demo);
		if (	(((app.current_demo == 2) && (num_demo == 5)) || 
				((app.current_demo == 5) && (num_demo == 2))) &&
				(null !== app.gauge)
			)
		{
			destroy_gauge();
			document.getElementById('occ_gauge').innerHTML = '';
		}

		// Set current demo
		app.current_demo = num_demo;
	};

	/*!
		\brief	Replace host + get new token


	*/
	var set_host = function(new_host, new_user, new_pass)
	{
		if ('undefined' === typeof new_host)
			new_host = document.getElementById('host').value;
		if ('undefined' === typeof new_user)
			new_user = document.getElementById('user').value;
		if ('undefined' === typeof new_pass)
			new_pass = document.getElementById('pass').value;

		// Hack destroy gauge : to update min/max/threshold to new sensor
		destroy_gauge();
		document.getElementById('occ_gauge').innerHTML = '';
		if (null !== app.chart)
			app.chart.destroy();

		app.host = 'http://' + new_host + '/';
		app.user = new_user;
		app.pass = new_pass;
		app.pass_err = false;
		app.file_path = document.getElementById('file_path').value;
		if ((app.file_path.charAt(app.file_path.length - 1) != '/') && (app.file_path.length > 1))	// Fix file path (must end with '/')
		{
			app.file_path += '/';
			document.getElementById('file_path').value = app.file_path;
		}
		app.old_mode = document.getElementById('old_mode').checked;

		// Set URL
		var new_url = window.location.origin + '/';
		new_url += '?host=' + encodeURI(new_host);
		new_url += '&user=' + encodeURI(new_user);
		new_url += '&pass=' + encodeURI(new_pass);
		new_url += '&file_path=' + encodeURI(app.file_path);
		if (app.old_mode)
			new_url += '&old_mode=' + 1;
		else
			new_url += '&old_mode=' + 0;
		new_url += '&demo=' + app.current_demo;
		new_url += '&hide_mouse=' + app.hide_mouse;
		new_url += '&hide_eurecam=' + app.hide_eurecam;
		new_url += '&hide_hour=' + app.hide_hour;
		new_url += '&hide_buttons=' + app.hide_buttons;
		if (app.occ_conf_forced.forced)
		{
			if (null != app.occ_conf_forced.min)
				new_url += '&occ_conf.min=' + app.occ_conf_forced.min;
			if (null != app.occ_conf_forced.max)
				new_url += '&occ_conf.max=' + app.occ_conf_forced.max;
			if (null != app.occ_conf_forced.thres_min)
				new_url += '&occ_conf.thres_min=' + app.occ_conf_forced.thres_min;
			if (null != app.occ_conf_forced.thres_max)
				new_url += '&occ_conf.thres_max=' + app.occ_conf_forced.thres_max;
		}
		// console.log('NEW url = '+new_url);
		history.pushState(null, null, new_url);

		// Get token and start the app
		check_update_tkn(check_update_tkn_callback);

		if ((1 == app.current_demo) || (5 == app.current_demo))
		{
			console.log('start demo without token');
			start_demo(null, null);
		}
	};

	/*!
		\brief	Add one CSS rules

		\param	selector				selector rule
		\param	rules					css rule
		\param	sheet					sheet to apply rule				(DEFAULT: document.styleSheets[0])
		\param	index					default rule index				(DEFAULT:0)
	*/
	var add_css_rule = function(selector, rules, sheet, index)
	{
		// apply default value
		if ('undefined' == typeof sheet)
			sheet = document.styleSheets[0];
		if ('undefined' == typeof index)
			index = 0;

		if ('insertRule' in sheet)
			sheet.insertRule(selector + '{' + rules + '}', index);
		else if ('addRule' in sheet)
			sheet.addRule(selector, rules, index);
	};

	/*!
		\brief	On load method

	*/
	var onload = function()
	{
		// Read url
		var tmp = window.location.search.replace('?', '').split('&');
		var tmp_host = '192.168.0.139';
		var tmp_demo = 2; 			// demo_2 is occupancy gauge
		var tmp_user = 'reader';
		var tmp_pass = 'reader';
		var tmp_file_path = '';
		var tmp_old_mode = false;
		var tmp_hide_buttons = 0;
		var tmp_hide_mouse = 0;
		var tmp_hide_eurecam = 0;
		var tmp_hide_hour = 0;
		for (var tmp_in in tmp)
		{
			var tmp_in_t = tmp[tmp_in].split('=');
			if (tmp_in_t.length >= 2)
			{
				// Read host
				if ('host' == tmp_in_t[0])
					tmp_host = tmp_in_t[1];
				else if ('demo' == tmp_in_t[0])
					tmp_demo = tmp_in_t[1];
				else if ('user' == tmp_in_t[0])
					tmp_user = tmp_in_t[1];
				else if ('pass' == tmp_in_t[0])
					tmp_pass = tmp_in_t[1];
				else if ('file_path' == tmp_in_t[0])
					tmp_file_path = tmp_in_t[1];
				else if ('old_mode' == tmp_in_t[0])
					tmp_old_mode = tmp_in_t[1];
				else if ('hide_buttons' == tmp_in_t[0])
					tmp_hide_buttons = tmp_in_t[1];
				else if ('hide_mouse' == tmp_in_t[0])
					tmp_hide_mouse = tmp_in_t[1];
				else if ('hide_eurecam' == tmp_in_t[0])
					tmp_hide_eurecam = tmp_in_t[1];
				else if ('hide_hour' == tmp_in_t[0])
					tmp_hide_hour = tmp_in_t[1];
				else if ('occ_conf.min' == tmp_in_t[0])
				{
					app.occ_conf_forced.forced = true;
					app.occ_conf_forced.min = tmp_in_t[1];
				}
				else if ('occ_conf.max' == tmp_in_t[0])
				{
					app.occ_conf_forced.forced = true;
					app.occ_conf_forced.max = tmp_in_t[1];
				}
				else if ('occ_conf.thres_min' == tmp_in_t[0])
				{
					app.occ_conf_forced.forced = true;
					app.occ_conf_forced.thres_min = tmp_in_t[1];
				}
				else if ('occ_conf.thres_max' == tmp_in_t[0])
				{
					app.occ_conf_forced.forced = true;
					app.occ_conf_forced.thres_max = tmp_in_t[1];
				}
			}
		}

		// console.log('app.occ_conf_forced');
		// console.log(app.occ_conf_forced);

		// Set host
		document.getElementById('host').value = tmp_host;
		document.getElementById('user').value = tmp_user;
		document.getElementById('pass').value = tmp_pass;
		document.getElementById('file_path').value = tmp_file_path;
		document.getElementById('old_mode').checked = false;
		if (tmp_old_mode > 0)
			document.getElementById('old_mode').checked = true;

		// Read file from fs
		app_relay.fs_read(fs_read_callback, 'toto.txt');

		// Show demo
		set_demo(parseInt(tmp_demo, 10));

		// Hide button
		if (tmp_hide_buttons > 0)
		{
			document.getElementById('buttons').setAttribute('hidden', ' ');
			app.hide_buttons = 1;
		}

		// Hide mouse
		if (tmp_hide_mouse > 0)
		{
			add_css_rule('.hide-mouse', 'cursor: none;');
			app.hide_mouse = 1;
		}

		// Hide Eurecam
		if (tmp_hide_eurecam > 0)
		{
			document.getElementById('footer_eurecam').setAttribute('hidden', ' ');
			app.hide_eurecam = 1;
		}

		// Hide Hour
		if (tmp_hide_hour > 0)
		{
			document.getElementById('footer_hour').setAttribute('hidden', ' ');
			app.hide_hour = 1;
		}

		set_host();
	};

	// ---------------------
	// Return public method
	// ---------------------

	return {
		adapt_user_pass_old_mode:	adapt_user_pass_old_mode,
		onload:							onload,
		start_demo:						start_demo,
		set_host:						set_host,
		set_demo:						set_demo
	};
})();
