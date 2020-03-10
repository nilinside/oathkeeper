package guard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ory/oathkeeper/pipeline"
	"github.com/ory/oathkeeper/pipeline/authn"
	"github.com/ory/x/httpx"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type UserRequest struct {
	Url    string            `json:"url"`
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   interface{}       `json:"body"`
}

type GuardianTeapot struct {
}

func (g *GuardianTeapot) Check(r *http.Request, session *authn.AuthenticationSession, config json.RawMessage, rule pipeline.Rule) (*http.Response, error) {
	var b bytes.Buffer
	header := make(map[string]string, 0)
	for h := range r.Header {
		header[h] = r.Header.Get(h)
	}
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.NewEncoder(&b).Encode(&UserRequest{
		Url:    fmt.Sprintf("%s://%s%s", r.URL.Scheme, r.URL.Host, r.URL.Path),
		Method: r.Method,
		Header: header,
		Body:   body,
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest("POST", "http://muster-teapot/api/v1/requests/guard", &b)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	client := httpx.NewResilientClientLatencyToleranceSmall(nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func NewGuardianTeapot() Guardian {
	return &GuardianTeapot{}
}
