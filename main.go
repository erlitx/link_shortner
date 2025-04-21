package main

import (
	"time"

	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/httpclient"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/wb"
)

func main() {
	// for i := 0; i < 1000; i++ {
	// go httpclient.Example()
	// }
	// time.Sleep(time.Second * 5)

	wb := wb.New("https://statistics-api.wildberries.ru/api/v1/supplier/orders?dateFrom=2025-02-20&flag=0")

	
}