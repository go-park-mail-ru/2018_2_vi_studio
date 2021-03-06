package handlers

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"io/ioutil"
	"main/middleware"
	"main/proto"
	"net/http"
)

const chunkSize = 1024

type UserAvatarHandler struct {
	sb  proto.AuthServices
	svc *s3.S3
}

func NewUserAvatarHandler(sb proto.AuthServices) *UserAvatarHandler {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("ru-msk"),
		Endpoint: aws.String("hb.bizmrg.com"),
	}))

	return &UserAvatarHandler{
		sb:  sb,
		svc: s3.New(sess),
	}
}

func (uah *UserAvatarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	session, ok := middleware.Session(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileUuid, _ := uuid.NewRandom()
	fileName := fileUuid.String()
	_, err = uah.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("vi-studio"),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(body),
		ACL          :aws.String("public-read"),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = uah.sb.Users.AddAvatar(context.Background(), &proto.AddAvatarRequest{
		UserUID: session.UserUID,
		Name:    fileName,
	})
	if err != nil {
		middleware.Logger(r.Context()).Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(`{"status":"ok"}`))
	w.WriteHeader(http.StatusOK)
}
