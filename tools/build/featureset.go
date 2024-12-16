package build

import (
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/go-build/core"
	"github.com/peter-mount/go-build/util/arch"
	"github.com/peter-mount/go-build/util/makefile"
	"github.com/peter-mount/go-build/util/makefile/target"
	"github.com/peter-mount/go-build/util/meta"
	"github.com/peter-mount/go-script/calculator"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"github.com/peter-mount/piweather.center/astro/chart"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/io"
	io2 "io"
	"path/filepath"
)

// FeatureSet compiles the standard FeatureSet's for star charts.
//
// Many of these are geojson files from https://github.com/dieghernan/celestial_data
type FeatureSet struct {
	Encoder *core.Encoder `kernel:"inject"`
	Build   *core.Build   `kernel:"inject"`
	Feature *string       `kernel:"flag,feature,Encode feature"`
}

func (s *FeatureSet) Start() error {
	s.Build.AddExtension(s.extension)
	s.Build.Makefile(0, s.makefile)
	if *s.Feature != "" {
		return s.encode()
	}
	return nil
}

func (s *FeatureSet) destDir(arch arch.Arch) string {
	return filepath.Join(arch.BaseDir(*s.Encoder.Dest), application.FileName(application.STATIC, "feature"))
}

func (s *FeatureSet) extension(arch arch.Arch, target target.Builder, meta *meta.Meta) {
	destDir := s.destDir(arch)

	destDirTarget := target.Target(destDir)

	for _, srcFile := range []string{"const.border", "const.line"} {
		src := filepath.Join("data", srcFile+".gz")
		dest := filepath.Join(destDir, srcFile+".feature")
		destDirTarget.
			Target(dest, src).
			MkDir(destDir).
			Echo("Feature", dest).
			BuildTool("-feature", src, "-d", dest)
	}
}

func (s *FeatureSet) makefile(root makefile.Builder, _ target.Builder, _ *meta.Meta) {
	// This adds a dependency to test for vsop87.
	//
	// With gnu-make, this works because it merges this test rule with the normal one
	// so when test is run, vsop87 is performed first so that the tests have access
	// to that data
	root.Rule("test", s.destDir(s.Build.BuildArch()))
}

func (s *FeatureSet) encode() error {
	var m map[string]interface{}

	if err := io.NewReader().
		Json(&m).
		Decompress().
		Open(*s.Feature); err != nil {
		return err
	}

	if set, err := s.importGeoJson(m); err != nil {
		return err
	} else {
		err = io.NewWriter(set.Write).
			Compress().
			CreateFile(*s.Encoder.Dest)
	}

	return nil
}

func (s *FeatureSet) importGeoJson(o map[string]interface{}) (catalogue.FeatureSet, error) {
	set := catalogue.NewFeatureSet()

	features, ok := util.GetJsonArray(o, "features")
	if !ok {
		return nil, io2.EOF
	}

	for _, f := range features {
		obj := f.(map[string]interface{})
		f := &catalogue.Feature{}
		if props, ok := util.GetJsonObject(obj, "properties"); ok {
			if n, ok := props["id"]; ok {
				f.SetId(n.(string))
			} else if n, ok = props["ids"]; ok {
				f.SetId(n.(string))
			}
			if n, ok := props["name"]; ok {
				f.SetName(n.(string))
			}
		}

		if geom, ok := util.GetJsonObject(obj, "geometry"); ok {
			// TODO MultiLineString here but MultiPolygon needed
			if coord, ok := util.GetJsonArray(geom, "coordinates"); ok {
				for _, srcLine := range coord {
					var line []chart.Point
					for _, pointDef := range srcLine.([]interface{}) {
						pointAry := pointDef.([]interface{})
						if len(pointAry) == 2 {
							x, err := calculator.GetFloat(pointAry[0])
							if err != nil {
								return nil, err
							}
							y, err := calculator.GetFloat(pointAry[1])
							if err != nil {
								return nil, err
							}
							line = append(line, chart.Point{X: x, Y: y})
						}
					}
					if len(line) > 0 {
						f.AddLine(line)
					}
				}
			} // coord
		} // geom

		set.AddFeature(f)
	} // features

	return set, nil
}
