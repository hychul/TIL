# Go

golang은 구글에서 만든 언어오 2009년에 처음 대중에게 발표되었다. 현재는 오픈 소스 프로젝트로 관리되고 있다.

C언어 이전 모델인 B언어를 만든 켄톰슨과 UTF-8을 만든 롭 파이크가 개발에 참여했다.

golang은 C언어와 닮았고 기본적으로 UTF-8을 사용한다. System 프로그래밍을 하기 위해 처음 개발된 언어이기 떄문에 굉장히 빠른 성능을 보여준다. (C언어보다는 느리다.)

# Go 설치

http://golang.org 에서 받을 수 있다.

설치 프로그램으로 golang을 설치한 설치 확인을 위해서 terminal을 열어 `go version`을 입력하면 설치된 golang에 대한 정보가 출력된다.

```terminal
go version
go version go1.14 windows/amd64
```

그리고 `go env`를 터미널에 쳤을때 GOPATH에 대한 디렉토리를 자동적으로 생성해주진 않기 때문에 이에 해당하는 디렉토리를 생성해주어야 한다.

# Hello World

vscode에 확장자가 '.go' 인 파일을 생성하면 golang에 관련된 플러그인들을 추천해준다 이를 모두 설치하면 intelli sence와 같은 간편한 기능들을 vscode를 통해서 제공 받을 수 있다.

다른 프로그래밍 언어들이 그러하듯이 "Hello World"를 출력하는 단순한 프로그램을 golang으로 만들어 보자.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

```

확장자가 '.go'인 파일을 생성한 뒤 위와 같이 코드를 작성하고 vscode의 menu bar/Debug/Run을 해주면 작성한 코드가 실행된다.

terminal에서 실행하길 원한다면 `go run [go format file]` 명령어를 입력하면 된다. 해당 명령어는 빌드를 하지 않고 동적 언어와 같이 golang을 실행하게 된다.

만약 기계어로의 빌드를 원한다면 소스 코드가 있는 디렉토리로 이동하여 `go build` 명령어 혹은 `go build [go format file]` 명령어를 실행하면 실행 가능한 파일을 생성한다.