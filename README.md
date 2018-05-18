# Golang-Encryption-AWSLambda
This is an encryption algorithm I developed and deployed as an aws lamda with the API Gate as a trigger

Execution details
-------------------------------------
ENCRYPT A MESSAGE WITH YOUR PASSWORD
-------------------------------------
    POST to this url:
    https://2z6ck5bqaf.execute-api.us-east-1.amazonaws.com/dev
    Headers:
    Content-Type: application/json
    Body:
    {
    	"action":"encrypt",
	"message":"your message",
	"password":"your password"
    }

-------------------------------------
DECRYPT A MESSAGE WITH YOUR PASSWORD
-------------------------------------
    POST to this url:
    https://2z6ck5bqaf.execute-api.us-east-1.amazonaws.com/dev
    Headers:
    Content-Type: application/json
    Body:
    {
	"action":"decrypt",
	"message":"your message",
	"password":"your password"
    }

------------------------------------
OPTIONAL FLAGS
------------------------------------
    Adding this extra parameter will show you the process execution logs
    Body:
    {
	"action":"encrypt",
	"message":"your message",
	"password":"your password",
        "debugon":true
    }


