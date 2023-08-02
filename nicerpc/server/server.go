package main

// total length = uint16
// function length = uint8
// function name ( bytes )
// first parameter length = uint8
// first parameter type = uint8 (slice, map, etc.. )
// first parameter value type = uint8 (int, float, etc )
// first parameter ( bytes )

var funcMap = make(map[string]func(interface{}))

func main() {

	go func() {
		funcName := "meow"
		interfaceFromBytes := "meow"
		funcMap[funcName](interfaceFromBytes)
	}()

}
