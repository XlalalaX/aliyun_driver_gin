package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jakeslee/aliyundrive"
	"log"

	"net/http"
	"time"

	"main/aliyun_services/global"
	"main/aliyun_services/util"
)

func Get_URL(ctx *gin.Context) {
	rand_string := ctx.Query("rand_string")

	re, err := global.Rdb.Get(rand_string).Result()
	if err != nil || re == "" {
		ctx.JSON(400, gin.H{"msg": "取件码失效或者文件不存在"})
		ctx.Abort()
		return
	}

	f_id, err := global.Aliyun.SearchNameInFolder(global.User, re, global.Sava_path_Id)
	if err != nil || len(f_id.Items) == 0 {
		ctx.JSON(500, gin.H{"err": err.Error()})
		ctx.Abort()
		return
	}

	resp, err := global.Aliyun.GetFolderFiles(global.User, &aliyundrive.FolderFilesOptions{FolderFileId: f_id.Items[0].FileId})
	if err != nil {
		ctx.JSON(400, gin.H{"err": err})
		ctx.Abort()
		return
	}
	ctx.JSON(200, gin.H{"resp": resp})
}

func Seach_info(ctx *gin.Context) {
	path := ctx.Query("path")

	re, err := global.Aliyun.SearchNameInFolder(global.User, path, global.Sava_path_Id)
	if err != nil {
		ctx.JSON(500, gin.H{"err": err.Error()})
		ctx.Abort()
	}
	if global.Debug {
		ctx.JSON(200, gin.H{"re": re})
	} else {
		ctx.JSON(200, gin.H{"f_id": re.Items[0].FileId})
	}

}

func Upload_file(ctx *gin.Context) {
	// 多文件
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(500, "上传文件到gin失败")
	}

	r_time_s := ctx.PostForm("r_time")
	r_time := time.Hour * 24
	if r_time_s != "" {
		r_time, _ = time.ParseDuration(r_time_s)
	}

	files := form.File["file"]

	uid := uuid.New()
	files_path, err := global.Aliyun.CreateDirectory(global.User, global.Sava_path_Id, uid.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"msg": "上传出错"})
		ctx.Abort()
		return
	}
	rand_string := ""
	for {
		rand_string = util.Get_rand_string()
		resp, err := global.Rdb.SetNX(rand_string, uid.String(), r_time).Result()
		fmt.Println("rand_string:", rand_string)
		if err != nil || !resp {
			fmt.Println("随机字符串出错:" + rand_string)
		} else {
			break
		}
	}
	//过期时要用到的，存到DB2
	global.Rdb_2.Set(uid.String(), "", r_time).Result()

	success_files := []string{}
	unsuccess_files := []string{}
	for _, file := range files {
		log.Println(file.Filename)
		read, err := file.Open()
		if err != nil {
			unsuccess_files = append(unsuccess_files, file.Filename)
			continue
		}
		_, err = global.Aliyun.UploadFile(global.User, &aliyundrive.UploadFileOptions{
			Name:         file.Filename,
			Size:         file.Size,
			ParentFileId: files_path.FileId,
			ProgressStart: func(info *aliyundrive.ProgressInfo) {
				fmt.Println("start:" + file.Filename)
			},
			ProgressDone: func(info *aliyundrive.ProgressInfo) {
				fmt.Println("end:" + file.Filename)
			},
			Reader: read,
		})
		if err != nil {
			unsuccess_files = append(unsuccess_files, file.Filename)
			continue
		}
		success_files = append(success_files, file.Filename)
		//spew.Dump(re)
		//// 上传文件到指定的路径
		//ctx.SaveUploadedFile(file, ".\\新建文件夹\\"+file.Filename)
	}
	ctx.JSON(http.StatusOK, gin.H{"取件码": rand_string, "dir": files_path, "file_path": uid.String(), "上传成功文件": success_files, "上传失败文件": unsuccess_files})
}

func Delete_file(ctx *gin.Context) {
	rand_string := ctx.PostForm("rand_string")
	if len(rand_string) != 6 {
		ctx.JSON(400, gin.H{"msg": "取件码错误"})
		ctx.Abort()
		return
	}

	re, err := global.Rdb.Del(rand_string).Result()
	if err != nil || re == 0 {
		ctx.JSON(400, gin.H{"msg": "取件码错误或者已失效"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"msg": "删除成功"})
}
