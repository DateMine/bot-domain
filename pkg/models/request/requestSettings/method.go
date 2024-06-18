package requestSettings

type Method int32

const (
	GET  Method = 1
	POST Method = 2
)

func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	default:
		return "GET"
	}
}

func (m Method) Value() int32 {
	switch m {
	case GET:
		return 1
	case POST:
		return 2
	default:
		return 1
	}
}
