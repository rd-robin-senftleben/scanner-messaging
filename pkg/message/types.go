package message

const (
	TOPIC_ENUMERATIONS  = "enumeration"
	TOPIC_PORT          = "port"
	TOPIC_VULNERABILITY = "vulnerability"
)

const (
	TYPE_ENUMERATION      = "enumeration"
	TYPE_ENUMERATION_DONE = "enumeration-done"
	TYPE_PERMUTATION      = "permutation"
	TYPE_PERMUTATION_DONE = "permutation-done"
	TYPE_PORTSCAN         = "portscan"
	TYPE_PORTSCAN_DONE    = "portscan-done"
	TYPE_HTTPCONNECT      = "httpconnect"
	TYPE_HTTPCONNECT_DONE = "httpconnect-done"
)

type RequestResponse struct {
	Type string
	Host string
}

type EnumerationRequest struct {
	RequestResponse
}
type EnumerationResponse struct {
	RequestResponse
}

type PermutationRequest struct {
	RequestResponse
}
type PermutationResponse struct {
	RequestResponse
}

type PortscanRequest struct {
	RequestResponse
}
type PortscanResponse struct {
	RequestResponse
	Ports []int
}

type HttpscanRequest struct {
	RequestResponse
}
type HttpscanResponse struct {
	RequestResponse
	Url  string
	Port int
}
