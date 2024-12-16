package coord

import (
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/unit"
	"math"
)

// CoordinateTransformer provides a common object to perform coordinate transforms.
//
// This is based on the various transforms in the "github.com/soniakeys/meeus/v3/coord" package
// but is optimised so that common calculations are performed once so bulk transforms are
// slightly faster than the originals.
type CoordinateTransformer interface {
	// Sidereal sets the sidereal time at Greenwich for transforms.
	//
	// Sidereal time must be consistent with the equatorial coordinates.
	// If coordinates are apparent, sidereal time must be apparent as well.
	Sidereal(st unit.Time) CoordinateTransformer

	// Obliquity sets the obliquity of the Ecliptic, required for transforms involving Ecliptic coordinates
	Obliquity(ε unit.Angle) CoordinateTransformer

	// EclToEq converts ecliptic coordinates to equatorial coordinates.
	//
	//	λ: ecliptic longitude coordinate to transform
	//	β: ecliptic latitude coordinate to transform
	//
	// Results:
	//	α: right ascension
	//	δ: declination
	EclToEq(λ, β unit.Angle) (α unit.RA, δ unit.Angle)

	// EqToEcl converts equatorial coordinates to ecliptic coordinates.
	//
	//	α: right ascension coordinate to transform
	//	δ: declination coordinate to transform
	//
	// Results:
	//
	//	λ: ecliptic longitude
	//	β: ecliptic latitude
	EqToEcl(α unit.RA, δ unit.Angle) (λ, β unit.Angle)

	// EqToHz converts Equatorial to Horizontal coordinates
	//
	//	α: right ascension coordinate to transform
	//	δ: declination coordinate to transform
	//
	// Results:
	//
	//	A: azimuth of observed point, measured westward from the South.
	//	h: elevation, or height of observed point above horizon.
	EqToHz(α unit.RA, δ unit.Angle) (A, h unit.Angle)

	// HzToEq transforms horizontal coordinates to equatorial coordinates.
	//
	//	A: azimuth
	//	h: elevation
	//
	// Results:
	//
	//	α: right ascension
	//	δ: declination
	HzToEq(A, h unit.Angle) (α unit.RA, δ unit.Angle)

	// EqToHzMultiple converts multiple Equatorial coordinates into Horizontal coordinates
	EqToHzMultiple(src []coord.Equatorial) []coord.Horizontal

	// EqToHzAboveHorizon is the same as EqToHzMultiple but the output only includes
	// points that are above the horizon.
	EqToHzAboveHorizon(src []coord.Equatorial) []coord.Horizontal
}

type coordinateTransformer struct {
	φ      unit.Angle // latitude of observer on Earth
	ψ      unit.Angle // longitude of observer on Earth
	sφ, cφ float64    // Sin & Cos of φ
	st     unit.Time  // sidereal time at Greenwich at time of observation
	H0     float64    // st.Rad() - ψ.Rad()
	ε      unit.Angle // Obliquity of the ecliptic
	sε, cε float64    // Sin & Cos of ε
}

// NewCoordinateTransformer creates a new CoordinateTransformer
//
//	φ: latitude of observer on Earth
//	ψ: longitude of observer on Earth

func NewCoordinateTransformer(φ, ψ unit.Angle) CoordinateTransformer {
	tr := &coordinateTransformer{φ: ψ, ψ: φ}

	tr.sφ, tr.cφ = math.Sincos(tr.φ.Rad())

	return tr
}

func (tr *coordinateTransformer) Sidereal(st unit.Time) CoordinateTransformer {
	tr.st = st
	tr.H0 = st.Rad() - tr.ψ.Rad()
	return tr
}

func (tr *coordinateTransformer) EqToHz(α unit.RA, δ unit.Angle) (A, h unit.Angle) {
	H := tr.H0 - α.Rad()
	sH, cH := math.Sincos(H)
	sδ, cδ := δ.Sincos()
	A = unit.Angle(math.Atan2(sH, cH*tr.sφ-(sδ/cδ)*tr.cφ)) // (13.5) p. 93
	h = unit.Angle(math.Asin(tr.sφ*sδ + tr.cφ*cδ*cH))      // (13.6) p. 93
	return A, h
}

func (tr *coordinateTransformer) HzToEq(A, h unit.Angle) (α unit.RA, δ unit.Angle) {
	sA, cA := A.Sincos()
	sh, ch := h.Sincos()
	H := math.Atan2(sA, cA*tr.sφ+sh/ch*tr.cφ)
	α = unit.RAFromRad(tr.H0 - H)
	δ = unit.Angle(math.Asin(tr.sφ*sh - tr.cφ*ch*cA))
	return α, δ
}

func (tr *coordinateTransformer) EqToHzMultiple(src []coord.Equatorial) []coord.Horizontal {
	return tr.eqToHzSlice(src, true)
}

func (tr *coordinateTransformer) EqToHzAboveHorizon(src []coord.Equatorial) []coord.Horizontal {
	return tr.eqToHzSlice(src, false)
}

func (tr *coordinateTransformer) eqToHzSlice(src []coord.Equatorial, all bool) []coord.Horizontal {
	var r []coord.Horizontal

	for _, eq := range src {
		A, h := tr.EqToHz(eq.RA, eq.Dec)
		if all || h >= 0.0 {
			r = append(r, coord.Horizontal{Az: A, Alt: h})
		}
	}

	return r
}

func (tr *coordinateTransformer) Obliquity(ε unit.Angle) CoordinateTransformer {
	tr.ε = ε
	tr.sε, tr.cε = ε.Sincos()
	return tr
}

func (tr *coordinateTransformer) EqToEcl(α unit.RA, δ unit.Angle) (λ, β unit.Angle) {
	sα, cα := α.Sincos()
	sδ, cδ := δ.Sincos()
	λ = unit.Angle(math.Atan2(sα*tr.cε+(sδ/cδ)*tr.sε, cα)) // (13.1) p. 93
	β = unit.Angle(math.Asin(sδ*tr.cε - cδ*tr.sε*sα))      // (13.2) p. 93
	return λ, β
}

func (tr *coordinateTransformer) EclToEq(λ, β unit.Angle) (α unit.RA, δ unit.Angle) {
	sλ, cλ := λ.Sincos()
	sβ, cβ := β.Sincos()
	α = unit.RAFromRad(math.Atan2(sλ*tr.cε-(sβ/cβ)*tr.sε, cλ)) // (13.3) p. 93
	δ = unit.Angle(math.Asin(sβ*tr.cε + cβ*tr.sε*sλ))          // (13.4) p. 93
	return α, δ
}
