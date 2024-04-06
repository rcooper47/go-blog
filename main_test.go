package main

import (
	"fmt"
	"go-blog/inits"
	"go-blog/models"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func runTestServer() *httptest.Server {
	inits.LoadEnv()
	inits.ConnectToDb()
	inits.SyncDB()

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	configStrs1.db = db

	if err != nil {
		panic("Failed to connect to DB")
	}

	configStrs1.db.AutoMigrate(&models.Blog{})

	return httptest.NewServer(SetupRouter())
}

func Test_Post_Api_Integration_tests_store_endpoint(t *testing.T) {
	ts := runTestServer()

	defer ts.Close()

	t.Run("it should return 200 when health ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/blog", nil) //http.Get(fmt.Sprintf("%s/blog", "http://localhost:3000"))
		// if err != nil {
		// 	t.Fatalf("expected no error, got %v", err)
		// }
		// w = httptest.NewRecorder()
		// ts.POST()

		fmt.Println(req)
	})
}
