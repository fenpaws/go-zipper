package helper

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
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

	if strings.Contains(folderPath, "zip") {
		err = os.RemoveAll(folderPath + ".zip")
		log.Printf("ZIP file %s deleated!", folderPath+".zip")
	} else {
		err = os.RemoveAll(folderPath)
		log.Printf("Temp folder %s deleated!", folderPath)
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

func ErrorFileToBig(err error, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if strings.Contains(err.Error(), "file is too big") {
		msg := "❌ - Im sorry but i cant download that file, its more then 20MB (API Limitation)"
		SendTelegramMessage(*bot, *update.Message, msg)
	}
}
