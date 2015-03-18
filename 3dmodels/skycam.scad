/*
 * openSCAD for the movable Sky Camera for the Mark II weather station.
 *
 * The idea of this camera is this:
 *
 * 1 The camera will be mounted within a dummy dome security camera,
 *   pointing upwards.
 * 2 It will be mounted on a 2 axis servo camera mount.
 *   This will be oriented so it will give a 180 view east to west.
 *
 * These are optional, if there is room in the camera dome.
 * 3 A 3D printed mount for this camera will polar align it.
 * 4 On top will be a 2x lens servo controlled so we can have 1x & 2x
 * 5 On top will be a solar filter, servo controlled.
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

// Include the telescope mount in the model
includeFront=0;

// Include the backplate in the model
includeBack=1;

includeFilter=0;

// ======================================================================
// Do not modify these

include <Common.scad>
include <PiCameraMount.scad>

// Camera mount specs
// width/height of inside the mounting
mountSize=30;
mountHalf=mountSize/2;

// Height of the 2 side forks
mountFork=15;

// Camera mount
width=40;

// Servo dimensions
servoWidth=13;
servoHeight=24;
// Offset is from top to center of servo shaft
servoOffset=5.5;
// Depth from flange to base
servoDepth=17;
// Height from base of flange to top (not including shaft)
servoTop=11;

// **********************************************************************
if(includeFront) front();
if(includeBack) back();
if(includeFilter) filter();

// **********************************************************************

// Filter wheel mount
module filter() {
	filterServo();
}

module filterServo() {
	translate([-6,0,0])
	difference() {
		// Camera mount + Servo mount
		union() {
			piCamera(width,2);
			translate([0,-width/2,(width+piCameraYOffset)/2])
				cube([5,width,20]);
		translate([0,-servoOffset-servoHeight/2,5+(width+piCameraYOffset)/2]) servo();
		}
		// Light path to camera
		translate([-1,0,piCameraYOffset/4])
			rotate([0,90,0]) 
			cylinder( h=24,r=12);
		// Servo mount
	}
}

// A servo. In reality used to cut out the servo from items
module servo() {
	translate([-servoTop,0,0]) union() {
		// The body - aka the hole for it
		cube([servoDepth+servoTop,servoHeight,servoWidth]);
		// The shaft
		rotate([0,90,0])
			translate([-servoOffset-1,servoHeight-servoOffset,-4])
			cylinder(h=4,r=2);
	}
}

// Camera mount front
module front() {
	piCamera(width,0);
}

// Camer mount back with fixing to the pan & tilt mount
module back() {
	difference() {
		union() {
			piCamera(width,3);
			translate([8,-mountHalf-2,-mountHalf]) cube([mountFork-1,6,mountSize]);
			translate([8,mountHalf-4,-mountHalf]) cube([mountFork-1,6,mountSize]);
			translate([8,-mountHalf,12]) cube([mountFork-1,mountSize,3]);
		}
		translate([5,0,0]) cameramount();
	}
}

// The camera mount - used to remove from the backing
module cameramount() {
	union() {
		translate([3,-mountHalf-3,-mountHalf/2]) cube([mountFork,3,mountFork]);
		translate([3,-mountHalf-1,-mountHalf/2]) cube([2,2,mountFork]);
		translate([3,mountHalf,-mountHalf/2]) cube([mountFork,3,mountFork]);
		translate([3,mountHalf-1,-mountHalf/2]) cube([2,2,mountFork]);
		translate([mountFork+3,-mountHalf-3,-mountHalf]) cube([2,mountSize+6,mountSize]);
	}
}
