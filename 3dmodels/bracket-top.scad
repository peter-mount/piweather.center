/*
 * openSCAD for the top bracket which holds the windvane and anemometer
 */

/*
 * The number of beams to generate.
 *
 * This must be at least 2 although higher than 5 is too many
 */
beams=2;

// Add an additional instrument at the center - 0 for no, 1 for yes
add_center_instrument=0;

/****************************************************************************
 * Do not change these settings
 ****************************************************************************/
// Beam length. The Maplin one is 130mm but thats too big for my printer but
// 100 will still fit the instruments
beam_length=100;


include <TextGenerator.scad>

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
			cylinder(h=41,r1=11,r2=11);
			cylinder(h=25-2,r1=11,r2=14);
			translate([0,0,25-2]) cylinder(h=2,r1=14,r2=14);
		}
		// Inner lug which goes into the instrument
		translate([0,0,25.1]) difference() {
			cylinder(h=20,r1=9.3,r2=9.3);
			// Inner lug
			difference() {
				cylinder(h=25,r1=5.25,r2=5.25);
				cylinder(h=25,r1=2.8,r2=2.8);
				cylinder(h=25,r1=2.8,r2=2.8);
				for(x=[0:3]) {
					rotate([0,0,x*(360/3)]) translate([0,4.5,0]) cylinder(h=25,r1=1.25,r2=1.25);
				}
			}
		}
		// M4 bolt hole
		translate([0,20,30]) rotate([90,0,0]) cylinder(h=40,r1=2,r2=2);
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
			translate([0,0,25]) cylinder(h=2,r1=15,r2=15);
			// Main body
			cylinder(h=25,r1=9.2,r2=9.2);
			// Lug
			translate([0,9.2,25-2]) cylinder(h=2,r1=2.5,r2=2.5);
			// Brace on to the body
			translate([0,0,25+2]) cylinder(h=25-2,r1=15,r2=15);
			// Optional cap on top
			//translate([0,0,50]) cylinder(h=5,r1=15,r2=0);
		}
		// Hole at bottom
		translate([0,0,-5]) cylinder(h=25,r1=6,r2=6);
		// M4 Bolt hole
		translate([0,20,12]) rotate([90,0,0]) cylinder(h=40,r1=2,r2=2);
	}
}

/*
 * Fit everything together
 */
for(beam=[0:(beams-1)]) {
	rotate([0,0,360*beam/beams]) union() {
		instrument();

		// Various bits of text on first 2 beams
		if(beam==0) {
			translate([18,-2,24.5]) scale([0.5,0.5,1.5]) drawtext("http://piweather.center");
		} else if(beam==1) {
			translate([18,-2,24.5]) scale([0.5,0.5,1.5]) drawtext("Mark II Weather Station");
		}
	}
}

// Enable to include an additional instrument holder at the centre
if(add_center_instrument) {instrument_holder();}

mast_attachment();
