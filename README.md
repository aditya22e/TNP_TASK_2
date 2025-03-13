# TNP_TASK_2
Automated Excel Template Filler

This Go project automates the process of filling template files using data from an Excel sheet. It reads an Excel file containing a list of people along with the places they have visited and plan to visit, then uses a predefined text template to generate personalized messages.

Features

Reads an Excel file (dummy.xlsx) containing travel data.

Fills a text template with data extracted from the Excel file.

Provides an API to upload an Excel file and get personalized output.

Uses excelize for reading Excel files and Go templates for text generation.


Project Structure

```text
/TNP_TASK_2
├── assets
│   └── dummy.xlsx            // Dummy Excel file with travel data.
├── cmd
│   └── main.go        // Entry point.
├── internal
│   ├── controllers
│   │   └── excel_controller.go  
│   ├── middleware
│   │   └── validate_excel.go    
│   ├── models
│   │   └── excel.go             
│   └── services
│       └── template_service.go  
├── pkg
│   └── utils
│       └── file_utils.go        
├── templates
│   └── template.txt           
├── uploads          
├── go.mod                     # Go module file.
├── go.sum                     # Dependencies file.
└── README.md                  # Project documentation.
```
