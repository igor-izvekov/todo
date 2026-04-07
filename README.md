# TODO
Мини-проект для лабораторной работы 9-10 по дисциплине «Методы отладки и тестирования ПО».

## Идея проекта
Todo приложение - удобный способ организовывать список дел, которые необходимо сделать

## Почему этот проект подходит для лабораторной
Проект удобно использовать для демонстрации сразу нескольких ролей в команду:
* backend-разработчик — реализует бизнес-логику, маршруты, БД и валидацию;
* frontend-разработчик — делает интерфейс, формы и отображение задач;
* тестировщик — готовит тест-кейсы, чек-листы, баг-репорты и проверяет исправления.

## Технологии:
* Golang 1.25
* Gin
* testify
* SQLite
* HTML/CSS/JS

## Структура проекта:
```
.
├── README.md
└── src
    ├── Makefile
    ├── cmd
    │   └── main.go
    ├── frontend
    │   ├── index.html
    │   └── todo.html
    ├── go.mod
    ├── go.sum
    ├── pkg
    │   ├── auth
    │   │   └── auth.go
    │   ├── database
    │   │   └── connect.go
    │   ├── handlers
    │   │   ├── complete_task.go
    │   │   ├── complete_task_test.go
    │   │   ├── create_task.go
    │   │   ├── create_task_test.go
    │   │   ├── delete_task.go
    │   │   ├── delete_task_test.go
    │   │   ├── get_task.go
    │   │   ├── get_task_test.go
    │   │   ├── get_tasks.go
    │   │   ├── get_tasks_test.go
    │   │   ├── task_test.go
    │   │   ├── update_task.go
    │   │   └── update_task_test.go
    │   ├── migrations
    │   │   └── auto_migrate.go
    │   └── models
    │       └── models.go
    └── todo.db
```

## Как запустить проект
```
cd src
make start
```

## Участники
Извеков Игорь БИВТ-22-СП-2
Ермилов Василий БИВТ-22-СП-2
