# inst2tg

## Overview
`inst2tg` is a Golang application designed to synchronize Instagram stories to Telegram. It utilizes `go-tdlib`, a Telegram client library, to interact with the Telegram Developer API.

## Features
- Synchronize Instagram stories directly to your Telegram account (todo: right now we are using url of remote file both photo or video).
- Utilize Telegram's robust API for secure and efficient transfers.
- Easy-to-use and straightforward setup.

## Requirements
- Go (latest version recommended)
- Telegram API credentials (API ID and Hash)
- Instagram account credentials (todo: later)

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/inst2tg.git
   ```
2. Navigate to the cloned directory:
   ```sh
   cd inst2tg
   ```
3. Install dependencies:
   ```sh
   go get .
   ```

## Usage
1. Set up your Telegram API credentials in the application.
2. Configure your Instagram account details.
3. Run the application:
   ```sh
   go run main.go
   ```

## Disclaimer
This project is not affiliated with Instagram or Telegram and uses publicly available APIs for synchronization.

