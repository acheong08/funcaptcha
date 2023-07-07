package funcaptcha

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

var (
	jar     = tls_client.NewCookieJar()
	options = []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(360),
		tls_client.WithClientProfile(tls_client.Firefox_110),
		tls_client.WithRandomTLSExtensionOrder(),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}
	client *tls_client.HttpClient
	proxy  = os.Getenv("http_proxy")
)

//goland:noinspection GoUnhandledErrorResult
func init() {
	cli, _ := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	client = &cli

	if proxy != "" {
		(*client).SetProxy(proxy)
	}
}

//goland:noinspection GoUnhandledErrorResult
func init() {
	cli, _ := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	client = &cli
	proxy := os.Getenv("http_proxy")
	if proxy != "" {
		(*client).SetProxy(proxy)
	}
}

//goland:noinspection GoUnusedExportedFunction
func SetTLSClient(cli *tls_client.HttpClient) {
	client = cli
}

func GetOpenAIToken() (string, string, error) {
	hex := randomHex(32)
	token, err := sendRequest(hex, "")
	return token, hex, err
}

func GetOpenAITokenWithBx(bx string) (string, error) {
	hex := randomHex(32)
	return sendRequest(hex, getBdaWitBx(bx))
}

//goland:noinspection SpellCheckingInspection,GoUnhandledErrorResult
func sendRequest(hex, bda string) (string, error) {
	if bda == "" {
		bda = getBDA(hex)
	}
	form := url.Values{
		"bda":          {base64.StdEncoding.EncodeToString([]byte(bda))},
		"public_key":   {"35536E1E-65B4-4D96-9D97-6ADB7EFF8147"},
		"site":         {"https://chat.openai.com"},
		"userbrowser":  {bv},
		"capi_version": {"1.5.2"},
		"capi_mode":    {"lightbox"},
		"style_theme":  {"default"},
		"rnd":          {strconv.FormatFloat(rand.Float64(), 'f', -1, 64)},
	}
	req, _ := http.NewRequest(http.MethodPost, "https://tcr9i.chat.openai.com/fc/gt2/public_key/35536E1E-65B4-4D96-9D97-6ADB7EFF8147", strings.NewReader(form.Encode()))
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://tcr9i.chat.openai.com")
	req.Header.Set("Referer", fmt.Sprintf("https://tcr9i.chat.openai.com/v2/1.5.2/enforcement.%s.html", hex))
	req.Header.Set("User-Agent", bv)
	resp, err := (*client).Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("status code " + resp.Status)
	}

	type arkoseResponse struct {
		Token string `json:"token"`
	}
	var arkose arkoseResponse
	err = json.NewDecoder(resp.Body).Decode(&arkose)
	if err != nil {
		return "", err
	}

	if !strings.Contains(arkose.Token, "sup=1|rid=") {
		return arkose.Token, errors.New("captcha required")
	}

	return arkose.Token, nil
}

//goland:noinspection SpellCheckingInspection
func getBDA(hex string) string {
	bx := fmt.Sprintf(bx_template,
		// getF(),
		getN(),
		// getWh(),
		// webglExtensions,
		// getWebglExtensionsHash(),
		// webglRenderer,
		// webglVendor,
		// webglVersion,
		// webglShadingLanguageVersion,
		// webglAliasedLineWidthRange,
		// webglAliasedPointSizeRange,
		// webglAntialiasing,
		// webglBits,
		// webglMaxParams,
		// webglMaxViewportDims,
		// webglUnmaskedVendor,
		// webglUnmaskedRenderer,
		// webglVsfParams,
		// webglVsiParams,
		// webglFsfParams,
		// webglFsiParams,
		// getWebglHashWebgl(),
		// hex,
		// getFe(),
		// getIfeHash(),
	)
	bt := getBt()
	bw := getBw(bt)
	return encrypt(bx, bv+bw)
}

func getBt() int64 {
	return time.Now().UnixMicro() / 1000000
}

func getBw(bt int64) string {
	return strconv.FormatInt(bt-(bt%21600), 10)
}

func getBdaWitBx(bx string) string {
	bt := getBt()
	bw := getBw(bt)
	return encrypt(bx, bv+bw)
}
