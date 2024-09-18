package tasks

import "main.go/app/core/tasks/syncronize"

func Run() {
	//run := decider.RunDecider()
	//
	//if !run {
	//	return
	//} else {
	//	download.Download()
	//	unzip.Unzip()
	//	syncronize.Syncronize()
	//}

	syncronize.Syncronize()
}
