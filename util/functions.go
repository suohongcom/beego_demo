package util

import (
	"crypto/md5"
	"fmt"
	"strings"
	"net/url"
	"encoding/base64"
	"crypto/rand" 
	"io"
) 


//加密
func Md5(str string) string  {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x",hash.Sum(nil))
}

//路由编译
func Rawurlencode(str string) string  {
	return strings.Replace(url.QueryEscape(str),"+","%20",-1)
}

func UniqueId() string {
	b := make([]byte, 48)
	if _,err := io.ReadFull(rand.Reader,b);err != nil{
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}
