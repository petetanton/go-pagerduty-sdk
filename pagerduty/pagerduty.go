package pagerduty

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultUrl             = "https://api.pagerduty.com"
	TypeEscalationPolicies = "escalation_policies"
	TypeEscalationPolicy   = "escalation_policy"
	TypeTeams              = "teams"
	TypeTeam               = "team"
	TypeTags               = "tags"
	TypeUser               = "user"
	TypeUsers              = "users"
	TypeService            = "service"
	TypeServices           = "services"
	TypeSchedule           = "schedule"
	TypeSchedules          = "schedules"
	TypeNotificationRule   = "notification_rule"
	TypeNotificationRules  = "notification_rules"
	TypeContactMethod      = "contact_method"
	TypeContactMethods     = "contact_methods"
	TypeResponsePlay       = "response_play"
	TypeResponsePlays      = "response_plays"
)

type Client struct {
	h   *http.Client
	cfg *Config
}

type Config struct {
	ApiToken string
	ApiUrl   string
	Logger   logrus.FieldLogger
}

type ApiObject struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self,omitempty"`
	HtmlUrl string `json:"html_url,omitempty"`
}

func NewClient(cfg *Config) (*Client, error) {
	if cfg.ApiUrl == "" {
		cfg.ApiUrl = "https://api.pagerduty.com"
	}

	if cfg.Logger == nil {
		cfg.Logger = logrus.New()
	}

	if cfg.ApiToken == "" {
		return nil, errors.New("please set an api token")
	}

	return &Client{h: &http.Client{Timeout: 10 * time.Second}, cfg: cfg}, nil
}

func (c *Client) get(url string, params interface{}) (*PagerDutyResponse, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodGet, url+"?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.do(request)
}

func (c *Client) post(url string, body io.Reader) (*PagerDutyResponse, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return c.do(request)
}

func (c *Client) put(url string, body io.Reader) (*PagerDutyResponse, error) {
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	return c.do(request)
}

func (c *Client) delete(url string) error {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = c.do(request)
	return err
}

func (c *Client) do(r *http.Request) (*PagerDutyResponse, error) {
	c.cfg.Logger.Debugf("%s: %s", r.Method, r.URL.String())
	r.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.cfg.ApiToken))
	r.Header.Add("Content-Type", "application/json")
	response, err := c.h.Do(r)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusTooManyRequests {
		c.cfg.Logger.Debug("retrying due to 429")
		time.Sleep(time.Second * 30)
		return c.do(r)
	}

	if response.StatusCode >= http.StatusBadRequest {
		b, _ := ioutil.ReadAll(response.Body)
		c.cfg.Logger.Errorf("Error on request to %s %s", r.Method, r.URL.String())
		return nil, fmt.Errorf("got a %d response from PagerDuty with error: %s", response.StatusCode, string(b))
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	body, err := gabs.ParseJSONBuffer(response.Body)
	if err != nil {
		return nil, err
	}

	return &PagerDutyResponse{body: body}, nil
}

func DefaultPagerDutyRequest() *PagerDutyRequest {
	return &PagerDutyRequest{
		Offset: 0,
		Limit:  100,
	}
}

type PagerDutyRequest struct {
	Offset   float64  `url:"offset"`
	Limit    float64  `url:"limit"`
	Includes []string `url:"include,omitempty,brackets"`
}
type PagerDutyResponse struct {
	body *gabs.Container
}

func (r *PagerDutyResponse) hasMore() bool {
	if r.body == nil {
		return true
	}

	return r.body.Path("more").Data().(bool)
}

func (r *PagerDutyResponse) nextOffset() float64 {
	if r.body == nil {
		return 0
	}

	return r.body.Path("limit").Data().(float64) + r.body.Path("offset").Data().(float64)
}

func (r *PagerDutyResponse) unmarshallResponse(out interface{}, path string) error {
	return json.Unmarshal(r.body.Path(path).Bytes(), out)
}

func (c *Client) objectToJson(in interface{}, path string) (io.Reader, error) {
	container := gabs.New()
	_, err := container.Set(in, path)
	if err != nil {
		return nil, err
	}
	b, err := container.MarshalJSON()
	if err != nil {
		return nil, err
	}

	c.cfg.Logger.Debug(string(b))
	return bytes.NewReader(b), nil
}
