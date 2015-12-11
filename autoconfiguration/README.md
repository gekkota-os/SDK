# Auto-configuration
An Innes equipment can be auto-configured.  
For this, the *innes.app-profile.addon-manager.\*.\*.\*.authorized* preference must be set to true.  
If auto-configuration is enabled, then the equipment downloads an auto-configuration file and executes it.  

## Auto-configuration file
The auto-configuration file contains some Javascript code and is located on a TFTP server or an USB storage volume.  
It can have the following names :

- *"<MAC\>.js"* where <MAC\> is the MAC address of the platform. It is a specific auto-configuration;
- *"00000000000.js"* which is for general autoconfiguration.  

Few examples of auto-configuration file are available in the *examples* directory.