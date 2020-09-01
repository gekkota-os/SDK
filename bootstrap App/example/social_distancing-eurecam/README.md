# Social distancing - Eurecam motion sensor

This App is an example to show a social distancing application with the `Eurecam` `COMPTIPIX 3D` motion sensor camera. 

It allows to limit automatically with a motion sensor the number of people in a room by displaying on a screen close to the access door:
- a message with a green background on a screen when the  max. threshold of people is not reached:
```
YOU CAN ENTER, 
Please respect social distancing measures
```      
- a message with a red background on a screen when the max. threshold of people is reached:
```
PLEASE WAIT, 
Please respect social distancing measures
```
The people threshold needs to be configured with the motion sensor configuration Web interface.

## Compatibility
It can run with any `Qeedji` device embedding `Gekkota OS 4.yy.zz` able to display some information on a screen.  

## App configuration
The App does not stop checking periodically the number of people in the room.
The App can have 3 states:
- `enter` state: the number of people estimated in the room is becoming appropriate.
	- play a sound (.mp3) once then 
	- until the number of people is not upper the defined threshold, does not stop displaying: 
		- the `YOU CAN ENTER` message for a while 
		- then, as option, a custom INFORMATION message for a while
- `wait` state: the number of people estimated in the room is becoming too high:
	- play a sound (.mp3) once then
	- until the number of people is equal or upper than the defined max. threshold, does not stop displaying: 
		- the `WAIT` message for a while 
		- then, as option, a custom INFORMATION message for a while
- `error` state: the App does not manage to connect to the `COMPTIPIX 3D` motion sensor.

The audio files are played at the beginning of the new state.
 
For the `wait` and `enter` states:
- one or two images can be displayed, 
- the duration for each image can be defined.

For the `error` state, only one image can be displayed.

The optional secondary image needs to be activated in the `app.html`.
The App can be customized by modifying some variables values in the `app.html` file.

```
//=============================================================================		
		var frequency = 1000; /* connection frequency to the motion sensor every <n> ms */  

		// configuration of the image behaviour for the enter state 
		var imageEnterMainDuration = 10000; /* main image duration in ms */
		var imageEnterSecondaryDuration = 5000; /* secondary image duration in ms */
		var imageEnterSecondaryActive = true; /* true: secondary image activated */
	
		// configuration of the image behaviour for the wait state  
		var imageWaitMainDuration = 10000; /* main image duration in ms*/
		var imageWaitSecondaryDuration = 5000; /* secondary image duration in ms */
		var imageWaitSecondaryActive = false; /* true: secondary image activated */

		// Debug mode  
		var DEBUG = false; /*true: App in debug mode */
```

**Resource naming pattern**

The App has 2 medias directories:
- `audio`:
    - `wait.mp3`: played once when pass into the `wait` state 
    - `enter.mp3`: played once when pass into the `enter` state
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

The `<device_output_resolution>` is the current Qeedji device resolution in the format `<width>x<height>`. For example, on the resolution of 1920x1080, the main files must be named as follows:
- `enter_main-1920x1080.jpg`
- `wait_main-1920x1080.jpg`

In case not found, the image fallback names are:
- `enter_main.jpg`
- `enter_secondary.jpg`
- `wait_main.jpg`
- `wait_secondary.jpg`

The error image naming pattern is `error-<device_output_resolution>.jpg`

If not all the required images are present, the application is entering in `error` state and displays an error message `Error loading images`, with a white background.

## Motion sensor configuration
Install the motion sensor as preconised by the `Eurecam` manufacturer. 
Ensure that the `COMPTIPIX 3D` motion sensor is running, available in your local network. Connect to its Web server, and in the menu `Occupancy` of the `Settings` pane, set the max. threshold of number of people authorized to enter in the room at the same time. 
Qeedji does not bring support on the `COMPTIPIX 3D` motion sensor configuration.
Check that the `COMPTIPIX 3D` motion sensor has a recent software version (V1.8 or above).

## Qeedji device configuration
Set the IP address of the `COMPTIPIX 3D` motion sensor camera (for example: `192.168.0.100`) as `field1` variable value of your Qeedji device.

## App logs
The `Gekkota OS` log can be activated in the `Logs` menu of the `Maintenance` pane of the device configuration Web interface menu: 
- Name : `app`
- Level: `DEBUG`   

For further information, refer to the `Qeedji` installation manual. 
