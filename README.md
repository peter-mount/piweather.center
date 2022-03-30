# piweather.center

This project is for a Weather Station based around the Raspberry PI 4B.

It originally started back in 2014 using a Raspberry PI 1A which operated for 5 years.
The project has been restarted in 2022 using current hardware.

The original code from 2014 is in the `original` branch of this repository.
The `master` branch contains the latest 2022 code.

## Setup

This requires a current version of Raspberry PI OS (formerly Raspbian) Bullseye running on a Raspberry PI 4B.
I'm using a 2Gb RAM version for my station, however you can use one with more ram if you need to.

For the sensors it uses the Pimoroni Weather Station Hat which handles the rain gauge and anemometer.

A separate temperature sensor is mounted outside the enclosure so that it is unaffected by the temperature of the PI.


## Cloud Camera

Although optional, a camera is used to measure cloud cover during the day.

With RaspiOS Bullseye it now uses the new libcamera framework which we use at the command level.
To keep SD card wear down to a minimum, this uses a ramdisk (hence why the minimum ram for the PI needs to be 2Gb).

### Issues with some third party cameras
I'm using a non-named camera of Amazon which connects to the CSI port like the standard cameras, however
the one I'm using is an 8MP camera with a 160° field of view.
My original 2014 station used a standard 5MP PI camera with a wide angle adaptor attached to it, but this camera is more compact.

However, it did not work with the new libcamera until I did the following:

In /boot/config.txt find the section

    # Automatically load overlays for detected cameras
    camera_auto_detect=1

and change it to:

    # Automatically load overlays for detected cameras
    camera_auto_detect=0
    dtoverlay=imx219

This worked as the camera is based on the imx219 chip but was not detected correctly.