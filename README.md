<h2 align="left">telegram bot for self use, which is able to parse RSS and use the openai api to generate a short summary</h2>

###

<img align="right" height="200" src="https://media.tenor.com/6OCEdkhjHKUAAAAC/d4dj-first-mix-d4dj.gif"  />

###

<div align="left">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

<div align="center">
  <a href="#" target="_blank">
    <img src="https://img.shields.io/static/v1?message=bot_ai_rss_man&logo=telegram&label=&color=2CA5E0&logoColor=white&labelColor=&style=for-the-badge" height="35" alt="telegram logo"  />
  </a>
</div>

###

cool telegram bot 
which is able to parse RSS and use the openai api to generate a short summary and a link to the source
it is also possible to add sources through commands in the telegram bot

## Installation

```
go get github.com/kenjitheman/rss_ai_tgbot
```

- or
```
git clone https://github.com/kenjitheman/rss_ai_tgbot 
```

```
go get go.mod
```

## usage
- create file with name "config.local.hcl" and inside the file create:
```
tolegram_bot_token = "your token"
openai_token = "your token"
telegram_channel_id = -1116543419111    // example id

openai_promt = "Make a summary of the text"   // to make summary
```
- you can run the app using ->
```
cd cmd
go run main.go
```

- or using docker ->
```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## Comands
- to add source:
```
/addsource {"name": "source name", "url": "source url"}
```
- to get sources:
```
/listsources
```
- to delete source:
```
/deletesource {"name": "source name"}
```
- to set source priority:
```
/setpriority 
```

## Contributing

- Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

- Please make sure to update tests as appropriate.

## License

- [MIT](https://choosealicense.com/licenses/mit/)
