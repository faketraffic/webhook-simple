package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your webhook URL: ")
    webhookURL, _ := reader.ReadString('\n')
    webhookURL = strings.TrimSpace(webhookURL)

    fmt.Print("Enter your message: ")
    message, _ := reader.ReadString('\n')
    message = strings.TrimSpace(message)

    message = strings.ReplaceAll(message, "@everyone", "@everyone\u200B")

    resp, err := http.Post(
        webhookURL,
        "application/json",
        strings.NewReader(`{"content":"`+message+`"}`),
    )
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}