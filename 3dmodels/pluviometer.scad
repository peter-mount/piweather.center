/*
 * openSCAD for a Pluviometer or rain gauge
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

includeTop=1;
includeFunnel=1;

includeTopPlate=1;
includeBody=1;

// **********************************************************************
// Do not modify these

// My printer can print up to 225x145x150mm so give room for raft etc
topRadius = 120/2;

containerRadius=35;
bodyHeight=50;

// **********************************************************************
if(includeTop) top();
if(includeFunnel) translate([0,0,-8]) funnel();
if(includeTopPlate) translate([0,0,-40]) topPlate();
if(includeBody) translate([0,0,-40-bodyHeight]) body();

// **********************************************************************
// The top plate
module top() {
	union() {
		// The outer ring
		difference() {
			union() {
				cylinder(3,r1=topRadius,r2=topRadius-5);
				translate([0,0,-3]) cylinder(3,r=topRadius);
			}
			translate([0,0,-4]) cylinder(10,r=topRadius-10);
		}
		// The inner mesh
		intersection() {
			union() {
				for(r=[0:2]) {
					rotate([0,0,60*r]) {
						for(i=[0:topRadius/4]) {
							translate([-topRadius+(i*9),-topRadius,-5]) cube([2,topRadius*2,5]);
						}
					}
				}
			}
			union() {
				cylinder(3,r1=topRadius,r2=topRadius-5);
				translate([0,0,-3]) cylinder(3,r=topRadius);
			}
		}
	}
}

// **********************************************************************
// The funnel which concentrates water to the center
module funnel() {
	difference() {
		union() {
			// Funnel
			cylinder(5,r=topRadius);
			translate([0,0,-30]) cylinder(30,r2=topRadius,r1=10);
			translate([0,0,-45]) cylinder(25,r=15);
			translate([0,0,-27]) cylinder(5,r=23);
			// Supports
			for(r=[0:2])
				rotate([0,0,120*r]) translate([-1,-5,-27]) cube([containerRadius-3,10,30]);
		}
		translate([0,0,-23]) cylinder(30,r2=topRadius,r1=10);
		translate([0,0,-46]) cylinder(27,r=10);

		// Support bolts
		for(r=[0:2])
			rotate([0,0,120*r]) translate([containerRadius-10,0,-30]) cylinder(8,r=2.5);

		// Remove to allow bolt from plate
		for(r=[0:2])
			rotate([0,0,60+120*r]) translate([containerRadius-10,0,-30]) cylinder(8,r=7);
	}
}

// **********************************************************************
// topPlate containing the bucket etc
module topPlate() {
	difference() {
		cylinder(5,r=containerRadius);
		translate([0,0,-1]) cylinder(7,r=16);

		// Support bolts
		for(r=[0:5])
			rotate([0,0,60*r]) translate([containerRadius-10,0,-2]) cylinder(8,r=2.85);
	}
}

module body() {
	difference() {
		cylinder(bodyHeight,r=containerRadius);
		translate([0,0,-1]) cylinder(bodyHeight+2,r=containerRadius-5);
	}
	for(r=[0:2])
		rotate([0,0,60+120*r]) translate([containerRadius-10,0,0]) union() {
			cylinder(bodyHeight,r=5);
			translate([-2,-5,0]) cube([10,10,bodyHeight]);
		}
}

