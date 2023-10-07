package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/hamdiBouhani/GopherNet-golang/dto"
	"github.com/hamdiBouhani/GopherNet-golang/mocks"
	"github.com/hamdiBouhani/GopherNet-golang/services"
	"github.com/hamdiBouhani/GopherNet-golang/storage/pg"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mockHttpService() *HttpService {

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	burrowServiceInstance := &services.BurrowService{Storage: &pg.DBConn{Db: db}}
	err = burrowServiceInstance.Storage.Drop()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	err = burrowServiceInstance.Storage.Migrate()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	svc := HttpService{
		BurrowServiceInstance: burrowServiceInstance,
	}

	svc.corsConfig = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	svc.devMode = os.Getenv("RUN_MODE") == "DEV"
	svc.testMode = os.Getenv("TEST_HTTP") == "TRUE"
	if svc.devMode {
		log.Println("-DEV_MODE enabled-")
	}

	err = svc.registerRoutes()
	if err != nil {
		fmt.Println(err)
		panic("failed to register routes")
	}

	return &svc
}

func TestPingAPI(t *testing.T) {
	httpmock := mockHttpService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ping", nil)
	httpmock.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestBurrowStatusAPI(t *testing.T) {
	httpmock := mockHttpService()

	for i := 0; i < 5; i++ {
		newBurrow := mocks.MockBurrow(false)
		err := httpmock.BurrowServiceInstance.Storage.CreateBurrow(newBurrow)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/burrows", nil)
	httpmock.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.IndexResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	burrows := resp.Results.([]interface{})
	if len(burrows) == 0 {
		t.Error("burrows not be empty")
		t.Fail()
		return
	}
}
