# Pack package interfaces
The Pack package function interfaces description.

## Introduction
Pack package is a functional package. It can encrypt files or data and generate check code with various algorithms. Related algorithms are mainly from the Go package.

## Feature of package
The package has two kinds of feature, one is pack or encrypt files or data to protect files or data security, the other is generate check code so that we can use this function to realize check function.

#### Pack or Encrypt files or data protect its security
  * Can pack or encrypt any type of files or data
  * Support encrypt various algorithms, like AES, DES, 3DES, RSA, BASE64, etc.
  * Support HTTP and HTTPS to call this function
  * You can know the process when pack or encrypt
  * Simple and useful
  
#### Generate check code so that we can check use it to realize check
  * Generate different kind of algorithms, such as MD5, SHA1, SHA256, SHA512
  * Support Gob stream and file generate, which used only by go application
  
## External definitions
There are some constant defined in file 'global/const.go', mainly about encrypt slice buffer size. You can redefine it in 'pack' package it you want to use this package independence.

## Usage of interfaces
Although many function was expose to external, we only need call function 'Pack(...)' to realize our function. You can also use other functions according to you demand.
```batch
// Pack function
// input src file list, output dest file path and algorithm which used in pack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// algorithm now support 'AES', 'DES', '3DES', 'RSA' and 'BASE64', you can send both up case and low case
// return err indicate the success or failure function execute
func Pack(src []string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "AES", "aes":
		err = PackAES(src, dest)
	case "DES", "des":
		err = PackDES(src, dest)
	case "3DES", "3des":
		err = Pack3DES(src, dest)
	case "RSA", "rsa":
		err = PackRSA(src, dest)
	case "BASE64", "base64":
		err = PackBase64(src, dest)
	default:
		s := fmt.Sprint("Undefined pack algorithm.")
		err = errors.New(s)
	}
	return err
}
```
You can call this function to achieve pack or encrypt like this:
```batch
src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
dest := "../test/data/pack/file_aes.txt"
algorithm := "AES"
err := Pack(src, dest, algorithm)
if err != nil {
    t.Fatal("Error Pack:", err)
}
```

Certainly, function 'Pack' mainly do the pack or encrypt. If you also want to check the process of pack or encrypt, you can use function 'WorkCalculate' to get the work value pack now. It show how much work remain.
```batch
// WorkCalculate function
// input src file list, algorithm which used in pack and output work value, return error info
// this function will called by calculate work
// algorithm now support 'AES', 'DES', '3DES', 'RSA' and 'BASE64', you can send both up case and low case
// work value is total work force that will be done
// return err indicate the success or failure function execute
func WorkCalculate(src []string, algorithm string, work *int64) (err error) {
	switch algorithm {
	case "AES", "aes":
		*work, err = PackAESWorkCalculate(src)
	case "DES", "des":
		*work, err = PackDESWorkCalculate(src)
	case "3DES", "3des":
		*work, err = PackDESWorkCalculate(src)
	case "RSA", "rsa":
		*work, err = PackRSAWorkCalculate(src)
	case "BASE64", "base64":
		*work, err = PackBase64WorkCalculate(src)
	default:
		s := fmt.Sprint("Undefined pack algorithm.")
		err = errors.New(s)
	}
	return err
}
```

You can use this function like this:
```batch
var work int64
src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
err := WorkCalculate(src, "aes", &work)
if err != nil {
    t.Fatal("Error Work Calculate:", err)
}
```

For other functions, you can also call them from external. Usage is similar to the function which description before.