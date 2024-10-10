package public

import (
	"context"
	"testing"
)

const (
	testUrl = "https://api.upbit.com"
	testVer = "/v1"
)

func testClient() *Client {
	return NewClient(Config{
		BaseUrl: testUrl,
		Version: testVer,
	})

}

func TestClient_GetMarkets(t *testing.T) {
	client := testClient()

	actual, err := client.GetMarkets(context.Background(), false)

	if err != nil {
		t.Errorf("GetMarkets() error = %v", err)
	}

	if len(actual) == 0 {
		t.Errorf("GetMarkets() got empty response")
	}
}

func TestClient_GetCandles(t *testing.T) {
	client := testClient()

	tests := []struct {
		name string
		req  CandleRequest
	}{
		{
			name: "GetCandles() with Minute",
			req: CandleRequest{
				Market:         "KRW-BTC",
				CandleInterval: Minute,
				UnitCount:      60,
				Count:          10,
			},
		},
		{
			name: "GetCandles() with Day",
			req: CandleRequest{
				Market:         "KRW-BTC",
				CandleInterval: Day,
				Count:          10,
			},
		},
		{
			name: "GetCandles() with Week",
			req: CandleRequest{
				Market:         "KRW-BTC",
				CandleInterval: Week,
				Count:          10,
			},
		},
		{
			name: "GetCandles() with Month",
			req: CandleRequest{
				Market:         "KRW-BTC",
				CandleInterval: Month,
				Count:          10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := client.GetCandles(context.Background(), tt.req)
			if err != nil {
				t.Errorf("GetCandles() error = %v", err)
			}
			if len(actual) == 0 {
				t.Errorf("GetCandles() got empty response")
			}
		})
	}
}

func TestClient_GetTradeTicks(t *testing.T) {
	client := testClient()

	tests := []struct {
		name string
		req  TradeTicksRequest
	}{
		{
			name: "GetTradeTicks()",
			req: TradeTicksRequest{
				Market:  "KRW-BTC",
				Count:   10,
				DaysAgo: 3,
				//Cursor:  "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := client.GetTradeTicks(context.Background(), tt.req)
			if err != nil {
				t.Errorf("GetTradeTicks() error = %v", err)
			}
			if len(actual) == 0 {
				t.Errorf("GetTradeTicks() got empty response")
			}
		})
	}
}

func TestClient_GetTickerPrice(t *testing.T) {
	client := testClient()

	tests := []struct {
		name    string
		markets []string
	}{
		{
			name:    "GetTickerPrice()",
			markets: []string{"KRW-BTC, BTC-ETH"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := client.GetTickerPrice(context.Background(), tt.markets)
			if err != nil {
				t.Errorf("GetTickerPrice() error = %v", err)
			}
			if len(actual) == 0 {
				t.Errorf("GetTickerPrice() got empty response")
			}
		})
	}
}

func TestClient_GetAllTickerPrices(t *testing.T) {
	client := testClient()

	tests := []struct {
		name            string
		quoteCurrencies []QuoteCurrency
	}{
		{
			name:            "GetAllTickerPrices()",
			quoteCurrencies: []QuoteCurrency{KRW, BTC, USDT},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := client.GetAllTickerPrices(context.Background(), tt.quoteCurrencies)
			if err != nil {
				t.Errorf("GetAllTickerPrices() error = %v", err)
			}
			if len(actual) == 0 {
				t.Errorf("GetAllTickerPrices() got empty response")
			}
		})
	}
}

func TestClient_GetOrderBook(t *testing.T) {
	client := testClient()

	tests := []struct {
		name    string
		markets []string
		level   float64
	}{
		{
			name:    "GetOrderBook()",
			markets: []string{"KRW-BTC, BTC-ETH"},
			level:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := client.GetOrderBook(context.Background(), tt.markets, tt.level)
			if err != nil {
				t.Errorf("GetOrderBook() error = %v", err)
			}
			if len(actual) == 0 {
				t.Errorf("GetOrderBook() got empty response")
			}
		})
	}
}

func TestClient_GetOrderBookSupportedLevels(t *testing.T) {
	client := testClient()

	tests := []struct {
		name    string
		markets []string
	}{
		{
			name:    "GetOrderBookSupportedLevels()",
			markets: []string{"KRW-BTC, BTC-ETH"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := client.GetOrderBookSupportedLevels(context.Background(), tt.markets)
			if err != nil {
				t.Errorf("GetOrderBookSupportedLevels() error = %v", err)
			}
			if len(actual) == 0 {
				t.Errorf("GetOrderBookSupportedLevels() got empty response")
			}
		})
	}
}
