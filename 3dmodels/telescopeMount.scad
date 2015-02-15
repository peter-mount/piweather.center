/*
 * openSCAD for an adapter for attaching the Raspberry PI Camera Module 
 * to a telescope.
 *
 * This model comes in two parts, the eyepiece mount which attaches directly
 * to the telescope (no eyepiece required) and the back plate.
 *
 * However it's intended that the backplate be incorporated into another model
 * so that the PI itself is mounted with the camera.
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
includeBack=0;

// The required eyepiece diameter in inches.
// Most telescopes these days are 1.75"
eyepiece=1.75;
// Some big scopes have 2" barrels
//eyepiece=2;
// Older ones have 0.925" barrels
//eyepiece=0.925;


// ======================================================================
// Do not modify these

// eyepiece diameter, adjusting for extruder diameter
eyediam=(eyepiece*25.4)-0.5;
eyerad=eyediam/2;
width=max(50,eyediam+10);


/***************************************************************************
 * Required modules
 */
include <Common.scad>
include <PiCameraMount.scad>

if(includeFront) front();
else if(includeBack) back();
else {
	translate([6,0,0]) rotate([0,0,180]) front();
	translate([-6,-width*3/2,0]) back();
}

// **********************************************************************

module front() {
	union() {
		piCamera(width,0);
		translate([-24,0,0]) rotate([0,90,0]) difference() {
			cylinder( h=24,r=eyerad);
			translate([0,0,-2]) cylinder( h=27,r=eyerad-2);
		}
	}
}

module back() {
	piCamera(width,3);
}
