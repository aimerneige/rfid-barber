# hardware

硬件相关内容。

## 硬件

- ESP32
- RFID-RC522
- SSD1306

## 连线

|         |     |     |     |     |     |     |     |
| ------- | --- | --- | --- | --- | --- | --- | --- |
| ESP32   | GND | 3V3 | 14  | 13  | TX2 | RX2 | 5   |
| SSD1306 | GND | VCC | D0  | D1  | RES | DC  | CS  |

|            |     |     |     |      |      |     |        |
| ---------- | --- | --- | --- | ---- | ---- | --- | ------ |
| ESP32      | GND | 3v3 | RST | 19   | 23   | 18  | 21     |
| RFID-RC522 | GND | VCC | 0   | MISO | MOSI | SCK | SS/SDA |