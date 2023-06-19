# VoiceRSS Cache Layer
A caching layer for the [VoiceRSS](https://www.voicerss.org/) API. It caches the results from unique requests to avoid consuming the API daily usage limit when making a request that was made before.

## Setup
### VoiceRss API KEY
You will need a [VoiceRSS API Key](https://www.voicerss.org/).

### `.env`
Create a `.env` in the app root directory. And then:
```env
VOICERSS_API_KEY="YOUR KEY"
SERVER_MODE="release" # options: debug, release
```

## Building and Running
```sh
go run . # to run the app

go build . # to compile the app into an executable
```

A `Dockerfile` is also available.

## Usage
You can check the api spec inside de `docs` folder.

### Example
```sh
curl -X 'GET' \
  'http://localhost:8080/api/v1/tts?src=hello&hl=en-us' \
  -H 'accept: text/plain'
```
The query parameters are the same as [VoiceRSS](https://www.voicerss.org/api/). Except for `key`, which you've set inside the `.env` file, and `b64`. You won't be using these two.
