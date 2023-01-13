package message

const (
	ASSET_MESSAGE_TYPE_ENUMERATION = "enumeration"
	ASSET_MESSAGE_TYPE_PERMUTATION = "permutation"
)

type AssetMessage struct {
	Type string
	Host string
}
