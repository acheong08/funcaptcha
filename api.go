package funcaptcha

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

const arkPreURL = "https://tcr9i.chat.openai.com/fc/gt2/"

var arkURLIns, _ = url.Parse(arkPreURL)

var initVer, initHex string

type arkReq struct {
	arkURL     string
	arkBx      string
	arkHeader  http.Header
	arkBody    url.Values
	arkCookies []*http.Cookie
	userAgent  string
}

var (
	jar     = tls_client.NewCookieJar()
	options = []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(360),
		tls_client.WithClientProfile(profiles.Chrome_117),
		tls_client.WithRandomTLSExtensionOrder(),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}
	client    *tls_client.HttpClient
	proxy     = os.Getenv("http_proxy")
	authArks  []*arkReq
	chat3Arks []*arkReq
	chat4Arks []*arkReq
)

type kvPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type cookie struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Expires string `json:"expires"`
}
type postBody struct {
	Params []kvPair `json:"params"`
}
type request struct {
	URL      string   `json:"url"`
	Headers  []kvPair `json:"headers,omitempty"`
	PostData postBody `json:"postData,omitempty"`
	Cookies  []cookie `json:"cookies,omitempty"`
}
type entry struct {
	StartedDateTime string  `json:"startedDateTime"`
	Request         request `json:"request"`
}
type logData struct {
	Entries []entry `json:"entries"`
}
type HARData struct {
	Log logData `json:"log"`
}

func readHAR() {
	// 指定需要读取的文件夹路径
	dirPath := "./harPool"
	var harPath []string
	// 打开文件夹并读取其内容
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是否为普通文件（非文件夹）
		if !info.IsDir() {
			// 获取文件后缀名
			ext := filepath.Ext(info.Name())
			if ext == ".har" {
				harPath = append(harPath, path)
			}
		}
		return nil
	})
	if err != nil {
		println("Error: please put HAR files in harPool directory!")
	}
	for _, path := range harPath {
		file, err := os.ReadFile(path)
		if err != nil {
			return
		}
		var harFile HARData
		err = json.Unmarshal(file, &harFile)
		if err != nil {
			println("Error: not a HAR file!")
			return
		}
		for _, v := range harFile.Log.Entries {
			if strings.HasPrefix(v.Request.URL, arkPreURL) {
				var tmpArk arkReq
				tmpArk.arkURL = v.Request.URL
				if v.StartedDateTime == "" {
					println("Error: no arkose request!")
					continue
				}
				t, _ := time.Parse(time.RFC3339, v.StartedDateTime)
				bw := getBw(t.Unix())
				fallbackBw := getBw(t.Unix() - 21600)
				tmpArk.arkHeader = make(http.Header)
				for _, h := range v.Request.Headers {
					// arkHeader except cookie & content-length
					if !strings.EqualFold(h.Name, "content-length") && !strings.EqualFold(h.Name, "cookie") && !strings.HasPrefix(h.Name, ":") {
						tmpArk.arkHeader.Set(h.Name, h.Value)
						if strings.EqualFold(h.Name, "user-agent") {
							tmpArk.userAgent = h.Value
						}
					}
				}
				tmpArk.arkCookies = []*http.Cookie{}
				for _, cookie := range v.Request.Cookies {
					expire, _ := time.Parse(time.RFC3339, cookie.Expires)
					if expire.After(time.Now()) {
						tmpArk.arkCookies = append(tmpArk.arkCookies, &http.Cookie{Name: cookie.Name, Value: cookie.Value, Expires: expire.UTC()})
					}
				}
				var arkType string
				tmpArk.arkBody = make(url.Values)
				for _, p := range v.Request.PostData.Params {
					// arkBody except bda & rnd
					if p.Name == "bda" {
						cipher, err := url.QueryUnescape(p.Value)
						if err != nil {
							panic(err)
						}
						tmpArk.arkBx = Decrypt(cipher, tmpArk.userAgent+bw, tmpArk.userAgent+fallbackBw)
					} else if p.Name != "rnd" {
						query, err := url.QueryUnescape(p.Value)
						if err != nil {
							panic(err)
						}
						tmpArk.arkBody.Set(p.Name, query)
						if p.Name == "public_key" {
							if query == "0A1D34FC-659D-4E23-B17B-694DCFCF6A6C" {
								arkType = "auth"
								authArks = append(authArks, &tmpArk)
							} else if query == "3D86FBBA-9D22-402A-B512-3420086BA6CC" {
								arkType = "chat3"
								chat3Arks = append(chat3Arks, &tmpArk)
							} else if query == "35536E1E-65B4-4D96-9D97-6ADB7EFF8147" {
								arkType = "chat4"
								chat4Arks = append(chat4Arks, &tmpArk)
							}
						}
					}
				}
				if tmpArk.arkBx != "" {
					println("success read " + arkType + " arkose")
				} else {
					println("failed to decrypt HAR file")
				}
			}
		}
	}
}

