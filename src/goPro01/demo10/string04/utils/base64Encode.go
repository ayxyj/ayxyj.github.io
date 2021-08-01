package utils

func EncodeBase64(plain string) (cipher string) {
	const key string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var (
		length        int = len(plain)
		loop_time     int = length / 3
		left          int = length % 3
		bytes_catnate int32
		index         int
	)
	for i := 0; i < loop_time; i++ {
		bytes_catnate = int32(plain[index])<<16 + int32(plain[index+1])<<8 + int32(plain[index+2])
		index += 3
		cipher += string(key[(bytes_catnate>>18)&0x3F])
		cipher += string(key[(bytes_catnate>>12)&0x3F])
		cipher += string(key[(bytes_catnate>>6)&0x3F])
		cipher += string(key[bytes_catnate&0x3F])
	}
	if left == 1 {
		bytes_catnate = int32(plain[index]) << 4
		cipher += string(key[(bytes_catnate>>6)&0x3F])
		cipher += string(key[bytes_catnate&0x3F])
		cipher += "=="
	}
	if left == 2 {
		bytes_catnate = (int32(plain[index])<<8 + int32(plain[index+1])) << 2
		cipher += string(key[(bytes_catnate>>12)&0x3F])
		cipher += string(key[(bytes_catnate>>6)&0x3F])
		cipher += string(key[bytes_catnate&0x3F])
		cipher += "="
	}

	return cipher
}

/**
Base64编码：
所谓Base64，就是说选出64个字符----小写字母a-z、大写字母A-Z、数字0-9、符号"+"、"/"（再加上作为垫字的"="，实际上是65个字符）
作为一个基本字符集。然后，其他所有符号都转换成这个字符集中的字符。

第一步，将每三个字节作为一组，一共是24个二进制位。
第二步，将这24个二进制位分为四组，每个组有6个二进制位。
第三步，在每组前面加两个00，扩展成32个二进制位，即四个字节。
第四步，根据表，得到扩展后的每个字节的对应符号，这就是Base64的编码值。

因为，Base64将三个字节转化成四个字节，因此Base64编码后的文本，会比原文本大出三分之一左右。

举一个具体的实例，演示英语单词Man如何转成Base64编码。



Text content	M	a	n
ASCII			77	97	110
Bit pattern		0	1	0	0	1	1	0	1	0	1	1	0	0	0	0	1	0	1	1	0	1	1	1	0
Index			19	22	5	46
Base64-Encoded	T	W	F	u

第一步，"M"、"a"、"n"的ASCII值分别是77、97、110，对应的二进制值是01001101、01100001、01101110，将它们连成一个24位的二进制字符串010011010110000101101110。
第二步，将这个24位的二进制字符串分成4组，每组6个二进制位：010011、010110、000101、101110。
第三步，在每组前面加两个00，扩展成32个二进制位，即四个字节：00010011、00010110、00000101、00101110。它们的十进制值分别是19、22、5、46。
第四步，根据上表，得到每个值对应Base64编码，即T、W、F、u。
因此，Man的Base64编码就是TWFu。

如果字节数不足三，则这样处理：
a）二个字节的情况：将这二个字节的一共16个二进制位，按照上面的规则，转成三组，最后一组除了前面加两个0以外，后面也要加两个0。
这样得到一个三位的Base64编码，再在末尾补上一个"="号。
比如，"Ma"这个字符串是两个字节，可以转化成三组00010011、00010110、00010000以后，
对应Base64值分别为T、W、E，再补上一个"="号，因此"Ma"的Base64编码就是TWE=。

b）一个字节的情况：将这一个字节的8个二进制位，按照上面的规则转成二组，最后一组除了前面加二个0以外，后面再加4个0。
这样得到一个二位的Base64编码，再在末尾补上两个"="号。
比如，"M"这个字母是一个字节，可以转化为二组00010011、00010000，
对应的Base64值分别为T、Q，再补上二个"="号，因此"M"的Base64编码就是TQ==。
*/