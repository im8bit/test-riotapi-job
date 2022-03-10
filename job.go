package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/im8bit/test-riotapi-library/aws"
	"github.com/im8bit/test-riotapi-library/riot"
)

func main() {
	var activeActId string = riot.GetActiveActId()

	fmt.Printf("Active Act ID: %s\n", activeActId)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	fmt.Println("Droping Table")
	aws.DropTable(svc)
	fmt.Println("Creating Table")
	aws.CreateTable(svc)

	var leaderboardDtoData riot.LeaderboardDto = riot.GetLeaderboard(activeActId)

	fmt.Println("Adding Items to Table")
	for _, player := range leaderboardDtoData.Players {
		aws.AddLeaderboardItem(svc, activeActId, player)
	}
}
