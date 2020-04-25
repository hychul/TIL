package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadTest(t *testing.T) {
	path := "./test-resources/dduky.jpg"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	if err != nil {
		t.Fatal("Failed!: create form file error", err)
	}
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("content-type", writer.FormDataContentType())

	uploadHandler(res, req)
	if res.Code != http.StatusOK {
		t.Fatal("Failed!: response code is not equal", res.Code)
	}

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	if err != nil {
		t.Fatal("Failed!: uploaded file has error", err)
	}

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	if bytes.Compare(uploadData, originData) != 0 {
		t.Fatal("Failed!: uploaded file is not equal with origin", uploadData)
	}
}
