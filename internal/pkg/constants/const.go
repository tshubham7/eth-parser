package constants

type EthApiMethod string

const (
	MethodBlockNumber      EthApiMethod = "eth_blockNumber"
	MethodGetBlockByNumber EthApiMethod = "eth_getBlockByNumber"
)

const (
	EnvEthEndpoint    = "ETH_ENDPOINT"
	EnvLogLevel       = "LOG_LEVEL"
	EnvAdminAuthToken = "ADMIN_AUTH_TOKEN"
)

const (
	HeaderAdminAuthToken = "Admin-Auth-Token"
)
