# Social distancing - WPAN push button 

This App is an example to show a social distancing application with the WPAN push-button.  

It allows to limit the number of people in a room by displaying on a screen close to the access door, some information message, commanded by a user pressing on the different keys of a `EnOcean` WPAN push button:
- a message&sup1; with a green background on a screen when the user is pressing on the `Empty circle` key of the push-button.
```
YOU CAN ENTER, 
Please respect social distancing measures
```      
- a message with a red background on a screen when the user is pressing on the `Full circle` key of the push-button.
```
PLEASE WAIT, 
Please respect social distancing measures
```
&sup1; *this message can be also displayed, only for a couple of time, by pressing on the `+` key of the push-button*. 

## Compatibility
- `Qeedji` device embedding `Gekkota OS 4.yy.zz`, not supporting WPAN natively
	- consequently, a `WPAN to USB` adapter needs to be used to work with this App. For further information, contact sales@qeedji.tech.
- `EnOcean` device 
	- example: `PTM 215B` push button. For further information, contact sales@qeedji.tech

## App configuration

The App can have 4 states:
- `wait` state: (default state) you cannot enter in the room for the moment:
	- play a sound (.mp3) once then 
	- does not stop displaying, : 
		- the `WAIT` message for a while, 
		- then, as option, a custom INFORMATION message for a while.
- `enter` state: you can enter in the room for the moment,
	- play a sound (.mp3) once then
	- does not stop displaying: 
		- the `YOU CAN ENTER` message for a while, 
		- then, as option, a custom INFORMATION message for a while.
- `re-enter` state: `enter` state only for a couple of time
	- play a sound (.mp3) once then
	- displays for a while : 
		- The `YOU CAN ENTER` message for a while, 
		- then, as option, a custom INFORMATION message for a while,
	- returns to `wait` state
- `error` state: the push-button MAC address is not properly defined or the `WPAN to USB` adapter is not plugged on the Qeedji device.

The different states are commanded by the push-buttons key values:
- `Full circle`  : entering in `wait` state,
- `Empty circle` : entering in `enter` state,
- `+`          : entering in `re-enter` state ( = `enter` state with a timeout duration).

For the `wait`, `enter`, and `re-enter` states, 
- one or two images can be displayed, 
- the duration for each image can be defined.

For the `error` state, only one image can be displayed.

The optional secondary image needs to be activated in the `app.html`.
The App can be customized by modifying some variables values in the `app.html` file.

```
//=============================================================================		
		// configuration of the image behaviour for the enter state 
		var imageEnterMainDuration = 10000; /* main image duration in ms */
		var imageEnterSecondaryDuration = 5000; /* secondary image duration in ms */
		var imageEnterSecondaryActive = true; /* true: secondary image activated */
		
		// configuration of the image behaviour for the wait state  
		var imageWaitMainDuration = 10000; /* main image duration in ms */
		var imageWaitSecondaryDuration = 5000; /* secondary image duration in ms */
		var imageWaitSecondaryActive = true; /* true: secondary image activated */

		// configuration of the reenter timeout duration
		var imageReEnterMainDuration = 5000; /* enter state timeout duration in ms */

		// Other default parameters
		var DEBUG = false;
```

**Resource naming pattern**

The App has 2 medias directories:
- `audio`:
    - `wait.mp3`: played once when pass in the `wait` state 
    - `enter.mp3`: played once when pass into the `enter` state or the `re-enter` state
- `images`:
    - images to be displayed with the `.jpg` format
    - has to match a file naming pattern depending on the `Qeedji` device resolution

- when 2 images are displayed, the `main` and `secondary` images named pattern are as follow:
    - For `enter` state: 
        - `enter_main-<device_output_resolution>.jpg`
        - `enter_secondary-<device_output_resolution>.jpg`
    - For `wait` state:
        - `wait_main-<device_output_resolution>.jpg`
        - `wait_secondary-<device_output_resolution>.jpg`
    
The `<device_output_resolution>` is the current `Qeedji` device resolution in the format `<width>x<height>`. For example, on the resolution of 1920x1080, the main files must be named as follows:
- `enter_main-1920x1080.jpg`
- `wait_main-1920x1080.jpg`

In case not found, the image fallback names are:
- `enter_main.jpg`
- `enter_secondary.jpg`
- `wait_main.jpg`
- `wait_secondary.jpg`

The error image naming pattern is `error-<device_output_resolution>.jpg` 

If not all the required images are present, the application is entering in `error` state and displays an error message `Error loading images`, with a white background.
    
## WPAN to USB adapter 
Plug the `WPAN to USB` adapter to an USB connector of the `Qeedji` media-player or the `Qeedji` tablet.

## Qeedji device configuration
It is possible to filter the key-pressed values of the `EnOcean` WPAN device by defining the MAC address of the target button in the `field1` variable value of the `Qeedji` media-player. 
If no value is defined in this field, the App does not apply filter and accept the key-pressed values of any push button.

That implies also to set the appropriate value for the `innes.adapters.serial.uart_1.syspath` user preference in the Qeedji device:

| Qeedji device | default value | value to support the `WPAN to USB` adapter
|:--|:--|:--
| Gekkota OS 4.13.11 for device SMT210  | /dev/ttyUSB0  | /dev/ttyACM0 
| Gekkota OS 4.13.10 for device SMA300  | /dev/ttyUSB0  | /dev/ttyACM0  
| Gekkota OS 4.13.10 for device DMB400  | /dev/ttyAS1   | /dev/ttyACM0 

For further information about this Qeedji `WPAN to USB` adapter, contact support@qeedji.tech  

## App logs
The `Gekkota OS` log can be activated in the `Logs` menu of the `Maintenance` pane of the device configuration Web interface menu: 
- Name : `app`
- Level: `DEBUG`   

For further information, refer to the `Qeedji` installation manual. 
