package main

import (
	"backend-app/Components"
	"backend-app/Reportes"
	"backend-app/parser"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type SintacticErrorListener struct {
	*antlr.DefaultErrorListener
}

type Error struct {
	Fila        int    `json:"fila"`
	Columna     int    `json:"columna"`
	Intruccion  string `json:"instruccion"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

type Gsymbol struct {
	Id         string `json:"id"`
	Line       int    `json:"line"`
	Col        int    `json:"col"`
	TypeSymbol string `json:"typesymbol"`
	Type       string `json:"type"`
	Env        string `json:"env"`
}

type Console struct {
	Key      int    `json:"key"`
	Terminal string `json:"terminal"`
}

var ListErrores []Error
var ListSymbol []Gsymbol
var ListConsole []Console

func (l *SintacticErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	posicionSol := strings.Index(msg, "extraneous")
	posicionDia := strings.Index(msg, "expecting ")
	cadenaModificada := msg[:posicionSol] + msg[posicionDia+len("expecting "):]

	Error := Error{
		Fila:        line,
		Columna:     column,
		Intruccion:  offendingSymbol.(antlr.Token).GetText(),
		Descripcion: cadenaModificada,
		Tipo:        "SINTACTICO",
	}
	ListErrores = append(ListErrores, Error)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./public")

	app.Post("/File", func(c *fiber.Ctx) error {
		cadenaRecibida := string(c.Body())

		ListErrores = []Error{}
		ListSymbol = []Gsymbol{}
		ListConsole = []Console{}

		is := antlr.NewInputStream(cadenaRecibida)
		lexico := parser.NewSintaxLexer(is)
		tokens := antlr.NewCommonTokenStream(lexico, antlr.TokenDefaultChannel)
		sintactico := parser.NewSintaxParser(tokens)

		sintactico.RemoveErrorListeners()
		sintactico.AddErrorListener(&SintacticErrorListener{})

		tree := sintactico.Inicio()
		eval := Components.NewVisitor(Components.NewEntorno(nil, "Global"))
		eval.Visit(tree)

		for _, env := range eval.ListErrores {
			Error := Error{
				Fila:        env.Fila,
				Columna:     env.Columna,
				Intruccion:  env.Intruccion,
				Descripcion: env.Descripcion,
				Tipo:        env.Tipo,
			}
			ListErrores = append(ListErrores, Error)
		}

		for key, env := range eval.TableGlobal {
			Gsymboll := Gsymbol{
				Id:         key,
				Line:       env.Line,
				Col:        env.Col,
				TypeSymbol: env.TypeSymbol,
				Type:       env.Type,
				Env:        env.Env,
			}
			ListSymbol = append(ListSymbol, Gsymboll)
		}

		for key, env := range eval.ListConsole {
			Console := Console{
				Key:      key,
				Terminal: env,
			}
			ListConsole = append(ListConsole, Console)
		}

		data := map[string]interface{}{
			"ListSymbol":  ListSymbol,
			"ListError":   ListErrores,
			"ListConsole": ListConsole,
		}

		Reportes.CreateCST(cadenaRecibida)

		return c.JSON(data)
	})

	app.Post("/AST", func(c *fiber.Ctx) error {
		svgData, err := ioutil.ReadFile("cst.svg")
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("ERROR AL GENERAR ARCHIVO")
		}
		c.Type("image/svg+xml")
		return c.Send(svgData)
	})

	app.Listen(":5000")
}
