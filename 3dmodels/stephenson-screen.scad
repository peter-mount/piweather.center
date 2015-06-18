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

outdiam=120;
height=20;

// Thickness of components
thick=5;

// The bolt size, for M6 use 3
boltSize=3;

/**********************************************************************
 * DO NOT CHANGE ANYTHING BELOW THIS POINT
 **********************************************************************/

boltMountHeight=height*0.5;

mountBoltSize=4;

// 2*thick
thick2=thick*2;

indiam=outdiam/2;

outrad=outdiam/2;
inrad=indiam/2;

// Uncomment one of these to generate each model
// Remember to comment out the exploded view first
//base();
//slat();
//plainTop();

/* Example of showing all in exploded format.
*/
translate([0,0,(1*6)+(height-4)*6]) plainTop();

for(s=[0:5]) translate([0,0,(1*s)+(height-5)*s]) slat();
translate([0,0,0]) slat();
translate([0,0,-thick-1]) base();


/*
 * A plain top, simply ensure rail runs off.
 *
 * This could be used as the basis for alternate tops, i.e. solar panel, light sensor etc
 */
module plainTop() {
	difference() {
		cylinder(h=height,r1=inrad+11,r2=0);
		translate([0,0,-1]) cylinder(h=height-5,r1=inrad+3,r2=0);
	}
	for(b=[0:3])
		rotate([0,0,90*b])
		translate([inrad+thick,0,-3])
		cylinder(h=4,r=boltSize);
}

/*
 * Create a slat.
 *
 * isTop=1 for when this is the top slat, 0 for all others below
 */
module slat() {
	// The cone
	difference() {
		cylinder(h=height,r1=outrad,r2=inrad);
		translate([0,0,-1]) union() {
			cylinder(h=height+(thick/2),r1=outrad-thick,r2=inrad-thick);
			slatMount();
		}
	}

	boltmounts();
}

module slatMount() {
	// The bolt holes, needed to go through the outer casing
	for(b=[0:3])
		rotate([0,0,90*b])
		translate([inrad+thick,0,height-boltMountHeight])
		cylinder(h=boltMountHeight+2,r=boltSize+0.25);
}

module boltmounts() {
	for(b=[0:3])
		assign(
			bh = b%2==0 ? (boltMountHeight*2)-5 : boltMountHeight,
			ofs = b%2==0 ? 5 : (height-boltMountHeight)
		) {
			rotate([0,0,90*b])
			translate([inrad+thick,0,ofs]) {
				difference() {
					cylinder(h=bh,r=boltSize+3);
					translate([0,0,-1])
						cylinder(h=bh+2,r=boltSize+0.2);
				}
				if(b%2==0)
					translate([0,0,-3])
					cylinder(h=6,r=boltSize);
			}
		}
}

/*
 * Base mounting plate
 */
module base() {
	difference() {
		union() {
			cylinder(h=thick,r=outrad);

			// Cable access
			translate([0,20,0]) cylinder(h=10,r=7);
		}

		// Cable access
		translate([0,20,-1]) cylinder(h=16,r=5);

		// Peg holes spaced 12mm apart and 3mm deep to allow for equipment to be
		// mounted on the base.
		//
		// Note:
		//		We don't put pegs near the shade mountings nor the cable access.
		//		The center two are also through holes for mounting to the mast bracket
		//
		rotate([0,0,90])
		for(y=[-3:3])
			translate([12*y,-12,0])
				union() {
					for(h=[-2:3])
						if( (y==0 && h>=-1 && h<=2) || (y!=0 && (y<0 || y>2 || h<0||h>1) ) )
							assign( th =  (y==0 && (h==0 || h==1 )) ? 6:3) {
								translate([0,6.5+(12*h),thick-th])
								cylinder(h=16,r=mountBoltSize/2);
							}
				}
	}

	for(b=[0:3])
		if(b%2==0) {
			rotate([0,0,90*b])
			translate([inrad+thick,0,thick])
			difference() {
				cylinder(h=boltMountHeight/2,r=boltSize+3);
				cylinder(h=boltMountHeight+2,r=boltSize+0.2);
			}
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
