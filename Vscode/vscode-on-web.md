iPad Pro를 사용중인데 컴퓨터를 대체할 폼펙터라고 한껏 광고를 해놓고는 개발도 못하는 잉여 컴퓨터로 유튜브를 보는데나 쓰고 있다가 vscode를 원격으로 접속해서 사용할 수 있는 레포지토리를 찾았다.

https://github.com/cdr/code-server/

리눅스 기반의 원격 서버에서 vscode를 실행하도록 만들어 주는 프로젝트다. 원격 vscode 실행을 위해 클라우드 서버를 한대 구해 실행을 할 수 있지만, 단지 원격 개발을 위해 비용이 드는 것을 피하고 싶기 때문에 WSL<sup>Windows Subsystem for Linus</sup>를 통해 vscode 리모트 서버를 설정했다.

<!-- 1. 도커 설치 팔요한가?
## 윈도우10에 도커 설치 https://blog.jayway.com/2017/04/19/running-docker-on-bash-on-windows/

윈도우10 우분투를 사용하는 경우 윈도우에도 도커가 설치되어있어야 제대로 동작을 한다.

- 먼저 https://www.docker.com/products/docker-desktop 로 접속하여 윈도우에 도커를 설치한다.

```terminal
$ docker -v
도커 버전 표시
$ docker-machine regenerate-certs dev
```

## 우분투에 도커 설치 https://www.bsidesoft.com/?p=7820
```terminal
$ sudo apt-get update
$ sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo apt-key fingerprint 0EBFCD88
$ sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
```

도커 CE 설치
```terminal
$ sudo apt-get update
$ sudo apt-get install -y docker-ce docker-ce-cli containerd.io
$ apt-cache madison docker-ce
```

## 윈도우 도커 데몬 Ubuntu 에서 접근하기 https://m.blog.naver.com/hahaysh/221233313670

윈도우 도커 설정에서 'Expose daemon on tcp://localhost:2375 without TLS' 옵션을 키고 윈도우 도커를 restart한다.

우분투에서 .bashrc 파일에 `export DOCKER_HOST=localhost:2375`를 추가한 후 bashrc 적용을 위해 재접속한다.

그럼 완성
-->

# WSL(Ubuentu) 설치

Windows 10 부터 WSL를 제공한다. WSL이라는 말 그대로 윈도우에서 리눅스를 사용할 수 있도록 만들어 준다.

1\. [제어판-프로그램-프로그램 및 기능-Window 기능 켜기/끄기] 에서 `Linux용 Windows 하위 시스템` 항목을 체크하여 리눅스 서브 시스템을 사용할 수 있도록 설정한다.

![vscode-on-ipad-001](https://user-images.githubusercontent.com/18159012/79640205-412c7c00-81cb-11ea-8b3a-b4344c1c02fb.PNG)

2\. MS Store에서 Ubuntu를 검색하여 설치한다.

![vscode-on-ipad-002](https://user-images.githubusercontent.com/18159012/79640214-4e496b00-81cb-11ea-8f87-5844716f53d8.PNG)

3\. 설치된 Ubuntu 어플리케이션을 실행한다.

![vscode-on-ipad-003](https://user-images.githubusercontent.com/18159012/79640230-65885880-81cb-11ea-8e38-0194315270a1.PNG)

<!-- https://eungbean.github.io/2019/11/04/remote-vscode/ -->
# code-server 설치 

Ubuntu를 실행한 후 원격 vscode를 위해 code-server를 설치한다. 도커로 설치할 수도 있지만 간단한 실행을 위해 바이너리를 다운받아서 설정했다.

```terminal
$ wget https://github.com/cdr/code-server/releases/download/2.1665-vsc1.39.2/code-server2.1665-vsc1.39.2-linux-x86_64.tar.gz
$ tar -xvf code-server2.1665-vsc1.39.2-linux-x86_64.tar.gz
$ mv code-server2.1665-vsc1.39.2-linux-x86_64 code-server
```

압축 해제된 폴더를 code-server로 설정했다. 

# code-server 실행

압축이 해제된 폴더 안에 code-server라는 파일이 존재한다. 해당 파일을 통해 vscode를 리눅스 서버로 실행하여 원격으로 접속할 수 있다.

```terminal
$ ./code-server/code-server [WORKSPACE DIR]
```

그냥 실행하면 실행될 때 마다 임의의 비밀번호를 만들어 내는데 외부에서 접근할 때 사용되기 때문에 서버를 띄울 때 마다 확인을 해줘야 하고, 랜덤한 긴 문자열을 만들어 내기 때문에 외워서 입력을 하기에도 어렵다. 때문에 커스텀 비밀번호를 설정하여 직접 만든 비밀번호로 접속할 수 있도록 했다.

```terminal
$ echo "export PASSWORD='[CUSTOM PASSWORD]]'" >> ~/.bashrc
$ source ~/.bashrc
```

<!-- https://prolite.tistory.com/1380 -->
## 포트 포워딩

위까지의 설정으로는 내부망에선 접속이 가능하지만, 외부망에선 원격으로 접속이 되지 않는다. 현대의 대부분의 PC는 공유기를 사용하기 때문에 외부망과 내부망이 분리가 되어있어, 공유기(NAT)에서 홀펀칭으로 포트 포워딩을 해주지 않으면 PC의 IP까지 접근이 어렵다.

내가 쓰는 공유기가 iptime이기 때문에 iptime을 기준으로 설명한다.

1\. 브라우저에서 http://192.168.0.1/ 로 접속하여 공유기 관리 페이지로 접속한다.

![vscode-on-ipad-004](https://user-images.githubusercontent.com/18159012/79640577-47bbf300-81cd-11ea-82b1-3dbd00dff1d6.PNG)

2\. [고급설정-NAT/라우터 관리]로 이동하여 포트포워딩 설정에 새로운 규칙을 추가 후 설정을 저장한다.

![vscode-on-ipad-005](https://user-images.githubusercontent.com/18159012/79640585-52768800-81cd-11ea-83ac-9e4badd08ca7.PNG)

포트포워딩 사용자정의를 HTTP로 설정  
내부 IP 주소를 Ubuntu가 동작하는 컴퓨터의 내부망 IP 주소  
외부포트를 기본으로 두고 (알아서 처리됨) 내부 포트를 8080~8080으로 설정

3\. [기본설정-시스템 요약 정보]로 이동하여 외부 IP 주소(NAT IP 주소)를 확인한다.
![vscode-on-ipad-006](https://user-images.githubusercontent.com/18159012/79640606-6621ee80-81cd-11ea-9f43-84599aad00c9.PNG)

4\. 외부 IP 주소에 포트 포워딩에서 설정한 외부 포트를 적용하여 외부망에서 접속한다.

![vscode-on-ipad-007](https://user-images.githubusercontent.com/18159012/79640614-6f12c000-81cd-11ea-9f08-2a085ede0527.PNG)

공유기의 IP의 경우 동적 IP 연결로 되어있는 경우 IP의 주소가 바뀔 가능성이 존재한다. 공유기에서 DDNS를 지원한다면 설정하여 더 간편하게 사용할 수 있다.