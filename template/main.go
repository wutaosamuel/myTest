package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"./utils"
	"./html"
)

func main() {
	fmt.Println("http start")
	testGen()
	// set static assert
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./html/CSS"))))
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
	j1 := html.Job{Name:"j1", ID:"j1-1234", Command:"run", Crontab:"crontab1", Stop:"button1"}
	j2 := html.Job{Name:"j2", ID:"j2-3456", Command:"run", Crontab:"crontab2", Stop:"button2"}
	jobs := []html.Job{j1, j2}
	h := html.GenerateJobs(jobs, string(templateJob), string(patternJob))
	utils.SaveHTML(filenameJob, h)

	filenameLog := "./html/logs.html"
	templateLog, _ := ioutil.ReadFile("./html/template/logs.html")
	patternLog, _ := ioutil.ReadFile("./html/pattern/log_pattern1.html")
	l1 := html.JobLog{Name:"j1", ID:"j1-1234", Command:"run", Crontab:"crontab1", Detail:"detail1", Delete:"delete"}
	ls := []html.JobLog{l1}
	lh := html.GenerateJobLogs(ls, string(templateLog), string(patternLog))
	utils.SaveHTML(filenameLog, lh)

	filenameD := "./html/detail.html"
	templateD, _ := ioutil.ReadFile("./html/template/detail.html")
	patternD, _ := ioutil.ReadFile("./html/pattern/detail_pattern1.html")
	d1 := html.Detail{Name:"j1", ID:"j1-1234", Command:"run", Crontab:"crontab1", Log:"Log here"}
	ds := []html.Detail{d1}
	dh := html.GenerateDetail(ds, string(templateD), string(patternD))
	utils.SaveHTML(filenameD, dh)
	fmt.Println("down")
}