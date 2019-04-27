package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/**
查询所有goods测试用例
*/
func TestFindAllCheckHandler(t *testing.T) {
	//创建一个请求
	req, err := http.NewRequest("GET", "/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 我们创建一个 ResponseRecorder (which satisfies http.ResponseWriter)来记录响应
	rr := httptest.NewRecorder()

	//直接使用findAllHandler，传入参数rr,req
	findAllHandler(rr, req)

	// 检测返回的状态码
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("return body:", rr.Body.String())
}

/**
添加goods测试用例
*/
func TestAddCheckHandler(t *testing.T) {
	reqData := Goods{ID: "1000", Name: "bccccd", Stock: 100}

	reqBody, _ := json.Marshal(reqData)

	fmt.Println("input:", string(reqBody))

	req := httptest.NewRequest(
		http.MethodPost,
		"/add",
		bytes.NewReader(reqBody),
	)

	rr := httptest.NewRecorder()
	addHandler(rr, req)

	result := rr.Result()

	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("return body:", string(body))

	if result.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			result.StatusCode, http.StatusOK)
	}
}

/**
根据id查找good
*/
func TestFindHandler(t *testing.T) {
	// body中  5cc40a1349b8335e670758e9
	reader := strings.NewReader(`{"id": "5cc42b8949b83365f5d5f421","name": "yyyy","Stock": 1000}`)
	//创建一个请求
	req := httptest.NewRequest(http.MethodPost, "/get", reader)
	//接受请求
	rr := httptest.NewRecorder()
	findHandler(rr, req)
	result := rr.Result()

	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("return body:", string(body))

	if result.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			result.StatusCode, http.StatusOK)
	}
}
/**
更新good
 */
func TestUpdateHandler(t *testing.T) {
	// body中  5cc40a1349b8335e670758e9
	reader := strings.NewReader(`{"id": "5cc42b8949b83365f5d5f421","name": "yyyy","Stock": 1000}`)
	//创建一个请求
	req := httptest.NewRequest(http.MethodPost, "/update", reader)
	//接受请求
	rr := httptest.NewRecorder()
	updateHandler(rr, req)
	result := rr.Result()

	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("return body:", string(body))

	if result.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			result.StatusCode, http.StatusOK)
	}
}

/**
删除good
*/
func TestDeleteHandler(t *testing.T) {
	// body中  5cc40a1349b8335e670758e9
	reader := strings.NewReader(`{"id": "5cc40a1349b8335e670758e9","name": "yyyy","Stock": 1000}`)
	//创建一个请求
	req := httptest.NewRequest(http.MethodPost, "/update", reader)
	//接受请求
	rr := httptest.NewRecorder()
	deleteHandler(rr, req)
	result := rr.Result()

	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("return body:", string(body))

	if result.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			result.StatusCode, http.StatusOK)
	}
}
