package tasks

import (
	"hotbed/tools/job"
)

func TaskInit() {
	job.NewJob("upload", "* * 3 * * *", job.JOB_START, uploadTask)
}
