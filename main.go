package main

import (
	"fmt"
	"log"
	"bytes"
	"errors"
	"net/http"
	"crypto/aes"
	"crypto/cipher"
	"html/template"
	"encoding/base64"
)

// 默认shiro key
var key_list = []string{"wGiHplamyXlVB11UXWol8g==", "7AvVhmFLUs0KTA3Kprsdag==", "cGhyYWNrY3RmREUhfiMkZA==", "NGk/3cQ6F5/UNPRh8LpMIg==", "6Zm+6I2j5Y+R5aS+5ZOlAA==", "r0e3c16IdVkouZgk1TKVMg==", "OY//C4rhfwNxCQAQCrQQ1Q==", "c+3hFGPjbgzGdrC+MHgoRQ==", "ikB3y6O9BpimrZLB3rca0w==", "lT2UvDUmQwewm6mMoiw4Ig==", "3qDVdLawoIr1xFd6ietnwg==", "9AvVhmFLUs0KTA3Kprsdag==", "3AvVhmFLUs0KTA3Kprsdag==", "A7UzJgh1+EWj5oBFi+mSgw==", "5J7bIJIV0LQSN3c9LPitBQ==", "CrownKey==a12d/dakdad", "HWrBltGvEZc14h9VpMvZWw==", "WuB+y2gcHRnY2Lg9+Aqmqg==", "O4pdf+7e+mZe8NyxMTPJmQ==", "2itfW92XazYRi5ltW0M2yA==", "U3ByaW5nQmxhZGUAAAAAAA==", "Z3VucwAAAAAAAAAAAAAAAA==", "cGljYXMAAAAAAAAAAAAAAA==", "Bf7MfkNR0axGGptozrebag==", "V2hhdCBUaGUgSGVsbAAAAA==", "d2ViUmVtZW1iZXJNZUtleQ==", "YI1+nBV//m7ELrIyDHm6DQ==", "ertVhmFLUs0KTA3Kprsdag==", "MPdCMZ9urzEA50JDlDYYDg==", "WkhBTkdYSUFPSEVJX0NBVA==", "s0KTA3mFLUprK4AvVhsdag==", "8AvVhmFLUs0KTA3Kprsdag==", "WcfHGU25gNnTxTlmJMeSpw==", "2A2V+RFLUs+eTA3Kpr+dag==", "4BvVhmFLUs0KTA3Kprsdag==", "L7RioUULEFhRyxM7a2R/Yg==", "ClLk69oNcA3m+s0jIMIkpg==", "rPNqM6uKFCyaL10AK51UkQ==", "fCq+/xW488hMTCD+cmJ3aQ==", "bWluZS1hc3NldC1rZXk6QQ==", "HoTP07fJPKIRLOWoVXmv+Q==", "Q01TX0JGTFlLRVlfMjAxOQ==", "Z3h6eWd4enklMjElMjElMjE=", "SDKOLKn2J1j/2BHjeZwAoQ==", "25BsmdYwjnfcWmnhAciDDg==", "5aaC5qKm5oqA5pyvAAAAAA==", "Is9zJ3pzNh2cgTHB4ua3+Q==", "empodDEyMwAAAAAAAAAAAA==", "0AvVhmFLUs0KTA3Kprsdag==", "GAevYnznvgNCURavBhCr1w==", "3JvYhmBLUs0ETA5Kprsdag==", "aU1pcmFjbGVpTWlyYWNsZQ==", "fsHspZw/92PrS3XrPW+vxw==", "sHdIjUN6tzhl8xZMG3ULCQ==", "4AvVhmFLUs0KTA3Kprsdag==", "yNeUgSzL/CfiWw1GALg6Ag==", "Y1JxNSPXVwMkyvES/kJGeQ==", "a3dvbmcAAAAAAAAAAAAAAA==", "cmVtZW1iZXJNZQAAAAAAAA==", "8BvVhmFLUs0KTA3Kprsdag==", "MzVeSkYyWTI2OFVLZjRzZg==", "f/SY5TIve5WWzT4aQlABJA==", "U3BAbW5nQmxhZGUAAAAAAA==", "2cVtiE83c4lIrELJwKGJUw==", "yeAAo1E8BOeAYfBlm4NG9Q==", "9FvVhtFLUs0KnA3Kprsdyg==", "IduElDUpDDXE677ZkhhKnQ==", "1tC/xrDYs8ey+sa3emtiYw==", "6ZmI6I2j5Y+R5aSn5ZOlAA==", "kPv59vyqzj00x11LXJZTjJ2UHW48jzHN", "ZmFsYWRvLnh5ei5zaGlybw==", "5AvVhmFLUS0ATA4Kprsdag==", "ZUdsaGJuSmxibVI2ZHc9PQ==", "6NfXkC7YVCV5DASIrEm1Rg==", "kPH+bIxk5D2deZiIxcaaaA==", "OUHYQzxQ/W9e/UjiAGu6rg==", "1AvVhdsgUs0FSA3SDFAdag==", "MTIzNDU2Nzg5MGFiY2RlZg==", "ZnJlc2h6Y24xMjM0NTY3OA==", "1QWLxg+NYmxraMoxAXu/Iw==", "vXP33AonIp9bFwGl7aT7rA==", "hBlzKg78ajaZuTE0VLzDDg==", "bWljcm9zAAAAAAAAAAAAAA==", "YTM0NZomIzI2OTsmIzM0NTueYQ==", "XgGkgqGqYrix9lI6vxcrRw==", "c2hpcm9fYmF0aXMzMgAAAA==", "bXRvbnMAAAAAAAAAAAAAAA==", "Jt3C93kMR9D5e8QzwfsiMw==", "i45FVt72K2kLgvFrJtoZRw==", "a2VlcE9uR29pbmdBbmRGaQ==", "6AvVhmFLUs0KTA3Kprsdag==", "6ZmI6I2j3Y+R1aSn5BOlAA==", "ZWvohmPdUsAWT3=KpPqda", "5AvVhmFLUs0KTA3Kprsdag==", "ZAvph3dsQs0FSL3SDFAdag==", "xVmmoltfpb8tTceuT5R7Bw==", "MTIzNDU2NzgxMjM0NTY3OA==", "NsZXjXVklWPZwOfkvk6kUA==", "bya2HkYo57u6fWh5theAWw==", "lxuEtAWbv+SgUOXREM+zrA==", "66v1O8keKNV3TTcGPK1wzg==", "2AvVhdsgUs0FSA3SDFAdag==", "SkZpbmFsQmxhZGUAAAAAAA==", "XTx6CKLo/SdSgub+OPHSrw==", "RVZBTk5JR0hUTFlfV0FPVQ=="}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func AesCBCDecrypt(encrypted []byte, key []byte) (err error, decrypted_data []byte){
	block, err := aes.NewCipher(key)   // 分组秘钥
	if err != nil {
		return err, []byte("")
	}                           
	blockSize := block.BlockSize() // 获取秘钥块的长度
	if len(encrypted) % blockSize != 0    {
		return errors.New("input not full blocks"), []byte("")
	}                       
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted_data = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted_data, encrypted)                 // 解密
	decrypted_data = pkcs5UnPadding(decrypted_data)  
	// fmt.Println(string(decrypted))                     // 去除补全码
	if len(decrypted_data) < 22 || bytes.Compare(decrypted_data[16:22], []byte{0xac, 0xed, 0x00, 0x05, 0x73, 0x72}) != 0 {//java 序列化开头 aced00057372
		return errors.New("maybe not valid java serilized data"), []byte("")
	}
	return nil, decrypted_data
}
func DecryptShiro(b64key, b64encrypt_str string) (string, error) {
	cipher_key, err := base64.StdEncoding.DecodeString(b64key)
	if err != nil {
		return "", err
	}
	cipher_data, err := base64.StdEncoding.DecodeString(b64encrypt_str)
	if err != nil {
		return "", err
	}

	err, decrypt_data := AesCBCDecrypt(cipher_data, cipher_key)
	if err == nil {
		iv_data_b64 := base64.StdEncoding.EncodeToString(decrypt_data[:16])
		decrypt_data_b64 := base64.StdEncoding.EncodeToString(decrypt_data[16:])
		ret_str := fmt.Sprintf("Key值: %s \n\nIV值: %s\n\n解密数据: %s \n\n解密数据（base64）: %s\n", b64key,  iv_data_b64, decrypt_data, decrypt_data_b64)
		return ret_str, nil
	}
	return "", err
}

