/**
 * Fonction: configuration-by-script
 * File Base: example_000000000000-V1.11.26.js
 * File name pattern
 *    The file name must match on of these these formats (ex: USB injection):
 *      - 000000000000.js
 * 		- <MAC>.js (upper case only for <MAC>)
 * 			ex: 001CE602FC01.js
 *    In PlugnCast G3 configuration, the file name must match on of these these formats :
 * 		- <Mac>.configuration.js (lower case as well as upper case for <MAC>)
 * 			ex: 00-1C-E6-02-FC-01.configuration.js or 00-1c-e6-02-fc-01.configuration.js
 * 		- <hostname>.configuration.js (lower case as well as upper case for <hostname>)
 * 		- <UUID>.configuration.js (lower case as well as upper case for <UUID>)
 * 		- configuration.js
 *
 *   If the script installation fails, go in player WebUI <player_IP>/.status/, open status.xml, and look at the error raised
 *
 */
 
const Cc = Components.classes;
const Ci = Components.interfaces;
const nsIPrefBranch = Ci.nsIPrefBranch;
Components.utils.import("resource://gre/modules/Services.jsm");

var logService = Cc["@innes/log4service;1"].getService(Ci.nsILog4Service);
var logger = new Logger("autoconf");
var systemManager = Cc["@innes/system/systemmanager;1"].getService(Ci.nsISystemManager);
logger.debug("Autoconf start");

// ---------------------------------------
// ---------------------------------------
// ---- BEGIN of the user configuration
// ---------------------------------------
// ---------------------------------------

var Platform = getPlatform();

// ----------------------------------------------------------------------------------
// ---- AVCmd: enable or disable DDC-CI profile on HDMI output
// ----------------------------------------------------------------------------------
let avCmdDdcCi = AvCmdGetProfileWithInstance("i2c_1", "");
AvCmdActivateProfile(avCmdDdcCi);

// ---- AVCmd: enable or disable DDC-CI features
AvCmdIsActivatedBrightness(avCmdDdcCi, true);

// ---- Save any previous preference changed : DO NOT COMMENT THE FOLLOWING LINE !!!
Services.prefs.savePrefFile(null);

// ---------------------------------------
// ---------------------------------------
// ---- END of the user configuration
// ---------------------------------------
// ---------------------------------------

/**
* Retrieve a "av-cmd" profile for the given device id and instance name.
* @param aDeviceId: (exemple "network" or "uart_<n>"
* @return the entry of profile or null
*/
function AvCmdGetProfileWithInstance(aDeviceId, aInstanceName) {
	if (isG4()) {
		let profiles = systemManager.instantiateApplicationProfileBindings(
			"av-cmd","", aDeviceId,"","");
		if (profiles.length > 0) {
			logger.debug("Has profile");
			let entry = profiles.queryElementAt(0, Ci.nsISystemAPBEntry);
			return entry;
		}
	}
	return null;
}

/**
* Activate a "av-cmd" profile.
* @param aProfileEntry: the entry for the "av-cmd" profile.
*/
function AvCmdActivateProfile(aProfileEntry) {
    logger.debug("Activate profile");
	if(aProfileEntry != null)
	{
		aProfileEntry.activated = true;
	}
}

function AvCmdIsActivatedBrightness(aProfileEntry, aActivated)
{
	if (isG4()) {
		getPlatform();
		if (Platform == "dmb400" || Platform == "sma300") {
			aProfileEntry.prefBranch.setBoolPref("features.brightness", aActivated);
		}
	}
}

/**
* Get platform
*/
function getPlatform() {
    if (Platform != undefined)
	   return Platform;
    if (isG4()) {
		Platform = Services.prefs.getCharPref("system.hw-platform");
		logger.debug("getPlatform : Platform = " + Platform);
	}
	else {
	var generalSettings = Cc["@innes/system/generalsettings;1"].getService(Ci.nsISystemGeneralSettings);
		Platform = generalSettings.platform;
		logger.debug("getPlatform : Platform = " + Platform);
   }
	  return Platform;
}

/**
* Test if Gekkota is G4 generation
*/
var gIsG4 = undefined;
function isG4() {
     if (gIsG4 == undefined)      {
		let version = Services.prefs.getCharPref("extensions.lastAppVersion"); 
		let tab = version.split(".");
		let major = parseInt(tab[0]);
		gIsG4 = (major == 4);
	}
	return gIsG4;
}

function log(str) {
    logger.debug(str, null);
}

function warn(str) {
    logger.warn(str, null);
}

function error(str) {
    logger.error(str, null);
}
