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

/*
 Printing notes:

 As defined in the repo, this file will generate a single component,
 the spar and it's associated bracket. However, you can print several in
 a single print.

 For a MakerBot/Wanhao Duplicator 4 (i.e. my printer) you can get 2 prints
 horizontally as designed here, so for that:

 includeSpar=1;
 includeBracket=1;
 showExploded=1;
 numComponents=2;
 componentModulo=2;

 However, if you want to print more than 2 you could (not tested, just theory)
 print up to 4, possibly 6 in one go, by breaking it into 2 prints, one
 for the spar & one for the bracket.

 So for 4 copies:

 // print 1 - the spar
 includeSpar=1;
 includeBracket=0;
 numComponents=4;
 componentModulo=2;

 // print 2 - the bracket
 includeSpar=0;
 includeBracket=1;
 numComponents=2;
 componentModulo=2;

 Then when to generate the STL and pass it into ReplicatorG (or equivalent)
 you then rotate it so the mast side of the component is flat against the
 build plate.

 */

// Include or exclude text on the model.
// I find that the text doesn't print well, so keep disabled until a better
// solution is found
includeText=0;

// Usually these are both set to 1. Setting one of them to 0 will not render it.
// Useful if you want to do a batch of brackets but not the spars.
includeSpar=1;
includeBracket=1;
includeMountBlock=1;

// Print legacy mounting?
//
// When set to 1 this will include a mounting block on the end that will fit
// existing instruments (i.e. a Maplin Rain Gauge).
//
// Setting this to 0 will make the top end of the spar flat for attaching a
// different mount - i.e. a crossbar for mounting a custom component.
//
// For mounting piweather.center components this should be set to 0 as mast
// mountable components will have support for this designed in.
//
legacyMounting=0;

// Show the components exploded (i.e. apart).
// Set to 1 for printing.
// Set to 0 is useful to see it set together.
// Not of much use if one of includeSpar or includeBracket is 0
showExploded=1;

// The number of components to print.
// This can be useful to print more than one part, if your printer has room.
numComponents=1;

// Only of use if numComponents > 1, defines how many components are shown width wise
componentModulo=2;

// The length of the spar.
// The original was sparLengthmm but if printed vertically thats way too tall
// hence -25 to fit the printer
sparLength=172-25;

// The thickness of the backing plate. Increase if it's needs extra strength.
// The minimum is 2, although after looking at the prototype print 5mm is better
plateThickness=5;

// The plate bolt size in mm, so use 6 for an M6 bolt
plateBoltSize=6;

// The outer diameter of the mast
mastSize=20.5;

// The instrument bolt size in mm, so use 4 for an M4 bolt
mountBoltSize=4;

// The diameter of the rain drainage holes in the spar
drainageHoleSize=1.5;

/***************************************************************************
 * Do not change these
 ***************************************************************************/

// Spacing between the components
componentSpacing=showExploded?5:0;

// Y & Z dimensions, used when numComponents>1
componentWidth=70;
componentHeight=70;

include <TextGenerator.scad>

// The main loop, layout & generate each component
for(comp=[0:numComponents-1]) {
	translate([0, (comp%componentModulo)*componentWidth, floor(comp/componentModulo)*componentHeight])
		generateComponent();
}

