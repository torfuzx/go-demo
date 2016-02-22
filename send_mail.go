package main

import "yto.net.cn/kefu/models"

var (
	Mo *models.Models
)

func init() {
	Mo = new(models.Models)

	if err := Mo.InitializeSmtp2("zaixiankefu@yto.net.cn", "zaixiankefu",
		"mail.yto.net.cn", 25); err != nil {
		panic(err)
	}

}

func main() {
	Mo.Mail.FromSet("zaixiankefu@yto.net.cn", "张学友")
	addr1 := Mo.Mail.AddressFormat("torfu@qq.com", "刘德华")
	addr2 := Mo.Mail.AddressFormat("chowmein@qq.com", "黎明")
	Mo.Mail.ToSet(addr1, addr2)
	Mo.Mail.CcSet("torfu@foxmail.com", "郭富城")
	Mo.Mail.SubjectSet("我来自Hong Kong")
	Mo.Mail.HtmlBodySet("hello <b>刘德华</b> and <i>黎明</i>!")

	// Mo.Mail.PlainBodySet("hello  刘德华 and 黎明")
	// Mo.Mail.HtmlAlternativeAdd("hello <b>刘德华</b> and <i>黎明</i>!")

	file, err := Mo.Mail.FileOpen("./image.jpg")
	if err != nil {
		panic(err)
	}

	// file := Mo.Mail.FileCreate("image1.jpg", []byte("Content 1"))

	Mo.Mail.Attach(file)

	// Mo.Mail.Embed(file)
	// Mo.Mail.BodySet("text/html", `<img src="cid:image.jpg" />`)

	if err := Mo.Mail.Send(); err != nil {
		panic(err)
	}
}
