package util

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Pager struct{
	Page int
	Totalnum int
	Pagesize int 
	urlpath string
	urlquery string
	nopath bool
}

func NewPager(page,totalnum,pagesize int,url string,nopath ...bool) *Pager  {
	p := new(Pager)
	p.Page = page
	p.Totalnum = totalnum
	p.Pagesize = pagesize
	arr := strings.Split(url,"?")
	p.urlpath = arr[0]
	if len(arr) > 1{
		p.urlquery = "?"+ arr[1]
	}else{
		p.urlquery = ""
	}
	if len(nopath)>0{
		p.nopath = nopath[0]
	}else{
		p.nopath = false
	}
	return p
}

func (p *Pager) url(page int) string  {
	if p.nopath{//不使用目录形式
		if p.urlquery !=""{
			return fmt.Sprintf("%s%s%[age=%d",p.urlpath,p.urlquery,page)
		}else{
			return fmt.Sprintf("%s?page=%d",p.urlpath,page)
		}
	}else{
		return fmt.Sprintf("%s/page/%d%s",p.urlpath,page,p.urlquery)
	}
}

func (p *Pager) ToString() string {
	if p.Totalnum<= p.Pagesize{
		return ""
	}
	var buf bytes.Buffer
	var from,to,linknum,offset,totalpage int
	offset = 5
	linknum = 15
	totalpage = int(math.Ceil(float64(p.Totalnum) / float64(p.Pagesize)))
	
	if totalpage < linknum {
		from = 1
		to = totalpage
	}else{
		from = p.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum -1 
		}else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}
	if p.Page > 1 {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-prev\" href=\"%s\">上一页</a></li>", p.url(p.Page-1)))
	} else {
		buf.WriteString("<span>上一页</span>")
	}

	if p.Page > linknum {
		buf.WriteString(fmt.Sprintf("<a href=\"%s\" class=\"laypage_first\">1...</a>", p.url(1)))
	}

	for i := from; i <= to; i++ {
		if i == p.Page {
			buf.WriteString(fmt.Sprintf("<span class=\"layui-laypage-curr\"><em class=\"layui-laypage-em\"></em><em>%d</em></span>", i))
		} else {
			buf.WriteString(fmt.Sprintf("<a href=\"%s\">%d</a>", p.url(i), i))
		}
	}

	if totalpage > to {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-last\" href=\"%s\">末页</a>", p.url(totalpage)))
	}

	if p.Page < totalpage {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-next\" href=\"%s\">下一页</a></li>", p.url(p.Page+1)))
	} else {
		buf.WriteString(fmt.Sprintf("<span>下一页</span>"))
	}
	return buf.String()
}