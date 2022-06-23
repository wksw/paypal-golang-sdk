package paypal

import (
	"time"
)

// Config paasport configuration
type Config struct {
	appId          string
	secret         string
	apiVersion     string
	userAgent      string
	proxyHost      string
	proxyUser      string
	proxyPwd       string
	HTTPTimeout    HTTPTimeout
	HTTPMaxConns   HTTPMaxConns
	sslVerify      bool
	isSandbox      bool
	debug          bool
	token          *TokenResponse
	getAccessToken func() string
	setAccessToken func(string)
}

// HTTPTimeout http timeout define
type HTTPTimeout struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	HeaderTimeout    time.Duration
	LongTimeout      time.Duration
	IdleConnTimeout  time.Duration
}

// HTTPMaxConns max idle connections
type HTTPMaxConns struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

// Configurer 配置
type Configurer func(conf *Config)

func (c *Config) withDefault() {
	if c.HTTPTimeout.ReadWriteTimeout == 0 {
		c.HTTPTimeout.ReadWriteTimeout = 5 * time.Second
	}
}

// WithSslVerify set sslVerify
func WithSslVerify(sslVerify bool) Configurer {
	return func(conf *Config) {
		conf.sslVerify = sslVerify
	}
}

// WithSandbox set sandbox
func WithSandbox(isSandbox bool) Configurer {
	return func(conf *Config) {
		conf.isSandbox = isSandbox
	}
}

// WithGetAccessToken set get access token func
func WithGetAccessToken(h func() string) Configurer {
	return func(conf *Config) {
		conf.getAccessToken = h
	}
}

// WithSetAccessToken set access token func
func WithSetAccessToken(h func(string)) Configurer {
	return func(conf *Config) {
		conf.setAccessToken = h
	}
}

// WithTimeout set timeout
func WithTimeout(httpTimeout HTTPTimeout) Configurer {
	return func(conf *Config) {
		conf.HTTPTimeout = httpTimeout
	}
}

// WithProxy set proxy
func WithProxy(proxyHost, proxyUser, proxyPwd string) Configurer {
	return func(conf *Config) {
		conf.proxyHost = proxyHost
		conf.proxyUser = proxyUser
		conf.proxyPwd = proxyPwd
	}
}

// WithMaxConnections set max connections
func WithMaxConnections(httpMaxConns HTTPMaxConns) Configurer {
	return func(conf *Config) {
		conf.HTTPMaxConns = httpMaxConns
	}
}
