package horizontal

// File       : core.go
// Path       : horizontal
// Time       : CST 2023/4/10 15:03
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var basicTypes = []string{
	"int", "int16", "int64", "int32", "float32", "float64", "string", "bool", "uint64", "uint32", "uint16",
}

// Verify whether the parameters meet the requirements
var validateParams = []string{"require", "pattern", "maxLength", "minLength", "maximum", "minimum", "maxItems", "minItems"}

const (
	taskNotExist     = 200001
	poolNotExist     = 400003
	subsysNotExist   = 500002
	volumeNotExist   = 600002
	snapshotNotExist = 700002
	accountNotExist  = 800001
	groupNotExist    = 800002

	poolAlreadyExists     = 400002
	subsysAlreadyExists   = 500001
	volumeAlreadyExists   = 600001
	snapshotAlreadyExists = 700001
)

var notExistsCodeSet = map[int]struct{}{
	poolNotExist:     {},
	subsysNotExist:   {},
	volumeNotExist:   {},
	snapshotNotExist: {},
	accountNotExist:  {},
	groupNotExist:    {},
	taskNotExist:     {},
}

var existsCodeSet = map[int]struct{}{
	poolAlreadyExists:     {},
	subsysAlreadyExists:   {},
	volumeAlreadyExists:   {},
	snapshotAlreadyExists: {},
}

// Request is used wrap http request
type Request struct {
	port       *int
	method     *string
	endpoint   *string
	path       *string
	apiVersion *string
	Headers    map[string]*string
	query      map[string]*string
	body       io.Reader
	requestId  *string
}

func NewRequest(requestId string) *Request {
	return &Request{requestId: String(requestId)}
}

func (request *Request) SetQuery(query map[string]*string) {
	request.query = query
}

func (request *Request) UpdateQuery(query map[string]*string) {
	request.query = query
}

func (request *Request) GetQuery() map[string]*string {
	return request.query
}

func (request *Request) SetBody(body interface{}) {
	jsonObj := ToJSONString(body)
	request.body = ToReader(jsonObj)
}

func (request *Request) SetPath(path string) {
	request.path = String(path)
}

func (request *Request) GetPath() *string {
	if request.path == nil {
		panic("request path no set")
	}
	return request.path
}

func (request *Request) SetPort(port *int) {
	request.port = port
}

func (request *Request) GetPort() *int {
	if request.port == nil {
		panic("request port no set")
	}
	return request.port
}

func (request *Request) SetApiVersion(apiVersion *string) {
	request.apiVersion = apiVersion
}

func (request *Request) GetApiVersion() *string {
	if request.apiVersion == nil {
		panic("request apiVersion no set")
	}
	return request.apiVersion
}

func (request *Request) SetEndpoint(endpoint *string) {
	request.endpoint = endpoint
}

func (request *Request) GetEndpoint() *string {
	if request.endpoint == nil {
		panic("request endpoint no set")
	}
	return request.endpoint
}

func (request *Request) GetRequestId() *string {
	if request.requestId == nil {
		request.requestId = String(uuid.New().String())
	}
	return request.requestId
}

func (request *Request) SetMethodGET() {
	request.method = String("GET")
}

func (request *Request) SetMethodPOST() {
	request.method = String("POST")
}

func (request *Request) SetMethodPUT() {
	request.method = String("PUT")
}

func (request *Request) SetMethodDELETE() {
	request.method = String("DELETE")
}

func (request *Request) GetMethod() *string {
	if request.method == nil {
		return String("GET")
	}

	return request.method
}

// Response is use d wrap http response
type Response struct {
	Body          io.ReadCloser
	StatusCode    *int
	StatusMessage *string
	Headers       map[string]*string
}

// SDKError struct is used save error code and message
type SDKError struct {
	StatusCode *int
	Code       *int
	Message    *string
	Data       *string
	errMsg     *string
	RequestId  *string
}

func (err *SDKError) SetErrMsg(msg string) {
	err.errMsg = String(msg)
}

func (err *SDKError) Error() string {
	if err.errMsg == nil {
		str := fmt.Sprintf("SDKError:\n   RequestId: %s\n   StatusCode: %d\n   Code: %d\n   Message: %s\n   Data: %s\n",
			StringValue(err.RequestId),
			IntValue(err.StatusCode),
			IntValue(err.Code),
			StringValue(err.Message),
			StringValue(err.Data))
		err.SetErrMsg(str)
	}
	return StringValue(err.errMsg)
}

