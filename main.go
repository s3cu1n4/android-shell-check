package main

import (
	"android-shell-check/common"
	"archive/zip"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/s3cu1n4/logs"
)

type myCloser interface {
	Close() error
}

func closeFile(f myCloser) {
	err := f.Close()
	if err != nil {
		logs.Error(err.Error())
	}
}

func main() {
	var zipFile string
	flag.StringVar(&zipFile, "f", "", "apk文件名")

	flag.Parse()
	if zipFile == "" {
		logs.Info("请输入api的文件路径或文件名")
		os.Exit(1)
	}

	var soinfo string

	zf, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer closeFile(zf)

	for _, file := range zf.File {
		info := file.FileInfo()

		soname := info.Name()
		match, _ := regexp.MatchString("^lib.*.so$", file.Name)
		if match {
			temp := common.GetLibInfo(soname)
			if temp != "" {
				soinfo = temp
				logs.Info(file.Name, soinfo)
			}
		}
	}
	if soinfo == "" {
		logs.Info("未知加固厂商或未加壳")
	}
}
