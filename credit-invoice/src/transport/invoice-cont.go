package transport

import (
	"devv-monteiro/go-digital-bank/credit-invoice/src/business"
	"encoding/json"
	"fmt"
	"net/http"
)

type InvoiceCont struct {
	serv *business.InvoiceServ
}

func NewInvoiceCont(serv *business.InvoiceServ) *InvoiceCont {
	return &InvoiceCont{serv: serv}
}

func (cont *InvoiceCont) GetCurrInvoice(resWr http.ResponseWriter, req *http.Request) {
	fmt.Println("Path: " + req.URL.Path)

	resWr.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodGet {
		resWr.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	req.ParseForm()
	customerId := req.Form.Get("customerId")
	fmt.Println("CustomerId: " + customerId)

	currentInvoice, err := cont.serv.GetCurrentInvoice(customerId)
	if err != nil {
		resWr.WriteHeader(err.StatusCode)
		json.NewEncoder(resWr).Encode(err)
		return
	}

	json.NewEncoder(resWr).Encode(currentInvoice)
}