func (err *SDKError) IsExists() bool {
	_, ok := existsCodeSet[*err.Code]
	return ok
}

func (err *SDKError) IsNotExists() bool {
	_, ok := notExistsCodeSet[*err.Code]
	return ok
}

// RuntimeObject is used for converting http configuration
type RuntimeObject struct {
	ConnectTimeout *int `json:"connectTimeout"`
	ReadTimeout    *int `json:"readTimeout"`
	Backoff        *int `json:"backoff"`
	Retry          *int `json:"retry"`
}

// NewResponse is create response with http response
func NewResponse(httpResponse *http.Response) (res *Response) {
	res = &Response{}
	res.Body = httpResponse.Body
	res.Headers = make(map[string]*string)
	res.StatusCode = Int(httpResponse.StatusCode)
	res.StatusMessage = String(httpResponse.Status)
	return
}

// NewSDKError is used for shortly create SDKError object
func NewSDKError(obj map[string]interface{}) *SDKError {
	err := &SDKError{}
	if val, ok := obj["code"].(int); ok {
		err.Code = Int(val)
	} else if val, ok := obj["code"].(string); ok {
		_val, _ := strconv.ParseInt(val, 10, 64)
		err.Code = Int(int(_val))
	}

	if obj["request_id"] != nil {
		err.RequestId = String(obj["request_id"].(string))
	}

	if obj["message"] != nil {
		err.Message = String(obj["message"].(string))
	}
	if data := obj["data"]; data != nil {
		r := reflect.ValueOf(data)
		if r.Kind().String() == "map" {
			res := make(map[string]interface{})
			tmp := r.MapKeys()
			for _, key := range tmp {
				res[key.String()] = r.MapIndex(key).Interface()
			}
			if statusCode := res["statusCode"]; statusCode != nil {
				if code, ok := statusCode.(int); ok {
					err.StatusCode = Int(code)
				} else if tmp, ok := statusCode.(string); ok {
					code, err_ := strconv.Atoi(tmp)
					if err_ == nil {
						err.StatusCode = Int(code)
					}
				} else if code, ok := statusCode.(*int); ok {
					err.StatusCode = code
				}
			}
		}
		byt, _ := json.Marshal(data)
		err.Data = String(string(byt))
	}

	if statusCode, ok := obj["statusCode"].(int); ok {
		err.StatusCode = Int(statusCode)
	} else if status, ok := obj["statusCode"].(string); ok {
		statusCode, err_ := strconv.Atoi(status)
		if err_ == nil {
			err.StatusCode = Int(statusCode)
		}
	}

	return err
}

func Recover(in interface{}) error {
	if in == nil {
		return nil
	}
	return errors.New(fmt.Sprint(in))
}

