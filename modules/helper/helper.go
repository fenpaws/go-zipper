package helper

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"os"
	"regexp"
)

var regex = regexp.MustCompile(`(?m)(?:.+\/)(.+)`)

func TempFolder(folderPath string, identifier string) (string, error) {
	// Create a unique temp download folder
	DownloadDir, err := os.MkdirTemp(folderPath, "*-"+identifier)
	if err != nil {
		return "", err
	}
	return DownloadDir, nil
}

func FileNameGenerator(fileURL string) string {
	return regex.FindAllStringSubmatch(fileURL, -1)[0][1]
}

func Clear(folderPath string) {
	var err error

	// Use os.Stat to get information about the file or folder
	info, err := os.Stat(folderPath)
	if err != nil {
		log.Printf("Error getting file info: %v", err)
		return
	}

	// Check if the file is a folder or a regular file
	if info.IsDir() {
		// It's a folder, so remove it with os.RemoveAll
		err = os.RemoveAll(folderPath)
		log.Printf("Temp folder %s deleated!", folderPath)
	} else {
		// It's a regular file, so remove it with os.Remove
		err = os.Remove(folderPath)
		log.Printf("ZIP file %s deleated!", folderPath)
	}

	if err != nil {
		log.Printf(err.Error())
	}
}

func SendTelegramMessage(bot tgbotapi.BotAPI, m tgbotapi.Message, message string) {
	msg := tgbotapi.NewMessage(m.Chat.ID, message)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf(err.Error())
	}
}

func RandSeq(n int) string {
	letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
