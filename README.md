# Crawl toàn bộ điểm thi THPTQG-2021
Một script nho nhỏ viết bằng Go để crawl toàn bộ điểm thi THPTQG-2021, mình đã crawl sẵn toàn bộ ở đây:<br>
https://drive.google.com/drive/u/0/folders/1IGb3n_ieBlsfOtND2nTvsB7nlFoYURXO<br>
Trong thư mục có 64 file tương ứng với 64 tỉnh thành, ngoài ra còn có file `total` gộp lại từ 64 file trên (vì một lý do tà thuật nào đấy mà mình không thể up file `total` với đuôi `.csv` lên Google Drive được nên bạn nào cần thì có thể tải xuống rồi đổi đuôi file nhé)

<img src="https://raw.githubusercontent.com/balldk/crawlscore-thptqg-2021/master/screenshots/demo.gif" width="700">

Dữ liệu mẫu:<br><br>
<img src="https://raw.githubusercontent.com/balldk/crawlscore-thptqg-2021/master/screenshots/sample.png" width="700">

## Yêu cầu
1. Đã tải source về
2. Đã cài đặt Go

## Cách dùng
Tải Dependencies
```bash
go mod download
```
Chạy chương trình
```bash
go run .
```
hoặc bạn cũng có thể build ra binary
```bash
go build .
./crawlscore
```

## Tuỳ chỉnh tham số
Một số tham số bạn có thể thay đổi trong file `.env`
```env
PATCH_SIZE=100
PATCH_DELAY=0.1
OUTPUT_FOLDER=data
TOTAL_FILENAME=total.csv
```
- Vì nguồn data để crawl có cơ chế chống DOS, do đó để hạn chế bị chặn thì mình đã cho chương trình chạy theo cơ chế crawl lần lượt theo từng patch với `PATCH_SIZE` là độ lớn của từng patch, `PATCH_DELAY` là thời gian chờ để crawl patch tiếp theo.
- `OUTPUT_FOLDER` là tên thư mục mà dữ liệu được xuất ra.
- `TOTAL_FILENAME` là tên của file tổng hợp tất cả dữ liệu từ 64 tỉnh thành.

## Một số lưu ý nhỏ
1. Chỉ có TP.HCM (Mã tỉnh thành 02) là có đầy đủ họ tên, ngày tháng năm sinh, giới tính (do chỉ có TP.HCM cung cấp).
2. Chỉ nên để giá trị `PATCH_SIZE` trong khoảng từ 100 đến 200, nhưng để ổn định nhất thì mình khuyên chỉ nên để 100 thôi, nó sẽ chạy xong sau một giấc ngủ trưa.
3. Mình không đảm bảo thuật toán tìm SBD của mình là hoàn toàn chính xác, nên có bất kì sai sót gì thì có thể báo cho mình nhé. Cụ thể thì có bạn phát hiện ở mã tỉnh 35 dữ liệu mình crawl không đủ, mình đã crawl bằng tay và cập nhật lại ở link Google Drive.
4. Trong quá trình chạy các bạn có thể bị chặn, cách duy nhất có lẽ là phải đổi địa chỉ IP thôi. Mình dùng WARP của Cloudflare, mỗi lần bị chặn thì chỉ cần vô `Preferences -> Connection -> Reset Encryption Keys` là được.
