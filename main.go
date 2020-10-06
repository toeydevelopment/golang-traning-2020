package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Photos []Photo

// https://jsonplaceholder.typicode.com/photos
func main() {

	dirName := strconv.Itoa(int(time.Now().UnixNano()))

	if err := os.Mkdir(dirName, 0777); err != nil {
		panic(err)
	}

	var photos Photos

	if err := getJson(&photos); err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	for _, v := range photos {

		wg.Add(1)
		go func(v Photo) {
			defer wg.Done()
			imageByte, err := downloadImage(v.ThumbnailURL)

			if err != nil {
				panic(err)
			}

			if err := saveImage(fmt.Sprintf("%s/%d", dirName, v.ID), imageByte); err != nil {
				panic(err)
			}

		}(v)

	}

	wg.Wait()

	fmt.Println("DONE")

}

func saveImage(fileName string, img []byte) error {
	f, err := os.Create(fileName + ".png")

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = io.Copy(f, bytes.NewReader(img))

	if err != nil {
		return err
	}

	return nil

}

func downloadImage(urlImg string) ([]byte, error) {
	resp, err := http.Get(urlImg)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return bodyByte, nil
}

func getJson(photos *Photos) error {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(bodyByte, photos); err != nil {
		return err
	}

	return nil
}
