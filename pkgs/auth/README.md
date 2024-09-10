**JWT (JSON Web Token)** là một chuẩn mở được sử dụng để truyền tải thông tin giữa các bên dưới dạng một đối tượng JSON an toàn. Thông tin trong JWT có thể được xác thực và tin cậy vì nó được ký bằng chữ ký số. JWT thường được sử dụng trong việc xác thực và phân quyền trong các ứng dụng web.

### **Cấu trúc của JWT**

JWT bao gồm 3 phần chính, được phân tách bằng dấu chấm (`.`):

1. **Header** (Đầu đề): 
   - Chứa thông tin về loại token và thuật toán ký số được sử dụng.
   - Ví dụ:
     ```json
     {
       "alg": "HS256",
       "typ": "JWT"
     }
     ```

2. **Payload** (Nội dung):
   - Chứa các thông tin (claims) mà bạn muốn truyền tải và xác thực. Ví dụ như ID người dùng, vai trò người dùng, thời gian hết hạn...
   - Ví dụ:
     ```json
     {
       "sub": "1234567890",
       "name": "John Doe",
       "admin": true
     }
     ```

3. **Signature** (Chữ ký):
   - Là phần chữ ký được tạo ra từ header, payload và một `secret key` (khóa bí mật). Chữ ký này dùng để xác minh tính toàn vẹn của JWT và đảm bảo rằng payload không bị thay đổi.

JWT trông như thế này:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

### **Tại sao phải sử dụng JWT?**

1. **Bảo mật**:
   - JWT được ký bằng một khóa bí mật, do đó người nhận có thể xác thực token mà không cần liên lạc với máy chủ xác thực. Điều này giúp giảm tải cho máy chủ và tăng hiệu suất hệ thống.

2. **Tính di động**:
   - JWT là một đối tượng JSON đơn giản, có thể được truyền qua các giao thức khác nhau như HTTP headers, URL, cookies mà không cần bất kỳ thay đổi nào.

3. **Phi trạng thái**:
   - JWT cho phép xây dựng hệ thống phi trạng thái (stateless). Máy chủ không cần phải lưu trữ session của người dùng, giúp tiết kiệm tài nguyên và tăng tính mở rộng.

4. **Đa nền tảng**:
   - JWT là chuẩn mở và được hỗ trợ bởi nhiều nền tảng khác nhau, giúp dễ dàng tích hợp với nhiều ngôn ngữ lập trình và framework.

5. **Quản lý phân quyền**:
   - JWT có thể chứa các thông tin phân quyền (claims), giúp việc kiểm soát truy cập trở nên dễ dàng hơn, đặc biệt khi cần kiểm tra quyền truy cập vào các API.

### **Khi nào nên sử dụng JWT?**

JWT đặc biệt hữu ích trong các tình huống sau:

- **API authentication (Xác thực API):** Khi bạn cần xác thực người dùng trước khi cho phép truy cập vào các tài nguyên trên server.
- **Single Sign-On (SSO):** JWT là một giải pháp tuyệt vời cho việc đăng nhập một lần, nơi một token có thể được sử dụng để xác thực trên nhiều ứng dụng hoặc domain.
- **Phân quyền trong ứng dụng web:** Bạn có thể sử dụng JWT để xác định và phân quyền cho các vai trò người dùng khác nhau.