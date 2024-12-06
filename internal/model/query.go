package model

type QueryParameter struct {
	ParameterId   int
	QueryId       int
	ParameterName string
	DataType      string
	Ordered       int
}

type Queries struct {
	QueryId          int
	QueryDescription string
	Query            string
	Params           []QueryParameter
}
