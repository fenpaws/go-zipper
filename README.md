# Telegram Zipper bot

Send some files to the bot, and it makes you a ZIP!

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

- Not everyting can be downloadad and zipped, see [Ideas & ToDo's](#ideas--todos)
- Cant download files bigger than 20MB (API limitation)
- Images cant be downloaded in the Original quality, send as file if you need original quality (API limitation)

## Ideas & ToDo's

- [ ] Allow user to set a Password for the ZIP
- [ ] Securely erase all files that have been downloaded
- [ ] pin code; for special access (larger file uploads and more)
- [ ] remove of files that are older than x minutes (1min?)
- [ ] compress in different formats? 7zip, rar and more?
- [ ] special compression algorithm, eg. zstd
- [ ] Allow more files to be processed:
   - [X] Pictures
   - [ ] Videos
   - [ ] Voice messages
   - [ ] Stickers
      - [ ] Individual
      - [ ] Packs
- [ ] Notifications, when files are finished downloading and compressed
- [X] Configuration through environment variables and .ENV file

---

## Used Library's

- [env](https://github.com/caarlos0/env) - Configuration
- [Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api) - Telegram API
