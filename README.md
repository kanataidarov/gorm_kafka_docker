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
curl --location '<Базовый URL>/applications' \
--request POST \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Канат Айдаров",
    "email": "kanataidarov@yahoo.com",
    "position": "Junior Go Developer"
}'
```
Пример отправки через Postman: \
![image](https://github.com/user-attachments/assets/a50972b9-3ab0-4974-8828-7c43a1b83749)

### Получение статистики
На `<Базовый URL>/application` отправить методом HTTP GET. \
Пример отправки через curl:
```shell
curl --location --request GET '<Базовый URL>/applications'
```

Пример отправки через Postman: \
![image](https://github.com/user-attachments/assets/184f367d-1a3a-4208-993e-7d0a8cdcae64)

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
