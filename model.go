package main

type unityData struct {
	Type string
	Data interface{}
}

type team struct {
	Id    int
	Name  string
	Rank  int
	Score int
}

type attack struct {
	From int
	To   int
}

type rank struct {
	Team []struct {
		Id    int
		Rank  int
		Score int
	} `binding:"required"`
}

type status struct {
	Id     int
	Status string `binding:"required"`
}

type round struct {
	Round int
}

type clock struct {
	Time int
}

type clearStatus struct {
	Id int
}
