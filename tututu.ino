#include <WiFi.h>
#include <ArduinoJson.h>
#include <TinyGPSPlus.h>
#include <SoftwareSerial.h>
#include <Adafruit_VL53L0X.h>
#include <HX711_ADC.h>
#include <WebSocketsClient.h>

#define RL 47         
#define m -0.263      
#define b 0.42        
#define Ro 10         
#define MQ_sensor 34  

HX711_ADC LoadCell(13, 12);
Adafruit_VL53L0X lox = Adafruit_VL53L0X();

const char* ssid = "Internet";
const char* password =  "0906200216092008?";

TinyGPSPlus gps;
SoftwareSerial ss(16, 17); // RX, TX

WebSocketsClient webSocket;
unsigned long lastConnectionAttempt = 0;

void webSocketEvent(WStype_t type, uint8_t * payload, size_t length) {
  if (type == WStype_DISCONNECTED) {
    Serial.println("WebSocket Disconnected! Reconnecting...");
  } else if (type == WStype_CONNECTED) {
    Serial.println("WebSocket Connected!");
  } else if (type == WStype_TEXT) {
    Serial.printf("Message from server: %s\n", payload);
  }
}

void sendData_WebSocket_Json(float remaining_fill, float ppm, double latitude, double longitude, float weight) {
  StaticJsonDocument<200> doc;
  doc["id"] = "0192867e-2ba2-769f-a3c6-07a7a58e70af";  
  doc["remaining_fill"] = remaining_fill;
  doc["air_quality"] = ppm;
  doc["latitude"] = latitude;
  doc["longitude"] = longitude;
  doc["weight"] = weight;

  String payload;
  serializeJson(doc, payload);
  webSocket.sendTXT(payload);  
}

void reconnectWebSocket() {
  if (millis() - lastConnectionAttempt > 5000) { // Thử kết nối lại sau mỗi 5 giây
    lastConnectionAttempt = millis();
    webSocket.beginSSL("www.waste.congmanh18.click", 443, "/wastebin/ws/update");
    webSocket.onEvent(webSocketEvent);
    webSocket.setReconnectInterval(5000);  // Tự động kết nối lại sau 5 giây nếu ngắt kết nối
  }
}

void checkWiFiConnection() {
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("WiFi disconnected! Reconnecting...");
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
      delay(1000);
      Serial.print(".");
    }
    Serial.println("\nWiFi reconnected");
  }
}

void setup() {
  Serial.begin(115200);
  ss.begin(9600);
  delay(20000);
  Wire.begin();
  
  if (!lox.begin()) {
    Serial.println(F("Failed to boot VL53L0X"));
    while (1);
  }

  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1500);
    Serial.println(F("Connecting to WiFi.."));
  }
  Serial.println(F("Connected to the WiFi network"));
  
  webSocket.beginSSL("www.waste.congmanh18.click", 443, "/wastebin/ws/update");
  webSocket.onEvent(webSocketEvent);
  webSocket.setReconnectInterval(5000);  // Thử kết nối lại WebSocket sau 5 giây nếu ngắt kết nối
  
  LoadCell.begin();
  LoadCell.start(2000, true);
  LoadCell.setCalFactor(248.34);
  if (LoadCell.getTareTimeoutFlag()) {
    Serial.println(F("Timeout, check MCU>HX711 wiring and pin designations"));
    while (1);
  } else {
    Serial.println(F("LoadCell startup is complete"));
  }
}

void loop() {
  checkWiFiConnection();  // Kiểm tra kết nối Wi-Fi
  reconnectWebSocket();    // Kiểm tra kết nối WebSocket và kết nối lại nếu cần
  webSocket.loop();        // Duy trì kết nối WebSocket

  VL53L0X_RangingMeasurementData_t measure;
  lox.rangingTest(&measure, false); 
  float distance = (measure.RangeStatus != 4) ? (measure.RangeMilliMeter - 20) / 10.0 : -1; 
  float remaining_fill = (distance > 0) ? ((24.0 - distance) / 24.0) * 100.0 : 0;
  remaining_fill = constrain(remaining_fill, 0, 100);

  Serial.print("Remaining Fill (%): ");
  Serial.println(remaining_fill);

  ///// MQ136
  float VRL = analogRead(MQ_sensor) * (5.0 / 4096.0); 
  float Rs = ((5.0 * RL) / VRL) - RL; 
  float ppm = pow(10, ((log10(Rs / Ro) - b) / m)); 
  Serial.print("H2S (ppm) = "); Serial.println(ppm);
  
  ///////////// GPS
  unsigned long start = millis();
  while (!gps.location.isUpdated() && millis() - start < 2000) {
    while (ss.available() > 0) {
      gps.encode(ss.read());
    }
  }
  double latitude = gps.location.isValid() ? gps.location.lat() : 0.0;
  double longitude = gps.location.isValid() ? gps.location.lng() : 0.0;

  Serial.print("Vi do = "); Serial.println(latitude, 6);
  Serial.print("Kinh do = "); Serial.println(longitude, 6);

  static boolean newDataReady = false;
  if (LoadCell.update()) newDataReady = true;
  float weight = (newDataReady) ? LoadCell.getData() : 0;
  weight = (weight < 0) ? 0 : weight;
  Serial.print("Load_cell output val: "); Serial.println(weight);
  Serial.println("=============================================");

  sendData_WebSocket_Json(remaining_fill, ppm, latitude, longitude, weight);

  delay(3000);  
}
