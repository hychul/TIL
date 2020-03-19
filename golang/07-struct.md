#Go struct

Golang에서 구조체<sup>struct</sup>은 여러가지 변수들을 하나의 타입으로 만들어 준다.

```go
type Person struct {
    name string
    age int
}
```

그리고 이렇게 유저를 통해 만들어진 구조체는 하나의 타입으로 사용될 수 있다. 

때문에 golang의 구조체는 하나의 객체로써 사용된다. 다른 언어의 생성자처럼 구조체의 멤버 변수들을 선언시에 설정할 수 있는데, 생성자를 굳이 명시하지 않아도 다음과 같이 초기값을 지정할 수 있다. 변수들의 값을 지정하지 않는 경우엔 default 값이 할당된다.

```go
func main() {
    var p Person // name : "", age: 0
    p1 := Person{} // name: "", age: 0
    p2 := Person{"Tom", 21} // name: "Tom", age: 21
    p3 := Person{name:"David", age:23} // name: "Davic", age: 23
    p4 := Person{name:"Alex"} // name: "Alex", age: 0
}
```

구조체는 멤버 변수 이외의 멤버 함수를 설정할 수 있다.

다른 언어와 다른 점은 멤버 함수를 구조체 내에 작성하는 것이 아닌 외부에 작성한다는 것이다.

```go
type Person struct {
    name string
    age int
}

func (p Person) PrintName() {
    fmt.Println(p.name)
}
```

위에서 작성한 멤버 함수는 다음의 함수와 동일하게 동작한다.

```go
func PrintName(p Person) {
    fmt.Println(p.name)
}
```

멤버 함수를 작성할 때 주의해야할 점이 있는데 05장에서 함수를 설명할 때 언급했던, 함수의 파라메터는 값을 복사하여 전달된다는 것이다. 때문에 다음의 코드는 의도한대로 값을 제대로 설정하지 못한다.

```go
type Person struct {
	name string
	age  int
}

func (p Person) SetValue(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	fmt.Println("Hello world")

	p := Person{name: "A", age: 0}

	p.SetValue("B", 1)

	fmt.Println(p) // {A 0}
}
```

앞서 언급한 것처럼 위에서 구조체의 값을 설정하는 멤버 함수인 `SetValue()`는 다음과 동일한 함수이다.

```go
func SetValue(p Person, name string, age int) {
    p.name = name
    p.age = age
}
```

그리고 golang에서 함수의 파라메터는 모두 값을 복사해서 가져오기 때문에 `SetValue()`의 컨텍스트 안의 `p Person`과 이를 호출하는 `main()` 함수의 `p Person`은 메모리 상에서 다른 주소값을 가진다.
 때문에 `SetValue()` 함수 안에서 Person 객체의 멤버 변수 값을 변경하더라도 다시 `main()` 의 Person 객체에는 영향을 주지 못한다.

이를 의도한 것처럼 구현하기 위해선 포인터를 사용해야한다.