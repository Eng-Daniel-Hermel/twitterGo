package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Eng-Daniel-Hermel/twitterGo.git/bd"
	"github.com/Eng-Daniel-Hermel/twitterGo.git/models"
)

func Registro(ctx context.Context) models.RespAPI {
	var t models.Usuario
	var r models.RespAPI
	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe espeificar el Email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "Debe espeificar una contrseña de al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario registrado com es Email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertoRegisstro(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar realizar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro OK"
	fmt.Println(r.Message)
	return r

}
