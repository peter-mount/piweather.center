/*
 * openSCAD for a bracket to hold an instrument on the side of a pole
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

include <TextGenerator.scad>

module main_spar() {
	translate([0,-7.5,0]) union() {
		difference() {
				union() {
					// The main spar
					difference() {
						union() {
							// Main spar but rounded outer edge for safety
							cube([172-5,25,15]);
							translate([172-5,25,5]) rotate([90,0,0]) cylinder(h=25,r1=5,r2=5);
							translate([172-5,25,10]) rotate([90,0,0]) cylinder(h=25,r1=5,r2=5);
							translate([172-5,0,5]) cube([5,25,5]);
						}
						// Cut out top grooves
						translate([-1,2,10.1]) cube([175,9.5,5]);
						translate([-1,13.5,10.1]) cube([175,9.5,5]);
						// Cut out bottom grooves
						translate([-1,2,-0.1]) cube([175,9.5,5]);
						translate([-1,13.5,-0.1]) cube([175,9.5,5]);
						// Inside of top mounting, repeated below
						translate([172-18.7,4.5,10]) cube([16,16,8]);
					}
					// Bracing on the bottom only
					for(x=[0:6]) {
						assign(tx=30+(x*17.3), ty=y*2*13.5) {
							translate([tx,1,0])
								linear_extrude(height=9.5) polygon(
									points=[[0,12],[10,0],[20,12],[3,11],[10,1.5],[17.5,11]],
									paths=[[0,1,2],[3,4,5]]);
							translate([tx,24,0]) scale([1,-1,1])
								linear_extrude(height=9.5) polygon(
									points=[[0,12],[10,0],[20,12],[3,11],[10,1.5],[17.5,11]],
									paths=[[0,1,2],[3,4,5]]);
						}
					}
					// Mounting bracket on top
					difference() {
						translate([172-18.7,4.5,10]) cube([16,16,8]);
						translate([172-16.2,7,10]) cube([10.8,10.8,9]);
					}
				}
				// Mounting hole underneath
				translate([172-16.2,7,0]) cube([10.8,10.8,8]);
				// M4 bolt hole
				translate([172-10.8,12.5,-0.1]) cylinder(h=16,r1=2,r2=2);
				// Rain Drainage holes
				translate([152,6.5,-0.1]) cylinder(h=16,r1=0.5,r2=0.5);
				translate([152,8.5,-0.1]) cylinder(h=16,r1=0.5,r2=0.5);
				translate([152,16.5,-0.1]) cylinder(h=16,r1=0.5,r2=0.5);
				translate([152,18.5,-0.1]) cylinder(h=16,r1=0.5,r2=0.5);
			}
		}
}

// The backing plate with the mast present
module mast_plate() {
	union() {
		//translate([0,-22.5,-15.5]) cube([2,55,45]);
		translate([0,-20,-15.5]) cube([2,50,45]);
		translate([0,-22.5,-13]) cube([2,20,37]);
		translate([0, 12.5,-13]) cube([2,20,37]);
		// Rounded corners
		translate([0,-20,-13]) rotate([0,90,0]) cylinder(h=2,r1=2.5,r2=2.5);
		translate([0,-20,24]) rotate([0,90,0]) cylinder(h=2,r1=2.5,r2=2.5);
		translate([0,30,-13]) rotate([0,90,0]) cylinder(h=2,r1=2.5,r2=2.5);
		translate([0,30,24]) rotate([0,90,0]) cylinder(h=2,r1=2.5,r2=2.5);

		// The mast
		translate([0,5,-15.5]) cylinder(h=45,r1=12.4,r2=12.4);
	}
}

module main_spar_bracket() {
	difference() {
		union() {
			main_spar();
			mast_plate();
			mast_bracket_flange();
			translate([6,6,5.5]) rotate([90,0,0]) linear_extrude(height=2) polygon(
				points=[[0,25],[32,0],[0,-25]],
				paths=[[0,1,2]]
			);
			translate([6,17.5,5.5]) rotate([90,0,0]) linear_extrude(height=2) polygon(
				points=[[-6,25],[32,0],[-6,-25]],
				paths=[[0,1,2]]
			);
			translate([6,-5.5,5.5]) rotate([90,0,0]) linear_extrude(height=2) polygon(
				points=[[-6,25],[32,0],[-6,-25]],
				paths=[[0,1,2]]
			);
		}
		// Mast space
		translate([0,5,-25]) cylinder(h=60,r1=9.4,r2=9.4);
		translate([-14,-23.5,-16.5]) cube([14,57,47]);
		translate([-1,-25,-25.5]) cube([14,57,10]);
		translate([-1,-25,26.5]) cube([14,60,10]);
	}
}

module mast_bracket_flange() {
	translate([0,5,-15.5]) cylinder(h=2,r1=17,r2=17);
	translate([0,5,5]) cylinder(h=2,r1=20,r2=20);
	translate([0,5,24.5]) cylinder(h=2,r1=17,r2=17);
}


module mast_bracket() {
	difference() {
		union() {
			mast_plate();
			mast_bracket_flange();
		}
		// Mast space
		translate([0,5,-25]) cylinder(h=60,r1=9.4,r2=9.4);
		translate([-14,-23.5,-16.5]) cube([14,57,47]);
		translate([-1,-25,-25.5]) cube([14,57,10]);
		translate([-1,-25,26.5]) cube([14,60,10]);
	}
}

// Everything together
difference() {
	union() {
		main_spar_bracket();
		// in translate set x to 0 to see it together, 5 for printing keeping the two
		// components separate
		rotate([0,0,180]) translate( [5,-10,0]) mast_bracket();
	}

	// M6 bolt holes
	union() {
		translate([-12.5,-16,-10]) rotate([0,90,0]) cylinder(h=20,r1=3,r2=3);
		translate([-12.5,-16,20]) rotate([0,90,0]) cylinder(h=20,r1=3,r2=3);
		translate([-12.5,26,-10]) rotate([0,90,0]) cylinder(h=20,r1=3,r2=3);
		translate([-12.5,26,20]) rotate([0,90,0]) cylinder(h=20,r1=3,r2=3);
	}
}

// Finally the text. DO NOT put this un main_spar_bracket() as that crashes OpenSCAD
rotate([90,0,0]) translate([30,4.5,6.5]) scale([1,1,3]) drawtext("Mark II Weather Station");
rotate([90,0,0]) translate([165,4.5,-19.75]) scale([-1,1,3]) drawtext("http://piweather.center");