// ReadBody is used read response body
func (response *Response) ReadBody() (body []byte, err error) {
	defer response.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)

	for {
		n, err := response.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func ToReader(obj interface{}) io.Reader {
	switch obj.(type) {
	case *string:
		tmp := obj.(*string)
		return strings.NewReader(StringValue(tmp))
	case []byte:
		return strings.NewReader(string(obj.([]byte)))
	case io.Reader:
		return obj.(io.Reader)
	default:
		panic("Invalid Body. Please set a valid Body.")
	}
}

func ToString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func getLocalAddr(localAddr string) (addr *net.TCPAddr) {
	if localAddr != "" {
		addr = &net.TCPAddr{
			IP: []byte(localAddr),
		}
	}
	return addr
}

func ToObject(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	byt, _ := json.Marshal(obj)
	err := json.Unmarshal(byt, &result)
	if err != nil {
		return nil
	}
	return result
}

func AllowRetry(retry interface{}, retryTimes *int) *bool {
	if IntValue(retryTimes) == 0 {
		return Bool(true)
	}
	retryMap, ok := retry.(map[string]interface{})
	if !ok {
		return Bool(false)
	}
	retryable, ok := retryMap["retryable"].(bool)
	if !ok || !retryable {
		return Bool(false)
	}

	maxAttempts, ok := retryMap["maxAttempts"].(int)
	if !ok || maxAttempts < IntValue(retryTimes) {
		return Bool(false)
	}
	return Bool(true)
}

func Merge(args ...interface{}) map[string]*string {
	finalArg := make(map[string]*string)
	for _, obj := range args {
		switch obj.(type) {
		case map[string]*string:
			arg := obj.(map[string]*string)
			for key, value := range arg {
				if value != nil {
					finalArg[key] = value
				}
			}
		default:
			byt, _ := json.Marshal(obj)
			arg := make(map[string]string)
			err := json.Unmarshal(byt, &arg)
			if err != nil {
				return finalArg
			}
			for key, value := range arg {
				if value != "" {
					finalArg[key] = String(value)
				}
			}
		}
	}

	return finalArg
}

func isNil(a interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(a)
	return vi.IsNil()
}

func ToMap(args ...interface{}) map[string]interface{} {
	isNotNil := false
	finalArg := make(map[string]interface{})
	for _, obj := range args {
		if obj == nil {
			continue
		}

		if isNil(obj) {
			continue
		}
		isNotNil = true

		switch obj.(type) {
		case map[string]*string:
			arg := obj.(map[string]*string)
			for key, value := range arg {
				if value != nil {
					finalArg[key] = StringValue(value)
				}
			}
		case map[string]interface{}:
			arg := obj.(map[string]interface{})
			for key, value := range arg {
				if value != nil {
					finalArg[key] = value
				}
			}
		case *string:
			str := obj.(*string)
			arg := make(map[string]interface{})
			err := json.Unmarshal([]byte(StringValue(str)), &arg)
			if err == nil {
				for key, value := range arg {
					if value != nil {
						finalArg[key] = value
					}
				}
			}
			tmp := make(map[string]string)
			err = json.Unmarshal([]byte(StringValue(str)), &tmp)
			if err == nil {
				for key, value := range arg {
					if value != "" {
						finalArg[key] = value
					}
				}
			}
		case []byte:
			byt := obj.([]byte)
			arg := make(map[string]interface{})
			err := json.Unmarshal(byt, &arg)
			if err == nil {
				for key, value := range arg {
					if value != nil {
						finalArg[key] = value
					}
				}
				break
			}
		default:
			val := reflect.ValueOf(obj)
			res := structToMap(val)
			for key, value := range res {
				if value != nil {
					finalArg[key] = value
				}
			}
		}
	}

	if !isNotNil {
		return nil
	}
	return finalArg
}

func structToMap(dataValue reflect.Value) map[string]interface{} {
	out := make(map[string]interface{})
	if !dataValue.IsValid() {
		return out
	}
	if dataValue.Kind().String() == "ptr" {
		if dataValue.IsNil() {
			return out
		}
		dataValue = dataValue.Elem()
	}
	if !dataValue.IsValid() {
		return out
	}
	dataType := dataValue.Type()
	if dataType.Kind().String() != "struct" {
		return out
	}
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, containsNameTag := field.Tag.Lookup("json")
		if !containsNameTag {
			name = field.Name
		} else {
			strs := strings.Split(name, ",")
			name = strs[0]
		}
		fieldValue := dataValue.FieldByName(field.Name)
		if !fieldValue.IsValid() || fieldValue.IsNil() {
			continue
		}
		if field.Type.String() == "io.Reader" || field.Type.String() == "io.Writer" {
			continue
		} else if field.Type.Kind().String() == "struct" {
			out[name] = structToMap(fieldValue)
		} else if field.Type.Kind().String() == "ptr" &&
			field.Type.Elem().Kind().String() == "struct" {
			if fieldValue.Elem().IsValid() {
				out[name] = structToMap(fieldValue)
			}
		} else if field.Type.Kind().String() == "ptr" {
			if fieldValue.IsValid() && !fieldValue.IsNil() {
				out[name] = fieldValue.Elem().Interface()
			}
		} else if field.Type.Kind().String() == "slice" {
			tmp := make([]interface{}, 0)
			num := fieldValue.Len()
			for i := 0; i < num; i++ {
				value := fieldValue.Index(i)
				if !value.IsValid() {
					continue
				}
				if value.Type().Kind().String() == "ptr" &&
					value.Type().Elem().Kind().String() == "struct" {
					if value.IsValid() && !value.IsNil() {
						tmp = append(tmp, structToMap(value))
					}
				} else if value.Type().Kind().String() == "struct" {
					tmp = append(tmp, structToMap(value))
				} else if value.Type().Kind().String() == "ptr" {
					if value.IsValid() && !value.IsNil() {
						tmp = append(tmp, value.Elem().Interface())
					}
				} else {
					tmp = append(tmp, value.Interface())
				}
			}
			if len(tmp) > 0 {
				out[name] = tmp
			}
		} else {
			out[name] = fieldValue.Interface()
		}

	}
	return out
}

