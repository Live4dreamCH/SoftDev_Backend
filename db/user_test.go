package db

import "testing"

func TestUser(t *testing.T) {
	var dbu DB_user
	t.Log(dbu.Search("Tom 李"))
	t.Log(dbu.LoginQuery("Tom 李", "password123456"))
	t.Log(dbu.Create("Tom 李", "password123456"))
	t.Log(dbu.CreateVote(1, 123, []int{1}))
	t.Log(dbu.GetName(1))
}
