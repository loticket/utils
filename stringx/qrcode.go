package stringx

import (
	"bytes"
	"encoding/json"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/loticket/utils/encry"
	"image/png"
)

//RSA加密生成二维码
func CreateQrCode(data string, pubKey string) ([]byte, error) {
	result, err := encry.RsaPublicEncrypt(data, pubKey)
	if err != nil {
		return nil, err
	}

	qrcode, err := qr.Encode(result, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	qrcode, _ = barcode.Scale(qrcode, 200, 200)

	encoder := png.Encoder{CompressionLevel: png.BestCompression}

	var b bytes.Buffer
	err = encoder.Encode(&b, qrcode)

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

//@action解析二维码
//@param 二维码扫描后获取的字符串  codeInfo 需要json反序列化  priKey 加密（rsa)的私钥
//@return error
func PaserQrCodeAction(param string, codeInfo interface{}, priKey string) error {
	result, err := ParseQrCodeToStr(param, priKey)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), codeInfo)
	if err != nil {
		return err
	}

	return nil
}

//@action 解密二维码字符串
//@param 二维码扫描后获取的字符串 priKey 加密（rsa)的私钥
//@return string json字符串 error 错误信息
func ParseQrCodeToStr(param string, priKey string) (string, error) {
	result, err := encry.RsaPriKeyDecrypt(param, priKey)
	if err != nil {
		return "", err
	}
	return result, nil
}

//@action创建多功能二维码 -- 先使用rsa 公钥加密后，生成二维码
//@parame data 需要加密的结构体 action 动作 pubkey 加密（rsa)的公钥
//return []byte 生成图片的字节 直接json处理返回即可 error 错误信息
func CreateQrCodeAction(data interface{}, action string, pubKey string) ([]byte, error) {
	var (
		pngs []byte
		err  error
	)

	if pngs, err = json.Marshal(&data); err != nil {
		return nil, err
	}

	result, err := encry.RsaPublicEncrypt(string(pngs), pubKey)
	if err != nil {
		return nil, err
	}

	var actionQrcode map[string]string = map[string]string{
		"action": action,
		"qrcode": result,
	}

	aQrcode, errs := json.Marshal(actionQrcode)
	if err != nil {
		return nil, errs
	}

	qrcode, err := qr.Encode(string(aQrcode), qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	qrcode, _ = barcode.Scale(qrcode, 200, 200)

	encoder := png.Encoder{CompressionLevel: png.BestCompression}

	var b bytes.Buffer
	err = encoder.Encode(&b, qrcode)

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
