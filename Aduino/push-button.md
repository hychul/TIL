```ino
void setup() {
  Serial.begin(9600);

  pinMode(8, OUTPUT);
  pinMode(7, INPUT);
}

void loop() {
  int readValue = digitalRead(7);
  Serial.println(readValue);

  if (readValue == HIGH) {
    digitalWrite(8, HIGH);
  } else {
    digitalWrite(8, LOW);
  }
}
```