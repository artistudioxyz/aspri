package main

import (
	"aspri/library"
	"flag"
	"fmt"
)

var (
	offsetFlag = flag.String("offset", "", "file and byte offset of identifier to be renamed, e.g. 'file.go:#123'.  For use by editors.")
	fromFlag   = flag.String("from", "", "identifier to be renamed; see -help for formats")
	toFlag     = flag.String("to", "", "new name for identifier")
	helpFlag   = flag.Bool("help", false, "show usage message")
)

func main() {
	fmt.Println("hello world")

	library.RefactorPlugins("/Users/muhammadsundoro/Desktop/Iseng", "MANTAP", "ISENG")
	//library.RefactorPlugins("/Users/muhammadsundoro/Desktop/Iseng", "ISENG", "MANTAP")
}
