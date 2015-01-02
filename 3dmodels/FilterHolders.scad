/*
 * openSCAD for a set of Filter holders
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
 * Module to render a simple Filter Holder
 *
 * To use translate it into the required position so that the filter is centered
 * at the origin & pointing along the x axis.
 *
 * Then call filterSimpleHolder(width,offset,radius)
 *
 * Where:
 *		width		the width of the component
 *		offset 	adjust on the z axis by this amount, usually 0
 *		radius		The radius of the opening of the filter
 *
 */
module filterSimpleHolder(width,offset,radius) {
	assign(half=width/2,oh=offset/2) {
		translate([0,0,oh]) union() {
			difference() {
				translate([0,-half,-half-offset]) cube([5,width,width+offset]);

				// the opening
				translate([-1,0,-oh]) rotate([0,90,0]) cylinder(h=7,r=radius);

				// M4 mounting holes
				for(h=[0:4])
						translate([-1,
							(h%2) ? -half+5 : half-5,
							(floor(h/2)) ? (-half+5-offset) : (half-5)
						])
					rotate([0,90,0])
					cylinder(h=16,r=2);
			}
		}
	}
}
