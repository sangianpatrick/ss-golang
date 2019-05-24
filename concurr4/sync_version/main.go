package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Order is struct
type Order struct {
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Filename    string  `json:"filename"`
	Height      int     `json:"height"`
	Width       int     `json:"widht"`
	Price       float32 `json:"price"`
	Rating      int     `json:"rating"`
}

// Orders contains list of order
type Orders []Order

// TotalPriceByType represent reduced price data of order.
type TotalPriceByType struct {
	Type       string  `json:"type"`
	TotalPrice float32 `json:"totalPrice"`
}

func main() {
	var orders Orders
	path, _ := filepath.Abs("./data.json")
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &orders)

	chanData := make(chan Orders)
	message := make(chan Orders)
	chanConsumer := make(chan int, 100)

	go broker(chanConsumer, chanData, message)
	go produce(orders, chanData)
	go consume(1, message, chanConsumer)
	// go consume(2, message, chanConsumer)
	time.Sleep(time.Minute * 2)
}

func broker(consumers chan int, data chan Orders, message chan Orders) {
	d := <-data
	<-consumers
	consumerSize := len(consumers)
	dataSize := len(d)
	perData := dataSize / consumerSize
	for i := 1; i <= consumerSize; i++ {
		message <- d[perData*i-perData : perData*i]
	}

}

func produce(data Orders, dataChan chan Orders) {
	dataChan <- data
}

func consume(id int, message chan Orders, chanConsumer chan int) {
	chanConsumer <- id
	close(chanConsumer)
	fmt.Printf("Consumer ke - %d: Received Data %d", id, len(<-message))
}
