package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func generateRandomString(length int) string {
	// Create a slice of runes with the given length.
	runes := make([]rune, length)

	// Fill the slice with random runes.
	for i := range runes {
		runes[i] = randomRune()
	}

	// Return the slice as a string.
	return string(runes)
}

// randomRune returns a random printable ASCII rune.
func randomRune() rune {
	// Choose a random number between 48 ('0') and 122 ('z').
	n := rand.Intn(74) + 48

	// If the number is between 58 and 64, add 7 to make it a uppercase letter.
	if n >= 58 && n <= 64 {
		n += 7
	}

	// Convert the number to a rune.
	return rune(n)
}

func UploadFilesTrip(next http.HandlerFunc) http.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		API_ENDPOINT      = os.Getenv("API_ENDPOINT")
		ACCESS_KEY_ID     = os.Getenv("ACCESS_KEY_ID")
		SECRET_ACCESS_KEY = os.Getenv("SECRET_ACCESS_KEY")
	)

	// tangani metode upload multiple data dengan S3 disini
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(API_ENDPOINT)
		log.Println(ACCESS_KEY_ID)
		log.Println(SECRET_ACCESS_KEY)
		ctx := r.Context()
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

			// log.Println(r.MultipartForm.Value)

			// Seed the random number generator with the current time.
			rand.Seed(time.Now().UnixNano())

			// Generate a random string of 10 characters.
			// randomString := generateRandomString(10)
			// Generate a random alphanumeric string with a length of 10
			b := make([]byte, 10)
			rand.Read(b)
			randomString := fmt.Sprintf("%x", b)

			// String
			// charset := "abcdefghijklmnopqrstuvwxyz"
			// Getting random character

			xt := strings.Split(fileHeader.Filename, ".")[1]
			var sep = "."
			filename := ("image-" + randomString + sep + xt)

			file, _ := fileHeader.Open()

			// Upload file ke S3
			_, err := minioClient.PutObject(ctx, os.Getenv("BUCKET"), filename, file, fileHeader.Size, minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")})
			if err != nil {
				log.Println("Disini")
				log.Println(err)
				return
			}
			fileNames = append(fileNames, filename)
			log.Println("Successfully uploaded bytes")

		}
		// log.Println("data", fileNames)
		//loop to image struct

		// add filename to ctx
		newCtx := context.WithValue(ctx, "file", fileNames)

		// fmt.Println("ini data ref", ref)

		// log.Println("read context here", newCtx)

		next.ServeHTTP(w, r.WithContext(newCtx))
	})

}
