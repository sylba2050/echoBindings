# echoBindings

echoのbindの動作確認

# 実行結果

## 全ての要素

### request

```bash
curl -D - -X POST http://153.126.139.150:8080/ -H 'Content-Type: application/json' -d '{"userid":"mstn","pw":"hoge"}'
```

or

```bash
curl -D - -G http://153.126.139.150:8080 -d userid=mstn -d pw=hoge
```

### response

```text
HTTP/1.1 200 OK
Date: Tue, 14 May 2019 02:41:40 GMT
Content-Length: 0
```

### server side

```text
mstn
```

## 過剰な要素

### request

```bash
curl -D - -X POST http://153.126.139.150:8080/ -H 'Content-Type: application/json' -d '{"userid":"mstn","pw":"hoge", "auth_code":"hoge"}'
```
or

```bash
curl -D - -G http://153.126.139.150:8080 -d userid=mstn -d pw=hoge -d auth_code=hoge
```

### response

```text
HTTP/1.1 200 OK
Date: Tue, 14 May 2019 02:46:35 GMT
Content-Length: 0
```

### server side

```text
mstn
```

## 要素が不足

### request

```bash
curl -D - -X POST http://153.126.139.150:8080/ -H 'Content-Type: application/json' -d '{"userid":"mstn"}'
```

or

```bash
curl -D - -G http://153.126.139.150:8080 -d userid=mstn
```

### response

```text
HTTP/1.1 200 OK
Date: Tue, 14 May 2019 02:47:40 GMT
Content-Length: 0
```

### server side

```text
mstn
```
