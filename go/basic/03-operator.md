# Go 연산자

<table class="table">
<tbody><tr>
  <th>연산자</th>
  <th>기능</th>
  <th>설명</th>
</tr>
<tr>
  <td>=</td>
  <td>대입</td>
  <td>
변수나 상수에 값을 대입합니다. 변수는 변수끼리 대입할 수 있습니다.

```go
var a int = 1
var b int = 2
var c int = b
const d string = "Hello, world!"
```

Go에선 다중 대입도 지원합니다.

```go
var a1, a2, a3, a4 = 1, 2, 3, 4

var b1, b2, b3, b4 = aFunctionThatReturnFourValue()
```

  </td>
</tr>
<tr>
  <td>:=</td>
  <td>변수 선언 및 대입</td>
  <td>변수를 선언하는 동시에 값을 대입합니다.

```go
a := 1 // int
b := 3.5 // float64
c := "Hello, world!" // string
```

  </td>
</tr>
<tr>
  <td>+</td>
  <td>덧셈</td>
  <td>
  두 값을 더합니다. 사용할 수 있는 자료형은 정수, 실수, 복소수, 문자열입니다.

```go
a := 1 + 2 // 3
b := 2 + 3 // 5
c := a + b // 8
d := "Hello, " + "world!" // Hello, world!
```

  </td>
</tr>
<tr>
  <td>-</td>
  <td>뺄셈</td>
  <td>두 값의 차이를 구합니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다.

```go
a := 3 - 2 // 1
b := 4 - 5 // -1
c := a - b // 2
```

  </td>
</tr>
<tr>
  <td>*</td>
  <td>곱셈</td>
  <td>두 값을 곱합니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다.

```go
a := 2 * 3 // 6
b := 9 * 21 // 189
c := a * b // 1134
```

  </td>
</tr>
<tr>
  <td>/</td>
  <td>나눗셈</td>
  <td>두 값을 나눕니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다.

```go
a := 2 * 3 // 6
b := 9 * 21 // 189
c := a * b // 1134
```

  </td>
</tr>
<tr>
  <td>%</td>
  <td>나머지</td>
  <td>두 값을 나눈 뒤 나머지를 구합니다. 사용할 수 있는 자료형은 정수입니다.

```go
a := 5 % 2 // 1
```

  </td>
</tr>
<tr>
  <td>+=</td>
  <td>덧셈 후 대입</td>
  <td>현재 변수와 값을 더한 다음 다시 변수에 대입합니다. 문자열은 현재 변수에 문자열을 붙인 다음 다시 변수에 대입합니다.

```go
a := 5 // 5
a += 2 // 7

b := "Hello, " // Hello,
b += "world!" // Hello, world!
```

  </td>
</tr>
<tr>
  <td>-=</td>
  <td>뺄셈 후 대입</td>
  <td>현재 변수에서 값을 뺀 다음 다시 변수에 대입합니다.

```go
a := 5 // 5
a -= 2 // 3
```

  </td>
</tr>
<tr>
  <td>*=</td>
  <td>곱셈 후 대입</td>
  <td>현재 변수와 값을 곱한 다음 다시 변수에 대입합니다.

```go
a := 5 // 5
a *= 2 // 10
```

  </td>
</tr>
<tr>
  <td>/=</td>
  <td>나눗셈 후 대입</td>
  <td>현재 변수를 값으로 나눈 다음 다시 변수에 대입합니다.

```go
a := 5 // 5
a /= 2 // 2
```

  </td>
</tr>
<tr>
  <td>%=</td>
  <td>나머지를 구한 후 대입</td>
  <td>현재 변수와 값의 나머지를 구한 다음 다시 변수에 대입합니다.

```go
a := 5 // 5
a %= 2 // 1
```

  </td>
</tr>
<tr>
  <td>&amp;</td>
  <td>AND 비트 연산</td>
  <td>두 값을 비트 단위로 AND 연산을 합니다. 사용할 수 있는 자료형은 정수입니다.

```go
a := 3 // 00000011
b := 2 // 00000010
c := a & b // 00000010
```

  </td>
</tr>
<tr>
  <td>|</td>
  <td>OR 비트 연산</td>
  <td>두 값을 비트 단위로 OR 연산을 합니다. 사용할 수 있는 자료형은 정수입니다.

```go
a := 3 // 00000011
b := 2 // 00000010
c := a | b // 00000011
```

  </td>
</tr>
<tr>
  <td>^</td>
  <td>XOR 비트 연산(다항)</td>
  <td>두 값을 비트 단위로 XOR 연산을 합니다. 사용할 수 있는 자료형은 정수입니다.

```go
a := 3 // 00000011
b := 2 // 00000010
c := a ^ b // 00000001
```

  </td>
</tr>
<tr>
  <td>&amp;^</td>
  <td>AND NOT 비트 연산</td>
  <td>두 값을 비트 단위로 AND NOT 연산을 합니다. 즉 다음과 같이 특정 비트를 끕니다(Bit clear). 사용할 수 있는 자료형은 정수입니다.

