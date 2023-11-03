package oss

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	//r.LoadHTMLGlob("view/index.html")

	// index页面显示
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 表单提交
	r.POST("/uploadfile", func(c *gin.Context) {

		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 10010,
				"msg":  err.Error(),
			})
			return
		}

		// 上传到七牛云
		//code, url := UploadToQiNiu(f)
		url, err := UploadData(&KuboConfig{Bucket: "video-clip2", SecretKey: "VAEiWagIhTqM7VlTDBg-zRFM4UIFWhr5kJbAzBXR", AccessKey: "2H0cu5Oo9QRD7pSWXX-gCjDXjyaKU4Mw7y2XSx5q", DomainName: "s3eki5poq.hn-bkt.clouddn.com"}, f, "test")

		c.JSON(http.StatusOK, gin.H{
			//"code": url,
			"msg": "OK",
			"url": url,
		})

	})
	// 运行，监听127.0.0.1:8080
	r.Run()
}
