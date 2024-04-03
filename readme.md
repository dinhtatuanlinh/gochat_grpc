## Cài đặt ban đầu
- [protobuf](https://github.com/protocolbuffers/protobuf/releases) chọn version win64.zip cho windows. Giải nén đưa vào ổ c lấy đường dẫn tới thư mục bin chứa file protoc.exe. Điền đường dẫn vào biến môi trường path
- [protobuf dependencies](https://grpc.io/docs/languages/go/quickstart/) cài đặt protobuf dependencies cho go
- protoc --go_out=internal/delivery/proto --go-grpc_out=internal/delivery/proto internal/delivery/proto/api.proto
