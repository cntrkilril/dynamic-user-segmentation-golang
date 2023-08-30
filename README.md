# Сервис динамического сегментирования пользователей

## Для запуска приложения

```sh
make build && make run
```

## Описание

### Swagger

[Ссылка на Swagger](http://localhost:8080/swagger)

### Методы

---

#### Сегменты

- ``POST`` ``body`` ``/api/segment/create`` ``Создание сегмента (с автоматическим добавлением к определенному проценту)``

| Имя                  | Тип    | Описание                                 | Дополнительное       |
|----------------------|--------|------------------------------------------|----------------------|
| slug                 | string | slug сегмента                            | required, >0         |
| autoAddToUserPercent | int    | процент пользователей для автодобавления | required, >=0, <=100 |

**Request**

```
{
    "slug": "TAG",
    "autoAddToUserPercent": 100
}
```

**Response**

```
{
    "data": {
        "segment": "TAG"
    },
    "errors": []
}
```

- ``DELETE`` ``body`` ``/api/segment/delete`` ``Удаление сегмента``

| Имя  | Тип    | Описание      | Дополнительное |
|------|--------|---------------|----------------|
| slug | string | slug сегмента | required, >0   |

**Request**

```
{
    "slug": "TAG",
}
```

**Response**

```
{
    "data": {},
    "errors": []
}
```

#### Сегменты пользователей

- ``POST`` ``body`` ``/api/segment/add-segments-to-user`` ``Добавление пользователей в список сегментов``

| Имя      | Тип      | Описание         | Дополнительное |
|----------|----------|------------------|----------------|
| userID   | int      | ID пользователя  | required, >0   |
| segments | []string | список сегментов | required       |

**Request**

```
{
    "userID": 2,
    "segments": ["VOICE", "TAG", "UI"]
}
```

**Response**

```
{
    "data": {
        "userID": 2,
        "segments": [
            "TAG" - P.S. сегменты, который были добавлены удачно
        ]
    },
    "errors": [
        "сегмент VOICE уже добавлен",
        "сегмент UI не найден"
    ]
}
```

- ``DELETE`` ``body`` ``/api/segment/delete-segments-to-user`` ``Удаление пользователей из списка сегментов``

| Имя      | Тип      | Описание         | Дополнительное |
|----------|----------|------------------|----------------|
| userID   | int      | ID пользователя  | required, >0   |
| segments | []string | список сегментов | required       |

**Request**

```
{
    "userID": 2,
    "segments": ["RENT", "UI", "1"]
}
```

**Response**

```
{
    "data": {
        "userID": 2,
        "segments": [
            "RENT",  - P.S. сегменты, который были удалены удачно
            "UI"
        ]
    },
    "errors": [
        "сегмент 1 не найден"
    ]
}
```

- ``GET`` ``params`` ``/api/segment/get-segments-by-user-id/{userID}`` ``Получение сегментов пользователя``

| Имя    | Тип | Описание        | Дополнительное |
|--------|-----|-----------------|----------------|
| userID | int | ID пользователя | required, >0   |

**Request**

```
 GET /api/segment/get-segments-by-user-id/2
```

**Response**

```
{
    "data": {
        "segments": [
            "VOICE",
            "TAG"
        ],
        "userID": 2
    },
    "errors": []
}
```

#### История

- ``GET`` ``query`` ``/api/segment/history/get-by-user-id-and-year-month`` ``Получение ссылки на csv-файл с историей операций над сегментами пользователя за определенный месяц года``

| Имя    | Тип | Описание        | Дополнительное          |
|--------|-----|-----------------|-------------------------|
| userID | int | ID пользователя | required, >0            |
| year   | int | год             | required, >=2000,<=3000 |
| month  | int | номер месяца    | required, >=1,<=12      |

**Request**

```
 GET api/segment/history/get-by-user-id-and-year-month?userID=2&year=2023&month=8
```

**Response**

```
{
    "data": {
        "url": "http://0.0.0.0:8080/2.csv"
    },
    "errors": []
}
```
