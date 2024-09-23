package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Response struct {
    Status      int    `json:"Status"`
    DevelopedBy string `json:"Devloped by"`
    Channel     string `json:"Channel"`
    Result      string `json:"Result"`
}

func main() {
    var userInput string
    fmt.Print("متن خود را وارد کنید: ")
    fmt.Scanln(&userInput)

    url := fmt.Sprintf("http://api.api-code.ir/api/ai-chatbot/?text=%s", userInput)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("خطا در ارسال درخواست:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("خطا در خواندن پاسخ:", err)
        return
    }

    var response Response
    if err := json.Unmarshal(body, &response); err != nil {
        fmt.Println("خطا در تجزیه JSON:", err)
        return
    }

    fmt.Println("نتیجه:", response.Result)
}
