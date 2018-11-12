package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/alexflint/go-arg"
)

var args struct {
	UserName string `arg:"positional,required"`
	Realm    string `arg:"positional,required"`
	Password string `arg:"positional,required"`
}

func main() {
	arg.MustParse(&args)
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s:%s:%s", args.UserName, args.Realm, args.Password)))
	fmt.Printf("%s:%s:%s", args.UserName, args.Realm, hex.EncodeToString(h.Sum(nil)))
}
