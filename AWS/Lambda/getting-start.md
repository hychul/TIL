# AWS Lambda
Lambda는 AWS<sup>Amazon Web Service</sup>에서 제공하는 FaaS<sup>Funtions As A Service</sup> 서비스이다. Lambda 말고도 다른 기업들에서 제공하는 비슷한 서비스들이 존재한다.

> Google Function
> Azure Function

하지만 아직까지는 Lambda의 사용률이 압도적이다.

![image](https://user-images.githubusercontent.com/18159012/75878504-09f14e00-5e5d-11ea-9dbb-b45cf8976e2c.png)

# Lambda 시작하기

AWS에 로그인을 하고 [Lambda Management Console](https://ap-northeast-2.console.aws.amazon.com/lambda/home?region=ap-northeast-2#/functions)에 들어가 함수 생성 버튼을 클릭한다.

함수 이름과 런타임(함수 언어)을 설정할 수 있다.

함수 생성 버튼을 누르면 Lambda 관리화면으로 전환된다.

# Lambda 관리하기

함수를 생성한 뒤 함수 목록에 생성한 Lambda가 보여진다. 각 Lambda를 클릭하면 함수 코드를 관리할 수 있다.
이때 NodeJS 런타임은 웹에서 코드를 수정할 수 있지만 Go 또는 Java 같은 런타임들은 패키지를 생성하여 업로드를 해야한다.

Lambda 관리 화면에서 테스트 버튼을 클릭하면 '테스트 이벤트 구성' 이라는 모달이 뜨게된다. 모달에 구성된 JSON은 Lambda 함수에서 이벤트를 구성할 객체다. 생성버튼을 클릭하고 관리 화면의 테스트 버튼을 다시 클릭하면 이벤트를 기반으로 테스트가 동작하고 결과값을 표출한다.
