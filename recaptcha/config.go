package recaptcha

import "errors"

var (
	in_sec      string
	in_site_sec string
)

// Config init config info
func Config(sec, site_sec string) {
	if sec == "" {
		panic(errors.New("require secret of reCAPTCHA"))
	} else if site_sec == "" {
		panic(errors.New("require site secret of reCAPTCHA"))
	}

	in_sec, in_site_sec = sec, site_sec
}
