# Connect API
- [upbit api docs](https://docs.upbit.com)

## Quotation API

- 시세 종목 조회
  - [x] 종목 코드 조회
    - url: `/market/all`
    - method: `GET`
- 시세 캔들 조회
  - [x] 분(Minute) 캔들
    - url: `https://api.upbit.com/v1/candles/minutes/{unit}
    - method: `GET`
  - [x] 일(Day) 캔들
    - url: `https://api.upbit.com/v1/candles/days`
    - method: `GET`
  - [x] 주(Week) 캔들
    - url: `https://api.upbit.com/v1/candles/weeks`
    - method: `GET`
  - [x] 월(Month) 캔들
    - url: `https://api.upbit.com/v1/candles/months`
    - method: `GET`
- 시세 체결 조회
  - [x] 최근 체결 내역
    - url: `https://api.upbit.com/v1/trades/ticks`
    - method: `GET`
- 시세 현재가(Ticker) 조회
  - [x] 종목 단위 현재가 정보
    - url: `https://api.upbit.com/v1/ticker`
    - method: `GET`
  - [x] 마켓 단위 현재가 정보
    - url: `https://api.upbit.com/v1/ticker/all`
    - method: `GET`
- 시세 호가 정보(Orderbook) 조회
  - [x] 호가 정보 조회
    - url: `https://api.upbit.com/v1/orderbook`
    - method: `GET`
  - [x] 호가 모아보기 단위 정보 조회
    - url: `https://api.upbit.com/v1/orderbook/supported_levels`
    - method: `GET`
