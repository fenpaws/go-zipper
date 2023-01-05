package fileHandler

import (
	"fmt"
	"github.com/fenpaws/go-zipper/modules/helper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

// AddFileToFiles adds the file in the given update message to the files map.
// It accepts files (raw images and so on), GIFs, voice messages, and single stickers.
// If the file has a caption, it will be used as the file name.
// Otherwise, a default file name will be generated.
// Returns an error if there was an issue getting the file URL or adding the file to the map.
func AddFileToFiles(bot *tgbotapi.BotAPI, update tgbotapi.Update) (err error, name string, url string) {

	// Return if there is no document, voice message, animation, or sticker in the update message
	if update.Message.Document == nil && update.Message.Voice == nil && update.Message.Animation == nil && update.Message.Sticker == nil {
		return nil, "", ""
	}

	// Get the file URL for the file in the update message
	var fileUrl string

	if update.Message.Document != nil {
		fileUrl, err = bot.GetFileDirectURL(update.Message.Document.FileID)
	}
	if update.Message.Voice != nil {
		fileUrl, err = bot.GetFileDirectURL(update.Message.Voice.FileID)
	}
	if update.Message.Sticker != nil {
		fileUrl, err = bot.GetFileDirectURL(update.Message.Sticker.FileID)
	}
	if err != nil {
		return err, "", ""
	}

	// Set the file name based on the type of file
	var fileName string
	if update.Message.Document != nil {
		if update.Message.Caption != "" {
			fileName = update.Message.Caption
		}

		if update.Message.Animation != nil {
			fileName = update.Message.Animation.FileName
		}

		fileName = update.Message.Document.FileName

	} else if update.Message.Voice != nil {
		if update.Message.Caption != "" {
			fileName = update.Message.Caption
		}
		now := time.Now().Format("2006-01-02")
		fileName = fmt.Sprintf("voice_%s_%s_%s", update.Message.Sticker.SetName, now, helper.RandSeq(3))

	} else if update.Message.Sticker != nil {
		if update.Message.Caption != "" {
			fileName = update.Message.Caption
		}
		now := time.Now().Format("2006-01-02")
		fileName = fmt.Sprintf("sticker_%s_%s_%s", update.Message.Sticker.SetName, now, helper.RandSeq(3))
	}

	return nil, fileName, fileUrl
}
