#!/usr/bin/php
<?php

//----------------------------------------------------------------------------------------------------------------------------------//
//			Eurecam Demo program of HTTP API usage with PHP curl, fell free to use in your product.
//			This is a Demo program, not a robust solution to use with your 100xx sensors to connect to concurrently.
//			This program is distributed in the hope that it will be useful,
//			but WITHOUT ANY WARRANTY; without even the implied warranty of
//			MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//----------------------------------------------------------------------------------------------------------------------------------//


//
// A simple PHP example that connect to an iComptipix/Concentrix/TinyCount and get counting data for sensor 1
// This program save a DATA file and a STATE file
// DATA file contains Entries/Exits data get after each connection.
// One data file per day is saved
// STATE file contains info since last connection and is used to compute new entry/exits since last connection
// only one state file is used
//
// USAGE: with php-cli installed run : ./http_api_demo.php --help
//
// IMPORTANT NOTE: This program assume that it will be runing at a resonably short interval (less than 5 minute) to ask sensor counting data (TinyCount or iComptipix)
// NOTE2: for iComptipix and Concentrix you can ask them counting file, but this can be used if you don't want to depend on writing on SDcard schedule
//

// Parameters
// -------------------------------------------------------------------------------

$CPX3_PROTOCOL=true; // true to use ComptipixV3 protocol (using token), false to use old protocol (Concetrix/iComptipix/TinyCount/Affix)

$id = '192.168.0.141';		// TODO adapt to the IP/DNS you want + I use the IP as identifiant, you may want to do something better
$url = 'http://'.$id.'/'; // URL to connect to

// Login parameter
$user = 'user';	// we log as user to get data : don't change this value : admin level is overkill to just get data
$pass = 'user';	// TODO adapt this if you changed user password
$do_logout = false;


// Basic arg reader
// -------------------------------------------------------------------------------

$arg_longopts  = array(
	'address:',		// Value needed
	'id:',   		// Value needed
	'pass:',			// Value needed
	'user:',			// Value needed
	'old',			// Value optionnal
	'logout',		// Value optionnal
	'help',			// Value optionnal
);
$arg_options = getopt('a:i:p:u:o::l::h::', $arg_longopts); // address - id - pass - user - old - logout - help

