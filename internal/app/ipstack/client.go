package ipstack

type ClientIP struct {
	Hostname  string
	AccessKey string
}

type QueryIP interface {
}

func NewClientIP(hostname, accessKey string) QueryIP {
	return &ClientIP{
		hostname,
		accessKey,
	}
}
