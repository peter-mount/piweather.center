/*
 * openSCAD for the solar telescope component for piweather.center
 *
 * This file is not for real use - it's just to draw up some ideas
 * I've had for some time now & won't be actually built until the
 * Mark II station is up and running.
 *
 * The principle is to have a PI camera mounted on the end of a
 * tube with a Solar Filter mounted on the front.
 * 
 * In theory the sun's diameter will be large enough, although a
 * 2x or 4x lens may be added in a future design.
 * This is based on observations seen with the prototype Sky Camera
 * during the summer of 2014.
 *
 * Copyright 2015 Peter T Mount
 * 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 
 * http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Center tube radius
tubeRadius=30;
tubeLength=100;

/***************************************************************************
 * Do not modify anything below this
 ***************************************************************************/

tubeDiameter=tubeRadius*2;

include <PiCameraMount.scad>
include <FilterHolders.scad>

/***************************************************************************/

// The tube assembly
translate([6,0,0]) union() {
	tube();
	rotate([180,180,0]) cameraHolder(0);
	translate([tubeLength-2,0,0]) filterHolder();
}

// The individual components
translate([-6,-70,0]) cameraHolder(1);
translate([0,70,0]) filterHolder();

/***************************************************************************/

// The main optical tube.

module tube() {
	rotate([0,90,0])
		union() {
			// The tube
			difference() {
				cylinder(h=tubeLength,r=tubeRadius);
				translate([0,0,-1]) cylinder(h=tubeLength+2,r=tubeRadius-3);
			}
			// Support Struts
			for(s=[0:11])
				if(s!=3 && s!=9)
					rotate([0,0,30*s])
					translate([-tubeRadius-2.2,0,0])
					cylinder(h=tubeLength,r=2);
		}
}

// The camera holder
module cameraHolder(mode) {
	piCamera(tubeDiameter,mode);
}

// The filter holder
module filterHolder() {
	filterSimpleHolder(tubeDiameter,piCameraYOffset,tubeRadius-5);
}

