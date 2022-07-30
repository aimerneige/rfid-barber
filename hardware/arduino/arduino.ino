/**
   SSD1306:    GND   VCC   D0   D1   RES   DC   CS
   ESP32:      GND   3V3   14   13   TX2   RX2  5
*/

#include <SPI.h>
#include <Wire.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>
#include <MFRC522.h>

#define debug false

#define SS_PIN 21
#define RST_PIN 0
MFRC522 mfrc522 = MFRC522(SS_PIN, RST_PIN);

#define SCREEN_WIDTH 128 // OLED display width, in pixels
#define SCREEN_HEIGHT 64 // OLED display height, in pixels

// Declaration for SSD1306 display connected using software SPI
#define OLED_MOSI   13
#define OLED_CLK    14
#define OLED_DC     16
#define OLED_CS     5
#define OLED_RESET  17

// SSD1306 OLED display
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT,
                         OLED_MOSI, OLED_CLK, OLED_DC, OLED_RESET, OLED_CS);

void rfidSetup() {
  // while (!Serial); // Do nothing if no serial port is opened (added for Arduinos based on ATMEGA32U4)
  SPI.begin(); // Init SPI bus
  mfrc522.PCD_Init(); // Init MFRC522
  mfrc522.PCD_DumpVersionToSerial(); // Show details of PCD - MFRC522 Card Reader details
  Serial.println(F("Scan PICC to see UID, SAK, type, and data blocks..."));
}

// setup ssd1306 display
void displaySetup() {
  Serial.println("Start Setup SSD1306!");
  if (! display.begin(SSD1306_SWITCHCAPVCC)) {
    Serial.println(F("SSD1306 allocation failed"));
    // loop forever
    while (1) {
      delay(10);
    }
  }
  display.clearDisplay();
  display.display();
  delay(1000);
  Serial.println("SSD1306 Setup Seccessful!");
}

void setup() {
  Serial.begin(115200);
  rfidSetup();
  displaySetup();
}

void loop() {
  // Reset the loop if no new card present on the sensor/reader. This saves the entire process when idle.
  if ( ! mfrc522.PICC_IsNewCardPresent()) {
    return;
  }
  // Verify if the NUID has been readed
  if ( ! mfrc522.PICC_ReadCardSerial()) {
    return;
  }

  // Dump debug info about the card; PICC_HaltA() is automatically called
  if (debug) {
    mfrc522.PICC_DumpToSerial(&(mfrc522.uid));
  }

  String id = getID();
  Serial.println(id);

  display.clearDisplay();
  display.setTextSize(1);
  display.setTextColor(SSD1306_WHITE);
  display.setCursor(0, 0);
  display.println("Copyright (c) 2022");
  display.println("by Aimer Neige");
  display.println("All rights reserved");
  display.println(id);
  display.display();

  delay(500);
}

String getID() {
  unsigned long hex_num;
  hex_num =  mfrc522.uid.uidByte[0] << 24;
  hex_num += mfrc522.uid.uidByte[1] << 16;
  hex_num += mfrc522.uid.uidByte[2] <<  8;
  hex_num += mfrc522.uid.uidByte[3];
  return String(hex_num);
}
