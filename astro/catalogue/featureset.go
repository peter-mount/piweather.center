package catalogue

import (
	"encoding/binary"
	"github.com/peter-mount/piweather.center/astro/chart"
	"github.com/peter-mount/piweather.center/util"
	"github.com/soniakeys/unit"
	"io"
	"math"
)

type FeatureSet interface {
	GetLayerAll(proj chart.Projection) chart.ConfigurableLayer
	GetLayerById(proj chart.Projection, id string) chart.ConfigurableLayer
	AddFeature(*Feature)
	Write(w io.Writer) error
	Read(r io.Reader) error
}

type featureSet struct {
	features map[string]*Feature
}

func NewFeatureSet() FeatureSet {
	return &featureSet{
		features: make(map[string]*Feature),
	}
}

func (f *featureSet) AddFeature(feature *Feature) {
	f.features[feature.id] = feature
}

func (f *featureSet) GetLayerAll(proj chart.Projection) chart.ConfigurableLayer {
	p := chart.NewPath(proj)
	for _, feature := range f.features {
		feature.addToPath(p)
	}
	return p
}

func (f *featureSet) GetLayerById(proj chart.Projection, id string) chart.ConfigurableLayer {
	p := chart.NewPath(proj)
	if f, exists := f.features[id]; exists {
		f.addToPath(p)
	}
	return p
}

type Feature struct {
	id    string
	name  string
	lines [][]chart.Point
}

func (f *Feature) Id() string {
	return f.id
}

func (f *Feature) Name() string {
	return f.name
}

func (f *Feature) SetId(id string) {
	f.id = id
}

func (f *Feature) SetName(name string) {
	f.name = name
}

func (f *Feature) AddLine(l []chart.Point) {
	f.lines = append(f.lines, l)
}

func (f *Feature) addToPath(p chart.Path) {
	for _, line := range f.lines {
		p.Start()
		for _, point := range line {
			p.Add(p.Project(unit.AngleFromDeg(point.X), unit.AngleFromDeg(point.Y)))
		}
		p.End()
	}
}

func (f *featureSet) Write(w io.Writer) error {
	var b []byte
	b = binary.LittleEndian.AppendUint16(b, uint16(len(f.features)))
	if _, err := w.Write(b); err != nil {
		return err
	}

	for _, feature := range f.features {
		if err := feature.write(w); err != nil {
			return err
		}
	}

	return nil
}

func (f *Feature) write(w io.Writer) error {
	le := binary.LittleEndian

	// Form the header
	var b []byte
	b = util.AppendString(b, f.Id())
	b = util.AppendString(b, f.Name())
	b = le.AppendUint16(b, uint16(len(f.lines)))

	// write header size
	if err := writeUint16(w, uint16(len(b))); err != nil {
		return err
	}

	// Now the header content
	if _, err := w.Write(b); err != nil {
		return err
	}

	for _, line := range f.lines {
		// Number of entries in line
		b = le.AppendUint16([]byte{}, uint16(len(line)))
		for _, point := range line {
			b = le.AppendUint64(b, math.Float64bits(point.X))
			b = le.AppendUint64(b, math.Float64bits(point.Y))
		}
		if _, err := w.Write(b); err != nil {
			return err
		}
	}

	return nil
}

func writeUint16(w io.Writer, v uint16) error {
	b := binary.LittleEndian.AppendUint16([]byte{}, v)
	if n, err := w.Write(b); err != nil {
		return err
	} else if n != len(b) {
		return io.EOF
	}
	return nil
}

func readUint16(r io.Reader) (uint16, error) {
	b := make([]byte, 2)
	if n, err := r.Read(b); err != nil {
		return 0, err
	} else if n != len(b) {
		return 0, io.EOF
	}

	return binary.LittleEndian.Uint16(b), nil
}

func (f *featureSet) Read(r io.Reader) error {
	numFeatures, err := readUint16(r)
	if err != nil {
		return err
	}

	f.features = make(map[string]*Feature)
	for i := 0; i < int(numFeatures); i++ {
		feat := &Feature{}
		err = feat.read(r)
		if err != nil {
			return err
		}
		f.AddFeature(feat)
	}

	return nil
}

func (f *Feature) read(r io.Reader) error {
	le := binary.LittleEndian

	headerSize, err := readUint16(r)
	if err != nil {
		return err
	}

	b := make([]byte, headerSize)
	if n, err := r.Read(b); err != nil {
		return err
	} else if n != int(headerSize) {
		return io.EOF
	}
	f.id, b = util.ReadString(b)
	f.name, b = util.ReadString(b)
	lineCount := le.Uint16(b)

	f.lines = nil
	for i := 0; i < int(lineCount); i++ {
		var line []chart.Point
		pointCount, err := readUint16(r)
		if err != nil {
			return err
		}

		b = make([]byte, pointCount*16)
		if n, err := r.Read(b); err != nil {
			return err
		} else if n != int(pointCount*16) {
			return io.EOF
		}
		for i := 0; i < int(pointCount); i++ {
			line = append(line, chart.Point{
				X: math.Float64frombits(le.Uint64(b[0:8])),
				Y: math.Float64frombits(le.Uint64(b[8:16])),
			})
			b = b[16:]
		}
		f.lines = append(f.lines, line)
	}

	return nil
}
