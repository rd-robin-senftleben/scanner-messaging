package types

const (
	TOPIC_ENUMERATIONS  = "enumeration"
	TOPIC_PORT          = "port"
	TOPIC_VULNERABILITY = "vulnerability"
)

func ALL_TOPICS() []string {
	return []string{TOPIC_ENUMERATIONS, TOPIC_VULNERABILITY, TOPIC_PORT}
}

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
	Type    string
	Host    string
	Payload map[string]any
}

type EnumerationRequestResponse struct {
	RequestResponse
}

type PermutationRequestResponse struct {
	RequestResponse
}

type PortscanRequestResponse struct {
	RequestResponse
	Ports []int
}

type HttpconnectRequestResponse struct {
	PortscanRequestResponse
}

type VulnerabilityRequestResponse struct {
	RequestResponse
	Url       string
	Templates []string
}
