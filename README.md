# Rssman

### Telegram bot, able to parse RSS, use the openai api to generate a summary.

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

- Telegram bot, which is able to parse RSS and use the openai api to generate a short summary and a link to the source.
- It's also possible to add sources through commands in the telegram bot.

## Installation

```shell
git clone https://github.com/kenjitheman/rssman
```

- Install all dependencies:

```shell
go mod tidy
```

## Usage

- Create file with name "config.local.hcl" and inside the file create:

```hcl
tolegram_bot_token = "your token"
openai_token = "your token"
telegram_channel_id = -1116543419111    // example id

openai_promt = "Make a summary of the text"   // to make summary
```

- Run the app using:

```shell
cd cmd
go run main.go
```

- Or using docker:

```shell
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## Commands

- To add source:

```go
/addsource {"name": "source name", "url": "source url"}
```

- To get sources:

```go
/listsources
```

- To delete source:

```go
/deletesource {"name": "source name"}
```

- To set source priority:

```go
/setpriority 
```

## Contributing

- Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

- Please make sure to update tests as appropriate.

## license

- [MIT](https://choosealicense.com/licenses/mit/)
