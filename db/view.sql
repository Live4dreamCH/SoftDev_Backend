use app;
create view full_act
as select a.act_id, a.act_name, a.org_id, u1.u_name org_name, a.act_stop, a.act_final, p.org_period, v.u_id voter_id, u2.u_name voter_name
from act a 
	inner join user_info u1 on a.org_id = u1.u_id 
	left outer join act_period p on a.act_id = p.act_id
	left outer join vote v on v.period_id = p.period_id
	left outer join user_info u2 on v.u_id = u2.u_id;