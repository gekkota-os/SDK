// -----------------------------------------
// Parse data file Module
// -----------------------------------------

/*!
	\brief	Module containing message utility function

*/
var app_data_file = (function()
{
	'use strict';// we are strict (see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions_and_function_scope/Strict_mode?redirectlocale=en-US&redirectslug=JavaScript%2FReference%2FFunctions_and_function_scope%2FStrict_mode)

	// Init object
	var parse_init_object = function()
	{
		var t_object = {};
		t_object.data = [];
		t_object.tick = [];
		//t_object.tick_10 = [];
		//t_object.tick_30 = [];
		t_object.min = 0;
		t_object.max = 0;

		// Generate hour tick : ['0h', '1h', '2h', '3h', '4h', '5h', '6h', '7h', '8h', '9h', '10h', '11h', '12h', '13h', '14h', '15h', '16h', '17h', '18h', '19h', '20h', '21h', '22h', '23h']
		for (var h=0; h<24; h++)
			t_object.tick[h] = h+'h';

		/*
		// Generate 10 minutes tick
		var i = 0;
		var minute = 0;
		for (h=0; h<144; h++)
		{
			if (minute > 0)
				t_object.tick_10[h] = i+'h'+minute;
			else
				t_object.tick_10[h] = i+'h0'+minute;

			minute += 10;
			if (minute > 50)
			{
				i++;
				minute = 0;
			}
		}
		//console.log(t_object.tick_10);//DEBUG

		// Generate 30 minutes tick
		i = 0;
		minute = 0;
		for (h=0; h<48; h++)
		{
			if (minute > 0)
				t_object.tick_30[h] = i+'h'+minute;
			else
				t_object.tick_30[h] = i+'h0'+minute;

			minute += 30;
			if (minute > 30)
			{
				i++;
				minute = 0;
			}
		}
		//console.log(t_object.tick_30);//DEBUG
		*/

		return t_object;
	};

	// Generate scale according to min and max
	var data_generate_scale = function(min, max)
	{
		min *= 1.2;
		max *= 1.2;

		if (max==min)
		{
			if (min<0)
				min -= 1;
			else
				max += 1;
		}

		var range = max-min;
		if (range<1)
			range = 1;

		var tick = 1;
		var scale = 1;

		while(range/tick/scale>9)
		{
			if (tick==1)
				tick = 2;
			else if (tick==2)
				tick = 5;
			else
			{
				tick = 1;
				scale *= 10;
			}
		}

		tick = tick*scale;
		min = Math.floor(min/tick)*tick;
		max = Math.ceil(max/tick)*tick;

		return {min:min, max:max, tick:tick};
	};

	// add data channel in data object
	var parse_add_data = function(object, nb_data, channel, value)
	{
		while(object.data.length<nb_data)
		{
			var t_data = {};
			t_data.name = '';

			for(var i=0;i<channel.length;++i)
				t_data[channel[i]] = [value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value, value];

			object.data.push(t_data);
		}
	};

	// Parse data into date and check date
	var parse_date_check = function(data, date)
	{
		var t_split = data.split('/');
		if (t_split.length!=3)
			return null;

		// check day

		var t_day = parseInt(t_split[0], 10);
		if ((t_day<1)||(t_day>31))
			return null;

		// check month

		var t_month = parseInt(t_split[1], 10);
		if ((t_month<1)||(t_month>12))
			return null;

		// check year

		var t_year = parseInt(t_split[2], 10);
		if (t_year<0)
			return null;

		if (date)
		{
			date.setFullYear(t_year);
			date.setMonth(t_month);
			date.setDate(t_day);
			return true;
		}

		return new Date(t_year, t_month, t_day, 0, 0, 0, 0);
	};

	// Parse data into hour and check hour
	var parse_hour_check = function(data, date)
	{
		var t_split = data.split(':');
		if (t_split.length!=3)
			return null;

		// check hour

		var t_hour = parseInt(t_split[0], 10);
		if ((t_hour<0)||(t_hour>23))
			return null;

		// check minute

		var t_minute = parseInt(t_split[1], 10);
		if ((t_minute<0)||(t_minute>59))
			return null;

		// check second

		var t_second = parseInt(t_split[2], 10);
		if ((t_second<0)||(t_second>59))
			return null;

		if (date)
		{
			date.setHours(t_hour);
			date.setMinutes(t_minute);
			date.setSeconds(t_second);
			return true;
		}

		return new Date(0, 0, 0, t_hour, t_minute, t_second);
	};

	// Parse counting file
	// returns an object with :
	// - data[]
	// 	- name
	// 	- entry[24]
	// 	- exit[24]
	// - tick[24]
	var parse_counting_file = function(data)
	{
		var t_object = parse_init_object();

		// add default series when no data

		if (!data)
		{
			parse_add_data(t_object, 1, ['entry', 'exit'], 0);
			return t_object;
		}

		// parse data

		var t_index = 0;

		while(1)
		{
			// get line

			var t_next = data.indexOf('\n', t_index);

			var t_line = data.substr(t_index, t_next-t_index);
			if (t_next<0)
				t_line = data.substr(t_index);

			// parse line

			var t_split = t_line.split(',');
			var t_channel;
			if (t_split.length>3)
			{
				// Counting data -> like '06/02/2017,00:00:00,0,0' or 'Date,Heure,E,S'

				var t_date = new Date();
				t_channel = Math.floor((t_split.length-2)/2); // channel is (E + S) * nb_channel + 2 element ('Date,Heure')
				parse_add_data(t_object, t_channel, ['entry', 'exit'], 0);

				if (parse_date_check(t_split[0], t_date) && parse_hour_check(t_split[1], t_date))
				{
					// Counting data only like '06/02/2017,00:00:00,0,0'

					var t_hour = t_date.getHours();

					// console.log(t_date);//DEBUG
					// console.log(t_hour);//DEBUG
					// console.log(t_channel);//DEBUG

					for(var t=0;t<t_channel;++t)
					{
						t_object.data[t].entry[t_hour] += parseInt(t_split[2+t*2], 10);
						t_object.data[t].exit[t_hour] += parseInt(t_split[3+t*2], 10);
					}
				}
			}
			else if (t_split.length==3)
			{
				// Channel name -> like '1,sumy,acces'
				// check data name

				t_channel = parseInt(t_split[0], 10);

				if (t_channel>0)
				{
					parse_add_data(t_object, t_channel, ['entry', 'exit'], 0);
					t_object.data[t_channel-1].name = t_split[1];
				}
			}

			if (t_next<0)
				break;
			t_index = t_next+1;
		}

		// add default series when no data

		if (0 === t_object.data.length)
			parse_add_data(t_object, 1, ['entry', 'exit'], 0);

		// update max and sum

		for(var j=0;j<t_object.data.length;++j)
		{
			t_object.data[j].entry_sum = 0;
			t_object.data[j].exit_sum = 0;

			for(var i=0;i<24;++i)
			{
				t_object.data[j].entry_sum += t_object.data[j].entry[i];
				if (t_object.data[j].entry[i]>t_object.max)
					t_object.max = t_object.data[j].entry[i];

				t_object.data[j].exit_sum += t_object.data[j].exit[i];
				if (t_object.data[j].exit[i]>t_object.max)
					t_object.max = t_object.data[j].exit[i];
			}
		}

		return t_object;
	};

	// Parse occupancy file
	// returns an object with :
	// - data[]
	// 	- name
	// 	- occupancy_min[24]
	// 	- occupancy_max[24]
	// - tick[24]
	var parse_occupancy_file = function(data)
	{
		var t_object = parse_init_object();

		// add default series when no data

		if (!data)
		{
			parse_add_data(t_object, 1, ['min', 'max'], 0);
			return t_object;
		}

		// parse data

		var t_index = 0;

		while(1)
		{
			// get line

			var t_next = data.indexOf('\n', t_index);

			var t_line = data.substr(t_index, t_next-t_index);
			if (t_next<0)
				t_line = data.substr(t_index);

			// parse line

			var t_split = t_line.split(',');
			if (t_split.length==7)
			{
				var t_date = new Date();
				parse_add_data(t_object, 1, ['min', 'max'], null);

				if (parse_date_check(t_split[0], t_date) && parse_hour_check(t_split[1], t_date))
				{
					var t_hour = t_date.getHours();
					var t_occ = parseInt(t_split[4], 10);

					if (null === t_object.data[0].min[t_hour])
						t_object.data[0].min[t_hour] = t_occ;
					else if (t_occ < t_object.data[0].min[t_hour])
						t_object.data[0].min[t_hour] = t_occ;

					if (null === t_object.data[0].max[t_hour])
						t_object.data[0].max[t_hour] = t_occ;
					else if (t_occ > t_object.data[0].max[t_hour])
						t_object.data[0].max[t_hour] = t_occ;
				}
			}
			else if (t_split.length==2)
			{
				parse_add_data(t_object, 1, ['min', 'max'], null);
				t_object.data[0].name = t_split[0];
			}

			if (t_next<0)
				break;
			t_index = t_next+1;
		}

		// replace null by 0

		for(var k=0; k<24; ++k)
		{
			if (null === t_object.data[0].min[k])
				t_object.data[0].min[k] = 0;

			if (null === t_object.data[0].max[k])
				t_object.data[0].max[k] = 0;
		}

		// update min-max

		for(var j=0; j<t_object.data.length; ++j)
		{
			for(var i=0; i<24; ++i)
			{
				if (t_object.data[j].max[i] > t_object.max)
					t_object.max = t_object.data[j].max[i];

				if (t_object.data[j].min[i] < t_object.min)
					t_object.min = t_object.data[j].min[i];
			}
		}

		return t_object;
	};

	// ---------------------
	// Return public method
	// ---------------------

	return {
		// parse_init_object: 		parse_init_object,
		data_generate_scale:		data_generate_scale,
		// parse_add_data:			parse_add_data,
		// parse_date_check:			parse_date_check,
		// parse_hour_check:			parse_hour_check,
		parse_counting_file:		parse_counting_file,
		parse_occupancy_file:	parse_occupancy_file
	};
})();
