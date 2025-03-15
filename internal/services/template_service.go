package services

import (
	"travel/m/internal/models"
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/xuri/excelize/v2"
)

// FillTemplate reads the template file and fills it with data from a PersonTravel record
func FillTemplate(record models.PersonTravel) (string, error) {
	// reads the template file from the templates folder
	tmplData, err := os.ReadFile("templates/template.txt")
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}

	// Parse the template.
	tmpl, err := template.New("travelTemplate").Parse(string(tmplData))
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template with the record data.
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, record); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.String(), nil
}

// ProcessExcelTemplate opens the Excel file, reads each row, and fills the template with its data
func ProcessExcelTemplate(filePath string) ([]string, error) {
	// open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	defer f.Close()

	// use the first sheet
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet: %w", err)
	}

	// Ensure there is at least one header row and one data row
	if len(rows) < 2 {
		return nil, fmt.Errorf("no data rows found in excel file")
	}

	// The first row is assumed to be the header
	header := rows[0]
	headerMap := make(map[string]int)
	for i, colName := range header {
		headerMap[colName] = i
	}

	var results []string

	// process each subsequent row as a record
	for _, row := range rows[1:] {
		person := models.PersonTravel{
			Name:     getCell(row, headerMap, "Name"),
			Visited1: getCell(row, headerMap, "Visited1"),
			Visited2: getCell(row, headerMap, "Visited2"),
			Visited3: getCell(row, headerMap, "Visited3"),
			Planned1: getCell(row, headerMap, "Planned1"),
			Planned2: getCell(row, headerMap, "Planned2"),
			Planned3: getCell(row, headerMap, "Planned3"),
		}

		// Fill the template with the data from the current record
		filled, err := FillTemplate(person)
		if err != nil {
			return nil, fmt.Errorf("error filling template for row: %w", err)
		}
		results = append(results, filled)
	}
	return results, nil
}

// getCell safely extracts a cell value
func getCell(row []string, headerMap map[string]int, key string) string {
	if idx, ok := headerMap[key]; ok && idx < len(row) {
		return row[idx]
	}
	return ""
}
