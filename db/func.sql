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
where act_id = ?