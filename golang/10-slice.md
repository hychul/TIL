# Go의 Silce

앞서 06에서 언급했던 길이를 지정한 배열을 정적배열<sup>Fixed Size Array</sup>이라한다. 쉽게 말해 정적배열은 길이가 바뀌지 않는 배열을 의미한다.

```go
var arr [10]int
```

동적배열<sup>Dynamic Array</sup>은 정적배열과 달리 배열의 길이가 바뀔 수 있는 배열을 의미한다. Golang에선 크기를 지정하지 않은 `[]Type` 으로 타입을 나타낸다.

```go
var arr []int
```

동적배열에서 배열의 길이가 늘어나야하는 경우 메모리 상에서 새로운 늘어난 크기 만큼의 배열의 메모리를 새로 할당받아 기존 배열의 값을을 복사하고 동적배열의 포인터를 새로운 메모리를 가리키도록 한다. 고정배열을 포인트하고 있는 것이다.

고정 배열을 선언할 떄 처럼 중괄호를 사용하여 초기값을 설정할 수 있다. 또한 Golang의 내장 키워드인 `make()`를 사용하여 동적배열 변수를 생성할 수 있다.

```go
arr1 := []int{1, 2, 3, 4, 5}
arr2 := make([]int, 5) // length가 5인 동적배열
arr3 := make([]int, 0, 8) // length는 0이고 capacity는 8인 동적배열
```

## length와 capacity

capacity는 동적배열이 포인트 하고 있는 고정배열의 메모리 상의 크기를 의미한다. 즉, 동적배열에서 새로운 아이템을 추가할때 capacity를 넘어가지만 않으면 포인트하고 있는 정적배열이 새로 메모리를 할당받고 값들을 복사하지 않아도 된다.

length는 정적배열과 비슷하게 현재 사용하고 있는 배열의 길이를 의미하지만 정적배열은 처음 선언될 떄 설정된 고정된 길이를 갖는다는 것과 동적 배열은 동적인 길이를 갖는다는 차이가 있다.

```go
arr := make([]int, 2, 4)

fmt.Println(len(arr), cap(arr))
```
```terminal
2, 4
```

정리하자면 동적 배열은 capacity, length 그리고 정적 배열을 멤버변수로 갖는 하나의 구조체라고 할 수 있다.

## append

동적배열은 길이가 변하기 때문에 새로운 아이템을 배열에 추가할 수 있는데 이때 `append(slice []Type, elems... Type)` 함수를 사용한다.

```go
arr := make([]int, 0, 5)

fmt.Println(len(arr), cap(arr))

arr = append(arr, 1)

fmt.Println(len(arr), cap(arr))
```
```terminal
0, 5
1, 5
```

`append()` 함수는 슬라이스를 파라메터와 리턴 타입으로 사용하는 것을 볼 수 있는데, 파라메터로 전달된 슬라이스와 리턴된 슬라이스는 다른 슬라이스이지만, 새로운 아이템을 append할 때 capacity가 충분한 경우 같은 정적 배열 메모리 공간을 가리킬 수 있다.

```go
arr1 := make([]int, 2, 3)
arr2 := append(arr1, 3)

fmt.Println(arr1)
fmt.Println(arr2)

fmt.Printf("%p %p\n", arr1, arr2)
fmt.Printf("%p %p\n", &arr1, &arr2)

arr3 := append(arr2, 4)

arr1[0] = 1
arr2[1] = 2

fmt.Println(arr1)
fmt.Println(arr2)
fmt.Println(arr3)

fmt.Printf("%p %p %p\n", arr1, arr2, arr3)
fmt.Printf("%p %p %p\n", &arr1, &arr2, &arr3)
```
```terminal
[0 0]
[0 0 3]
0xc000010440 0xc000010440
0xc0000044c0 0xc0000044e0
[1 2]
[1 2 3]
[0 0 3 4]
0xc000010440 0xc000010440 0xc00000a300
0xc0000044c0 0xc0000044e0 0xc000004580
```

때문에 기존 슬라이스와 append한 슬라이스를 모두 사용해야하는 경우, `copy(dst, source []Type)` 함수를 통해 다른 메모리를 가리키게 한 후 사용하는게 안전하다.

```go
arr1 := make([]int, 2, 3)
arr2 := make([]int, len(arr1), cap(arr1))
copy(arr2, arr1)

arr2 = append(arr2, 3)

arr1[0] = 1
arr2[1] = 2

fmt.Println(arr1) 
fmt.Println(arr2)

fmt.Printf("%p %p\n", arr1, arr2) 
fmt.Printf("%p %p\n", &arr1, &arr2) 
```
```terminal
[1 0]
[0 2 3]
0xc000010440 0xc000010460
0xc0000044c0 0xc0000044e0
```

## Slice 슬라이스하여 사용하기

슬라이스는 이름처럼 일부를 잘라내어 사용할 수 있다. 배열의 대괄호와 콜론을 사용하여 구간을 나타내는데, `[시작 인덱스:끝 인덱스]`로 표시하고 시작 인덱스의 경우 inclusive, 끝 인덱스는 exclusive 하다.

```go
arr1 := []int{1, 2, 3, 4, 5}
arr2 := arr1[1:4]

fmt.Println(arr1)
fmt.Println(arr2)
```
```terminal
[1 2 3 4 5]
[2 3 4]
```

시작 인데스 혹은 끝 인덱스 중 하나를 생략할 수 있는데, 이 경우엔 시작이 생략된경우 0을 끝이 생략된 경우 배열의 길이가 주어진 것과 같다.

```go
arr1 := []int{1, 2, 3, 4, 5}
arr2 := arr1[:4]
arr3 := arr1[1:]

fmt.Println(arr1) 
fmt.Println(arr2) 
fmt.Println(arr3)
```
```terminal
[1 2 3 4 5]
[1 2 3 4]
[2 3 4 5]
```

주의할 점은 슬라이스의 잘라낸 경우 동일한 메모리 영역을 공유한다는 점이다. 슬라이스를 잘랐을때 새로운 메모리에 복사하여 사용하는 것이 아니라 기존 슬라이스의 메모리의 일부분을 가리켜 사용하도록 한다.

```go
arr1 := []int{1, 2, 3, 4, 5}
arr2 := arr1[1:4]

arr1[1] = 0
arr2[3] = 0

fmt.Println(arr1) 
fmt.Println(arr2)
```
```terminal
[1 0 3 0 5]
[0 3 0]
```
