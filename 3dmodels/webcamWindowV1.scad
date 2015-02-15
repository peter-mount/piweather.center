/*
 * openSCAD for a WebCam using the Raspberry PI Camera Module and a
 * Raspberry PI which is housed in it's own case (in this instance a
 * Pimoroni case is used).
 *
 * Now this is for an internal webcam - specifically this one is used
 * on http://maidstoneweather.com for displaying the western horizon
 * visible outside my bedroom window, so this one sits on a window sill.
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

// Enable each component. When printing have just 1 of these defined

// Front panel
includeFront=0;

// The main body
includeBody=0;

// The rear panel
includeRear=0;

// Camera & Wide Angle Lens Mount
includeCamera=1;

// Wide Angle Lens cap - holds the lens in position
includeCap=0;

// Stray Light shield
includeShield=0;

// Set to 1 to include a hole for an ethernet cable, 0 if using WiFi
includeEthernet=0;

// The overall height of the camera
webCamH=125;
webCamW=115;


/***************************************************************************
 * Do not modify these
 */
webCamWH=webCamW/2;
webCamHH=webCamH/2;

// The camera plate size, keep at 50mm
webCamPS=50;
webCamPSH=webCamPS/2;

// The height from the base of where the webcam is mounted
webCamCH=webCamH-webCamPSH-10;

/***************************************************************************
 * Required modules
 */
include <Common.scad>
include <PiCameraMount.scad>

if(includeFront) frontPanel();
if(includeBody) body();
if(includeRear) rear();
if(includeCamera) camera();
if(includeCap) cameraCap();
if(includeShield) shield();

/***************************************************************************/

// The main front panel
module frontPanel() {
	difference() {
		// Base plate
		//translate([-webCamWH,-1,0]) cube([webCamW,6,webCamH]);
		union() {
			translate([-webCamWH,-1,0])
				cube([webCamW,6,webCamH-webCamPS]);
			translate([0,5,webCamHH+webCamPS/3])
				rotate([90,90,0])cylinder(h=6,r=webCamWH);
		}

		// Cutout where the camera will be
		translate([-webCamPSH,-1.5,webCamCH-webCamPSH-4]) cube([webCamPS,9,webCamPS+8]);
		// The bottom bolts
		translate( [-webCamWH+15, 13,10]) rotate([90,90,0]) cylinder(h=20,r=M4);
		translate( [ webCamWH-15, 13,10]) rotate([90,90,0]) cylinder(h=20,r=M4);
	}

	// The camera mount
	difference() {
		// The camera back plate
		translate([0,-7,webCamCH])
			rotate([0,0,90])
			piCamera(webCamPS,1);
		// Cutout to feed camera cable
		translate([-piCameraRibbonWidth/2,-1,webCamCH-webCamPSH])
			cube([piCameraRibbonWidth,9,7]);
	}
}

// The main body
module body() {
	// Base
	translate( [-webCamWH, 5,0]) cube([webCamW,80,6]);

	// top
	translate([0,85,webCamHH+webCamPS/3])
		difference() {
			rotate([90,90,0])cylinder(h=80,r=webCamWH);
			translate([0,9,0]) union() {
				rotate([90,90,0])cylinder(h=100,r=webCamWH-5);
				translate([-webCamWH,-90,-webCamHH-webCamPS/7]) cube([webCamW,100,webCamHH]);
			}
		}

	// Sides
	translate( [-webCamWH, 5,0]) cube([6,80,6+webCamHH+webCamPS/7]);
	translate( [webCamWH-6, 5,0]) cube([6,80,6+webCamHH+webCamPS/7]);

	// Top bolts
	difference() {
		intersection() {
			union() {
				translate([0,5,webCamCH]) rotate([0,0,90]) piCamera(webCamPS,2);
				translate( [-webCamWH, 5,webCamH-5]) cube([webCamW,6,12]);
				translate( [23, 5,webCamH-6-30]) cube([webCamWH-23,6,webCamHH-23]);
				translate( [-webCamWH, 5,webCamH-6-30]) cube([webCamWH-23,6,webCamHH-23]);
			}
			translate([0,95,webCamHH+webCamPS/3])
				rotate([90,90,0]) cylinder(h=100,r=webCamWH);
		}
		translate([-26,4,webCamCH-(webCamPS/2)-8]) cube([webCamPS*2,10,7+webCamPS/2]);
	}