// Set parameters
foreach ($arg_options as $arg_key => $arg_value)
{
	$id_defined = false;
	if ( ('a' == $arg_key) || ('address' == $arg_key) )
	{
		// echo $arg_key.' = '.$arg_value.' !'.PHP_EOL;
		$url = $arg_value;

		// Check there is 'http://' or 'HTTPS://' at the beginings of url, add it if not
		$url_http = false;
		if (0 === strpos($url, 'http://'))
			$url_http = true;
		if (0 === strpos($url, 'HTTP://'))
			$url_http = true;
		if (0 === strpos($url, 'https://'))
			$url_http = true;
		if (0 === strpos($url, 'HTTPS://'))
			$url_http = true;
		if (false === $url_http)
			$url = 'http://'.$url.'/';

		// Check string ends with '/' ... because for some reasoon I can't explain PHPcurl do error otherwise
		$url_last_char = mb_substr($url, -1); // why mb_substr() ?  →  because php we may have UTF8 char ... the real question is Why php keep old useless funtion like substr() ?
		if ('/' != $url_last_char)
			$url .= '/';

		// Set id with url (if id is not set already)
		if (false == $id_defined)// Don't want url to overide id, if we have defined it
			$id = preg_replace('/[^A-Za-z0-9,.]/', '', $url);// remove all char other than numerical or alphabetiacl char
	}
	else if ( ('i' == $arg_key) || ('id' == $arg_key) )
	{
		// echo $arg_key.' = '.$arg_value.' !'.PHP_EOL;
		$id = $arg_value;
		$id_defined = true;
	}
	else if ( ('p' == $arg_key) || ('pass' == $arg_key) )
	{
		// echo $arg_key.' = '.$arg_value.' !'.PHP_EOL;
		$pass = $arg_value;
	}
	else if ( ('u' == $arg_key) || ('user' == $arg_key) )
	{
		// echo $arg_key.' = '.$arg_value.' !'.PHP_EOL;
		$user = $arg_value;
	}
	else if ( ('o' == $arg_key) || ('old' == $arg_key) )
	{
		// echo $arg_key.' = true !'.PHP_EOL;
		$CPX3_PROTOCOL = false;
	}
	else if ( ('l' == $arg_key) || ('logout' == $arg_key) )
	{
		// echo $arg_key.' = true !'.PHP_EOL;
		$do_logout = true;
	}
	else if ( ('h' == $arg_key) || ('help' == $arg_key) )
	{
		echo 'This program build counting files with timestamped data, by doing HTTP request'.PHP_EOL;
		echo 'It does not get files saved in SD card !! (so it can work with a TinyCount or an Affix, that do not has SD card)'.PHP_EOL;
		echo 'It get only new data since last connection to sensor and do a sum'.PHP_EOL;
		echo 'If you want a counting data resolution of 1 minute, you must call the program each minute; 5 minutes if you want 5 ...etc'.PHP_EOL;
		echo PHP_EOL.'--- HELP: ---'.PHP_EOL;
		echo '-a or -address : define the url address to connect to (like "http://192.168.0.141/")'.PHP_EOL;
		echo '-i or -id : define the id used to save file, this must be unique for each sensor (DEFAULT: id is address)'.PHP_EOL;
		echo '-p or -pass : define the user password (DEFAULT: pass is "user")'.PHP_EOL;
		echo '-u or -user : define the user level (DEFAULT: user is "user")'.PHP_EOL; // Note that eurecam product have only "user" and "admin" level (+ comptipixV3 has "reader") --> "user" is suffisent to read data
		echo '-o or -old : to use old protocol (Affix/Concentrix/TinyCount/iComptpix) (DEFAULT: no old)'.PHP_EOL; // Note that this will cause to always try POST request with user/password
		echo '-l or -logout : use it to do a logout after all request (DEFAULT: no logout)'.PHP_EOL; // Note that this will cause to always try POST request with user/password
		echo '-h or -help : print this help'.PHP_EOL.PHP_EOL;
		echo 'USAGE EXAMPLE:'.PHP_EOL;
		echo './http_api_demo.php -a http://192.168.0.141/ -p pass_for_user'.PHP_EOL;
		echo PHP_EOL.'NOTE: Limitation:'.PHP_EOL;
		echo 'This is a demo program distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.'.PHP_EOL;
		echo 'Please note 2 limitations:'.PHP_EOL;
		echo '* for a concentrix, this program only save first sensor counting (capteur_etat_1)'.PHP_EOL;
		echo '* for a comptipixV3 or an iComptipix, this program only save sum of all lines counting'.PHP_EOL;

		return; // if we display help we do nothing else after
	}
}



// Functions
// -------------------------------------------------------------------------------

/*!
	\brief	Check that php curl is installed

	\return	false if not, true if yes
*/
function _is_curl_ok()
{
	return function_exists('curl_version');
}

/*!
	\brief	Do a curl connection

	\param	$turl 		URL to connect
	\param	$post 		true to do post request, false for get request
	\param	$post_user	optionnal post data
	\param	$post_pass	optionnal post data
	\return	array containing server answer + CURL connection info http://php.net/manual/fr/function.curl-getinfo.php
*/
function do_curl_request($turl, $post=false, $post_user=null, $post_pass=null)
{
	// Check curl
	if (false === _is_curl_ok())
	{
		echo 'php-curl MUST BE INSTALLED !!!!!!! !!!!!!! !!!!! !!!!!!! !!!!!! !!!!!! !!!!!! '.PHP_EOL;
		echo 'apt-get install php-curl or php5-curl'.PHP_EOL;
		return false;
	}

	// Do http curl connection
	$ch_con = curl_init();
	curl_setopt($ch_con, CURLOPT_URL, $turl);					// set url to send this request
	curl_setopt($ch_con, CURLOPT_RETURNTRANSFER, true);	// Receive server response ...
	if ($post)
	{
		curl_setopt($ch_con, CURLOPT_POST, true); 				// set request to be POST
		curl_setopt($ch_con, CURLOPT_POSTFIELDS,					// set user&pass field
			http_build_query(array(										// use http build that will escape/URLencode etc our paramameters
					'user' => $post_user,
					'pass' => $post_pass,
				)
			)
		);
	}
	else
		curl_setopt($ch_con, CURLOPT_POST, false); 				// set request to be GET

	// Send query + close when done
	$ch_con_output = curl_exec($ch_con); // Send
	$ch_con_info = curl_getinfo($ch_con);// Get connexion info before closing
	curl_close($ch_con);	// Close

	return array($ch_con_output, $ch_con_info);
}



