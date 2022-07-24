# TRTMP

A bot for streaming in Telegram live streams.

## Get Started

### Use a Cloud Platform

[![Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/callsmusic/trtmp)

### Manually

#### Install

1. Create a new directory.

```bash
mkdir trtmp && cd $_
```

2. Download a binary from the
   [latest release](https://github.com/callsmusic/trtmp/releases/latest). If
   you’re on Debian or a Debian-based Linux distribution, you can just run the
   following:

```bash
curl -s https://raw.githubusercontent.com/callsmusic/trtmp/main/scripts/install.debian.sh | bash
```

#### Configure

In order to configure the bot, you have to set the following environmental
variables:

- `BOT_TOKEN`
- `RTMP_KEY`
- `RTMP_URL`

Or put them in a `.env` in the directory.

#### Run

```bash
chmod +x ./your_binary # Only for Unix-like systems.
./your_binary
```

## Commands

> Note: Commands are triggered with the "/" prefix.

### stream

Starts streaming the provided video or URI leading to a media.

- You can stream from 100+ platforms. Including popular ones like YouTube,
  Twitter, Facebook, etc.

- There’s no limit in the duration of the provided media, it can be one second
  or a live stream.

- If streaming, the stream should be stopped with `/stop` before using this
  command.

### stop

Stops the active stream.

### now

Tells who is currently streaming what.

## License

[You just DO WHAT THE FUCK YOU WANT TO.](./LICENSE)
