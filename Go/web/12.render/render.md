render는 Go web에서 JSON, XML, binary data 그리고 HTML template과 같은 response를 간편하게 설정할 수 있는 오픈소스 패키지다.

```terminal
$ go get github.com/unrolled/render
```

`render`를 사용해서 `HTML()` 메서드를 사용하는 경우 기본적으로 root 디렉토리에서 './templates' 폴더 안에 존재하는 파일들 중 '.tmpl' 확장자를 갖는 템플릿만 인식한다. 확장자가 '.html'인 파일도 사용하고 싶다면, `render` 를 `New()` 메서드로 선언할 때, `render.Options` struct를 통해서 확장자를 인식하도록 설정해야 한다.

```go
rd = render.New(render.Options{
    Extensions: []string{".html", ".tmpl"},
})
```

확장자와 마찬가지로 기본 설정된 디렉토리를 사용하지 않고 직접 이름을 지정한 디렉토리를 사용하기 위해선 `render.Options`에 Directory를 사용한다.

```go
rd = render.New(render.Options{
    Directory: "htmls",
    Extensions: []string{".html", ".tmpl"},
})
```