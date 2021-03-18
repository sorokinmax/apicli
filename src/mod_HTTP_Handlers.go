package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func runHandler(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var program = ctx.PostForm("program")
	var args = ctx.PostForm("args")

	argsArray, err := parseCommandLine(args)
	if err != nil {
		log.Errorf("Cannot parse \"%s\": %s\n", args, err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stdout, stderr, err := run(program, argsArray)
	if err != nil {
		log.Errorf("Run failed with: %s\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stdout": stdout, "stderr": stderr})
}

func pingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
