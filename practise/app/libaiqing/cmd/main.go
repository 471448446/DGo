package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type SimpleDownload func()
/**
下载李伯清老师的评书
 */
func main() {
	println("start")
	//findCatalog("大话60年", "http://pingshu.zgpingshu.com/pingshu/5108/#play")
	findCatalog("舌尖上的四川", "http://pingshu.zgpingshu.com/pingshu/5109/")
	//findDownloadLink("下载 01", "http://www.zgpingshu.com/down/5108/")
	//download("下载 01", "http://dops81.zgpingshu.com/%E6%9D%8E%E4%BC%AF%E6%B8%85/%E5%A4%A7%E8%AF%9D60%E5%B9%B4%2848%E5%9B%9E%29%28%E6%BC%94%E6%92%AD-%E6%9D%8E%E4%BC%AF%E6%B8%85%29%2832kb%29/797F24E555.mp3")
}

func findCatalog(name string, url string) {
	e := os.Mkdir(name, os.ModePerm)
	//https://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-exists
	if nil != e && !os.IsExist(e) {
		fmt.Println("fail create dir just return all", e)
		return
	}
	var wg sync.WaitGroup

	fmt.Println("Hello")
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		link := e.Attr("href")
		var linksCatalog = make(map[string]string)
		if ConvertString(e.Text) == "下载" {
			name := ConvertString(e.Attr("title"))
			//name = strings.Replace(name, ConvertString("下载 "), "", 1)

			//  //www.zgpingshu.com/down/5108/
			if strings.HasPrefix(link, "//") {
				link = "http:" + link
			}
			fmt.Printf("——————: %q -> %s\n", name, link)
			// 所有的目录
			linksCatalog[name] = link
		}
		time.Sleep(1000)

		if len(linksCatalog) != 0 {
			index := 0
			for i, v := range linksCatalog {
				wg.Add(1)
				go func() {
					findDownloadLink(i, v, name, func() {
						wg.Done()
					})
				}()

				time.Sleep(10000)
				if 0 == index {
					break
				}
				index++
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	ee := c.Visit(url)
	if nil != ee {
		fmt.Println("error_", ee)
	}
	//https://mozillazg.com/2014/10/go-wait-all-goroutines-end.html
	wg.Wait()
}

func findDownloadLink(name string, link string, dir string, call SimpleDownload) {
	//	——————: "下载 01回" -> http://www.zgpingshu.com/down/5108/
	fmt.Printf("findDownloadLink:%q -> %s\n", name, link)

	c := colly.NewCollector()
	c.OnHTML("#down", func(element *colly.HTMLElement) {
		downloadLink := element.Attr("href")
		fmt.Println("findDownloadLink suc:", name, downloadLink)
		download(name, downloadLink, dir, call)
	})
	c.OnError(func(response *colly.Response, e error) {
		fmt.Println("findDownloadLink fail:", name)
		call()
	})
	e := c.Visit(link)
	fmt.Println("findDownloadLink visit:", name, link, "ok?", e)
}

func download(name string, url string, dir string, call SimpleDownload) {
	// http://dops81.zgpingshu.com/%E6%9D%8E%E4%BC%AF%E6%B8%85/%E5%A4%A7%E8%AF%9D60%E5%B9%B4%2848%E5%9B%9E%29%28%E6%BC%94%E6%92%AD-%E6%9D%8E%E4%BC%AF%E6%B8%85%29%2832kb%29/797F24E555.mp3

	defer call()

	fmt.Println("download____", "prepare", name)
	res, e := http.Get(url)
	if nil != e {
		fmt.Println("download____", "err get", name, e.Error())
		return
	}
	file, e := os.Create(dir + "/" + name + ".mp3")
	if nil != e {
		fmt.Println("download____", "error file ", name, e.Error())
		return
	}
	defer file.Close()
	written, e := io.Copy(file, res.Body)
	if nil != e {
		fmt.Println("download____", "error save failed", name, written, e.Error())
	} else {
		fmt.Println("download____", "save ok ", name, written)
	}
}

func ConvertString(src string) string {
	return ConvertToString(src, "gbk", "utf-8")
}

/**
https://www.kancloud.cn/liupengjie/go/999877
*/
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
