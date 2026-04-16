package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/songquanpeng/one-api/common/env"

	"github.com/google/uuid"
)

var SystemName = "瓦兰卡"
var ServerAddress = "http://walankaai.com"
var Footer = ""
var Logo = ""

var ChatLink = ""
var QuotaPerUnit = 500 * 1000.0 // $0.002 / 1K tokens
var DisplayInCurrencyEnabled = true
var DisplayTokenStatEnabled = true

// Any options with "Secret", "Token" in its key won't be return by GetOptions

var SessionSecret = uuid.New().String()

var OptionMap map[string]string
var OptionMapRWMutex sync.RWMutex

var ItemsPerPage = 10
var MaxRecentItems = 100

var PasswordLoginEnabled = true
var PasswordRegisterEnabled = true
var EmailVerificationEnabled = false
var GitHubOAuthEnabled = false
var OidcEnabled = false
var WeChatAuthEnabled = false
var TurnstileCheckEnabled = false
var RegisterEnabled = true

var EmailDomainRestrictionEnabled = false
var EmailDomainWhitelist = []string{
	"gmail.com",
	"163.com",
	"126.com",
	"qq.com",
	"outlook.com",
	"hotmail.com",
	"icloud.com",
	"yahoo.com",
	"foxmail.com",
}

var DebugEnabled = strings.ToLower(os.Getenv("DEBUG")) == "true"
var DebugSQLEnabled = strings.ToLower(os.Getenv("DEBUG_SQL")) == "true"
var MemoryCacheEnabled = strings.ToLower(os.Getenv("MEMORY_CACHE_ENABLED")) == "true"

var LogConsumeEnabled = true

var SMTPServer = ""
var SMTPPort = 587
var SMTPAccount = ""
var SMTPFrom = ""
var SMTPToken = ""

var GitHubClientId = ""
var GitHubClientSecret = ""

var LarkClientId = ""
var LarkClientSecret = ""

var OidcClientId = ""
var OidcClientSecret = ""
var OidcWellKnown = ""
var OidcAuthorizationEndpoint = ""
var OidcTokenEndpoint = ""
var OidcUserinfoEndpoint = ""

var WeChatServerAddress = ""
var WeChatServerToken = ""
var WeChatAccountQRCodeImageURL = ""

var MessagePusherAddress = ""
var MessagePusherToken = ""

var TurnstileSiteKey = ""
var TurnstileSecretKey = ""

var StripeSecretKey = ""
var StripeWebhookSecret = ""

var AlipayAppId = "2088480387025401"
var AlipayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnlDOjVsB0kEUIkaqDeR/F+jx+/rg2KD7mqan/Ql6ILRnPR9ZTqd46+poOG6pijsbhK+hMmSj+HkKo60ek2oDwPW2I79VQxj9krQufS9JOTysBFlrj6GmSlzNguvVSOp0BZSHOaPcIHAp80TQVp1e5ptOtS/0idbPBkfvqSmvWECUWLBnO0P3Vnvtyow08w789GWeS9HE+X5SpxATABM8hd8Gy4w74f28AJSgWmIcQqzlzeCfke/zxiDrQ7sqNqHOehVQ+82+03E8MqCNaFy11/R0+3NDi/EAYDCdPDpuJdhaDWXd1TYGlUecbtpBV+7SvNzPPhX+gRbWS9ELlaQLUwIDAQAB"
var AlipayPrivateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCeUM6NWwHSQRQiRqoN5H8X6PH7+uDYoPuapqf9CXogtGc9H1lOp3jr6mg4bqmKOxuEr6EyZKP4eQqjrR6TagPA9bYjv1VDGP2StC59L0k5PKwEWWuPoaZKXM2C69VI6nQFlIc5o9wgcCnzRNBWnV7mm061L/SJ1s8GR++pKa9YQJRYsGc7Q/dWe+3KjDTzDvz0ZZ5L0cT5flKnEBMAEzyF3wbLjDvh/bwAlKBaYhxCrOXN4J+R7/PGIOtDuyo2oc56FVD7zb7TcTwyoI1oXLXX9HT7c0OL8QBgMJ08Om4l2FoNZd3VNgaVR5xu2kFX7tK83M8+Ff6BFtZL0QuVpAtTAgMBAAECggEAGjPnq++3InSQ/4dQmBIMkwmEG5+PXWtvmU4iGbha1VTmjXBF2MXRNsLKUyTFvNJBoLls5alQlkx6XVLG78EpM+O3LL17QCq/tWtLEn8kEGwhUGr4aYJowoAPp66e032yjEXkB78+LMRDvTbTgLJ1RZfI9CYmxDWgeDqpyQbUNbffSbTxhTXneGpuh5XjY4+5V1uzUU8/130WkJ97qEl9V6s4AfpIBqZuafrKjDZz+8teO5vnqDYkttMS75Fn74nnVyLOJQGZwlF19gZzmYOFBPpUrqrDIhELtBipGZgJm59VzDTeGGt14RC7jBaUDgcZgrY/gJkI2vrPWBF67wOr8QKBgQDnnWT+pevZjtyduc/SRASEvJT8KLCSemX7kVpyp1vCGVvBn0xngosbnU66lNJHywFZ6x6t5CEZBXxBaYf6fozbQQqcZOLa9kcIfYlE+3gisdvZuVgp2BQV70FydzBXfHhoar2LOM6elCdxVbybumM48ubN6I8xgPDoypbsHEJ/KQKBgQCu+9BxJZtpCSoilfosahmHgnC5/2JnOi6MJYtr0O/Mhyss4f5oVEAGQ3pHgNv8WOjigIDerWlZLw29o3AEk+fBzSiBwealv35DG3xl/4vwwyD+K4hmvYL0R4eN1zV8ad2oDuxVaaYPUrMqHHDKNMzqZAixMFGX6+4UhSYiAunSGwKBgBx4hImo66z6mrPou1slcUi/xbCZb9sRoKej3nJpkCXz6AuNAV9X9LGYTK0yzgZ1Nd1PwZ2uhUMGIZgI2OY52Ca7gAppfFleHK02gUExiDr7kgLZfbTnEtD/cBQaAp8+da6gMFyExyFHJPIRj/W0m63MbgKxq6hyKSr0fEjZ0HLBAoGAGHnSozD+dwe4JBRUZQgGQCUnvWySiBvkTOgng1I7aKFzkZie7Fr0havEm+HTY43QLXaKEBuzg60IQAFvdsR1g289/kBwEbkiYSKkGORQ38F7iPHv52cUvTSQKm/y5E8umQZVWnEnsDcCJp7JzA7sptCNQrOehiCTMb2aIuHDcQcCgYBfPA8AegGMqWEFba7qEECdMofHmyzu7QsYYbcfrFSOdNRdYkfD2bFfeYlEBUeRdTDwlgbQ9g+WvF6tmLvUVLAKOHDtwXhET6qr7sKwzwjEh1AZBh+m3klwfdx4zv/CP4t6P8DE5sA8bfk1kESPmq8GIGEfxg+cvloLFa05ZTheTw=="
var PayRateRMB float64 = 7.2


