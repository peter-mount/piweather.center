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

/*
 * Measured dimensions of a PI Camera Module.
 *
 * Width & Depth are the exterior dimensions. 0.25 is to give some manovering room
 * y offset is the height from the bottom of the center of the camera lens
 * ribbon width is just that, if required
 */
piCameraWidth=25.2;
piCameraDepth=24.2;
piCameraYOffset=9.4;
piCameraRibbonWidth=16.5;

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
 */

module piCamera(width,mode) {
	assign(
		tubeRadius=width/2,
		offset=piCameraYOffset
	) {
		translate([0,0,piCameraYOffset/2]) union() {
			// Front panel
			difference() {
				if(mode)
					translate([6,-tubeRadius,-tubeRadius-offset]) cube([6,width,width+offset]);
				else
					translate([0,-tubeRadius,-tubeRadius-offset]) cube([6,width,width+offset]);

				// Holder for the pi camera module
				translate([5,-piCameraWidth/2,-piCameraWidth/2])
					cube([5,piCameraWidth,piCameraWidth]);

				// Hole for the camera lens
				translate([-0.5,-2.5,-piCameraYOffset+2.5]) cube([6,5,5]);
		
				// Space for ribbon cable
				translate([5,-piCameraRibbonWidth/2,-tubeRadius-1-offset])
					cube([5,piCameraRibbonWidth,tubeRadius+1+offset]);

				// M4 mounting holes
				for(h=[0:4])
						translate([-1,
							(h%2) ? (-tubeRadius+5) : (tubeRadius-5),
							(floor(h/2)) ? (-tubeRadius+5-offset) : (tubeRadius-5)
						])
					rotate([0,90,0])
					cylinder(h=16,r=2);
			}
		}

	}
}

/*
 * Standalone, generate a model of just the camera module with the specified radius
 */
module piCameraStandalone(radius) {
	translate([6,0,0]) rotate([0,0,180]) piCamera(radius,0);
	translate([-6,-radius*3/2,0]) piCamera(radius,1);
}
