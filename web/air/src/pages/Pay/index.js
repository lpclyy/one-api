import React, {useState, useEffect} from 'react';
import {API, showError, showInfo, showSuccess} from '../../helpers';
import {renderNumber, renderQuota} from '../../helpers/render';
import {Col, Layout, Row, Typography, Card, Button, Form, Divider, Space, Modal, Spin, message} from "@douyinfe/semi-ui";
import Title from "@douyinfe/semi-ui/lib/es/typography/title";
import Text from '@douyinfe/semi-ui/lib/es/typography/text';

const Pay = () => {
    const [amount, setAmount] = useState(0);
    const [customAmount, setCustomAmount] = useState(0);
    const [paymentMethod, setPaymentMethod] = useState('alipay');
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [showPayModal, setShowPayModal] = useState(false);
    const [payUrl, setPayUrl] = useState('');
    const [qrCode, setQrCode] = useState('');
    const [orderId, setOrderId] = useState('');
    const [payStatus, setPayStatus] = useState('');
    const [polling, setPolling] = useState(null);

    // 预设额度档位
    const amountOptions = [
        {value: 10, label: '10元'},
        {value: 50, label: '50元'},
        {value: 100, label: '100元'},
        {value: 500, label: '500元'},
        {value: 1000, label: '1000元'}
    ];

    // 选择预设额度
    const handleAmountSelect = (value) => {
        setAmount(value);
        setCustomAmount(0);
    };

    // 自定义额度
    const handleCustomAmountChange = (value) => {
        if (value > 0) {
            setCustomAmount(value);
            setAmount(0);
        }
    };

    // 获取实际支付金额
    const getActualAmount = () => {
        return amount > 0 ? amount : customAmount;
    };

    // 创建支付订单
    const createPayOrder = async () => {
        const actualAmount = getActualAmount();
        if (actualAmount < 0.1) {
            showInfo('充值金额最低为0.1元');
            return;
        }
        if (actualAmount > 100000) {
            showInfo('单次充值金额不能超过100000元');
            return;
        }

        setIsSubmitting(true);
        try {
            let apiUrl = '';
            if (paymentMethod === 'alipay') {
                apiUrl = '/api/user/pay/alipay/create';
            } else if (paymentMethod === 'wechat') {
                // 后续添加微信支付接口
                showInfo('微信支付暂未开放');
                setIsSubmitting(false);
                return;
            }

            const res = await API.post(apiUrl, {
                amount: actualAmount
            });

            if (res.data.success) {
                setPayUrl(res.data.data);
                setOrderId('order_' + Date.now());
                setShowPayModal(true);
                setPayStatus('pending');
                // 开始轮询支付状态
                startPolling();
            } else {
                showError(res.data.message);
            }
        } catch (err) {
            showError('创建支付订单失败');
            console.error(err);
        } finally {
            setIsSubmitting(false);
        }
    };

    // 开始轮询支付状态
    const startPolling = () => {
        const interval = setInterval(async () => {
            try {
                // 调用后端支付状态查询接口
                const res = await API.get(`/api/user/pay/status?orderId=${orderId}`);
                if (res.data.success) {
                    if (res.data.data.status === 'success') {
                        setPayStatus('success');
                        clearInterval(interval);
                        showSuccess('支付成功！');
                        setTimeout(() => {
                            setShowPayModal(false);
                            // 跳转到充值成功页面
                            window.location.href = '/topup?success=true';
                        }, 2000);
                    } else if (res.data.data.status === 'failed') {
                        setPayStatus('failed');
                        clearInterval(interval);
                        showError('支付失败，请重试');
                    }
                }
            } catch (err) {
                console.error('查询支付状态失败', err);
            }
        }, 3000);

        setPolling(interval);
    };

    // 取消支付
    const handleCancelPay = () => {
        if (polling) {
            clearInterval(polling);
        }
        setShowPayModal(false);
        setPayStatus('');
        setPayUrl('');
        setOrderId('');
    };

    // 跳转到支付页面
    const redirectToPay = () => {
        if (payUrl) {
            window.open(payUrl, '_blank');
        }
    };

    return (
        <div>
            <Layout>
                <Layout.Header>
                    <h3>在线充值</h3>
                </Layout.Header>
                <Layout.Content>
                    <div style={{marginTop: 20, display: 'flex', justifyContent: 'center'}}>
                        <Card
                            style={{width: '500px', padding: '20px'}}
                        >
                            <Title level={3} style={{textAlign: 'center'}}>选择充值金额</Title>
                            
                            <div style={{marginTop: 20}}>
                                <Divider>
                                    预设额度
                                </Divider>
                                <Space wrap style={{marginBottom: 20}}>
                                    {amountOptions.map((option) => (
                                        <Button
                                            key={option.value}
                                            type={amount === option.value ? 'primary' : 'default'}
                                            theme={amount === option.value ? 'solid' : 'borderless'}
                                            onClick={() => handleAmountSelect(option.value)}
                                            style={{padding: '0 20px', margin: '5px'}}
                                        >
                                            {option.label}
                                        </Button>
                                    ))}
                                </Space>
                            </div>

                            <div style={{marginTop: 20}}>
                                <Divider>
                                    自定义额度
                                </Divider>
                                <Form.Input
                                    field={'customAmount'}
                                    label={'自定义金额'}
                                    placeholder={'请输入充值金额，最低0.1元'}
                                    name='customAmount'
                                    type={'number'}
                                    value={customAmount}
                                    suffix={'元'}
                                    min={0.1}
                                    max={100000}
                                    step={0.1}
                                    onChange={(value) => handleCustomAmountChange(value)}
                                />
                            </div>

                            <div style={{marginTop: 20}}>
                                <Divider>
                                    支付方式
                                </Divider>
                                <Space>
                                    <Button
                                        type={paymentMethod === 'alipay' ? 'primary' : 'default'}
                                        theme={paymentMethod === 'alipay' ? 'solid' : 'borderless'}
                                        onClick={() => setPaymentMethod('alipay')}
                                    >
                                        支付宝
                                    </Button>
                                    <Button
                                        type={paymentMethod === 'wechat' ? 'primary' : 'default'}
                                        theme={paymentMethod === 'wechat' ? 'solid' : 'borderless'}
                                        onClick={() => setPaymentMethod('wechat')}
                                        disabled={true}
                                    >
                                        微信支付（暂未开放）
                                    </Button>
                                </Space>
                            </div>

                            <div style={{marginTop: 30, textAlign: 'center'}}>
                                <Button
                                    type={'primary'}
                                    theme={'solid'}
                                    onClick={createPayOrder}
                                    disabled={isSubmitting || getActualAmount() < 0.1}
                                    size={'large'}
                                >
                                    {isSubmitting ? '处理中...' : `确认支付 ${getActualAmount()} 元`}
                                </Button>
                            </div>
                        </Card>
                    </div>
                </Layout.Content>
            </Layout>

            {/* 支付模态框 */}
            <Modal
                title="支付确认"
                visible={showPayModal}
                onCancel={handleCancelPay}
                maskClosable={false}
                size={'medium'}
                centered={true}
                footer={null}
            >
                {payStatus === 'pending' && (
                    <div style={{textAlign: 'center'}}>
                        <p>订单号：{orderId}</p>
                        <p>支付金额：{getActualAmount()} 元</p>
                        <p>支付方式：{paymentMethod === 'alipay' ? '支付宝' : '微信支付'}</p>
                        <Divider />
                        <p>请点击下方按钮跳转到支付页面</p>
                        <Button
                            type={'primary'}
                            theme={'solid'}
                            onClick={redirectToPay}
                            style={{marginTop: 20}}
                        >
                            跳转到支付页面
                        </Button>
                        <p style={{marginTop: 20, color: '#666'}}>支付完成后请耐心等待系统确认...</p>
                        <Spin style={{marginTop: 20}} />
                    </div>
                )}
                
                {payStatus === 'success' && (
                    <div style={{textAlign: 'center'}}>
                        <p style={{color: '#52c41a', fontSize: '18px', fontWeight: 'bold'}}>支付成功！</p>
                        <p>订单号：{orderId}</p>
                        <p>支付金额：{getActualAmount()} 元</p>
                        <p style={{marginTop: 20}}>系统正在处理，请稍候...</p>
                        <Spin style={{marginTop: 20}} />
                    </div>
                )}
                
                {payStatus === 'failed' && (
                    <div style={{textAlign: 'center'}}>
                        <p style={{color: '#ff4d4f', fontSize: '18px', fontWeight: 'bold'}}>支付失败</p>
                        <p>订单号：{orderId}</p>
                        <p>请检查支付信息后重试</p>
                        <Button
                            type={'primary'}
                            theme={'solid'}
                            onClick={handleCancelPay}
                            style={{marginTop: 20}}
                        >
                            关闭
                        </Button>
                    </div>
                )}
            </Modal>
        </div>
    );
};

export default Pay;