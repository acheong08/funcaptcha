package funcaptcha

import http "github.com/bogdanfinn/fhttp"

var headers = http.Header{
	"Accept":           []string{"*/*"},
	"Accept-Encoding":  []string{"gzip, deflate, br"},
	"Accept-Language":  []string{"en-US,en;q=0.5"},
	"Cache-Control":    []string{"no-cache"},
	"Connection":       []string{"keep-alive"},
	"Content-Type":     []string{"application/x-www-form-urlencoded; charset=UTF-8"},
	"Host":             []string{"client-api.arkoselabs.com"},
	"Origin":           []string{"https://client-api.arkoselabs.com"},
	"User-Agent":       []string{bv},
	"X-Requested-With": []string{"XMLHttpRequest"},
}

const bx_template string = `
[{"key":"api_type","value":"js"},{"key":"p","value":1},{"key":"f","value":"7496d4f710164195fca37d1429b078ea"},{"key":"n","value":"%s"},{"key":"wh","value":"80b13fd48b8da8e4157eeb6f9e9fbedb|5ab5738955e0611421b686bc95655ad0"},{"key":"enhanced_fp","value":[{"key":"webgl_extensions","value":null},{"key":"webgl_extensions_hash","value":null},{"key":"webgl_renderer","value":null},{"key":"webgl_vendor","value":null},{"key":"webgl_version","value":null},{"key":"webgl_shading_language_version","value":null},{"key":"webgl_aliased_line_width_range","value":null},{"key":"webgl_aliased_point_size_range","value":null},{"key":"webgl_antialiasing","value":null},{"key":"webgl_bits","value":null},{"key":"webgl_max_params","value":null},{"key":"webgl_max_viewport_dims","value":null},{"key":"webgl_unmasked_vendor","value":null},{"key":"webgl_unmasked_renderer","value":null},{"key":"webgl_vsf_params","value":null},{"key":"webgl_vsi_params","value":null},{"key":"webgl_fsf_params","value":null},{"key":"webgl_fsi_params","value":null},{"key":"webgl_hash_webgl","value":null},{"key":"user_agent_data_brands","value":null},{"key":"user_agent_data_mobile","value":null},{"key":"navigator_connection_downlink","value":null},{"key":"navigator_connection_downlink_max","value":null},{"key":"network_info_rtt","value":null},{"key":"network_info_save_data","value":null},{"key":"network_info_rtt_type","value":null},{"key":"screen_pixel_depth","value":24},{"key":"navigator_device_memory","value":null},{"key":"navigator_languages","value":"en-US,en"},{"key":"window_inner_width","value":0},{"key":"window_inner_height","value":0},{"key":"window_outer_width","value":0},{"key":"window_outer_height","value":0},{"key":"browser_detection_firefox","value":true},{"key":"browser_detection_brave","value":false},{"key":"audio_codecs","value":"{\"ogg\":\"probably\",\"mp3\":\"maybe\",\"wav\":\"probably\",\"m4a\":\"maybe\",\"aac\":\"maybe\"}"},{"key":"video_codecs","value":"{\"ogg\":\"probably\",\"h264\":\"probably\",\"webm\":\"probably\",\"mpeg4v\":\"\",\"mpeg4a\":\"\",\"theora\":\"\"}"},{"key":"media_query_dark_mode","value":false},{"key":"headless_browser_phantom","value":false},{"key":"headless_browser_selenium","value":false},{"key":"headless_browser_nightmare_js","value":false},{"key":"document__referrer","value":""},{"key":"window__ancestor_origins","value":null},{"key":"window__tree_index","value":[1]},{"key":"window__tree_structure","value":"[[],[]]"},{"key":"window__location_href","value":"https://tcr9i.chat.openai.com/v2/1.5.2/enforcement.64b3a4e29686f93d52816249ecbf9857.html#35536E1E-65B4-4D96-9D97-6ADB7EFF8147"},{"key":"client_config__sitedata_location_href","value":"https://chat.openai.com/"},{"key":"client_config__surl","value":"https://tcr9i.chat.openai.com"},{"key":"mobile_sdk__is_sdk"},{"key":"client_config__language","value":null},{"key":"audio_fingerprint","value":"35.73833402246237"}]},{"key":"fe","value":["DNT:1","L:en-US","D:24","PR:1","S:0,0","AS:false","TO:0","SS:true","LS:true","IDB:true","B:false","ODB:false","CPUC:unknown","PK:Linux x86_64","CFP:-1429721473","FR:false","FOS:false","FB:false","JSF:Arial,Arial Narrow,Bitstream Vera Sans Mono,Bookman Old Style,Century Schoolbook,Courier,Courier New,Helvetica,MS Gothic,MS PGothic,Palatino,Palatino Linotype,Times,Times New Roman","P:Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,PDF Viewer,WebKit built-in PDF","T:0,false,false","H:2","SWF:false"]},{"key":"ife_hash","value":"df082a66c67a12f2e8d144c622c26115"},{"key":"cs","value":1},{"key":"jsbd","value":"{\"HL\":2,\"NCE\":true,\"DT\":\"\",\"NWD\":\"false\",\"DOTO\":1,\"DMTO\":1}"}]
`

// LinkedHashMap
var fe = []map[string]interface{}{
	{"DNT": "1"},
	{"L": "en-US"},
	{"D": 24},
	{"PR": 1}, // this.getPixelRatio();
	{"S": "0;0"},
	{"AS": false},
	{"TO": 0},
	{"SS": true},
	{"LS": true},
	{"IDB": true},
	{"B": false},
	{"ODB": true},
	{"CPUC": "unknown"},
	{"PK": "Linux x86_64"},
	{"CFP": cfp},
	{"FR": false},
	{"FOS": false},
	{"FB": false},
	{"JSF": "Arial,Arial Narrow,Bitstream Vera Sans Mono,Bookman Old Style,Century Schoolbook,Courier,Courier New,Helvetica,MS Gothic,MS PGothic,Palatino,Palatino Linotype,Times,Times New Roman"},
	{"P": p},
	{"T": "0;false;false"},
	{"H": 2},
	{"SWF": false},
}

const (
	p  = "Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,PDF Viewer,WebKit built-in PDF"
	bv = "Mozilla/5.0 (X11; Linux x86_64; rv:114.0) Gecko/20100101 Firefox/114.0"
)

const (
	cfp = "-1471542875"
)
