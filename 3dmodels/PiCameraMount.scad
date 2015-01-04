/*
 * openSCAD for a simple mount for the Raspberry PI Camera Module
 *
 * This file is not for standalone use - it's included in to another
 * openSCAN file for including in to another model.
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

// Uncomment this line to render/print just the holder as-is.
// Leave it commented out so that this file can be included in other modules
//piCameraStandalone(50);

/*
 * Measured dimensions of a PI Camera Module.
 *
 * These are based on actual measurements and those manually measured by Gert van Loo
 * and published on Scribd:
 *
 *		http://www.scribd.com/doc/142718448/Raspberry-Pi-Camera-Mechanical-Data
 *
 * Width & Depth are the exterior dimensions, width being the edge with the ribbon.
 *
 * We add 1.5 to these to give some room and account for the printer as it seems
 * to give a better fit.
 *
 * y offset is the height from the bottom of the center of the camera lens.
 *
 * ribbon width is just that, if required
 */
piCameraWidth=25 +1.5;
piCameraDepth=23.9 +1.5;
piCameraRibbonWidth=16.2 +0.5;
piCameraLensSize=8 +1.5;
piCameraYOffset=5.1+(piCameraLensSize/2);

/*
 * Module to render a PI camera mount.
 *
 * To use translate it into the required position so that the camera lens
 * is at the origin & pointing along the x axis.
 *
 * Then call piCameraMount(radius,mode)
 *
 * Where:
 *		width	the width of the component
 *		mode	The mode of the render
 *
 * The width must be >30mm as it's used to include the mounting hardware.
 * However you can enlarge it.
 *
 * The mode defines how it's rendered:
 *		0	Render the front panel
 *		1	Render the back panel
 *		2	Render a plain panel with just the mounting holes
 */

module piCamera(width,mode) {
	assign(
		tubeRadius=width/2,
		piCameraLensHalf=piCameraLensSize/2,
		piCameraWHalf=piCameraWidth/2,
		piCameraDHalf=piCameraDepth/2,
		offset=piCameraYOffset
	) {
		translate([0,0,piCameraYOffset/2]) union() {
			// Front panel
			difference() {
				if(mode==1)
					translate([6,-tubeRadius,-tubeRadius-offset]) cube([6,width,width+offset]);
				else
					translate([0,-tubeRadius,-tubeRadius-offset]) cube([6,width,width+offset]);

				if(mode!=2) {
					// Holder for the pi camera module
					translate([5,-piCameraWHalf,-piCameraDHalf])
						cube([1.1,piCameraWidth,piCameraDepth]);

					// Holder for components on rear of module
					translate([6,1-piCameraWHalf,1-piCameraDHalf])
						cube([3,piCameraWidth-2,piCameraDepth-2]);
					translate([6,-piCameraWHalf,-piCameraDHalf])
						cube([3,piCameraWidth,9.5]);

					// Hole for the camera lens
					translate([-0.5,-piCameraLensHalf,2-piCameraLensSize])
						cube([6,piCameraLensSize,piCameraLensSize]);

					// Indent for ribbon holder on front of the module
					translate([2.5,-piCameraLensSize*2/3,-0.5])
						cube([3,piCameraLensSize*4/3,piCameraDHalf+1]);
		
					// Space for ribbon cable
					translate([5,-piCameraRibbonWidth/2,-tubeRadius-1-offset])
						cube([3,piCameraRibbonWidth,tubeRadius+1+offset]);
				}

				// M4 mounting holes
				piCameraHoles(width,15);
			}
		}

	}
}

module piCameraHoles(width,h) {
	assign(tubeRadius=width/2) {
		for(h=[0:3])
			translate([-1,
				(h%2) ? (-tubeRadius+5) : (tubeRadius-5),
				(floor(h/2)) ? (-tubeRadius+5-piCameraYOffset) : (tubeRadius-5)
			])
			rotate([0,90,0])
			cylinder(h=15,r=2);
	}
}

/*
 * Standalone, generate a model of just the camera module with the specified radius
 */
module piCameraStandalone(radius) {
	translate([6,0,0]) rotate([0,0,180]) piCamera(radius,0);
	translate([-6,-radius*3/2,0]) piCamera(radius,1);
}
