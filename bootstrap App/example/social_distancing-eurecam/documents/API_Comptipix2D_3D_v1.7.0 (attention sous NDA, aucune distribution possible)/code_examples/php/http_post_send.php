<?php

//----------------------------------------------------------------------------------------------------------------------------------//
//			Eurecam Demo program receiving our POST protocol, fell free to use in your product.
//			This is a Demo program.
//			This program is distributed in the hope that it will be useful,
//			but WITHOUT ANY WARRANTY; without even the implied warranty of
//			MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//----------------------------------------------------------------------------------------------------------------------------------//

// This program is a simple example of POST send protocol no login and no basic auth support
// For a more complet example see http_post_send_full.php

/*!
 	\file				http_post_send file

	\brief			Simple example to receive data from a comptipix 2D/3D with http post or push method

 	\author			Benjamin Silvestre
	\date				2015-02-16
*/

// config

$g_data = "data";

// program

if (isset($_GET['serial']) && isset($_GET['file']) && isset($_GET['size']))
{
	$t_serial = $_GET['serial'];
	$t_file = $_GET['file'];
	$t_size = $_GET['size'];

	if (isset($_GET['type']))
		$t_dir = $_GET['serial'].'-'.$_GET['type'];
	else
		$t_dir = $_GET['serial'];

	if (isset($_GET['check']))
	{
		// TODO do something to check file

		header('Content-Length: 1'); // IMPORTANT: you must set the content length
// 		echo "2";	// skip all remaining files
// 		echo "1";	// skip this file
		echo "0";	// send this file
	}
	else if (isset($_GET['data']))
	{
		// create data directory

		if (!file_exists($g_data))
			mkdir($g_data);

		if (!file_exists($g_data."/".$t_dir))
			mkdir($g_data."/".$t_dir);

		// post data size should be equal to data in post

		$t_data = file_get_contents("php://input");
		if (strlen($t_data)==$t_size)
		{
			// save data

			$t_write = fopen($g_data."/".$t_dir."/".$t_file, 'w');
			fwrite($t_write, $t_data);
			fclose($t_write);

			header('Content-Length: 1'); // IMPORTANT: you must set the content length
			echo "1";
		}
		else
		{
			$t_err_size = "0, size error";
			header('Content-Length: '.strlen($t_err_size)); // IMPORTANT: you must set the content length
			echo $t_err_size;
		}
	}
	else
	{
		$t_err_req = "0, request error";
		header('Content-Length: '.strlen($t_err_req)); // IMPORTANT: you must set the content length
		echo $t_err_req;
	}
}
else if (isset($_GET['serial']) && isset($_GET['type']) && isset($_GET['push']))
{
	$t_dir = $_GET['serial'].'-'.$_GET['type'];

	if (isset($_GET['jpeg']))
		$t_file = date("Ymd-His").".jpeg";
	else if (isset($_GET['bmp']))
		$t_file = date("Ymd-His").".bmp";
	else
		$t_file = date("Ymd-His").".txt";

	// create data directory

	if (!file_exists($g_data))
		mkdir($g_data);

	if (!file_exists($g_data."/".$t_dir))
		mkdir($g_data."/".$t_dir);

	// write post data to file

	$t_data = file_get_contents("php://input");

	$t_write = fopen($g_data."/".$t_dir."/".$t_file, 'w');
	fwrite($t_write, $t_data);
	fclose($t_write);

	echo "OK";
}
else
{
	$t_err_param = "0, parameter error";
	header('Content-Length: '.strlen($t_err_param)); // IMPORTANT: you must set the content length
	echo $t_err_param;
}

?>
