package middleware

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	API_ENDPOINT      = "bcd797c2633d9e281fe6c0f11c7fa212.r2.cloudflarestorage.com"
	ACCESS_KEY_ID     = "8902998fd11e1e5232a7311462046333"
	SECRET_ACCESS_KEY = "9cff4eebc2b40c79cdcb23ecfe48138687990aef35dfa57e6109a266df2b4f1b"
)

func UploadFilesTrip(next http.HandlerFunc) http.HandlerFunc {
	// ctx := context.Background()

	// tangani metode upload multiple data dengan S3 disini
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Initialize minio client object.
		minioClient, err := minio.New(API_ENDPOINT, &minio.Options{
			Creds:  credentials.NewStaticV4(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ""),
			Secure: true,
		})
		if err != nil {
			log.Println("disini")
			log.Fatalln(err)
		}

		// tangkap kondisi multiple data form

		// Upload file
		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}
		//loleksi data kedalam array string
		var fileNames = []string{}

		// looping upload files
		for _, fileHeader := range r.MultipartForm.File["images"] {

			log.Println(r.MultipartForm.Value)

			rand.Seed(time.Now().UnixNano())

			// String
			charset := "abcdefghijklmnopqrstuvwxyz"
			// Getting random character
			c := charset[rand.Intn(len(charset))]
			c2 := charset[rand.Intn(len(charset))]
			c3 := charset[rand.Intn(len(charset))]
			c4 := charset[rand.Intn(len(charset))]
			c5 := charset[rand.Intn(len(charset))]

			xt := strings.Split(fileHeader.Filename, ".")[1]
			var sep = "."
			filename := (string(c) + string(c2) + string(c3) + string(c4) + string(c5) + sep + xt)

			file, _ := fileHeader.Open()

			// Upload file ke S3
			uploadInfo, err := minioClient.PutObject(context.Background(), "baharsah-s3", filename, file, fileHeader.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
			if err != nil {
				log.Println("Disini")
				log.Println(err)
				return
			}
			fileNames = append(fileNames, filename)
			log.Println("Successfully uploaded bytes: ", uploadInfo)

		}
		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFileNames", fileNames)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
