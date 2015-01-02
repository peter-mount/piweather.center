/*
 * openSCAD for The Mark II Stephenson Screen
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

// Set to 1 to enable, 0 to disable

// The components, when printing make certain only one is enabled
// The main body
includeBody=0;
// The door, usually print separate to the body
includeDoor=1;

// Valid only when includeBody=1
includeSensor=1;
includeAccess=1;

// The Base size of the screen, limited by the size of the printer
basesize=120;

// When includeAccess=1 the inner diameter of the cable access in mm
accessDiameter=8;

// Various bolt sizes, for M4 use 4, M6 6 etc
roofBoltSize=4;
doorHingeBoltSize=4;
doorLockBoltSize=4;

/**********************************************************************
 * DO NOT CHANGE ANYTHING BELOW THIS POINT
 **********************************************************************/

// The roof, usually printed with the body but can be separated
includeRoof=includeBody;

basehalf=basesize/2;
frameheight=2*basesize/3;
framehalf=frameheight/2;

doorwidth=basesize-28;
doorhalf=doorwidth/2;
doorheight=frameheight+35;

// By default only include roof mounting if one not both of body & roof are included
// i.e. if printing as one unit (minus the door) then no need for the mounts
includeRoofMounting = 0;//!(includeBody&&includeRoof) && (includeBody||includeRoof);

// Now render everything

//doorframe();
difference() {
	union() {
		difference() {
			union() {
				if(includeBody) {
					body();
					if(includeSensor) {sensor();}
				}
				if(includeRoof) {roof();}
				if(includeDoor) {doorFrameInner();}
			}

			if(includeRoof || includeBody) doorframe();
		}
		if(includeDoor) door();
	}

	bolts();

	if(includeBody && includeAccess) {
		access();
	}
}

module base() {
	// Base platform
	difference() {
		union() {
			translate([-basehalf,-basehalf,-5]) cube([basesize,basesize,5]);

			// Access lip
			if(includeAccess)
				translate([(2*basehalf/3)-10,35,0]) cylinder(h=5,r=(accessDiameter/2)+2);
		}

		// Access hole
		if(includeAccess)
			translate([(2*basehalf/3)-10,35,-6]) cylinder(h=12,r=accessDiameter/2);
	}

	for(c=[0:4])
		rotate([0,0,90*c]) translate([-basehalf,-basehalf,0]) cube([15,15,frameheight]);
}

// Renders the left, right & rear screens
module screen() {
	// Edges
	translate([-basehalf+15,-7.2,0]) cube([basesize-28,19.2,10.5]);
	translate([-basehalf+15,-7.2,frameheight-16.5]) cube([basesize-28,15,16.5]);
	
	translate([-basehalf+10,-7.2,0]) cube([8,19,frameheight]);
	translate([ basehalf-15,-7.2,0]) cube([5,19.2,frameheight]);
	
	// center support
	translate([ -7.5,-7.2,2.5]) cube([15,19.2,frameheight-15]);
	
	// interim supports
	translate([-25.5,-7.2,2.5]) cube([3,19.2,frameheight-15]);
	translate([ 25.5,-7.2,2.5]) cube([3,19.2,frameheight-15]);
	
	// Slats
	for(y=[0:6]) {
		slat(basesize-24.7,20+8*y);
	}

}

module slat(w,y) {
	translate([-basehalf+13,2.5,y])
		rotate([90,0,90])
		linear_extrude(height=w)
		polygon(
			points=[ [-9.5,-11.5], [0,0], [9.5,-11.5], [9.5,-9.5], [0,2], [-9.5,-9.5] ],
			paths=[ [0,1,2,3,4,5] ]
		);
}


// Renders the main body
module body() {
	base();
	// The left, right & rear screens
	for(s=[0:3]) {
		if(s!=1)
			rotate([0,0,-90+(90*s)])
			translate([0,-basehalf+7.2,0])
			screen();
	}

	// Roof mounting
	for(s=[0:3])
		if(s!=1)
			rotate([0,0,-90+(90*s)])
			translate([0,-basehalf+7.2,0])
			roofmounting();
}

// Inner door frame
module doorFrameInner() {
	if(includeBody||includeRoof) {
		translate([-doorhalf-2,-basehalf+5,0]) cube([10,10,doorheight]);
		translate([doorhalf-5,-basehalf+5,0]) cube([10,10,doorheight]);
		translate([-doorhalf-2,-basehalf+5,doorheight-15]) cube([doorwidth,10,15]);
		translate([-doorhalf-2,-basehalf+5,0]) cube([doorwidth,10,15]);
	}
}

