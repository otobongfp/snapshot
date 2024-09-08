package main

import (
	"fmt"

	snapshot "github.com/otobongfp/snapshot/cmd"
	// "path/filepath"
	// "runtime"
	// "strconv"
	// "time"
)

func main() {
	fmt.Println("Welcome to the Snapshot CLI tool!")
	snapshot.Capture()
	//snapshot.GenerateMeta()
}
