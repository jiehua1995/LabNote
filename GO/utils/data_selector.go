package utils

import (
    "fmt"
    "time"
)

func GetDate() time.Time {
    for {
        var userInput string
        fmt.Print("Press Enter to use the current system date, or enter a date (format: YYYY-MM-DD): ")
        fmt.Scanln(&userInput)

        if userInput == "" {
            current := time.Now()
            fmt.Printf("The current system date is: %s\n", current.Format("2006-01-02"))
            return current
        } else {
            userDate, err := time.Parse("2006-01-02", userInput)
            if err == nil {
                fmt.Printf("You entered the date: %s\n", userDate.Format("2006-01-02"))
                return userDate
            } else {
                fmt.Println("Invalid date format. Please use the format YYYY-MM-DD. Try again.")
            }
        }
    }
}
