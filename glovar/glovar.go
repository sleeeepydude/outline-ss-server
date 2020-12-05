package glovar




var uploadData[] UploadParams
type User struct {
	u int64
	d int64
	id int64
}

type Server struct {
	rate float32
}

type UploadParams struct {
	u int64
	d int64
	user_id int
}

func addTraffic(d int64, u int64, userID int) {
	//fmt.Println("流量")
	userPos := findPosById(uploadData, userID)
	if userPos == -1 {
		var data = UploadParams{ d: d, u: u, user_id: userID}
		slice := append(uploadData, data)
		uploadData = slice

	}else{
		uploadData[userPos].d = uploadData[userPos].d + d
		uploadData[userPos].u = uploadData[userPos].u + u
	}

}

func getTraffic() [] UploadParams {
	return uploadData
}

func findPosById(items []UploadParams, item int) int {
	for i, eachItem := range items {
		if eachItem.user_id == item {
			//fmt.Println("找到" +  strconv.Itoa(item)+ "在第" + strconv.Itoa(i) + "个")
			return i
		}
	}
	return -1
}
