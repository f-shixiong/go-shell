package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var installPath = os.Getenv("GOPATH") + "/src/github.com/f-shixiong/go-shell/.plugin-lib"
var defaultPluginLib = os.Getenv("GOPATH") + "/src/github.com/f-shixiong/go-shell/lib/plugin-lib"

func autoImport(path string) error {
	codeCachePath := os.Getenv("GOPATH") + "/src/github.com/f-shixiong/go-shell/.code-cache-lib/" + path
	Debug("codeCachePath -> %s", codeCachePath)
	Debug("installPath -> %s", installPath)
	createFile(installPath)
	createFile(codeCachePath)
	srcs := []string{os.Getenv("GOROOT"), os.Getenv("GOPATH")}
	for _, src := range srcs {
		if src == "" {
			continue
		}
		srcPath := src + "/src/" + path
		vendorPath := src + "/src/vendor"
		if isExist(srcPath) {
			Debug("begin copy")
			files, _ := ioutil.ReadDir(srcPath)
			packageName := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
			for _, f := range files {
				if f.IsDir() {
					continue
				}
				narr := strings.Split(f.Name(), ".")
				if len(narr) != 2 || strings.Contains(narr[0], "_test") || narr[1] != "go" {
					continue
				}
				Debug("f -> %#v,package -> %s,mode -> %s", f, packageName, f.Mode().String())
				newFile := codeCachePath + "/" + f.Name()
				oldFile := srcPath + "/" + f.Name()
				Debug("new -> %s, old -> %s", newFile, oldFile)
				_, err := CopyFile(newFile, oldFile)
				if err != nil {
					panic(err)
				}
				buf, err := ioutil.ReadFile(newFile)
				if err != nil {
					panic(err)
				}
				content := string(buf)
				reg, _ := regexp.Compile("package[ ]+" + packageName + "[ ]*")
				content = reg.ReplaceAllString(content, "package main")
				err = ioutil.WriteFile(newFile, []byte(content), 0)
				if err != nil {
					panic(err)
				}
			}
			//TODO 20
			pluginName := strings.Replace(path, "/", "@@@@@@", 20) + ".so"
			installPlugin(pluginName, codeCachePath, installPath, vendorPath)
			break
		}
	}
	Debug("goroot = %s ,gopath = %s", os.Getenv("GOROOT"), os.Getenv("GOPATH"))

	return nil
}

func installPlugin(pName, cPath, iPath, vPath string) {
	Debug("pluginName -> %s, codeCachePath-> %s,installPath-> %s", pName, cPath, iPath)
	copyVendorCmd := fmt.Sprintf("cp -r %s %s", vPath, cPath+"/vendor")

	buildCmd := fmt.Sprintf("cd %s && go build -buildmode=plugin -o %s && mv %s %s", cPath, pName, pName, iPath+"/"+pName)
	Debug("copeC = %s", copyVendorCmd)
	Debug("buildC = %s", buildCmd)
	os.Setenv("GOPATH", os.Getenv("GOPATH")+":"+vPath)
	cc := exec.Command("cp", "-r", vPath, cPath+"/vendor")
	err := cc.Run()
	if err != nil {
		Error("error %+v", err)
	}
	hc := exec.Command("go", "build", "-buildmode=plugin", "-o", pName)
	hc.Dir = cPath + "/"
	err = hc.Run()
	if err != nil {
		Error("error %+v", err)
	}
	mc := exec.Command("mv", pName, iPath+"/"+pName)
	mc.Dir = cPath + "/"
	err = mc.Run()
	if err != nil {
		Error("error %+v", err)
	}
}
func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	/*if isExist(dstName) {
		os.Remove(dstName)
	}*/
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
