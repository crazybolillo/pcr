package phone

import "regexp"

var avayaJ100Regex = regexp.MustCompile(`AVAYA/(J\d{3})-(?:\d\.?)+ \(MAC:([abcdef0-9]{12})\)$`)

func avayaJSeries(ua string) Identity {
	matches := avayaJ100Regex.FindStringSubmatch(ua)
	if matches == nil {
		return Identity{}
	}

	return Identity{
		Model: matches[1],
		MAC:   matches[2],
	}
}
