# tutorial-grpc-rest-api-go

Running tutorial in create api rest over grpc out of the box

Based in:

- https://www.tabnews.com.br/rschio/tutorial-rest-api-com-go-e-grpc
- https://blog.logrocket.com/guide-to-grpc-gateway/#using-grpc-gateway-with-gin

Out last version:

Client:

```zsh
$ grpcurl -plaintext -d '{"name": "foo", "price": 10.20}'  localhost:8081 product.v1.ProductService/AddProduct
{
  "productId": "3f4fbf05-e83e-42e3-b832-48c7e5ae5195"
}
```

Server:

```zsh
$ make runcomplete
2024/03/18 01:43:06 http[200]-- 630.728µs -- /api/v1/list
```
