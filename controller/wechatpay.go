package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/songquanpeng/one-api/common/config"
	"github.com/songquanpeng/one-api/common/ctxkey"
	"github.com/songquanpeng/one-api/common/logger"
	"github.com/songquanpeng/one-api/model"
)

type WechatPayCheckoutRequest struct {
	Amount float64 `json:"amount"` // Amount in RMB, e.g. 10.0 for ¥10
}

func CreateWechatPayCheckoutSession(c *gin.Context) {
	// 暂时返回未开放的提示，后续可以集成微信支付SDK
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "管理员尚未配置微信支付，请联系管理员开通在线支付。",
	})
	return

	// 以下是微信支付的实现示例，需要集成微信支付SDK
	/*
	if config.WechatPayAppId == "" || config.WechatPayMchId == "" || config.WechatPayApiKey == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "管理员尚未配置微信支付，请联系管理员开通在线支付。",
		})
		return
	}

	var req WechatPayCheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "请求参数错误"})
		return
	}

	if req.Amount < 0.1 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "充值金额最低为 ¥0.1"})
		return
	}
	if req.Amount > 100000 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "单次充值金额不能超过 ¥100,000"})
		return
	}

	userId := c.GetInt(ctxkey.Id)
	tradeNo := fmt.Sprintf("wechat_%d_%d", userId, time.Now().UnixNano())

	// 这里需要集成微信支付SDK，创建支付订单
	// 示例代码：
	// client := wechatpay.NewClient(...)
	// order := client.CreateOrder(...)

	// 模拟返回支付链接
	payUrl := fmt.Sprintf("wechat://pay?trade_no=%s&amount=%.2f", tradeNo, req.Amount)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    payUrl,
	})
	*/
}

func WechatPayWebhook(c *gin.Context) {
	// 暂时返回成功，后续可以集成微信支付SDK处理回调
	c.String(http.StatusOK, "success")
	return

	// 以下是微信支付回调的实现示例，需要集成微信支付SDK
	/*
	// 解析微信支付回调
	// 验证签名
	// 处理支付结果
	
	userId := 0 // 从回调中获取用户ID
	amountRMB := 0.0 // 从回调中获取支付金额
	sessionId := "" // 从回调中获取订单号

	rate := config.PayRateRMB
	if rate <= 0 {
		rate = 7.2 // Default failover
	}

	amountUSD := amountRMB / rate
	amountCents := int64(math.Round(amountUSD * 100))

	quota := int64(math.Round(amountUSD * (config.QuotaPerUnit / 0.002)))

	if model.IsStripeSessionProcessed(sessionId) {
		logger.SysLog(fmt.Sprintf("wechat pay webhook session %s already processed, skipping", sessionId))
		c.String(http.StatusOK, "success")
		return
	}

	if err := model.CreateStripePaymentRecord(sessionId, userId, amountCents, quota); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			c.String(http.StatusOK, "success")
			return
		}
		logger.Error(c.Request.Context(), "wechat pay webhook failed to record payment: "+err.Error())
		c.String(http.StatusInternalServerError, "fail")
		return
	}

	if err := model.IncreaseUserQuota(userId, quota); err != nil {
		logger.Error(c.Request.Context(), fmt.Sprintf("wechat pay webhook failed to credit quota for user %d: %s", userId, err.Error()))
		c.String(http.StatusOK, "success")
		return
	}

	model.RecordTopupLog(
		c.Request.Context(),
		userId,
		fmt.Sprintf("WeChat 扫码充值 ¥%.2f，到账 %d 配额", amountRMB, quota),
		int(quota),
	)

	logger.SysLog(fmt.Sprintf("wechat pay webhook successfully credited %d quota to user %d (session %s)", quota, userId, sessionId))
	c.String(http.StatusOK, "success")
	*/
}