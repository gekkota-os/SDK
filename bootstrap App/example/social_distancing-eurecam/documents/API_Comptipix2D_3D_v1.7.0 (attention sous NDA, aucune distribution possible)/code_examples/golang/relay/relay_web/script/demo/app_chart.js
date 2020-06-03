// -----------------------------------------
// Check functions Module
// -----------------------------------------


/*!
	\brief	Module containing all function to plot or generate html table from counting or log file

*/
var app_chart = (function()
{
	'use strict';// we are strict   (see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions_and_function_scope/Strict_mode?redirectlocale=en-US&redirectslug=JavaScript%2FReference%2FFunctions_and_function_scope%2FStrict_mode)

	// ---------------------
	// Private methods
	// ---------------------


	// ---------------------
	// Public methods
	// ---------------------

	/*!
		\brief	Plot data using chartjs in container id

		\note 	This function add to charts :
		 			- better scale option
					- possibility to hide 1/2 ticks on xAxis

		\param 	container_id		the html container id
		\param 	files_src 			array of file source : first file is entry/exit, second is occupancy (Optional)
		\param 	chart_title 		null for no title																															DEFAULT: null
		\param 	conf					config to get title and message																										DEFAULT: null
		\param 	xtick_filter_nb	number of x tick to be filtered																										DEFAULT: 2
		\param 	chart_type 			chartjs type ('bar' , 'horizontalBar' , 'line' , 'bubble' , 'doughnut' , 'pie' , 'radar' , 'polarArea')		DEFAULT: 'bar'
		\return 	chartjs object created, null if nothing created
	*/
	var plot = function(container_id, files_src, chart_title, conf, xtick_filter_nb, chart_type)
	{
		// Apply default parameters
		if ('undefined' === typeof chart_type)
			chart_type = 'bar';// NOTE: the only tested for now is bar type
		if ('undefined' === typeof xtick_filter_nb)
			xtick_filter_nb = 2;
		if ('undefined' === typeof chart_title)
			chart_title = null;
		if ('undefined' === typeof conf)
			conf = null;

		// Check input (file_src)
		if (files_src.length < 1)
			return null;

		// Get plot container
		var plot_elm = document.getElementById(container_id);
		if (null === plot_elm)
			return null;

		// Set message titles
		var title_y_axis_1 = 'Counting';
		var title_y_axis_2 = 'Occupancy';
		var title_entry = 'Entry';
		var title_exit = 'Exit';
		var title_occ_max = 'Occupancy max';
		var title_occ_min = 'Occupancy min';
		if (null !== conf)
		{
			if ('undefined' !== typeof conf.titles)
			{
				if ('undefined' !== typeof conf.titles.y_axis_1)
					title_y_axis_1 = conf.titles.y_axis_1;
				if ('undefined' !== typeof conf.titles.y_axis_2)
					title_y_axis_2 = conf.titles.y_axis_2;
				if ('undefined' !== typeof conf.titles.entry)
					title_entry = conf.titles.entry;
				if ('undefined' !== typeof conf.titles.exit)
					title_exit = conf.titles.exit;
				if ('undefined' !== typeof conf.titles.occ_max)
					title_occ_max = conf.titles.occ_max;
				if ('undefined' !== typeof conf.titles.occ_min)
					title_occ_min = conf.titles.occ_min;
			}
		}

		// Parse data
		var countings = app_data_file.parse_counting_file(files_src[0]);
		var c_counting_scale = app_data_file.data_generate_scale(countings.min, countings.max);
		var occupancy = null;
		var c_occupancy_scale =  null;
		if (files_src.length > 1)
		{
			if ((null !== files_src[1])&&('undefined' !== typeof files_src[1])) // check file is not empty
			{
				occupancy = app_data_file.parse_occupancy_file(files_src[1]);
				c_occupancy_scale = app_data_file.data_generate_scale(occupancy.min, occupancy.max);
			}
		}


		// Build datasets
		// ---

		// Always add entries/exits
		var dataset_entries = {
			type: 'bar',
			label: title_entry + ' ('+countings.data[0].entry_sum+')',
			data: countings.data[0].entry,
			backgroundColor: 'rgba(153,255,51,0.4)'
		};
		var dataset_exits = {
			type: 'bar',
			label: title_exit + ' ('+countings.data[0].exit_sum+')',
			data: countings.data[0].exit,
			backgroundColor: 'rgba(255,153,0,0.4)'
		};

		// If we plot occupancy too, add y-axis-id
		if (null !== occupancy)
		{
			dataset_entries.yAxisID = 'y-axis-1';
			dataset_exits.yAxisID = 'y-axis-1';
		}

		var c_datasets = [dataset_entries, dataset_exits];

		// Add occupancy if there is a second file
		if (null !== occupancy)
		{
			// Occupancy max
			c_datasets.push({
				type: 'line',
				fill: true,
				lineTension: 0,
				label: title_occ_max + ' ('+occupancy.max+')',
				data: occupancy.data[0].max,
				borderColor: 'rgba(10,20,250,0.4)',
				backgroundColor: 'rgba(10,20,250,0.1)',
				yAxisID: 'y-axis-2'
			});

			// Occupancy min
			c_datasets.push({
				type: 'line',
				fill: true,
				lineTension: 0,
				label: title_occ_min + ' ('+occupancy.min+')',
				data: occupancy.data[0].min,
				borderColor: 'rgba(0,255,255,0.4)',
				backgroundColor: 'rgba(0,255,255,0.1)',
				yAxisID: 'y-axis-2'
			});
		}


		// Build data
		// ---

		var c_data = {
			labels: countings.tick,
			datasets: c_datasets
		};

		// Build yAxes
		// ---

		// entries/exit yAxes
		var y_axes_1 = {
			position: 'left',
			ticks: {
				max: c_counting_scale.max,
				min: c_counting_scale.min,
				stepSize: c_counting_scale.tick
			},
			scaleLabel: {
				display: true,
				labelString: title_y_axis_1
			}
		};

		// Add occupancy second y axis
		var y_axes_2 = null;
		if (null !== occupancy)
		{
			y_axes_2 = {
				position: 'right',
				ticks: {
					max: c_occupancy_scale.max,
					min: c_occupancy_scale.min,
					stepSize: c_occupancy_scale.tick,
					fontColor: 'rgba(10,20,250,1)'
				},
				gridLines: {
					drawOnChartArea: false
				},
				scaleLabel: {
					display: true,
					labelString: title_y_axis_2,
					fontColor: 'rgba(10,20,250,1)'
				}
			};

			// Set id
			y_axes_1.id = 'y-axis-1';
			y_axes_2.id = 'y-axis-2';
		}

		var y_axes = [y_axes_1];
		if (null !== y_axes_2)
			y_axes.push(y_axes_2);


		// Build option
		// ---

		// Better responisivity on all screen + scale on yAxes + better labels
		var c_options = {
			responsive: true,
			maintainAspectRatio: false, // Because the graph should adapt to any screen (this is mandatory for a client that use smarthandle in paysage mode), so we don't care of ratio
			scales: {
				yAxes: y_axes
			},
			tooltips: {
				mode: 'index',		// print all tooltip of the current index
				intersect: true,	// when mouse is hover a point or bar
				callbacks: {
					label: function(tooltipItems, data, values) {
						// Because with chartjs the datasets label is used for all label, and contains the sum
						// but we don't want to print the sum in all labels, so we remove all inside parenthesis
						var regex = new RegExp('\\([^)]*\\)', 'g'); // Remove everything inside parenthesis like : '(blahblah)'
						return data.datasets[tooltipItems.datasetIndex].label.replace(regex, '') +': ' + tooltipItems.yLabel;
					}
				}
			}
		};

		// Add fileter on xAxes
		if (xtick_filter_nb > 1)
		{
			// Add tick filter
			c_options.scales.xAxes = [{
				afterTickToLabelConversion: function(data){
					var xLabels = data.ticks;
					xLabels.forEach(function (labels, i) {
						if (i % xtick_filter_nb == 1)
							xLabels[i] = ''; // return an empty string to hide label, we can also return null to hide the grid
               });
				}
			}];
		}

		// DEBUG
		// console.log(chart_type);
		// console.log('countings =');
		// console.log(countings);
		// console.log('c_counting_scale =');
		// console.log(c_counting_scale);
		// console.log('c_datasets = ');
		// console.log(c_datasets);

		// Add title
		if (null !== chart_title)
		{
			c_options.title = {
				display: true,
				fontSize: 17,
				text: chart_title
			};
		}

		// Build graphjs object
		var ctx = plot_elm.getContext('2d');
		var chart_obj = new Chart(ctx, {
			type: chart_type,
			data: c_data,
			options: c_options
		});

		// return the chart object (need to be destroyed before a redraw, or can be used to update data)
		return chart_obj;
	};


	// ---------------------
	// Return public method
	// ---------------------

	return {
		plot:		plot
	};
})();
