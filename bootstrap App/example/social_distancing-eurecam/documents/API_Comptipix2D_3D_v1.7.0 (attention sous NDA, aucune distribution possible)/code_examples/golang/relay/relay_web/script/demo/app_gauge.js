/*!
	\brief	specialize a page object to a occupancy page object

	\param	page	the page object to specialize
*/
var app_gauge = (function()
{
	'use strict';// we are strict (see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions_and_function_scope/Strict_mode?redirectlocale=en-US&redirectslug=JavaScript%2FReference%2FFunctions_and_function_scope%2FStrict_mode)

	// ---------------------
	// Definition
	// ---------------------

	// definition of occupancy_config_0 bitmask parameters
	var define_option = {
		CONFIG_ENABLE:				0x01,
		CONFIG_CLIP_MIN:			0x02,
		CONFIG_CLIP_MAX:			0x04,
		CONFIG_THRESHOLD_MAX:	0x10,
		CONFIG_THRESHOLD_MIN:	0x20
	};

	// definition of threshold type
	var thres_type = {
		NO:							0,
		MAX:							1,
		MIN:							2
	};

	/*!
		\brief	Define an ocupancy config object

	*/
	var occ_config_obj = function()
	{
		this.enabled = false;						// occupancy enabled
		this.thres_type = thres_type.NO;			// threshold type
		this.thres_value = 0;						// number
		this.max_enabled = false;					// boolean
		this.max_value = 0;							// number
		this.min_enabled = false;					// boolean
		this.min_value = 0;							// number
		// this.reset_value = 0;						// number
		// this.auto_rst_enabled = [];				// array of boolean
		// this.auto_rst_value = [];					// array of number
	};

	// ---------------------
	// Private method
	// ---------------------

	/*!
		\brief	Read occupancy config from comptipix

		\param 	occ_conf		occupancy_config as array
		\return  an object containing occupancy config
	*/
	var read_occupancy_config = function(occ_conf)
	{
		// apply default values
		if ('undefined' === typeof occ_conf)
			return null;

		var occ_config_tmp = new occ_config_obj();

		// enabled occupancy ?
		occ_config_tmp.enabled = define_option.CONFIG_ENABLE&parseInt(occ_conf[0], 10);	// occupancy enabled

		// Threshold
		if (define_option.CONFIG_THRESHOLD_MAX&parseInt(occ_conf[0], 10)) // test thres max
			occ_config_tmp.thres_type = thres_type.MAX;
		else if (define_option.CONFIG_THRESHOLD_MIN&parseInt(occ_conf[0], 10)) // test thres min
			occ_config_tmp.thres_type = thres_type.MIN;
		else
			occ_config_tmp.thres_type = thres_type.NO;
		occ_config_tmp.thres_value = parseInt(occ_conf[3], 10); // threshold value

		// Max and Min
		occ_config_tmp.max_enabled = (define_option.CONFIG_CLIP_MAX&parseInt(occ_conf[0], 10))>0?true:false;// boolean
		occ_config_tmp.max_value = parseInt(occ_conf[2], 10);// max value
		occ_config_tmp.min_enabled = (define_option.CONFIG_CLIP_MIN&parseInt(occ_conf[0], 10))>0?true:false;// boolean
		occ_config_tmp.min_value = parseInt(occ_conf[1], 10);// min value

		return occ_config_tmp;
	};


	// ---------------------
	// Public method
	// ---------------------

	/*!
		\brief	Plot a gauge, reuse a gauge object to update an existing gauge

		\param 	gauge_ref				justgage object to use
		\param 	occupancy_state		occupancy_state from comptipix as array
		\param 	occupancy_config		occupancy_config from comptipix as array
		\param 	id							dom id to plot gauge								DEFAULT: 'occ_gauge'
		\param 	title_label				title label to use								DEFAULT: 'Occupancy'
		\param 	people_label			people label to use								DEFAULT: 'people'
		\param 	big_num_sep				big number separator								DEFAULT: ' '
		\return  justgage object used
	*/
	var plot = function(gauge_ref, occupancy_state, occupancy_config, id, title_label, people_label, big_num_sep)
	{
		if (('undefined' === typeof gauge_ref) || (null === gauge_ref))
			gauge_ref = null;
		if (('undefined' === typeof id) || (null === id))
			id = 'occ_gauge';
		if (('undefined' === typeof title_label) || (null === title_label))
			title_label = 'Occupancy';
		if (('undefined' === typeof people_label) || (null === people_label))
			people_label = 'people';
		if (('undefined' === typeof big_num_sep) || (null === big_num_sep))
			big_num_sep = ' ';

		// value used for gauge
		var gauge_max = 0;
		var gauge_min = 0;
		var gauge_showmin = false;
		var gauge_showmax = false;
		var gauge_hideminmax = false;
		var gauge_colors = [];

		var occ_conf = read_occupancy_config(occupancy_config);
		if (null === occ_conf)
		{
			console.warn('app_gauge.plot occupancy_config WRONG -> ABORT');
			return;
		}

		// console.log('PLOT('+id+') ');
		// console.log(occ_conf);
		// console.log(occupancy_state);

		// gauge max
		if (occ_conf.max_enabled)
		{
			gauge_max = occ_conf.max_value;
			gauge_showmax = true;
		}
		else if ((thres_type.MAX == occ_conf.thres_type) || (thres_type.MIN == occ_conf.thres_type))
		{
			gauge_max = occ_conf.thres_value;
			gauge_showmax = true;
		}
		else
		{
			// There is no max : so max = current+20 and we don't show min and max
			gauge_max = parseInt(occupancy_state[1], 10)+20;
			gauge_showmax = false;
		}

		// gauge min
		if (occ_conf.min_enabled)
		{
			gauge_min = occ_conf.min_value;
			gauge_showmin = true;
		}
		else
		{
			// There is no min : so min = current-20 (if currzent <0) and we don't show min and max
			if (parseInt(occupancy_state[1], 10) >= 0)
				gauge_min = 0;
			else
				gauge_min = parseInt(occupancy_state[1], 10)-20;
			gauge_showmin = false;
		}

		// justgage has just a showminmax (no showmin and showmax separated)
		if (gauge_showmin||gauge_showmax)
			gauge_hideminmax = false;
		else
			gauge_hideminmax = true;

		// gauge color + revert
		var reverse_gauge = false;
		if (thres_type.MIN == occ_conf.thres_type)
		{
			reverse_gauge = true;
			gauge_colors = [{
							color : "#FF0000",
							lo : gauge_min,
							hi : gauge_min+(10*gauge_min/100)
						},{
							color : "#FFA500",
							lo : gauge_min+(10*gauge_min/100),
							hi : gauge_min+(25*gauge_min/100)
						},{
							color : "#FFFF00",
							lo : gauge_min+(25*gauge_min/100),
							hi : gauge_min+(50*gauge_min/100)
						},{
							color : "#00B900",
							lo : gauge_min+(50*gauge_min/100),
							hi : gauge_min+(100*gauge_min/100)
					}];
		}
		else if (thres_type.MAX == occ_conf.thres_type)
		{
			gauge_colors = [{
							color : "#FF0000",
							lo : gauge_max-(10*gauge_max/100),
							hi : gauge_max
						},{
							color : "#FFA500",
							lo : gauge_max-(30*gauge_max/100),
							hi : gauge_max-(10*gauge_max/100)
						},{
							color : "#FFFF00",
							lo : gauge_max-(50*gauge_max/100),
							hi : gauge_max-(30*gauge_max/100)
						},{
							color : "#00B900",
							lo : gauge_max-(100*gauge_max/100),
							hi : gauge_max-(50*gauge_max/100)
					}];
		}
		else
		{
			gauge_colors = [{
							color : "#00B900",
							lo : -9007199254740990, // this is the value of MAX_SAFE_INTEGER see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/MAX_SAFE_INTEGER
							hi : 9007199254740990
					}];
		}

		// console.log('occ_conf');
		// console.log(occ_conf);

		// build/update occupancy gauge
		var gauge_elm = document.getElementById(id);
		if (null !== gauge_elm)
		{
			// Test gauge already exist, but have to be rebuild because of config change (hideminmax or color -> min and max can be updated)
			if (null !== gauge_ref) // Gauge is initialized
			{
				if (gauge_hideminmax != gauge_ref.config.hideMinMax) // internal justgage is in 'config'
				{
					// Destroy gauge to force a rebuild
					gauge_ref = null;
					gauge_elm.innerHTML = '';
				}
			}

			if (null === gauge_ref) // Gauge is uninitialized
			{
				// Get default format number
				var format_number = true;
				var format_number_symbol = '';
				if ('' !== big_num_sep)
				{
					format_number = true;
					if (',' != big_num_sep)
						format_number_symbol = big_num_sep; // For french but need to patch justgage
				}
				else
					format_number = false;

				// Justgage conf object
				// console.log('Build max/min(revert) : '+gauge_max+'/'+gauge_min+'('+reverse_gauge+')');
				var tmp_justgage_conf = {
					id: id,// this id must correspond to html id
					value: parseInt(occupancy_state[1], 10),
					symbol: '',
					min: gauge_min,
					max: gauge_max,
					hideMinMax : gauge_hideminmax,
					title: title_label,
					label: people_label,
					relativeGaugeSize: true,
					labelFontColor: '#1A1A1A',
					// levelColors: gauge_colors,
					customSectors: {
						percents: false,
						range: gauge_colors
					},
					counter: false,
					gaugeWidthScale: 0.2,
					formatNumber: format_number,
					formatNumberSymbol: format_number_symbol,
					reverse: reverse_gauge,
					donut: false
				};

				// Init justgage
				gauge_ref = new JustGage(tmp_justgage_conf);
			}
			else {
				// console.log('Update max/min(reverse) : '+gauge_max+'/'+gauge_min+'('+reverse_gauge+')');
				if (reverse_gauge)
					gauge_ref.refresh(parseInt(occupancy_state[1], 10)); // Refresh justgage with min and max is buggy when gauge is reverted -> TODO fix it
				else
					gauge_ref.refresh(parseInt(occupancy_state[1], 10), gauge_max, gauge_min);
			}
		}
		else
			console.war('app_gauge.plot id:'+id+' NOT FOUND !!!');

		return gauge_ref;
	};


	// ---------------------
	// Return public method
	// ---------------------

	return {
		define_option:	define_option,
		thres_type:		thres_type,
		plot:				plot
	};
})();
