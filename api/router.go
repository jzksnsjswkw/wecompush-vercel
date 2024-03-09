package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jzksnsjswkw/wecom-push"
)

var router *gin.Engine

func init() {
	r := gin.Default()

	r.POST("", func(ctx *gin.Context) {
		s := struct {
			Corpid     string `json:"corpid" binding:"required"`
			Corpsecret string `json:"corpsecret" binding:"required"`
			Touser     string `json:"touser"`
			AgentID    int    `json:"agentID" binding:"required"`
			Content    string `json:"content" binding:"required"`
		}{}
		if err := ctx.BindJSON(&s); err != nil {
			return
		}
		w := wecom.New(s.Corpid, s.Corpsecret)
		err := w.Text(&wecom.TextInfo{
			Touser:  s.Touser,
			AgentID: s.AgentID,
			Content: s.Content,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		}
		ctx.Status(http.StatusOK)
	})
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
