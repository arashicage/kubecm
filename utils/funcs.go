package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

// CheckErr 对于命令行类的应用，检测到出错时，打印错误并退出就行了
func CheckErr(err error) {
	if err != nil {
		color.Red(err.Error())
		println()
	}
}

func PrettifyPrint(v interface{}) {

	fmt.Println(awsutil.Prettify(v))
}

func FileExists(name string) bool {

	info, err := os.Lstat(name)
	if err == nil {
		return !info.IsDir()
	}
	return !os.IsNotExist(err)
}

func Cat(filename string) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		CheckErr(errors.WithMessagef(err, "failed to get contents of file: %s\n", filename))
	} else {
		fmt.Println(string(content))
	}

}
