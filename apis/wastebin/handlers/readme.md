``` go
var lastUpdate time.Time
var debounceDuration = 5 * time.Second

func (w WasteBinHandler) WebSocketUpdateWasteBin(c *websocket.Conn) {
    defer c.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var updateWasteBinReq models.CreateWasteBinReq
    if err := c.ReadJSON(&updateWasteBinReq); err != nil {
        c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error reading JSON: %v", err)))
        return
    }

    // Kiểm tra nếu thời gian cập nhật cuối quá gần với lần trước
    if time.Since(lastUpdate) < debounceDuration {
        c.WriteMessage(websocket.TextMessage, []byte("Update skipped due to debounce"))
        return
    }

    lastUpdate = time.Now()

    // Tiếp tục xử lý cập nhật và lưu vào database
    // ...
}
```