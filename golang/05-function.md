# Go의 함수

go의 함수는 다음의 형식에 따라 정의된다.

```
func [함수명]([파라메터 이름] [파라메터 타입]) [리턴 타입] {
    [구현부]
}
```

두 정수를 받아 덧샘을 하는 함수를 구현하면 다음과 같다. 파라메터의 타입이 같은 경우엔 `add1()` 함수처럼 파라메터를 표현도 가능하다.

```go
func add(x int, y int) int {
    return x + y
}

func add1(x, y int) int {
    return x + y
}
```


Golang에선 리턴값을 여러개로 설정할 수 있다.

```go
func main() {
    a, b := change(0, 1)

    fmt.Println(a, b) // print: 1 0
}

func change(x, y int) (int, int) {
    return y, x
}
```

주의할 점은 golang에서 파라메터는 항상 값으로서 전달된다. 레퍼런스 참조를 사용하기 위해선 포인터를 명시적으로 사용해야한다.