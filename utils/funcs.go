package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

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
		cobra.CheckErr(errors.WithMessagef(err, "failed to get contents of file: %s\n", filename))
	} else {
		fmt.Println(string(content))
	}

}