	// Bottom bolts
	difference() {
		translate( [-webCamWH, 5,6]) cube([webCamW,6,10]);
		translate( [-webCamWH+15, 13,10]) rotate([90,90,0]) cylinder(h=20,r=M6);
		translate( [webCamWH-15, 13,10]) rotate([90,90,0]) cylinder(h=20,r=M6);
	}

	// Back panel bolts
	rearBolts(68,1);
}

module rear() {
	translate([0,86,0]) {
		difference() {
			union() {
				translate([-webCamWH,-1,0])
					cube([webCamW,6,webCamH-webCamPS]);
				translate([0,5,webCamHH+webCamPS/3])
					rotate([90,90,0])cylinder(h=6,r=webCamWH);
			}
			// Cut out the bolts from the rear
			rearBolts(-2,0);
			rearBolts(2,0);
			rearBolts(5,0);

			// rear ventilation holes
			for(s=[0:6])
				translate([-webCamWH+(webCamPS*2/3),-10,(3*webCamHH/3)+(6*s)])
					cube([webCamPS,20,3]);

			// Power access
				translate([webCamWH-35,-10,15])
					cube([10,20,10]);

			if(includeEthernet)
				translate([-webCamWH+25,-10,15])
					cube([15,20,10]);
		}

		rearBolts(-1,1);
	}
}

// Common to body() & rear() - used to add the rear mounting bolts
module rearBolts(y,h) {
	translate([ webCamWH-19,y,6]) difference() {
		cube([15,6,10]);
		if(h) translate([7.5,25,5]) rotate([90,90,0]) cylinder(h=50,r=M6);
	}
	translate([-webCamWH+ 6,y,6]) difference() {
		cube([15,6,10]);
		if(h) translate([7.5,25,5]) rotate([90,90,0]) cylinder(h=50,r=M6);
	}
	translate([-7.5,y,webCamH-5]) difference() {
		cube([15,6,10]);
		if(h) translate([7.5,25,5]) rotate([90,90,0]) cylinder(h=50,r=M6);
	}
}

// The camera mounting
module camera() {
	translate([0,-7,webCamCH])
	difference() {
		union() {
		// Camera front plate
		rotate([0,0,90]) piCamera(webCamPS,0);

		// Normal plate has space for ribbon out of bottom, this masks it out
		// as we have our own
		translate([-piCameraRibbonWidth/2,0,-webCamPSH-5])
			cube([piCameraRibbonWidth,6,5]);
		
		// Pimoroni Wide Angle Lens Holder
		translate([0,0,2.5]) rotate([90,90,0]) cylinder(h=13,r=37.5/2);
	}
	translate([0,1,2.5]) rotate([90,90,0]) cylinder(h=20,r=35.5/2);
	}
}

// Cap that holds the wide angle lens in position
module cameraCap() {
	translate([0,-30,webCamCH+2.5])
		rotate([90,90,0])
		difference() {
			translate([0,0,-5]) cylinder(h=15,r=44/2);
			translate([0,0,-6]) cylinder(h=13,r=40/2);
			translate([0,0,6.5]) cylinder(h=3,r1=33/2,r2=30/2);
			translate([0,0,1]) cylinder(h=10,r=30/2);
		}	
}

// Stray Light Shield
module shield() {
	translate([0,-13,webCamCH])
		difference() {
			union() {
				rotate([0,0,90]) piCamera(webCamPS,2);
				rotate([90,90,0]) translate([0,0,60]) scale([0.75,0.75,1]) sphere(r=60);
				rotate([90,90,0]) translate([0,0,0]) cylinder(h=15,r=45/2);
				rotate([90,90,0]) translate([0,0,50+9]) cylinder(h=5,r=50);
			}
			rotate([90,90,0]) union() {
				translate([0,0,-7]) cylinder(h=72,r=40/2);
				translate([0,0,64]) scale([0.75,0.75,1]) sphere(r=60);
			}
		}	
}

