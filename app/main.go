package main

import (
	"log"
	"sync"

	"patientdx-backend-web-go/db"
	"patientdx-backend-web-go/patient_dx/delivery"
	"patientdx-backend-web-go/patient_dx/repository"
	"patientdx-backend-web-go/patient_dx/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	Init()
	// initEnv()
	listenPort := ":4000"
	// appName := os.Getenv("APP_NAME")

	patientRepo := repository.NewPostgrePatient(db.GormClient.DB)

	timeoutContext := fiber.Config{}.ReadTimeout

	patientUseCase := usecase.NewPatientUseCase(patientRepo, timeoutContext)

	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${green} ${status} ${white} | ${latency} | ${ip} | ${green} ${method} ${white} | ${path} | ${yellow} ${body} ${reset} | ${magenta} ${resBody} ${reset}\n",
		TimeFormat: "02 January 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(cors.New())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {

		//call delivery http here
		delivery.NewPatientHandler(app, patientUseCase)
		log.Fatal(app.Listen(listenPort))
		wg.Done()
	}()
	wg.Wait()

}

func Init() {
	InitEnv()
	InitDB()
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}
}

func InitDB() {
	db.NewGormClient()
}
