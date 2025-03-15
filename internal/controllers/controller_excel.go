package controllers

import (
	"travel/m/internal/services"
	"travel/m/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// UploadExcel handles the Excel file upload, processes it, and returns the filled templates
func UploadExcel(w http.ResponseWriter, r *http.Request) {
	// retrieve the file from the form-data.
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "failed to get.....", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// rnsure the uploads folder exists
	uploadDir := "./uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	// saves the upload
	filePath := filepath.Join(uploadDir, header.Filename)
	if err := utils.SaveUploadedFile(file, filePath); err != nil {
		http.Error(w, "failed to save....", http.StatusInternalServerError)
		return
	}

	// Process the Excel file and fill the template.
	filledTemplates, err := services.ProcessExcelTemplate(filePath)
	if err != nil {
		log.Println("error processing template:", err)
		http.Error(w, "Error processing file", http.StatusInternalServerError)
		return
	}

	// Return the filled templates as plain text.
	for i, content := range filledTemplates {
		fmt.Fprintf(w, "Record %d:\n%s\n\n", i+1, content)
	}
}
