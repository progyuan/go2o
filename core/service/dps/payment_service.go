/**
 * Copyright 2015 @ z3q.net.
 * name : payment_service.go
 * author : jarryliu
 * date : 2016-07-03 13:24
 * description :
 * history :
 */
package dps

import (
	"go2o/core/domain/interface/order"
	"go2o/core/domain/interface/payment"
)

type paymentService struct {
	_rep      payment.IPaymentRep
	_orderRep order.IOrderRep
}

func NewPaymentService(rep payment.IPaymentRep, orderRep order.IOrderRep) *paymentService {
	return &paymentService{
		_rep:      rep,
		_orderRep: orderRep,
	}
}

// 根据编号获取支付单
func (p *paymentService) GetPaymentOrder(id int) *payment.PaymentOrder {
	po := p._rep.GetPaymentOrder(id)
	if po != nil {
		v := po.GetValue()
		return &v
	}
	return nil
}

// 根据支付单号获取支付单
func (p *paymentService) GetPaymentOrderByNo(paymentNo string) *payment.PaymentOrder {
	if v := p._rep.GetPaymentOrderByNo(paymentNo); v != nil {
		v2 := v.GetValue()
		return &v2
	}
	return nil
}

// 创建支付单
func (p *paymentService) CreatePaymentOrder(v *payment.PaymentOrder) (int, error) {
	o := p._rep.CreatePaymentOrder(v)
	return o.Commit()
}

// 调整支付单金额
func (p *paymentService) AdjustOrder(paymentNo string, amount float32) error {
	o := p._rep.GetPaymentOrderByNo(paymentNo)
	if o == nil {
		return payment.ErrNoSuchPaymentOrder
	}
	return o.Adjust(amount)
}

func (p *paymentService) SetPrefixOfTradeNo(id int, prefix string) error {
	o := p._rep.GetPaymentOrder(id)
	if o == nil {
		return payment.ErrNoSuchPaymentOrder
	}
	return o.TradeNoPrefix(prefix)
}

// 积分抵扣支付单
func (p *paymentService) IntegralDiscountForPaymentOrder(orderId int,
	integral int, ignoreOut bool) (float32, error) {
	o := p._rep.GetPaymentOrder(orderId)
	if o == nil {
		return 0, payment.ErrNoSuchPaymentOrder
	}
	return o.IntegralDiscount(integral, ignoreOut)
}

// 余额抵扣
func (p *paymentService) BalanceDiscountForPaymentOrder(orderId int, remark string) error {
	o := p._rep.GetPaymentOrder(orderId)
	if o == nil {
		return payment.ErrNoSuchPaymentOrder
	}
	err := o.BalanceDiscount(remark)
	if err == nil {
		_, err = o.Commit()
	}
	return err
}

// 赠送账户支付
func (p *paymentService) PresentAccountPayment(orderId int, remark string) error {
	o := p._rep.GetPaymentOrder(orderId)
	if o == nil {
		return payment.ErrNoSuchPaymentOrder
	}
	return o.PresentAccountPayment(remark)
}

// 完成支付单支付，并传入支付方式及外部订单号
func (p *paymentService) FinishPayment(tradeNo string, spName string,
	outerNo string) error {
	o := p._rep.GetPaymentOrderByNo(tradeNo)
	if o == nil {
		return payment.ErrNoSuchPaymentOrder
	}
	err := o.PaymentFinish(spName, outerNo)
	if err == nil {
		//更改订单支付完成
		if orderId := o.GetValue().OrderId; orderId > 0 {
			m := p._orderRep.Manager()
			err = m.PaymentForOnlineTrade(o.GetValue().OrderId)
		}
	}
	return err
}
