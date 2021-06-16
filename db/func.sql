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

