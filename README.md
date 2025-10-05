# Практическое задание 3

## ЭФМО-02-25 Мишкин Артём Дмитриевич
---
# Информация о проекте
https - выполняет функции листа с задачами, которые можно добавлять, фильтровать и выводить полный список.
## Цели занятия
 - Освоить базовую работу со стандартной библиотекой net/http без сторонних фреймворков.
 - Научиться поднимать HTTP-сервер, настраивать маршрутизацию через http.ServeMux.
 - Научиться обрабатывать параметры запроса (query, path), тело запроса (JSON/form-data) и формировать корректные ответы (код статуса, заголовки, JSON).
 - Научиться базовому логированию запросов и обработке ошибок.

## Планируемые результаты
1.	Запускать простой HTTP-сервер и настраивать маршруты.
2.	Реализовывать обработчики (http.HandlerFunc) для GET/POST.
3.	Читать query-параметры, path-параметры (простым способом), тело запроса, декодировать JSON.
4.	Возвращать JSON-ответы с корректными кодами статуса и заголовками.
5.	Писать простую прослойку логирования (middleware-подход на http.Handler).
6.	Тестировать API с помощью curl/PowerShell/HTTPie/Postman

## Файловая структура проекта:

<img width="217" height="247" alt="image" src="https://github.com/user-attachments/assets/346dddf6-cb11-4f44-96be-822327e8b024" />

# Примечания по конфигурации и требования

Для запуска требуется:

Go: версия 1.21 и выше

# Команды запуска/сборки
Для запуска http нужно выполнить 4 шага:
## 1) Клонировать данный репозиторий в удобную для вас папку:
```Powershell
git clone https://github.com/ArtyomMishkin/http
```
## 2) Перейти в папку http:
```Powershell
cd http
```
## 3) Загрузка зависимостей:
```Powershell
go mod tidy
```
## 4) Команда запуска
```Powershell
go run .cmd/server/main.go
```

# Команда сборки
Для сборки бинарника и запуска .exe файла используются данные программы

```Powershell
go build -o http.exe ./cmd/server
.\http.exe
```
# Проверка работоспособности веб-сервиса

## Проверка "/health"

<img width="974" height="300" alt="image" src="https://github.com/user-attachments/assets/8ca9044a-04f4-4f0f-98cb-dddedd6f9a7d" />

## Проверка "создание задания"

<img width="974" height="262" alt="image" src="https://github.com/user-attachments/assets/57daf29f-213d-4c95-b848-3fa99f359fd2" />

## Проверка "вывод задания по ID"

<img width="974" height="242" alt="image" src="https://github.com/user-attachments/assets/04252f76-c90c-44c4-bfd4-d139b3eb217b" />

## Проверка "Вывод списка"

<img width="974" height="294" alt="image" src="https://github.com/user-attachments/assets/27fa5fff-2122-41fa-8357-17af26dc3ef6" />

## Проверка "Фильтра по слову"

<img width="974" height="267" alt="image" src="https://github.com/user-attachments/assets/c5ee0f42-f55d-4bd5-9a08-199169c6e863" />

## Задание на "звёздочку"

make.ps1  в данном проекте - файл со скриптом с целями run, build, test

```Powershell
.\make.ps1 run # равен go run ./cmd/server 
.\make.ps1 build # равен go build -o bin/server.exe ./cmd/server 
.\make.ps1 test # равен go test -v ./internal/api/...
```
ДЛЯ РАБОТЫ МОЖЕТ ПОНАДОБИТЬСЯ ВЫДАТЬ РАЗРЕШЕНИЕ КОМАНДОЙ

```Powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```
Проверка работы скрипта

### test

<img width="382" height="198" alt="image" src="https://github.com/user-attachments/assets/b873a4bd-eb8d-4a59-90c8-a1c864bc0454" />


### build

<img width="248" height="31" alt="image" src="https://github.com/user-attachments/assets/dbb07532-a531-47a8-88e6-05e1bd36b8e6" />

### run

<img width="301" height="31" alt="image" src="https://github.com/user-attachments/assets/22c1af3f-e529-4f91-a0bd-11f722aa28fd" />

Также был создан фалй handlers_test.go для проверки хэндлеров

- TestCreateTask - успешное создание задачи
- TestCreateTask_EmptyTitle - создание с пустым заголовком
- TestListTasks - получение списка задач
- TestListTasks_WithFilter - фильтрация задач
- TestCreateTask_InvalidLength - валидация длины заголовка

<img width="337" height="184" alt="image" src="https://github.com/user-attachments/assets/3eff87dc-9029-4445-9066-0b5628af26f8" />
