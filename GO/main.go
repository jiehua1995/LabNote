package main

import (
    "fmt"
    "os"
    "LabNote_GO/utils"
)

func main() {
    selectedDate := utils.GetDate()
    formated_selectedDate := selectedDate.Format("2006-01-02")
    fmt.Printf("Selected date is: %s\n", selectedDate.Format("2006-01-02"))
    fileType := utils.GetCategory()
    fmt.Printf("Selected category is: %s\n", fileType)

    // merge file name
    filename := fileType + "_" + formated_selectedDate + ".md"
    // create markdown file
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Could not create the markdown file:", err)
        return
    } 
    defer file.Close()
    fmt.Printf("File created: %s\n", filename)

    // write yaml information
    file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Could not open the file: ", err)
        return
    }
    defer file.Close()

    yaml := "---\nDate: " + formated_selectedDate + "\nauthor: Jie Hua\nemail: Jie.Hua@lmu.de\ncategory: "+ fileType + "\n---"
    _, err = file.WriteString(yaml)
    if err != nil {
        fmt.Println("Error in writing into the file:", err)
        return
    }
    fmt.Println("Basic information was written successfully")
}