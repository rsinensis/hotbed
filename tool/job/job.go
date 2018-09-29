package job

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var jobs sync.Map

const ANY = -1
const JOB_START = 1
const JOB_STOP = 2

type Job struct {
	Name                 string
	Status               int
	Month, Day, Weekday  int8
	Hour, Minute, Second int8
	Task                 func(time.Time)
}

func init() {
	go processJobs()
}

func NewJob(name string, date string, status int, task func(time.Time)) error {

	dates := strings.Split(date, " ")
	if len(dates) != 6 {
		return fmt.Errorf("data parse fail")
	}

	var idates []int8
	for _, v := range dates {
		i, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("strconv fail:%v", err)
		}
		idates = append(idates, int8(i))
	}

	if idates[0] != -1 {
		if idates[0] > 12 || idates[0] < 1 {
			return fmt.Errorf("data month fail")
		}
	}

	if idates[1] != -1 {
		if idates[1] > 31 || idates[1] < 1 {
			return fmt.Errorf("data day fail")
		}
	}

	if idates[2] != -1 {
		if idates[2] > 7 || idates[2] < 0 {
			return fmt.Errorf("data weekday fail")
		}
	}

	if idates[3] != -1 {
		if idates[3] > 23 || idates[3] < 0 {
			return fmt.Errorf("data hour fail")
		}
	}

	if idates[4] != -1 {
		if idates[4] > 59 || idates[4] < 0 {
			return fmt.Errorf("data minute fail")
		}
	}

	if idates[5] != -1 {
		if idates[5] > 59 || idates[5] < 0 {
			return fmt.Errorf("data second fail")
		}
	}

	job := Job{Name: name, Status: status, Task: task, Month: idates[0], Day: idates[1], Weekday: idates[2], Hour: idates[3], Minute: idates[4], Second: idates[5]}

	// _, ok := jobs.Load(name)
	// if ok {
	// 	return fmt.Errorf("job exist")
	// }

	jobs.Store(name, job)

	return nil

}

func ChangeJobStatus(name string, status int) bool {

	vv, ok := jobs.Load(name)
	if !ok {
		return false
	}

	j, ok := vv.(Job)
	if !ok {
		return false
	}

	j.Status = status

	jobs.Store(name, j)

	return true

}

func DeleteJob(name string) bool {

	vv, ok := jobs.Load(name)
	if !ok {
		return false
	}

	_, ok = vv.(Job)
	if !ok {
		return false
	}

	jobs.Delete(name)

	return true

}

func (cj Job) Matches(t time.Time) (ok bool) {
	ok = (cj.Month == ANY || cj.Month == int8(t.Month())) &&
		(cj.Day == ANY || cj.Day == int8(t.Day())) &&
		(cj.Weekday == ANY || cj.Weekday == int8(t.Weekday())) &&
		(cj.Hour == ANY || cj.Hour == int8(t.Hour())) &&
		(cj.Minute == ANY || cj.Minute == int8(t.Minute())) &&
		(cj.Second == ANY || cj.Second == int8(t.Second()))

	return ok
}

func processJobs() {
	for {

		now := time.Now()

		jobs.Range(func(k, v interface{}) bool {

			j, ok := v.(Job)
			if !ok {
				return false
			}

			if j.Matches(now) && j.Status == JOB_START {
				go j.Task(now)
			}

			return true
		})

		time.Sleep(time.Second)
	}
}
