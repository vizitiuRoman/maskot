## Run application in dev mode

```bash
$ docker-compose -f ./ops/docker/docker-compose.yaml -f docker-compose.override.yaml up -d
```

And if you see the following output then services is started ok.
```bash
Running 3/3
  Container docker-db-1      Healthy                                                                                                                6.3s
  Container docker-app-1     Started                                                                                                                0.8s
  Container docker-flyway-1  Started
```

## Some other commands

```bash
$ docker-compose ps

CONTAINER ID   IMAGE           COMMAND                  CREATED          STATUS                    PORTS                                                                                    NAMES
c380f007dec7   postgres:13.3   "docker-entrypoint.s…"   34 seconds ago   Up 33 seconds (healthy)   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp                                                docker-db-1
7970b55a9eb8   docker-app      "/bin/sh -c /entrypo…"   34 seconds ago   Up 33 seconds             0.0.0.0:8080->8080/tcp, :::8080->8080/tcp, 0.0.0.0:49153->3000/tcp, :::49153->3000/tcp   docker-app-1

$ docker logs docker-app-1

2022/10/04 14:27:25 Server started and listening on port: 8080
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

{"jsonrpc":"2.0","method":"getBalance","params":{"callerId":1,"playerName":"player1","currency":"EUR","gameId":"riot"},"id":0}

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

{"jsonrpc":"2.0","method":"withdrawAndDeposit","params":{"callerId":1,"playerName":"player1","withdraw":400,"deposit":200,"currency":"EUR","transactionRef":"1:UOwGgNHPgq3OkqRE","gameRoundRef":"1wawxl:39","gameId":"riot","reason":"GAME_PLAY_FINAL","sessionId":"qx9sgvvpihtrlug","spinDetails":{"betType":"spin","winType":"standart"}},"id":0}

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

{"jsonrpc":"2.0","method":"rollbackTransaction","params":{"callerId":1,"playerName":"player1","transactionRef":"1:UOwGgNHPgq3OkqRE","gameId":"riot","sessionId":"qx9sgvvpihtrlug","gameRoundRef":"1wawxl:39"},"id":0}

---

HTTP/1.1 200 OK
Connection: keep-alive
Content-Type: application/json
Date: Mon, 22 Aug 2022 00:00:00 GMT
Vary: Accept-Encoding

{"jsonrpc":"2.0","id":0,"result":null} // success response
```
