package build

import (
	"fmt"
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/go-build/core"
	"github.com/peter-mount/go-build/util/arch"
	"github.com/peter-mount/go-build/util/makefile"
	"github.com/peter-mount/go-build/util/makefile/target"
	"github.com/peter-mount/go-build/util/meta"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	"github.com/peter-mount/go-script/calculator"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"github.com/peter-mount/piweather.center/astro/chart"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/io"
	io2 "io"
	"os"
	"path/filepath"
	"strings"
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

	featureDir := filepath.Join("data", "feature")
	destDirTarget := target.Target(destDir)

	_ = walk.NewPathWalker().
		Then(func(_ string, fi os.FileInfo) error {
			srcFile := fi.Name()
			src := filepath.Join(featureDir, srcFile)
			dest := filepath.Join(destDir, strings.TrimSuffix(srcFile, ".gz")+".feature")
			destDirTarget.
				Target(dest, src).
				MkDir(destDir).
				Echo("Feature", dest).
				BuildTool("-feature", src, "-d", dest)
			return nil
		}).
		PathHasSuffix(".gz").
		IsFile().
		Walk(featureDir)
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
				if str, err := calculator.GetString(n); err == nil {
					f.SetId(str)
				}
			} else if n, ok = props["ids"]; ok {
				if str, err := calculator.GetString(n); err == nil {
					f.SetId(str)
				}
			}
			if n, ok := props["name"]; ok {
				if str, err := calculator.GetString(n); err == nil {
					f.SetName(str)
				}
			}
		}

		if geom, ok := util.GetJsonObject(obj, "geometry"); ok {
			geomType := geom["type"].(string)

			if coord, ok := util.GetJsonArray(geom, "coordinates"); ok {
				// TODO MultiLineString here but MultiPolygon needed
				switch geomType {
				case "MultiLineString":
					for _, srcLine := range coord {
						if err := parseCoordLine(f, srcLine.([]interface{})); err != nil {
							return nil, err
						}
					}

				case "LineString":
					if err := parseCoordLine(f, coord); err != nil {
						return nil, err
					}

				case "Polygon":
					f.SetPolygon(true)
					if err := parseCoordLine(f, coord); err != nil {
						return nil, err
					}

				case "MultiPolygon":
					f.SetPolygon(true)
					for _, srcLine := range coord {
						for _, polyLine := range srcLine.([]interface{}) {
							if err := parseCoordLine(f, polyLine.([]interface{})); err != nil {
								return nil, err
							}
						}
					}
				default:
					return nil, fmt.Errorf("unsupported geometry type: %v", geomType)
				}
			} // coord
		} // geom

		set.AddFeature(f)
	} // features

	return set, nil
}

func parseCoordLine(f *catalogue.Feature, srcLine []interface{}) error {
	var line []chart.Point
	for _, pointDef := range srcLine {
		pointAry := pointDef.([]interface{})
		if len(pointAry) == 2 {
			x, err := calculator.GetFloat(pointAry[0])
			if err != nil {
				return err
			}
			y, err := calculator.GetFloat(pointAry[1])
			if err != nil {
				return err
			}
			line = append(line, chart.Point{X: x, Y: y})
		}
	}

	if len(line) > 0 {
		f.AddLine(line)
	}

	return nil
}
