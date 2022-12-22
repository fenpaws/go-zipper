# Telegram Zipper bot
![docker-build](https://github.com/fenpaws/go-zipper/actions/workflows/docker-build-publish.yml/badge.svg)

This bot allows you to send files to it and receive a ZIP archive of those files in return. Simply write the bot [@GoZipper_bot](http://t.me/GoZipper_bot) and forward or send it some files, then write /zip to receive your ZIP archive.

## Hosting

If you want to host your own bot

1. Copy the [docker-compose.yml](./docker-compose.yml) file
2. Create a config.env file and copy the snippet below.

   ```env
   BOT_TOKEN=""
   CUSTOM_TELEGRAM_API_SERVER_URL=""
   DEBUG="true"
   ```

3. Create a bot with [BotFather](http://t.me/BotFather) and copy your token
4. Paste your token into the ``BOT_TOKEN`` section of config.env. This is also the time to set up commands 
   - see [commands.md](commands.md) for more information
5. Run the bot with ``docker-compose up -d``


## Features & Limitations

### Features

- Can ZIP files, either as a message or as a forward
- Deletes all files immediately after creating the ZIP archive

### Limitations

- Not all file types can be downloaded and zipped. See [Ideas & ToDo's](#ideas--todos) for more information.
- Cannot download files larger than 20MB (API limitation)
- Images cannot be downloaded in their original quality. If you need the original quality, send the file as a document (API limitation)

## Ideas & ToDo's

- [ ] Allow user to set a Password for the ZIP
  - [X] Core Functionality
- [ ] Securely erase all files that have been downloaded
- [ ] pin code; for special access (larger file uploads and more)
- [ ] remove of files that are older than x minutes (1min?)
- [ ] compress in different formats? 7zip, rar and more?
- [ ] special compression algorithm, eg. zstd
- [ ] Allow more files to be processed:
   - [X] Pictures
   - [X] Videos
   - [X] Voice messages
   - [X] Stickers
      - [X] Individual
      - [ ] Packs
- [ ] Notifications, when files are finished downloading and compressed
- [X] Configuration through environment variables and .ENV file

---

## Used Library's

- [env](https://github.com/caarlos0/env) - Configuration
- [Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api) - Telegram API
