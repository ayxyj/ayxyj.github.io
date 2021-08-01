package main
 
import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/hex"
    "fmt"
    "hash"
 
    "golang.org/x/crypto/md4"
    "golang.org/x/crypto/ripemd160"
)
 
func main() {
    var val,tp string = "md5" , "123456"
    var flag bool = false
    fmt.Println("请输入加密类型：")
    fmt.Scanln(&tp)
    fmt.Println("请输入加密数据：")
    fmt.Scanln(&val)
    fmt.Println("数据类型是否为十六进制：")
    fmt.Scanln(&flag)
    res := HASH(val, tp , flag)
    fmt.Println("采用"+tp+"加密后的结果为：")
    fmt.Println(res)
}
 
// HASH HASH
func HASH(text string, hashType string, isHex bool) string {
    var hashInstance hash.Hash
    switch hashType {
    case "md4":
        hashInstance = md4.New()
    case "md5":
        hashInstance = md5.New()
    case "sha1":
        hashInstance = sha1.New()
    case "sha256":
        hashInstance = sha256.New()
    case "sha512":
        hashInstance = sha512.New()
    case "ripemd160":
        hashInstance = ripemd160.New()
    }
    if isHex {
        arr, _ := hex.DecodeString(text)
        hashInstance.Write(arr)
    } else {
        hashInstance.Write([]byte(text))
    }
 
    bytes := hashInstance.Sum(nil)
    return fmt.Sprintf("%x", bytes)
}
 
// MD4 MD4
func MD4(text string, isHex bool) string {
    var hashInstance hash.Hash
    hashInstance = md4.New()
    if isHex {
        arr, _ := hex.DecodeString(text)
        fmt.Println(arr)
        hashInstance.Write(arr)
    } else {
        hashInstance.Write([]byte(text))
    }
 
    bytes := hashInstance.Sum(nil)
    return fmt.Sprintf("%x", bytes)
}
 
// MD5 MD5
func MD5(text string, isHex bool) string {
    var hashInstance hash.Hash
    hashInstance = md5.New()
    if isHex {
        arr, _ := hex.DecodeString(text)
        fmt.Println(arr)
        hashInstance.Write(arr)
    } else {
        hashInstance.Write([]byte(text))
    }
 
    bytes := hashInstance.Sum(nil)
    return fmt.Sprintf("%x", bytes)
}
 
 
 
//实现双哈希算法
func sha256Double(text string, isHex bool) []byte {
    hashInstance := sha256.New()
    if isHex {
        arr, _ := hex.DecodeString(text)
        hashInstance.Write(arr)
    } else {
        hashInstance.Write([]byte(text))
    }
    bytes := hashInstance.Sum(nil)
    hashInstance.Reset()
    hashInstance.Write(bytes)
    bytes = hashInstance.Sum(nil)
    return bytes
}
func sha256DoubleString(text string, isHex bool) string {
    bytes := sha256Double(text, isHex)
    return fmt.Sprintf("%x", bytes)
}