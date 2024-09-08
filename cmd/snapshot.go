package snapshot

import (
	"fmt"
	"os"

	"github.com/otobongfp/snapshot/internal/fileops"
	"github.com/otobongfp/snapshot/internal/metadata"
)

func GenerateMeta() {
	fmt.Println("Snaphot process started....")
	// Call GetDetails() from the metadata pkg
	meta := metadata.GetDetails()
	// Print out the metadata in a well-formatted manner
	fmt.Println("Snapshot Metadata:")
	fmt.Printf("Filename: %s\n", meta.Filename)
	fmt.Printf("Title: %s\n", meta.Description)
	fmt.Printf("Hash: %s\n", meta.Hash)
	fmt.Printf("Operating System: %s\n", meta.Os)
	fmt.Printf("Created At: %s\n", meta.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("Snapshot process completed.")
}

func Capture() {
	fmt.Println("Snapshot Capture Wizard")
	fmt.Println("(Filepath must start from current dir)")

	currentPath, _ := os.Getwd()
	path := ""

	fmt.Println("Path to file: ")
	fmt.Scan(&path)
	filepath := currentPath + "/" + path
	fmt.Println(filepath)

	fmt.Println("Capture the entire file [yes=1, no=2] default(1):")
	choice := 1
	fmt.Scan(&choice)

	dirPath, _ := os.Getwd()

	if choice == 1 {
		content, _ := fileops.ReadFile(filepath)
		fileops.WriteTo(dirPath, "text.txt", content)
	} else if choice == 2 {
		var startLine int
		var endLine int
		fmt.Println("Start from line?: ")
		fmt.Scan(&startLine)

		fmt.Println("End at line? (0 for end of Line): ")
		fmt.Scan(&endLine)

		if startLine == 0 {
			fmt.Println("Cannot start from line 0")
			return
		}

		content, _ := fileops.ReadLIneInRange(filepath, startLine, endLine)
		fileops.WriteTo(dirPath, "text.txt", content)

	}

}
