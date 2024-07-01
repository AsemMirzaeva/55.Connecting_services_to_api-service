### Uyga Vazifa: gRPC Cancellation Vazifa Boshqaruv Tizimida Tatbiq Etish

#### Maqsad
Ushbu vazifaning maqsadi `gRPC` cancellation funksiyasini vazifa boshqaruv tizimida tushunish va tatbiq etishdir. Vazifa boshqaruv tizimi foydalanuvchilarga yangi vazifalar yaratish, ro'yxatga olish va cancellation imkonini beradi.

#### Vazifalar

1. **Loyihani Tayyorlash**:
   - Task xizmati uchun kerakli `.proto` fayllarini aniqlash.

2. **Task Serviceni Aniqlash**:
   - Quyidagi RPC usullariga ega bo'lgan xizmat yaratish:
     - `CreateTask` - yangi vazifa yaratish.
     - `ListTasks` - barcha vazifalarni ro'yxatga olish.
     - `CancelTask` - muayyan vazifani cancellation.
   
   Misol uchun `.proto` file:
   ```proto
   syntax = "proto3";

   option go_package="./taskpb";

   service TaskService {
     rpc CreateTask (TaskRequest) returns (TaskResponse);
     rpc ListTasks (Empty) returns (TaskList);
     rpc CancelTask (CancelRequest) returns (CancelResponse);
   }

   message TaskRequest {
     string task_description = 1;
   }

   message TaskResponse {
     string task_id = 1;
     string status = 2;
   }

   message TaskList {
     repeated TaskResponse tasks = 1;
   }

   message CancelRequest {
     string task_id = 1;
   }

   message CancelResponse {
     string status = 1;
   }

   message Empty {}
    ```

3. **Serverni Tatbiq Etish**:
- Unikal task IDlari bilan xotirada vazifalarni saqlash

4. **Clientni Tatbiq Etish:**:

### Talablar
- TaskService unikal `ID`lar bilan tasklarni boshqarishi kerak.
- Task yaratish uzoq davom etadigan operatsiyani simulyatsiya qilishi kerak.
- Task bekor qilish taskning holatini "`Cancelled`" qilib o'zgartirishi kerak.
- Client task operatsiyalarini bekor qilish uchun `context.Context`dan foydalanishi kerak.
- Taskning hayot aylanishi voqealarini kuzatish uchun to'g'ri loglash tatbiq etilishi kerak.