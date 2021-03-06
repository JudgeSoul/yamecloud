package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	message = "message"
	data    = "data"
	errors  = "errors"
)

func RequestParametersError(g *gin.Context, err error) {
	g.JSON(http.StatusBadRequest,
		gin.H{data: err.Error(), message: err.Error(), errors: err.Error()},
	)
}

func InternalServerError(g *gin.Context, _data interface{}, err error) {
	g.JSON(http.StatusInternalServerError,
		gin.H{data: _data, message: err, errors: err.Error()},
	)
}
