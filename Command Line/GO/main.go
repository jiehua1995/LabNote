package main

import (
    "fmt"
    "os"
    "time"
    "LabNote_GO/utils"
)

func main() {
    // date selection
    selectedDate := utils.GetDate()
    formated_selectedDate := selectedDate.Format("2006-01-02")
    fmt.Printf("Selected date is: %s\n", selectedDate.Format("2006-01-02"))

    // category selection
    fmt.Println("=================\nPlease select one category (1-4):")
	fmt.Println("1. Experimental log")
	fmt.Println("2. Discussion")
	fmt.Println("3. Reading")
	fmt.Println("4. Inspiration\n=================")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Inavilable selection:", err)
		os.Exit(1)
	}

	var category string
    var content string
	switch choice {
	case 1:
		category = "Experimental log"
		content = "## Experiments\n\n"
	case 2:
		category = "Discussion"
        content = "## Content\n\n"
	case 3:
		category = "Reading"
        content = "**Title:**\n\n**DOI:**\n\n**URL:**\n\n**Published:**\n\n**Journal:**\n\n##Notes\n\n"
	case 4:
		category = "Inspiration"
        content = "## Content\n\n"
	default:
		fmt.Println("Inavilable input, please select from 1 to 4。")
		os.Exit(1)
	}

	fmt.Printf("Selected category is: %s\n", category)

	yaml := "---\nDate: " + formated_selectedDate + "\nAuthor: Jie Hua\nEmail: Jie.Hua@lmu.de\nCategory: "+ category + "\n---\n\n[toc]\n\n## Basic Information\n\n"

	MdContent := yaml + content

    // merge file name
    filename := category + "_" + formated_selectedDate + ".md"
    // create markdown file
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Could not create the markdown file:", err)
        return
    } 
    defer file.Close()
    fmt.Printf("File created: %s\n", filename)

    // write markdown information
    file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Could not open the file: ", err)
        return
    }
    defer file.Close()


    _, err = file.WriteString(MdContent)
    if err != nil {
        fmt.Println("Error in writing into the file:", err)
        return
    }
    fmt.Println("Markdown file generated successfully")

    fmt.Println("Software will be closed after ")
    fmt.Println("2 seconds...............")
    time.Sleep(1 * time.Second)
    fmt.Println("1 second..........")
    time.Sleep(1 * time.Second)
}