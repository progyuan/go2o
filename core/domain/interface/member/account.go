/**
 * Copyright 2015 @ z3q.net.
 * name : account
 * author : jarryliu
 * date : 2015-07-24 08:48
 * description :
 * history :
 */
package member

const (
	// 余额账户
	AccountBalance = 1
	// 积分账户
	AccountIntegral = 2
	// 赠送账户
	AccountPresent = 3
	// 流通金账户
	AccountFlow = 4
)

const (
	// 用户充值
	ChargeByUser = 1
	// 系统自动充值
	ChargeBySystem = 2
	// 客服充值
	ChargeByService = 3
	// 退款充值
	ChargeByRefund = 4
)

const (
	// 会员充值
	KindBalanceCharge = 1
	// 系统充值
	KindBalanceSystemCharge = 2
	// 支付抵扣
	KindBalanceDiscount = 3
	// 退款
	KindBalanceRefund = 4
	// 转入
	KindBalanceTransferIn = 5
	// 转出
	KindBalanceTransferOut = 6
	// 失效
	KindBalanceExpired = 7
	// 冻结
	KindBalanceFreeze = 8
	// 解冻
	KindBalanceUnfreeze = 9

	// 客服充值
	KindBalanceServiceCharge = 15
	// 客服扣减
	KindBalanceServiceDiscount = 16
)

const (
	// 赠送金额
	KindPresentAdd = 1
	// 抵扣奖金
	KindPresentDiscount = 2
	// 转入
	KindPresentTransferIn = 5
	// 转出
	KindPresentTransferOut = 6
	// 失效
	KindPresentExpired = 7
	// 冻结
	KindPresentFreeze = 8
	// 解冻
	KindPresentUnfreeze = 9
	// 提现到余额
	KindPresentTakeOutToBalance = 11
	// 提现到银行卡(人工提现)
	KindPresentTakeOutToBankCard = 12
	// 提现到第三方
	KindPresentTakeOutToThirdPart = 13
	// 提现退还到银行卡
	KindPresentTakOutRefund = 14

	// 客服赠送
	KindPresentServiceAdd = 21
	// 客服扣减
	KindPresentServiceDiscount = 22
)

const (
	KindGrow = 7 // 增利

	//KindCommission = 9 // 手续费

	// 赠送
	//KindBalancePresent = 3

	// 流通账户
	KindBalanceFlow = 4 // 账户流通

	// 提现
	//KindBalanceApplyCash = 11
	// 转账
	KindBalanceTransfer = 12

	// 提现并充值到余额
	TypeApplyCashToCharge = 1
	// 提现到银行卡
	TypeApplyCashToBank = 2
	// 提现到第三方服务提供商（如：Paypal,支付宝等)
	TypeApplyCashToServiceProvider = 3

	// 退款到银行卡
	TypeBackToBank = 1
	// 退款到第三方
	TypeBackToServiceProvider = 2

	// 提现请求已提交
	StateApplySubmitted = 0
	// 提现已经确认
	StateApplyConfirmed = 1
	// 提现未通过
	StateApplyNotPass = 2
	// 提现完成
	StateApplyOver = 3

	StatusNormal = 0
	StatusOK     = 1
)

const (
	// 赠送
	TypeIntegralPresent = 1
	// 积分抵扣
	TypeIntegralDiscount = 2
	// 积分冻结
	TypeIntegralFreeze = 3
	// 积分解冻
	TypeIntegralUnfreeze = 4
	// 购物赠送
	TypeIntegralShoppingPresent = 5
	// 支付抵扣
	TypeIntegralPaymentDiscount = 6
)

