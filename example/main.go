package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/flagz"
)

func main() {
	var stringz flagz.Flagz
	var boolz flagz.Flagz
	var intz flagz.Flagz
	var floatz flagz.Flagz

	flag.Var(&stringz, "strings", "strings")
	flag.Var(&boolz, "bools", "bools")
	flag.Var(&intz, "ints", "ints")
	flag.Var(&floatz, "floats", "floats")

	flag.String("string", "default", "string flag")
	flag.Bool("bool", false, "bool flag")
	flag.Int("int", 0, "int flag")
	flag.Float64("float", 0, "float flag")

	flag.Parse()

	strings := stringz.Stringz()
	bools, boolErr := boolz.Boolz()
	ints, intErr := intz.Intz()
	floats, floatErr := floatz.Floatz()

	fmt.Println("-------------- multi-value flag getters")
	fmt.Printf("%+v\n", strings)
	fmt.Printf("%+v\n%+v\n", bools, boolErr)
	fmt.Printf("%+v\n%+v\n", ints, intErr)
	fmt.Printf("%+v\n%+v\n", floats, floatErr)

	s, stringErr := flagz.String("string")
	b, boolErr := flagz.Bool("bool")
	i, intErr := flagz.Int("int")
	f, floatErr := flagz.Float("float")

	fmt.Println("-------------- post-parse flag getters")
	fmt.Printf("%+v\n%+v\n", s, stringErr)
	fmt.Printf("%+v\n%+v\n", b, boolErr)
	fmt.Printf("%+v\n%+v\n", i, intErr)
	fmt.Printf("%+v\n%+v\n", f, floatErr)
}
