# Social distancing - WPAN push button 

This App is an example to show a social distancing application with the WPAN push-button.  

It allows to a people to limit the number of people in a room by displaying on a screen, located near the access door, to display different information messages thanks to an `EnOcean` WPAN push button :
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
An .mp3 sound is played at each state transitions.

&sup1; *this message can be also displayed, only for a couple of time, by pressing on the `+` key of the push-button*. 

## Compatibility
- `Qeedji` tablet
	- Gekkota OS 4.13.13 (or above) 
		- for device SMT210
- `Qeedji` media-players
	- Gekkota OS 4.13.10 (or above)
		- for device SMA300
		- for device DMB400
`USB to WPAN` adapter
	- With this App, an `USB to WPAN` adapter needs to be plugged on the device to work properly else an error is raised. 
- `EnOcean` device 
	- Example: `PTM 215B` push button. For further information, contact sales@qeedji.tech

## Qeedji device configuration  

Plug the `USB to WPAN` adapter to an USB connector of the `Qeedji` media-player or the `Qeedji` tablet. For further information about the Qeedji `USB to WPAN` adapter, contact support@qeedji.tech. 

Then connect to the device configuration Web interface:
- in the `Preferences > Maintenance` pane, set the appropriate value for the `innes.adapters.serial.uart_1.syspath` user preference.

| Qeedji device | user preference | default value | value to support the `USB to WPAN` adapter
|:--|:--|:--|:--
| SMT210  |`innes.adapters.serial.uart_1.syspath` | /dev/ttyUSB0  | /dev/ttyACM0 
| SMA300  |`innes.adapters.serial.uart_1.syspath` | /dev/ttyUSB0  | /dev/ttyACM0  
| DMB400  |`innes.adapters.serial.uart_1.syspath` | /dev/ttyAS1   | /dev/ttyACM0  

- to deal only with only one specific push button, in the `Configuration > Variables` pane, , it is advised to enter the push button MAC address in the `field1` variable input.

- then reboot the `Qeedji` device. 

## App customization

**The App is already configured to work properly without error. Anyway it is possible to customize it according to your needs.** 

The App can have 4 states:
- `wait` state (default state):
	- does not stop displaying: 
		- the `WAIT` message for a while, 
		- then, as option, a custom INFORMATION message for a while.
- `enter` state:
	- does not stop displaying: 
		- the `YOU CAN ENTER` message for a while, 
		- then, as option, a custom INFORMATION message for a while.
- `re-enter` state: 
	- displays: 
		- the `YOU CAN ENTER` message for a while, 
		- then, as option, a custom INFORMATION message for a while,
	- then returns to `wait` state.
- `error` state: 
	- either the `USB to WPAN` adapter is not plugged on the `Qeedji` device, 
	- or the `innes.adapters.serial.uart_1.syspath` user preference has a wrong value, 
	- or the embedded Gekkota_OS version does not support the `USB to WPAN` adapter.    

As option, a .mp3 sound can be played at each state transitions. 

The different states are driven by the push-buttons key values:
- `full circle` button: entering in `wait` state,
- `empty circle` button: entering in `enter` state,
- `+` button: entering in `re-enter` state ( = `enter` state with a timeout duration).

For each state, except for the `error` state: 
- one or two images can be displayed, 
- the duration for each image can be defined.

For the `error` state:
- only one image can be displayed.

The optional secondary image containing the custom INFORMATION message needs to be activated in the `app.html`.
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
    - `wait.mp3`: played once when passing into the `wait` state, 
    - `enter.mp3`: played once when passing into the `enter` state or the `re-enter` state.
- `images`:
    - images to be displayed with the `.jpg` format,
    - has to match a file naming pattern depending on the `Qeedji` device resolution.

- when 2 images are displayed, the `main` and `secondary` images named pattern are as follow:
    - for `enter` state: 
        - `enter_main-<device_output_resolution>.jpg`,
        - `enter_secondary-<device_output_resolution>.jpg`.
    - for `wait` state:
        - `wait_main-<device_output_resolution>.jpg`,
        - `wait_secondary-<device_output_resolution>.jpg`.
    
The `<device_output_resolution>` is the current `Qeedji` device resolution in the format `<width>x<height>`. For example, on the resolution of 1024x600, the main files must be named as follows:
- `enter_main-1024x600.jpg`
- `wait_main-1024x600.jpg`

In case not found, the image fallback names are:
- `enter_main.jpg`
- `enter_secondary.jpg`
- `wait_main.jpg`
- `wait_secondary.jpg`

The error image naming pattern is `error-<device_output_resolution>.jpg` 

If not all the required images are present, the application is entering in `error` state and displays an error message `Error loading images`, with a white background.