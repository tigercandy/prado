package utils

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverGigit   *base64Captcha.DriverDigit
}

func DriverStringFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4, "123456789abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, nil, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

func DriverGigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverGigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	driver := e.DriverGigit
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}
