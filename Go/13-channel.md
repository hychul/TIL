# Go의 channel

채널<sup>channel</sup>은 Golang에서 기본 타입으로 제공하는 정해진 버퍼 크기를 갖는 스레드 세이프<sup>Thread Safe<sup>한 큐이다. 스레드 세이프하기 때문에 다른 스레드에서 채널에 들어있는 데이터에 접근 하더라도 데드락과 같은 스레드 이슈가 발생하지 않는다.

채널을 초기화 할때 채널의 버퍼의 크기를 지정할 수 있다. 지정하지 않는 경우엔 버퍼 크기가 0을 가지게 된다.

```go
var a chan Type
b := make(chan Type)
c := make(chan Type, 10)
```

채널은 기본적으로 양방향 채널을 의미한다. 때문에 `<-` 연산자를 통해서 아이템을 보낼 수도, 받을 수도 있다. 하지만 단방향 채널도 제공한다.

```go
c := make(chan int)
var cr <-chan int = c // receive only
var cs chan<- int = c // send only
```

## 채널의 동작

채널에 아이템을 넣고 빼는 것은 `<-` 연산자를 통해 동작한다.

```go
c := make(chan int)

c <- 1
a := <-c

fmt.Println(a)
```
```terminal
fatal error: all goroutines are asleep - deadlock!
```

하지만 위의 경우, 결과값을 보면 알 수 있듯이 모든 고루틴이 멈췄다는 에러를 발생시킨다. 그 이유는 채널의 버퍼 사이즈는 채널에서 아이템을 보낼 때 사용할 버퍼의 갯수를 의미하기 때문이다. 보내는 쪽에서 남은 채널의 버퍼의 갯수가 0이라면 받는 쪽에서 아이템을 받아서 처리할 때까지 블로킹<sup>Blocking</sup>된다.[^go-cahnnel-blocking]

[^go-cahnnel-blocking]:https://golang.org/src/runtime/chan.go #219

 그리고 이렇게 채널에서 블로킹이 된 경우, 다른 고루틴이 실행된다. 다른 고루틴을 실행하려 했지만 실행 가능한 고루틴이 존재하지 않았기 때문에 위의 에러 메시지가 하나의 고루틴이 아닌, 모든 고루틴이 멈췄다고 표시를 하는 것이다.

```go
c := make(chan int)

go func() {
    fmt.Println("send")
    c <- 1 // send 1
    fmt.Println("send end")
}()

time.Sleep(1000 * time.Millisecond)
fmt.Println("receive")
a := <-c                        // receive 1
fmt.Println("receive end :", a) // 1

fmt.Scanln()
```
```terminal
send
receive
receive end : 1
send end
```

채널의 버퍼 크기가 0이기 때문에 "send end"가 "receive"보다 나중에 출력되게 된다. 버퍼 크기를 0보다 크게 초기화를 한다면 그 순서는 반대가 될 것이다.

## 1:1 단방향 채널

앞서 채널의 남은 버퍼 갯수가 0인 경우 보내는 쪽에서 받는 쪽에서 처리할 때까지 대기한다고 하였다. 때문에 1:1로 채널을 통해 아이템을 전달한다면 보내는 수와 받는 수를 맞춰야한다. 이를 받는 쪽에서 보내는 쪽에서 몇번을 보내는지 모르더라도 항상 동작하게 만들면 다음과 같다.

```go
c := func() <-chan int {
    c:= make(chan int)
    go func() {
        defer close(c)
        c <- 1
        c <- 2
        c <- 3
    }()
    return c
}()

for item := range c {
    fmt.Println(item)
}
```
```terminal
1
2
3
```

받는 쪽에서 단방향 채널을 전달받도록 제한하여 받는 쪽에서 채널에 아이템을 보내 고루틴이 멈추게 되는 것을 예방하였다.

## 버퍼가 있는 채널

앞서 살펴본 버퍼의 크기가 0인, 즉, 버퍼가 없는 채널은 물리적으로 동시레 수행되지 않는 경우가 많다. 받는 쪽이 준비가 되지 않아 아이템을 받지 못하는 경우 보내는 쪽에서 받는 쪽이 아이템을 받아갈 때까지 블로킹이되기 때문이다.

받는 쪽이 준비가 되지 않더라도 보내는 쪽에서 미리 아이템을 큐에 넣어놓기 위해선 채널을 만들때 버퍼 크기를 지정해주면 된다. 

```go
c := make(chan Type, 10)
```