// Main
// -------------------------------------------------------------------------------

// File name
$file_state = 'STATE_'.$id.'.txt'; // STATE file : will be used to compute Entries/exits since last connection to sensor
$file_data = 'DATA_'.$id;

// this value will be used to avoid impossible amount of counting
// * 10 for a TinyCount is reasonable
// * 25 for an iComptipix is reasonable
// * 25x8 for a Concentrix is reasonable (if the Concentrix store 8 sensors)
$max_counting_per_s = 10; // TODO IMPORTANT: this value can be 25 for Comptipix/iComptipix and 25*8 for a Concentrix, but for a TinyCount 10 is already big
if ($CPX3_PROTOCOL)
	$max_counting_per_s = 25;


// Check if we need to do the post request with credential
$skeep_login = false;
$last_token = '';
$req_test = array(0,0);
if ((false == $do_logout) && (file_exists($file_state))) // For ComptipixV3 we may re-use last token
{
	if ($CPX3_PROTOCOL)
	{
		// Get last token (CPX3 only)

		// Read state file
		$f_state = fopen($file_state, 'r');
		$arr_preceding_state = explode(',', fgets($f_state));
		if (count($arr_preceding_state) > 4)
		{
			// Check last token is valid
			$req_test = do_curl_request($url.'CONFIG?tkn='.$arr_preceding_state[4]);
			if (200 == $req_test[1]['http_code'])
			{
				$last_token = $arr_preceding_state[4];
				$skeep_login = true;
			}
		}
	}
	else
	{
		// Affix/TinyCount/Concentrix/iComptipix case
		$req_test = do_curl_request($url.'CONFIG?uptime'); // Check our authentication is still valid
		if (200 == $req_test[1]['http_code'])
			$skeep_login = true;
	}
}

$req_login = array(0,0);
if ($skeep_login)
{
	if ($CPX3_PROTOCOL)
	{
		// Skip login for ComptipixV3 token re-use (CPX3 only)
		$req_login = $req_test;
		$req_login[0] = $last_token;

		// Print that we re-use same token
		echo 'SKIP login using last tkn : '.$last_token.PHP_EOL;
	}
	else
	{
		// Affix/TinyCount/Concentrix/iComptipix case
		$req_login = $req_test;

		// Print that we re-use same connection authentication
		echo 'SKIP login'.PHP_EOL;
	}
}
else
{
	// Do the login request
	$url_post = $url;
	if ($CPX3_PROTOCOL)
		$url_post = $url.'CONFIG?get_tkn'; // get a token for ComptipixV3 protocol

	$req_login = do_curl_request($url_post, true, $user, $pass);
}

