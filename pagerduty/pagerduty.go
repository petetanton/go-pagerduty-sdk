package pagerduty

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/google/go-querystring/query"
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/cache"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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

	maxRetries = 5
)

type Client struct {
	h                     *http.Client
	cfg                   *Config
	userCache             *cache.UserCache
	escalationPolicyCache *cache.EscalationPolicyCache
}

type Config struct {
	ApiToken     string
	ApiUrl       string
	FromHeader   string
	Logger       logrus.FieldLogger
	TypesToCache []string
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

	return &Client{h: &http.Client{Timeout: 20 * time.Second}, cfg: cfg, userCache: cache.NewUserCache(cfg.Logger), escalationPolicyCache: cache.NewEscalationPolicyCache()}, nil
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

	return c.do(request, maxRetries)
}

func (c *Client) post(url string, body io.Reader) (*PagerDutyResponse, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return c.do(request, maxRetries)
}

func (c *Client) put(url string, body io.Reader) (*PagerDutyResponse, error) {
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	return c.do(request, maxRetries)
}

func (c *Client) delete(url string) error {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = c.do(request, maxRetries)
	return err
}

func (c *Client) do(r *http.Request, retry uint) (*PagerDutyResponse, error) {
	//Error: Get "https://api.pagerduty.com/users/P6GNAHL?limit=100&offset=0": http2: server sent GOAWAY and closed the connection; LastStreamID=199, ErrCode=NO_ERROR, debug=""

	c.cfg.Logger.Debugf("%s: %s", r.Method, r.URL.String())
	r.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.cfg.ApiToken))
	r.Header.Add("Content-Type", "application/json")

	if c.cfg.FromHeader != "" && strings.Contains(r.URL.String(), TypeResponsePlay) {
		r.Header.Add("From", c.cfg.FromHeader)
	}

	response, err := c.h.Do(r)
	if err != nil {
		if strings.Contains(err.Error(), "GOAWAY") {
			c.cfg.Logger.Errorf("error contains 'GOAWAY' so retrying: %v", err)
			time.Sleep(time.Minute)
			return c.do(r, retry)
		}

		if strings.Contains(err.Error(), "Client.Timeout") {
			c.cfg.Logger.Errorf("error contains 'Client.Timeout' so retrying: %v", err)
			time.Sleep(time.Minute)
			return c.do(r, retry)
		}
		return nil, err
	}

	defer response.Body.Close()
	c.cfg.Logger.Debugf("%s: %s status: %d", r.Method, r.URL.String(), response.StatusCode)

	if response.StatusCode == http.StatusTooManyRequests {
		c.cfg.Logger.Info("retrying due to 429")
		time.Sleep(time.Second * 30)
		return c.do(r, retry)
	}

	if response.StatusCode >= http.StatusInternalServerError {
		b, _ := ioutil.ReadAll(response.Body)
		c.cfg.Logger.Errorf("Error on request to %s %s", r.Method, r.URL.String())
		if retry > 0 {
			c.cfg.Logger.Info("retrying due to 500")
			time.Sleep(time.Second * 30)
			return c.do(r, retry-1)
		}

		return nil, fmt.Errorf("[%s] %s: got a %d response from PagerDuty with error: %s", r.Method, r.URL, response.StatusCode, string(b))
	}

	if response.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	body, err := gabs.ParseJSONBuffer(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= http.StatusBadRequest {
		c.cfg.Logger.Errorf("%s: %s status: %d,body: %s", r.Method, r.URL.String(), response.StatusCode, parseErrorResponse(body))

		return nil, fmt.Errorf("got a %d response from PagerDuty: %s", response.StatusCode, parseErrorResponse(body))
	}

	return &PagerDutyResponse{body: body}, nil
}

func parseErrorResponse(body *gabs.Container) string {
	if body.ExistsP("error.errors") {
		var errStr []string
		for _, container := range body.Path("error.errors").Children() {
			errStr = append(errStr, container.Data().(string))
		}

		return strings.Join(errStr, ",")
	}

	return body.String()
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
	if r == nil || r.body == nil {
		return true
	}

	return r.body.Exists("more") && r.body.Path("more").Data().(bool)
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

func (c *Client) shouldCacheType(tp string) bool {
	for _, s := range c.cfg.TypesToCache {
		if s == tp {
			return true
		}
	}

	return false
}
