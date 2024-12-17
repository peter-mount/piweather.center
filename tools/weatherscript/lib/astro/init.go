package astro

import "github.com/peter-mount/go-script/packages"

func init() {
	packages.Register("angle", &Angle{})
	packages.Register("astro", &Astro{})
	packages.Register("astroTime", &Time{})
	packages.Register("geo", &Geo{})
}
