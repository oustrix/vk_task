# Профильное задание Go-разработчик

## Задание
```
Создать простого чат бота на базе ВКонтакте API 
(допускается как альтернатива Telegram Bot API)

Функциональные требования:
Бот должен уметь отправлять приветственное сообщение, иметь 
минимум 4 кнопки на первом слое и минимум по 2 во втором.

Технические требования:
Нужно реализовать на языке Golang с минимальным использованием 
готовых пакетов
```

## Использованные внешние библиотеки
-  [Cleanenv](https://github.com/ilyakaznacheev/cleanenv) - чтение конфигурации из файла config.yml и переменных окружения
- [Godotenv](https://github.com/joho/godotenv) - чтение переменных окружения из файла .env
- [Query](https://github.com/google/go-querystring) - парсинг структур в query-параметры
- [Mapstructure](https://github.com/mitchellh/mapstructure) - преобразование ответов API в структуры

## Запуск бота
1. Загрузить проект
```bash
git clone https://github.com/oustrix/vk_task
```
2. Создать файл .env в корне проекта и заполнить переменные окружения (см [example.env](https://github.com/oustrix/vk_task/blob/main/example.env))
3. Собрать docker-образ проекта
```bash
docker build -t vk_task .
```
4. Запустить docker-образ
```bash
docker run vk_task
```