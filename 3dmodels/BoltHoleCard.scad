/*
 * openSCAD utility model to enable a modeller to determine the correct
 * hole sizes in their own models.
 *
 * This model came about when I kept wasting prints due to not getting
 * the hole sizes correctly, usually too small. You'd have thought that
 * an M4 bolt who's radius is 2mm would have a radius of 2mm, but you
 * have to adjust for the size of the fillament in the print - hence this card.
 *
 * To use, print out this card. Then when modelling bolt holes use the
 * card to determine the correct sizes for your printer (assuming each
 * printer or print settings affect the sizes).
 *
 * For each bolt size there are 4 holes represented, each one taking
 * the given radius plus 0, 0.25, 0.5 and 0.75mm respectively.
 *
 * Normally for a tight fit (you will find the bolt hard to fit & needs
 * screwing in) will be the +0.25mm hole. Use the larger sizes if you want
 * the bolt's to just slide in.
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

// The bolt's to generate
boltStart=4;

// Width of each square
w=50;

// Do not change these
boltEnd=boltStart+3;
wh=w/2;
wo=w/4;
h=wo+w*floor((2+boltEnd-boltStart)/2);
d=4;

echo(h);

// Generate the card
translate([-(wo+h)/2,-w,0]) difference() {
	cardBase();
	for(i=[boltStart:boltEnd]) {
		bolt(i);
	}
}

// Labels
include <TextGenerator.scad>

translate([-(wo+h)/2,-w,0])
translate([10,4,d])
rotate([0,0,90])
drawtext(str("M",boltStart,"-M",boltEnd," bolt hole"));

for(i=[boltStart:boltEnd]) {
	translate([-(wo+h)/2,-w,0])
	translate([wo+w*(floor((i-boltStart)/2)),w*(((i-boltStart)%2)),d])
	union() {
		translate([15,7,0])
			scale(1.5)
			rotate([0,0,90])
			drawtext(str("M",i));
		for(o=[1:3]) {
			translate([
				wh+(floor(o/2)?4:-3),
				wh+(wo*(o%2?1:-1)) +(o%2?-10:-8),
				0])
			rotate([0,0,90])
			scale(0.8)
			drawtext( str( (i/2)+(o/4) ));
		}
	}
}

module cardBase() {
	cube([h,w*2,d]);
	translate([h-1,0,d]) cube([1,w*2,2]);
	for(i=[boltStart:boltEnd]) {
		translate([wo+w*(floor((i-boltStart)/2)),w*(((i-boltStart)%2)),d])
		union() {
			cube([1,w,2]);
			cube([w,1,2]);
			translate([0,w-1,0]) cube([w,1,2]);
		}
	}
}

module bolt(i) {
	translate([wo+w*(floor((i-boltStart)/2)),w*(((i-boltStart)%2)),d])
	union() {
		for(o=[1:3]) {
			translate([
				wh+(wo*(floor(o/2)?1:-1)),
				wh+(wo*(o%2?1:-1)),
				-d-d])
			cylinder(4*d,r=(i/2)+(o/4));
		}
	}
}

module boltCenter(i) {
}
