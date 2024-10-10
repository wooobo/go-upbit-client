package private

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// GetAccounts 전체 계좌 조회
func (c *Client) GetAccounts(ctx context.Context) ([]Account, error) {
	path := "/accounts"

	var resp []Account
	err := c.Get(ctx, path, nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOrderChance 마켓별 주문 가능 정보를 확인한다.
func (c *Client) GetOrderChance(ctx context.Context, market string) (OrderChance, error) {
	path := "/orders/chance"

	values := url.Values{}
	values.Set("market", market)

	var resp OrderChance
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return OrderChance{}, fmt.Errorf("failed to fetch accounts: %w", err)
	}

	return resp, nil
}

// GetFilledOrder 개별 주문 조회
func (c *Client) GetFilledOrder(ctx context.Context, uuid string) (FilledOrder, error) {
	path := "/order"
	values := url.Values{}
	values.Set("uuid", uuid)

	var resp FilledOrder
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return FilledOrder{}, err
	}

	return resp, nil
}

// GetOrdersByIdentifier 주문 UUID 로 주문 정보 조회
func (c *Client) GetOrdersByIdentifier(ctx context.Context, req OrderSearchRequest) ([]Order, error) {
	path := "/orders/uuids"
	values := url.Values{}
	values.Set("market", req.Market)
	values.Set("OrderBy", req.OrderBy)

	for _, uuid := range req.UUIDs {
		values.Add("uuids[]", uuid)
	}

	for _, identifier := range req.Identifiers {
		values.Add("identifiers[]", identifier)
	}

	var resp []Order
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOpenOrders 미체결 주문 조회
func (c *Client) GetOpenOrders(ctx context.Context, req OrderQueryParams) ([]Order, error) {
	path := "/orders/open"

	values := url.Values{}
	values.Set("market", req.Market)
	values.Set("limit", fmt.Sprint(req.Limit))

	if req.Limit > 0 {
		values.Set("limit", fmt.Sprint(req.Limit))
	}
	if req.Page > 0 {
		values.Set("page", fmt.Sprint(req.Page))
	}
	if req.OrderBy != "" {
		values.Set("order_by", req.OrderBy.String())
	}
	if req.States != nil {
		for _, state := range req.States {
			values.Add("states[]", state.String())
		}
	} else {
		values.Add("state", req.State.String())
	}

	var resp []Order
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetClosedOrder 완료 주문 조회
// start_time 과 end_time 은 Time Zone 이 포함된 ISO-8601 포맷(ex. 2024-03-13T00:00:00+09:00) 이어야 합니다.
// start_time, end_time 이 둘 다 정의되지 않은 경우 현재 시각으로부터 최대 1시간 전까지의 주문이 조회됩니다.
// start_time 만 정의한 경우 start_time 으로부터 최대 1시간 후까지의 주문이 조회됩니다.
// end_time 만 정의한 경우 end_time 으로부터 최대 1시간 전까지의 주문이 조회됩니다.
// start_time 과 end_time 을 둘 다 정의할 경우 최대 1시간 범위까지의 주문만 조회가 가능합니다.
// *조회 시간 내의 주문 건이라도 limit 개수를 초과한 범위일 경우 조회되지 않으니 이 경우 나누어서 조회하여야 합니다.
func (c *Client) GetClosedOrder(ctx context.Context, req CompletedOrderRequest) ([]Order, error) {
	path := "/orders/closed"

	values := url.Values{}
	values.Set("market", req.Market)
	values.Set("limit", strconv.Itoa(req.Limit))
	values.Set("order_by", req.OrderBy.String())

	if req.States != nil {
		for _, state := range req.States {
			values.Add("states[]", state.String())
		}
	} else {
		values.Add("state", req.State.String())
	}

	if req.StartTime != "" {
		values.Set("start_time", req.StartTime)
	}

	if req.EndTime != "" {
		values.Set("end_time", req.EndTime)
	}

	var resp []Order
	err := c.Get(ctx, path, values, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PlaceOrder 주문하기
func (c *Client) PlaceOrder(ctx context.Context, order PlaceOrderRequest) (PlaceOrder, error) {
	path := "/orders"

	values := url.Values{}
	values.Set("market", order.Market)
	values.Set("side", order.Side.String())
	values.Set("volume", order.Volume)
	values.Set("price", order.Price)
	values.Set("ord_type", order.OrdType)

	if order.TimeInForce != "" {
		values.Set("time_in_force", order.TimeInForce)
	}
	if order.Identifier != "" {
		values.Set("identifier", order.Identifier)
	}

	var resp PlaceOrder
	err := c.Post(ctx, path, values, &resp)
	if err != nil {
		return PlaceOrder{}, fmt.Errorf("failed to place order: %w", err)
	}

	return resp, nil
}

// CancelOrder 주문 취소 접수
func (c *Client) CancelOrder(ctx context.Context, req CancelOrderRequest) (Order, error) {
	path := "/order"
	values := url.Values{}
	values.Set("uuid", req.UUID)
	values.Set("identifier", req.Identifier)

	var resp Order
	err := c.Delete(ctx, path, values, &resp)
	if err != nil {
		return Order{}, err
	}

	return resp, nil
}
