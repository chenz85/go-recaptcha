package recaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	var data = map[string]string{
		"secret":   in_sec,
		"response": token,
	}
	if remoteip != "" {
		data["remoteip"] = remoteip
	}

	jsonData, je := json.Marshal(data)
	if je != nil {
		err = je
		return
	}

	resp, re := httpCli.Post(reqUrl, "application/json", bytes.NewReader(jsonData))
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
