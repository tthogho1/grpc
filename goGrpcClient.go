package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <image_file>", os.Args[0])
	}

	// サーバーへの接続を確立
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEmbeddingServiceClient(conn)

	// 画像ファイルを読み込む
	imageData, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to read image file: %v", err)
	}

	// コンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// リクエストを作成して送信
	req := &pb.ImageRequest{
		ImageData: imageData,
		Filename:  os.Args[1],
	}

	resp, err := client.GetEmbedding(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get embedding: %v", err)
	}

	if !resp.Success {
		log.Fatalf("Server error: %s", resp.Error)
	}

	fmt.Printf("Successfully got embeddings of length: %d\n", len(resp.Embeddings))
	if len(resp.Embeddings) >= 5 {
		fmt.Println("First 5 embedding values:", resp.Embeddings[:5])
	}
}