var QuotaForNewUser int64 = 0
var QuotaForInviter int64 = 0
var QuotaForInvitee int64 = 0
var ChannelDisableThreshold = 5.0
var AutomaticDisableChannelEnabled = false
var AutomaticEnableChannelEnabled = false
var QuotaRemindThreshold int64 = 1000
var PreConsumedQuota int64 = 500
var ApproximateTokenEnabled = false
var RetryTimes = 0

var RootUserEmail = ""

var IsMasterNode = os.Getenv("NODE_TYPE") != "slave"

var requestInterval, _ = strconv.Atoi(os.Getenv("POLLING_INTERVAL"))
var RequestInterval = time.Duration(requestInterval) * time.Second

var SyncFrequency = env.Int("SYNC_FREQUENCY", 10*60) // unit is second

var BatchUpdateEnabled = false
var BatchUpdateInterval = env.Int("BATCH_UPDATE_INTERVAL", 5)

var RelayTimeout = env.Int("RELAY_TIMEOUT", 0) // unit is second

var GeminiSafetySetting = env.String("GEMINI_SAFETY_SETTING", "BLOCK_NONE")

var Theme = env.String("THEME", "default")
var ValidThemes = map[string]bool{
	"default": true,
	"berry":   true,
	"air":     true,
}

// All duration's unit is seconds
// Shouldn't larger then RateLimitKeyExpirationDuration
var (
	GlobalApiRateLimitNum            = env.Int("GLOBAL_API_RATE_LIMIT", 480)
	GlobalApiRateLimitDuration int64 = 3 * 60

	GlobalWebRateLimitNum            = env.Int("GLOBAL_WEB_RATE_LIMIT", 240)
	GlobalWebRateLimitDuration int64 = 3 * 60

	UploadRateLimitNum            = 10
	UploadRateLimitDuration int64 = 60

	DownloadRateLimitNum            = 10
	DownloadRateLimitDuration int64 = 60

	CriticalRateLimitNum            = 20
	CriticalRateLimitDuration int64 = 20 * 60
)

var RateLimitKeyExpirationDuration = 20 * time.Minute

var EnableMetric = env.Bool("ENABLE_METRIC", false)
var MetricQueueSize = env.Int("METRIC_QUEUE_SIZE", 10)
var MetricSuccessRateThreshold = env.Float64("METRIC_SUCCESS_RATE_THRESHOLD", 0.8)
var MetricSuccessChanSize = env.Int("METRIC_SUCCESS_CHAN_SIZE", 1024)
var MetricFailChanSize = env.Int("METRIC_FAIL_CHAN_SIZE", 128)

var InitialRootToken = os.Getenv("INITIAL_ROOT_TOKEN")

var InitialRootAccessToken = os.Getenv("INITIAL_ROOT_ACCESS_TOKEN")

var GeminiVersion = env.String("GEMINI_VERSION", "v1")

var OnlyOneLogFile = env.Bool("ONLY_ONE_LOG_FILE", false)

var RelayProxy = env.String("RELAY_PROXY", "")
var UserContentRequestProxy = env.String("USER_CONTENT_REQUEST_PROXY", "")
var UserContentRequestTimeout = env.Int("USER_CONTENT_REQUEST_TIMEOUT", 30)

var EnforceIncludeUsage = env.Bool("ENFORCE_INCLUDE_USAGE", false)
var TestPrompt = env.String("TEST_PROMPT", "Output only your specific model name with no additional text.")
