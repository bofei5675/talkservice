package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "net/http"
    "talkapp/biz/model"
    "talkapp/biz/service"
)

func CompareHandler(c *gin.Context) {
    var params model.CompareParameter
    if err := c.ShouldBindJSON(&params); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    method := c.Param("method")
    logrus.Infof("received parameters %v and method %s", params, method)
    score, err := service.CompareTwoString(params.Source, params.Target, method)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success", "score": score})
}