package public

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// GetMarkets 종목 코드 조회
// isDetails : 유의종목 필드과 같은 상세 정보 노출 여부 (default: false)
func (c *Client) GetMarkets(ctx context.Context, isDetails bool) ([]Market, error) {
	path := "/market/all"
	values := url.Values{}
	if isDetails {
		values.Set("isDetails", "true")
	}

	var markets []Market
	err := c.Get(ctx, path, values, &markets)

	if err != nil {
		return nil, err
	}

	return markets, nil
}

// GetCandles 캔들 조회
func (c *Client) GetCandles(ctx context.Context, req CandleRequest) ([]Candle, error) {
	path := ""
	switch req.CandleInterval {
	case Minute:
		path = fmt.Sprintf("%s%s", "/candles/minutes/", strconv.Itoa(req.UnitCount))
	case Day:
		path = "/candles/days"
	case Week:
		path = "/candles/weeks"
	case Month:
		path = "/candles/months"
	}

	values := url.Values{}
	values.Set("market", req.Market)
	values.Set("count", fmt.Sprintf("%d", req.Count))
	if req.To != "" {
		values.Set("to", req.To)
	}

	var candles []Candle
	err := c.Get(ctx, path, values, &candles)

	if err != nil {
		return nil, err
	}

	return candles, nil
}

// GetTradeTicks 파라미터 To 는 UTC 기준으로 조회해야함
func (c *Client) GetTradeTicks(ctx context.Context, req TradeTicksRequest) ([]TradeTick, error) {
	path := "/trades/ticks"
	values := url.Values{}
	values.Set("market", req.Market)
	values.Set("count", fmt.Sprintf("%d", req.Count))
	if req.To != "" {
		values.Set("to", req.To)
	}
	if req.Cursor != "" {
		values.Set("cursor", req.Cursor)
	}
	if req.DaysAgo != 0 {
		values.Set("daysAgo", fmt.Sprintf("%d", req.DaysAgo))
	}

	var tradeTicks []TradeTick
	err := c.Get(ctx, path, values, &tradeTicks)

	if err != nil {
		return nil, err
	}

	return tradeTicks, nil
}

// GetTickerPrice 종목 단위 현재가 정보 조회
// 반점으로 구분되는 종목 코드 (ex. KRW-BTC, BTC-ETH)
func (c *Client) GetTickerPrice(ctx context.Context, markets []string) ([]TickerSnapshot, error) {
	path := "/ticker"
	values := url.Values{}
	values.Set("markets", strings.Join(markets, ","))

	var resp []TickerSnapshot
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type QuoteCurrency string

const (
	KRW  QuoteCurrency = "KRW"
	BTC  QuoteCurrency = "BTC"
	USDT QuoteCurrency = "USDT"
)

// GetAllTickerPrices 마켓 단위 종목들의 스냅샷을 반환합니다.
func (c *Client) GetAllTickerPrices(ctx context.Context, quoteCurrencies []QuoteCurrency) ([]TickerSnapshot, error) {
	path := "/ticker/all"
	quoteCurrencyStrings := make([]string, len(quoteCurrencies))
	for i, qc := range quoteCurrencies {
		quoteCurrencyStrings[i] = string(qc)
	}

	// 문자열 슬라이스를 쉼표로 구분된 문자열로 변환
	quoteCurrenciesCsv := strings.Join(quoteCurrencyStrings, ", ")
	values := url.Values{}
	values.Set("quoteCurrencies", quoteCurrenciesCsv)

	var resp []TickerSnapshot
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOrderBook 호가 정보 조회
func (c *Client) GetOrderBook(ctx context.Context, markets []string, level float64) ([]OrderBook, error) {
	path := "/orderbook"
	values := url.Values{}
	values.Set("markets", strings.Join(markets, ", "))
	if level != 0 {
		values.Set("level", strconv.FormatFloat(level, 'f', -1, 64))
	}

	var resp []OrderBook
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOrderBookSupportedLevels 호가 모아보기 단위 정보 조회
// 호가 모아보기 기능은 원화마켓(KRW)에서만 지원하므로 BTC, USDT 마켓의 경우 0만 존재합니다.
func (c *Client) GetOrderBookSupportedLevels(ctx context.Context, markets []string) ([]SupportedLevels, error) {
	path := "/orderbook/supported_levels"
	values := url.Values{}
	values.Set("markets", strings.Join(markets, ", "))

	var resp []SupportedLevels
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
