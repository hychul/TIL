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
    Directory:  "htmls",
    Extensions: []string{".html", ".tmpl"},
})
```

`redner`를 사용하면 레이아웃을 지정하여 템플릿을 재사용할 수 있다. 홈페이지를 구성하는 html 에서 페이지마다 중복을 사용되는 부분을 지정하고 그 페이지 안에서 다른 템플릿을 가져올때 처럼 `{{ yield }}` 구문을 삽입하여 사용된다. yield를 통해 다른 템플릿으로 구성할 수 있게 된다.

```html
<html>
<head>
<title>Index</title>
</head>
<body>
Hello Web 
{{ yield }}
</body>
</html>
```

```go
rd = render.New(render.Options{
    Directory:  "htmls",
    Extensions: []string{".html", ".tmpl"},
    Layout:     "index",
})
```

또한 레이아웃 안에서 구성되는 이너 템플릿 마다 다른 값을 주고 싶을 수 있다. 예를 들면 html의 title이 이너 템플릿 마다 다른 이름으로 하고 싶은 경우 `partial` 을 사용한다. 다음과 같이 partial 뒤에 quotes를 붙인 문자열로 키값을 지정하고 템플릿이 보관된 디렉토리 안에 '[키]-[이너 템플릿 파일]'로 템플릿을 생성하면 이너 템플릿에 해당하는 템플릿이 `partial`에 적용된다.

```html
<html>
<head>
<title>{{ partial "title" }}</title>
</head>
<body>
Hello Web 
{{ yield }}
</body>
</html>
```