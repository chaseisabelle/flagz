package flagz

import (
	"flag"
	"fmt"
	"strconv"
	"testing"
)

func TestStringz(t *testing.T) {
	var stringz Flagz

	flag.Var(&stringz, "strings", "strings")

	err := flag.Set("strings", "str1")

	if err != nil {
		t.Error(err)
	}

	err = flag.Set("strings", "str2")

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	strings := stringz.Stringz()

	if len(strings) != 2 {
		t.Errorf("expectected %d, got %d", 2, len(strings))
	}

	if strings[0] != "str1" {
		t.Errorf("expected %s, got %s", "str1", strings[0])
	}

	if strings[1] != "str2" {
		t.Errorf("expected %s, got %s", "str2", strings[1])
	}
}

func TestBoolz(t *testing.T) {
	var boolz Flagz

	flag.Var(&boolz, "bools", "bools")

	err := flag.Set("bools", "1")

	if err != nil {
		t.Error(err)
	}

	err = flag.Set("bools", "0")

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	bools, err := boolz.Boolz()

	if err != nil {
		t.Error(err)
	}

	if len(bools) != 2 {
		t.Errorf("expectected %d, got %d", 2, len(bools))
	}

	if !bools[0] {
		t.Errorf("expected %v, got %v", true, bools[0])
	}

	if bools[1] {
		t.Errorf("expected %v, got %v", false, bools[1])
	}
}

func TestIntz(t *testing.T) {
	var intz Flagz

	flag.Var(&intz, "ints", "ints")

	err := flag.Set("ints", "1")

	if err != nil {
		t.Error(err)
	}

	err = flag.Set("ints", "2")

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	ints, err := intz.Intz()

	if err != nil {
		t.Error(err)
	}

	if len(ints) != 2 {
		t.Errorf("expectected %d, got %d", 2, len(ints))
	}

	if ints[0] != 1 {
		t.Errorf("expected %d, got %d", 1, ints[0])
	}

	if ints[1] != 2 {
		t.Errorf("expected %d, got %d", 2, ints[1])
	}
}

func TestFloatz(t *testing.T) {
	var floatz Flagz

	flag.Var(&floatz, "floats", "floats")

	err := flag.Set("floats", "3.14")

	if err != nil {
		t.Error(err)
	}

	err = flag.Set("floats", "0.5")

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	floats, err := floatz.Floatz()

	if err != nil {
		t.Error(err)
	}

	if len(floats) != 2 {
		t.Errorf("expectected %d, got %d", 2, len(floats))
	}

	if floats[0] != 3.14 {
		t.Errorf("expected %f, got %f", 3.14, floats[0])
	}

	if floats[1] != 0.5 {
		t.Errorf("expected %f, got %f", 0.5, floats[1])
	}
}

func TestString(t *testing.T) {
	exp := "bar"
	key := "string"

	flag.String(key, "", "test")

	err := flag.Set(key, exp)

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	act, err := String(key)

	if err != nil {
		t.Error(err)
	}

	if act != exp {
		t.Errorf("expected %s, got %s", exp, act)
	}
}

func TestBool(t *testing.T) {
	key := "bool"
	exp := true

	flag.Bool(key, false, "test")

	err := flag.Set(key, "1")

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	act, err := Bool(key)

	if err != nil {
		t.Error(err)
	}

	if act != exp {
		t.Errorf("expected %v, got %v", exp, act)
	}
}

func TestInt(t *testing.T) {
	key := "int"
	exp := 5

	flag.Int(key, 0, "test")

	err := flag.Set(key, strconv.Itoa(exp))

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	act, err := Int(key)

	if err != nil {
		t.Error(err)
	}

	if act != exp {
		t.Errorf("expected %v, got %v", exp, act)
	}
}

func TestFloat(t *testing.T) {
	key := "float"
	exp := 1.23

	flag.Float64(key, 0, "test")

	err := flag.Set(key, fmt.Sprintf("%f", exp))

	if err != nil {
		t.Error(err)
	}

	flag.Parse()

	act, err := Float(key)

	if err != nil {
		t.Error(err)
	}

	if act != exp {
		t.Errorf("expected %v, got %v", exp, act)
	}
}
