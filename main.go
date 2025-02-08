package main

import (
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	// Get list of files from images directory
	files, err := os.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	pngFiles := make([]string, 0)
	for _, file := range files {
		pngFiles = append(pngFiles, "./images/"+file.Name())
	}

	pdf := gofpdf.New("P", "cm", "A4", "")

	for _, file := range pngFiles {
		pdf.AddPage()
		info := pdf.RegisterImage(file, "")
		pdf.Image(file, 0, 0, info.Width(), info.Height(), false, "", 0, "")
	}

	// Save the PDF to a file
	err = pdf.OutputFileAndClose("./output/output.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
