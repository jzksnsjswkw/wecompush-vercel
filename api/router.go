package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jzksnsjswkw/wecom-push"
)

var router *gin.Engine

func init() {
	r := gin.Default()

	r.POST("/:touser/:agentID", func(ctx *gin.Context) {
		s := struct {
			// AppID      string `json:"appID"`
			// AppName string `json:"appName"`
			// DeviceName string `json:"deviceName"`

			Title    string `json:"title"`
			// Subtitle string `json:"subtitle"`
			Message  string `json:"message"`

			// Date  string `json:"date"`
			// Image string `json:"image"`
			// Icon string `json:"icon"`
		}{}
		if err := ctx.ShouldBindJSON(&s); err != nil {
			ctx.String(http.StatusBadRequest, "invalid body")
			return
		}
		s2 := struct {
			Touser  string `uri:"touser"`
			AgentID int    `uri:"agentID"`
		}{}
		if err := ctx.ShouldBindUri(&s2); err != nil {
			ctx.String(http.StatusBadRequest, "invalid uri")
			return
		}

		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.String(http.StatusBadRequest, "authorization is empty")
			return
		}
		ss := strings.Split(auth, ":")
		if len(ss) != 2 {
			ctx.String(http.StatusBadRequest, "invalid Authorization")
			return
		}

		w := wecom.New(ss[0], ss[1])
		err := w.Text(&wecom.TextInfo{
			Touser:  s2.Touser,
			AgentID: s2.AgentID,
			Content: fmt.Sprintf("%s\n%s", s.Title, s.Message),
		})
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		ctx.String(http.StatusOK, "OK")
	})
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
