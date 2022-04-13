package utils

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/pkg/errors"
)

// CheckErr 工具类应用中检测到出错时，简单退出打印并退出即可
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrettifyPrint(v interface{}) {

	fmt.Println(awsutil.Prettify(v))
}

func Cat(filename string) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		CheckErr(errors.WithMessagef(err, "failed to get contents of file: %s\n", filename))
	} else {
		fmt.Println(string(content))
	}

}
