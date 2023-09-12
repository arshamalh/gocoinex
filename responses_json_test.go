package gocoinex

var JSONResp = map[string]string{
	"AllMarketList": `{"code":0,"data":["LTCBCH","ZECBCH","DASHBCH"],"message":"Ok"}`,
	"AllMarketInfo": `{"code":0,"message":"Ok","data":{"XRPBTC":{"taker_fee_rate":"0.001","pricing_name":"BTC","trading_name":"XRP","min_amount":"0.001","name":"XRPBTC","trading_decimal":8,"maker_fee_rate":"0.001","pricing_decimal":8},"CETUSDC":{"taker_fee_rate":"0.001","pricing_name":"USDC","trading_name":"CET","min_amount":"0.001","name":"CETUSDC","trading_decimal":8,"maker_fee_rate":"0.001","pricing_decimal":8}}}`,
}
