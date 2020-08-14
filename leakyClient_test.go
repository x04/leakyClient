package leakyClient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

func TestLeakyClient_Get(t *testing.T) {
	client := New(4)
	wg := sync.WaitGroup{}
	wg.Add(2)

	start := time.Now()

	go func() {
		defer wg.Done()
		resp, err := client.Get("https://httpbin.org/get")
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	go func() {
		defer wg.Done()
		resp, err := client.Get("https://httpbin.org/get")
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	wg.Wait()

	if time.Now().Sub(start).Milliseconds() < 250 {
		t.Error("Time to complete both requests was <250ms, should have been >250ms.")
	}
}

func TestLeakyClient_Post(t *testing.T) {
	client := New(4)
	wg := sync.WaitGroup{}
	wg.Add(2)

	start := time.Now()

	jsonData, err := json.Marshal(map[string]string{"leakyClient": "testPost"})
	if err != nil {
		t.Error("Failed to marshal json data:", err)
	}

	go func() {
		defer wg.Done()
		resp, err := client.Post("https://httpbin.org/post", "application/json", bytes.NewReader(jsonData))
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	go func() {
		defer wg.Done()
		resp, err := client.Post("https://httpbin.org/post", "application/json", bytes.NewReader(jsonData))
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	wg.Wait()

	if time.Now().Sub(start).Milliseconds() < 250 {
		t.Error("Time to complete both requests was <250ms, should have been >250ms.")
	}
}

func TestLeakyClient_PostData(t *testing.T) {
	client := New(4)
	wg := sync.WaitGroup{}
	wg.Add(2)

	start := time.Now()

	values := url.Values{}
	values.Add("leakyClient", "testPost")

	go func() {
		defer wg.Done()
		resp, err := client.PostForm("https://httpbin.org/post", values)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	go func() {
		defer wg.Done()
		resp, err := client.PostForm("https://httpbin.org/post", values)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	wg.Wait()

	if time.Now().Sub(start).Milliseconds() < 250 {
		t.Error("Time to complete both requests was <250ms, should have been >250ms.")
	}
}

func TestLeakyClient_Do(t *testing.T) {
	client := New(4)
	wg := sync.WaitGroup{}
	wg.Add(2)

	start := time.Now()

	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		t.Error("Failed to create request:", err)
	}

	go func() {
		defer wg.Done()
		values := url.Values{}
		values.Add("leakyClient", "testPost")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	go func() {
		defer wg.Done()
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal("Unexpected response code:", resp.StatusCode)
		}
	}()

	wg.Wait()

	if time.Now().Sub(start).Milliseconds() < 250 {
		t.Error("Time to complete both requests was <250ms, should have been >250ms.")
	}
}
