package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Index(rw http.ResponseWriter, r *http.Request){

	var err error
	//rooms:=&Roomss{}
	//rooms.Roooms= make([]*Room,0,200)
	//
	//rows, err := dbc.Query("select * from `#####` ")
	//if err != nil {
	//	panic(err)
	//}
	//for rows.Next() {
	//	room:=&Room{}
	//	err = rows.Scan(&room.IDr,&room.Game,&room.Players)
	//	if err != nil {
	//		panic(err)
	//	}
	//	rooms.Roooms=append(rooms.Roooms,room)
	//
	//}
	//rows.Close()

	var rooms [] Room

	for i:=1;i<=10;i++{
		room:=Room{}
		room.ID=strconv.Itoa(i)
		room.Game="Pocker"
		for j:=1;j<=3;j++{
			player:=Player{}
			player.ID=strconv.Itoa(j)
			player.Name="BOB" + strconv.Itoa(j)
			room.Players= append(room.Players, &player)

		}
		rooms=append(rooms,room)
	}

	//fmt.Printf("%v\n",rooms)

	err = tpl.ExecuteTemplate(rw, "index.html", rooms)
	if err != nil {
		fmt.Printf("INDEX execute template:%s \n", err)
	}

}