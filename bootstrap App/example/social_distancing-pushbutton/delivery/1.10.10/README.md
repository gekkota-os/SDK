# Social distancing - WPAN push button 

This App is an example to show a social distancing application with the a WPAN push-button.  

It allows to limit the number of people in a room by displaying on a screen close to at the access door, some information message, commanded by a user pressing on the different keys of a `EnOcean` WPAN push button:
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

It can run with any tablet or media player `Qeedji` device embedding `Gekkota OS 4.yy.zz`.

The WPAN is not supported by default on the `Qeedji` media player or `Qeedji` tablet. So a `WPAN to USB` adapter needs to be provided to work with this App. For further information, contact sales@qeedji.tech.
`EnOcean` device example: `PTM 215B` push button. For further information, contact sales@qeedji.tech

## App configuration

The App has 2 medias directories:
- `audio`:
    - `wait.mp3`: played once when pass in the `wait` state 
    - `enter.mp3`: played once when pass into the `enter` state or the `re-enter` state
- `images`:
    - images to be displayed with the `.jpg` format
    - has to match a file naming pattern depending on the `Qeedji` device resolution

The App can be customized by modifying some variables values in the `app.html` file:
```
// configuration of the image behaviour for the enter state 
var imageEnterMainDuration = 10000; /* main image duration */
var imageEnterSecondaryDuration = 5000; /* secondary image duration */
var imageEnterSecondaryActive = true; /* true: secondary image activated */

// configuration of the image behaviour for the wait state  
var imageWaitMainDuration = 10000; /* main image duration */
var imageWaitSecondaryDuration = 5000; /* secondary image duration */
var imageWaitSecondaryActive = true; /* true: secondary image activated */

// configuration of the re-enter timeout duration
var imageReEnterMainDuration = 5000; /* enter state timeout duration */

// Other default parameters
var DEBUG = false; /*true: App in debug mode */
```

The App can behave with 4 states:
- `wait` state: (default state) you cannot enter in the room for the moment,
- `enter` state: you can enter in the room for the moment
- `re-enter` state: `enter` state only for a couple of time  
- `error` state: the push-button MAC address is not properly defined or the `WPAN to USB` adapter is not plugged on the Qeedji device.

For the `wait` and `enter` states, 
- one or two images can be displayed, 
- the duration for each image can be defined.

For the `re-enter` state, the `main` enter image is displayed for a defined duration. 

When the `re-enter` state is finished, the App goes into the `wait` state.

For the `error` state, only one image can be displayed.

The different states are commanded by the push-buttons key values:
- `Full circle`  : `wait` state,
- `Empty circle` : `enter` state,
- `+`          : `re-enter` state ( = `enter` state with a timeout duration).

The secondary images are required only if the secondary image is enabled.

If not all the required images are present, the application will not start and display an error message `Error loading images`.

**Resource naming pattern**

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

In case the appropriate images cannot be found, a message *Error loading images* is displayed.
    
## WPAN to USB adapter 
Plug the `WPAN to USB` adapter to an USB connector of the `Qeedji` media-player or the `Qeedji` tablet.

## Qeedji device configuration
It is possible to filter the key-pressed values of the `EnOcean` WPAN device by defining the MAC address of the target button in the `field1` variable value of the `Qeedji` media-player. 
If no value is defined in this field, the App does not apply filter and accept to the key-pressed values of the push button.

That implies also to set the appropriate value for the `innes.adapters.serial.uart_1.syspath` user preference in the Qeedji device:

| Qeedji device | default value | value to support the `WPAN to USB` adapter
|:--|:--|:--
| Gekkota OS 4.13.11 for device SMT210  | /dev/ttyUSB0  | /dev/ttyACM0 
| Gekkota OS 4.13.10 for device SMA300  | /dev/ttyUSB0  | /dev/ttyACM0  
| Gekkota OS 4.13.10 for device DMB400  | /dev/ttyAS1   | /dev/ttyACM0 

For further information about this Qeedji `WPAN to USB` adapter, contact support@qeedji.tech  
