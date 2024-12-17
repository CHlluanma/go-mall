package service

import (
	"context"

	"github.com/chhz0/go-mall-kitex/app/checkout/infra/mq"
	"github.com/chhz0/go-mall-kitex/app/checkout/infra/rpc"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/cart"
	checkout "github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/checkout"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/email"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/order"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/payment"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResp == nil || cartResp.Items == nil || len(cartResp.Items) == 0 {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var oi []*order.OrderItem
	var total float32

	for _, carItem := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: carItem.ProductId})
		if err != nil {
			klog.Error(err)
			return nil, err
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(carItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: carItem.ProductId,
				Quantity:  carItem.Quantity,
			},
			Cost: cost,
		})
	}

	orderReq := &order.PlaceOrderReq{
		UserId: req.UserId,
		Items:  oi,
		Email:  req.Email,
	}

	if req.Address != nil {
		orderReq.Address = &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		}
	}

	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		klog.Error(err.Error())
		return nil, err
	}
	// empty cart
	emptyCartResp, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
		return nil, err
	}

	klog.Info(emptyCartResp)

	var orderId string
	if orderResp != nil || orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	paymentResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@eamil.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "Order Confirmation",
		Content:     "Your order has been confirmed",
	})

	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResp)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResp.TransactionId,
	}

	return
}
