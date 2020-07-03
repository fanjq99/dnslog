package api

import (
	"net/http"

	"github.com/fanjq99/common/db"
	"github.com/fanjq99/common/log"
	"github.com/fanjq99/dnslog/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type HttpServer struct {
	server      *http.Server
	redisClient *redis.Client
	addr        string
	saveTime    int64
}

func NewHttpServer(c config.YmlConfig) *HttpServer {
	client, err := db.GetRedisClient(c.Redis.Addr,c.Redis.Password,c.Redis.Database)
	if err != nil {
		log.Fatal(err)
	}
	hs := &HttpServer{
		addr:        c.ApiAddr,
		redisClient: client,
		saveTime:    c.SaveTime,
	}
	server := &http.Server{
		Addr:    c.ApiAddr,
		Handler: hs.buildHandler(),
	}

	hs.server = server
	return hs
}

func (h *HttpServer) Run() {
	log.Info("api Listening on ", h.addr)
	if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen:%s", err)
	}
}

func (h *HttpServer) buildHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/")

	api.GET("/", h.Index)
	api.GET("/status", h.Status)
	api.GET("/ddffdd", h.PHPInclude)
	api.GET("/xss", h.Xss)
	api.GET("/vul-verify.php", h.VulVerifyHttp)
	api.GET("/dns-verify.php", h.VulVerifyDNS)

	return router
}
