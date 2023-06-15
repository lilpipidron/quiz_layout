package file_work

import (
	"gopkg.in/src-d/go-git.v4"
	examples "gopkg.in/src-d/go-git.v4/_examples"
)

func Update(system int) {
	var url string
	if system == 0 {
		url = "https://github.com/lipipidronstudy/updateRepo"
	} else {
		url = "https://github.com/lipipidronstudy/updateRepo"
	}

	_, err := git.PlainClone("/home/vadim/go", true, &git.CloneOptions{
		URL: url,
	})
	examples.CheckIfError(err)
	return
}
