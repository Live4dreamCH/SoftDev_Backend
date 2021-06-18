package main

type GetAct_json struct {
	Sid   string `json:"SessionID" binding:"required"`
	ActID int    `json:"Actid" binding:"required"`
}

/*func (j *GetAct_json) json2Act() (a obj.Act, err error) {
	uid, err := sid_manager.get(j.Sid)
	if err != nil {
		return
	}
	a.Aid = j.ActID
	a.Uid = uid
	return
}*/