```go
a := 255 // 11111111
b := 68 // 01000100
c := a &^ b // 10111011
```

  </td>
</tr>
<tr>
  <td>&amp;=</td>
  <td>AND 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 AND 연산한 다음 다시 변수에 대입합니다.

```go
a := 3 // 00000011
b := 2 // 00000010
a &= b // 00000010
```

  </td>
</tr>
<tr>
  <td>|=</td>
  <td>OR 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 OR 연산한 다음 다시 변수에 대입합니다.

```go
a := 3 // 00000011
b := 2 // 00000010
a |= b // 00000011
```

  </td>
</tr>
<tr>
  <td>^=</td>
  <td>XOR 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 XOR 연산한 다음 다시 변수에 대입합니다.

```go
a := 3 // 00000011
b := 2 // 00000010
a ^= b // 00000001
```

  </td>
</tr>
<tr>
  <td>&amp;^=</td>
  <td>AND NOT 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 AND NOT 연산한 다음 다시 변수에 대입합니다. 이 연산자는 특정 플래그를 끌 때 주로 사용합니다.

```go
a := 255 // 11111111
b := 68 // 01000100
a &^= b // 10111011
```

  </td>
</tr>
<tr>
  <td>&lt;&lt;</td>
  <td>비트를 왼쪽으로 이동</td>
  <td>현재 값의 비트를 특정 횟수만큼 왼쪽으로 이동합니다. 사용할 수 있는 자료형은 정수입니다.

```go
a := 7 // 00000111
b := a << 2 // 00011100
```

  </td>
</tr>
<tr>
  <td>&gt;&gt;</td>
  <td>비트를 오른쪽으로 이동</td>
  <td>현재 값의 비트를 특정 횟수만큼 오른쪽으로 이동합니다. 사용할 수 있는 자료형은 정수입니다.

```go
a := 112 // 01110000
b := a >> 3 // 00001110
```

  </td>
</tr>
<tr>
  <td>&lt;&lt;=</td>
  <td>비트를 왼쪽으로 이동 후 대입</td>
  <td>현재 변수를 특정 횟수만큼 왼쪽으로 이동한 다음 다시 변수에 대입합니다.

```go
a := 7 // 00000111
a <<= 2 // 00011100
```

  </td>
</tr>
<tr>
  <td>&gt;&gt;=</td>
  <td>비트를 오른쪽으로 이동 후 대입</td>
  <td>현재 변수를 특정 횟수만큼 오른쪽으로 이동한 다음 다시 변수에 대입합니다.

```go
a := 112 // 01110000
a >>= 3 // 00001110
```

  </td>
</tr>
<tr>
  <td>^</td>
  <td>비트 반전(단항)</td>
  <td>비트를 반전시킵니다(Bitwise complement, 1의 보수). 즉 0은 1로 1은 0으로 바꿉니다.

```go
var a uint8 = 3 // 00000011
b := ^a // 11111100
```

  </td>
</tr>
<tr>
  <td>+</td>
  <td>양수 부호(단항)</td>
  <td>값에 양수 부호를 붙입니다.

```go
a := 3
b := -2
c := +a // 3
d := +b // -2
```

  </td>
</tr>
<tr>
  <td>-</td>
  <td>음수 부호 (단항)</td>
  <td>값에 음수 부호를 붙입니다.

```go
a := 3
b := -2
c := -a // -3
d := -b // 2
```

  </td>
</tr>
<tr>
  <td>==</td>
  <td>같다</td>
  <td>두 값이 같은지 비교합니다.
  <ul>
  <li>실수형은 값을 연산한 뒤에는 오차가 발생하므로 ==로 비교할 때 주의해야 합니다.</li>
  <li>문자열은 내용이 같은지 비교합니다.</li>
  <li>배열은 요소의 내용이 모두 같은지 비교합니다.</li>
  <li>슬라이스와 맵은 배열과는 달리 내용을 비교할 수 없고, 메모리에 할당되어 있는지 확인합니다.</li>
  <li>포인터는 주소가 같은지 비교합니다.</li>
  </ul>

```go
fmt.Println(1 == 2) // false
fmt.Println(3.5 == 5.5) // false
fmt.Println("Hello" == "world") // false
fmt.Println("Hello" == "Hello") // true

a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
fmt.Println(a == b) // true

c := []int{1, 2, 3}
fmt.Println(c == nil) // false

// check memory allocation
d := map[string]int{"Hello": 1}
fmt.Println(d == nil) // false

// compare memory address
e := 1
f := 1
var p1 int = &e
var p2 int = &f
fmt.Println(p1 == p2) // false
```

  </td>
</tr>
<tr>
  <td>!=</td>
  <td>같지 않다</td>
  <td>두 값이 다른지 비교합니다.

