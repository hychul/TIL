직접 회원가입이나 로그인을 제공하지 않으면서 다른 플랫폼의 로그인 기능을 통해 로그인을 제공하는 표준을 OAuth라고 한다. 

아이디나 비밀번호와 같은 개인정보를 취급하는 것에 있어서 스타트업과 같은 작은 규모의 프로젝트에서 관련 법규를 모두 지키면서 관리하기랑 쉽지 않다. 때문에 구글이나 네이버, 카카오와 같이 큰 플랫폼을 가진 회사의 로그인 데이터를 통해서 로그인 기능을 제공하면 부담이 적다.

그리고 유저 입장에서도 하나의 계정을 통해 로그인을 할 수 있기 때문에 편리하고, 해당 기능을 제공하는 큰 플랫폼의 회사의 경우에도 자신들의 서비스를 포탈로 사용하도록 유도할 수 있기 때문에 이득이 된다.

하지만 개인정보에 대한 비용이 발생하기 때문에 민감하지 않거나 해당 기능에 대한 요청이 적은 경우 무료로 제공하고 있다.

User - sign in -> My Web -- OAuth ---> Google  
&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;┃&ensp;&ensp;<-&ensp;callback&ensp;━━┛  
&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;┃&ensp;&ensp;&ensp;&ensp;&ensp;{refresh key, api key} // api key를 통해 user 정보에 접근 가능  
&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;┃  



Google OAuth를 사용하기 위해선 'https://console.developers.google.com/'에 접속하여 OAuth 동의화면으로 이동하여 자신의 앱 이름을 설정하고 저장한다.

Google Developers 페이지에서 사용자 인증 정보 탭을 클릭하고 '사용자 인증 정보 만들기'를 클릭하여 OAuth 2.0 클라이언트 ID를 생성한다. 테스트를 위해 '승인된 리디렉션 URI'에 'http://localhost:3000'을 추가하고, '승인된 리디렉션 URI'에 리디렉션을 받아서 처리한 URI인 'http://localhost:3000/auth/google/callback' URI를 추가한다.

Go 프로젝트에서 환경변수 접근을 통해 클라이언트 ID를 가져와 사용할 예정이기 때문에 해당 키값들을 환경변수로 지정한다. (윈도우에선 재부팅 필요)

OAuth를 사용하기 위한 패키지를 추가한다.

```terminal
$ go get golang.org/x/oauth2
$ go get cloud.google.com/go
```