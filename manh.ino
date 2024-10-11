#include <WiFi.h>
#include <HTTPClient.h>
#include <ArduinoJson.h>
#include <TinyGPSPlus.h>
#include <SoftwareSerial.h>
#include "Adafruit_VL53L0X.h"
#include <HX711_ADC.h>

#define RL 47         // Giá trị của điện trở RL là 47K
#define m -0.263      // Nhập giá trị Slope đã tính toán cho H2S
#define b 0.42        // Nhập giá trị giao điểm đã tính toán cho H2S
#define Ro 10         // Nhập giá trị Ro đã tìm cho H2S (đây là ví dụ, bạn cần hiệu chuẩn nó)
#define MQ_sensor 34  // Cảm biến được kết nối với chân 34

HX711_ADC LoadCell(13, 12);
Adafruit_VL53L0X lox = Adafruit_VL53L0X();

const char* ssid = "Internet";
const char* password =  "0906200216092008?";

const int MQ135_AO_PIN = 35;

String serverAddress = "https://postman-echo.com/";

TinyGPSPlus gps;
SoftwareSerial ss(16, 17); // RX, TX

unsigned long t = 0;

void sendData_HttpPost_Json(float filled_level, float ppm, double latitude, double longitude, float weight) {
  Serial.println(F("Sending data to server using HTTP POST request and json data"));

  String url = serverAddress + "post";
  Serial.println("Url: " + url);

  HTTPClient http;
  http.begin(url);
  // Specify content-type header: application/json
  http.addHeader("Content-Type", "application/json");

  StaticJsonDocument<200> doc;
  doc["filled_level"] = filled_level;
  doc["air_quality"] = ppm;
  doc["latitude"] = latitude;
  doc["longitude"] = longitude;
  doc["weight"] = weight;

  String payload;
  serializeJson(doc, payload);
  Serial.println("Data: " + payload);

  int httpCode = http.POST(payload);
  if (httpCode == 200) {
    String responsePayload = http.getString();
    StaticJsonDocument<200> jsonResponse;
    deserializeJson(jsonResponse, responsePayload);
    Serial.println("Response Payload: ");
    Serial.println("  \"data\": {");
    Serial.print("    \"filled_level\": ");
    Serial.println(jsonResponse["data"]["filled_level"].as<float>(), 6);
    Serial.print("    \"air_quality\": ");
    Serial.println(jsonResponse["data"]["air_quality"].as<float>(), 6);
    Serial.print("    \"latitude\": ");
    Serial.println(jsonResponse["data"]["latitude"].as<double>(), 6);
    Serial.print("    \"longitude\": ");
    Serial.println(jsonResponse["data"]["longitude"].as<double>(), 6);
    Serial.print("    \"weight\": ");
    Serial.println(jsonResponse["data"]["weight"].as<double>(), 6);
    Serial.println("  }");
  } else {
    Serial.print(F("Send data failed, HTTP code: "));
    Serial.println(httpCode);
  }
  http.end();
}

void setup() {
  Serial.begin(115200);
  ss.begin(9600);
  delay(20000);
  Wire.begin();
  // Khởi tạo cảm biến VL53L0X
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
  VL53L0X_RangingMeasurementData_t measure;
  lox.rangingTest(&measure, false); // pass in 'true' to get debug data printout!
  float distance = 0;
  float filled_level = 0;
  if (measure.RangeStatus != 4) {  // phase failures have incorrect data
    distance = measure.RangeMilliMeter - 20; // Điều chỉnh offset nếu cần
    distance = distance / 10.0; // Chuyển đổi từ mm sang cm
    Serial.print("Distance (cm): "); 
    Serial.println(distance);
    // Tính toán filled_level
    filled_level = ((24.0 - distance) / 24.0) * 100.0;
    if (filled_level < 0) filled_level = 0;          // Nếu khoảng cách lớn hơn chiều cao thùng
    if (filled_level > 100) filled_level = 100;      // Nếu có lỗi trong phép đo
    Serial.print("Filled Level (%): ");
    Serial.println(filled_level);
  } else {
    Serial.println(" out of range ");
  }
  ///// MQ136
  float VRL = analogRead(MQ_sensor) * (5.0 / 4096.0); 
  float Rs = ((5.0 * RL) / VRL) - RL; 
  float ratio = Rs / Ro; 
  float ppm = pow(10, ((log10(ratio) - b) / m)); 
  Serial.print("H2S (ppm) = "); Serial.println(ppm);
  ///////////// GPS
  unsigned long start = millis();
  while (!gps.location.isUpdated() && millis() - start < 2000) { // timeout sau 2 giây
    while (ss.available() > 0) {
      gps.encode(ss.read());
    }
  }
  double latitude = gps.location.isValid() ? gps.location.lat() : 0.0;
  double longitude = gps.location.isValid() ? gps.location.lng() : 0.0;
  Serial.print("Vi do = "); Serial.println(latitude, 6);
  Serial.print("Kinh do = "); Serial.println(longitude, 6);

  static boolean newDataReady = 0;
  if (LoadCell.update()) newDataReady = true;
  float weight = 0;
  if (newDataReady) {
    weight = LoadCell.getData();
    if (weight < 0) weight = 0;
    Serial.print("Load_cell output val: "); Serial.println(weight);
    newDataReady = 0;
  }
  
  // Send data to server via HTTP request
  if (WiFi.status() == WL_CONNECTED) {
    sendData_HttpPost_Json(filled_level, ppm, latitude, longitude, weight);
  } else {
    Serial.println("WiFi Disconnected");
  }
  if (LoadCell.getTareStatus() == true) {
    Serial.println(F("Tare complete"));
  }
  // Delay for next loop iteration
  delay(3000);
}