```go
fmt.Println(1 != 2) // true
fmt.Println(3.5 != 5.5) // true
fmt.Println("Hello" != "world") // true
fmt.Println("Hello" != "Hello") // false


a := [3]int{1, 2, 3}
b := [3]int{3, 2, 1}
fmt.Println(a != b) // true

c := []int{1, 2, 3}
fmt.Println(c != nil) // true

// checking memory allocation
d := map[string]int{"Hello": 1}
fmt.Println(d != nil) // true

// comparing memeory address
e := 1
f := 1
var p1 int = &e
var p2 int = &f
fmt.Println(p1 != p2) // true
```

  </td>
</tr>
<tr>
  <td>&lt;</td>
  <td>작다</td>
  <td>앞의 값이 작은지 비교합니다. 문자열은 ASCII 코드 값을 기준으로 판단합니다. 또한, 첫 글자가 같다면 그 다음 글자부터 차례대로 비교하여 최종 값을 구합니다.

```go
fmt.Println(1 < 2) // true
fmt.Println(3.5 < 5.5) // true
```

  </td>
</tr>
<tr>
  <td>&lt;=</td>
  <td>작거나 같다</td>
  <td>앞의 값이 작거나 같은지 비교합니다.

```go
fmt.Println(2 <= 2) // true
fmt.Println(3.5 <= 5.5) // true
```

  </td>
</tr>
<tr>
  <td>&gt;</td>
  <td>크다</td>
  <td>앞의 값이 큰지 비교합니다.

```go
fmt.Println(2 > 1) // true
fmt.Println(5.5 > 3.5) // true
```

  </td>
</tr>
<tr>
  <td>&gt;=</td>
  <td>크거나 같다</td>
  <td>앞의 값이 크거나 같은지 비교합니다.

```go
fmt.Println(2 >= 2) // true
fmt.Println(5.5 >= 3.5) // true
```

  </td>
</tr>
<tr>
  <td>&amp;&amp;</td>
  <td>AND 논리 연산</td>
  <td>두 불 값이 모두 참인지 확인합니다.

```go
fmt.Println(true && true) // true
fmt.Println(true && false) // false
fmt.Println(false && false) // false
```

  </td>
</tr>
<tr>
  <td>||</td>
  <td>OR 논리 연산</td>
  <td>두 불 값 중 한 개라도 참인지 확인합니다.

```go
fmt.Println(true || true) // true
fmt.Println(true || false) // true
fmt.Println(false || false) // false
```

  </td>
</tr>
<tr>
  <td>!</td>
  <td>NOT 논리 연산</td>
  <td>불값을 반대로 연산합니다.

```go
fmt.Println(!true) // false
fmt.Println(!false) // true
```

  </td>
</tr>
<tr>
  <td>&amp;</td>
  <td>참조(레퍼런스) 연산</td>
  <td>현재 변수의 메모리 주소를 구합니다.

```go
a := 1
b := &a // a의 메모리 주소를 b에 대입
fmt.Println(b) // 0xc0820062d0 (메모리 주소)
```

  </td>
</tr>
<tr>
  <td>*</td>
  <td>역참조 연산</td>
  <td>현재 포인터 변수에 저장된 메모리에 접근하여 값을 가져오거나 저장합니다.

```go
a := new(int)
*a = 1 // a에 저장된 메모리에 접근하여 1을 저장
fmt.Println(*a) // a에 저장된 메모리에 접근하여 값을 가져옴
```
```terminal
1
```

  </td>
</tr>
<tr>
  <td>&lt;-</td>
  <td>채널 수신 연산</td>
  <td>채널에 값을 보내거나 값을 가져옵니다.

```go
c := make(chan int)
go func() {
c <- 1 // 채널 c에 1을 보냄
}()
a := <-c // 채널 c에서 값을 가져와서 a에 대입

fmt.Println(a)
```
```terminal
1
```

  </td>
</tr>
<tr>
  <td>++</td>
  <td>증가</td>
  <td>변수의 값을 1 증가시킵니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다. 

```go
a := 1
a++ // 2

b := 1.5
b++ // 2.5

c := 1 + 2i
c++ // 2+2i
```

Go 언어에서는 ++ 연산자를 사용한 뒤 값을 대입할 수 없고, 변수 뒤에서만 사용할 수 있습니다. 따라서 ++ 연산자는 단독으로 사용하거나 <b>if</b> 조건문, <b>for</b> 반복문 안에서 주로 사용합니다.

```go
a := 1
b := a++ // compile error
c := ++a // compile error
++a // compile error
```

  </td>
</tr>
<tr>
  <td>--</td>
  <td>감소</td>
  <td>변수의 값을 1 감소시킵니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다. 

```go
a := 1
a-- // 0

b := 1.5
b-- // 0.5

c := 1 + 2i
c-- // 0+2i
```

Go 언어에서는 -- 연산자를 사용한 뒤 값을 대입할 수 없고, 변수 뒤에서만 사용할 수 있습니다. 따라서 -- 연산자는 단독으로 사용하거나 <b>if</b> 조건문, <b>for</b> 반복문 안에서 주로 사용합니다.

```go
a := 1
b := a-- // compile error
c := --a // compile error
--a // compile error
```

  </td>
</tr>
</tbody></table>