package main

import (
	"errors"
	"os/exec"
	"runtime"
)

var openFuncs = map[string]func(string) error{
	"windows": openWebByWindows,
	"linux":   openWebByLinux,
	"wsl":     openWebByWSL,
	"darwin":  openWebByMac,
}

func openWeb(url string, isWSL bool) error {
	runtimeOS := runtime.GOOS
	if isWSL {
		runtimeOS = "wsl"
	}

	openFunc, ok := openFuncs[runtimeOS]
	if !ok {
		return errors.New("That OS is not yet supported....")
	}
	return openFunc(url)
}

func openWebByWindows(url string) error {
	_, err := exec.Command("cmd.exe", "/c", "start", url).CombinedOutput()
	return err
}

func openWebByLinux(url string) error {
	// TODO: インタラクティブモードで開かれた場合どうするか。errはnil
	//       戻り値に返ってきてるっぽい
	_, err := exec.Command("xdg-open", url).CombinedOutput()
	return err
}

func openWebByWSL(url string) error {
	_, err := exec.Command("cmd.exe", "/c", "start", url).CombinedOutput()
	return err
}

func openWebByMac(url string) error {
	_, err := exec.Command("open", url).CombinedOutput()
	return err
}
