package main

import (
	// "fmt"
	// "time"
	"os"
	//"encoding/json"
	"path/filepath"
)

func main() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir component: %v", filepath.Dir(path))
	Printfln("File component: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))
}
