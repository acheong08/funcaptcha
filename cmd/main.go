package main

import (
	"fmt"
	"log"

	"github.com/acheong08/funcaptcha"
)

func main() {
	//token, err := funcaptcha.GetOpenAITokenWithBx(`[{"key":"api_type","value":"js"},{"key":"p","value":1},{"key":"f","value":"cdb0697262ceb9aa4e938ae5d9697efd"},{"key":"n","value":"MTY4OTQ0MjY0Mw=="},{"key":"wh","value":"15fb0b896f4534b3970269c17e188322|5ab5738955e0611421b686bc95655ad0"},{"key":"enhanced_fp","value":[{"key":"webgl_extensions","value":"ANGLE_instanced_arrays;EXT_blend_minmax;EXT_color_buffer_half_float;EXT_float_blend;EXT_frag_depth;EXT_shader_texture_lod;EXT_sRGB;EXT_texture_compression_bptc;EXT_texture_compression_rgtc;EXT_texture_filter_anisotropic;OES_element_index_uint;OES_fbo_render_mipmap;OES_standard_derivatives;OES_texture_float;OES_texture_float_linear;OES_texture_half_float;OES_texture_half_float_linear;OES_vertex_array_object;WEBGL_color_buffer_float;WEBGL_compressed_texture_etc;WEBGL_compressed_texture_s3tc;WEBGL_compressed_texture_s3tc_srgb;WEBGL_debug_renderer_info;WEBGL_debug_shaders;WEBGL_depth_texture;WEBGL_draw_buffers;WEBGL_lose_context"},{"key":"webgl_extensions_hash","value":"ccc5c4979d89351fef1dcc0582cdb3d2"},{"key":"webgl_renderer","value":"NVIDIA GeForce GTX 980/PCIe/SSE2"},{"key":"webgl_vendor","value":"Mozilla"},{"key":"webgl_version","value":"WebGL 1.0"},{"key":"webgl_shading_language_version","value":"WebGL GLSL ES 1.0"},{"key":"webgl_aliased_line_width_range","value":"[1, 10]"},{"key":"webgl_aliased_point_size_range","value":"[1, 2047]"},{"key":"webgl_antialiasing","value":"yes"},{"key":"webgl_bits","value":"8,8,24,8,8,0"},{"key":"webgl_max_params","value":"16,192,32768,1024,32768,32,32768,32,16,32,1024"},{"key":"webgl_max_viewport_dims","value":"[32768, 32768]"},{"key":"webgl_unmasked_vendor","value":"NVIDIA Corporation"},{"key":"webgl_unmasked_renderer","value":"NVIDIA GeForce GTX 980/PCIe/SSE2"},{"key":"webgl_vsf_params","value":"23,127,127,23,127,127,23,127,127"},{"key":"webgl_vsi_params","value":"0,24,24,0,24,24,0,24,24"},{"key":"webgl_fsf_params","value":"23,127,127,23,127,127,23,127,127"},{"key":"webgl_fsi_params","value":"0,24,24,0,24,24,0,24,24"},{"key":"webgl_hash_webgl","value":"5890c452638eafb176df4e27cce6e5a3"},{"key":"user_agent_data_brands","value":null},{"key":"user_agent_data_mobile","value":null},{"key":"navigator_connection_downlink","value":null},{"key":"navigator_connection_downlink_max","value":null},{"key":"network_info_rtt","value":null},{"key":"network_info_save_data","value":null},{"key":"network_info_rtt_type","value":null},{"key":"screen_pixel_depth","value":24},{"key":"navigator_device_memory","value":null},{"key":"navigator_languages","value":"en-US,en"},{"key":"window_inner_width","value":0},{"key":"window_inner_height","value":0},{"key":"window_outer_width","value":631},{"key":"window_outer_height","value":1039},{"key":"browser_detection_firefox","value":true},{"key":"browser_detection_brave","value":false},{"key":"audio_codecs","value":"{\"ogg\":\"probably\",\"mp3\":\"maybe\",\"wav\":\"probably\",\"m4a\":\"maybe\",\"aac\":\"maybe\"}"},{"key":"video_codecs","value":"{\"ogg\":\"probably\",\"h264\":\"probably\",\"webm\":\"probably\",\"mpeg4v\":\"\",\"mpeg4a\":\"\",\"theora\":\"\"}"},{"key":"media_query_dark_mode","value":true},{"key":"headless_browser_phantom","value":false},{"key":"headless_browser_selenium","value":true},{"key":"headless_browser_nightmare_js","value":false},{"key":"document__referrer","value":"http://127.0.0.1:8000/"},{"key":"window__ancestor_origins","value":null},{"key":"window__tree_index","value":[0]},{"key":"window__tree_structure","value":"[[]]"},{"key":"window__location_href","value":"https://tcr9i.chat.openai.com/v2/1.5.2/enforcement.64b3a4e29686f93d52816249ecbf9857.html#35536E1E-65B4-4D96-9D97-6ADB7EFF8147"},{"key":"client_config__sitedata_location_href","value":"http://127.0.0.1:8000/arkose.html"},{"key":"client_config__surl","value":"https://tcr9i.chat.openai.com"},{"key":"mobile_sdk__is_sdk"},{"key":"client_config__language","value":null},{"key":"audio_fingerprint","value":"35.73833402246237"}]},{"key":"fe","value":["DNT:unspecified","L:en-US","D:24","PR:1","S:1920,1080","AS:1920,1080","TO:-480","SS:true","LS:true","IDB:true","B:false","ODB:false","CPUC:unknown","PK:Linux x86_64","CFP:699685943","FR:false","FOS:false","FB:false","JSF:Arial,Courier New,Times New Roman","P:Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,PDF Viewer,WebKit built-in PDF","T:0,false,false","H:16","SWF:false"]},{"key":"ife_hash","value":"63562827dd8cdcf172844085452eb5f4"},{"key":"cs","value":1},{"key":"jsbd","value":"{\"HL\":1,\"NCE\":true,\"DT\":\"\",\"NWD\":\"true\",\"DOTO\":1,\"DMTO\":1}"}]`, "", "")
	//log.Println(token)
	//
	//if err == nil {
	//	return
	//}
	token := "155179050361eafc5.5354759702|r=us-west-2|meta=3|metabgclr=transparent|metaiconclr=%23757575|guitextcolor=%23000000|pk=0A1D34FC-659D-4E23-B17B-694DCFCF6A6C|at=40|ag=101|cdn_url=https%3A%2F%2Ftcr9i.chat.openai.com%2Fcdn%2Ffc|lurl=https%3A%2F%2Faudio-us-west-2.arkoselabs.com|surl=https%3A%2F%2Ftcr9i.chat.openai.com|smurl=https%3A%2F%2Ftcr9i.chat.openai.com%2Fcdn%2Ffc%2Fassets%2Fstyle-manager"
	hex := "fbfc14b0d793c6ef8359e0e4b4a91f67"
	// Start a challenge
	session, err := funcaptcha.StartChallenge(token, hex)
	if err != nil {
		log.Fatalf("error starting challenge: %v\n", err)
	}
	log.Println("Challenge started!")

	apiBreaker, err := session.RequestChallenge(false)
	if err != nil {
		log.Fatalf("error requesting challenge: %v\n", err)
	}
	log.Println(session.ConciseChallenge)
	log.Println("Downloading challenge")
	_, err = funcaptcha.DownloadChallenge(session.ConciseChallenge.URLs, false)
	if err != nil {
		log.Fatalf("error downloading challenge: %v\n", err)
	}
	log.Println("Challenge downloaded!")
	// User input here
	fmt.Println("Please enter the index of the image based on the following instructions:")
	fmt.Println(session.ConciseChallenge.Instructions)

	ints := make([]int, len(session.ConciseChallenge.URLs))
	fmt.Println("Enter the integers:")
	for i := 0; i < len(ints); i++ {
		fmt.Scan(&ints[i])
	}
	log.Println(ints)
	err = session.SubmitAnswer(ints, false, apiBreaker)
	if err != nil {
		log.Fatalf("error submitting answer: %v\n", err)
	}
}
