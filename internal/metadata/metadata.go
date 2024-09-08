package metadata

import (
	"fmt"
	"runtime"
	"time"
)

type meta struct {
	Filename    string
	Description string
	Hash        string
	Os          string
	CreatedAt   time.Time
}

func newMetaData(filename, des string) meta {
	return meta{
		Filename:    filename,
		Description: des,
		Hash:        "",
		Os:          runtime.GOOS,
		CreatedAt:   time.Now(),
	}
}

func GetDetails() meta {
	fmt.Println("What do you want to name the file (useful for search): ")
	var f string
	fmt.Scan(&f)

	fmt.Println("Describe what/why/info about this file (useful for search): ")
	var d string
	fmt.Scan(&d)

	return newMetaData(f, d)
}
