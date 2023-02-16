package constants

const (
	APIServiceName         = "douyin.api"
	BaseServiceName        = "douyin.base"
	InteractionServiceName = "douyin.interaction"
	SocialServiceName      = "douyin.social"
	EtcdAddress            = "127.0.0.1:2379"
	BaseTCPAddr            = "127.0.0.1:8889"
	InteractionTCPAddr     = "127.0.0.1:8890"
	SocialTCPAddr          = "127.0.0.1:8891"
	MySQLDefaultDSN        = "douyin:douyin-7355608@tcp(localhost:9910)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	UploadAddr             = "http://192.168.85.149:8888/upload/" //客户端测试时ip地址改为自己的无线局域网ipv4地址
	VideoCountLimit        = 30
)
