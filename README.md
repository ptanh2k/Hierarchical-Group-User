**Đề bải**

Hệ thống giúp người dùng quản lý group và các user trong group, thiết kế DB cho tính
năng này.

Mỗi tổ chức bao gồm nhiều phòng ban/bộ phận khác nhau tạo thành 1 cây phân cấp như hình minh họa
phía dưới (phần được highlight bằng khung trắng). Ví dụ: Viettel Tickets là root, trong Viettel Tickets bao
gồm các bộ phận như Tier 1 (chuyên xử lý alerts), Tier 3 (chuyên xử lý tickets), v.v. Trong group Tier 1 lại
có các group A, B, C khác để đảm nhận các nhiệm vụ riêng, trong group A lại có thể có các group con
khác.

**Yêu cầu**:

- Dựng model, init dữ liệu sample.
- Cho tên 1 group ví dụ Tier 1, thực hiện truy vấn tìm ra cây thư mục của group này, bao gồm các
  group con ở tất cả các cấp và level của chúng để portal có thể hiển thị được cây thư mục như
  trên. Sử dụng Pgadmin hoặc các tool tương tự thực hiện truy vấn.
- Cũng với yêu cầu trên viết 1 service api thực hiện nghiệp vụ này. (Có thể sử dụng Python hoặc
  Golang client).

**Chạy chương trình Demo**

- Kết nối PostgreSQL, thực thi các câu lệnh trong file script.sql.

- Tại thư mục /api/main, chạy câu lệnh:

```
go run . --host=<db-hostname> --port=<db-port> --user=<db-username>  --password=<db-password> --db=<db-to-use>
```

Chương trình sẽ chạy trên http://localhost:8080

**API**

- Group

```
GET /groups                     Lấy thông tin các group
GET /groups/:cur_group_id/sub   Lấy thông tin các group con (tên, path, cấp)
POST /groups                    Tạo group mới
```

- User

```
GET /users              Lấy thông tin toàn bộ người dùng
GET /users/:id          Lấy thông tin người dùng dựa trên ID
```

**Diagram**

![Group-user drawio](https://user-images.githubusercontent.com/50461553/187580085-bd94c42b-3af5-4474-9dae-d5cc97cd08a4.png)
