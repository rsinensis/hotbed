package task

import (
	"hotbed/tool/job"
)

func TaskInit() {
	job.NewJob("upload", "* * 3 * * *", job.JOB_START, uploadTask)
}
