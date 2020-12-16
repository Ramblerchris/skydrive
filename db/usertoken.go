package db

import (
	"fmt"
	"github.com/skydrive/config"
	mysql "github.com/skydrive/db/mysqlconn"
	"github.com/skydrive/utils"
	"time"
)

const (
	selectUTokenInfoByPhone = "select id ,uid,phone,user_token from tbl_user_token where phone=? "
	selectUTokenInfoByUId   = "select id ,uid,phone,user_token,expiretime from tbl_user_token where uid=?"
	selectUTokenInfoByToken = "select id ,uid,phone,user_token,expiretime from tbl_user_token where user_token=?"
	saveUToken              = "insert into tbl_user_token( uid,phone,user_token,expiretime) values(?,?,?,?)"
	deleteUToken            = "delete from  tbl_user_token where id=?"
)

func (uToken TableUToken) String() string {
	return fmt.Sprintf("Tid:%d Uid:%d  Phone: %s  User_token: %s  Expiretime: %d ",
		uToken.Tid.Int64, uToken.Uid.Int64, uToken.Phone.String, uToken.User_token.String, uToken.Expiretime.Int64)
}

/*func (uToken *TableUToken) Stringer() string{
	return fmt.Sprintf("Tid:%d Uid:%d  Phone: %s  User_token: %s  Expiretime: %d ",
		uToken.Tid, uToken.Uid, uToken.Phone, uToken.User_token, uToken.Expiretime)
}*/

//删除token表中一条记录
func DeleteUserTokenByTid(tid int64) bool {
	stmt, error := mysql.DbConnect().Prepare(deleteUToken)
	if error != nil {
		fmt.Println("failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(tid)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，但未添加，
			fmt.Println("failed  exec", error)
			return false
		}
	}
	return true
}
func GetUserTokenInfoByToken(token string) (utoken TableUToken, err error) {
	stmt, error := mysql.DbConnect().Prepare(selectUTokenInfoByToken)
	if error != nil {
		fmt.Println("failed to prepare statement error:", error.Error())
		return utoken, error
	}
	defer stmt.Close()
	utoken = TableUToken{}
	error = stmt.QueryRow(token).Scan(
		&utoken.Tid,
		&utoken.Uid,
		&utoken.Phone,
		&utoken.User_token,
		&utoken.Expiretime,
	)
	if error != nil {
		fmt.Println("failed to QueryRow error:", error)
		return utoken, error
	}
	fmt.Println("token data :",utoken)
	return utoken, nil
}

func GetUserTokenInfoListByUid(uid int32) (tokenlist []TableUToken, err error) {
	stmt, error := mysql.DbConnect().Prepare(selectUTokenInfoByUId)
	if error != nil {
		fmt.Println("failed to prepare statement error:", error.Error())
		return tokenlist, error
	}
	defer stmt.Close()
	rows, error := stmt.Query(uid)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return tokenlist, error
	}
	tokenlist = make([]TableUToken, 0)
	//顺序要保持一直
	for rows.Next() {
		var utoken TableUToken
		if error := rows.Scan(&utoken.Tid, &utoken.Uid, &utoken.Phone, &utoken.User_token); error != nil {
			fmt.Println("error is", error)
			continue
		}
		tokenlist = append(tokenlist, utoken)
	}
	fmt.Println("token list data", tokenlist)
	return tokenlist, nil
}

//创建用户一个新的token
func CreateUserTokenByUidPhone(uid int32, phone string) (token string, err error) {
	stmt, error := mysql.DbConnect().Prepare(saveUToken)
	if error != nil {
		fmt.Println("failed to prepare statement error:", error.Error())
		return token, error
	}
	defer stmt.Close()
	uuid := utils.BuildUUID()
	//毫秒
	//haomiao := time.Now().AddDate(0, 0, 7).UnixNano()/1e6
	haomiao := time.Now().Add(config.Token_ExpriseTime).UnixNano()/1e6
	fmt.Println(" Expiretime :", haomiao)
	exec, error := stmt.Exec(uid, phone, uuid, haomiao)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return token, error
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，但未添加，
			fmt.Println("failed  exec", error)
			return token, error
		}
	}
	return uuid, nil
}
