기존의 웹의 경우 정적웹을 지원하면서 request에 대한 response document를 반환하는 것을 의미했다. 하지만 현대로 넘어오면서 동적 웹에대한 필요성이 생기면서 websocket과 event-source가 html5 표준에서 추가되었다.

websocket은 기존 response처럼 보내고 tcp 연결을 끊지 않고 양방향 연결을 유지한 상태에서 클라이언트와 서버가 서로 데이터를 주고 받는 것을 의미한다.

event-source는 websocket이 양방향 연결인 것과 달리 서버에서 클라이언트로의 데이터 전송만을 단방향 연결을 통해서 전달한다. server-sent event에 대한 웹 인터페이스를 의미한다.