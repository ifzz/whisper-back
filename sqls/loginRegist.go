package sqls

import "fmt"

//Login SQL 从数据库中获取邮箱对应的密码，返回给serve
func Login(mail string) (string, int) {
	var (
		pw    string
		power int
	)
	row := Db.QueryRow("select password,power from user where mail=?", mail)
	row.Scan(&pw, &power)
	return pw, power
}

//Regist SQL 先查询是否存在重复的邮箱或昵称，再插入
func Regist(name, mail, pw string) (result string, code int) {
	//传入的参数分别对应着mail userName password userName mail
	pre, err := Db.Prepare(`INSERT INTO user (mail, userName, password)
		SELECT ?,?,?
		from DUAL
		WHERE not exists (
				SELECT *
				from user
				WHERE userName = ? or mail = ? LIMIT 1);`)
	if err != nil {
		fmt.Println("预编译表达式出错", err.Error())
	}
	effects, err := pre.Exec(mail, name, pw, name, mail)
	if err != nil {
		fmt.Println("写入用户数据，执行SQL出错", err.Error())
	}
	//如果rownum==0，说明没有插入数据，即用户名邮箱已存在
	rownum, _ := effects.RowsAffected()
	if rownum == 1 {
		result = "注册成功"
		code = 200
	} else {
		result = "用户名或邮箱已注册"
		code = 201
	}
	return
}
