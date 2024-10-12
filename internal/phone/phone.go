package phone

type Identity struct {
	Model string
	MAC   string
}

type Values struct {
	DisplayName string
	Username    string
	Password    string
}

type Identifier func(ua string) Identity

var identifiers = []Identifier{
	avayaJSeries,
}

// Identify a phone based on its User-Agent.
func Identify(ua string) Identity {
	for _, identifier := range identifiers {
		identity := identifier(ua)
		if (Identity{}) != identity {
			return identity
		}
	}

	return Identity{}
}
