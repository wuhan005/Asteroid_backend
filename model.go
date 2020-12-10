package main

type unityData struct {
	Type string
	Data interface{}
}

type greet struct {
	Title string
	Time  int
	Round int
	Team  []team
}

type team struct {
	Id    int
	Name  string
	Rank  int
	Image string
	Score int
}

type attack struct {
	From int
	To   int
}

type rank struct {
	Team []team
}

type status struct {
	Id     int
	Status string
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
