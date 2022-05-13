피에조 부저는 소리에 대한 원하느는 음을 낼 수 있는 부저다.

```ino
/*
 * CONNECTION
 *
 * (+) --> 13
 * (-) --> GND
 */

int Buzzer_Pin = 13; // 부저의 +측에 연결된 핀                  
 
int Tones[8] = {261, 294, 330, 349, 392, 440, 494, 523};              
int Tones_Number; // Tones에 들어갈 변수
 
void setup() {
  // put your setup code here, to run once:
}
 
void loop() {
/*도 레 미 파 솔 라 시 도를 연속적으로 연주한다.*/
  for(Tones_Number=0; Tones_Number<8; Tones_Number++){
     tone(Buzzer_Pin,Tones[Tones_Number]); // tone(출력 핀(Pin) 번호, 주파수 값)
     delay(500);
  }
  noTone(Buzzer_Pin);
  delay(1000);
}
```