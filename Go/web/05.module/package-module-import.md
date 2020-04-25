# Module

기본적으로 $GOPATH 디렉토리 하위에서 Go의 개발을 하는 것을 권장한다. 하지만 프로젝트가 커지거나 하는 등의 상황에 놓일 땐 $GOPATH 외부의 디렉토리에서 개발을 해야할 때가 온다.

문제가 되는 것은 전체 프로젝트를 모듈화 하여 개발을 한다고 하더라도, 각각의 모듈의 맡은 개발자의 개발환경이 일정하지 않을 수 있기 때문에 모듈의 실행 환경 등의 차이가 생기고 이는 전체 서비스의 안정성에 영향을 줄 수 있다.



Draft

04 장의 소스코드를 다른 디렉토리에 가져와 실행을 하려고 하면 main.go 에서 import한 myapp 패키지가 GOPATH의 절대경로에 존재하지 않기 때문에 실행시에 해당 패키지를 찾을 수 없다는 오류가 발생한다.

```termianl
$ go build
main.go:6:2: cannot find package "github.com/hychul/myapp" in any of:
        c:\go\src\github.com\hychul\myapp (from $GOROOT)
        C:\Users\hychul\go\src\github.com\hychul\myapp (from $GOPATH)
```

1. 모듈로 지정하지 않고서는 패키지를 상대 경로로 지정할 수 있다.
2. 모듈로 지정한 경우엔 패키지를 상대 경로로 지정할 수 없다. 대신 모듈을 설정한 패키지 경로를 루트로 사용하여야 한다.

모듈을 지정하는 방법

```terminal
$ go mod init example.com/module
```

go.mod 파일이 생성된다.

```mod
module example.com/module

go 1.14
```

명령어를 통해 지정한 'example.com/module' 이라는 경로가 루트 경로로 사용된다. 때문에 해당 go.mod 파일이 존재하는 디렉토리 하위에 디렉토리를 만드는 경우 루트 경로의 상대 경로로 패키지를 import 할 수 있다.

> Go 1.11 이후의 모듈 지원 : https://blog.golang.org/using-go-modules