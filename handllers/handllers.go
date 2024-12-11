package handllers

import (
	"context"
	"fmt"

	"github.com/Eng-Daniel-Hermel/twitterGo.git/models"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("Voy a processar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespAPI
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string){
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
			
		}
		//
	
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
				
		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
				
		}
		//
	}

	r.Message="Method Invalid"
	return r

}
