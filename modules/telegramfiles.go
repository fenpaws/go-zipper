package modules

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

// AddPhotoToFiles adds a photo to the files map if it is present in the update message
// The map key is the file name and the value is the file URL
// It returns an error if there is an issue getting the file URL
func AddPhotoToFiles(bot *tgbotapi.BotAPI, update tgbotapi.Update, files map[string]string) error {
	// Return if there is no photo in the update message
	if update.Message.Photo == nil {
		return nil
	}

	// Get the file ID of the largest photo in the update message
	photoSize := update.Message.Photo
	photoID := photoSize[len(photoSize)-1].FileID

	// Get the file URL for the photo
	photoUrl, err := bot.GetFileDirectURL(photoID)
	if err != nil {
		return err
	}

	// Add the file to the files map
	files[FileNameGenerator(photoUrl)] = photoUrl
	return nil
}

// AddFileToFiles adds the file in the given update message to the files map.
// It accepts files (raw images and so on), GIFs, voice messages, and single stickers.
// If the file has a caption, it will be used as the file name.
// Otherwise, a default file name will be generated.
// Returns an error if there was an issue getting the file URL or adding the file to the map.
func AddFileToFiles(bot *tgbotapi.BotAPI, update tgbotapi.Update, files map[string]string) error {

	// Return if there is no document, voice message, animation, or sticker in the update message
	if update.Message.Document == nil && update.Message.Voice == nil && update.Message.Animation == nil && update.Message.Sticker == nil {
		return nil
	}

	// Get the file URL for the file in the update message
	var fileUrl string
	var err error
	if update.Message.Document != nil {
		fileUrl, err = bot.GetFileDirectURL(update.Message.Document.FileID)
	} else if update.Message.Voice != nil {
		fileUrl, err = bot.GetFileDirectURL(update.Message.Voice.FileID)
	} else if update.Message.Sticker != nil {
		fileUrl, err = bot.GetFileDirectURL(update.Message.Sticker.FileID)
	}
	if err != nil {
		return err
	}

	// Set the file name based on the type of file
	var fileName string
	if update.Message.Document != nil {
		if update.Message.Caption != "" {
			fileName = update.Message.Caption
		}

		if update.Message.Animation.FileName != "" {
			fileName = update.Message.Animation.FileName
		}

		fileName = update.Message.Document.FileName
	} else if update.Message.Voice != nil {
		if update.Message.Caption != "" {
			fileName = update.Message.Caption
		}
		now := time.Now().Format("2006-01-02")
		fileName = "voice_" + now + "_" + randSeq(3)
	} else if update.Message.Sticker != nil {
		if update.Message.Caption != "" {
			fileName = update.Message.Caption
		}
		now := time.Now().Format("2006-01-02")
		fileName = "sticker_" + update.Message.Sticker.SetName + "_" + now + "_" + randSeq(3)
	}

	// Add the file to the files map
	files[fileName] = fileUrl
	return nil
}
