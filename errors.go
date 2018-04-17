package gautils

//Error status code
const(
	GeneralError 			= -1
	InternetError 			= -2
	NormalStatus 			= 0
)

func CheckGenError(err error) int{
	if err!=nil{
		Error.Printf(err.Error())
		return GeneralError
	}else{
		return NormalStatus
	}
}

func CheckFatalError(err error){
	if err != nil{
		Error.Fatalln(err.Error())
	}
}

func CheckInternetError(err error) int {
	if err != nil{
		Error.Println(err.Error())
		return InternetError
	}else{
		return NormalStatus
	}
}