/**
 * File : autoconf.js
 * Description :
 *   Example for a JavaScript platform auto-configuration.
 *
 *   This file must be nammed as :
 *      - "000000000000.js" for a default auto-configuration,
 *      - "<MAC>.js" for a specific auto-configuration (with <MAC> is the
 *             MAC adress of the platform with no separator character;
 *             Example : "00E04B4124DB.js").
 *  
 *   If the script fails, it's possible to activate a log.
 *   For that, add into the file "<playzilla install path>/res/log4xpcom.xml" the following section :
        <logger name="launcher.profile.addon-manager">
                <level value="DEBUG"/>
        </logger>
 *   Then, you can retrieve the log in the file 'player.log', which can be retrieved
 *   at different locations, depending on the platform type an playzilla version.
 *
 *	 To use the auto-configuration Javascript you need to have a firmware version 3.10.11 or higher
 */

const Cc = Components.classes;
const Ci = Components.interfaces;
const nsIPrefBranch=Ci.nsIPrefBranch;
Components.utils.import("resource://gre/modules/Services.jsm");

var logService = Cc["@innes/log4service;1"].getService(Ci.nsILog4Service);
var logger = new Logger("autoconf");
var systemManager = Cc["@innes/system/systemmanager;1"].getService(Ci.nsISystemManager);
logger.debug("Autoconf start");

var Platform=undefined;
setInitGPIO();

// ---------------------------------------
// ---------------------------------------
// ---- BEGIN of the user configuration
// ---------------------------------------
// ---------------------------------------

// Select the GPIOs you want to configure, and comment/uncomment the correct lines


// GPIO1 : uncomment one line below corresponding to your choice 
setPhoenixGPIO("1", false, "in",  100000000);   // GPIO1 : Input, Debouncing 100ms 
//setPhoenixGPIO("1", false, "in",  0);         // GPIO1 : Input, No Debouncing 
//setPhoenixGPIO("1", true,  "in",  100000000); // GPIO1 : Input, Invert, Debouncing 100ms 
//setPhoenixGPIO("1", true,  "in",  0);         // GPIO1 : Input, Invert, No Debouncing 
//setPhoenixGPIO("1", false, "out", 0);         // GPIO1 : Output 
//setPhoenixGPIO("1", false, "disable", 0);     // GPIO1 : Disable

// GPIO2 : uncomment one line below corresponding to your choice 
setPhoenixGPIO("2", false, "in",  100000000);   // GPIO2 : Input, Debouncing 100ms 
//setPhoenixGPIO("2", false, "in",  0);         // GPIO2 : Input, No Debouncing 
//setPhoenixGPIO("2", true,  "in",  100000000); // GPIO2 : Input, Invert, Debouncing 100ms 
//setPhoenixGPIO("2", true,  "in",  0);         // GPIO2 : Input, Invert, No Debouncing 
//setPhoenixGPIO("2", false, "out", 0);         // GPIO2 : Output 
//setPhoenixGPIO("2", false, "disable", 0);     // GPIO2 : Disable

// GPIO3 : uncomment one line below corresponding to your choice 
setPhoenixGPIO("3", false, "in",  100000000);   // GPIO3 : Input, Debouncing 100ms 
//setPhoenixGPIO("3", false, "in",  0);         // GPIO3 : Input, No Debouncing 
//setPhoenixGPIO("3", true,  "in",  100000000); // GPIO3 : Input, Invert, Debouncing 100ms 
//setPhoenixGPIO("3", true,  "in",  0);         // GPIO3 : Input, Invert, No Debouncing 
//setPhoenixGPIO("3", false, "out", 0);         // GPIO3 : Output 
//setPhoenixGPIO("3", false, "disable", 0);     // GPIO3 : Disable

// GPIO4 : uncomment one line below corresponding to your choice 
setPhoenixGPIO("4", false,  "in", 100000000);   // GPIO4 : Input, Debouncing 100ms 
//setPhoenixGPIO("4", false, "in",  0);         // GPIO4 : Input, No Debouncing 
//setPhoenixGPIO("4", true,  "in",  100000000); // GPIO4 : Input, Invert, Debouncing 100ms 
//setPhoenixGPIO("4", true,  "in",  0);         // GPIO4 : Input, Invert, No Debouncing 
//setPhoenixGPIO("4", false, "out", 0);         // GPIO4 : Output 
//setPhoenixGPIO("4", false, "disable", 0);     // GPIO4 : Disable

