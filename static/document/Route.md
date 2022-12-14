# 项目启动

## USER

```shell
go run service/user/rpc/user.go -f service/user/rpc/etc/user.yaml
go run service/user/api/user.go -f service/user/api/etc/user.yaml
```

## PRODUCT

```shell
go run service/product/rpc/product.go -f service/product/rpc/etc/product.yaml
go run service/product/api/product.go -f service/product/api/etc/product.yaml
```

## ORDER

```shell
go run service/order/rpc/order.go -f service/order/rpc/etc/order.yaml
go run service/order/api/order.go -f service/order/api/etc/order.yaml
```

## PAY

```shell
go run service/pay/rpc/pay.go -f service/pay/rpc/etc/pay.yaml
go run service/pay/api/pay.go -f service/pay/api/etc/pay.yaml
```

# 路由请求

## USER

> /api/user/userinfo

![userinfo](../picture/userinfo.png)

## PRODUCT

> /api/product/create

![create](../picture/create.png)

> /api/product/update

![update](../picture/update.png)

> /api/product/remove

![remove](../picture/remove.png)

> /api/product/detail

![detail](../picture/detail.png)

## ORDER

> /api/order/create

![detail](../picture/order_create.png)

> /api/order/detail

![detail](../picture/order_detail.png)

> /api/order/update

![detail](../picture/order_update.png)

> /api/order/remove

![detail](../picture/order_remove.png)

> /api/order/list

![detail](../picture/order_list.png)

## PAY

> /api/pay/create

![detail](../picture/pay_create.png)

> /api/pay/detail

![detail](../picture/pay_detail.png)

> /api/pay/callback

![detail](../picture/pay_callback1.png)

![detail](../picture/pay_callback2.png)
