package testing

import (
	"fmt"
	"io"
	"os"
)

func TestingCopy() {
	destName := "/mnt/e/copied.txt"
	srcName := "/mnt/e/test.txt"
	total, err := copyFile(destName, srcName)
	// total, err := copyFile2(destName, srcName)
	if err != nil {
		fmt.Println("复制文件过程中产生错误，，", err)
	} else {
		fmt.Println("文件拷贝成功，总量是：", total)
	}
}

//实现文件的复制，返回值拷贝的数据总量，以及拷贝过程中产生的错误
func copyFile(destName, srcName string) (int, error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	destFile, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()
	//复制数据
	bs := make([]byte, 1024)
	count := 0 //每次实际读入数据量
	total := 0 // 用于统计读取的数据总量
	for {
		count, err = srcFile.Read(bs)
		if err != nil && err != io.EOF { //EOF
			return total, err
		} else if err == io.EOF {
			fmt.Println("已经到达文件末尾。。")
			break
		}
		destFile.Write(bs[:count])
		total += count
	}
	return total, nil

}

func copyFile2(destName, srcName string) (int64, error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	destFile, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()
	return io.Copy(destFile, srcFile)
}
