insert into user_info(u_name, u_psw)
values (?, ?);

select u_id
from user_info
where u_name = ? and u_psw = ?;

select u_id
from user_info
where u_name = ?;

select count(*)
from act
where act_id = ?;

insert into act(act_id, org_id,act_name,act_len,act_des,act_stop)
values (?, ?, ?, ?, ?, ?);

update act
set act_stop=? ,act_final=?
where act_id = ?;

select org_period, count(u_id)
from act_period,vote
where act_id=? and act_period.period_id=vote.period_id
group by org_period;

select count(*)
from act,act_period
where act.act_id=? and org_id=? and act_period.org_period=? and act.act_id=act_period.act_id

/*
select act_stop,act_len
from act
where act_id = ?

insert into vote(period_id, act_id, org_period)
values (?, ?, ?)

select act_stop,act_len,org_id,act_des
from act,user_info
where act_id = ? and org_id == u_id

select act_stop,act_len,org_id,act_des,act_name
		from act,user_info
		where act_id = ? and org_id == u_id

select unique act_id
from act_period,vote
where u_id = ? and act_period.period_id == vote.period_id //query

select org_period,period_id
from act,act_period
where act_id = ? and act.act_id == act_period.act_id //query
*/

select ap.period_id
from act_period ap
where ap.act_id = ? and ap.org_period = ?
