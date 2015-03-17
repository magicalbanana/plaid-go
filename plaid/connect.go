package plaid

import (
	"bytes"
	"encoding/json"
)

// POST /connect
// Submits a new user given a set of credentials
func (c client) ConnectAddUser(username, password, pin, institutionType string,
	options *ConnectOptions) (postRes *postResponse, mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectJson{
		c.clientID,
		c.secret,
		institutionType,
		username,
		password,
		pin,
		options,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect", bytes.NewReader(jsonText))
}

// POST /connect/step
// Submits an mfa "send_method", e.g. {"mask":"xxx-xxx-5309"}
func (c client) ConnectStepSendMethod(accessToken, key, value string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	sendMethod := map[string]string{key: value}
	jsonText, err := json.Marshal(connectStepSendMethodJson{
		c.clientID,
		c.secret,
		accessToken,
		connectStepOptions{sendMethod},
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect/step", bytes.NewReader(jsonText))
}

// POST /connect/step
// Submits an mfa answer
func (c client) ConnectStep(accessToken, answer string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectStepJson{
		c.clientID,
		c.secret,
		accessToken,
		answer,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect", bytes.NewReader(jsonText))
}

// POST /connect/get
// Retrieves account and transaction data for an access token
func (c client) ConnectGet(accessToken string, options *ConnectGetOptions) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectGetJson{
		c.clientID,
		c.secret,
		accessToken,
		options,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect/get", bytes.NewReader(jsonText))
}

// PATCH /connect
// Update a users credentials
func (c client) ConnectUpdate(username, password, pin, accessToken string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectUpdateJson{
		c.clientID,
		c.secret,
		username,
		password,
		pin,
		accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.patchAndUnmarshal("/connect", bytes.NewReader(jsonText))
}

// PATCH /connect/step
// Send MFA for updating a user
func (c client) ConnectUpdateStep(username, password, pin, mfa, accessToken string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectUpdateStepJson{
		c.clientID,
		c.secret,
		username,
		password,
		pin,
		mfa,
		accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.patchAndUnmarshal("/connect/step", bytes.NewReader(jsonText))
}

// DELETE /connect
// Deletes data associated with an access token
func (c client) ConnectDelete(accessToken string) (deleteRes *deleteResponse, err error) {
	jsonText, err := json.Marshal(connectDeleteJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	return c.deleteAndUnmarshal("/connect", bytes.NewReader(jsonText))
}

type ConnectOptions struct {
	Webhook   string `json:"webhook,omitempty"`
	Pending   bool   `json:"pending,omitempty"`
	LoginOnly bool   `json:"login_only,omitempty"`
	List      bool   `json:"list,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}
type connectJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	Type     string `json:"type"`

	Username string          `json:"username"`
	Password string          `json:"password"`
	PIN      string          `json:"pin,omitempty"`
	Options  *ConnectOptions `json:"options,omitempty"`
}

type connectStepOptions struct {
	SendMethod map[string]string `json:"send_method"`
}
type connectStepSendMethodJson struct {
	ClientID    string             `json:"client_id"`
	Secret      string             `json:"secret"`
	AccessToken string             `json:"access_token"`
	Options     connectStepOptions `json:"options"`
}

type connectStepJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`

	MFA string `json:"mfa"`
}

type ConnectGetOptions struct {
	Pending bool   `json:"pending,omitempty"`
	Account string `json:"account,omitempty"`
	GTE     string `json:"gte,omitempty"`
	LTE     string `json:"lte,omitempty"`
}
type connectGetJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`

	Options *ConnectGetOptions `json:"options,omitempty"`
}

type connectUpdateJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	AccessToken string `json:"access_token"`
}

type connectUpdateStepJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	MFA         string `json:"mfa"`
	AccessToken string `json:"access_token"`
}

type connectDeleteJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
