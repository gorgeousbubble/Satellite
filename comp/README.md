# Compress package interfaces
The Compress package function interfaces description.

## Introduction
Compress package is a functional package. It can compress files to an archive. Related algorithms are mainly from the Go package. Now support tar, tar.gz, zip, gzip, etc.

## Feature of package
The package is mainly used for compress files which has been operated by 'comp' package.

#### Achieve files by compress
  * The purpose of compress is to achieve files and decrease it's volume 
  * Support compress various algorithms which has been operated by 'comp' package, like gzip, tar, tar.gz, zip, zlib, etc.
  * GUI support
  * Simple and useful

## Usage of interfaces
Although many function was expose to external, we only need call function 'Compress(...)' to realize our function. You can also use other functions according to you demand.
```batch
// Compress function
// input src file list, output dest file path and algorithm which used in pack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.tar.gz' or '../test/data/package.tar.gz'
// algorithm now support 'tar', 'tar.gz', 'zip', you can send both up case and low case
// return err indicate the success or failure function execute
func Compress(src []string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = CompressTar(src, dest)
	case "tar.gz":
		err = CompressTarGz(src, dest)
	case "zip":
		err = CompressZip(src, dest)
	default:
		s := fmt.Sprint("Undefined compress algorithm.")
		err = errors.New(s)
	}
	return err
}
```
