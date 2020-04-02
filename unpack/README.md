# Unpack package interfaces
The Unpack package function interfaces description.

## Introduction
Unpack package is a functional package. It can decrypt files or data. Related algorithms are mainly from the Go package.

## Feature of package
The package is mainly used for decrypt data or unpack file which has been operated by 'pack' package.

#### Unpack or Decrypt files or data protect its security
  * Can unpack or decrypt any type of files or data
  * Support decrypt various algorithms which has been operated by 'pack' package, like AES, DES, 3DES, RSA, BASE64, etc.
  * Support HTTP and HTTPS to call this function
  * You can know the process when unpack or decrypt
  * Simple and useful

## External definitions
There are some constant defined in file 'global/const.go', mainly about decrypt slice buffer size. You can redefine it in 'unpack' package it you want to use this package independence.

## Usage of interfaces
Although many function was expose to external, we only need call function 'Unpack(...)' to realize our function. You can also use other functions according to you demand.
```batch
// Unpack function
// input src file which was packed or encrypted, output dest file path, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// algorithm now support 'AES', 'DES', '3DES', 'RSA' and 'BASE64', but you don't need to care it~
// return err indicate the success or failure function execute
func Unpack(src string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAES(src, dest)
	case "DES", "des":
		err = UnpackDES(src, dest)
	case "3DES", "3des":
		err = Unpack3DES(src, dest)
	case "RSA", "rsa":
		err = UnpackRSA(src, dest)
	case "BASE64", "base64":
		err = UnpackBase64(src, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}
```
You can call this function to achieve unpack or decrypt like this:
```batch
src := "../test/data/unpack/file_aes.txt"
dest := "../test/data/unpack/"
err := Unpack(src, dest)
if err != nil {
	t.Fatal("Error Unpack:", err)
}
```

Certainly, function 'Unpack' mainly do the unpack or decrypt. There are some other function in package which you can call it for different situation.
If the package file is too large that the program may generate amount of go-routine. Too much go-routine may not a good thing for speed up unpack, it will squeeze memory and cause program break down. You can call this function to restrict go-routine when unpack package.
```batch
// UnpackConfine function
// unpack file with restrict goroutine(if we do not restrict goroutine, memory will soon be occupied)
// you can adjust confine file and confine buffer when you need change
// other function is same as 'Unpack'
func UnpackConfine(src string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAESConfine(src, dest)
	case "DES", "des":
		err = UnpackDESConfine(src, dest)
	case "3DES", "3des":
		err = Unpack3DESConfine(src, dest)
	case "RSA", "rsa":
		err = UnpackRSAConfine(src, dest)
	case "BASE64", "base64":
		err = UnpackBase64Confine(src, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}
```