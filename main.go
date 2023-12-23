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
    U, _ := reader.ReadString('\n')
    U = strings.TrimSpace(webhookURL)

    fmt.Print("Enter your message: ")
    m, _ := reader.ReadString('\n')
    m = strings.TrimSpace(message)

    m = strings.ReplaceAll(m, "@everyone", "@everyone\u200B")

    resp, err := http.Post(
        U,
        "application/json",
        strings.NewReader(`{"content":"`+m+`"}`),
    )
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}
