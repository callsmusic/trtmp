FROM debian:latest
RUN apt update && apt upgrade -y && apt install ffmpeg curl wget -y
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl && chmod a+rx /usr/local/bin/youtube-dl
WORKDIR /app
CMD curl -s https://raw.githubusercontent.com/callsmusic/trtmp/main/scripts/install.debian.sh | bash && chmod +x linux-amd64 && ./linux-amd64
