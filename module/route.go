package module

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"server.pluto.com/log"
	"server.pluto.com/module/service"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func RegisterRouter(r *gin.Engine) *gin.Engine {
	log.Info("访问主页")
	r.GET("/", Index)

	log.Info("get接口测试")
	r.GET("/get-test", GetTest)

	log.Info("post接口测试")
	r.POST("/post-test", PostTest)

	log.Info("获取服务器运行资源信息")
	r.GET("/resources", GetServerResource)

	log.Info("查询Minecraft用户")
	r.GET("/query", QueryUser)

	log.Info("新增Minecraft用户")
	r.GET("/insert", InsertUser)

	log.Info("删除Minecraft用户")
	r.GET("/del", DeltUser)
	return r
}
func DeltUser(c *gin.Context) {
	rc := c.Query("name")
	log.Info("删除Minecraft用户，username=", rc)
	process := service.User{}
	c.JSON(0, process.DelUser(rc))
}

func InsertUser(c *gin.Context) {
	rc := c.Query("name")
	log.Info("新增Minecraft用户，username=", rc)
	process := service.User{}
	c.JSON(0, process.InsertUser(rc))
}

func QueryUser(c *gin.Context) {
	log.Info("查询Minecraft用户")
	result := service.User{}
	c.JSON(0, result.QueryUserList())
}

func PostTest(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	fmt.Println("姓名name=" + name + ",年龄age=" + age)
}

func Index(c *gin.Context) {
	result := c.Param("name")
	fmt.Println(result)
	c.JSON(0, "Hello Pluto")
}

func GetTest(c *gin.Context) {
	rc := c.Query("name")
	fmt.Println(rc)
	c.JSON(0, "name的值为"+rc)
}

func GetServerResource(c *gin.Context) {
	log.Info("获取服务器运行资源信息")
	receive := service.SeverRecsource{}
	c.JSON(0, receive.GetServerResources())
}
