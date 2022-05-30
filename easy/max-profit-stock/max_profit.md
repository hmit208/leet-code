## idea
### Kadane's Algorithm

- có thể đưa bài toán về sử dụng Kadane's Algorithm như sau:
- Giả sử có mảng ban đầu là a0, a1, a2, a3, a4
- Ta chuyển về mảng mới b0, b1, b2, b3, b4
  - b[i] = 0 if i == 0
  - b[i] = a[i] - a[i-1] if i > 0
- Sau đó để tìm lợi nhuận lớn nhất thì tương đương với việc tìm dãy con có tổng lớn nhất của b. Ví dụ nếu lợi nhuận lớn nhất từ a là a2 và a4 thì max_profit = a4-a2
 b4 = a4 - a3
 b3 = a3 - a2
 b3 + b4 = a4- a3 + a3 - a2 = a4 - a2 = max_profit


==> tại đây ta có thể sử dụng kadane để giải quyết bài toán



### Duyệt bình thường
- Ý tưởng là duyệt mảng a, lấy phần tử đang a[i] trừ đi phần từ nhỏ nhất trong a[0]-> a[i]
- sử dụng thêm một biến là minPrice để lưu giá trị nhỏ nhất trong mảng a, init = a[0]