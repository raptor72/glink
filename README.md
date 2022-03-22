# glink
HTML link parser 

Модуль для получения всех ссылок с HTML страницы.

## Использование

Импортируйте модуль "github.com/raptor72/glink"

Для получения всех ссылок с конкретной страницы передайте в функцию glink.Parse обьект типа io.Reader, например для ответа на http запрос:

```golang
resp, err := http.Get(requested_url)

links, err := glink.Parse(resp.Body)
if err != nil {
	...
}
```

Ссылки будут возвращены в виде структуры Link:

```golang
type Link struct {
    Href, Text string
}
```

В модуле представлено несколько примеров HTML страниц в виде файлов. Чтобы посмотреть их нужно указать конкретный html файл, например:

```bash
go run ./exhamples/main.go --filename=./exhamples/ex2.html
[{Href:https://www.twitter.com/joncalhoun Text:Check me out on twitter} {Href:https://github.com/gophercises Text:Gophercises is on Github !}]
```

Для получения справки выполните:

```bash
go run ./exhamples/main.go --help
```

Для выполнения тестов выполните:
```bash
go test
PASS
ok      github.com/raptor72/glink       0.006s
```

Идея взята из курса Джона Колтона: https://courses.calhoun.io/lessons/les_goph_16

Гитхаб оригинального проекта: https://github.com/gophercises/link
