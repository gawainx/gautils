package gautils

func ICStr(src interface{}) *string{
	//convert interface{} to string
	res,ok := src.(string)
	if ok {
		return &res
	}else{
		return nil
	}
}

func ICInt(src interface{}) *int{
	//convert interface{} to int
	res,ok := src.(int)
	if ok{
		return &res
	}else{
		return nil
	}
}