type (
	IAccount interface {
		// 获取领域对象编号
		GetDomainId() int

		// 获取账户值
		GetValue() *Account

		// 保存
		Save() (int, error)

		// 设置优先(默认)支付方式, account 为账户类型
		SetPriorityPay(account int, enabled bool) error

		// 根据编号获取余额变动信息
		GetBalanceInfo(id int) *BalanceInfo

		// 根据号码获取余额变动信息
		// GetBalanceInfoByNo(no string) *BalanceInfo

		// 保存余额变动信息
		SaveBalanceInfo(*BalanceInfo) (int, error)

		// 获取赠送账户日志
		GetPresentLog(id int) *PresentLog

		// 充值,客服操作时,需提供操作人(relateUser)
		ChargeForBalance(chargeType int, title string, outerNo string, amount float32, relateUser int) error

		// 扣减余额
		DiscountBalance(title string, outerNo string, amount float32, relateUser int) error

		// 冻结余额
		Freeze(title string, outerNo string, amount float32, relateUser int) error

		// 解冻金额
		Unfreeze(title string, outerNo string, amount float32, relateUser int) error

		// 赠送金额,客服操作时,需提供操作人(relateUser)
		ChargeForPresent(title string, outerNo string, amount float32, relateUser int) error

		// 赠送金额(指定业务类型)
		ChargePresentByKind(kind int, title string, outerNo string, amount float32, relateUser int) error

		// 扣减奖金,mustLargeZero是否必须大于0, 赠送金额存在扣为负数的情况
		DiscountPresent(title string, outerNo string, amount float32,
			relateUser int, mustLargeZero bool) error

		// 冻结赠送金额
		FreezePresent(title string, outerNo string, amount float32, relateUser int) error

		// 解冻赠送金额
		UnfreezePresent(title string, outerNo string, amount float32, relateUser int) error

		// 流通账户余额变动，如扣除,amount传入负数金额
		ChargeFlowBalance(title string, tradeNo string, amount float32) error

		// 支付单抵扣消费,tradeNo为支付单单号
		PaymentDiscount(tradeNo string, amount float32, remark string) error

		//　增加积分
		AddIntegral(iType int, outerNo string, value int, remark string) error

		// 积分抵扣
		IntegralDiscount(logType int, outerNo string, value int, remark string) error

		// 冻结积分,当new为true不扣除积分,反之扣除积分
		FreezesIntegral(value int, new bool, remark string) error

		// 解冻积分
		UnfreezesIntegral(value int, remark string) error

		// 退款
		RequestBackBalance(backType int, title string, amount float32) error

		// 完成退款
		FinishBackBalance(id int, tradeNo string) error

		// 申请提现,applyType：提现方式,返回info_id,交易号 及错误
		RequestTakeOut(applyType int, title string, amount float32, commission float32) (int, string, error)

		// 确认提现
		ConfirmTakeOut(id int, pass bool, remark string) error

		// 完成提现
		FinishTakeOut(id int, tradeNo string) error

		// 将冻结金额标记为失效
		FreezeExpired(accountKind int, amount float32, remark string) error

		// 转账
		TransferAccounts(accountKind int, toMember int, amount float32,
			csnRate float32, remark string) error

		// 接收转账
		ReceiveTransfer(accountKind int, fromMember int, tradeNo string,
			amount float32, remark string) error

		// 转账余额到其他账户
		TransferBalance(kind int, amount float32, tradeNo string, toTitle, fromTitle string) error

		// 转账返利账户,kind为转账类型，如 KindBalanceTransfer等
		// commission手续费
		TransferPresent(kind int, amount float32, commission float32, tradeNo string,
			toTitle string, fromTitle string) error

		// 转账活动账户,kind为转账类型，如 KindBalanceTransfer等
		// commission手续费
		TransferFlow(kind int, amount float32, commission float32, tradeNo string,
			toTitle string, fromTitle string) error

		// 将活动金转给其他人
		TransferFlowTo(memberId int, kind int, amount float32, commission float32,
			tradeNo string, toTitle string, fromTitle string) error
	}

	// 余额变动信息
	BalanceInfo struct {
		Id       int    `db:"id" auto:"yes" pk:"yes"`
		MemberId int    `db:"member_id"`
		TradeNo  string `db:"trade_no"`
		Kind     int    `db:"kind"`
		Type     int    `db:"type"`
		Title    string `db:"title"`
		// 金额
		Amount float32 `db:"amount"`
		// 手续费
		CsnAmount float32 `db:"csn_amount"`
		// 引用编号
		RefId      int   `db:"ref_id"`
		State      int   `db:"state"`
		CreateTime int64 `db:"create_time"`
		UpdateTime int64 `db:"update_time"`
	}

	// 余额日志
	BalanceLog struct {
		Id       int    `db:"id" auto:"yes" pk:"yes"`
		MemberId int    `db:"member_id"`
		OuterNo  string `db:"outer_no"`
		// 业务类型
		BusinessKind int    `db:"kind"`
		Title        string `db:"title"`
		// 金额
		Amount float32 `db:"amount"`
		// 手续费
		CsnFee float32 `db:"csn_fee"`
		// 关联操作人,仅在客服操作时,记录操作人
		RelateUser int `db:"rel_user"`
		// 状态
		State int `db:"state"`
		// 备注
		Remark string `db:"remark"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 赠送账户日志
	PresentLog struct {
		Id       int    `db:"id" auto:"yes" pk:"yes"`
		MemberId int    `db:"member_id"`
		OuterNo  string `db:"outer_no"`
		// 业务类型
		BusinessKind int    `db:"kind"`
		Title        string `db:"title"`
		// 金额
		Amount float32 `db:"amount"`
		// 手续费
		CsnFee float32 `db:"csn_fee"`
		// 关联操作人,仅在客服操作时,记录操作人
		RelateUser int `db:"rel_user"`
		// 状态
		State int `db:"state"`
		// 备注
		Remark string `db:"remark"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 账户值对象
	Account struct {
		// 会员编号
		MemberId int `db:"member_id" pk:"yes" json:"memberId"`
		// 积分
		Integral int `db:"integral"`
		// 不可用积分
		FreezeIntegral int `db:"freeze_integral"`
		// 余额
		Balance float32 `db:"balance" json:"balance"`
		// 不可用余额
		FreezeBalance float32 `db:"freeze_balance" json:"freezesFee"`
		// 失效的账户余额
		ExpiredBalance float32 `db:"expired_balance"`
		//奖金账户余额
		PresentBalance float32 `db:"present_balance" json:"presentBalance"`
		//冻结赠送金额
		FreezePresent float32 `db:"freeze_present" json:"FreezePresent"`
		//失效的赠送金额
		ExpiredPresent float32 `db:"expired_present"`
		//总赠送金额
		TotalPresentFee float32 `db:"total_present_fee" json:"totalPresentFee"`
		//流动账户余额
		FlowBalance float32 `db:"flow_balance" json:"flowBalance"`
		//当前理财账户余额
		GrowBalance float32 `db:"grow_balance" json:"growBalance"`
		//理财总投资金额,不含收益
		GrowAmount float32 `db:"grow_amount" json:"growAmount"`
		//当前收益金额
		GrowEarnings float32 `db:"grow_earnings" json:"growEarnings"`
		//累积收益金额
		GrowTotalEarnings float32 `db:"grow_total_earnings" json:"growTotalEarnings"`
		//总消费金额
		TotalConsumption float32 `db:"total_consumption" json:"totalFee"`
		//总充值金额
		TotalCharge float32 `db:"total_charge" json:"totalCharge"`
		//总支付额
		TotalPay float32 `db:"total_pay" json:"totalPay"`
		// 优先(默认)支付选项
		PriorityPay int `db:"priority_pay"`
		//更新时间
		UpdateTime int64 `db:"update_time" json:"updateTime"`
	}

	// 积分记录
	IntegralLog struct {
		// 编号
		Id int `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int `db:"member_id"`
		// 类型
		Type int `db:"type"`
		// 关联的编号
		OuterNo string `db:"outer_no"`
		// 积分值
		Value int `db:"value"`
		// 备注
		Remark string `db:"remark"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
	}
)