func Retryable(err error) *bool {
	if err == nil {
		return Bool(false)
	}
	if realErr, ok := err.(*SDKError); ok {
		if realErr.StatusCode == nil {
			return Bool(false)
		}
		code := IntValue(realErr.StatusCode)
		return Bool(code >= http.StatusInternalServerError)
	}
	return Bool(true)
}

func GetBackoffTime(backoff interface{}, retrytimes *int) *int {
	backoffMap, ok := backoff.(map[string]interface{})
	if !ok {
		return Int(0)
	}
	policy, ok := backoffMap["policy"].(string)
	if !ok || policy == "no" {
		return Int(0)
	}

	period, ok := backoffMap["period"].(int)
	if !ok || period == 0 {
		return Int(0)
	}

	maxTime := math.Pow(2.0, float64(IntValue(retrytimes)))
	return Int(rand.Intn(int(maxTime-1)) * period)
}

func Sleep(backoffTime *int) {
	sleeptime := time.Duration(IntValue(backoffTime)) * time.Second
	time.Sleep(sleeptime)
}

func Validate(params interface{}) error {
	if params == nil {
		return nil
	}
	requestValue := reflect.ValueOf(params)
	if requestValue.IsNil() {
		return nil
	}
	err := validate(requestValue.Elem())
	return err
}

