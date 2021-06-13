insert into user_info(u_name, u_psw)
values (?, ?);

select u_id
from user_info
where u_name = ? and u_psw = ?;

select u_id
from user_info
where u_name = ?;