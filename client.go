package paypal

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Client paasport client define
type Client struct {
	sync.RWMutex
	conf   *Config
	client *http.Client
}

// NewClient create a new paasport client
func NewClient(appId, secret string, configures ...Configurer) (*Client, error) {
	conf := &Config{
		secret: secret,
		appId:  appId,
	}
	for _, configure := range configures {
		configure(conf)
	}
	conf.withDefault()
	client := Client{
		conf: conf,
	}
	return client.init()
}

// Debug open or close debug
func (c *Client) Debug(debug bool) {
	c.conf.debug = debug
}

func (c *Client) init() (*Client, error) {
	transport, err := c.newTransport()
	if err != nil {
		return c, err
	}
	c.client = &http.Client{
		Transport: transport,
	}
	return c, nil
}

func (c *Client) newTransport() (*http.Transport, error) {
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(network, addr, c.conf.HTTPTimeout.ConnectTimeout)
			if err != nil {
				return nil, err
			}
			return newConn(conn, c.conf.HTTPTimeout.ReadWriteTimeout,
				c.conf.HTTPTimeout.LongTimeout), nil

		},
		MaxIdleConns:          c.conf.HTTPMaxConns.MaxIdleConns,
		MaxIdleConnsPerHost:   c.conf.HTTPMaxConns.MaxIdleConnsPerHost,
		IdleConnTimeout:       c.conf.HTTPTimeout.IdleConnTimeout,
		ResponseHeaderTimeout: c.conf.HTTPTimeout.HeaderTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: !c.conf.sslVerify},
	}
	if c.conf.proxyHost != "" {
		proxyUrl, err := url.Parse(c.conf.proxyHost)
		if err != nil {
			return nil, err
		}
		if c.conf.proxyUser != "" {
			if c.conf.proxyPwd != "" {
				proxyUrl.User = url.UserPassword(c.conf.proxyUser, c.conf.proxyPwd)
			} else {
				proxyUrl.User = url.User(c.conf.proxyUser)
			}
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	return transport, nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, in interface{}) (reqBody []byte, req *http.Request, err error) {
	if in != nil {
		reqBody, err = json.Marshal(in)
		if err != nil {
			return reqBody, nil, err
		}
	}
	req, err = http.NewRequestWithContext(ctx, method, c.requestPath(path), bytes.NewBuffer(reqBody))
	c.setDefaultHeader(req)
	return reqBody, req, err
}

// Do send http request
func (c *Client) Do(ctx context.Context, method, path string, in, out interface{}) error {
	reqBody, req, err := c.newRequest(ctx, method, path, in)
	if err != nil {
		return err
	}
	if c.conf.getAccessToken != nil {
		req.Header.Set("Authorization", "Bearer "+c.conf.getAccessToken())
	} else {
		accessToken, err := c.getAccessToken()
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
	return c.do(reqBody, req, out)
}

// DoWithReq send http request
func (c *Client) DoWithReq(reqBody []byte, req *http.Request, out interface{}) error {
	if c.conf.getAccessToken != nil {
		req.Header.Set("Authorization", "Bearer "+c.conf.getAccessToken())
	} else {
		accessToken, err := c.getAccessToken()
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
	return c.do(reqBody, req, out)
}

func (c *Client) getAccessToken() (string, error) {
	c.RLock()
	if c.conf.token != nil {
		if int64(c.conf.token.ExpiresIn) > time.Now().Unix() {
			c.RUnlock()
			return c.conf.token.Token, nil
		}
	}
	c.RUnlock()
	return c.setAccessToken()
}

func (c *Client) setAccessToken() (string, error) {
	c.Lock()
	defer c.Unlock()
	resp, err := c.GetAccessToken(context.Background())
	if err != nil {
		return "", err
	}
	resp.ExpiresIn = time.Now().Unix() + resp.ExpiresIn
	c.conf.token = resp
	return resp.Token, nil
}

// send request
func (c *Client) do(reqBody []byte, req *http.Request, out interface{}) error {
	var respBody []byte
	var err error
	var resp *http.Response
	defer func() {
		c.debug(reqBody, respBody, req, resp)
	}()
	resp, err = c.client.Do(req)
	if err != nil {
		return err
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("%s", string(respBody))
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(respBody, out)
}

func (c *Client) debug(reqBody, respBody []byte, req *http.Request, resp *http.Response) {
	if c.conf.debug {
		if req != nil {
			fmt.Printf("> %s %s %s\n", req.Method, req.URL.RequestURI(), req.Proto)
			fmt.Printf("> Host: %s\n", req.Host)
			for key, header := range req.Header {
				for _, value := range header {
					fmt.Printf("> %s: %s\n", key, value)
				}
			}
			fmt.Println(">")
			fmt.Println(string(reqBody))
			fmt.Println(">")
		}
		if resp != nil {
			fmt.Printf("< %s %s\n", resp.Proto, resp.Status)
			for key, header := range resp.Header {
				for _, value := range header {
					fmt.Printf("< %s: %s\n", key, value)
				}
			}

			fmt.Println("< ")
			fmt.Println(string(respBody))
			fmt.Println("< ")
		}
	}
}

// set default headers
func (c *Client) setDefaultHeader(req *http.Request) {
	req.Header.Add(HTTPHeaderUserAgent, userAgent())
	req.Header.Add(HTTPHeaderContentType, "application/json")
	req.Header.Add("Prefer", "return=representation")
}

// get request path
func (c *Client) requestPath(path string) string {
	endpoint := LIVE_ENDPOINT
	if c.conf.isSandbox {
		endpoint = SANDBOX_ENDPOINT
	}
	return fmt.Sprintf("%s/%s",
		strings.TrimRight(endpoint, "/"),
		strings.TrimLeft(path, "/"))
}
