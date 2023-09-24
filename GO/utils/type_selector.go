package utils

import (
    "fmt"
    "os"
)

func GetCategory() string {
	fmt.Println("请选择一个类别:")
	fmt.Println("1. Experimental log")
	fmt.Println("2. Discussion")
	fmt.Println("3. Reading")
	fmt.Println("4. Inspiration")

	var choice int
	fmt.Print("Please select document category (1-4): ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Inavilable selection:", err)
		os.Exit(1)
	}

	var category string
	switch choice {
	case 1:
		category = "Experimental log"
	case 2:
		category = "Discussion"
	case 3:
		category = "Reading"
	case 4:
		category = "Inspiration"
	default:
		fmt.Println("Inavilable input, please select from 1 to 4。")
		os.Exit(1)
	}

	return category
}
