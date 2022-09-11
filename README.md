## Run application in dev mode

```bash
docker-compose -f ./ops/docker/docker-compose.yaml -f docker-compose.override.yaml up -d
```

## Example requests

- **GetBalance**
```
POST /mascot/seamless HTTP/1.1
Host: localhost:8080
Content-Length: 126
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"jsonrpc":"2.0","method":"rpc.GetBalance","params":{"callerId":1,"playerName":"player1","currency":"EUR","gameId":"riot"},"id":0}

---

HTTP/1.1 200 OK
Connection: keep-alive
Content-Type: application/json
Date: Mon, 22 Aug 2022 00:00:00 GMT
Vary: Accept-Encoding

{"jsonrpc":"2.0","id":0,"result":{"balance":10000}} // 100 EUR
```

- **WithdrawAndDeposit** 
```
POST /mascot/seamless HTTP/1.1
Host: localhost:8080
Content-Length: 339
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"jsonrpc":"2.0","method":"rpc.WithdrawAndDeposit","params":{"callerId":1,"playerName":"player1","withdraw":400,"deposit":200,"currency":"EUR","transactionRef":"1:UOwGgNHPgq3OkqRE","gameRoundRef":"1wawxl:39","gameId":"riot","reason":"GAME_PLAY_FINAL","sessionId":"qx9sgvvpihtrlug","spinDetails":{"betType":"spin","winType":"standart"}},"id":0}

---

HTTP/1.1 200 OK
Connection: keep-alive
Content-Type: application/json
Date: Mon, 22 Aug 2022 00:00:00 GMT
Vary: Accept-Encoding

{"jsonrpc":"2.0","id":0,"result":{"newBalance":9800,"transactionId":"1413628395"}}
```

- **RollbackTransaction**
```
POST /mascot/seamless HTTP/1.1
Host: localhost:8080
Content-Length: 213
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"jsonrpc":"2.0","method":"rpc.RollbackTransaction","params":{"callerId":1,"playerName":"player1","transactionRef":"1:UOwGgNHPgq3OkqRE","gameId":"riot","sessionId":"qx9sgvvpihtrlug","gameRoundRef":"1wawxl:39"},"id":0}

---

HTTP/1.1 200 OK
Connection: keep-alive
Content-Type: application/json
Date: Mon, 22 Aug 2022 00:00:00 GMT
Vary: Accept-Encoding

{"jsonrpc":"2.0","id":0,"result":null} // success response
```
