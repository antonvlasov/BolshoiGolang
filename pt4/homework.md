# HW 4

Для понимания всех терминов смотрите в requirements.md в папке project_requirements. 

Глобально пока мы работаем только со статускодами 200 (http.StatusOK), 404 (http.StatusNotFound) и 400 (http.StatusBadRequest)

Детальнее прочитать про значение статусов можно [тут](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status).

Для передачи и получения значения используйте json следующего вида (прочитайте, про [struct tags](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go), чтобы сделать value с маленькой буквы)

```json
{"value":"some-value"}
```

## 1. Добавить /health endpoint

Необходимо добавить эндпойнт /health, который возвращает статус 200 (OK) на GET запрос

```
curl -v http://localhost:8090/health
```

```
hexzedels@hexzedels:~/Work/Teaching/BolshoiGolang$ curl -v http://localhost:8090/health
*   Trying ::1:8090...
* TCP_NODELAY set
* Connected to localhost (::1) port 8090 (#0)
> GET /health HTTP/1.1
> Host: localhost:8090
> User-Agent: curl/7.68.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 15 Oct 2024 21:40:21 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```

## 2. Добавить GET /scalar/get/:key endpoint

Необходимо добавить эндпойнт /scalar/get/:key, который возвращает json, содержащий скаляр, соответствующий ключу key

Request:

```
curl --request GET
  --url http://localhost:8090/scalar/get/123
```

Response (200 или 400 статус):

```
{"Value":"456"}
```

## 3. Добавить PUT /scalar/set/:key + json body endpoint

Необходимо добавить эндпойнт /scalar/set/:key, в который можно послать json в body, содержащий скаляр, который нужно поставить для ключа key

Request:

```
curl --request PUT \
  --url http://localhost:8090/scalar/set/123 \
  --header 'Content-Type: application/json' \
  --data '{
	"Value":"456"
}'
```

Response:

200 или 404 статус
