package middleware

import (
	"github.com/tongvinh/celeritas"
	"myapp/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
