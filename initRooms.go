package main

func initRooms() map[string]Room {
	roomsMap := map[string]Room{}

	doorOpen := false
	roomsMap["room"] = StartRoom{map[string]string{"door": "end"}, &doorOpen, map[string]bool{"key": true}, map[string]string{"key": "There is a #Bkey#B on the ground."}}
	roomsMap["end"] = Basicroom{"There is an old #Bman#B sleeping on a rocking chair. You can go back where you woke up in the \033[4mroom\033[0m.", map[string]string{"room": "room"}, map[string]bool{}, map[string]string{}}

	return roomsMap
}
