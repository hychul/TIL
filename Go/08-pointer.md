#Go의 포인터

포인터는 메모리 상에서 변수의 주소값을 나타낸다. C 나 C++ 같은 언어들은 포인터를 명시적으로 사용할 수 있도록 하여 유저들에게 자유도를 주었지만 메모리에 직접 접근하기 때문에 시스템 자체에 위험할 수 있다.
때문에 비교적 최근 언어라고 할 수 있는 C#이나 Java에선 포인터의 개념을 명시적으로 사용할 수 없고 레퍼런스 타입의 개념으로 묵시적으로 사용된다.

Golang에선 다른 언어에서처럼 묵시적으로 포인터를 사용하는 것이 아닌 직접적으로 포인터를 이해하고 사용하도록 하면서도, 포인터의 연산이나 캐스팅을 막아 직관적으로 포인터를 사용하면서도 안전성을 높혔다.

컴퓨터의 메모리 주소는 8byte를 갖기 때문에 포인터 변수는 8byte의 크기를 갖는다.

```go
a := 5
var p *int

p = &a

fmt.Println(a, *p, p) // 5 5 0xc000126090(a의 메모리의 주소값)
```

포인터는 메모리에서의 주소값을 통해 동일한 변수에 접근 할 수 있기 때문에 함수에서 파라메터가 값을 복사하여 전달하더라도 함수 외부의 변수에 접근이 가능하다.

```go
func main() {
	var a int
	a = 0

	Increase(a)

	fmt.Println(a) // 0

	IncreateWithPointer(&a)

	fmt.Println(a) // 1
}

func Increase(a int) {
	a++
}

func IncreateWithPointer(p *int) {
	*p++
}
```

이를 통해 07장에서 하지 못했던 객체의 멤버 함수를 통한 멤버 변수 값의 설정도 가능하다.

```go
type Person struct {
    name string
    age int
}

func (p *Person) SetValue(name string, age int) {
    p.name = name
    p.age = age
}

func main() {
    var p Person

    p.SetValue("A", 1)

    fmt.Println(p) // {A 1}
}
```