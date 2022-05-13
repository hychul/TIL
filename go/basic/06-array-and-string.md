# Go의 배열

```go
// 타입의 default 값을 사용한 배열의 선언
var a [10]int
a := [10]int{}

// 초기화를 사용한 배열의 선언
a := [10]int{0,1,2,3,4,5,6,7,8,9}
```

배열의 길이

```go
a := [10]int{0,1,2}
fmt.Println(len(a)) // 10
```

# Go의 문자열

Go의 `string`도 배열이다. 하지만 주의해야할 점은 Go에선 ASCII 대신 UTF-8을 기본으로 사용하기 때문에 다른 언어와 같이 string의 각 문자들이 항상 1byte로 표현되지 않는다. <sup>영어, 숫자의 경우엔 UTF-8에서도 1byte로 표현되기 때문에 1byte로 봐도 무방하다.</sup>

실제로 `string`을 배열로 접근하여 출력을 할 경우엔 각 UTF-8 코드에 해당하는 값이 출력된다.

```go
str := "Hello world"

for i := 0; i < len(str); i++ {
	fmt.Print(str[i], ",")
} // 72,101,108,108,111,32,119,111,114,108,100,
```

`string`을 배열로 사용하여 각 문자를 출력하기 위해선 `string`으로의 변환이 이뤄져야한다.

```go
for i := 0; i < len(str); i++ {
	fmt.Print(string(str[i]), ",")
}
```

UTF-8을 사용하기 때문에 `string`을 배열 처럼 사용할때 UTF-8에서 1byte가 아닌 문자를 사용할 땐 배열로서의 접근에서 주의해야한다. UTF-8에선 모든 문자가 1byte를 갖지 않기 때문이다.

만약 "Hello 월드"의 문자열을 string으로 캐스팅하여 출력하는 경우 원하는 것 처럼 각 문자가 출력되지 않는다.

```go
str := "Hello 월드"

for i := 0; i < len(str); i++ {
    fmt.Print(string(str[i]), ",")
} // H,e,l,l,o, ,ì,,,ë,,,
```

`len(str)`의 경우 위의 코드에서 예상한것처럼 8의 값을 리턴하지 않는다. `len()` 은 문자열을 파라메터로 받았을 때 byte 단위로 크기를 구하기 때문이다. 

Go에선 문자를 `rune`이라는 1~3byte 사이의 크기를 갖는 타입으로 나타내기 때문에 문자의 갯수를 구하기 위해선 `[]rune` 배열로 변환해야한다.

```go
str := "Hello 월드"
runeArray := []rune(str)

for i := 0; i < len(runeArray); i++ {
	fmt.Print(string(runeArray[i]), ",")
} // H,e,l,l,o, ,월,드,
```

Go의 대입 연산자 '=' 가 다중 대입을 지원하기 때문에 배열을 통해 다른 언어와 다른 방식으로 구현할 수 있는 것들이 존재한다.

하나의 예로 배열의 순서를 바꾸는 것을 좀더 적은 코드와 변수로 작성할 수 있다.

일반적인 대입 연산자가 하나의 대입만을 지원할때 배열의 순서를 거꾸로 하기 위해선 다음과 같이 변수를 하나 선언하여 임시로 배열의 값을 저장하고 넘겨주는 방식으로 코드를 작성해야한다.

```go
arr := []int{1, 2, 3, 4, 5}
var tmp int

for i := 0; i < len(arr)/2; i++ {
	tmp = arr[len(arr)-1-i]
	arr[len(arr)-1-i] = arr[i]
	arr[i] = tmp
}
```

하지만 이 코드를 다중 대입을 사용하여 작성하면 다음과 같이 작성할 수 있다.

```go
arr := []int(1, 2, 3, 4, 5)

for i := 0; i < len(arr)/2; i++ {
	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
} 
```