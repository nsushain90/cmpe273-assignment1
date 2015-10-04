package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type StockandSymbol struct {
	StockSymbolAndPercentage string
	Budget                   float64
}

type SReply struct {
	TradeId        int
	Stocks         string
	UnvestedAmount float64
}

func Transaction(stock string, budget float64) SReply {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	defer client.Close()

	args := records{Stock, budget}
	var reply StockReply

	err = client.Call("Stock.Transaction", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}

func currentTransaction(transactionId int) TReply {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	defer client.Close()

	val := TransactionArgs{transactionId}
	var reply TransactionReply
	err = client.Call("currentTransaction.Stock", val, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}

func main() {
	url := "http://finance.yahoo.com/d/quotes.csv?s=AAPL+GOOG&f=nsb"
	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}
	//fmt.Scanf("Stock:%s:Percentage:%s\n", row[0], row[1])

	r := mux.NewRouter()

	s := rpc.NewServer()
	stock := new(stock)
	s.RegisterService(stock, "")
	s.RegisterCodec(json.NewCodec(), "application/json")
	stock := new(Stock)
	s.RegisterService(stock, "")

	chain := alice.New(
		func(h http.Handler) http.Handler {
			return handlers.CombinedLoggingHandler(os.Stdout, h)
		},
		handlers.CompressHandler,
		func(h http.Handler) http.Handler {
			return recovery.Handler(os.Stderr, h, true)
		})

	r.Handle("/rpc", chain.Then(s))
	log.Fatal(http.ListenAndServe(":8080", r))
	var records []Stock
	for _, row := range data {
		i, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Println("err")

		}
		records = append(records, Stock{row[0], i, row[1]})
		fmt.Println(records)
		val := StockandSymbol{"GOOG:100%", 1000}
		var reply StockReply
		err = client.Call("Stock.Transaction", val, &reply)
		if err != nil {
			log.Fatal("error:", err)
		}
		fmt.Println(reply)
	}

}
