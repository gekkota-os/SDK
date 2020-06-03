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

$CPX3_PROTOCOL=true; // true to use ComptipixV3 protocol (using token), false to use old protocol (Concetrix/iComptipix)

$id = '192.168.0.141';		// TODO adapt to the IP/DNS you want + I use the IP as identifiant, you may want to do something better
$url = 'http://'.$id.'/'; // URL to connect to

// Login parameter
$user = 'reader';	// we log as user to get data : don't change this value : admin level is overkill to just get data
$pass = 'reader';	// TODO adapt this if you changed user password
$do_logout = false;
$pass_set_by_user = false;
$file_name = '';


// Basic arg reader
// -------------------------------------------------------------------------------

$arg_longopts  = array(
	'address:',		// Value needed
	'id:',   		// Value needed
	'pass:',			// Value needed
	'user:',			// Value needed
	'file',			// Value optionnal
	'old',			// Value optionnal
	'logout',		// Value optionnal
	'help',			// Value optionnal
);
$arg_options = getopt('a:i:p:u:f:o::l::h::', $arg_longopts); // address - id - pass - user - old - logout - help

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
		$pass_set_by_user = true;
	}
	else if ( ('u' == $arg_key) || ('user' == $arg_key) )
	{
		// echo $arg_key.' = '.$arg_value.' !'.PHP_EOL;
		$user = $arg_value;
	}
	else if ( ('f' == $arg_key) || ('file' == $arg_key) )
	{
		// echo $arg_key.' = '.$arg_value.' !'.PHP_EOL;
		$file_name = $arg_value;
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
		echo 'This program get files saved on SD card'.PHP_EOL;
		echo '↳ So it can NOT work with a TinyCount or an Affix ! Because they haven\'t any SD card !'.PHP_EOL;
		echo PHP_EOL.'--- HELP: ---'.PHP_EOL;
		echo '-a or -address : define the url address to connect to (like "http://192.168.0.141/")'.PHP_EOL;
		echo '-i or -id : define the id used to save file, this must be unique for each sensor (DEFAULT: id is address)'.PHP_EOL;
		echo '-p or -pass : define the user password (DEFAULT: pass is "user")'.PHP_EOL;
		echo '-u or -user : define the user level (DEFAULT: user is "user")'.PHP_EOL; // Note that eurecam product have only "user" and "admin" level (+ comptipixV3 has "reader") --> "user" is suffisent to read data
		echo '-f or -file : define file to get (Example: 20161001.csv), if empty program will get today data'.PHP_EOL;
		echo '-o or -old : to use old protocol (Concentrix/iComptpix) (DEFAULT: no old)'.PHP_EOL; // Note that this will cause to always try POST request with user/password
		echo '-l or -logout : use it to do a logout after all request (DEFAULT: no logout)'.PHP_EOL; // Note that this will cause to always try POST request with user/password
		echo '-h or -help : print this help'.PHP_EOL.PHP_EOL;
		echo 'USAGE EXAMPLE:'.PHP_EOL;
		echo './http_api_demo.php -a http://192.168.0.141/ -p pass_for_reader'.PHP_EOL;
		echo PHP_EOL.'NOTE: Limitation:'.PHP_EOL;
		echo 'This is a demo program distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.'.PHP_EOL;

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

// Set file name if user don't pass it
if ('' == $file_name)
{
	$file_name =  date("Ymd").'.csv';
	echo 'Get default today file : '.$file_name.PHP_EOL;
}

// File name
$file_state = 'STATE_'.$id.'.txt'; // STATE file : will be used to store token used in last connection to sensor

// Adapt password/login to old protocol (for iComptipix/Concentrix)
if (!$CPX3_PROTOCOL)
{
	// Warn that we are using the old protocol
	echo 'Using old protocol (Concentrix/iComptipix) only !!!'.PHP_EOL;

	// For old protocol, we don't have reader level, so use the user level
	if ('reader' == $user)
	{
		$user = 'user';	// we log as user to get data : don't change this value : admin level is overkill to just get data
		if (!$pass_set_by_user)
			$pass = 'user';
	}
}


// Check if we need to do the post request with credential
$skip_login = false;
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
		if (count($arr_preceding_state) > 1)
		{
			// Check last token is valid
			$req_test = do_curl_request($url.'CONFIG?tkn='.$arr_preceding_state[1]);
			if (200 == $req_test[1]['http_code'])
			{
				$last_token = $arr_preceding_state[1];
				$skip_login = true;
			}
		}
	}
	else
	{
		// Concentrix/iComptipix case
		$req_test = do_curl_request($url.'CONFIG?uptime'); // Check our authentication is still valid
		if (200 == $req_test[1]['http_code'])
			$skip_login = true;
	}
}

