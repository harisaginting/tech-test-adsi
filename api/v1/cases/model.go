package cases

type Response struct{
	Odd 	int 		`json:"odd"`
}

type Payload struct{
	Data 	[]int 		`json:"data"`
}