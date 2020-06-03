<?php

//----------------------------------------------------------------------------------------------------------------------------------//
//			Eurecam Demo program receiving our POST protocol, fell free to use in your product.
//			This is a Demo program.
//			This program is distributed in the hope that it will be useful,
//			but WITHOUT ANY WARRANTY; without even the implied warranty of
//			MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//----------------------------------------------------------------------------------------------------------------------------------//

// This program is a complete example of POST send protocol with login and basic auth support
// For a smaller exemple easier to understand see http_post_send.php

// Global configuration
$g_data_parent_dir = "../data"; 					// or "./data" to write directly inside this dir
$g_check_file_size = true;							// true to check file size : in this case file with same size will be skipped
$g_use_type_info = true;							// true to use type information send since version 1.3.0
$g_type_default = "CPX3";							// if type info is used, even a Comptipix version < 1.3.0 will use CPX3
$g_type_serial_separator = "-";					// the separator to use for type-serial directory for example here result will be like: 'CPX3-119063'
$g_log_file = $g_data_parent_dir."/log.log";	// empty to not use a log file
$g_log_skip = false;									// true to log when a file send is skipped, false otherwise
$g_skip_all_remaning_files = true;				// true to skip all remaining file when first file size is the same (support since Comptipix version 1.3.0)
$g_use_basic_auth = false;							// true to use basic auth, false otherwise
$g_basic_auth_user = "user";						// Basic auth user sender must set (if $g_use_basic_auth is true), change it to suit your need
$g_basic_auth_pass = "pass";						// Basic auth pass sender must set (if $g_use_basic_auth is true), change it to suit your need
$g_basic_auth_log_err = true;						// true to log basic auth error, false otherwise

/*!
	\brief			Force a basic auth

*/
function force_authenticate() {
	// Send header to force Basic auth
	$t_err_msg = "0, You need to provide a valid User/Pass basic auth.\n";
	header('WWW-Authenticate: Basic realm="Test Authentication System"');
	header('HTTP/1.0 401 Unauthorized');
	header('Content-Length: '.strlen($t_err_msg)); // IMPORTANT: you must set the content length
	echo $t_err_msg;
	exit; // if there is error on login/pass, stop here
}

