# Qbit Go
* The project aims to leverage the download of torrent file through a telegram bot.
Simply add the download link to the bot and it will start downloading the torrent file in local mode using Qbit Torrent client. The downloaded file will be streamed in the Plex server that will be accessed through the local network along with other torrent file available.

### Set password
Add this in `config/config/qBittorrent.conf`. This will set the username `admin` and password `adminadmin`

```conf
[Preferences]
WebUI\Password_PBKDF2="@ByteArray(ARQ77eY1NUZaQsuDHbIMCA==:0WMRkYTUWVT9wVvdDtHAjU9b3b7uB8NR1Gur2hmQCvCDpm39Q+PsJRJPaCU51dEiz+dTzh8qbPsL8WkFljQYFQ==)"
```

 The idea of implementing the bot was from [Rakibul Yeasin](github.com/dreygur). The implementation of the bot is also available in Rust. You can check out here [ChatQbit]([https://github.com/dreygur/ChatQBit).