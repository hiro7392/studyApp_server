package main

import (
	"database/sql"
	"fmt"
	"log"
	//"os"
	"github.com/aws/aws-sdk-go/aws"
  	//"github.com/aws/aws-sdk-go/gen/ec2"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	
	dbName:="studyappaws2"
	dbUser:="user2"
    dbHost := "studyappaws2.coyiikoe4d90.ap-northeast-1.rds.amazonaws.com"
    dbPort := 3307
    dbEndpoint := fmt.Sprintf("%s:%d", dbHost, dbPort)
    region := "ap-northeast-1d"
	
    //creds := credentials.NewCredentials();
	creds := aws.IAMCreds()//credentials.NewStaticCredentials(aws_access_key_id,aws_secret_access_key,x509);
	credValue,err:=creds.Get()
	if err !=nil{
		log.Println(err.Error())
	}
	fmt.Println("checkpoint1",credValue)

    authToken, err := rdsutils.BuildAuthToken(dbEndpoint, region, dbUser,creds)
    if err != nil {
        log.Println(err.Error())
    }
	fmt.Println("checkpoint 2")

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&allowCleartextPasswords=true",
        dbUser, authToken, dbEndpoint, dbName,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

	
}