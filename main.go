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

	pdf := gofpdf.New("L", "cm", "A3", "")
	pageWidth, pageHeight := pdf.GetPageSize()

	for _, file := range pngFiles {
		pdf.AddPage()

		pdf.RegisterImage(file, "")
		// Get the dimensions of the image
		imageInfo := pdf.GetImageInfo(file)
		if imageInfo == nil {
			log.Printf("Failed to get image info for %s", file)
			continue
		}
		imgWidth, imgHeight := imageInfo.Width(), imageInfo.Height()

		// Calculate the aspect ratio
		aspectRatio := imgWidth / imgHeight
		var newWidth, newHeight float64
		if aspectRatio > pageWidth/pageHeight {
			newWidth = pageWidth
			newHeight = pageWidth / aspectRatio
		} else {
			newHeight = pageHeight
			newWidth = pageHeight * aspectRatio
		}
		pdf.Image(file, 0, 0, newWidth, newHeight, false, "", 0, "")
	}

	// Save the PDF to a file
	err = pdf.OutputFileAndClose("./output/output.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
