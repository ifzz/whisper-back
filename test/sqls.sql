SELECT  userid  from user 
WHERE (userName='深蓝' or mail='mail@ri-co.cn');


/* 如果不存在内循环的值，就执行插入语句 */
INSERT INTO user (`mail`, `userName`, `password`)
SELECT 'mal@ri.cn',
    'rrw',
    '123456'
from DUAL
WHERE not exists (
        SELECT *
        from user
        WHERE userName = '蓝'
            or mail = 'mal@ri-co.cn'
        LIMIT 1
    );