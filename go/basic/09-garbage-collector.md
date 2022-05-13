# Go의 Garbage Collector

가비지 콜렉터는 일반적으로 GC<sup>Garbage Collector<sup>라고 부르면 메모리에서 사용되지 않는 가비지 값들을 다시 사용할 수 있는 빈 메모리 공간으로 정리한다.

## GC의 탄생 배경

C언어에선 메모리를 크게 스택 메모리와 힙 메모리로 구분하여 사용한다. 스택의 경우 일반적으로 선언한 변수들을 저장하는 공간이고, 힙 메모리는 `void malloc(size_t size)` 함수를 통해 메모리를 할당하여 사용하는 변수들이 저장되는 공간이다.

스택 메모리 공간의 경우 함수의 영역안에서만 존재하는 변수들을 저장하기 때문에 함수가 `return`과 함께 끝나는 경우 메모리를 다시 반환한다.

하지만 힙 메모리의 경우 유저가 명시적으로 메모리를 할당해서 사용하기 때문에 프로그램에서 언제 이 메모리가 반환되어야 하는지 알 수 없어, 유저가 직접 할당한 메모리를 `void free(void *)` 함수를 통해 반환을 해야한다.

때문에 유저가 메모리의 반환을 제대로 하지 않는다면 가비지 메모리가 많아져 프로그램이 제대로 동작하지 않게 된다. 이렇게 유저가 만들어낸 버그로 인해 메모리의 관리가 제대로 안되는 상황을 피하기 위해 언어 레벨에서 root로 부터 참조가 존재하지 않는, 즉, 레퍼런스 카운터가 0인 메모리를 자동으로 반환할 수 있도록 GC가 고안되었다.

## GC의 한계

처음 자바의 GC는 힙 영역에 존재하는 변수들을 전수 조사하기 때문에 임의의 시간에 프로세스를 모두 멈추고 GC 프로세스가 수행되었다. 하지만 현재는 GC가 전수조사를 위해 오랜 기간 동안 프로세스를 멈추지 않고 잠깐 잠깐 눈치채기 어려운 만큼씩 수행된다.

하지만 GC가 있다고 메모리 릭<sup>Memory Leak</sup>이 발생하지 않는 것이 아니다. GC는 아무도 참조하지 않는 변수인 Dangling Pointer를 관리하지만, 특정 객체에서 사용하지 않을 변수들을 들고있는다면 GC는 참조하고 있는 변수라고 생각하고 해당 메모리를 반환하지 않는다.

## Go의 GC

앞서 언급한 Java와 비교했을 때 Go의 GC는 약간의 차이를 갖는다.

### 압축<sup>Compaction</sup>

GC는 크게 정적 유형과 동적 유형으로 나뉜다.

정적 유형의 GC는 heap 내의 변수들을 재배치하지 않는다. Go는 Mark & Sweep을 통한 GC를 제공하는데 이는 정적 유형 GC에 속한다. heap 내의 변수들을 재배치하지 않는 경우 메모리의 할당과 해제(GC)가 반복될 경우 메모리가 파편화 되어 성능이 악화되는 문제가 있다.

동적 유형의 GC는 heap 내의 변수를 재배치하여 heap을 압축한다. HotSpot VM의 GC에서 사용되는 copy GC의 경우 동적 유형의 GC이다. 동적 유형의 GC에서는 메모리 단편화를 피할 수 있고, 새로운 메모리 할당이 필요한 경우 압축된 heap의 마지막 영역부터 메모리를 할당할 수 있어 고속 메모리 할당이 가능하다.

#### Go이 힙영역을 압축하지 않는 이유

Go도 처음엔 장점을 갖고있는 동적 유형의 GC로 기획되었지만 1.4 버전에서 C와 어셈블리로 작성된 Go의 코드들을 Go로 재작성하는 과정을 거치면서 개발 기간의 제약이 걸리게 되었다. 그리고 추후에 TCMalloc 기반의 메모리 할당자를 도입하여 메모리 단편화와 할당 속도 문제를 해결했다.

> [In 1.4, much of the code has been translated to Go ](https://Go.org/doc/go1.4#runtime)
>
> [This was originally based on tcmalloc, but has diverged quite a bit. – malloc.go](http://goog-perftools.sourceforge.net/doc/tcmalloc.html)

### 세대별 GC<sup>Generational GC</sup>

세대별 GC란 heap 내의 변수를 수명(GC에서 살아남은 횟수 등)에 따라 분류하여 GC의 효율을 향상시키는 것이다.

일반적인 어플리케이션에서 새로 할당된 변수는 대부분 일찍 죽는다는 가설(세대별 가설)이 존재한다. 이 가설을 GC에 적용하면 수명이 긴 변수를 여러번 스캔할 필요가 없어 GC의 효율이 향상된다.

- 신규 할당 영역에선 GC를 자주 수행 (Minor GC)
- 오래된 할당 영역에 대해선 승격시켜 GC 빈도가 낮은 열력으로 이동 (Major GC)

Java8 HotSpot VM에선 모든 GC가 세대별 GC를 제공한다.

하지만 세대별 GC는 GC를 수행하지 않을 때도 어플리케이션에 오버헤드를 발생시킨다는 단점이 존재한다.

세대별 GC에 일반적인 레퍼런스 카운터만 적용 되어있을 때, 신규 영역과 올드 영역이 존재한다고 가정해보자. 이때 올드 영역에 있는 변수 A에서 멤버 변수 등으로 신규 영역에 있는 변수 a를 참조하고 있을 수 있다. root에서 신규 영역으로의 연결이 없는 변수를 스캔할 때 a변수는 root로부터 직접적인 연결이 없기 때문에 메로리가 반환되고 변수 A에선 반환해버린 잘못된 메모리를 참조하게 된다.

따라서 세대별 GC를 위의 예시와 같은 문제를 해결하기 위해선 root로 부터의 레퍼런스 카운터 뿐만 아니라 올드 영역에서 신규 영역간의 참조에 대한 추가적인 정보가 필요하다. 이러한 부수적인 처리를 Write Barrier라고 한다.

#### Go에서 세대별 GC가 없는 이유

일반적으로 Write Barrier라는 오버헤드가 존재하지만 세대병 GC를 적용하면 이보다 큰 효용성을 얻을 수 있기 때문에 이를 적용하지만 Go에서 write barrier에대한 오버헤드를 허용하지 않았다.

또한, Go는 컴파일러 수준에서 사용되지 않는 변수들을 분석하는 성능이 우수하고, 필요한 경우 heap 영역에 메모리가 할당되지 않도록 유저가 제어할 수 있도록 하기 때문에 일반적으로 수명이 짧은 변수는 힙이 아닌 스택영역에 할당된다. 때문에 세대별 GC로 얹을 수 있는 이익이 일반적인 runtime GC에 비해 적다.

물론 수명이 긴 변수에 대한 GC에서 매번 GC 스캔을 하는 소모적인 작업이 남아있지만, Go의 GC는 병렬적으로 수행되기 때문에 합리적인 결정이라고 생각한다고 Google의 lan Lance Taylor가 [Why Go garbage-collector not implement Generational and Compact gc?](https://groups.google.com/forum/#!topic/Go-nuts/KJiyv2mV2pU)에서 언급했다.