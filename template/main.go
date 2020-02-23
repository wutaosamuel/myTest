package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"log"
	"os"

	"./html"
	"./job"
	"./utils"
)

func main() {
	fmt.Println("test start")
	//testExec()
	var l *log.Logger
	logName := "test.log"
	WriteLog(l, logName, "log test")
	fn := func(g *log.Logger) {
		g.Println("test func 1")
		g.Println("test func 2")
	}
	WriteLogFunc(l, logName, fn)
	read, _ := ReadLog(logName)
	fmt.Println(read)
}

// WriteLog write into log
// Println only
func WriteLog(logger *log.Logger, logName, logInfo string) {
	// open log file
	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger = log.New(f, "", log.Ldate | log.Ltime | log.LUTC)
	logger.Println(logInfo)
	return
}

// LogActCallback is
type LogActCallback func(*log.Logger)

// WriteLogFunc write into log by func
func WriteLogFunc(logger *log.Logger, logName string, logFunc LogActCallback) {
	// open log file
	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger = log.New(f, "", log.Ldate|log.Ltime|log.LUTC)
	logFunc(logger)
	return
}

// ReadLog read log
func ReadLog(logName string) (string, error) {
	f, err := ioutil.ReadFile(logName)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

// testExec test exec
func testExec() {
	testExec := job.NewExec()
	testExec.Name = "testExec"
	logName := testExec.GetLogName()
	testExec.Command = "ls -l"
	testExec.SetCronTime("*/1", "*", "*", "*", "*")
	fmt.Println(testExec)
	fmt.Println("Cron Start")
	testExec.StartCron()
	//testExec.DoExec()
	time.Sleep(time.Duration(122) * time.Second)
	testExec.StopCron()
	log.Println("In Log done")
	fmt.Println("Cron Stop")

	f, err := ioutil.ReadFile(logName)
	if err != nil {
		fmt.Println("Read log error")
	}
	fmt.Println(string(f))

	//if err := testExec.DeleteLog(); err != nil {
	//fmt.Println("remove file fail")
	//}
	fmt.Println("done")
}

// testServer for server
func testServer() {
	fmt.Println("http start")
	// set static assert
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./html/CSS"))))
	// hand func
	http.HandleFunc("/", html.HandleIndex)
	http.HandleFunc("/shell.html", html.HandleShell)
	http.HandleFunc("/jobs.html", html.HandleJobs)
	http.HandleFunc("/logs.html", html.HandleLogs)
	fmt.Println("listening...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// testGen for html
func testGen() {
	fmt.Println("start to generate jobs.html")
	filenameJob := "./html/jobs.html"
	templateJob, _ := ioutil.ReadFile("./html/template/jobs.html")
	patternJob, _ := ioutil.ReadFile("./html/pattern/job_pattern1.html")
	j1 := html.Job{Name: "j1", ID: "j1-1234", Command: "run", Crontab: "crontab1", Stop: "button1"}
	j2 := html.Job{Name: "j2", ID: "j2-3456", Command: "run", Crontab: "crontab2", Stop: "button2"}
	jobs := []html.Job{j1, j2}
	h := html.GenerateJobs(jobs, string(templateJob), string(patternJob))
	utils.SaveHTML(filenameJob, h)

	filenameLog := "./html/logs.html"
	templateLog, _ := ioutil.ReadFile("./html/template/logs.html")
	patternLog, _ := ioutil.ReadFile("./html/pattern/log_pattern1.html")
	l1 := html.JobLog{Name: "j1", ID: "j1-1234", Command: "run", Crontab: "crontab1", Detail: "detail1", Delete: "delete"}
	ls := []html.JobLog{l1}
	lh := html.GenerateJobLogs(ls, string(templateLog), string(patternLog))
	utils.SaveHTML(filenameLog, lh)

	filenameD := "./html/detail.html"
	templateD, _ := ioutil.ReadFile("./html/template/detail.html")
	patternD, _ := ioutil.ReadFile("./html/pattern/detail_pattern1.html")
	d1 := html.Detail{Name: "j1", ID: "j1-1234", Command: "run", Crontab: "crontab1", Log: "Log here"}
	ds := []html.Detail{d1}
	dh := html.GenerateDetail(ds, string(templateD), string(patternD))
	utils.SaveHTML(filenameD, dh)
	fmt.Println("down")
}
