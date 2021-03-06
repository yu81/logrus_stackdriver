package logrus_stackdriver

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/google-api-go-wrapper/stackdriver/logging"
)

const (
	fieldMessage      = "message"
	fieldLogName      = "log_name"
	fieldHTTPRequest  = "http_request"
	fieldHTTPResponse = "http_response"
)

type dataField struct {
	defaultLogName string
	data           logrus.Fields
	logLevel       logrus.Level
	omitList       map[string]struct{}
}

func newDataField(logName string, entry *logrus.Entry) *dataField {
	if _, ok := entry.Data[fieldMessage]; !ok {
		entry.Data[fieldMessage] = entry.Message
	}

	return &dataField{
		defaultLogName: logName,
		data:           entry.Data,
		logLevel:       entry.Level,
		omitList:       make(map[string]struct{}),
	}
}

func (d *dataField) len() int {
	return len(d.data)
}

func (d *dataField) isOmit(key string) bool {
	_, ok := d.omitList[key]
	return ok
}

func (d *dataField) getRequest() *http.Request {
	if req, ok := d.data[fieldHTTPRequest].(*http.Request); ok {
		d.omitList[fieldHTTPRequest] = struct{}{}
		return req
	}
	return nil
}

func (d *dataField) getResponse() *http.Response {
	if resp, ok := d.data[fieldHTTPResponse].(*http.Response); ok {
		d.omitList[fieldHTTPResponse] = struct{}{}
		return resp
	}
	return nil
}

func (d *dataField) getLogName() string {
	if name, ok := d.data[fieldLogName].(string); ok {
		return name
	}
	return d.defaultLogName
}

func (d *dataField) getSeverity() logging.Severity {
	switch d.logLevel {
	case logrus.DebugLevel:
		return logging.SeverityDebug
	case logrus.InfoLevel:
		return logging.SeverityInfo
	case logrus.WarnLevel:
		return logging.SeverityWarning
	case logrus.ErrorLevel:
		return logging.SeverityError
	case logrus.PanicLevel:
		return logging.SeverityCritical
	case logrus.FatalLevel:
		return logging.SeverityAlert
	default:
		return logging.SeverityDefault
	}
}
