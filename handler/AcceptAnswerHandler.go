package handler

import (
	"github.com/cinience/sdc/lib"
	"github.com/cinience/sdc/model"
	"strconv"
)

type AcceptAnswerHandler struct {
	lib.AuthHandler
}

func (self *AcceptAnswerHandler) Get() {
	aid, _ := self.GetInt64(":aid")
	qid, _ := self.GetInt64(":qid")
	uid, _ := self.GetSession("userid").(int64)
	role, _ := self.GetSession("userrole").(int64)
	if aid > 0 && qid > 0 && uid > 0 {
		if err := model.SetAcceptAnswer(qid, aid, uid, role); err == nil {
			self.Redirect("/"+strconv.Itoa(int(qid))+"/", 302)
		} else {
			self.Redirect("/", 302)
		}
	} else {
		self.Redirect("/", 302)
	}
}
