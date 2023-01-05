package fileHandler

import (
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AddPhotoToFiles adds a photo to the files map if it is present in the update message
// The map key is the file name and the value is the file URL
// It returns an error if there is an issue getting the file URL
func AddPhotoToFiles(bot *tgbotapi.BotAPI, update tgbotapi.Update) (err error, name string, url string) {
	// Return if there is no photo in the update message
	if update.Message.Photo == nil {
		return nil, "", ""
	}

	// Get the file ID of the largest photo in the update message
	photoSize := update.Message.Photo
	photoID := photoSize[len(photoSize)-1].FileID

	// Get the file URL for the photo
	photoUrl, err := bot.GetFileDirectURL(photoID)
	if err != nil {
		return err, "", ""
	}

	// Add the file to the files map
	return nil, helper.FileNameGenerator(photoUrl), photoUrl
}
