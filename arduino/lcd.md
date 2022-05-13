16x2의 LCD 디스플레이 모듈(LCD 1602)에 I2C 칩을 결합한 형태로 사용했다. I2C는 기존 1602 LCD가 여러개의 핀을 사용하던 것을 4개의 핀만으로 아두이노와 결합하여 사용할 수 있도록 도와준다.

```ino
/*
 * CONNECTION
 * 
 * VCC --> 5v
 * GND --> GND
 * SDA --> A4
 * SCL --> A5
 */

#include <Wire.h> // i2C 통신을 위한 라이브러리
#include <LiquidCrystal_I2C.h> // LCD 1602 I2C용 라이브러리
 
LiquidCrystal_I2C lcd(0x27,16,2); // 접근주소: 0x3F or 0x27
 
byte pac_man1[8] = { // 팩맨 출력 입벌림, 얼굴방향 오른쪽
  0b01110,
  0b11011,
  0b11110,
  0b11100,
  0b11100,
  0b11110,
  0b11111,
  0b01110,
};
byte pac_man2[8] = { // 팩맨 출력 입닫음, 얼굴방향 오른쪽
  0b01110,
  0b11011,
  0b11111,
  0b11111,
  0b11100,
  0b11111,
  0b11111,
  0b01110,
};
 
byte pac_man3[8] = { // 팩맨 출력 입벌림,  얼굴방향 왼쪽
  0b01110,
  0b11011,
  0b01111,
  0b00111,
  0b00111,
  0b01111,
  0b11111,
  0b01110,
};
byte pac_man4[8] = { // 팩맨 출력 입닫음,  얼굴방향 왼쪽
  0b01110,
  0b11011,
  0b11111,
  0b11111,
  0b00111,
  0b11111,
  0b11111,
  0b01110,
};
 
void setup()
{
  lcd.init(); // LCD 초기화
  lcd.backlight(); // 백라이트 켜기
  lcd.createChar(1, pac_man1); // 팩맨 얼굴1
  lcd.createChar(2, pac_man2); // 팩맨 얼굴2
  lcd.createChar(3, pac_man3); // 팩맨 얼굴3
  lcd.createChar(4, pac_man4); // 팩맨 얼굴4
  
}
void loop(){
  // 첫번째 줄
  for(int first_line=0; first_line<16; first_line++) {
    if(first_line == 0) {        // 프로그램 처음 실행 시
      
      lcd.setCursor(0, 0);  // 커서 위치 0,0    첫번째 줄
      lcd.write(byte(1));       // 팩맨 얼굴 출력
 
      lcd.setCursor(1, 0);    // 커서 위치 1,0    첫번째 줄
      lcd.print("Hello Yujin");// Hello, Eduino!! 출력        
            
      lcd.setCursor(0, 1);    // 커서 위치 0,1    두번째 줄
      lcd.print("I love you!");  // Have a nice day! 출력
    } else {
      for(int blank = 0; blank < first_line; blank++) {  // 팩맨이 지나온 길을 빈칸으로 만들기
        lcd.setCursor(blank, 0);  // 빈칸
        lcd.print(" ");
      }
      int place = first_line;  // fisrt_line 변화를 막기 위해 place 변수 만듦
      lcd.setCursor(place++, 0);  // 팩맨 한칸씩 이동
      if(place % 2 == 1) {
        lcd.write(byte(1));     // 팩맨 1
      } else { 
        lcd.write(byte(2));     // 팩맨 2
      }
    }
    delay(500);
  }
  
  for(int second_line = 15; second_line >= 0; second_line --) {      // 팩맨 오른쪽부터 시작
    if(second_line == 15) {    // for문 맨 처음 시작 시
      for(int blank = 0; blank < 16; blank++) {  // 첫번째 줄 모두 다 빈칸으로 만들기
        lcd.setCursor(blank, 0);  // 빈칸
        lcd.print(" ");
      }
      lcd.setCursor(15, 1); // 맨 오른쪽에 팩맨 위치
      lcd.write(byte(3));      
    } else {
      for(int blank = 15; blank >= second_line; blank--) { // 팩맨이 지나온 길 빈칸으로
        lcd.setCursor(blank, 1);  // 빈칸
        lcd.print(" ");
      }
      int place = second_line;   // second_line 변화를 막기 위해 place 변수 만듦
      lcd.setCursor(place++, 1);  // 팩맨 한칸씩 이동
      if(place % 2 == 0) {
        lcd.write(byte(3)); // 팩맨3
      } else {
        lcd.write(byte(4));     // 팩맨4
      }
    }
    delay(500);
  }
}
```