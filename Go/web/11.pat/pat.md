Go 에선 mux를 API 메서드 별로 제공할 수 있도록, gorilla에서 제공하는 Mux 패키지로 사용할 수 있지만 Restful API를 위한 메서드를 `string` 파라메터로 전달해야하기 때문에 불편하고 코드도 아름답지 않다.

이를 대체하기 위해  gorilla에서 pat이라는 패키지를 만들었는데, 이 패키지는 Restful API 에 대한 핸들러 등록 메서드를 메서드 형태로 제공한다. pat은 mux 패키지와 같이 gorilla에서 만들었고 mux 패키지를 확장했기 때문에 기존에 사용하던 mux와 관련된 구현을 사용할 수 있다.

```terminal
$ go get github.com/gorilla/pat
```