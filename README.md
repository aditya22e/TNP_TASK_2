Automated Excel Template Filler

This Go project automates the process of filling template files using data from an Excel sheet. It reads an Excel file containing a list of people along with the places they have visited and plan to visit, then uses a predefined text template to generate personalized messages.

Features

- Reads an Excel file (dummy.xlsx) containing travel data.

- Fills a text template with data extracted from the Excel file.

- Provides an API to upload an Excel file and get personalized output.

- Uses excelize for reading Excel files and Go templates for text generation.

The excel file have 7 columns and storing the name of the person and 3 places that they've visited and 3 places that they plan to visit
while template.txt shows those stuff as
```text
Hello {{.Name}},

Here are the top 3 places you visited:
1. {{.Visited1}}
2. {{.Visited2}}
3. {{.Visited3}}

And here are the top 3 places you plan to visit:
1. {{.Planned1}}
2. {{.Planned2}}
3. {{.Planned3}}

Happy Travels!
```

Project Structure

```text
/TNP_TASK_2
├── assets
│   └── dummy.xlsx            # Sample Excel file with travel data.
├── cmd
│   └── main.go               # Entry point.
├── internal
│   ├── controllers
│   │   └── excel_controller.go  # API controller for file upload.
│   ├── middleware
│   │   └── validate_excel.go    # Middleware for validating Excel file.
│   ├── models
│   │   └── excel.go             # Data model for Excel records.
│   └── services
│       └── template_service.go  # Core logic to process Excel and templates.
├── pkg
│   └── utils
│       └── file_utils.go        # Utility functions for file handling.
├── templates
│   └── template.txt           # Predefined text template.
├── uploads                    # Directory for uploaded files.
├── go.mod                     # Go module file.
├── go.sum                     # Dependencies file.
└── README.md                  # Project documentation.
```
