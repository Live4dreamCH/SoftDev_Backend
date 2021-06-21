package db

import "testing"

func TestAct(t *testing.T) {
	var dba DB_act
	t.Log(dba.ActinfoQuery(1))
	t.Log(dba.OnepersonActQuery(1))
	t.Log(dba.ActivityQuery(1))
	t.Log(dba.PeriodQuery(1))
	t.Log(dba.GetPeriodID(1, "2021-6-21 0:0:0"))
	t.Log(dba.Search(1))
	t.Log(dba.Search_aid(1))
	t.Log(dba.Search_aid_uid(1, 1))
	t.Log(dba.Search_vote(1))
	t.Log(dba.Search_period(1, 1, "2021-6-21 0:0:0"))
	t.Log(dba.Create(1, 1, "", 1, ""))
	t.Log(dba.Create_period(1, []string{"2021-6-21 0:0:0"}))
	t.Log(dba.Update_final(1, "2021-6-21 0:0:0"))
}
