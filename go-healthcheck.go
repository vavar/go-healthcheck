package main

import (
    "log"
    "fmt"
    "time"
    "os"
    "go-healthcheck/healthcheck"
)

func main()  {

    args := os.Args[1:]
    filePath := args[0]

    h, err := healthcheck.NewChecker(filePath)
    if (err != nil) {
        log.Fatal(err)
        return
    }

    fmt.Println("Perform website checking...")
    result := h.Check()
    //Send report via api
    if err := h.Report(result); err != nil {
        fmt.Println(err)
    }
    fmt.Println("Done!")

    //Print operation result to console
    fmt.Printf("\nChecked websites: %d\n", result.TotalWebsites)
    fmt.Printf("Successful websites: %d\n", result.Success)
    fmt.Printf("Failure websites: %d\n", result.Failure)
    fmt.Printf("Total times to finished checking website: %v\n",
        time.Unix(0,result.TotalTime).Sub(time.Unix(0,0)))
}