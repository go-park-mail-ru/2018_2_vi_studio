package handlers

import (
	"io"
	"main/proto"
	"net/http"
)

const chunkSize = 1024

type UserAvatarHandler struct {
	sb proto.AuthServices
}

func NewUserAvatarHandler(sb proto.AuthServices) *UserAvatarHandler {
	return &UserAvatarHandler{
		sb: sb,
	}
}

func (uah *UserAvatarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//session, ok := middleware.Session(r.Context())
	//if !ok {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	//
	//fileUuid, _ := uuid.NewRandom()
	//fileName := fileUuid.String() + ".jpg"
	//file, err := os.Create("./media/images/" + fileName)
	//if err != nil {
	//	middleware.Logger(r.Context()).Panic(err.Error())
	//	return
	//}
	//
	//defer func() {
	//	if err := file.Close(); err != nil {
	//		middleware.Logger(r.Context()).Panic(err.Error())
	//		return
	//	}
	//}()

	chunk := make([]byte, chunkSize)
	for {
		// read a chunk
		n, err := r.Body.Read(chunk)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		//if n == chunkSize {
		//	_, err = file.Write(chunk)
		//} else {
		//	_, err = file.Write(chunk[:n])
		//}
		if err != nil {
			panic(err)
		}
	}

	//err = uah.sb.Users.AddAvatar(*session.UserId, fileName)
	//if err != nil {
	//	middleware.Logger(r.Context()).Error(err.Error())
	//}
}
