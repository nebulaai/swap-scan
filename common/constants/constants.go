package constants

const (
	SWAN_PAYMENT_ABI_JSON = "on-chain/contracts/abi/SwanPayment.json"
	DEFAULT_SELECT_LIMIT  = "100"

	URL_EVENT_PREFIX = "events"

	HTTP_STATUS_SUCCESS = "success"
	HTTP_STATUS_FAIL    = "fail"
	HTTP_STATUS_ERROR   = "error"

	HTTP_CODE_200_OK                    = "200" //http.StatusOk
	HTTP_CODE_400_BAD_REQUEST           = "400" //http.StatusBadRequest
	HTTP_CODE_401_UNAUTHORIZED          = "401" //http.StatusUnauthorized
	HTTP_CODE_500_INTERNAL_SERVER_ERROR = "500" //http.StatusInternalServerError

	NETWORK_TYPE_GOERLI  = "goerli"
	NETWORK_TYPE_POLYGON = "polygon"
	NETWORK_TYPE_NBAI    = "nbai"
	NETWORK_TYPE_BSC     = "bsc"
	NETWORK_TYPE_ETH     = "eth"

	TRANSACTION_STATUS_SUCCESS = "success"
	TRANSACTION_STATUS_FAIL    = "fail"
	RANSACTION_STATUS_PENDING  = "pending"

	PRIVATE_KEY_NAME_FOR_BSC_ADMIN_WALLET  = "privateKeyForBscAdminWallet"
	PRIVATE_KEY_NAME_FOR_NBAI_ADMIN_WALLET = "privateKeyForNbaiAdminWallet"

	NETWORK_INFO_UUID_FOR_NBAI = "93d16508-01b4-4910-ae6a-3fef936493e6"
	NETWORK_INFO_UUID_FOR_BSC  = "980c82e2-9de4-4ae0-be29-66d8a423f717"
	NETWORK_INFO_UUID_FOR_ETH  = "2cb44291-0c35-448b-a121-97075492d92e"

	ZERO_18                         = "000000000000000000"
	TX_HASH_NIL_VALUE_FOR_CHARGED   = "0x111111111111111111111111111tx" // Empty transactions that have been charged for handling fees
	TX_HASH_NIL_VALUE_FOR_NO_CHARGE = "0x222222222222222222222222222tx" // Empty transactions that haven't been charged for handling fees
	WALLET_NIL_VALUE                = "0x111111111111111111111111111wa"
)
