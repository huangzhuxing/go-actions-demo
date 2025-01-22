package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect database: %v", err)
	}
	db.AutoMigrate(&User{})
}

func TestCreateUser(t *testing.T) {
	setupTestDB(t)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/users", createUser)

	// 创建测试用户数据
	user := User{
		Username: "testuser",
		Email:    "test@example.com",
	}
	jsonValue, _ := json.Marshal(user)

	// 创建测试请求
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// 检查响应状态码
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// 解析响应
	var response User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// 验证响应数据
	if response.Username != user.Username {
		t.Errorf("Expected username %s, got %s", user.Username, response.Username)
	}
	if response.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, response.Email)
	}
}
