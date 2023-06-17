package main

import (
	"errors"
	"os/exec"
	"runtime"
)

var openFuncs = map[string]func(string) error{
	"windows": openWebByWindows,
	"linux":   openWebByLinux,
	"darwin":  openWebByMac,
}

func openWeb(url string) error {
	openFunc, ok := openFuncs[runtime.GOOS]
	if !ok {
		return errors.New("That OS is not yet supported....")
	}
	return openFunc(url)
}

func openWebByWindows(url string) error {
	_, err := exec.Command("start", url).CombinedOutput()
	return err
}

func openWebByLinux(url string) error {
	// TODO: WSLどうしよう。これで実行すると、ブラウザ起動されないけどerrはnilで、だからurlの出力もしない
	//       インタラクティブモードで開かれてるようで、多分戻り値に返ってきてる
	_, err := exec.Command("xdg-open", url).CombinedOutput()
	// TODO: WSLなら以下で開ける
	// _, err := exec.Command("cmd.exe", "/c", "start", url).CombinedOutput()
	return err
}

func openWebByMac(url string) error {
	_, err := exec.Command("open", url).CombinedOutput()
	return err
}