// GPIO5 : uncomment one line below corresponding to your choice 
setPhoenixGPIO("5", false, "in",  100000000);   // GPIO5 : Input, Debouncing 100ms 
//setPhoenixGPIO("5", false, "in",  0);         // GPIO5 : Input, No Debouncing 
//setPhoenixGPIO("5", true,  "in",  100000000); // GPIO5 : Input, Invert, Debouncing 100ms 
//setPhoenixGPIO("5", true,  "in",  0);         // GPIO5 : Input, Invert, No Debouncing 
//setPhoenixGPIO("5", false, "out", 0);         // GPIO5 : Output 
//setPhoenixGPIO("5", false, "disable", 0);     // GPIO5 : Disable

// GPIO6 : uncomment one line below corresponding to your choice 
setPhoenixGPIO("6", false, "in",  100000000);   // GPIO6 : Input, Debouncing 100ms 
//setPhoenixGPIO("6", false, "in",  0);         // GPIO6 : Input, No Debouncing 
//setPhoenixGPIO("6", true,  "in",  100000000); // GPIO6 : Input, Invert, Debouncing 100ms 
//setPhoenixGPIO("6", true,  "in",  0);         // GPIO6 : Input, Invert, No Debouncing 
//setPhoenixGPIO("6", false, "out", 0);         // GPIO6 : Output 
//setPhoenixGPIO("6", false, "disable", 0);     // GPIO6 : Disable

// Jack GPIO : uncomment one line below corresponding to your choice 
setJackGPIO(false, "in",  100000000);   		// Jack GPIO : Input, Debouncing 100ms 
//setJackGPIO(false, "in",  0);         		// Jack GPIO : Input, No Debouncing 
//setJackGPIO(true,  "in",  100000000); 		// Jack GPIO : Input, Invert, Debouncing 100ms 
//setJackGPIO(true,  "in",  0);         		// Jack GPIO : Input, Invert, No Debouncing 
//setJackGPIO(false, "out", 0);         		// Jack GPIO : Output 
//setJackGPIO(false, "disable", 0);     		// Jack GPIO : Disable


// ---- Save any previous preference changed : DO NOT COMMENT THE FOLLOWING LINE !!!
Services.prefs.savePrefFile(null);


// ---------------------------------------
// ---------------------------------------
// ---- END of the user configuration
// ---------------------------------------
// ---------------------------------------



// ---- Functions section

/**
* Init somes GPIOs to avoid conflict
*/
function setInitGPIO()
{
	var systemPref = Cc["@innes/systemprefservice;1"].getService(Ci.nsIPrefBranch);
	var generalSettings = Cc["@innes/system/generalsettings;1"].getService(Ci.nsISystemGeneralSettings);
    Platform = generalSettings.platform;
	logger.debug("setInitGPIO : Platform = " + Platform);
	if (Platform == "dmc200") 
	{
		systemPref.setBoolPref("system/connector/phoenix_1/3/rs485/uart_1/enable", false);
		systemPref.setBoolPref("system/connector/phoenix_1/15/rs422/uart_1/enable", false);
	}
	else if (Platform == "smt210")
	{
		Services.prefs.setBoolPref("innes.app-profile.gpio-input.phoenix-gpio_1.phoenix_1.1.authorized", false);
		Services.prefs.setBoolPref("innes.app-profile.gpio-output.phoenix-gpio_1.phoenix_1.1.authorized", true);
		systemPref.setBoolPref("system/connector/phoenix_1/1/io/phoenix-gpio_1/enable", true);
	}
}


/**
* Set the GPIO on Phoenix connector
* @param aPort : GPIO number
* @param aInvert : true to invert data input
* @param aDirection : direction on the GPIO1, in for input, out for output, disable to don't set the drection
* @param aDebouncing : to set the duration of the debouncing (in ns), maximum=10000000000 (10s) mimimum=0 (no deboucning)
*/
function setPhoenixGPIO(aPort, aInvert, aDirection, aDebouncing)   
{
	if (aPort == 1)
	{
		PortNumber = 1;
	}
	else if (aPort == 2)
	{
		PortNumber = 2;
	}
	else if (aPort == 3)
	{
		PortNumber = 4;
	}
	else if (aPort == 4)
	{
		PortNumber = 8;		
	}
	else if (aPort == 5)
	{
		PortNumber = 16;		
	}	
	else if (aPort == 6)
	{
		PortNumber = 32;		
	}
	
	if ((Platform == "dmc200") && (aPort >= 1) && (aPort <= 4))
	{
		setPhoenixGPIOfunction(PortNumber, aInvert, aDirection, aDebouncing);
	}
	
	if ((Platform == "dmb300") && (aPort >= 1) && (aPort <= 6))
	{
		setPhoenixGPIOfunction(PortNumber, aInvert, aDirection, aDebouncing);
	}
	
	if ((Platform == "smt210") && (aPort >= 2) && (aPort <= 3))
	{
		setPhoenixGPIOfunction(PortNumber, aInvert, aDirection, aDebouncing);
	}
}

