/**
 * Fonction: configuration-by-script
 * File Base: example_000000000000-V1.11.12.js
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

// Jack GPIO : uncomment one line below corresponding to your choice
var Platform = getPlatform();
setJackGPIO(false, "in", 0);         		// Jack GPIO : Input, No Debouncing
//setJackGPIO(false, "out", 0);         // Jack GPIO : Output, No Debouncing

// ---- Save any previous preference changed : DO NOT COMMENT THE FOLLOWING LINE !!!
Services.prefs.savePrefFile(null);

// ---------------------------------------
// ---------------------------------------
// ---- END of the user configuration
// ---------------------------------------
// ---------------------------------------

/**
* Set the GPIO on Phoenix connector
* @param aInvert : true to invert data input
* @param aDirection : direction on the GPIO, in for input, out for output, disable to don't set the drection
* @param aDebouncing : to set the duration of the debouncing (in ns), maximum=10000000000 (10s) mimimum=0 (no deboucning)
*/
function setJackGPIO(aInvert, aDirection, aDebouncing) {
  if (Platform == "dmb400" || Platform == "sma300" || Platform == "smh300" || Platform == "smt210") {
    if (Platform == "sma300" || Platform == "smh300" || Platform == "smt210") {
      if (aDirection == "disable") {
        Services.prefs.setBoolPref("system.connector.jack35_1.1.io.uart_1.enabled", true);
      }
      else {
        Services.prefs.setBoolPref("system.connector.jack35_1.1.io.uart_1.enabled", false);
      }
    }

    // Set the direction : input or output
    if (aDirection == "out") {
      Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.authorized", false);
      Services.prefs.setBoolPref("innes.app-profile.gpio-output.jack35-gpio_1.jack35_1.*.authorized", true);
      Services.prefs.setBoolPref("system.connector.jack35_1.1.io.jack35-gpio_1.enabled", true);
    }
    else if (aDirection == "in") {
      Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.authorized", true);
      Services.prefs.setBoolPref("innes.app-profile.gpio-output.jack35-gpio_1.jack35_1.*.authorized", false);
      Services.prefs.setBoolPref("system.connector.jack35_1.1.io.jack35-gpio_1.enabled", true);
    }
    else if (aDirection == "disable") {
      Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.authorized", false);
      Services.prefs.setBoolPref("innes.app-profile.gpio-output.jack35-gpio_1.jack35_1.*.authorized", false);
      Services.prefs.setBoolPref("system.connector.jack35_1.1.io.jack35-gpio_1.enabled", false);
    }

    // Set invert or not
    Services.prefs.setBoolPref("innes.app-profile.gpio-input.jack35-gpio_1.jack35_1.*.invert-value", aInvert);

    // Set the debouncing time
    if (aDirection == "in") {
      Services.prefs.setIntPref("system.connector.jack35_1.1.io.jack35-gpio_1.debouncing.period", aDebouncing);

      if (aDebouncing == 0) {
        Services.prefs.setBoolPref("system.connector.jack35_1.1.io.jack35-gpio_1.debouncing.enabled", false);
      }
      else {
        Services.prefs.setBoolPref("system.connector.jack35_1.1.io.jack35-gpio_1.debouncing.enabled", true);
      }
    }
    else {
      Services.prefs.setBoolPref("system.connector.jack35_1.1.io.jack35-gpio_1.debouncing.enabled", false);
    }
  }
  else if (Platform == "dme204") {
    // Set the direction : input or output
    if (aDirection == "out") {
      Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.authorized", false);
      Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.jack35_1.*.authorized", true);
    }
    else if (aDirection == "in") {
      Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.authorized", true);
      Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.jack35_1.*.authorized", false);
    }
    else if (aDirection == "disable") {
      Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.authorized", false);
      Services.prefs.setBoolPref("innes.app-profile.gpio-output.epld_1.jack35_1.*.authorized", false);
    }
    // Set invert or not
    Services.prefs.setBoolPref("innes.app-profile.gpio-input.epld_1.jack35_1.*.invert-value", aInvert);
    // Set the debouncing time
    if (aDirection == "in") {
      Services.prefs.setIntPref("system.connector.jack35_1.1.io.epld_1.debouncing.period", aDebouncing);
      if (aDebouncing == 0) {
        Services.prefs.setBoolPref("system.connector.jack35_1.1.io.epld_1.debouncing.enabled", false);
      }
      else {
        Services.prefs.setBoolPref("system.connector.jack35_1.1.io.epld_1.debouncing.enabled", true);
      }
    }
    else {
      Services.prefs.setBoolPref("system.connector.jack35_1.1.io.epld_1.debouncing.enabled", false);
    }
  }
}

/**
* Get platform
*/
function getPlatform() {
  if (Platform != undefined)
    return Platform;
  Platform = Services.prefs.getCharPref("system.hw-platform");
  logger.debug("getPlatform : Platform = " + Platform);
  return Platform;
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