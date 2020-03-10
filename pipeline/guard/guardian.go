package guard

import (
	"encoding/json"
	"github.com/ory/oathkeeper/pipeline"
	"github.com/ory/oathkeeper/pipeline/authn"
	"net/http"
)

type Guardian interface {
	Check(r *http.Request, session *authn.AuthenticationSession, config json.RawMessage, rule pipeline.Rule) (*http.Response, error)
}
