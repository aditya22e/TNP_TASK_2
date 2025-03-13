package main

import (
	"fmt"
	"log"

	"TNP_TASK_2/internal/services"
)

func main() {
	// path to dummy
	dummyExcelPath := "assets/dummy.xlsx"

	// process dummy
	filledTemplates, err := services.ProcessExcelTemplate(dummyExcelPath)
	if err != nil {
		log.Fatalf("error processing dummy excel file: %v", err)
	}

	// print the filled templates
	for i, tmpl := range filledTemplates {
		fmt.Printf("Record %d:\n%s\n\n", i+1, tmpl)
	}
}