// Verify whether the parameters meet the requirements
func validate(dataValue reflect.Value) error {
	if strings.HasPrefix(dataValue.Type().String(), "*") { // Determines whether the input is a structure object or a pointer object
		if dataValue.IsNil() {
			return nil
		}
		dataValue = dataValue.Elem()
	}

	dataType := dataValue.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		valueField := dataValue.Field(i)
		for _, value := range validateParams {
			err := validateParam(field, valueField, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateParam(field reflect.StructField, valueField reflect.Value, tagName string) error {
	tag, containsTag := field.Tag.Lookup(tagName) // Take out the checked regular expression
	if containsTag && tagName == "require" {
		err := checkRequire(field, valueField)
		if err != nil {
			return err
		}
	}
	if strings.HasPrefix(field.Type.String(), "[]") { // Verify the parameters of the array type
		err := validateSlice(field, valueField, containsTag, tag, tagName)
		if err != nil {
			return err
		}
	} else if valueField.Kind() == reflect.Ptr { // Determines whether it is a pointer object
		err := validatePtr(field, valueField, containsTag, tag, tagName)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateSlice(field reflect.StructField, valueField reflect.Value, containsregexpTag bool, tag, tagName string) error {
	if valueField.IsValid() && !valueField.IsNil() { // Determines whether the parameter has a value
		if containsregexpTag {
			if tagName == "maxItems" {
				err := checkMaxItems(field, valueField, tag)
				if err != nil {
					return err
				}
			}

			if tagName == "minItems" {
				err := checkMinItems(field, valueField, tag)
				if err != nil {
					return err
				}
			}
		}

		for m := 0; m < valueField.Len(); m++ {
			elementValue := valueField.Index(m)
			if elementValue.Type().Kind() == reflect.Ptr { // Determines whether the child elements of an array are of a basic type
				err := validatePtr(field, elementValue, containsregexpTag, tag, tagName)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func validatePtr(field reflect.StructField, elementValue reflect.Value, containsregexpTag bool, tag, tagName string) error {
	if elementValue.IsNil() {
		return nil
	}
	if isFilterType(elementValue.Elem().Type().String(), basicTypes) {
		if containsregexpTag {
			if tagName == "pattern" {
				err := checkPattern(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "maxLength" {
				err := checkMaxLength(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "minLength" {
				err := checkMinLength(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "maximum" {
				err := checkMaximum(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "minimum" {
				err := checkMinimum(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}
		}
	} else {
		err := validate(elementValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkRequire(field reflect.StructField, valueField reflect.Value) error {
	name, _ := field.Tag.Lookup("json")
	strs := strings.Split(name, ",")
	name = strs[0]
	if !valueField.IsNil() && valueField.IsValid() {
		return nil
	}
	return errors.New(name + " should be setted")
}

func checkPattern(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		value := valueField.String()
		r, _ := regexp.Compile("^" + tag + "$")
		if match := r.MatchString(value); !match { // Determines whether the parameter value satisfies the regular expression or not, and throws an error
			return errors.New(value + " is not matched " + tag)
		}
	}
	return nil
}

func checkMaxItems(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		maxItems, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if maxItems < length {
			errMsg := fmt.Sprintf("The length of %s is %d which is more than %d", field.Name, length, maxItems)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMinItems(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() {
		minItems, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if minItems > length {
			errMsg := fmt.Sprintf("The length of %s is %d which is less than %d", field.Name, length, minItems)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMaxLength(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		maxLength, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if valueField.Kind().String() == "string" {
			length = strings.Count(valueField.String(), "") - 1
		}
		if maxLength < length {
			errMsg := fmt.Sprintf("The length of %s is %d which is more than %d", field.Name, length, maxLength)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMinLength(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() {
		minLength, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if valueField.Kind().String() == "string" {
			length = strings.Count(valueField.String(), "") - 1
		}
		if minLength > length {
			errMsg := fmt.Sprintf("The length of %s is %d which is less than %d", field.Name, length, minLength)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMaximum(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		maximum, err := strconv.ParseFloat(tag, 64)
		if err != nil {
			return err
		}
		byt, _ := json.Marshal(valueField.Interface())
		num, err := strconv.ParseFloat(string(byt), 64)
		if err != nil {
			return err
		}
		if maximum < num {
			errMsg := fmt.Sprintf("The size of %s is %f which is greater than %f", field.Name, num, maximum)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMinimum(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		minimum, err := strconv.ParseFloat(tag, 64)
		if err != nil {
			return err
		}

		byt, _ := json.Marshal(valueField.Interface())
		num, err := strconv.ParseFloat(string(byt), 64)
		if err != nil {
			return err
		}
		if minimum > num {
			errMsg := fmt.Sprintf("The size of %s is %f which is less than %f", field.Name, num, minimum)
			return errors.New(errMsg)
		}
	}
	return nil
}

// Determines whether realType is in filterTypes
func isFilterType(realType string, filterTypes []string) bool {
	for _, value := range filterTypes {
		if value == realType {
			return true
		}
	}
	return false
}

func IsUnset(val interface{}) *bool {
	if val == nil {
		return Bool(true)
	}

	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		return Bool(v.IsNil())
	}

	valType := reflect.TypeOf(val)
	valZero := reflect.Zero(valType)
	return Bool(valZero == v)
}

func ToJSONString(a interface{}) *string {
	switch v := a.(type) {
	case *string:
		return v
	case string:
		return String(v)
	case []byte:
		return String(string(v))
	case io.Reader:
		byt, err := io.ReadAll(v)
		if err != nil {
			return nil
		}
		return String(string(byt))
	}
	byt, err := json.Marshal(a)
	if err != nil {
		return nil
	}
	return String(string(byt))
}

func AssertAsMap(a interface{}) (_result map[string]interface{}, _err error) {
	r := reflect.ValueOf(a)
	if r.Kind().String() != "map" {
		return nil, errors.New(fmt.Sprintf("%v is not a map[string]interface{}", a))
	}

	res := make(map[string]interface{})
	tmp := r.MapKeys()
	for _, key := range tmp {
		res[key.String()] = r.MapIndex(key).Interface()
	}

	return res, nil
}

func ReadAsJSON(body io.Reader) (result interface{}, err error) {
	byt, err := io.ReadAll(body)
	if err != nil {
		return
	}
	if string(byt) == "" {
		return
	}
	r, ok := body.(io.ReadCloser)
	if ok {
		_ = r.Close()
	}
	d := json.NewDecoder(bytes.NewReader(byt))
	d.UseNumber()
	err = d.Decode(&result)
	return
}