버퍼가 있는 채널은 보내는 쪽과 받는 쪽이 같은 고루틴이어도 동작하고 다른 고루틴에서 동작할 때 속도의 차이가 존재하더라도 계속 동작할 수 있기 때문에 성능향상을 기대할 수 있다.

하지만 논리적 오류등으로 버퍼가 가득 차게 되면 보내는 쪽이 영원히 멈춰버리기 때문에 버퍼 크기만 늘리는 것은 능사가 아니다.

## 닫힌 채널

닫힌 채널에서 아이텐을 가져오려고 하면 채널을 기다리지 않고 아이템 타입의 default 값을 반환한다.

```go
c := make(chan int, 1)

go func() {
    defer close(c)
}()

time.Sleep(1000 * time.Millisecond)
a := <-c

fmt.Println(a) // 0
```

## 파이프 라인 채널

단방향 채널을 응용하면 하나의 아이템이 여러 단계를 순서대로 거쳐 동작하도록 만들 수 있다.

```go
type IntPipe func(<-chan int) <-chan int

func main() {
	in := make(chan int)
	plusTwo := Chain(PlusOne, PlusTen)
	out := plusTwo(in)

	fmt.Println("start with :", 1)
	in <- 1

	fmt.Println("finish with :", <-out)
}

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			num++
			fmt.Println("plus one :", num)
			out <- num
		}
	}()
	return out
}

func PlusTen(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			num += 10
			fmt.Println("plus ten :", num)
			out <- num
		}
	}()
	return out
}
```
```terminal
start with : 1
plus one : 2
plus ten : 12
finish with : 12
```

## 팬아웃

팬아웃<sup>Fan-out</sup>은 하나의 출력이 여러 입력으로 들어가는 것을 말한다. 1:1 단방향 채널을 통해서 구현된 파이프 라인은 한번에 하나씩 순서대로 아이템이 처리되기 중간 과정이 보내는 쪽보다 시간이 오래걸린다면 중간 과정이 처리될 때까지 대기하게 된다. 때문에 중간 과정을 병렬로 처리 할 수 있도록 채널 하나를 여럿에게 공유하면 효율적으로 아이템을 처리할 수 있다.

여러 고루틴에서 한 채널의 아이템을 받아가려 할 때, 보내는 쪽에서 아이템을 보내면 여로 고루틴 중 하나의 고루틴에만 아이템이 전달된고 나머지는 대기하게 된다.

```go
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		fmt.Println("set receiver :", i)
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println("received :", i, n)
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("send :", i)
		c <- i
	}

	time.Sleep(1000 * time.Millisecond)

	close(c)
}
```
```terminal
set receiver : 0
set receiver : 1
set receiver : 2
set receiver : 3
set receiver : 4
send : 0
send : 1
send : 2
send : 3
send : 4
received : 0 0
received : 3 3
received : 4 4
received : 2 1
received : 1 2
```

순차적으로 받는 쪽을 설정하고 아이템을 보냈지만, 받아서 처리하는 리시버의 인덱스는 순차적이지 않은 걸 알 수 있다. 아이템을 랜덤한 고루틴에서 받아 처리하기 때문이다.

## 팬인

팬인<sup>Fan-in</sup>은 팬아웃과 반대로 여러 출력이 하나의 입력으로 들어가는 것을 의미한다.

```go
func main() {
	in1 := make(chan int)
	in2 := make(chan int)
	in3 := make(chan int)

	out := FanIn(in1, in2, in3)

	in1 <- 1
	in2 <- 2
	in3 <- 3

	go func() {
		for num := range out {
			fmt.Println(num)
		}
	}()
}

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(ins))

	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```
```terminal
1
2
3
end
```

## 분산처리

팬아웃과 팬인을 조합하여 사용하면 분산처리를 할 수 있다. 앞서 파이프 라인에서 정의한 `IntPipe`를 파라메터로 받아 n개로 분산 처리하는 함수는 다음과 같이 구현된다.

```go
// IntPipe와 FanIn 함수 생략

func main() {
	in := make(chan int)
	distributPlus := Distribution(Peek, 5)
	out := distributPlus(in)

	in <- 1
	in <- 2
	in <- 3
	in <- 4

	go func() {
		for num := range out {
			fmt.Println("receive :", num)
		}
	}()

	time.Sleep(1000 * time.Millisecond)
}

func Distribution(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

func Peek(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			fmt.Println("peek :", num)
			out <- num
		}
	}()
	return out
}
```
```terminal
peek : 1
peek : 3
receive : 1
receive : 3
peek : 4
receive : 4
peek : 2
receive : 2
```