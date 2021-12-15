# vuc-ui

# Запросы/Ответы

host: "/"

### GET

"Content-Type": "application/json"

Ответ

```json
{ "connections": [ {
	"surname": "Frozen Yogurt4", 
	"class": 8080, "port": 8089,
	"status": "Online", 
	"session": "asdas21e909"
	}, ] }
```

### POST

"Content-Type": "application/json"

Запрос

```json
{
	"surname": "Frozen Yogurt4", 
	"class": 8080
}
```

Ответ

```json
{ "message":"Recosrd Not Found" }
```

### PUT

"Content-Type": "application/json"

Запрос

```json
{
	"session": "asdas21e909",
	"surname": "Frozen Yogurt4", 
	"class": 8080
}
```

Ответ

```json
{ "message":"Recosrd Not Found" }
```

### DELETE

"Content-Type": "application/json"

Запрос

```json
{
	"session": "asdas21e909"
}
```

Ответ

```json
{ "message":"Recosrd Not Found" }
```
## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
