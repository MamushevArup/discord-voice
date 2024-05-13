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
I use clean code architecture approach. Where my code consist of two layer usecase, handler. And third-party api integration is in adapters.
Communcate through the interfaces on the consumer side.

To make sure code is properly written you can check it through the linter. Configuration for the linter is provided in this code also.

This project is not provide a repo layer and any storage implementation. As long as there is no need to store any data.

To download it the conversation for particular voice room chat bot send message with some id to download and listen.

Here is how it look likes -- 

![Screenshot 2024-05-13 163842](https://github.com/MamushevArup/discord-voice/assets/93328884/665854ea-ad60-4933-af8e-96fb6c35482a)

## Notes
All variables can be changed in the code vars.go file

When voice channel created bot joins and wait for 5 second before start recording.
After five second it record 15 second or until room is empty (means no real user in there) also if in the room will two or three or more bots it will also stop record.

In aws presigned object is accessible 48 hours. 




