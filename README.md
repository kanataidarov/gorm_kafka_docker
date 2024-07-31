# messaggio_assignment
## Тестовое задание от Messaggio на вакансию Junior Go Developer
Разработать микросервис на Go, который будет принимать сообщения через HTTP API, сохранять их в PostgreSQL, 
а затем отправлять в Kafka для дальнейшей обработки. Обработанные сообщения должны помечаться. 
Сервис должен также предоставлять API для получения статистики по обработанным сообщениям.

## Удаленный запуск
### Пререквизиты
- HTTP-клиент или браузер.

### Запуск
API доступно для вызова на ` `. \
Далее этот URL будет помечаться как `<Базовый URL>`.

### Отправка через HTTP API
На `<Базовый URL>/applications`, методом HTTP POST, отправить тестовый payload: 
```json
{
    "name": "Канат Айдаров",
    "email": "kanataidarov@yahoo.com",
    "position": "Junior Go Developer"
}
```

Пример отправки через curl:
```shell
curl --location 'https://messaggio-assignment-99487939e4bb.herokuapp.com/applications' \
--request POST \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Канат Айдаров",
    "email": "kanataidarov@yahoo.com",
    "position": "Junior Go Developer"
}'
```
Пример отправки через Postman: \
![image](https://github.com/user-attachments/assets/47f8d232-0c63-4b69-8ca6-f9fd5ed17aa4)

### Получение статистики
На `<Базовый URL>/application` отправить методом HTTP GET. \
Пример отправки через curl:
```shell
curl --location 'https://messaggio-assignment-99487939e4bb.herokuapp.com/applications' \
--request POST --header 'Content-Type: application/json'
```

Пример отправки через Postman: \
![image](https://github.com/user-attachments/assets/6f877029-1cef-44ce-a792-b3cfd1f46bb2)

## Локальный запуск

### Пререквизиты
- Golang версии 1.22+. [Инструкции по установке](https://go.dev/doc/install)
- Docker. [Инструкции по установке](https://docs.docker.com/get-docker/)
- docker-compose. [Инструкции по установке](https://docs.docker.com/compose/install/standalone/)

### Запуск
- Для `Windows`, в CMD, набрать `./start.bat`, находясь в корне проекта. 
- Для `Linux / MacOS`, в терминале, набрать `./start.sh`, находясь в корне проекта.

После этого, API будет доступно по базовому URL: http://localhost:44049. \
Использование API аналогично разделам удаленного запуска. 
