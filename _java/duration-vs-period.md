Java 8에 추가된 Time 패키지에는 시간의 차이를 나타내기 위한 클래스인 `Duration`과 `Period`가 존재한다. 비슷해 보이는 두 클래스는 조금 다른 성격을 갖는다.

# Duration

`Duration` 클래스는 두 시간 사이의 차이를 초와 나노초를 저장한다. `Duration.between()` 정적 메서드를 통해 시작과 끝 시간을 파라메터로 전달하면 그 차이를 `Duration` 인스턴스로 반환한다. 주의해야할 점은 해당 메서드가 `Temporal` 형으로 파라메터를 전달 받지만, 초단위로 차이를 계산하기 때문에 파라메터 타입이 초단위를 포함하지 않는 `LocalDate`등의 클래스라면 `UnsupportedTemporalTypeException`을 발생시킨다.

```java
LocalDateTime startDateTime = LocalDateTime.of(2020, 6, 25, 9, 30, 0);
LocalDateTime endDateTime = LocalDateTime.of(2020, 6, 30, 9, 30, 10);

Duration duration = Duration.between(startDateTime, endDateTime);

System.out.println("Seconds : " + duration.getSeconds());
System.out.println("Nano : " + duration.getNano());
System.out.println("Days : " + duration.toDays());
```
```terminal
Seconds : 432010
Nano : 0
Days : 5
```

# Period

`Period` 클래스는 두 시간의 차이를 초단위 뿐만 아니라, 년/월/일 단위로 나타낸다.

```java
LocalDate startDate = LocalDate.of(2011, 3, 1);
LocalDate endDate = LocalDate.of(2020, 6, 25);

Period period = Period.between(startDate, endDate);

System.out.println("Year : " + period.getYears());
System.out.println("Month : " + period.getMonths());
System.out.println("Day : " + period.getDays());
```
```terminal
Year : 9
Month : 3
Day : 24
```