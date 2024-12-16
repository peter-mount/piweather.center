package catalogue

import (
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/piweather.center/util/io"
	"sync"
)

type Manager struct {
	mutex    sync.Mutex
	ybsc     *Catalog
	features map[string]FeatureSet
}

func (m *Manager) YaleBrightStarCatalog() (*Catalog, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.ybsc == nil {
		cat := &Catalog{}
		if err := io.NewReader(cat.Read).
			Decompress().
			Open(application.FileName(application.STATIC, "bsc5.bin")); err != nil {
			return nil, err
		}
		m.ybsc = cat
	}
	return m.ybsc, nil
}

func (m *Manager) Feature(name string) (FeatureSet, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.features == nil {
		m.features = make(map[string]FeatureSet)
	}

	if feature, ok := m.features[name]; ok {
		return feature, nil
	}

	feature := NewFeatureSet()
	if err := io.NewReader(feature.Read).
		Decompress().
		Open(application.FileName(application.STATIC, "feature", name+".feature")); err != nil {
		return nil, err
	}

	m.features[name] = feature
	return feature, nil
}
