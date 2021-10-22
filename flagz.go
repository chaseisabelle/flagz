package flagz

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Flagz []string

func (f *Flagz) String() string {
	return strings.Join(f.Array(), ", ")
}

func (f *Flagz) Set(str string) error {
	*f = append(*f, str)

	return nil
}

func (f *Flagz) Array() []string {
	return *f
}

func (f *Flagz) Stringz() []string {
	return f.Array()
}

func (f *Flagz) Intz() ([]int, error) {
	strs := f.Stringz()
	ints := make([]int, len(strs))

	for index, str := range strs {
		int, err := strconv.Atoi(str)

		if err != nil {
			return ints, err
		}

		ints[index] = int
	}

	return ints, nil
}

func (f *Flagz) Boolz() ([]bool, error) {
	strs := f.Stringz()
	boolz := make([]bool, len(strs))

	for index, str := range strs {
		b, err := strconv.ParseBool(str)

		if err != nil {
			return boolz, err
		}

		boolz[index] = b
	}

	return boolz, nil
}

func (f *Flagz) Floatz() ([]float64, error) {
	strs := f.Stringz()
	floatz := make([]float64, len(strs))

	for index, str := range strs {
		f, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return floatz, err
		}

		floatz[index] = f
	}

	return floatz, nil
}

func lookup(name string) (interface{}, error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	option := flag.Lookup(name)

	if option == nil {
		return nil, fmt.Errorf("failed to find flag %s", name)
	}

	value := option.Value

	if value == nil {
		return nil, fmt.Errorf("failed to find value for flag %s", name)
	}

	getter, ok := value.(flag.Getter)

	if !ok {
		return nil, fmt.Errorf("failed to convert value to getter for flag %s", name)
	}

	return getter.Get(), nil
}

func String(name string) (string, error) {
	val, err := lookup(name)

	if err != nil {
		return "", err
	}

	str, ok := val.(string)

	if !ok {
		err = fmt.Errorf("invalid type for flag %s", name)
	}

	return str, err
}

func Int(name string) (int, error) {
	val, err := lookup(name)

	if err != nil {
		return 0, err
	}

	i, ok := val.(int)

	if !ok {
		err = fmt.Errorf("invalid type for flag %s", name)
	}

	return i, err
}

func Bool(name string) (bool, error) {
	val, err := lookup(name)

	if err != nil {
		return false, err
	}

	b, ok := val.(bool)

	if !ok {
		err = fmt.Errorf("invalid type for flag %s", name)
	}

	return b, err
}

func Float(name string) (float64, error) {
	val, err := lookup(name)

	if err != nil {
		return 0, err
	}

	f, ok := val.(float64)

	if !ok {
		err = fmt.Errorf("invalid type for flag %s", name)
	}

	return f, err
}