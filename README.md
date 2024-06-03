# Telegram Bot for qBitTorrent and Plex Integration.

This project is a Telegram bot that seamlessly integrates with the qBittorrent client and Plex media server, enabling easy torrent management and local streaming. Users can interact with the bot to add magnet links to the download queue, and once downloaded, the files are available for streaming on a local Plex media server.

## Features

- Seamless integration with Telegram for easy torrent management.
- Automatic downloading of magnet links using qBittorrent.
- Local streaming of downloaded content via Plex media server.
- Easy deployment using Docker and Docker Compose.

## Prerequisites

Before running the Telegram bot, ensure you have the following prerequisites installed:

- Docker
- Docker Compose

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone git@github.com:PrantaDas/Qbit-Go.git
    ```
2. Navigate to the project directory:
    ```bash
    cd Qbit-Go
    ```
3. Build and start the services using Docker Compose:
    ```bash
    docker-compose up -d
    ```
## Usage
1. Start the Telegram bot by sending /start command in the Telegram chat.
2. Add a magnet link to the download queue using the /magnet <link> command.
3. Monitor the download progress in the qBittorrent web interface (accessible at http://localhost:8090).
4. Once the download is complete, the file will be available for streaming on the Plex media server (accessible at http://localhost:8000).

## Configuration
1. Telegram Bot Token: Set your Telegram bot token in the docker-compose.yml file under the `app` service environment variables.

2. qBittorrent Web UI Credentials: Update the username and password in the `docker-compose.yml` file under the app service environment variables if required.
3. Plex Media Server Configuration: Customize Plex configuration options in the `docker-compose.yml` file under the `plex` service environment variables.


### Set password
Add this in `config/config/qBittorrent.conf`. This will set the username `admin` and password `adminadmin`

```conf
[Preferences]
WebUI\Password_PBKDF2="@ByteArray(ARQ77eY1NUZaQsuDHbIMCA==:0WMRkYTUWVT9wVvdDtHAjU9b3b7uB8NR1Gur2hmQCvCDpm39Q+PsJRJPaCU51dEiz+dTzh8qbPsL8WkFljQYFQ==)"
```

## Acknowledgements
Special thanks to [Rakibul Yeasin](github.com/dreygur) for his guidance and contributions to this project. The implementation of the bot is also available in Rust. You can check out here [ChatQbit](https://github.com/dreygur/ChatQBit).