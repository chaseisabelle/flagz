# flagz
simple go package for multiple flags with the same name/key, and getting flags after `flag.Parse()` has been called.

---
## example

The code...
```go
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
```

Running it...
```shell
./main --strings=str1 --strings=str2 --string=str --ints=1 --ints=2 --int=3 --bools=1 --bools=0 --bool=1 --floats=1.23 --floats=4.56 --float=0.5
-------------- multi-value flag getters
[str1 str2]
[true false]
<nil>
[1 2]
<nil>
[1.23 4.56]
<nil>
-------------- post-parse flag getters
str
<nil>
true
<nil>
3
<nil>
0.5
<nil>
```