package command

import (
	"fmt"
	"os/exec"
)

func GetSiteIp(name string) string {
	prg := "dig"

	arg1 := "+short"
	arg2 := name
	arg3 := "@resolver1.opendns.com"

	cmd := exec.Command(prg, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())

		return err.Error()
	}

	return string(stdout)
}
