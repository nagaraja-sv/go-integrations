package cloudstorageexample

/*func (signupController) KycDetails(ctx *gin.Context) {

	utils.Log.Info(map[string]interface{}{
		"RequestURI": ctx.Request.RequestURI,
		"method":     ctx.Request.Method,
		"remoteAddr": ctx.Request.RemoteAddr,
		"event":      "KycDetails",
	}, "kyc request received")

	var userkyc models.Userkyc

	ctx.Writer.Header().Set("Content-Type", "application/json")

	file, header, err := ctx.Request.FormFile("upload")
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./tmp/" + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	if err := ctx.BindJSON(&userkyc); err != nil {
		log.Println(err)
		errcode := 11
		ctx.JSON(services.Error[errcode].HTTPCode, gin.H{
			"err": errcode,
			"str": services.Error[errcode].PublicError,
		})
		return
	}

	fmt.Println("userkyc data is:", userkyc)

	errcode, uid := services.Auth.IsValidToken(ctx.GetHeader("Authorization"))
	if errcode != 51 {
		ctx.JSON(services.Error[errcode].HTTPCode, gin.H{
			"err":  errcode,
			"str":  services.Error[errcode].PublicError,
			"auth": "",
		})
		return
	}

	userkyc.UID = uid

	errcode = services.Signup.KycDetails(&userkyc)

	ctx.JSON(services.Error[errcode].HTTPCode, gin.H{
		"err": errcode,
		"str": services.Error[errcode].PublicError,
	})

}
*/
