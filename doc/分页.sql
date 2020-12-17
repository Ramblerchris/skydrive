# 针对正常分页
SELECT * from  tbl_file limit 10 ,8;

# 针对动态内容分页
SELECT * from  tbl_file where id>31 limit 1;


SELECT * from  tbl_file where  id BETWEEN 30 and 40  limit 10;



select * from skydrive.tbl_user_file  where id < 133  order by id desc limit 4;
select * from skydrive.tbl_user_file  where id < 129  order by id desc limit 40;
select max(id) from skydrive.tbl_user_file  ;
select max(id) from skydrive.tbl_user_file order by id desc ;