/*!
 	\file				http_post_send file

	\brief			Simple example to receive data from a comptipix V3 with http post method

 	\author			Benjamin Silvestre
	\date				2015-02-16
*/
if (isset($_GET['serial']) && isset($_GET['file']) && isset($_GET['size']))
{
	// Get info sent
	$t_serial = $_GET['serial'];
	$t_file = $_GET['file'];
	$t_size = $_GET['size'];
	$t_type = "";// Default type empty
	if ($g_use_type_info)
	{
		$t_type = $g_type_default.$g_type_serial_separator;
		if (isset($_GET['type']))
			$t_type = $_GET['type'].$g_type_serial_separator;
	}

	// Check Basic auth (if used)
	if ($g_use_basic_auth)
	{
		// Check a basic auth is provided
		if ((!isset($_SERVER['PHP_AUTH_USER'])) || ($_SERVER['PHP_AUTH_USER'] != $g_basic_auth_user) || ($_SERVER['PHP_AUTH_PW'] != $g_basic_auth_pass))
		{
			// Log error
			if (("" !== $g_log_file) && ($g_basic_auth_log_err))
			{
				$t_log_err_f = fopen($g_log_file, 'a');
				if (false !== $t_log_err_f)
				{
					$t_log_err = date("Y-m-d H:i:s") . ", Sensor:" . $t_type.$t_serial . " -> Error auth:";
					if (!isset($_SERVER['PHP_AUTH_USER']))
						$t_log_err .= "Missing basic auth header".PHP_EOL;
					else
						$t_log_err .= "Wrong user/pass -> (" .$_SERVER['PHP_AUTH_USER']. "/" .$_SERVER['PHP_AUTH_PW']. ")".PHP_EOL;
					fwrite($t_log_err_f, $t_log_err);
					fclose($t_log_err_f);
				}
			}

			// Force auth and exit
			force_authenticate();
		}
	}

	if (isset($_GET['check']))
	{
		// Check file size if enabled
		$t_result = "0";
		$t_f_size = 0;
		if ($g_check_file_size)
		{
			if (!file_exists($g_data_parent_dir)) // data directory not created
				$t_result = "0"; // send this file
			else if (!file_exists($g_data_parent_dir."/".$t_type.$t_serial)) // serial sub-directory not created
				$t_result = "0"; // send this file
			else
			{
				$t_f_size = filesize($g_data_parent_dir."/".$t_type.$t_serial."/".$t_file);
				if ($t_f_size == $t_size)
				{
					$t_result = "1"; // skip file (same size)
					if ($g_skip_all_remaning_files)
						$t_result = "2"; // skip all remaining files
				}
				else
					$t_result = "0"; // send file (size different)
			}
		}
		else
		{
			// Configured to not check file size
			$t_result = "0";// send this file
		}

		// Log send command
		if ("" !== $g_log_file)
		{
			$t_log = date("Y-m-d H:i:s") . ", Sensor:" . $t_type.$t_serial . " -> Check File(size):" . $t_file . "(".$t_size.") <- ";
			if ($g_check_file_size)
			{
				if ($t_f_size == $t_size)
					$t_log .= $t_result.":skip file, same size";
				else if ($t_f_size > 0)
					$t_log .= $t_result.":send file, size different (" . $t_f_size . ")";
				else
					$t_log .= $t_result.":send file, file not present";
			}
			else
				$t_log .= "send file, always";

			$t_log .= PHP_EOL;

			// Write a log file
			if ( ("0" == $t_result) || (true === $g_log_skip) )
			{
				$t_log_f = fopen($g_log_file, 'a');
				if (false !== $t_log_f)
				{
					fwrite($t_log_f, $t_log);
					fclose($t_log_f);
				}
			}
		}

		// NOTE answer possible are :
		// echo "2";	// skip all remaining files
		// echo "1";	// skip this file
		// echo "0"; 	// send this file

		// Send result answer
		header('Content-Length: '.strlen($t_result)); // IMPORTANT: you must set the contant length
		echo $t_result;
	}
	else if (isset($_GET['data']))
	{
		// Log save operation
		$t_receive_result = "0";
		$t_log_f_receive = false; // No log by default
		if ("" !== $g_log_file)
			$t_log_f_receive = fopen($g_log_file, 'a');

		// Create data directory
		if (!file_exists($g_data_parent_dir))
			mkdir($g_data_parent_dir);

		if (!file_exists($g_data_parent_dir."/".$t_type.$t_serial))
			mkdir($g_data_parent_dir."/".$t_type.$t_serial);

		// Post data size should be equal to data in post
		$t_data = file_get_contents("php://input");
		if (strlen($t_data)==$t_size)
		{
			// Save data
			$t_write = fopen($g_data_parent_dir."/".$t_type.$t_serial."/".$t_file, 'w');
			$t_write_result = fwrite($t_write, $t_data);
			$t_close_result = fclose($t_write);

			// Check data saved
			if ((false !== $t_write_result) && (false !== $t_close_result))
				$t_receive_result = "1"; // file succefully saved
			else
				$t_receive_result = "0, save error"; // error in saving file

			// Log operation
			if (false !== $t_log_f_receive)
			{
				$t_log_receive = date("Y-m-d H:i:s") . ", Sensor:" . $t_type.$t_serial . " <- Received File(size):" . $t_file . "(".$t_size."), OK".PHP_EOL;
				if ("1" !== $t_receive_result)
					$t_log_receive = date("Y-m-d H:i:s") . ", Sensor:" . $t_type.$t_serial . " <- Received File(size):" . $t_file . "(".$t_size."), ERROR save file failed".PHP_EOL;

				// if (false === strstr($t_file, ".jpg")) // DEBUG: Very verbose
				// 	$t_log_receive .= $t_data.PHP_EOL."---------------------".PHP_EOL; // DEBUG: Very verbose -> write saved file in log

				fwrite($t_log_f_receive, $t_log_receive);
			}
		}
		else
		{	// File size Error
			// Log operation
			if (false !== $t_log_f_receive)
			{
				$t_log_receive = date("Y-m-d H:i:s") . ", Sensor:" . $t_type.$t_serial . " <- Received File(size):" . $t_file . "(".$t_size."), ERROR file size error (size should be: ".strlen($t_data).")".PHP_EOL;
				fwrite($t_log_f_receive, $t_log_receive);
			}
		}

		// Close log
		if (false !== $t_log_f_receive)
			fclose($t_log_f_receive);

		// Send result
		header('Content-Length: '.strlen($t_receive_result)); // IMPORTANT: you must set the content length
		echo $t_receive_result;
	}
	else
	{
		$t_error_req = "0, request error";
		header('Content-Length: '.strlen($t_error_req)); // IMPORTANT: you must set the content length
		echo $t_error_req;
	}
}
else
{
	$t_error_param = "0, parameter error";
	header('Content-Length: '.strlen($t_error_param)); // IMPORTANT: you must set the content length
	echo $t_error_param;
}

?>