// The main spar
module main_spar() {
	translate([0,-7.5,0]) union() {
		difference() {
				union() {
					// The main spar
					difference() {
						union() {
							// Main spar
							cube([sparLength-5,25,15]);
							
							// rounded outer edge for safety
							translate([sparLength-5,0,5]) cube([5,25,5]);
							translate([sparLength-5,25,5]) rotate([90,0,0]) cylinder(h=25,r1=5,r2=5);

							// Legacy mounting is also rounded on top
							if(legacyMounting)
								translate([sparLength-5,25,10]) rotate([90,0,0]) cylinder(h=25,r=5);
						}
						// Cut out top grooves
						translate([-1,2,10.1]) cube([sparLength+3,9.5,5]);
						translate([-1,13.5,10.1]) cube([sparLength+3,9.5,5]);
						// Cut out bottom grooves
						translate([-1,2,-0.1]) cube([sparLength+3,9.5,5]);
						translate([-1,13.5,-0.1]) cube([sparLength+3,9.5,5]);

						// Inside of top mounting, repeated below
						if(legacyMounting)
							translate([sparLength-18.7,4.5,10]) cube([16,16,8]);
						else
							translate([sparLength-20,-1,10]) cube([25,27,8]);
					}

					// Bracing on the bottom only
					for(x=[0:5]) {
						assign(tx=30+(x*17.3), ty=0) {
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

					if(legacyMounting) difference() {
						// Mounting bracket on top
						translate([sparLength-18.7,4.5,10]) cube([16,16,14]);
						translate([sparLength-16.2,7,10]) cube([10.8,10.8,15]);
					} else {
						// Non legacy has a cross bar on the top
						translate([sparLength-23,0,10]) cube([3,25,5]);
					}
				}

				// Instrument mounting bolt(s)
				if(legacyMounting) {
					// Cutout underneath
					translate([sparLength-16.2,7,-3.1]) cube([10.8,10.8,8]);
					// A single central bolt
					translate([sparLength-10.8,12.5,-0.1]) cylinder(h=16,r=mountBoltSize/2);
				} else {
					// No cutout as we have 2 bolts between the spars
					translate([sparLength-10.8,6.5,-0.1]) cylinder(h=16,r=mountBoltSize/2);
					translate([sparLength-10.8,18.5,-0.1]) cylinder(h=16,r=mountBoltSize/2);
				}

				// Rain Drainage holes - aligned to not break the spars underneath
				if(legacyMounting) {
					translate([sparLength-22,6,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
					translate([sparLength-22,10.5,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
					translate([sparLength-22,14.5,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
					translate([sparLength-22,19,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
				} else {
					translate([sparLength-25,4,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
					translate([sparLength-25,10.5,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
					translate([sparLength-25,14.5,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
					translate([sparLength-25,21,-0.1]) cylinder(h=16,r=drainageHoleSize/2);
				}
			}
		}
}

// The backing plate with the mast present
module mast_plate() {
	union() {
		translate([0,-20,-15.5]) cube([plateThickness,50,45]);
		translate([0,-22.5,-13]) cube([plateThickness,20,37]);
		translate([0, 12.5,-13]) cube([plateThickness,20,37]);

		// Rounded corners
		translate([0,-20,-13]) rotate([0,90,0]) cylinder(h=plateThickness,r1=2.5,r2=2.5);
		translate([0,-20,24]) rotate([0,90,0]) cylinder(h=plateThickness,r1=2.5,r2=2.5);
		translate([0,30,-13]) rotate([0,90,0]) cylinder(h=plateThickness,r1=2.5,r2=2.5);
		translate([0,30,24]) rotate([0,90,0]) cylinder(h=plateThickness,r1=2.5,r2=2.5);

		// The mast interface
		translate([0,5,-15.5]) cylinder(h=45,r=(mastSize/2)+3);
	}
}

// used in a union, the flanges onto the plate to add strength
module mast_bracket_flange() {
	translate([0,5,-15.5]) cylinder(h=2,r1=17,r2=17);
	translate([0,5,5]) cylinder(h=2,r1=20,r2=20);
	translate([0,5,24.5]) cylinder(h=2,r1=17,r2=17);
}

// Used in a difference, removes everything where the mast would be.
module mast_space() {
		translate([0,5,-25]) cylinder(h=60,r=mastSize/2);
		translate([-24,-23.5,-16.5]) cube([24,57,47]);
		translate([-1,-25,-25.5]) cube([14,57,10]);
		translate([-1,-25,26.5]) cube([14,60,10]);
}

// The actual complete spar side of the component
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
		mast_space();
	}
}

// The complete bracket side of the component
module mast_bracket() {
	difference() {
		union() {
			mast_plate();
			mast_bracket_flange();
		}
		mast_space();
	}
}

// Everything together
module generateComponent() {
	difference() {

		union() {
			if(includeSpar)
				main_spar_bracket();

			if(includeMountBlock)
				mountBlock();

			if(includeBracket)
				rotate([0,0,180]) translate( [componentSpacing,-10,0]) mast_bracket();
		}

		// Remove the plate bolt holes. Doing this here means both sides match up
		union() {
			translate([-12.5,-16,-10]) rotate([0,90,0]) cylinder(h=20,r=plateBoltSize/2);
			translate([-12.5,-16,20]) rotate([0,90,0]) cylinder(h=20,r=plateBoltSize/2);
			translate([-12.5,26,-10]) rotate([0,90,0]) cylinder(h=20,r=plateBoltSize/2);
			translate([-12.5,26,20]) rotate([0,90,0]) cylinder(h=20,r=plateBoltSize/2);
		}
	}

	// Finally the text. DO NOT put this in the difference as that crashes OpenSCAD
	if(includeText && includeSpar) {
		rotate([90,0,0]) translate([30,4.5,6.5]) scale([1,1,3]) drawtext("Mark II Weather Station");
		rotate([90,0,0]) translate([165,4.5,-19.75]) scale([-1,1,3]) drawtext("http://piweather.center");
	}
}

/**
 * Optional spacer block for non-legacy mounts used to attach items to the bracket
 */
module mountBlock() {
	if(!legacyMounting) {
		difference() {
			translate([sparLength-19,-7.5,10]) cube([18,25,5]);
			translate([0,-7.5,0]) union() {
					// No cutout as we have 2 bolts between the spars
					translate([sparLength-10.8,6.5,-0.1]) cylinder(h=16,r=mountBoltSize/2);
					translate([sparLength-10.8,18.5,-0.1]) cylinder(h=16,r=mountBoltSize/2);
			}
		}
	}
}
