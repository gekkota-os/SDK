# Social distancing - Eurecam motion sensor

This App is an example to show a social distancing application with a `Eurecam` motion sensor (for example: the `COMPTIPIX 3D` camera). 

It allows to limit automatically with a motion sensor the number of people in a room by displaying on a screen close to at the access door:
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
The people threshold need to be configured with the motion sensor configuration Web interface.

## Compatibility
It can run with any tablet or media-player `Qeedji` devices embedding `Gekkota OS 4.yy.zz`.  

## App configuration
The App has 2 medias directories:
- `audio`:
    - `wait.mp3`: played once when pass in the wait state 
    - `enter.mp3`: played once when pass into the enter state
- `images`:
    - images to be displayed with the `.jpg` format
    - has to match a file naming pattern depending on the `Qeedji` device resolution
  
The App can be customized by modifying some variables values in the `app.html` file:
```
// polling frequency of motion sensor 
var frequency = 1000;

// configuration of the image behaviour for the enter state 
var imageEnterMainDuration = 10000; /* main image duration */
var imageEnterSecondaryDuration = 5000; /* secondary image duration */
var imageEnterSecondaryActive = true; /* true: secondary image activated */
    
// configuration of the image behaviour for the wait state  
var imageWaitMainDuration = 10000; /* main image duration */
var imageWaitSecondaryDuration = 5000; /* secondary image duration */
var imageWaitSecondaryActive = false; /* true: secondary image activated */

// Debug mode  
var DEBUG = false; /*true: App in debug mode */
```

The App can behave with 3 states:
- `enter` state: the number of people estimated in the room is under the defined number of people threshold.
- `wait` state: the number of people estimated in the room is equal or upper that the defined number of people threshold.
- `error` state: the App does not manage to connect to the `COMPTIPIX 3D` motion sensor.

The audio files are played at the beginning of the new state.
 
For the `wait` and `enter` states:
- one or two images can be displayed, 
- the duration for each image can be defined.

The secondary images are required only if the secondary image is enabled.

If not all the required images are present, the application will not start and display an error message `Error loading images`.

For the `error` state, only one image can be displayed.

**Resource naming pattern**

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

## Motion sensor configuration
Install the motion sensor as preconised by the `Eurecam` manufacturer. Ensure that the `COMPTIPIX 3D` motion sensor is running, available in your local network. Connect to its Web server, and in the menu `Occupancy` of the `Settings` pane, set max. threshold of people authorized to enter in the room at the same time. Qeedji does not bring support on the `COMPTIPIX 3D` motion sensor configuration.

## Qeedji device configuration
Set the IP address of the COMPTIPIX 3D device (for example: `192.168.0.100`) as `field1` variable value of your Qeedji device.
