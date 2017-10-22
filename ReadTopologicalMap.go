package main

import(
)

type Cell struct {
	longititude float64
	latitude float64
	altitude float64
	terranType string
}

func ReadTopologicalMap() [][]Cell {
	mapname := os.Args[1]
	maphandle,_ := os.Open(mapname)
  //read in all the map variable inc
  
	maphandle.Close()
}



func main() {

  ReadTopologicalMap()

}
