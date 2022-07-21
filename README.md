# level0

> Структура проекта

1. docker-compose.yml - поднимает базу данных и ants-streaming
2. main.go - точка входа приложения
3. model - модели
4. service - содержит классы хранилища (БД и Кэш)
5. handler - обработчики запросов
6. пакет test - main.go для отправки данных в nats, mem_test.go юнит тест кеша

> Запуск 

1. docker-compose up
2. go run main.go
3. go run test/main.go

> Использование

Доступны 2 роута:

1. localhost:8080/order/{id} - получить заказ по ид
2. localhost:8080/orders - все заказы