//goland:noinspection GoUnhandledErrorResult
func init() {
	initVer = "1.5.4"
	initHex = "cd12da708fe6cbe6e068918c38de2ad9" // should be fixed associated with version.
	readHAR()
	cli, _ := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	client = &cli
	if proxy != "" {
		(*client).SetProxy(proxy)
	}
}

//goland:noinspection GoUnusedExportedFunction
func SetTLSClient(cli *tls_client.HttpClient) {
	client = cli
}

func GetOpenAIAuthToken(puid string, proxy string) (string, error) {
	token, err := sendRequest(0, "", puid, proxy)
	return token, err
}

func GetOpenAIAuthTokenWithBx(bx string, puid string, proxy string) (string, error) {
	token, err := sendRequest(0, getBdaWitBx(bx), puid, proxy)
	return token, err
}

func GetOpenAIToken(version int, puid string, proxy string) (string, error) {
	token, err := sendRequest(version, "", puid, proxy)
	return token, err
}

func GetOpenAITokenWithBx(version int, bx string, puid string, proxy string) (string, error) {
	token, err := sendRequest(version, getBdaWitBx(bx), puid, proxy)
	return token, err
}

//goland:noinspection SpellCheckingInspection,GoUnhandledErrorResult
func sendRequest(arkType int, bda string, puid string, proxy string) (string, error) {
	var tmpArk *arkReq
	if arkType == 0 {
		tmpArk = authArks[0]
		authArks = append(authArks[1:], authArks[0])
	} else if arkType == 3 {
		tmpArk = chat3Arks[0]
		chat3Arks = append(chat3Arks[1:], chat3Arks[0])
	} else {
		tmpArk = chat4Arks[0]
		chat4Arks = append(chat4Arks[1:], chat4Arks[0])
	}
	if tmpArk == nil || tmpArk.arkBx == "" || len(tmpArk.arkBody) == 0 || len(tmpArk.arkHeader) == 0 {
		return "", errors.New("a valid HAR file required")
	}
	if proxy != "" {
		(*client).SetProxy(proxy)
	}
	if bda == "" {
		bda = getBDA(tmpArk)
	}
	tmpArk.arkBody.Set("bda", base64.StdEncoding.EncodeToString([]byte(bda)))
	tmpArk.arkBody.Set("rnd", strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	req, _ := http.NewRequest(http.MethodPost, tmpArk.arkURL, strings.NewReader(tmpArk.arkBody.Encode()))
	req.Header = tmpArk.arkHeader.Clone()
	(*client).GetCookieJar().SetCookies(arkURLIns, tmpArk.arkCookies)
	if puid != "" {
		req.Header.Set("cookie", "_puid="+puid+";")
	}
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
	// Check if rid is empty
	if !strings.Contains(arkose.Token, "sup=1|rid=") {
		return arkose.Token, errors.New("captcha required")
	}

	return arkose.Token, nil
}

//goland:noinspection SpellCheckingInspection
func getBDA(arkReq *arkReq) string {
	var bx string = arkReq.arkBx
	if bx == "" {
		bx = fmt.Sprintf(bx_template,
			getF(),
			getN(),
			getWh(),
			webglExtensions,
			getWebglExtensionsHash(),
			webglRenderer,
			webglVendor,
			webglVersion,
			webglShadingLanguageVersion,
			webglAliasedLineWidthRange,
			webglAliasedPointSizeRange,
			webglAntialiasing,
			webglBits,
			webglMaxParams,
			webglMaxViewportDims,
			webglUnmaskedVendor,
			webglUnmaskedRenderer,
			webglVsfParams,
			webglVsiParams,
			webglFsfParams,
			webglFsiParams,
			getWebglHashWebgl(),
			initVer,
			initHex,
			getFe(),
			getIfeHash(),
		)
	} else {
		re := regexp.MustCompile(`"key"\:"n","value"\:"\S*?"`)
		bx = re.ReplaceAllString(bx, `"key":"n","value":"`+getN()+`"`)
	}
	bt := getBt()
	bw := getBw(bt)
	return Encrypt(bx, arkReq.userAgent+bw)
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
	return Encrypt(bx, bv+bw)
}
