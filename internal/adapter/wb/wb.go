package wbadapter

import "gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/wb"

type API struct {
	Host   string // basic WB URL
	Orders string // api path for orders
	Sales  string // api path for sales
}

type WBConn struct {
	Client *wb.Client
	API    API
}

func New(c *wb.Client) *WBConn {
	api := API{
		Orders: orders,
		Sales:  sales,
		Host:   host,
	}

	return &WBConn{
		Client: c,
		API:    api,
	}
}
