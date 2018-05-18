package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

var debugOn bool = false
var debugOut []string

type MyEvent struct {
	Message  string `json:"message,omitempty"`
	Password string `json:"password"`
	Action   string `json:"action"`
	Debug    bool   `json:"debugon,omitempty"`
}

type MyResponse struct {
	Message     string   `json:"result"`
	DebugOutput []string `json:"debugOutput,omitempty"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {

	debugOn = event.Debug
	var resp string
	debugOut = []string{}

	//check debug
	if debugOn {
		debugOut = []string{"INFO: Start debug", event.Action, event.Message, event.Password}
	}

	//decrypt or encrypt
	if event.Action == "encrypt" {
		fixedstr := strings.Replace(event.Message, " ", "_", -1)
		resp = kylecrypt(fixedstr, event.Password)

	} else if event.Action == "decrypt" {
		underscorestr := kyleuncrypt(event.Message, event.Password)
		resp = strings.Replace(underscorestr, "_", " ", -1)
	} else {
		return MyResponse{Message: fmt.Sprintf("Action %s not allowed, please use 'encrypt' or 'decrypt' in this field. For more info, check the documentation! jk, there is none", event.Action), DebugOutput: debugOut}, nil
	}

	//return with or without debug info
	if debugOn {
		return MyResponse{Message: resp, DebugOutput: debugOut}, nil
	}
	return MyResponse{Message: resp}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)

	// msgPtr := flag.String("message", "Hello, World!", "your message")
	// pwdPtr := flag.String("password", "password", "your password")
	// kylecrypttrue := flag.Bool("kylecrypt", false, "encrypt message")
	// kyleuncrypttrue := flag.Bool("kyleuncrypt", false, "decrypt message")
	// debug := flag.Bool("debug", false, "print stuff")

	// flag.Parse()
	// debugOn = *debug

	// if debugOn {
	// 	fmt.Println(*msgPtr, *pwdPtr, *kylecrypttrue, *kyleuncrypttrue)
	// }

	// if *kylecrypttrue {
	// 	fixedstr := strings.Replace(*msgPtr, " ", "_", -1)
	// 	fmt.Println(kylecrypt(fixedstr, *pwdPtr))
	// }

	// if *kyleuncrypttrue {

	// }

}

func kylecrypt(message, password string) string {
	encr := []byte{}
	for i, l := range message {
		lbyte := uint(l)
		if (i % 2) > 0 {
			lbyte += uint(len(password))
		} else {
			lbyte -= uint(len(password))
		}
		if lbyte >= 126 {
			return "Does not compile"
		}
		encr = append(encr, byte(lbyte))
		if debugOn {
			debugOut = append(debugOut, string(string(int(i))+" | char "+string(l)+" | crypt "+string(lbyte)+" | bytes "+string(encr)))
		}
	}
	return string(encr)
}

func kyleuncrypt(message, password string) string {
	dencr := []byte{}
	for i, l := range message {
		lbyte := uint(l)
		if (i % 2) > 0 {
			lbyte -= uint(len(password))
		} else {
			lbyte += uint(len(password))
		}
		dencr = append(dencr, byte(lbyte))
		if debugOn {
			debugOut = append(debugOut, string(string(int(i))+" | char "+string(l)+" | decrypt "+string(lbyte)+" | bytes "+string(dencr)))
		}
	}
	return string(dencr)
}
