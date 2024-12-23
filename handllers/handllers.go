package handllers

import (
	"context"
	"fmt"

	"github.com/Eng-Daniel-Hermel/twitterGo.git/jwt"
	"github.com/Eng-Daniel-Hermel/twitterGo.git/models"
	"github.com/Eng-Daniel-Hermel/twitterGo.git/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("Voy a processar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespAPI
	r.Status = 400

	isOk, statusCode, msg, _ := validoAuthorization(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
	}

	r.Message = "Method Invalid"
	return r

}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOk, msg, err := jwt.ProcessoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOk {
		if err != nil {
			fmt.Println("Error en el Token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el Tpken " + msg)
			return false, 402, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim
}
