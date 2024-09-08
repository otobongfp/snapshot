package fileops

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", fmt.Errorf("unable to open %w", err)
	}
	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
	}
	return buffer.String(), err
}

func ReadLIneInRange(filepath string, startLine, endLine int) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("error opeing file: %w", err)
	}
	defer file.Close()

	var buffer bytes.Buffer

	scanner := bufio.NewScanner(file)
	currentLine := 0

	if endLine == 0 {
		for scanner.Scan() {
			currentLine++
			if currentLine < startLine {
				continue
			}

			if currentLine >= startLine {
				buffer.WriteString(scanner.Text() + "\n")
			}
		}
	} else {
		for scanner.Scan() {
			currentLine++
			if currentLine < startLine {
				continue
			}

			if currentLine >= startLine && currentLine <= endLine {
				buffer.WriteString(scanner.Text() + "\n")
			}

			if currentLine > endLine {
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return buffer.String(), nil
}

func WriteTo(dirPath string, fileName string, content string) error {
	fullPath := filepath.Join(dirPath, fileName)

	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}

	fmt.Println("File written to:", fullPath)
	return nil
}
