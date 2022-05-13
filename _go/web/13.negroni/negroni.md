Negroni는 크게 recovery, logging 그리고 static file 전송 기능을 제공하는 패키지다. 

`http.FileServer(http.Dir("public"))` 구문과 동일하게 동작하는데, 아래와 같이 코드를 작성하는 경우 'public' 디렉토리의 'index.html' 파일을 root api로 제공한다. mux를 `UseHandler()` 메서드로 랩핑하는 경우 mux에서 root에 대한 api를 정의 했더라도 오버라이딩이 되기 때문에 주의해야한다.

```go
mux := pat.New()

mux.Get("/", getIndexHandler)
mux.Get("/users", getUserHandler)
mux.Post("/users", createUserHandler)

n := negroni.Classic()
n.UseHandler(mux)

http.ListenAndServe(":3000", n)
```

그리고 negroni 패키지의 핸들러를 사용하는 경우 기본적으로 로그 기능을 제공하여 api 호출에 대한 로그를 보여준다.