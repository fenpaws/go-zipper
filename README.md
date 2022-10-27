# Telegram Zipper bot

Send some files to the bod, and it makes you a ZIP!

## Usage

Just write the bot [@GoZipper_bot](http://t.me/GoZipper_bot) and forward or send him some files, then just write `/zip` and you will get your data as a ZIP

## Hosting

If you want to host your own bot then its superr siple.

1. Copy the [docker-compose.yml](./docker-compose.yml) file
2. Create a config.env file and copy the snippet below.

   ```env
   BOT_TOKEN=""
   CUSTOM_TELEGRAM_API_SERVER_URL=""
   DEBUG="true"
   ```

3. Create a bot with [BotFather](http://t.me/BotFather) and copy your token
4. Paste your token into the `BOT_TOKEN` section.
5. now run the bot with `docker-compose up -d`

## Features & Limitations

### Features

- Can ZIP files, directly as a message or as a forward
- This bot will delete all your files Immediately after you make the ZIP

### Limitations

- Cant download files bigger than 20MB (API limitation)
- Can only download files that got send to the bot as files

## Ideas & ToDo's

- [ ] Allow user to set a Password for the ZIP
- [ ] Securely erase all files that have been downloaded
- [ ] pin code; for special access (larger file uploads and more)
- [ ] remove of files that are older than x minutes (1min?)
- [ ] compress in different formats? 7zip, rar and more?
- [ ] special compression algorithm, eg. zstd
- [ ] Allow all files (stickers, normal images, voice messages and more) to be downloadable
- [ ] Notifications, when files are finished downloading and compressed
- [x] Configuration through environment variables and .ENV file

---

## Used Library's

- [env](https://github.com/caarlos0/env) - Configuration
- [Bot API](github.com/go-telegram-bot-api/telegram-bot-api) - Telegram API
