package recaptcha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	in_verify_url_f = "https://%s/recaptcha/api/siteverify"
)

var (
	httpCli *http.Client
)

func init() {
	httpCli = &http.Client{}
}

// VerifyResult result data
type VerifyResult struct {
	Success     bool     `json:"success"`
	ChallengeAt string   `json:"challenge_ts"` // timestamp of the challenge load (ISO format yyyy-MM-dd'T'HH:mm:ssZZ)
	HostName    string   `json:"hostname"`     // the hostname of the site where the reCAPTCHA was solved
	ErrCodes    []string `json:"error-codes"`  // optional
}

// Verify verify token (send from client)
// remoteip can be empty
// ref: https://developers.google.com/recaptcha/docs/verify
func Verify(token string, remoteip string) (vr *VerifyResult, err error) {
	reqUrl := fmt.Sprintf(in_verify_url_f, in_host)

	var data = make([]string, 0, 3)
	data = append(data, fmt.Sprintf("secret=%s", in_sec))
	data = append(data, fmt.Sprintf("response=%s", token))
	if remoteip != "" {
		data = append(data, fmt.Sprintf("remoteip=%s", remoteip))
	}

	resp, re := httpCli.Post(reqUrl, "application/x-www-form-urlencoded", strings.NewReader(strings.Join(data, "&")))
	if re != nil {
		err = re
		return
	}
	defer resp.Body.Close()

	if rdata, re := ioutil.ReadAll(resp.Body); re != nil {
		err = re
	} else if je := json.Unmarshal(rdata, &vr); je != nil {
		err = je
	}
	return
}
