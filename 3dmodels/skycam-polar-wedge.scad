/*
 * openSCAD for a simple Equatorial Wedge to hold the Pan & Tilt camera
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

// The required latitude in degrees.
// Note:
//  1 this is a positive value, so 51.5N and 51.5S is still 51.5 here.
//  2 There's a bug in the math for latitudes 25 to 35 where the supports
//    don't reach the mount plate. You may have to modify the height manually
latitude=51.271749;

/* **********************************
 * * Do not modify below this point *
 * **********************************/

// Base plate to attach to the exterior case
difference() {
	union() {
		cube([50,60,5]);
		// Bolt mounts
		for(c=[0:3])
			translate([50*(c%2),60*(floor(c/2)%2),0])
			cylinder(h=6,r1=8,r2=6);
	}
	translate([7.5,5,-1]) cube([35,50,7]);

	// Bolt holes
	for(c=[0:3])
		translate([50*(c%2),60*(floor(c/2)%2),0])
		translate([0,0,-1]) cylinder(h=8,r=3.25);
}

// The camera mount plate
translate([0,20,0]) assign(
	alpha=90-latitude,
	height=55*sin(90-latitude),
	offset=55*cos(90-latitude)//-7.5
) {
	// Mount plate
	translate([7.5,4,1])
		rotate([alpha,0,0])
		cube([35,55,5]);

	// Front mount
	translate([7.5,0,0]) cube([35,8,5]);

	// Rear mount - usually part of rear of basePlate but accounts for
	// lower latitudes when it's closer to the front
	//translate([7.5,offset,0]) cube([35,5,5]);

	// Center support
	translate([22.5,0,0]) cube([5,35,5]);

	// Risers to rear of the mount plate
	assign( off2=35-offset)
		assign(
			// Height of supports
			h = offset<30 ? sqrt(off2*off2+height*height)*sin(alpha)
				: offset>35 ? sqrt(35*35+height*height)*sin(alpha)
			 	: height,
			// Prevent supports leaning backwards
			a=offset<30?atan(off2/height):0
		)
			for(i=[0,1,2])
				translate([7.5+(15*i),35,0])
					rotate([a,0,0])
					cube([5,5,h]);
}

