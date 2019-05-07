package handlers

import (
	//"database/sql"
	"fmt"
	"net/http"
)

func Rooms(rw http.ResponseWriter, r *http.Request){

	//listPlayer:=&Room{}
	//listPlayer.Players=make([]*Player,0,4)
	//var err error
	//var rows *sql.Rows
	//var id int

	idString := r.FormValue("id")
	fmt.Printf("id from room: %s\n", idString)

	//id,err: = strconv.Atoi(idString)
	//if err != nil {
	//	fmt.Printf("id room error:%s; ID:%s\n", err, idString)
	//	http.Redirect(rw, r, "index.html", 302)
	//	return
	//}





	if r.Method == "GET"{

		//rows,err=dbc.Query("select `room` where `id`= ?",id)
		//if err!= nil{
		//	fmt.Printf("dbc.Query room error: %s",err)
		//	http.Redirect(rw,r,"index.html",302)
		//	return
		//}
		//for rows.Next() {
		//	player := &Player{}
		//	err := rows.Scan(&player.ID, &player.Name, /*TODO &player.Cards*/)
		//	if err != nil {
		//		panic(err)
		//	}
		//	listPlayer.Players = append(listPlayer.Players, player)
		//
		//}
		//rows.Close()
		//
		//err = tpl.ExecuteTemplate(rw, "room.html", listPlayer)
		//if err != nil {
		//	fmt.Printf("ListPlayer ExecuteTemplate error: %s\n", err)
		//}
		//return



		//var rooms=[]*Room{}
		//var room=[]Room{}
		//for _,rr:= range rooms{
		//	if rr.ID!= idString{
		//		continue
		//	}
		//	fmt.Printf("%s %s %s\n",rr.ID,rr.Game,rr.Players)
		//	room=append(room,*rr)
		//}
		p:=&Player{"111","BANG",nil}
		p1:=&Player{"222","NIG",nil}
		p2:=&Player{"333","JOB",nil}
		room:=&Room{"2", "Pocker" ,[]*Player{
				p,p1,p2,
			},
		}

		err := tpl.ExecuteTemplate(rw, "room.html", room)
		if err != nil {
			fmt.Printf("POST room error: %s\n", err)
		}

	}else if r.Method == "POST"{

		player1_card1:=r.PostFormValue("player1_card1")
		fmt.Printf("P1C1: %s\n",player1_card1)

		player1_card2:=r.PostFormValue("player1_card2")
		fmt.Printf("P1C2: %s\n",player1_card2)

		player1_card3:=r.PostFormValue("player1_card3")
		fmt.Printf("P1C3: %s\n",player1_card3)

		player2_card1:=r.PostFormValue("player2_card1")
		fmt.Printf("P2C1: %s\n",player2_card1)

		player2_card2:=r.PostFormValue("player2_card2")
		fmt.Printf("P2C2: %s\n",player2_card2)

		player2_card3:=r.PostFormValue("player2_card3")
		fmt.Printf("P2C3: %s\n",player2_card3)


	}




}