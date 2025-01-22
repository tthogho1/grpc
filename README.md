# gRPC Image Upload and Embedding Example

This repository contains sample code for sending an uploaded image file from a Go language client to a Python server using gRPC, and then embedding it with OpenAI's CLIP.

## Set Up

1. Install the necessary packages:

   For Python:
   ```bash
   pip install grpcio grpcio-tools torch transformers Pillow
   ```

   For Go:
   ```bash
   go get google.golang.org/grpc
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

2. Compile the Protocol Buffers:

   For Python:
   ```bash
   python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. embedding.proto
   ```

   For Go:
   You need to download `protoc` and set the path.
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative embedding.proto
   ```

3. Start the server:
   ```bash
   python server.py
   ```

4. Run the client:
   ```bash
   go run client.go path/to/image.jpg
   ```