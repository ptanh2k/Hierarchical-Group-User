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