package main

import (
	"fmt"
	"path"
	"path/filepath"
	"os"
	"aqwari.net/xml/xsdgen"
	"strings"
)

func main(){

	var files []string
	var err error

	files, err = filepath.Glob("xsd_files/all_xsd_files/*.xsd")
	if err != nil {
		fmt.Println("Could not get details of the files - Error:  %v", err.Error())
		os.Exit(1)
	}

	for _, f := range files {
		fmt.Printf("Finished: %v \n", f)
		generateFile(f)
	}
}

func generateFile(f string){

	var cfg xsdgen.Config

	// Windows:
	fn := strings.Replace(f, "\\", "/", -1)

	// File and folder name:
	_, fn = path.Split(fn)
	fn = strings.Replace(fn, ".", "_", -1)
	fn = strings.TrimRight(fn, "_xsd")
	folder := fmt.Sprintf("generated_files/%s", fn)
	pt := fmt.Sprintf("%s/%s.go", folder, fn)

	// Make directory and generate file:
	_ = os.Mkdir(folder, 0666)

	cfg.Option(
		xsdgen.IgnoreAttributes("id", "href", "offset"),
		xsdgen.IgnoreElements("comment"),
		xsdgen.PackageName(fmt.Sprintf("iso20022_%s", fn)),
		xsdgen.Replace("_", ""),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
	)

	if err := cfg.GenCLI("-o", pt, f); err != nil {
		fmt.Printf("%v", err)
	}
}