$req_login = array(0,0);
if ($skip_login)
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
		// Concentrix/iComptipix case
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
	// echo 'LOGIN : '.$url_post.' : '.$user.'/'.$pass.' : '.$req_login[0].PHP_EOL;
}

// Check login success
if (200 == $req_login[1]['http_code']) // All Eurecam product answer only HTTP code 200 on success, so there is only that code to check
{
	// Set counting parameters
	$read_param[0] = 'FICHIER?info='.$file_name;
	$read_param[1] = 'FICHIER?lecture='.$file_name;
	$tkn = '';
	if ($CPX3_PROTOCOL)
	{
		$tkn = $req_login[0];					// ComptipixV3 token
		$read_param[0] = 'CONFIG?tkn='.$tkn.'&sdcard_info='.$file_name;
		$read_param[1] = 'CONFIG?tkn='.$tkn.'&sdcard_read='.$file_name;
	}

	// Check there is file
	$req_file_info = do_curl_request($url.$read_param[0]);

	// Check request is okay
	if (200 != $req_file_info[1]['http_code'])
	{
		echo '!!! ERR in read file info GET request : '.$req_file_info[1]['url'].' → answer HTTP code: '.$req_file_info[1]['http_code'].PHP_EOL;
		return;
	}

	// Check file exist
	$arr_answer = explode('=', $req_file_info[0]);
	if (intval($arr_answer[1]) <= 0)
	{
		echo 'There is no file : '.$file_name.PHP_EOL;
		return;
	}

	// File exist : get the file
	$req_file_read = do_curl_request($url.$read_param[1]);

	// Check request is okay
	if (200 != $req_file_read[1]['http_code'])
	{
		echo '!!! ERR in GET file request : '.$req_file_read[1]['url'].' → answer HTTP code: '.$req_file_read[1]['http_code'].PHP_EOL;
		return;
	}

	// DEBUG:
	// echo 'File ('.$file_name.') read: '.$url.$read_param[0].PHP_EOL;
	// echo $req_file_read[0];

	// Save STATE file timestamp,tkn in a file
	// NOTE: you may want to use your favorite SQL database to store this
	$now = time(); // just to store date of last token
	$fd_state = fopen($file_state, 'w') or die('Unable to open file '.$file_state.' !');
	if ( ($CPX3_PROTOCOL) && (false === $do_logout) )
		fwrite($fd_state, $now.','.$tkn);// For comptipixV3 you can save token and try to re-use it at next connection (if we don't do logout)
	fclose($fd_state);

	// Save DATA file
	// NOTE: you may want to use your favorite SQL database to store this
	$fd_data = fopen($file_name, 'w') or die('Unable to open file '.$file_name.' !');
	fwrite($fd_data, $req_file_read[0]);
	fclose($fd_data);

	// Do logout
	if ($do_logout)
	{
		if ($CPX3_PROTOCOL)
			do_curl_request($url.'logout.html?tkn='.$tkn); // always use token for CPX3
		else
			do_curl_request($url.'logout.html');
	}

	// All ok
	echo 'File : '.$file_name.' saved OK'.PHP_EOL;
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
		echo '↳ Or you are trying to connect to an Concentrix/iComptipix without -o option !'.PHP_EOL;
		echo '↳ Or you are trying to connect to a ComptipixV3 with -o option !'.PHP_EOL;
		echo ' --- --- --- --- --- --- ---'.PHP_EOL.PHP_EOL;
		echo 'DEBUG INFO : '.PHP_EOL;
		var_dump($req_login[1]);
	}
}

?>
