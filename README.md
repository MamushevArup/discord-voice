# Discord Voice Program

This program allows you to capture and record voice conversations in Discord voice channels, as well as convert audio files between different formats. It also provides functionality to send recorded audio back to the voice channel for downloading.

Download format supported -> wav, ogg, mp3

*Command available -> help, download-mp3, download-wav, download-ogg*

__See help.txt file to get more familiar with proper commands__

### Prerequisites

- Go installed on your system. You can download and install it from the official Go website: https://golang.org/

- FFmpeg installed on your system. You can download and install it from the official FFmpeg website: https://ffmpeg.org/
- FFmpeg mus be on your system path if you run on windows

### Steps

1. Clone this repository to your local machine:
2. Navigate to the project directory:
3. Install dependency with ```go mod tidy```
4. Set up your Discord bot:

- Create a new Discord bot on the Discord Developer Portal: https://discord.com/developers/applications
- Copy the bot token.
- Add bot to your server
- Set the bot token as an environment variable named `TOKEN`. .env.example is provided
5. Set up your aws s3 bucket
  - Create a new bucket and get secret keys, and also specify aws-region
  - Allow to aws to public read from the other hosts
  - Specify the required values as in the .env.example shows
6. Run the program ```go run cmd/main/main.go``` command

### Explanation

#### Approach 
I use clean code architecture approach. Where my


