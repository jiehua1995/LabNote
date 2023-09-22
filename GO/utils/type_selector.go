package utils

import (
    "fmt"
    "strings"
)

func GetFileType() string {
    fmt.Print("Enter the file type (e.g., txt, pdf, jpg): ")
    var fileType string
    fmt.Scanln(&fileType)
    return strings.ToLower(fileType)
}
