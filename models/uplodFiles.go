package models

import (
	"DataCertPlatfom/db_mysql"
	"fmt"
)
type UploadFile struct {
	Id        int
	UserId    int
	FileName  string
	FileSize  int64
	FileCert  string
	FileTitle string
	CertTime  string
}
func (u *UploadFile) AddFiles() (int64, error) {
	result, err := db_mysql.Db.Exec("INSERT INTO upload_record" +
		"(user_id,file_name,file_size,file_cert,file_title,cert_time)" +
		"value (?,?,?,?,?,?)",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)
	if err != nil {
		fmt.Println("数据库存储文件出错，请重试！")
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rows, nil
}
func QueryRecordsByUserId(userId int)([]UploadFile,error)  {
	fmt.Println("当前登录用户的ID为:",userId)
	rows,err := db_mysql.Db.Query("select id,user_id,file_name,file_size,file_cert,file_title,cert_time from upload_record where user_id =?",userId)
	if err != nil {
		return nil,err
	}
	var records = make([]UploadFile, 0)
	var record =UploadFile{}
	for rows.Next()  {
		err = rows.Scan(&record.Id,&record.UserId,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTitle,&record.CertTime)
		if err != nil {
			return nil,err
		}
		records = append(records, record)
	}
	return records,nil
}
