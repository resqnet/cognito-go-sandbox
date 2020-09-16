package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {

	mySession := session.Must(session.NewSessionWithOptions(session.Options{Profile: "default"}))
	svc := cognitoidentityprovider.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))

	paramsConfirm := &cognitoidentityprovider.ConfirmSignUpInput{
		Username:         aws.String("USERNAME"),
		ConfirmationCode: aws.String("Verification Code"),
		ClientId:         aws.String("ClientID"),
	}

	_, err := svc.ConfirmSignUp(paramsConfirm)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
