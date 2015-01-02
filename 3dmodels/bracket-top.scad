/*
 * openSCAD for the top bracket which holds the windvane and anemometer
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

/*
 * The number of beams to generate.
 *
 * This must be at least 2 although higher than 5 is too many
 */
beams=2;

// Add an additional instrument at the center - 0 for no, 1 for yes
add_center_instrument=0;
// Add a spike on the top of the center to direct rain away
// mutually exclusive with add_center_instrument
add_center_spike=1;

// Optional inner lug on the instrument holder
includeInstrumentLug=0;

// Original had a small lug to orient the pole, this will enable that
includeMastLug=0;

// Include text? As it doesn't work in the prototype disable for now
includeText=0;

// Various bolt sizes, 4=M4, 6=M6 etc
instrumentBoltSize=4;
mastBoltSize=4;

/****************************************************************************
 * Do not change these settings
 ****************************************************************************/
// Beam length. The Maplin one is 130mm but thats too big for my printer but
// 100 will still fit the instruments
beam_length=100;

include <TextGenerator.scad>

// Fit everything together
for(beam=[0:(beams-1)]) {
	rotate([0,0,360*beam/beams]) union() {
		instrument();

		if(includeText) {
			// Various bits of text on first 2 beams
			if(beam==0) {
				translate([18,-2,24.5]) scale([0.5,0.5,1.5]) drawtext("http://piweather.center");
			} else if(beam==1) {
				translate([18,-2,24.5]) scale([0.5,0.5,1.5]) drawtext("Mark II Weather Station");
			}
		}
	}
}

// Enable to include an additional instrument holder at the centre
if(add_center_instrument) instrument_holder();

mast_attachment();

/*
 * A spar, drawn from the center to the outer extremity.
 * We do this rather than just one so later we can do 3 or 4 arm versions
 */
module instrument_spar() {
	union() {
		difference() {
			// Main spar
			translate([0,-7.5,0]) cube([beam_length,15,25]);
			// Cutout to make it a T-Beam
			for(y=[0:1]) {
				translate([5, 3-(11*y),2]) cube([beam_length-10,5.5,21]);
			}
		}
		// Bracing
		translate([0,-7.5,11.5]) cube([beam_length,15,2]);
		for( x=[0:4]) {
			assign(tx=3+(x*17.3), ty=7.5) {
				translate([tx,ty,12]) rotate([90,0,0])
				linear_extrude(height=15) polygon(
					points=[[0,13],[10,0],[20,13],[3,11],[10,1.5],[17.5,11]],
					paths=[[0,1,2],[3,4,5]]);
				translate([tx,ty-15,13]) rotate([-90,0,0])
				linear_extrude(height=15) polygon(
					points=[[0,13],[10,0],[20,13],[3,11],[10,1.5],[17.5,11]],
					paths=[[0,1,2],[3,4,5]]);
			}
		}
	}
}

/*
 * Module for defining an instrument holder
 */
module instrument_holder() {
	difference() {
		// Outer sheaf
		union() {
			cylinder(h=41,r=11);
			cylinder(h=25-2,r1=11,r2=14);
			translate([0,0,25-2]) cylinder(h=2,r=14);
		}

		translate([0,0,25.1]) difference() {
			cylinder(h=20,r=9.3);
			if(includeInstrumentLug) {
				// Inner lug which goes into the instrument
				difference() {
					cylinder(h=25,r=5.25);
					cylinder(h=25,r=2.8);
					for(x=[0:3])
						rotate([0,0,x*(360/3)]) translate([0,4.5,0]) cylinder(h=25,r=1.25);
				}
			}
		}
		// M4 bolt hole
		translate([0,20,30]) rotate([90,0,0]) cylinder(h=40,r=instrumentBoltSize/2);
	}
}

/*
 * Combines a spar and an instrument
 */
module instrument() {
		instrument_spar();
		translate([beam_length,0,0]) instrument_holder();
}

/*
 * The mast attachment at the centre
 */
module mast_attachment() {
	translate([0,0,-25]) difference() {
		union() {
			// Flange
			translate([0,0,25]) cylinder(h=2,r=15);

			// Main body
			cylinder(h=25,r=9.2);

			// Lug
			if(includeMastLug) translate([0,9.2,25-2]) cylinder(h=2,r=2.5);

			// Brace on to the body
			translate([0,0,25+2]) cylinder(h=25-2,r=15);

			// Optional cap on top
			if(add_center_spike) translate([0,0,50]) cylinder(h=5,r1=15,r2=0);
		}
		// Hole at bottom
		translate([0,0,-5]) cylinder(h=25,r=6);
		// M4 Bolt hole
		translate([0,20,12]) rotate([90,0,0]) cylinder(h=40,r=mastBoltSize/2);
	}
}
