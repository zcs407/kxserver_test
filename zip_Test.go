package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//oldZipPath := "./oldSource.zip"
	oldZipDir := "./oldTestDir"
	//newZipPath := "./newSource.zip"
	//newZipDir := "./newTestDir"
	//err:=fileUnzip(oldZipPath, oldZipDir)
	//if err != nil {
	//	log.Println("can't unzip this file :",oldZipPath," error: ",err)
	//	return
	//}
	//err=fileUnzip(newZipPath, newZipDir)
	//if err != nil {
	//	log.Println("can't unzip this file :",newZipPath," error: ",err)
	//	return
	//}

	//pwd,_ := os.Getwd()
	//获取文件或目录相关信息
	//fileInfoList, err := ioutil.ReadDir(newZipDir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := range fileInfoList {
	//	srcFile := newZipDir + "/" + fileInfoList[i].Name()
	//	dstFile := oldZipDir + "/" + fileInfoList[i].Name()
	//	log.Println("开始替换文件,将new file:", srcFile, " 替换掉: ", dstFile)
	//	_, err = fileCopy(srcFile, dstFile)
	//	if err != nil {
	//		log.Println("开始替换文件,将new file:", srcFile, " 替换掉: ", dstFile, " error: ", err)
	//		return
	//	}
	//}
	err := fileZip(oldZipDir, "./newSource001.zip")
	if err != nil {
		log.Println("无法压缩为zip文件")
		return
	}
	//filepathNames,err := filepath.Glob(filepath.Join(pwd,"*"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := range filepathNames {
	//	log.Println(filepathNames[i]) //打印path
	//}
	//
	//filepath.Walk(pwd,func(path string, info os.FileInfo, err error) error{
	//	log.Println(path) //打印path信息
	//	log.Println(info.Name()) //打印文件或目录名
	//	return nil
	//})

}
func fileZip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		//header.Name = path
		if info.IsDir() {
			header.Name += "/"
			header.Method = zip.Deflate
			log.Println(" head name: ", header.Name)
		} else {
			header.Method = zip.Deflate
			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				_, err = io.Copy(writer, file)
			}
			return err
		}
		return err
	})
	return err
}

//一个是输入源zip文件 另一个是输出路径
func fileUnzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		//fmt.Pringln("DIR:",path)
		return nil
	}

	//文件类型需要进行过滤

	//buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	//content := string(buf)

	//替换
	//newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	//ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}

func fileCopy(srcFileFullPath, dstFileFullPath string) (int64, error) {
	sourceFileStat, err := os.Stat(srcFileFullPath)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", srcFileFullPath)
	}

	source, err := os.Open(srcFileFullPath)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dstFileFullPath)
	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main2() {
	type aaa struct {
		Name string
		Age  int
	}
	bbb := aaa{
		Name: "123",
		Age:  100,
	}
	log.Println(bbb)
}