// Door body
module door() {
	if(includeDoor) {
		// Inner door frame which slots into the front
		translate([-doorhalf+.5,-basehalf-1,1]) cube([10,11,doorheight-1.5]);
		translate([doorhalf-10.5,-basehalf-1,1]) cube([10,11,doorheight-1.5]);
		translate([-doorhalf+.5,-basehalf-1,doorheight-15.5]) cube([doorwidth-1,11,15]);
		translate([-doorhalf+.5,-basehalf-1,1]) cube([doorwidth-1,11,14.5]);

		// Outer door frame which fits over everything
		difference() {
			translate([-basehalf+10,-basehalf-10,-1]) cube([basesize-20,10,doorheight+2]);
			translate([-doorhalf+10,-basehalf-11,15]) cube([doorwidth-20,12,doorheight-30]);
		}

		// Door slats
		translate([-5,-basehalf-10,15]) cube([10,20,doorheight-30]);
		translate([10,-basehalf-3,5]) for(y=[0:10]) {
			slat(basesize-44.7,20+8*y);
		}
	}
}

// used in a difference, cuts the door out of both body & frame.
// It also ensures holes line up
module doorframe() {
	// Cut out the door
	translate([-doorhalf,-basehalf-1.5,0]) cube([doorwidth,11,doorheight]);
}

// Optional sensor frame - i.e. attach thermometer & barometer on this
module sensor() {

	// Side struts
	for(s=[0:1]) {
		translate([ s ? ((basehalf/3)-5) : (-basehalf/3), -2.5,-0.5])
		union() {
			cube([5,5,framehalf]);
			rotate([90,0,90])
				linear_extrude(height=5)
				polygon(
					points=[ [0,0], [-15,0], [0,framehalf/2], [5,framehalf/2], [20,0] ],
					paths=[ [0,1,2,3,4] ]
				);
			translate([5*s,s?0:5,0])
				rotate([90,0,s*180])
				linear_extrude(height=5)
				polygon(
					points=[ [0,0], [-15,0], [0,framehalf/3] ],
					paths=[ [0,1,2] ]
				);
		}
	}

	// Mounting slats
	for(y=[0:2])
		translate([-(basehalf/3),-2.5,framehalf-0.75-(15*y)]) cube([2*basehalf/3,5,5]);
}

// The roof mounting - only applies if roof & body are not created
module roofmounting() {
	if(includeRoofMounting)
		translate([-7.5,0,frameheight-2]) union() {
			translate([0,2.7,0]) cube([15,22.3,2]);
			translate([0,2.7,-10]) cube([15,12,10]);
			for(i=[0:1])
				translate([11*i,0,0])
				rotate([90,0,90])
				linear_extrude(height=4)
				polygon(
					points=[ [14.6,0], [25,0], [14.6,-10] ],
					paths=[ [0,1,2] ]
				);
		}
}

// Renders the roof
module roof() {
	difference() {
		// The main roof
		union() {
			// Roof front
			translate([-basehalf,-basehalf,frameheight]) cube([basesize,5,47.3]);

			// Roof
			translate([-basehalf-5,basehalf+15,frameheight])
				rotate([160,0,0])
				linear_extrude(height=6) polygon(
					points=[ [0,0], [basesize+10,0], [basesize+10, basesize+30], [0,basesize+30] ],
					paths=[ [0,1,2,3] ]
				);
	
			// Roof frame
			if(includeRoofMounting)
				for(s=[0:3])
					if(s!=1)
						rotate([0,0,-90+(90*s)])
						translate([-basehalf+5,-basehalf+5,frameheight])
						cube([basesize-5,13,5]);

			// Roof mounting
			for(s=[0:3])
				if(s!=1)
					translate([0,0,frameheight*2])
					rotate([0,0,-90+(90*s)])
					translate([0,-basehalf+7.2,0])
					rotate([0,-180,0])
					roofmounting();

		// Sides
			for(s=[0:1])
				translate([s?basehalf:(6-basehalf),-basehalf,frameheight])
				rotate([0,-90,0])
				linear_extrude(height=6) polygon(
					points=[ [0,0], [0,basesize], [5,basesize], [49,0] ],
					paths=[ [0,1,2,3] ]
			);
		}

		// Remove cruft from above the front roof
		translate([-basehalf-5,basehalf+15,frameheight+16.3])
			rotate([160,0,0])
			linear_extrude(height=15.35) polygon(
				points=[ [0,0], [basesize+10,0], [basesize+10, basesize+30], [0,basesize+30] ],
				paths=[ [0,1,2,3] ]
			);
	}
}

module bolts() {
	// roof bolts if required
	if(includeRoofMounting)
		for(s=[0:3])
			rotate([0,0,-90+(90*s)])
			translate([ 0, 32.5, frameheight-10])
			cylinder(h=17,r=roofBoltSize/2);

	// Door bolts - recess to hold a M6 nut & associated hole for bolt
	if(includeBody||includeDoor)
		for(b=[0:1])
			translate([0,-basehalf+11,b?doorheight-7.5:7.5])
				rotate([90,0,0])
				union() {
					// Nut size is ~10mm so 12mm will give room for adjustment
					cylinder(h=10,r=6);
					// M6 bolt hole
					translate([0,0,-16]) cylinder(h=40,r=3);
				}
}
