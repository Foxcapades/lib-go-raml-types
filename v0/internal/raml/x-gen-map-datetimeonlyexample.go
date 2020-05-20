package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func NewDatetimeOnlyExampleMap(log *logrus.Entry) *DatetimeOnlyExampleMap {
	return &DatetimeOnlyExampleMap{
		log:   xlog.WithType(log, "internal.DatetimeOnlyExampleMap"),
		index: make(map[string]*raml.DatetimeOnlyExample),
	}
}

// DatetimeOnlyExampleMap generated @ 2020-05-20T01:05:35.571783841-04:00
type DatetimeOnlyExampleMap struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	index map[string]*raml.DatetimeOnlyExample
}

func (o *DatetimeOnlyExampleMap) Len() uint {
	return uint(len(o.slice))
}

func (o *DatetimeOnlyExampleMap) Put(key string, value raml.DatetimeOnlyExample) raml.DatetimeOnlyExampleMap {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

func (o *DatetimeOnlyExampleMap) PutNonNil(key string, value raml.DatetimeOnlyExample) raml.DatetimeOnlyExampleMap {
	if !util.IsNil(value) {
		return o.Put(key, value)
	}
	return o
}

func (o *DatetimeOnlyExampleMap) Replace(key string, value raml.DatetimeOnlyExample) raml.DatetimeOnlyExample {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *DatetimeOnlyExampleMap) ReplaceOrPut(key string, value raml.DatetimeOnlyExample) raml.DatetimeOnlyExample {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
		return nil
	}

	out := *o.index[key]
	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *DatetimeOnlyExampleMap) Get(key string) raml.DatetimeOnlyExample {
	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *DatetimeOnlyExampleMap) At(index uint) (key option.String, value raml.DatetimeOnlyExample) {
	tmp := &o.slice[index]
	key = option.NewString(tmp.Key.(string))
	value = tmp.Value.(raml.DatetimeOnlyExample)

	return
}

func (o *DatetimeOnlyExampleMap) IndexOf(key string) option.Uint {
	if !o.Has(key) {
		return option.NewEmptyUint()
	}
	for i := range o.slice {
		if o.slice[i].Key == key {
			return option.NewUint(uint(i))
		}
	}
	panic("invalid map state, index out of sync")
}

func (o *DatetimeOnlyExampleMap) Has(key string) bool {
	_, ok := o.index[key]
	return ok
}

func (o *DatetimeOnlyExampleMap) Delete(key string) raml.DatetimeOnlyExample {
	if !o.Has(key) {
		return nil
	}

	out := *o.index[key]
	delete(o.index, key)

	for i := range o.slice {
		if o.slice[i].Key == key {
			o.slice = append(o.slice[:i], o.slice[i+1:]...)
			return out
		}
	}
	panic("invalid map state, index out of sync")
}

func (o DatetimeOnlyExampleMap) ForEach(fn func(string, raml.DatetimeOnlyExample)) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o DatetimeOnlyExampleMap) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *DatetimeOnlyExampleMap) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.DatetimeOnlyExampleMap.UnmarshalRAML")
	yml, err := assign.AsMapSlice(val)

	if err != nil {
		return xlog.Error(log, err)
	}

	for i := range yml {
		tmp := &yml[i]
		l2 := xlog.AddPath(log, tmp.Key)

		key := ""

		if err = assign.AsString(tmp.Key, &key, l2); err != nil {
			return xlog.Error(l2, err)
		}

		tmpVal := NewDatetimeOnlyExample(l2)
		if err = tmpVal.UnmarshalRAML(tmp.Value, l2); err != nil {
			return xlog.Error(l2, err)
		}

		o.Put(key, tmpVal)
	}

	return nil
}
