// Copyright 2017 Vckai Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snatchs

import (
	"strings"
	"testing"
)

func TestBQIsBookUrl(t *testing.T) {
	bq := NewBiquge()

	url := "http://www.biqiuge.com/book/4772/"
	if !bq.IsBookURL(url) {
		t.Fatal("小说简介页面匹配失败")
	}

	url = "http://www.biqiuge.com/book/4772/222.html"
	if bq.IsBookURL(url) {
		t.Fatal("小说简介页面匹配失败")
	}

	url = "http://www.biqiuge.com/xuanhuanxiaoshuo/"
	if bq.IsBookURL(url) {
		t.Fatal("小说简介页面匹配失败")
	}
}

func TestBQIsCrawlerUrl(t *testing.T) {
	bq := NewBiquge()

	url := "http://www.biqiuge.com/book/4772/"
	if bq.IsCrawlerURL(url) {
		t.Fatal("爬虫URL匹配失败")
	}

	url = "http://www.biqiuge.com/book/4772/222.html"
	if bq.IsCrawlerURL(url) {
		t.Fatal("爬虫匹配失败")
	}

	url = "http://www.biqiuge.com/xuanhuanxiaoshuo/"
	if !bq.IsCrawlerURL(url) {
		t.Fatal("爬虫匹配失败")
	}
}

func TestBQGetNovel(t *testing.T) {
	bq := NewBiquge()

	url := "http://www.biqiuge.com/book/4772/"
	info, err := bq.GetNovel(url)
	if err != nil {
		t.Fatal(err.Error())
	}

	nov := info.Nov

	if nov == nil {
		t.Fatal("获取小说失败")
	}

	if nov.Name != "圣墟" {
		t.Fatal("获取小说标题错误", nov.Name)
	}

	if !strings.Contains(nov.Desc, "败中崛起，在寂灭中复苏") {
		t.Fatal("获取小说内容错误", nov.Desc)
	}

	if nov.Author != "辰东" {
		t.Fatal("获取小说作者错误", nov.Author)
	}

	if nov.CateName != "玄幻小说" {
		t.Fatal("获取小说分类名称错误", nov.CateName)
	}

	if nov.CateId != 1 {
		t.Fatal("匹配小说分类ID错误", nov.CateId)
	}

	if nov.Cover != "http://www.biqiuge.com/files/article/image/4/4772/4772s.jpg" {
		t.Fatal("获取小说章节封面错误", nov.Cover)
	}

	if info.ChapterUrl != url {
		t.Fatal("获取小说章节链接错误", info.ChapterUrl)
	}
}

func TestBQFindNovel(t *testing.T) {
	bq := NewBiquge()
	info, err := bq.FindNovel("圣墟")
	if err != nil {
		t.Fatal(err.Error())
	}

	nov := info.Nov

	if nov == nil {
		t.Fatal("获取小说失败")
	}

	if nov.Name != "圣墟" {
		t.Fatal("获取小说标题错误")
	}

	if !strings.Contains(nov.Desc, "在寂灭中复苏") {
		t.Fatal("获取小说内容错误")
	}

	if nov.Author != "辰东" {
		t.Fatal("获取小说作者错误")
	}

	if info.Url != "http://www.biqiuge.com/book/4772/" {
		t.Fatal("获取小说链接错误")
	}

	if info.ChapterUrl != "http://www.biqiuge.com/book/4772/" {
		t.Fatal("获取小说章节链接错误")
	}
}

func TestBQGetChapter(t *testing.T) {
	bq := NewBiquge()

	info, err := bq.GetChapter("http://www.biqiuge.com/book/4772/2940354.html")
	if err != nil {
		t.Fatal(err.Error())
	}
	chap := info.Chap

	if chap.Title != "第一章 沙漠中的彼岸花" {
		t.Fatal("获取章节标题错误", chap.Title)
	}

	if !strings.Contains(chap.Desc, "的大漠，空旷而高远，壮阔") {
		t.Fatal("获取章节内容错误", chap.Desc)
	}

	if info.PreUrl != "" && !strings.Contains(info.PreUrl, "2940353.html") {
		t.Fatal("获取上一页连接错误", info.PreUrl)
	}

	if info.NextUrl != "http://www.biqiuge.com/book/4772/2941694.html" {
		t.Fatal("获取下一页连接错误", info.NextUrl)
	}

}

func TestBQGetChapters(t *testing.T) {
	bq := NewBiquge()

	links, err := bq.GetChapters("http://www.biqiuge.com/book/4772/")
	if err != nil {
		t.Fatal("获取章节列表失败", err.Error())
	}
	if len(links) == 0 {
		t.Fatal("获取章节列表失败")
	}

	if links[0].Chap.Link != "http://www.biqiuge.com/book/4772/2940354.html" {
		t.Fatal("获取小说链接错误", links[0].Chap.Link)
	}
}