// Check login success
if (200 == $req_login[1]['http_code']) // All Eurecam product answer only HTTP code 200 on success, so there is only that code to check
{
	// Set counting parameters
	$counting_param = 'capteur_etat_1'; // NOTE for a concentrix you may want to add capteur_etat_1&capteur_etat_2&capteur_etat_3&...&capteur_etat_8 , and add all sensors counting
	$tkn = '';
	if ($CPX3_PROTOCOL)
	{
		$counting_param = 'line_counting';	// ComptipixV3 counting parameter
		$tkn = $req_login[0];					// ComptipixV3 token
	}

	// Param we want
	$params = array(
		$counting_param,	// to get counting + sensor state
		'uptime'				// to get uptime -> we will use this to check if sensor rebooted
	);

	// Build request
	$req = implode('&', $params);
	$req_get = $req;
	if ($CPX3_PROTOCOL)
		$req_get = 'tkn='.$tkn.'&'.$req;//CPX3 use a token at each request

	// Get req with curl
	$req_counting = do_curl_request($url.'CONFIG?'.$req_get);

	// Check request is okay
	if (200 != $req_counting[1]['http_code'])
	{
		echo '!!! ERR in second GET request : '.$req_counting[1]['url'].' → answer HTTP code: '.$req_counting[1]['http_code'].PHP_EOL;
		return;
	}

	// Get answer as array
	$arr_answer = explode(PHP_EOL, $req_counting[0]);

	// Build an associative array param_name→[value1, value2, ...]
	$params_answer = array();
	foreach ($arr_answer as &$value)
	{
		$tmp_param_name = substr($value, 0, strpos($value, '='));
		$tmp_param_answer = substr($value, strpos($value, '=')+1);
		$params_answer[$tmp_param_name] = explode(',', $tmp_param_answer);
	}

	// DEBUG: Simple print array answer
	// echo 'ANSWER to '.$req_get.'  →  ';
	// var_dump($params_answer);

	// Get counting
	$current_e = 0; // will store entry counting since last connection
	$current_x = 0; // will store exit counting since last connection
	$sum_e = intval($params_answer[$counting_param][1]); // entry counting since boot
	$sum_x = intval($params_answer[$counting_param][2]); // exit counting since boot
	$now = time(); // will be used to compute elapsed time since last connection to sensor

	if (file_exists($file_state)) // Check this is not the first time we connect to sensor
	{
		// Get last state to get current counting
		$f_state = fopen($file_state, 'r');
		if ($f_state)
		{
			// Read state file
			$arr_preceding_state = explode(',', fgets($f_state));
			$last_sum_e = intval($arr_preceding_state[2]); // Last connection entry sum
			$last_sum_x = intval($arr_preceding_state[3]); // Last connection exit sum

			// time elapsed (in second) since last connection to sensor
			$elapsed_s = $now - intval($arr_preceding_state[0]);

			// Check sensor has reboot
			if (intval($params_answer['uptime'][0]) < intval($arr_preceding_state[1])) // NOTE: no need to check for Uptime counter overflow, because Uptime counter will overflow at 4294967295 ... which is more than 136 year, there is absolutly no chance the sensor never rebooted
			{
				// Sensor rebooted since last connection
				echo 'Sensor rebooted during the last '.$elapsed_s.'s !! '.PHP_EOL;

				// Restart with this sum as current counting
				$current_e = $sum_e;	// Current entries
				$current_x = $sum_x;	// Current exits
			}
			else
			{
				// Get current entry sum
				if ($last_sum_e <= $sum_e)
					$current_e = $sum_e - $last_sum_e; // No counting overflow
				else
				{
					// On HTTP protocol Counter overflow at 2147483647 for a TinyCount/ComptipixV3 and 4294967295 for Concentrix/iComptipix
					// NOTE that for UDP all product use a short, so the overflow is 65535
					// So if it happen the new counting is OVERFLOW-last + current
					if ($last_sum_e > 2147483647)
						$current_e = 4294967295 - $last_sum_e + $sum_e; // Concentrix/iComptipix case
					else
						$current_e = 2147483647 - $last_sum_e + $sum_e; // TinyCount/ComptipixV3 case
				}

				// Get current exits sum
				if ($last_sum_x <= $sum_x)
					$current_x = $sum_x - $last_sum_x; // No counting overflow
				else
				{
					// On HTTP protocol Counter overflow at 2147483647 for a TinyCount/ComptipixV3 and 4294967295 for Concentrix/iComptipix
					// NOTE that for UDP all product use a short, so the overflow is 65535
					// So if it happen the new counting is OVERFLOW-last + current
					if ($last_sum_x > 2147483647)
						$current_x = 4294967295 - $last_sum_x + $sum_x; // Concentrix/iComptipix case
					else
						$current_x = 2147483647 - $last_sum_x + $sum_x; // TinyCount/ComptipixV3 case
				}
			}

			// Protection against suspicious too many counting since last request
			// as internet connection can be in error for some times, we should check that counting we store is possible
			$current_suspect_e = $current_e;
			$current_suspect_x = $current_x;
			if ($current_e/$elapsed_s > $max_counting_per_s) // check entries counting per second don't exeed the value defined
			{
				// Too many counting entries in a short period
				echo 'More than '.$max_counting_per_s.' counting per second ('.($current_e/$elapsed_s).') is too suspect → RESET current E'.PHP_EOL;
				$current_e = 0;
			}
			if ($current_x/$elapsed_s > $max_counting_per_s) // check exits counting per second don't exeed the value defined
			{
				// Too many counting exits in a short period
				echo 'More than '.$max_counting_per_s.' counting per second ('.($current_x/$elapsed_s).') is too suspect → RESET current X'.PHP_EOL;
				$current_x = 0;
			}
			if ( ($current_suspect_e != $current_e) || ($current_suspect_x != $current_x) )
			{
				// Counting has been cut because of suspicious counting ... on the safe side we save the suspicious counting file
				// NOTE that it shouldn't happen any time soon, it's just a precaution
				$fd_data = fopen($file_data.'_'.date('Ymd').'_suspect.csv', 'a') or die('Unable to open file '.$file_data.'_'.date('Ymd').'_suspect.csv !');
				fwrite($fd_data, date('d/m/Y,h:i:s').','.$current_suspect_e.','.$current_suspect_x.PHP_EOL); // TODO: you may want to use your favorite database to store this
				fclose($fd_data);
			}
			else
			{
				// All is OK
				echo 'current COUNTING: '.$current_e.'/'.$current_x;
				echo ' - ';
				echo 'since boot COUNTING: '.$sum_e.'/'.$sum_x.PHP_EOL;
			}

			// Close file state
			fclose($f_state);
		}
		else
			echo '!!! OOOps open file ERR'.PHP_EOL; // this is not normal
	}
	else
	{
		// This is the first counting (there is no STATE file)
		$current_e = $sum_e;
		$current_x = $sum_x;

		// Print first counting
		echo 'first COUNTING: '.$current_e.'/'.$current_x.PHP_EOL;
	}

	// Save STATE file timestamp,uptime,E,X in a file
	// NOTE: you may want to use your favorite SQL database to store this
	$fd_state = fopen($file_state, 'w') or die('Unable to open file '.$file_state.' !');
	if ( ($CPX3_PROTOCOL) && (false === $do_logout) )
		fwrite($fd_state, $now.','.$params_answer['uptime'][0].','.$sum_e.','.$sum_x.','.$tkn);// For comptipixV3 you can save token and try to re-use it at next connection (if we don't do logout)
	else
		fwrite($fd_state, $now.','.$params_answer['uptime'][0].','.$sum_e.','.$sum_x);
	fclose($fd_state);

	// Add data to DATA file
	// NOTE: you may want to use your favorite SQL database to store this
	$fd_data = fopen($file_data.'_'.date('Ymd').'.csv', 'a') or die('Unable to open file '.$file_data.'_'.date('Ymd').'.csv !');
	fwrite($fd_data, date('d/m/Y,h:i:s').','.$current_e.','.$current_x.PHP_EOL);
	fclose($fd_data);

	// Do logout
	if ($do_logout)
	{
		if ($CPX3_PROTOCOL)
			do_curl_request($url.'logout.html?tkn='.$tkn); // always use token for CPX3
		else
			do_curl_request($url.'logout.html');
	}
}
else
{
	// handle some error
	if (false === $req_login) // Curl not properly installed
		echo '!!! Ooops : Your CURL installation is not working'.PHP_EOL;
	else
	{
		// if reach here : misconfigured password $pass
		echo '!!! Ooops : login to '.$url.' FAILED (user='.$user.' & pass='.$pass.')  →  $pass may be wrong ! server answer HTTP code: '.$req_login[1]['http_code'].'  →  '.PHP_EOL;
		echo '↳ Or you are trying to connect to an iComptipix/TinyCount/Affix without -o option !'.PHP_EOL;
		echo '↳ Or you are trying to connect to a ComptipixV3 with -o option !'.PHP_EOL;
		echo ' --- --- --- --- --- --- ---'.PHP_EOL.PHP_EOL;
		echo 'DEBUG INFO : '.PHP_EOL;
		var_dump($req_login[1]);
	}
}

?>
