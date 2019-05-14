# echoBindings

echoのbindの動作確認

# 実行結果

## 全ての要素をPOST

```bash
curl -D - -X POST http://153.126.139.150:8080/ -H 'Content-Type: application/json' -d '{"userid":"mstn","pw":"hoge"}'
```

```text
HTTP/1.1 200 OK
Date: Tue, 14 May 2019 02:41:40 GMT
Content-Length: 0
```

## 過剰な要素をポスト

```bash
curl -D - -X POST http://153.126.139.150:8080/ -H 'Content-Type: application/json' -d '{"userid":"mstn","pw":"hoge", "auth_code":"hoge"}'
```

```text
HTTP/1.1 200 OK
Date: Tue, 14 May 2019 02:46:35 GMT
Content-Length: 0
```

## 要素が不足しているパターン

```bash
curl -D - -X POST http://153.126.139.150:8080/ -H 'Content-Type: application/json' -d '{"userid":"mstn"}'
```

```text
HTTP/1.1 200 OK
Date: Tue, 14 May 2019 02:47:40 GMT
Content-Length: 0
```
