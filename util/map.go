package util

type StringMap map[string]string

func NewStringMap() StringMap {
	return make(map[string]string)
}

func (a *StringMap) Add(k, v string) bool {
	if a != nil {
		if _, exists := (*a)[k]; !exists {
			(*a)[k] = v
			return true
		}
	}
	return false
}

func (a *StringMap) Get(k string) string {
	if a == nil {
		return ""
	}
	return (*a)[k]
}

func (a *StringMap) Contains(k string) bool {
	if a == nil {
		return false
	}
	_, exists := (*a)[k]
	return exists
}

func (a *StringMap) Size() int {
	if a == nil {
		return 0
	}
	return len(*a)
}

func (a *StringMap) IsEmpty() bool {
	return a == nil || len(*a) == 0
}

func (a *StringMap) Keys() []string {
	var r []string
	for k, _ := range *a {
		r = append(r, k)
	}
	return r
}

func (a *StringMap) ForEach(f func(string, string) error) error {
	if a != nil {
		for k, v := range *a {
			if err := f(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
