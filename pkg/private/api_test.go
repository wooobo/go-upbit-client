package private

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testUrl   = "https://api.upbit.com"
	testVer   = "/v1"
	accessKey = ""
	secretKey = ""
)

func testClient() *Client {
	return NewClient(Config{
		BaseUrl:      testUrl,
		Version:      testVer,
		PublicApiKey: accessKey,
		SecretApiKey: secretKey,
	})

}

func TestClient_GetAccounts(t *testing.T) {
	client := testClient()

	actual, err := client.GetAccounts(context.Background())

	if err != nil {
		t.Errorf("GetAccounts() error = %v", err)
	}

	if len(actual) == 0 {
		t.Errorf("GetAccounts() got empty response")
	}
}

func TestClient_PlaceOrder(t *testing.T) {
	client := testClient()

	tests := []struct {
		name string
		req  PlaceOrderRequest
	}{
		{
			name: "PlaceOrder() with Limit",
			req: PlaceOrderRequest{
				Market:      "KRW-BTC",
				Side:        OrderSideBid,
				Volume:      "0.0001",
				Price:       "80000000",
				OrdType:     "limit",
				Identifier:  "",
				TimeInForce: "",
			},
		},
		{
			name: "PlaceOrder() with Limit",
			req: PlaceOrderRequest{
				Market:      "KRW-BTC",
				Side:        OrderSideAsk,
				Volume:      "0.0001",
				Price:       "99000000",
				OrdType:     "limit",
				Identifier:  "",
				TimeInForce: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := client.PlaceOrder(context.Background(), tt.req)

			if err != nil {
				t.Errorf("PlaceOrder() error = %v", err)
			}

			if actual.UUID == "" {
				t.Errorf("PlaceOrder() got empty response")
			}
		})
	}
}

func TestClient_FetchCancelOrder(t *testing.T) {
	client := testClient()

	placeOrder := PlaceOrderRequest{
		Market:      "KRW-BTC",
		Side:        OrderSideBid,
		Volume:      "0.0001",
		Price:       "80000000",
		OrdType:     "limit",
		Identifier:  "",
		TimeInForce: "",
	}

	order, _ := client.PlaceOrder(context.Background(), placeOrder)

	actual, err := client.CancelOrder(context.Background(), CancelOrderRequest{
		UUID:       order.UUID,
		Identifier: "",
	})

	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestClient_GetOrderChance(t *testing.T) {
	client := testClient()

	actual, err := client.GetOrderChance(context.Background(), "KRW-BTC")

	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestClient_GetFilledOrder(t *testing.T) {
	client := testClient()

	placeOrder := PlaceOrderRequest{
		Market:      "KRW-BTC",
		Side:        OrderSideBid,
		Volume:      "0.0001",
		Price:       "80000000",
		OrdType:     "limit",
		Identifier:  "",
		TimeInForce: "",
	}

	order, _ := client.PlaceOrder(context.Background(), placeOrder)

	actual, err := client.GetFilledOrder(context.Background(), order.UUID)

	assert.NoError(t, err)
	assert.NotNil(t, actual)

	_, err = client.CancelOrder(context.Background(), CancelOrderRequest{
		UUID:       order.UUID,
		Identifier: "",
	})
	if err != nil {
		t.Errorf("Failed to cancel order: %s", err.Error())
		return
	}
}

func TestClient_GetOrdersByIdentifier(t *testing.T) {
	client := testClient()

	placeOrder := PlaceOrderRequest{
		Market:  "KRW-BTC",
		Side:    OrderSideBid,
		Volume:  "0.0001",
		Price:   "80000000",
		OrdType: "limit",
		//Identifier:  "test",
		TimeInForce: "",
	}

	order, orderErr := client.PlaceOrder(context.Background(), placeOrder)
	if orderErr != nil {
		t.Errorf("Failed to place order: %s", orderErr.Error())
		return
	}

	actual, err := client.GetOrdersByIdentifier(context.Background(), OrderSearchRequest{
		Market: "KRW-BTC",
		UUIDs:  []string{order.UUID},
	})

	assert.NoError(t, err)
	assert.NotNil(t, actual)

	_, err = client.CancelOrder(context.Background(), CancelOrderRequest{
		UUID: order.UUID,
	})
}

func TestClient_GetOpenOrders(t *testing.T) {
	client := testClient()

	actual, err := client.GetOpenOrders(context.Background(), OrderQueryParams{
		Market: "KRW-BTC",
		//State:  StateWatch,
		States: []State{StateWait, StateWatch},
		Limit:  10,
	})

	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestClient_GetClosedOrders(t *testing.T) {
	client := testClient()

	loc, _ := time.LoadLocation("Asia/Seoul")
	startTime := time.Date(2024, 9, 19, 7, 0, 0, 0, loc)
	endTime := time.Date(2024, 9, 19, 8, 0, 0, 0, loc)

	actual, err := client.GetClosedOrder(context.Background(), CompletedOrderRequest{
		Market:    "KRW-BTC",
		StartTime: startTime.Format("2006-01-02T15:04:05Z"),
		EndTime:   endTime.Format("2006-01-02T15:04:05Z"),
		State:     StateCompletedOrderDone,
		Limit:     10,
	})

	assert.NoError(t, err)
	assert.NotNil(t, actual)
}
