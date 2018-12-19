package handlers

import (
	"com.tony666.answer/activity"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type userOption struct {
	UId    int             `json:"u_id"`
	Option activity.Option `json:"option"`
}

var upgrader = websocket.Upgrader{} // use default options

func DoAnswer(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		_, _ = w.Write([]byte("错误的协议, 仅支持websocket协议."))
		return
	}
	defer c.Close()
	incomeUid, _ := strconv.Atoi(r.FormValue("uid"))
	activity.RegisterChan(incomeUid)
	go writeMessage(incomeUid, c)
	readMessage(incomeUid, c)
}

func readMessage(uid int, c *websocket.Conn) {
	for {
		o := userOption{}
		err := c.ReadJSON(&o)
		if err != nil {
			activity.ReleaseChan(uid)
			log.Println("客户端连接已关闭...")
			return
		}
		log.Println("接受到客户端消息: ", o)
		if activity.CrtActivity.Processing {
			if activity.CrtActivity.CardState.CanReply {
				activity.CrtActivity.Answer(o.UId, o.Option)
			} else {
				_ = c.WriteMessage(1, []byte("脑子不好使啦？？现在已经不能继续答题啦！!"))
			}
		} else {
			_ = c.WriteMessage(1, []byte("本次活动已经结束，感谢参与!"))
		}
	}
}

func writeMessage(uId int, c *websocket.Conn) {
	for {
		select {
		case question, opened := <-activity.QuestionChan:
			if opened {
				err := c.WriteJSON(question)
				if err != nil {
					return
				}
			}
		case explanation, opened := <-activity.ExplanationChan:
			if opened {
				err := c.WriteJSON(explanation)
				if err != nil {
					return
				}
			}
		case result, opened := <-activity.UserChan[uId]:
			if opened {
				err := c.WriteJSON(result)
				if err != nil {
					activity.ReleaseChan(uId)
					return
				}
			}
		case <-activity.EndChan:
			log.Println("本次问答已结束...")
			_ = c.WriteMessage(1, []byte("本次问答已经结束，感谢参与!!"))
			return
		}
	}
}
