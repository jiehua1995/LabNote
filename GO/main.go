package main

import (
    "fmt"
    "LabNote_GO/utils"
)

func main() {
    selectedDate := utils.GetDate()
    fmt.Printf("Selected date is: %s\n", selectedDate.Format("2006-01-02"))
	fileType := utils.GetFileType()

    switch fileType {
    case "txt":
        fmt.Println("You selected a text file.")
        // 在这里执行与文本文件相关的操作
    case "pdf":
        fmt.Println("You selected a PDF file.")
        // 在这里执行与PDF文件相关的操作
    case "jpg":
        fmt.Println("You selected a JPG file.")
        // 在这里执行与JPG文件相关的操作
    default:
        fmt.Println("Unknown file type.")
    }
}
