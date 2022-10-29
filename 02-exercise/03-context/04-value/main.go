package main

import (
	"fmt"
	"context"
)

type database map[string]bool

var db database = database{
	"jane": true,
}
type userIDType string

func main() {
	processRequest("jane")
}

func processRequest(userid string) {
	// TODO: send userID information to checkMemberShip through context for
	// map lookup.
	key:= userIDType("userIDKey")
	ctx:= context.WithValue(context.Background(),key,userid)
	ch := checkMemberShip(ctx, key)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShip - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do map lookup
// sends the result on the returned channel.
func checkMemberShip(ctx context.Context, key userIDType) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		userid:= ctx.Value(key).(string)
		status := db[userid]
		ch <- status
		
		
		
	}()
	return ch
}
