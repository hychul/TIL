# Go의 select

셀렉트<sup>select</sup>를 이용하면 동시에 여러 채널과 통신할 수 있다. 셀렉트의 형태는 swith 문과 비슷하게 여러 채널 중 아이템을 받은 채널에 대한 케이스가 수행되고, 모든 채널에서 보낸 아이템이 없는 경우 default 문이 수행돤다.

```go
select {
    case v := <-c1:
        fmt.Println("channel 1")
    case v := <-c2:
        fmt.Println("channel 2")
    default:
        fmt.Println("default")
}
```

셀렉트를 사용하면 손쉽게 채널에 대한 팬인을 구현할 수 있다.

```go
select {
    case v := <-c1: out <- v
    case v := <-c2: out <- v
    case v := <-c3: out <- v
}
```

채널을 그냥 사용한다면 채널에서 아이템을 전달하지 않는다면 받는 쪽에선 계속 대기를 해야하지만 셀렉트를 사용한다면 기다리지 않고 무시하고 다른 처리를 할 수 있다.

```go
select {
    case v := <-c:
        fmt.Println("receive item")
    default:
        fmt.Println("item is not ready. skipping...)
}
```

또한 `time.After()` 함수를 이용하여 지정한 시간 만큼만 기다리도록 할 수 있다.

```go
select {
    case v := <-c:
        fmt.Pintln("receive item")
    case <-time.After(5 * time.Second)
        fmt.Println("no item received for 5 second.")
        return
}
```
