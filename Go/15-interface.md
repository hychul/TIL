# Go의 OOP

Go는 객체지향<sup>Object Oriented Programming</sup>을 완전하게 지원하지 않는다. 하지만 Go에서 지원하는 구조체를 통해 기존의 객체지향의 언어보다 더 유연하게 오브젝트를 설계하고 사용할 수 있다.

## 다형성

관용적인 개념의 다형성은 같은 코드를 통해서 다른 행위를 하는 것을 의미한다. 일반적인 언어에선 같은 함수를 통해서 다른 로직을 수행하는 것을 의미한다.

객체지향 프로그래밍의 특징 중 하나인 다형성은 상속을 통해 기능을 확장하거나 변경하는 것을 가능하게 한다. Go는 상속을 지원하지 않지만 덕 타이핑<sup>Duck Typing</sup>을 지원한다.

덕 타이핑이란 어떤 물체가 오리인지 확인하기 위해선 오리인지 검사하는 것이 아닌 오리처럼 울고 걷고 행동하는지 검증하라는 개념으로 시작한다. Go는 덕 타이핑을 인터페이스<sup>interface</sup>를 통해서 지원한다.

### 인터페이스

Go의 인터페이스를 통해서 다형성을 위한 덕 타이핑을 지원한다고 했는데, 이 때문에 Go의 인터페이스는 다른 객체지향형 언어와 달리 명시적으로 인터페이스를 구현한다는 표시를 할 필요 없이 인터페이스 내의 함수를 구현하는 것만으로 그 인터페이스를 구현하는 것으로 인식된다.

```go
type Duck interface {
    Cry() string
    Walk()
}
```

위의 `Duck` 인터페이스는 선언하면, `Cry()`와 `Walk()` 함수를 갖는 구조체는 `Duck` 인터페이스 타입으로 사용할 수 있게 된다.

이런 Go의 인터페이스의 장점은 다른 객체지향 언어들이 인터페이스가 그 인터페이스를 구현하는 쪽에서 명시적으로 선언하는 것과 달리 암묵적으로 인터페이스를 상속하도록 하기 때문에 다른 사람이 만든 패키지의 구조체를 사용할 때 내가 만든 인터페이스 타입으로 덮어씌어 사용할 수 있는 유연한 확장성을 지원한다.

<!-- 예시가 있으면 좋을듯 -->

## 상속

객체지향에서 상속은 어떤 클래스의 구현들을 재사용하기 위해 사용된다. 때문에 인터페이스가 아닌 클래스나 추상 클래스를 상속하여 그 기능을 변형, 확장하여 사용한다.

Go의 경우 클래스 상속, 즉, 구조체 상속을 지원하지 않지만 인터페이스와 내장 구조체를 통해 기존의 구조체를 확장하여 사용할 수 있다.

```go
type Duck {}

func (d Duck) Cry() string {
    return "quack"
}

func (d Duck) Walk() {
    fmt.Println("walk")
}

type Chicken struct { Duck }

func main() {
	c := Chicken{}

	fmt.Println(c.Cry())
	c.Walk()
}
```
```terminal
quack
walk
```

### 메서드 추가

내장 구조체를 통해 새로운 구조체를 만들고 해당 구조체에 새로운 메서드를 추가하면 된다.

```go
...
type Chicken struct { Duck }

func (c Chicken) Fly() {
    fmt.Println("fly")
}

func main() {
    c := Chicken{}

	fmt.Println(c.Cry())
    c.Walk()
    c.Fly()
}
```
```terminal
quack
walk
fly
```

### 오버라이딩

내장 구조체를 통해 내장된 구조체의 메서드 중 오버라이딩하려는 메서드를 새로 작성하면된다.


```go
...
type Chicken struct { Duck }

func (c Chicken) Cry() string {
    return "cock-a-doodle-doo"
}

func main() {
    c := Chicken{}

	fmt.Println(c.Cry())
    c.Walk()
}
```
```terminal
cock-a-doodle-doo
walk
```

## 캡슐화

객체안에 있는 정보를 외부로 부터 숨겨 보호하기 위해 사용하는 것이 캡슐화이다. 일반적인 언어들에서 `public` 이나 `private` 과 같은 접근자를 통해 이런 객체 내부로의 접근을 제한하지만 Go의 경우 이러한 접근자가 존재하지 않는다.

대신 01장에서 언급한 것과 같이 대문자 혹은 소문자를 통해 다른 패키지로부터의 접근을 제한한다.

1.4부턴 패키지의 접근 제한[^1.4-internal-package]도 추가되었다. 'internal' 디렉토리를 생성하고 그 안에 package를 개발하면, 'internal' 디렉토리가 위치하는 바로 위 루트의 모든 하위 package들에서만 internal의 public 함수를 사용할 수 있게 된다. 이러한 rule은 go command에 의해 강제된다.

이를 통해 각 패키지 내부에서 공통적으로 사용되는 public 함수들이 여러 패키지를 사용하여 개발을 할때에도 충돌하지 않고 동작될 수 있게 되었다.

[^1.4-internal-package]:[1.4 internal package Google Reference](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit)