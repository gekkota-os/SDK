# Settings script
An Innes equipment can be configured via a script.  
If the *innes.app-profile.addon-manager.\*.\*.\*.authorized* preference is set to true, then the equipment downloads a settings script and executes its Javascript code as an application extension.   

The process to  get the settings script via the network is based on the DHCP auto configuration feature.   
The option 66 of the DHCP server provides the TFTP server address.  
When the IP address is given or renewed by the DHCP server, the equipment tries to get the settings file on the TFTP server.  

The settings file can also be located on a USB storage.  
It can have the following names :

- "<*MAC*>.js" where <*MAC*> is the MAC address of the platform. It is a specific configuration;
- "00000000000.js" which is for general configuration.  

Few examples of settings script are available in the *examples* directory.  
