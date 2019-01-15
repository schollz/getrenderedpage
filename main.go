package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var getpagejs = []byte(`const puppeteer = require('puppeteer');

(async() => {
	const browser = await puppeteer.launch({headless:true});
	const page = await browser.newPage();
	await page.goto(process.argv[2]);
    await page.waitFor(500);
	let content = await page.content();
	console.log(content);
	browser.close();
})();`)

func getRenderedPage(urlToGet string) (html []byte, err error) {
	if !strings.HasPrefix(urlToGet, "http") {
		err = fmt.Errorf("bad url")
		return
	}
	log.Println(urlToGet)
	tmpfile, err := ioutil.TempFile(".", "getpage.*.js")
	if err != nil {
		return
	}
	fname := tmpfile.Name()
	defer os.Remove(fname)

	if _, err = tmpfile.Write(getpagejs); err != nil {
		return
	}
	if err = tmpfile.Close(); err != nil {
		return
	}

	cmd := exec.Command("node", fname, urlToGet)
	html, err = cmd.CombinedOutput()
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	urlToGet := strings.Replace(r.URL.String()[1:], " ", "", -1)
	html, err := getRenderedPage(urlToGet)
	if err == nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(html)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
