package models

import (
	"errors"
	"rest_api/project/db"
	"rest_api/project/utils"
)

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error{
	query:="INSERT INTO users(email,password) VALUES(?,?)"
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	
	defer stmt.Close()
	pass,err:=utils.HashGenerator(user.Password)
	if err!=nil{
		return err
	}
	result,err:=stmt.Exec(user.Email,pass)
    if err!=nil{
        return err
    }
    id,err:=result.LastInsertId()
    user.ID=id
    
    return err

}

func (user *User) ValidateCreds()error{
	query:="SELECT id,password from users where email=?"
	row:=db.DB.QueryRow(query,user.Email)
	var retrivedpassword string
	err:=row.Scan(&user.ID,&retrivedpassword)
	if err!=nil{
		return err
	}
	passwdvalid:=utils.CheckHashPassword(user.Password,retrivedpassword)
	if !passwdvalid{
		return errors.New("Email or password is Wrong!")
	}
	return nil
}