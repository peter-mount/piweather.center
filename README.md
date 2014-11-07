piweather.center
================

Raspberry PI based Weather Station

This is an experimental standalone application based on raspistill for running a
webcam using the Raspberry PI Camera (Normal or NOIR).

Prerequisites:

* The PI must have gcc etc installed and up to date.
* libmicrohttpd must be installed
* imagemagick

sudo apt-get install libmicrohttpd10 imagemagick

Next edit /etc/default/tmpfs as root and change

#RAMTMP=no

to

RAMTMP=yes

In /boot/config.txt

change gpu_mem to 64

add disable_camera_led=1

Now reboot.

For PiWeather Board support:

sudo vi /etc/modprobe.d/raspi-blacklist.conf
comment out the line blacklist i2c-bcm2708

sudo vi /etc/modules
add i2c-dev on a new line
sudo apt-get update
sudo apt-get install i2c-tools
sudo apt-get install libi2c-dev
sudo adduser pi i2c
sudo reboot

Once back run:
i2cdetect -y 1

You should see 4e appear within that list, indicating that the i2c board is present and is being detected.

Netbeans config:

If using netbeans on another machine with the PI as the build slave then you must
have libmicrohttpd installed there as well.

Either use apt-get on that machine (or you're distribution's installer).

e.g. on Gentoo it's: sudo emerge libmicrohttpd

