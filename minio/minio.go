package main

import (
	"os"
	"log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"encoding/json"
	"context"
)
type minioCredentials struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Endpoint  string `json:"url"`
}
func minioi() {
	creadentialFile,err:=os.Open("credentials.json")
	if err!=nil{
		log.Fatal(err)
	}
	defer creadentialFile.Close()
	values,_:=io.ReadAll(creadentialFile)
	var mc minioCredentials
	json.Unmarshal(values,&mc)
	minioClient,err:=minio.New(mc.Endpoint,&minio.Options{
		Creds:  credentials.NewStaticV4(mc.AccessKey,mc.SecretKey,""),
        Secure: false,
	})
	if err!=nil{
		log.Fatal(err)
	}

	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	for _, bucket := range buckets {
		log.Println(bucket)
	}
	minioClient.MakeBucket(context.Background(),"go-bucket",minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully created mybucket.")
}