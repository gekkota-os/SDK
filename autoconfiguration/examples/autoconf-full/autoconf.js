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
     */
    
    const Cc = Components.classes;
    const Ci = Components.interfaces;
    const nsIPrefBranch=Ci.nsIPrefBranch;
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
    
    
    
    // ---------------------------------------
    // ---- Administrator configuration
    // ---------------------------------------
    
    // ---- Defines hostname : uncomment the line after
    //setHostname("sma200-test");
    
    // ---- Defines admin login : uncomment the line after
    //Services.prefs.setCharPref("innes.webserver.username", "admin");
    
    // ---- Defines admin password : uncomment the line after
    //Services.prefs.setCharPref("innes.webserver.password", "admin");
    
    
    // ---------------------------------------
    // ---- LAN configuration
    // ---------------------------------------
    
    // ---- Retrieve lan adapter : DO NOT COMMENT THE FOLLOWING LINE !!!
    let lan = getNetAdapter("LAN");
    
    
    // ---- Choose static or DHCP for IPV4 : uncomment one of the 2 lines after
    //enableDhcpv4(lan); // This one for DHCP
    //disableDhcpv4(lan); // This one for static
    
    // ---- Set static IP address, netmask and gateway for LAN adapter : uncomment the line after
    //setIPv4StaticAddress(lan, "192.168.0.2", "255.255.255.0", "192.168.0.1");
    
    // ---- Choose if UPnP discover gateway is activated or not
    // ---- Not that if DHCP for IPv4 is enabled by enableDhcpv4(), the gateway will be defined by DHCP, not by UPnP
    // ---- uncomment one of the 2 lines after
    //enableGatewayUpnp(lan); // ---- This one for activating UPnP discover
    //disableGatewayUpnp(lan); // ---- This one for gateway static (if disableDhcpv4() is used) or DHCP (if enableDhcpv4() is used)
    
    // ---- Set the IGMP version : uncomment one of the 2 lines after
    //setIgmpVersion(lan,3); // ---- This one for IGMP V3
    //setIgmpVersion(lan,2); // ---- This one for IGMP V2
    
    // ---- set the maximun bitrate value for LAN adapter. Possible values are :
    // ---- MAX_BITRATE_VALUE_56_KBPS
    // ---- MAX_BITRATE_VALUE_128_KBPS
    // ---- MAX_BITRATE_VALUE_256_KBPS
    // ---- MAX_BITRATE_VALUE_512_KBPS
    // ---- MAX_BITRATE_VALUE_1024_KBPS
    // ---- MAX_BITRATE_VALUE_2048_KBPS
    // ---- MAX_BITRATE_VALUE_5120_KBPS
    // ---- MAX_BITRATE_VALUE_10240_KBPS
    // ---- uncomment the line after
    //setMaxBitrateValue(lan, Ci.nsISystemAdapterIP.MAX_BITRATE_VALUE_10240_KBPS);
    
    // ---- Choose if the max bitrate limitation is activated or not : uncomment one of the 2 lines after
    //disableMaxBitrate(lan); // ---- This one for disabling
    //enableMaxBitrate(lan); // ---- This one for enabling
    
    
    // ---------------------------------------
    // ---- Output configuration
    // ---------------------------------------
    
    // ---- Retrieve display : DO NOT COMMENT THE FOLLOWING LINE !!!
    let display = getDisplayAdapter();
    
    
    // ---- Change display mode : uncomment the line after
    //setDisplayMode(display, "1280x720 60Hz SMPTE (720p)");
    
    // ---- Change display rotation : uncomment the line after
    //display.rotation = 0; // ---- Possible values are 0, 90, 180, 270
    // --- Change the display area (overscan)
    //setOverscan(100, 50, 200, 300); // ---- This one for enabling and change the parameters
    //disableOverscan(); // ---- This one for disabling
    
    // ---- Choose if audio is muted : uncomment one of the 2 lines after
    //setAudioIsDisable(true); // ---- This one for audio mute
    //setAudioIsDisable(false); // ---- This one for audio activated
    
    // ---- Defines the audio volume : uncomment the line after
    //setAudioVolume(50, 50);
    
    
    // ---------------------------------------
    // ---- Servers configuration
    // ---------------------------------------
    
    // ---- Configuring a Plugncast G2 server : uncomment the 2 lines after
    //setPlugnCastG2("http://192.168.1.29/", 1);
    //disablePlugnCastG3();
    
    // ---- Configuring a Plugncast G3 server : uncomment the 2 lines after
    //setPlugnCastG3("http://192.168.1.73/.playo-1ut", 1, "admin", "admin");
    //disablePlugnCastG2();
    
    // ---- Configuring no plugncast server (Screen Composer use for example) : uncomment the 2 lines after
    //disablePlugnCastG2();
    //disablePlugnCastG3();
    
    // ---- Choose static or DHCP for DNS : uncomment one of the 2 lines after
    //enableDnsDhcp(lan); // ---- This one for DHCP
    //disableDnsDhcp(lan); // ---- This one for static
    
    // ---- set static DNS servers : uncomment the line after
    //setDns(lan, "192.168.0.1", "192.168.0.1");
    
    // ---- Defines NTP settings : uncomment the line after
    //setNtpSettings("fr.pool.ntp.org", 20, 10);
    
    // ---- Choose if NTP is enabled : uncomment one of the 2 lines after
    //setNtpIsEnable(true); // ---- This one for enabling
    //setNtpIsEnable(false); // ---- This one for disabling
    
    // ---- Choose if proxy is used : uncomment one of the 2 lines after
    //setProxyIsEnable(true); // ---- This one for enabling
    //setProxyIsEnable(false); // ---- This one for disabling
    
    // ---- Defines 'No proxy for ' : uncomment the line after
    //setNoProxyFor("localhost, 127.0.0.1");
    
    // ---- Defines and activate http proxy : uncomment the line after
    //setProxyHttp(true, "proxy-url", 3128, "login", "password");
    
    // ---- Deactivate http proxy : uncomment the line after
    //setProxyHttp(false, "", 0, "", "");
    
    // ---- Defines and activate https proxy : uncomment the line after
    //setProxyHttps(true, "proxy-url", 3128, "login", "password");
    
    // ---- Deactivate http proxy : uncomment the line after
    //setProxyHttps(false, "", 0, "", "");
    
    // ---- Defines and activate ftp proxy : uncomment the line after
    //setProxyFtp(false, "proxy-url", 3128, "login", "password");
    
    // ---- Deactivate ftp proxy : uncomment the line after
    //setProxyFtp(false, "", 0, "", "");
    
    
    // ---------------------------------------
    // ---- Maintenance configuration
    // ---------------------------------------
    
    // ---- Disable test card mode : uncomment the line after
    //Services.prefs.setBoolPref("innes.player.mire", false);
    
    // ---- Enable test card mode : uncomment the line after
    //Services.prefs.setBoolPref("innes.player.mire", true);
    
    
    // ---- Save any previous preference changed : DO NOT COMMENT THE FOLLOWING LINE !!!
    Services.prefs.savePrefFile(null);
    
    
    
    // ---------------------------------------
    // ---------------------------------------
    // ---- END of the user configuration
    // ---------------------------------------
    // ---------------------------------------
    
    
    // ---- Functions section
    
    /**
     * Configuring a Plugncast G2 server.
     * @param aBaseuri : baseuri of Plugncast server
     * @param aHeartbeat : heartbeat of download in minutes */
    function setPlugnCastG2 (aBaseuri, aHeartbeat)
    {
    	Services.prefs.setCharPref("innes.app-profile.manifest-downloader.*.*.*.class-name", "g2");
        let basepref="innes.app-profile.manifest-downloader:g2.*.*.*.";
    	Services.prefs.setBoolPref(basepref + "authorized", true);
        if (aHeartbeat <= 0)
    	   aHeartbeat = 1;
    	Services.prefs.setIntPref(basepref + "heartbeat", aHeartbeat*60);
    	if (aBaseuri.charAt(aBaseuri.length-1) == '/')
        {
           aBaseuri = aBaseuri.substr(0,aBaseuri.length-1);
    	   log ("setPlugnCastG2 after aBaseuri = " + aBaseuri);
        }
        if (aBaseuri.lastIndexOf("ftp:") == 0)
        {
            Services.prefs.setCharPref(basepref + "base-uri", aBaseuri);
            let adr=aBaseuri.substr(4);
            Services.prefs.setCharPref(basepref + "wsman.base-uri", 
    		    "http:" + adr + "/wsman");
        }
    	else
        {
            Services.prefs.setCharPref(basepref + "base-uri", 
    		    aBaseuri + "/plugnCast/");
            Services.prefs.setCharPref(basepref + "wsman.base-uri", 
    		    aBaseuri + "/wsman");
        }
    }
    /**
     * Disable a Plugncast G2 server.  */
    function disablePlugnCastG2 ()
    {
        let basepref="innes.app-profile.manifest-downloader:g2.*.*.*.";
    	Services.prefs.setBoolPref(basepref + "authorized", false);
    }
    /**
     * Configuring a Plugncast G3 server.
     * @param aBaseuri : baseuri of Plugncast server
     * @param aHeartbeat : heartbeat of download in minutes
     * @param aUsername : username for authentication
     * @param aPassword : password for authentication */
    function setPlugnCastG3 (aBaseuri, aHeartbeat, aUsername, aPassword)
    {
    	Services.prefs.setCharPref("innes.app-profile.manifest-downloader.*.*.*.class-name", "g3");
        let basepref="innes.app-profile.manifest-downloader:g3.*.*.*.";
    	Services.prefs.setBoolPref(basepref + "authorized", true);
        if (aHeartbeat <= 0)
    	   aHeartbeat = 1;
    	Services.prefs.setIntPref(basepref + "heartbeat", aHeartbeat*60);
        Services.prefs.setCharPref(basepref + "base-uri", aBaseuri);
        Services.prefs.setCharPref(basepref + "username", aUsername);
        Services.prefs.setCharPref(basepref + "password", aPassword);
    }
    /**
     * Disable a Plugncast G3 server.  */
    function disablePlugnCastG3 ()
    {
        let basepref="innes.app-profile.manifest-downloader:g3.*.*.*.";
    	Services.prefs.setBoolPref(basepref + "authorized", false);
    }
    /**
     * Set and active the maximun bitrate value for a network adapter.
     * @param aAdapter : the network adapter to be modified
     * @param aMaxBirateValue : the maximun bitrate value in kbps */
     function setMaxBitrateValue(aAdapter, aMaxBirateValue)
     {
        aAdapter.isMaxBitrateEnabled = true;
        aAdapter.maxBitrateValue = aMaxBirateValue;
     }
    /**
     * Disable the maximun bitrate feature. */
     function disableMaxBitrate(aAdapter)
     {
        aAdapter.isMaxBitrateEnabled = false;
     }
    /**
     * Disable the maximun bitrate feature. */
     function enableMaxBitrate(aAdapter)
     {
        aAdapter.isMaxBitrateEnabled = true;
     }
    /**
     * Enable the DHCP for IPv4.
     * @param aAdapter : the network adapter to be modified */
     function enableDhcpv4(aAdapter)
     {
        aAdapter.isDhcpv4Enabled=true;
     }
    /**
     * Disable the DHCP for IPv4. Use static IPV4 adress.
     * @param aAdapter : the network adapter to be modified */
     function disableDhcpv4(aAdapter)
     {
        aAdapter.isDhcpv4Enabled=false;
     }
    /**
     * Set the static IPV4 adress for a network adapter.
     * @param aAdapter : the network adapter to be modified
     * @param aAdress : the string for the IP adress
     * @param aNetmask : the string for the netmask
     * @param aGateway : the string for the gateway adress */
     function setIPv4StaticAddress(aAdapter,aAdress,aNetmask,aGateway)
     {
        aAdapter.setIpv4Address(
    	   createIPv4Adress(aAdress),
    	   createIPv4Adress(aNetmask),
    	   createIPv4Adress(aGateway),
    	   Ci.nsISystemAdapterIP.METRIC_AUTOMATIC);
     }
    /**
     * Enable the UPnP discover gateway.
     * Not that if DHCP for IPv4 is enabled, the gateway will be defined by DHCP, not by UPnP.
     * @param aAdapter : the network adapter to be modified */
     function enableGatewayUpnp(aAdapter)
     {
        aAdapter.isGatewayUPnPDiscovered=true;
     }
    /**
     * Disable the UPnP discover gateway.
     * @param aAdapter : the network adapter to be modified */
     function disableGatewayUpnp(aAdapter)
     {
        aAdapter.isGatewayUPnPDiscovered=false;
     }
    /**
     * Enable the DHCP for DNS.
     * @param aAdapter : the network adapter to be modified */
     function enableDnsDhcp(aAdapter)
     {
        aAdapter.isDnsDhcpEnabled=true;
     }
    /**
     * Disable the DHCP for DNS. Use static DNS
     * @param aAdapter : the network adapter to be modified */
     function disableDnsDhcp(aAdapter)
     {
        aAdapter.isDnsDhcpEnabled=false;
     }
    /**
     * Set the primary en secondary DNS server for a network adapter.
     * @param aAdapter : the network adapter to be modified
     * @param aAdress : the string adress for the primary DNS server
     * @param aAdress : the string adress for the secondary DNS server
    */
    function setDns(aAdapter,aDns1,aDns2)
    {
       aAdapter.setDns(createIPv4Adress(aDns1),createIPv4Adress(aDns2));
    }
    /**
     * Set the IGMP version for a network adapter.
     * @param aAdapter : the network adapter to be modified
     * @param aValue : the version
    */
    function setIgmpVersion(aAdapter,aValue)
    {
       aAdapter.igmpVersion=aValue;
    }
    
    /**
     * Retrieve the component for a network adapter
     * @param aName : type of the adapter (lan, wlan, wwan)
     * @return component of interface nsISystemIP for a network adapter */
    function getNetAdapter(aName)
    {
        let iids = {
           LAN: Ci.nsISystemAdapterLAN,
           WAN: Ci.nsISystemAdapterWLAN,
           WWAN: Ci.nsISystemAdapterWWAN 
        };
        let id = iids[aName];
    	log ("id = " + id);
    	if (id == undefined)
        {
    	   error ("undefined network adapter type");
           throw new Error("undefined network adapter type");
        }
    	let list = systemManager.getAdaptersByIId(id);
        if (list.lenght == 0)
           throw new Error("no adapter found");
        let adapter = list.queryElementAt(0, id);
        adapter = adapter.registered;
    	log ("getNetAdapter OK adapter = " + adapter);
        return adapter;
    }
    /**
     * Retrieve the first display adapter
     * @return component of interface nsISystemAdapterDisplayOutput for the display adapter */
    function getDisplayAdapter()
    {
       let id = Ci.nsISystemAdapterDisplayOutput;
       let list = systemManager.getAdaptersByIId(id);
       if (list.lenght == 0)
           throw new Error("no adapter found");	
       let adapter = list.queryElementAt(0, id);
       return adapter.registered;
    }
    /**
     * Set the display mode of a display adapter 
     * @param aDisplay : the display adapter to be modified
     * @param aModeLabel : label of mode to be setted
     */
    function setDisplayMode(aDisplay, aModeLabel)
    {
       let list = aDisplay.displayModesList;
       let l = list.length;
       for (let i = 0; i < l; i++)
       {
          let mode = list.queryElementAt(i, Ci.nsISystemDisplayMode);
          if (mode.label.toUpperCase() == aModeLabel.toUpperCase())
          {
             aDisplay.displayMode = mode;
    		 return;
          }
       }
       throw new Error("mode not found");
    }
    /**
    * Set the display area (overscan).
    * @param aTop The top  position of the display area (can be negative)
    * @param aLeft The left position of the display area (can be negative)
    * @param aWidth the width of display area. Use 0 for complete width of display 
    * @param aHeight the height of display area. Use 0 for complete height of display  */
    function setOverscan(aTop, aLeft, aWidth, aHeight)
    {
    	Services.prefs.setIntPref("innes.player.display.top", aTop);
    	Services.prefs.setIntPref("innes.player.display.left", aLeft);
    	Services.prefs.setIntPref("innes.player.display.width", aWidth);
    	Services.prefs.setIntPref("innes.player.display.height", aHeight);
    }
    /**
    * Disable the overscan */
    function disableOverscan()
    {
       setOverscan(0,0,0,0);
    }
    function createIPv4Adress (aAdress)
    {
       let ip = Cc["@innes/system/ipv4structure;1"].createInstance(Ci.nsISystemIPv4Structure);
       log("createIPv4Adress aAdress = " + aAdress);
       ip.stringAddress = aAdress;
       return ip;
    }
    /**
     * Set the hostname of the platforme
    * @param aName : the new hostname
    */
    function setHostname(aName)
    {
    	let settings = Cc["@innes/system/generalsettings;1"].getService(Ci.nsISystemGeneralSettings);
    // Verify hostname
    	log("setHostname begin aName =" + aName);
        let pattern="[A-Za-z0-9][A-Za-z0-9.-]{";
    	pattern += aName.length-1 + "}";
    	let re = new RegExp(pattern);
    		log ("setHostname re = " + re);
        if (!re.test(aName))
        {
           throw new Error("unauthorized character in hostname");
        }
    	settings.hostname = aName;
    	log("setHostname end");
    }
    /**
    * Set the audio volume
    * @param aVolumeL : volume for left channel between 0 to 100
    * @param aVolumeR : volume for rigth channel between 0 to 100
    */
    function setAudioVolume(aVolumeL, aVolumeR)
    {
       var systemPref = Cc["@innes/systemprefservice;1"].getService(Ci.nsIPrefBranch);
       let volume = aVolumeL | (aVolumeR << 8);
       systemPref.setIntPref("system/display/audio/volume", volume);
    }
    /**
    * Enable or disable audio
    * @param aDisable : false to enable audio, true to mute audio
    */
    function setAudioIsDisable(aDisable)
    {
       var systemPref = Cc["@innes/systemprefservice;1"].getService(Ci.nsIPrefBranch);
       systemPref.setBoolPref("system/display/audio/mute", aDisable);
    }
    /**
    * Enable or disable NTP service
    * @param aEnable : true to enable the NTP service, false otherwise
    */
    function setNtpIsEnable(aEnable)
    {
    	var systemPref = Cc["@innes/systemprefservice;1"].getService(Ci.nsIPrefBranch);
       systemPref.setBoolPref("system/network/ntp_enabled", aEnable);
    }
    /**
    * Set the NTP configuration
    * @param aServer : the adresse of NTP server
    * @param aNbTries : the maximum number of try
    * @param aTimeout : the delay between 2 tries in secondes
    */
    function setNtpSettings(aServer,aNbTries,aTimeout)
    {
    	var systemPref = Cc["@innes/systemprefservice;1"].getService(Ci.nsIPrefBranch);
       systemPref.setCharPref("system/network/ntp_server", aServer);
       systemPref.setIntPref("system/network/ntp_nbtries", aNbTries);
       systemPref.setIntPref("system/network/ntp_timeout", aTimeout);
    }
    /*
    * Enable or disable proxies usage.
    * @param aEnable : true to use proxies, false otherwise
    **/
    function setProxyIsEnable(aEnable)
    {
    	Services.prefs.setIntPref("network.proxy.type", (aEnable) ? 1 : 0);
    }
    /*
    * Set list of URLs with direct access.
    * @param aNoProxyFor: list of URL separate by ',' caractere
    */ 
    function setNoProxyFor(aNoProxyFor)
    {
    	Services.prefs.setCharPref("network.proxy.no_proxies_on", aNoProxyFor);
    }
    /*
    * Set the HTTP proxy configuration
    * @param aEnabled:  true to use proxy, false otherwise
    * @param aServer:  URL of proxy server
    * @param aPort:  Port of proxy server
    * @param aLogin:  Login to proxy connexion
    * @param aPassword:  Passord to proxy connexion
    */ 
    function setProxyHttp(aEnabled,aServer,aPort,aLogin,aPassword)
    {
    	setProxy("http", aEnabled,aServer,aPort,aLogin,aPassword);
    }
    /*
    * Set the HTTPS proxy configuration
    * @param aEnabled:  true to use proxy, false otherwise
    * @param aServer:  URL of proxy server
    * @param aPort:  Port of proxy server
    * @param aLogin:  Login to proxy connexion
    * @param aPassword:  Passord to proxy connexion
    */ 
    function setProxyHttps(aEnabled,aServer,aPort,aLogin,aPassword)
    {
    	setProxy("ssl", aEnabled,aServer,aPort,aLogin,aPassword);
    }
    /*
    * Set the FTP proxy configuration
    * @param aEnabled:  true to use proxy, false otherwise
    * @param aServer:  URL of proxy server
    * @param aPort:  Port of proxy server
    * @param aLogin:  Login to proxy connexion
    * @param aPassword:  Passord to proxy connexion
    */ 
    function setProxyFtp(aEnabled,aServer,aPort,aLogin,aPassword)
    {
    	setProxy("ftp", aEnabled,aServer,aPort,aLogin,aPassword);
    }
    function setProxy(aPrefix, aEnabled,aServer,aPort,aLogin,aPassword)
    {
    	let prefix = "network.proxy.";
    	let prefixInnes =  "innes.network.proxy.";
    	Services.prefs.setBoolPref(prefixInnes + aPrefix + "_enabled", aEnabled);
    	Services.prefs.setCharPref(prefix + aPrefix, (aEnabled) ? aServer : "");
    	Services.prefs.setIntPref(prefix + aPrefix + "_port", (aEnabled) ? aPort : 0);
    	Services.prefs.setCharPref(prefixInnes + aPrefix, aServer);
    	Services.prefs.setIntPref(prefixInnes + aPrefix + "_port", aPort);
    	Services.prefs.setCharPref(prefixInnes + aPrefix + "_login", aLogin);
    	Services.prefs.setCharPref(prefixInnes + aPrefix + "_password", aPassword);
    }
    function log(str)
    {
        logger.debug(str, null);
    }
    function warn(str)
    {
        logger.warn(str, null);
    }
    function error(str)
    {
        logger.error(str, null);
    }
    