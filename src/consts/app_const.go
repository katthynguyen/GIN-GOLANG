package consts

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

const (
	AppName = "cms-comment"
)

var (
	MDbDatabase string
)

func init() {
	MDbDatabase = os.Getenv("BOX_SUB_MONGO_URI")
}
