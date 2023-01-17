package message

const (
	ASSET_MESSAGE_TYPE_ENUMERATION      = "enumeration"
	ASSET_MESSAGE_TYPE_ENUMERATION_DONE = "enumeration-done"
	ASSET_MESSAGE_TYPE_PERMUTATION      = "permutation"
	ASSET_MESSAGE_TYPE_PERMUTATION_DONE = "permutation-done"
	ASSET_MESSAGE_TYPE_PORTSCAN         = "portscan"
	ASSET_MESSAGE_TYPE_PORTSCAN_DONE    = "portscan-done"
	ASSET_MESSAGE_TYPE_HTTPCONNECT      = "httpconnect"
	ASSET_MESSAGE_TYPE_HTTPCONNECT_DONE = "httpconnect-done"
)

type AssetMessage struct {
	Type  string
	Host  string
	Ports []int
}
