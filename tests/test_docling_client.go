package main

import (
	"fmt"
	"os"

	"github.com/TechnoDamo/embedding_service_go.git/internal/client"
)

func main() {
	doclingURL := "http://localhost:5001"
	dc := client.NewDoclingClient(doclingURL)

	filePath := "Test.pdf"
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	extractedText, err := dc.ExtractText(filePath, fileData)
	if err != nil {
		fmt.Println("Docling extraction failed:", err)
		return
	}

	fmt.Println("Extracted text from file:")
	fmt.Println(extractedText)
}
