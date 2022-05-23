package main

import (
	"main/aliyun_services/action"
	"main/aliyun_services/global"
	"main/aliyun_services/middle"
)

func main() {

	global.Init()
	global.Root.POST("/reflash_token", global.Reflash_token)

	group := global.Root.Group("/action")
	group.Use(middle.Token_exist_mid)
	group.GET("/get", action.Get_URL)
	group.GET("/search", action.Seach_info)
	group.POST("/upload_file", action.Upload_file)
	group.POST("/delete", action.Delete_file)

	global.Root.Run(":" + global.Port)

	//a := al.New()

	//drive := aliyundrive.NewClient(&aliyundrive.Options{
	//	AutoRefresh: true,
	//	UploadRate:  2 * 1024 * 1024, // 限速 2MBps
	//})
	//
	//cred, err := drive.AddCredential(aliyundrive.NewCredential(&aliyundrive.Credential{
	//	RefreshToken: "45e308a8e80a41028d161421d50a002c",
	//}))
	////drive.RefreshToken(cred)
	////spew.Dump(cred)
	//
	//root, err := drive.GetByPath(cred, "test")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//file, err := os.OpenFile("0M~0@S}F3CCG%E_[`[FVF{C.jpg", os.O_RDONLY, 0)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//defer file.Close()
	//
	//f_info, _ := file.Stat()
	//fileRapid, rapid, err := drive.UploadFileRapid(cred, &aliyundrive.UploadFileRapidOptions{
	//	UploadFileOptions: aliyundrive.UploadFileOptions{
	//		Name:         f_info.Name(),
	//		Size:         f_info.Size(),
	//		ParentFileId: root.FileId,
	//	},
	//	File: file,
	//})

	//fileurl, err := drive.GetDownloadURL(cred, fileRapid.FileId)
	//
	//spew.Dump(fileRapid)
	//spew.Dump(rapid)
	//spew.Dump(fileurl)

	//url := "https://bj29.cn-beijing.data.alicloudccp.com/U6CKCoRO%2F60651%2F6280f2666239958cc1634bb9bdab3164f06344a4%2F6280f266aec7b102c7de4d8b86757824028bf53f?di=bj29&dr=60651&f=6280f2666239958cc1634bb9bdab3164f06344a4&response-content-disposition=attachment%3B%20filename%2A%3DUTF-8%27%27name%25288%2529&u=654f5ad9325c4b5dbe152ffed488ae2d&x-oss-access-key-id=LTAI5t8sJLSvMtxoes9pGyTv&x-oss-additional-headers=referer&x-oss-expires=1652632232&x-oss-signature=sWvqiAxq4O%2F4qL0R6yCPDP6isYN55EhXTYBAg8MRb6k%3D&x-oss-signature-version=OSS2"
	//r, err := http.NewRequest("GET", url, nil)
	//if err != nil {
	//	panic(err)
	//}
	//r.Header.Set("Referer", " https://www.aliyundrive.com/")
	//
	//c := http.Client{}
	//re, err := c.Do(r)
	//b, _ := ioutil.ReadAll(re.Body)
	//f, _ := os.Create("name.jpg")
	//f.Write(b)
	//f.Close()
	// ...

	//msg, err := drive.RemoveFile(cred, root.FileId)
	//spew.Dump(msg)
}
