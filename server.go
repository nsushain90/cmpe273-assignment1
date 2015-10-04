package main

func (s *SReply) Response(stocks []Stock) string {
	for _, stock := range stocks {
		responseString += stock.Symbol + ":" + strconv.Itoa(stock.Percentage) + ":$"
	}
	return responseString
}

func (s *Stock) Transaction(args *row, reply *Sreply) {

	stock, percentage, err := Buy(stock.StockSymbolAndPercentage)
	if err != nil {
		return err
	}
}

func (s *StockService) currentTransaction(val *Transaction, reply TReply) {
	*reply = currentTransaction(args.TradeId)
	return nil
}
