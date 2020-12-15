# 针对正常分页
SELECT * from  tbl_file limit 10 ,8;

# 针对动态内容分页
SELECT * from  tbl_file where id>31 limit 1;


SELECT * from  tbl_file where  id BETWEEN 30 and 40  limit 10;


