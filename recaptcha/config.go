package recaptcha

import "errors"

var (
	in_sec  string
	in_host string
)

// Config init config info
func Config(sec string) {
	ConfigWithHost(sec, "www.google.com")
}

// ConfigWithHost init config info
// host should be one of:
// 	www.google.com
// 	www.recaptcha.net
func ConfigWithHost(sec string, host string) {
	if sec == "" {
		panic(errors.New("require secret of reCAPTCHA"))
	} else if host == "" {
		panic(errors.New("require host of reCAPTCHA"))
	}

	in_sec, in_host = sec, host
}
