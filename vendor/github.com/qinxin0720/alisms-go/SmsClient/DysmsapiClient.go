package SmsClient

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"io/ioutil"
	"encoding/hex"
)

type keepAliveAgent struct {
	defaultPort    int
	KeepAlive      bool
	KeepAliveMsecs int
	protocol       string
}

func newKeepAliveAgent(defaultPort int, KeepAlive bool, KeepAliveMsecs int, protocol string) *keepAliveAgent {
	return &keepAliveAgent{
		defaultPort:    defaultPort,
		KeepAlive:      KeepAlive,
		KeepAliveMsecs: KeepAliveMsecs,
		protocol:       protocol,
	}
}

const (
	apiVersion = "2017-05-25"
)

type dysmsapiClient struct {
	accessKeyID,
	secretAccessKey,
	endpoint string
}

var makeNonce func() string

func newDysmsapiClient(accessKeyID, secretAccessKey, endpoint string) (*dysmsapiClient, error) {
	var ep string
	if accessKeyID == "" {
		return nil, errors.New("accessKeyId is empty")
	}
	if secretAccessKey == "" {
		return nil, errors.New("secretAccessKey is empty")
	}
	if endpoint == "" {
		return nil, errors.New("endpoint is empty")
	}
	if endpoint[len(endpoint)-1] == 0x2F { //0x2F 为 ASCII /
		ep = endpoint[:len(endpoint)-1]
	} else {
		ep = endpoint
	}
	makeNonce = makeNonce_()
	return &dysmsapiClient{
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		endpoint:        ep,
	}, nil
}

func (dsc *dysmsapiClient) SendSms(params Params, accessKeyID, secretAccessKey string) (int, string, error) {
	if params.PhoneNumbers == "" {
		return http.StatusBadRequest, `{"RequestId" : "", "Code" : "get_PhoneNumber_error", "Message" : "缺少参数，需要PhoneNumbers", "BizId" : ""}`, errors.New("parameter \"PhoneNumbers\" is required")
	}
	if params.SignName == "" {
		return http.StatusBadRequest, `{"RequestId" : "", "Code" : "get_signName_error", "Message" : "缺少参数，需要SignName", "BizId" : ""}`, errors.New("parameter \"SignName\" is required")
	}
	if params.TemplateCode == "" {
		return http.StatusBadRequest, `{"RequestId" : "", "Code" : "get_templateCode_error", "Message" : "缺少参数，需要TemplateCode", "BizId" : ""}`, errors.New("parameter \"TemplateCode\" is required")
	}
	if accessKeyID == "" {
		return http.StatusBadRequest, `{"RequestId" : "", "Code" : "get_accessKeyID_error", "Message" : "缺少参数，需要accessKeyID", "BizId" : ""}`, errors.New("parameter \"accessKeyId\" is required")
	}
	statusCode, data, err := request("SendSms", params, accessKeyID, secretAccessKey, dsc.endpoint)
	return statusCode, data, err
}

type mapList struct {
	l          []string
	normalized map[string]interface{}
}

func request(action string, param Params, accessKeyID, secretAccessKey, endpoint string) (int, string, error) {
	defaults := buildParams(accessKeyID)
	ml := mapList{make([]string, 0, 24), make(map[string]interface{})}
	ml.l = append(ml.l, "AccessKeyId")
	ml.l = append(ml.l, "Action")
	ml.l = append(ml.l, "Format")
	ml.l = append(ml.l, "PhoneNumbers")
	ml.l = append(ml.l, "SignName")
	ml.l = append(ml.l, "SignatureMethod")
	ml.l = append(ml.l, "SignatureNonce")
	ml.l = append(ml.l, "SignatureVersion")
	ml.l = append(ml.l, "TemplateCode")
	ml.l = append(ml.l, "TemplateParam")
	ml.l = append(ml.l, "Timestamp")
	ml.l = append(ml.l, "Version")

	var params struct {
		Action,
		Format,
		SignatureMethod,
		SignatureNonce,
		SignatureVersion,
		Timestamp,
		AccessKeyId,
		Version,
		PhoneNumbers,
		SignName,
		TemplateCode,
		TemplateParam string
	}
	params.Action = action
	params.Format = defaults.Format
	params.SignatureMethod = defaults.SignatureMethod
	params.SignatureNonce = defaults.SignatureNonce
	params.SignatureVersion = defaults.SignatureVersion
	params.Timestamp = defaults.Timestamp
	params.AccessKeyId = defaults.AccessKeyID
	params.Version = defaults.Version
	params.PhoneNumbers = param.PhoneNumbers
	params.SignName = param.SignName
	params.TemplateCode = param.TemplateCode
	params.TemplateParam = param.TemplateParam

	ml.normalized = normalize(params)
	for k, v := range ml.normalized {
		ml.normalized[k] = url.QueryEscape(v.(string))
	}
	canonicalized := strings.Replace(canonicalize(&ml), "+", "%20", -1)

	stringToSign := "GET&" + url.QueryEscape("/") + "&" + url.QueryEscape(canonicalized)
	key := secretAccessKey + "&"
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	ml.l = append(ml.l, "Signature")
	ml.normalized["Signature"] = url.QueryEscape(signature)
	urls := endpoint + "/?" + canonicalize(&ml)
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return http.StatusBadRequest, `{"RequestId" : "", "Code" : "http_request_error", "Message" : "HTTP请求失败", "BizId" : ""}`, err
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	c := http.Client{}
	var resp *http.Response
	resp, err = c.Do(req)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(data), err
}

type buildParam struct {
	Format,
	SignatureMethod,
	SignatureNonce,
	SignatureVersion,
	Timestamp,
	AccessKeyID,
	Version string
}

func buildParams(accessKeyID string) *buildParam {
	return &buildParam{
		"JSON", "HMAC-SHA1", makeNonce(), "1.0", timeStap(), accessKeyID, apiVersion,
	}
}

func makeNonce_() func() string {
	var counter = 0
	var last float64
	machine, _ := os.Hostname()
	pid := os.Getpid()
	rand.Seed(time.Now().UnixNano())
	return func() string {
		val := math.Floor(float64(rand.Float64() * 1000000000000))
		if val == last {
			counter++
		} else {
			counter = 0
		}
		last = val
		uid := machine + strconv.Itoa(pid) + strconv.FormatFloat(val, 'f', -1, 64) + strconv.Itoa(counter)
		m := md5.New()
		io.WriteString(m, uid)
		return hex.EncodeToString(m.Sum(nil))
	}
}

func timeStap() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func normalize(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func canonicalize(ml *mapList) string {
	var params string
	for _, v := range ml.l {
		params += v
		params += "="
		params += ml.normalized[v].(string)
		params += "&"
	}
	params = params[:len(params)-1]
	return params
}
