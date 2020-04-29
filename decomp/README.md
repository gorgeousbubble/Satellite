# Decompress package interfaces
The Decompress package function interfaces description.

## Introduction
Decompress package is a functional package. It can decompress an archive to files. Related algorithms are mainly from the Go package. Now support tar, tar.gz, zip, gzip, bzip2, etc.

## Feature of package
The package is mainly used for decompress achive which has been operated by 'decomp' package.

#### Achieve files by decompress
  * The purpose of decompress is to restore the achive file
  * Support compress various algorithms which has been operated by 'decomp' package, like gzip, tar, tar.gz, zip, bzip2, etc.
  * GUI support
  * Simple and useful

## Usage of interfaces
Although many function was expose to external, we only need call function 'Decompress(...)' to realize our function. You can also use other functions according to you demand.
```batch
// DeCompress function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.tar.gz' or '../test/data/file.tar.gz'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// algorithm now support 'tar', 'tar.gz', 'zip', you can send both up case and low case
// return err indicate the success or failure function execute
func DeCompress(src string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = DeCompressTar(src, dest)
	case "tar.gz":
		err = DeCompressTarGz(src, dest)
	case "zip":
		err = DeCompressZip(src, dest)
	default:
		s := fmt.Sprint("Undefined decompress algorithm.")
		err = errors.New(s)
	}
	return err
}
```