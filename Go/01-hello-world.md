# Hello World 코드 까보기

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

```

## package

패키지의 선언을 의미한다. 모든 golang으로 만들어지는 코드에 필요하다. 자바의 그것과 비슷하다.

프로그램의 시작점을 포함한 패키지는 항상 main으로 작성되어야 한다.

그자체로 실행이 필요한 프로그램은 main 패키지로 시작되어야 하지만 라이브러리로 작성된 golang 코드는 main을 가질 필요가 없다.

> 기능 묶음
> 라이브러리 : 기능을 묶어놓은 것
> 모듈 : 라이브러리 보다 좀 더 특정 영역에 포커싱되어 있음
> 프레임 워크 : 기능을 묶어놓은 것, 절차가 필요함
> 엔진 : 기능 뿐만 아니라 다른 프로그램이나 툴의 묶음

## import

c 언어의 `#include ~` 구문 혹은 자바의 그것과 같이 다른 패키지를 가져와 사용하고 싶을 때 사용한다.

fmt 패키지는 입출력을 담당하는 golang의 표준 패키지 중 하나이다.

하나의 경우에는 필요 없지만, import를 통해 여러개의 패키지를 가져오고 싶은 경우에는 braket을 사용하면된다.

```go
import (
    "fmt"
    "sort"
    ...
)
```

## func

function의 약어로 함수를 의미한다. (Kotlin에서 fun을 함수로 나타내는데 이보다는 더 세련됐다 :P)

## fmt.Println("Hello World")

fmt 패키지에 정의된 `Println()` 함수를 사용하여 "Hello World" 문자열을 화면에 출력한다.

언어마다 코드 컨벤션이 존재하기 때문에 왜 fmt 패키지에서 사용한 함수는 대문자로 시작하고 `main()` 함수는 소문자로 시작하는지 의문이 생길 수 있다.

golang function naming convention에 따르면 다른 패키지에 노출이 되는 함수는 대문자로 시작하고 그렇지 않은 함수는 소문자로 작성된다고 한다.