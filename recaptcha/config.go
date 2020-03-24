package recaptcha

import "errors"

var (
	in_sec      string
	in_site_sec string
	in_host     string
)

// Config init config info
func Config(sec, site_sec string) {
	ConfigWithHost(sec, site_sec, "www.google.com")
}

// ConfigWithHost init config info
// host should be one of:
// 	www.google.com
// 	www.recaptcha.net
func ConfigWithHost(sec, site_sec string, host string) {
	if sec == "" {
		panic(errors.New("require secret of reCAPTCHA"))
	} else if site_sec == "" {
		panic(errors.New("require site secret of reCAPTCHA"))
	} else if host == "" {
		panic(errors.New("require host of reCAPTCHA"))
	}

	in_sec, in_site_sec, in_host = sec, site_sec, host
}
