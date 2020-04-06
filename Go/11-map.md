# Go의 map

맵<sup>Map</sup>은 리스트(Slice)와 더불어 일반적이 언어들에서 가장 널리 쓰이는 자료구조이다. 때문에 Go도 맵을 기본적으로 제공한다.

```go
var m1 map[string]string
var m2 map[int]string
var m3 map[Key]Value
```

맵을 위 처럼 선언만 하는 경우 `nil` 값을 갖기 때문에 초기화를 해주어야 한다.

```go
m1 := make(map[string]string)
m2 := map[string]string{}
m3 := map[string]string{"key1": "value1", "key2": "value2"}
```

## Map에 아이템 추가와 가져오기

맵에 새로운 아이템을 추가하고 읽는 방법은 슬라이스와 비슷한데, 인덱스 대신 키값을 이용한다는 것이 다르다. 만약 추가하지 않은 키값에 대한 값을 얻으려고 한다면 deafult값이 출력된다.

```go
m := make(map[string]int)

m["first"] = 1

fmt.Println(m) // map[first:1]
fmt.Println(m["first"]) // 1
fmt.Println(m["second"]) // 0
```

이렇게 된 경우 한가지 문제가 발생하는데, 만약 어떤 키에 대한 밸류로 해당 타입의 default 값과 같은 값을 설정한다면 나중에 밸류를 받아왔을 때 이것이 맵에 존재하지 않는 값인지 아니면 존재하지만 default와 같은 값을 설정한 건지 알 수가 없다.

때문에 Go의 맵은 키값으로 밸류를 받아올때 다중 대입을 통해 맵에 존재하는지게 대한 정보를 함께 반환한다.

```go
m := make(map[string]int)

m["first"] = 1
m["second"] = 0

v1, ok1 := m["first"]
v2, ok2 := m["second"]
v3, ok3 := m["third"]

fmt.Println(v1, ok1) // 1 true
fmt.Println(v2, ok2) // 0 ture
fmt.Println(v3, ok3) // 0 false
```

## Map에 아이템 삭제

맵으로 부터 아이템을 삭제하고 싶은 경우엔 `delete(m map[Type]Type1, key Type)` 함수를 사용한다.

```go
m := make(map[string]int)

m["first"] = 1
delete(m, "first")

v1, ok1 := m["first"]
v2, ok2 := m["second"]

fmt.Println(v1, ok1) // 0 false
fmt.Println(v2, ok2) // 0 false
```

## Map의 순회

맵의 추가된 모든 키, 밸류 값을 순회하고 싶을 때 사용한다. Go에서 순회할 때 사용하는 `for`를 사용하여 표현한다. 키값에 대해서 순차적으로 반환되지 않고 랜덤한 순서에 따라 키, 밸류가 반환된다.

```go
m := map[string]int{"first": 1, "second": 2}

for key, value := range m {
	fmt.Print(key, " ", value, ", ")
} // first 1, second 2,
```

## Map으로 만든 Set

셑<sup>Set</sup>은 맵의 키처럼 자료구조 상에 하나만 존재할 수 있는 집합을 의미한다. 그렇기 때문에 Java에서도 `Set`의 내부 구현을 보면 `Map`을 사용하는 것을 확인 할 수 있다. 때문에 Go에선 셑을 직접적으로 제공하진 않지만 맵을 통해서 간단하게 구현할 수 있다.

```go
s := make(map[int]bool)

s[0] = true

fmt.Println(s[0], s[1]) // true false
```

위와 같이 간단한 모양으로 셑을 만들어 사용할 수 있다. 하지만 셑에 존재하는지 확인해야 하는 키값의 대한 메모리 뿐만 아니라 `bool` 타입에 대한 메모리가 추가적으로 사용된다는 문제가 있다. 이는 빈 구조체를 사용하여 벨류에 대한 메모리를 사용하지 않게 할 수 있다. 이는 앞서 설명했듯이 맵은 키값이 맵에 존재하는지에 대한 정보를 함께 반환하는 덕분에 가능한 일이다.

```go
s := make(map[int]struct{})

s[1] = struct{}{}

_, ok1 := s[1] 
_, ok2 := s[2]

fmt.Println(ok1, ok2) // true false
```

위의 코드에서 셑에 키를 추가할때 `struct{}{}`의 중괄호가 두번 붙는 이유는 `struct{}`가 빈 구조체를 나타내는 타입이기 때문이다. `strcut{}` 타입에 대해서 초기화를 할때 붙이는 `{}`를 한번 더 붙여 초기화 된 빈 구조체를 밸류로 추가하는 것이다.