function setPhoenixGPIOfunction(aPort, aInvert, aDirection, aDebouncing)   
{
	var systemPref = Cc["@innes/systemprefservice;1"].getService(Ci.nsIPrefBranch);

	if (Platform == "dmc200")
	{
		//Set the direction : input or output
		if (aDirection == "out")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.io_expander_1.phoenix_1." + aPort + ".authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.io_expander_1.phoenix_1." + aPort + ".authorized", true);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/io_expander_1/enable", true);
		}
		else if (aDirection == "in")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.io_expander_1.phoenix_1." + aPort + ".authorized", true);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.io_expander_1.phoenix_1." + aPort + ".authorized", false);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/io_expander_1/enable", true);		
		}
		else if (aDirection == "disable")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.io_expander_1.phoenix_1." + aPort + ".authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.io_expander_1.phoenix_1." + aPort + ".authorized", false);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/io_expander_1/enable", false);
		}
		//Set invert or not
		Services.prefs.setBoolPref("innes.app-profile.gpio-input.io_expander_1.phoenix_1." + aPort + ".invert-value", aInvert);
		// Set the debouncing time
		if (aDirection == "in") 
		{
			systemPref.setIntPref("system/connector/phoenix_1/" + aPort + "/io/io_expander_1/debouncing/time", aDebouncing);
			if (aDebouncing == 0)
			{
				systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/io_expander_1/debouncing/enable", false);
			}
			else
			{
				systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/io_expander_1/debouncing/enable", true);
			}
		}
		else
		{
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/io_expander_1/debouncing/enable", false);
		}
	}
	else if (Platform == "dmb300")
	{
		//Set the direction : input or output
		if (aDirection == "out")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.phoenix_1." + aPort + ".authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.phoenix_1." + aPort + ".authorized", true);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/enable", true);
		}
		else if (aDirection == "in")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.phoenix_1." + aPort + ".authorized", true);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.phoenix_1." + aPort + ".authorized", false);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/enable", true);		
		}
		else if (aDirection == "disable")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.phoenix_1." + aPort + ".authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.phoenix_1." + aPort + ".authorized", false);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/enable", false);
		}
		//Set invert or not
		Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.phoenix_1." + aPort + ".invert-value", aInvert);
		// Set the debouncing time
		if (aDirection == "in") 
		{
			systemPref.setIntPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/debouncing/time", aDebouncing);
			if (aDebouncing == 0)
			{
				systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/debouncing/enable", false);
			}
			else
			{
				systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/debouncing/enable", true);
			}
		}
		else
		{
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/debouncing/enable", false);
		}
	}
	else if (Platform == "smt210")
	{
			//Set the direction : input or output
		if (aDirection == "out")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.phoenix-gpio_1.phoenix_1." + aPort + ".authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.phoenix-gpio_1.phoenix_1." + aPort + ".authorized", true);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/phoenix-gpio_1/enable", true);
		}
		else if (aDirection == "in")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.phoenix-gpio_1.phoenix_1." + aPort + ".authorized", true);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.phoenix-gpio_1.phoenix_1." + aPort + ".authorized", false);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/phoenix-gpio_1/enable", true);		
		}
		else if (aDirection == "disable")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.phoenix-gpio_1.phoenix_1." + aPort + ".authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.phoenix-gpio_1.phoenix_1." + aPort + ".authorized", false);
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/phoenix-gpio_1/enable", false);
		}
		//Set invert or not
		Services.prefs.setBoolPref("innes.app-profile.gpio-input.phoenix-gpio_1.phoenix_1." + aPort + ".invert-value", aInvert);
		// Set the debouncing time
		if (aDirection == "in") 
		{
			systemPref.setIntPref("system/connector/phoenix_1/" + aPort + "/io/phoenix-gpio_1/debouncing/time", aDebouncing);
			if (aDebouncing == 0)
			{
				systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/phoenix-gpio_1/debouncing/enable", false);
			}
			else
			{
				systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/phoenix-gpio_1/debouncing/enable", true);
			}
		}
		else
		{
			systemPref.setBoolPref("system/connector/phoenix_1/" + aPort + "/io/epld_1/debouncing/enable", false);
		}
	}
}


/**
* Set the GPIO on Phoenix connector
* @param aInvert : true to invert data input
* @param aDirection : direction on the GPIO, in for input, out for output, disable to don't set the drection
* @param aDebouncing : to set the duration of the debouncing (in ns), maximum=10000000000 (10s) mimimum=0 (no deboucning)
*/
function setJackGPIO(aInvert, aDirection, aDebouncing) 
{	
	var systemPref = Cc["@innes/systemprefservice;1"].getService(Ci.nsIPrefBranch);
	
	if (Platform == "dmc200" || Platform == "dmb300")
	{
		if (Platform == "dmc200")
		{
			systemPref.setBoolPref("system/connector/jack35_1/1/io/uart_1/enable", false);
		}
		else if (Platform == "dmb300")
		{
			systemPref.setBoolPref("system/connector/jack35_1/1/io/uart_2/enable", false);
		}
		// Set the direction : input or output
		if (aDirection == "out")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.jack35_1.*.authorized", true);
			systemPref.setBoolPref("system/connector/jack35_1/1/io/epld_1/enable", true);
		}
		else if (aDirection == "in")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.authorized", true);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.jack35_1.*.authorized", false);
			systemPref.setBoolPref("system/connector/jack35_1/1/io/epld_1/enable", true);			
		}
		else if (aDirection == "disable")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.jack35_1.*.authorized", false);
			systemPref.setBoolPref("system/connector/jack35_1/1/io/epld_1/enable", false);
		}
		// Set invert or not
		Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.invert-value", aInvert);
		// Set the debouncing time
		if (aDirection == "in") 
		{
			systemPref.setIntPref("system/connector/jack35_1/1/io/epld_1/debouncing/time", aDebouncing);
			if (aDebouncing == 0)
			{
				systemPref.setBoolPref("system/connector/jack35_1/1/io/epld_1/debouncing/enable", false);
			}
			else
			{
				systemPref.setBoolPref("system/connector/jack35_1/1/io/epld_1/debouncing/enable", true);
			}
		}
		else
		{
			systemPref.setBoolPref("system/connector/jack35_1/1/io/epld_1/debouncing/enable", false);
		}	
	}
	else if (Platform == "smt210" || Platform == "sma200")
	{
		systemPref.setBoolPref("system/connector/jack35_1/1/io/uart_1/enable", false);
		// Set the direction : input or output
		if (aDirection == "out")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.jack35-gpio_1.jack35_1.*.authorized", true);
			systemPref.setBoolPref("system/connector/jack35_1/1/io/jack35-gpio_1/enable", true);
		}
		else if (aDirection == "in")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.authorized", true);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.jack35-gpio_1.jack35_1.*.authorized", false);
			systemPref.setBoolPref("system/connector/jack35_1/1/io/jack35-gpio_1/enable", true);			
		}
		else if (aDirection == "disable")
		{
			Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.authorized", false);
			Services.prefs.setBoolPref("innes.app-profile.gpio-output.jack35-gpio_1.jack35_1.*.authorized", false);
			systemPref.setBoolPref("system/connector/jack35_1/1/io/jack35-gpio_1/enable", false);
		}
		// Set invert or not
		Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.invert-value", aInvert);
		// Set the debouncing time
		if (aDirection == "in") 
		{
			systemPref.setIntPref("system/connector/jack35_1/1/io/jack35-gpio_1/debouncing/time", aDebouncing);
			if (aDebouncing == 0)
			{
				systemPref.setBoolPref("system/connector/jack35_1/1/io/jack35-gpio_1/debouncing/enable", false);
			}
			else
			{
				systemPref.setBoolPref("system/connector/jack35_1/1/io/jack35-gpio_1/debouncing/enable", true);
			}
		}
		else
		{
			systemPref.setBoolPref("system/connector/jack35_1/1/io/jack35-gpio_1/debouncing/enable", false);
		}	
	}
}