func BatchDecryptShiro(encrypt_data string) (ret_str string){
	count := 0
	var result_str string
	for _, key := range key_list {
		ret_str, err := DecryptShiro(key, encrypt_data)
		if err == nil {
			count += 1
			result_str = ret_str
		}
	}
	if count == 0{
		return "很遗憾，没法解密"
	}else{
		return result_str
	}
}


var  htmltpl string = `<html>
	<head>
	<title>Shiro解密小工具</title>
	</head>
	<body>
	<div align="center">
		<form action="/" method="post">
			shiro rememberMe加密数据:<br>
			<textarea name="encrypt_text" rows="12" cols="100">{{ .Inputs }}</textarea>
			<br><br>
			<input type="submit" value="解密">
		</form>
		解密结果:<br>
		<textarea name="decrypt_text" rows="12" cols="100">{{ .Results }}</textarea>
	</div>
	</body>
	</html>
`
type ShiroInfo struct {
	Inputs		string
	Results    	string
}
// 渲染页面并输出
func renderHTML(w http.ResponseWriter, data interface{}) {
	t, err := template.New("test").Parse(htmltpl)
	if err != nil {
		log.Println("parse template err:", err)
	}
	t.Execute(w, data)
}
func decrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
        err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
        if err != nil {
          log.Fatal("ParseForm err: ", err)
		}

		si := &ShiroInfo{}
		_, b64_err := base64.StdEncoding.DecodeString(r.Form["encrypt_text"][0])
		if len(r.Form["encrypt_text"][0]) <= 400  {
			si.Inputs = r.Form["encrypt_text"][0]
			si.Results = "输入数据无效"
			renderHTML(w, si)
		}else if b64_err != nil {
			si.Inputs = r.Form["encrypt_text"][0]
			si.Results = "输入base64数据无效"
			renderHTML(w, si)
		}else{
			ret_str := BatchDecryptShiro(r.Form["encrypt_text"][0])
			si.Inputs = r.Form["encrypt_text"][0]
			si.Results = ret_str
			renderHTML(w, si)
		}
    }else {
		renderHTML(w, nil)
	}
}

func main() {
	log.Println("Will listening on 127.0.0.1:9090, please use browser to access it.")
	http.HandleFunc("/", decrypt)	// 设置访问的路由
	err := http.ListenAndServe("127.0.0.1:9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	
}
