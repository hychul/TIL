# Go 조건 연산자

Golang 의 if 연산자는 참 혹은 거짓을 확인하기 위한 조건 절에 괄호를 사용하지 않는다.

```go
var a = 3
if a > 0 {
    fmt.Println("Positive number")
} else if a < 0 {
    fmt.Println("Negative number")
} else {
    fmt.Println("Zero")
} // Positive number
```

또한 `if` 절에 조건문뿐만 아니라 전처리문을 명시할 수 있다. 전처리문에 정의된 변수의 life scope은 전처리문이 명시된 `if` 의 scope와 일치한다.

```go
if a := 1; true {
	fmt.Println(a)
} // 1
```

# Go의 switch 문

Golang의 switch는 break를 명시적으로 설정하지 않더라도 각 case 별로 break가 설정된다. Case문은 코드로 작성된 위에서 아래 순서로 조건을 확인하며 동작한다.

```go
var a = 0
switch {
case a > 0:
    fmt.Println("Positive number")
case a < 0:
    fmt.Println("Negative number")
case a == 0:
    fmt.Println("Zero")
} // Zero

var b = "+"
switch b {
case "+":
    fmt.Println("Plus")
case "-":
    fmt.Println("Minus")
case "*":
    fmt.Println("Multiply")
case "/":
    fmt.Println("Division")
} // Plus
```

# Go의 반복문

Golang에서 반복문은 `for` 문만 제공된다. 하지만 `for` 문을 다른 언어의 `while` 처럼 사용할 수 있다.

```go
var i = 0
for i < 10 {
    fmt.Println(i)
    i++
} // print 0 to 9
```

`for` 문의 조건문을 명시하지 않는다면 `for true {...}`과 같이 무한 루프로 동작한다. 때문에 이 경우엔 `break` 키워드를 잘 사용해야한다.

```go
var i = 0
for {
    if (i == 10)
        break
    fmt.Println(i)
} // print 0 to 9
```

물론 익숙한 for 문처럼 `for 전처리문; 조건문; 후처리문`의 형식으로 사용할 수도 있다.

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
} // print 0 to 9
```