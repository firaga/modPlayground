package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	// Register image formats
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	// ioutil.WriteFile("texture1.jpg", []byte(texture1), 0666)
	// ioutil.WriteFile("texture2.png", []byte(texture2), 0666)

	log.Println("Texture 1")
	log.Println("size =", len(texture1))
	a := []byte(texture1)
	fmt.Println(a)
	r1 := bytes.NewReader(a)
	img1, format1, err1 := image.Decode(r1)
	log.Println("format =", format1)
	log.Println("err =", err1)
	log.Println()

	log.Println("Texture 2")
	log.Println("size =", len(texture2))
	r2 := bytes.NewReader([]byte(texture2))
	img2, format2, err2 := image.Decode(r2)
	log.Println("format =", format2)
	log.Println("err =", err2)

	_, _ = img1, img2
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const texture1 = "" +
	"\xff\xd8\xff\xe0\x00\x10\x4a\x46\x49\x46\x00\x01\x01\x00\x00\x01" +
	"\x00\x01\x00\x00\xff\xdb\x00\x43\x00\x05\x03\x04\x04\x04\x03\x05" +
	"\x04\x04\x04\x05\x05\x05\x06\x07\x0c\x08\x07\x07\x07\x07\x0f\x0b" +
	"\x0b\x09\x0c\x11\x0f\x12\x12\x11\x0f\x11\x11\x13\x16\x1c\x17\x13" +
	"\x14\x1a\x15\x11\x11\x18\x21\x18\x1a\x1d\x1d\x1f\x1f\x1f\x13\x17" +
	"\x22\x24\x22\x1e\x24\x1c\x1e\x1f\x1e\xff\xdb\x00\x43\x01\x05\x05" +
	"\x05\x07\x06\x07\x0e\x08\x08\x0e\x1e\x14\x11\x14\x1e\x1e\x1e\x1e" +
	"\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e" +
	"\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e" +
	"\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\xff\xc0" +
	"\x00\x11\x08\x00\x10\x00\x10\x03\x01\x22\x00\x02\x11\x01\x03\x11" +
	"\x01\xff\xc4\x00\x1f\x00\x00\x01\x05\x01\x01\x01\x01\x01\x01\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09" +
	"\x0a\x0b\xff\xc4\x00\xb5\x10\x00\x02\x01\x03\x03\x02\x04\x03\x05" +
	"\x05\x04\x04\x00\x00\x01\x7d\x01\x02\x03\x00\x04\x11\x05\x12\x21" +
	"\x31\x41\x06\x13\x51\x61\x07\x22\x71\x14\x32\x81\x91\xa1\x08\x23" +
	"\x42\xb1\xc1\x15\x52\xd1\xf0\x24\x33\x62\x72\x82\x09\x0a\x16\x17" +
	"\x18\x19\x1a\x25\x26\x27\x28\x29\x2a\x34\x35\x36\x37\x38\x39\x3a" +
	"\x43\x44\x45\x46\x47\x48\x49\x4a\x53\x54\x55\x56\x57\x58\x59\x5a" +
	"\x63\x64\x65\x66\x67\x68\x69\x6a\x73\x74\x75\x76\x77\x78\x79\x7a" +
	"\x83\x84\x85\x86\x87\x88\x89\x8a\x92\x93\x94\x95\x96\x97\x98\x99" +
	"\x9a\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xb2\xb3\xb4\xb5\xb6\xb7" +
	"\xb8\xb9\xba\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xd2\xd3\xd4\xd5" +
	"\xd6\xd7\xd8\xd9\xda\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xf1" +
	"\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xff\xc4\x00\x1f\x01\x00\x03" +
	"\x01\x01\x01\x01\x01\x01\x01\x01\x01\x00\x00\x00\x00\x00\x00\x01" +
	"\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\xff\xc4\x00\xb5\x11\x00" +
	"\x02\x01\x02\x04\x04\x03\x04\x07\x05\x04\x04\x00\x01\x02\x77\x00" +
	"\x01\x02\x03\x11\x04\x05\x21\x31\x06\x12\x41\x51\x07\x61\x71\x13" +
	"\x22\x32\x81\x08\x14\x42\x91\xa1\xb1\xc1\x09\x23\x33\x52\xf0\x15" +
	"\x62\x72\xd1\x0a\x16\x24\x34\xe1\x25\xf1\x17\x18\x19\x1a\x26\x27" +
	"\x28\x29\x2a\x35\x36\x37\x38\x39\x3a\x43\x44\x45\x46\x47\x48\x49" +
	"\x4a\x53\x54\x55\x56\x57\x58\x59\x5a\x63\x64\x65\x66\x67\x68\x69" +
	"\x6a\x73\x74\x75\x76\x77\x78\x79\x7a\x82\x83\x84\x85\x86\x87\x88" +
	"\x89\x8a\x92\x93\x94\x95\x96\x97\x98\x99\x9a\xa2\xa3\xa4\xa5\xa6" +
	"\xa7\xa8\xa9\xaa\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xc2\xc3\xc4" +
	"\xc5\xc6\xc7\xc8\xc9\xca\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xe2" +
	"\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9" +
	"\xfa\xff\xda\x00\x0c\x03\x01\x00\x02\x11\x03\x11\x00\x3f\x00\xec" +
	"\xbe\x34\xf8\xa7\x56\xd7\x75\xd6\xd3\x36\xdd\x69\x49\xa7\x5c\x89" +
	"\x2d\xf1\xf2\x4b\x1c\xa3\xa4\xb9\xf5\xf4\xed\x8a\xee\xbe\x12\x7c" +
	"\x56\x5d\x62\xf6\xdf\xc2\xfe\x28\x2b\x6b\xaf\x32\x1f\x22\x6e\x04" +
	"\x57\xc0\x77\x4f\x47\xf5\x4f\xca\xba\x8f\x89\x3e\x06\xb4\xf1\x4d" +
	"\x87\x9f\x10\x48\x35\x38\x94\xf9\x33\x63\xef\x0f\xee\x3f\xb7\xf2" +
	"\xaf\x2b\xf0\x2f\xc3\x5b\xad\x6b\xc4\xa5\x35\xfb\x49\x20\xb4\xd3" +
	"\x24\x0e\x47\x47\x32\x83\x91\x83\xdb\xb1\xe2\xb3\xbb\x40\x7f\xff" +
	"\xd9"

const texture2 = "" +
	"\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52" +
	"\x00\x00\x00\x80\x00\x00\x00\x80\x08\x00\x00\x00\x00\xe6\x55\x3e" +
	"\x17\x00\x00\x00\x02\x74\x52\x4e\x53\x00\x65\x4b\x4b\x58\xef\x00" +
	"\x00\x03\xc0\x49\x44\x41\x54\x78\x01\xed\x95\x4d\x6b\x14\x41\x10" +
	"\x86\x7b\x7a\x23\x1e\xcc\x45\x08\x5e\x3d\x08\x22\xc4\x8b\xb3\xff" +
	"\x40\xbc\x2b\x5e\xe2\x41\x8c\xd2\xec\xcf\x2a\x24\xbf\xa8\x31\x47" +
	"\x25\x07\x15\x04\x41\x30\x88\x22\xd1\x40\x9c\xdd\xf9\xec\x99\xe9" +
	"\x9e\xae\xaa\xfe\x10\x74\x2f\x3b\x9a\x9d\x7e\xde\x7e\xaa\xba\xba" +
	"\xd8\x08\xce\x07\xea\x2f\x45\x5f\x41\xb2\xf8\xaf\xd9\x7c\xb1\x97" +
	"\x77\xff\x3c\x03\x0d\xff\x86\xc8\x14\xa0\xe1\x8b\x67\x99\x02\xb4" +
	"\x7c\x56\x01\x18\x01\x80\xc7\x65\x07\xe8\xf8\x2a\x4f\x00\x08\x66" +
	"\x42\xb2\xfd\x33\x13\xac\xd6\xec\xfa\xeb\x4f\x77\x13\x1b\x18\xef" +
	"\xf9\x33\x24\x35\x00\xba\x79\x50\xe5\xd9\x45\x2b\xa1\x4c\x17\x60" +
	"\xd8\xfe\x87\xe7\xe7\x6d\x82\xb3\x43\x62\x00\xc9\x3a\x7e\x0f\xbb" +
	"\x43\xf8\x0d\xd2\x04\xe8\xc6\x9f\x9a\x8c\x01\x62\x02\x5c\x09\x66" +
	"\xc6\x6f\xf9\xe1\x67\xf3\xa4\x49\x9d\x20\xd9\xe3\xff\x09\x4f\x42" +
	"\xb1\xf1\xfe\xe9\xc9\xa5\x6d\xfa\x76\x7f\x21\x4c\x66\x7f\x03\x60" +
	"\xe5\x8b\x63\x86\x04\xef\x1e\x70\xdf\x3e\x65\x3b\x1c\x84\x7e\xf3" +
	"\x00\x15\xc0\xb7\x04\x2d\x7f\xef\x78\x79\x3e\xaa\x08\x01\xc0\x63" +
	"\x6d\x5a\x27\xc8\x60\x7c\x62\x27\x78\xf5\x00\x78\x6e\x6c\xd0\x09" +
	"\xde\x37\xa4\x4f\x09\x00\x21\x16\xdd\x09\x32\x2c\x1f\x3f\x9a\x97" +
	"\x0d\x00\xb6\xb1\x70\x12\x64\x70\x3e\x52\xc2\x92\x01\x20\x8d\x58" +
	"\x84\x04\xb7\x01\xa0\xf1\x31\x12\x9c\x06\x80\x7c\xc5\xf8\x4b\x90" +
	"\x91\xf8\xde\x12\x1c\x06\x80\x81\x37\xc9\x37\x9f\x52\x0c\xb0\xf9" +
	"\xfd\x9b\x5f\x81\x60\x80\xcf\x37\x24\xdc\x7e\x84\x34\x10\x86\xdf" +
	"\xbf\xfe\x1e\x70\x01\x02\xf1\xab\x05\x6e\x2d\xf4\xe2\x6c\x09\x58" +
	"\xed\x6f\x5f\xed\xfa\x73\x4f\x03\x61\xf9\x42\xed\x37\x0f\x17\xe0" +
	"\x17\x20\x30\x5f\x88\x23\xe5\x28\xc3\x6a\x1d\x9d\x5f\x7d\xca\xd3" +
	"\xab\xfa\x41\xeb\x72\xc9\x40\x0c\xbe\x10\x2f\xad\x12\xc6\x06\xe2" +
	"\xf0\xb7\x12\x74\xf3\xa0\x3f\xde\x1b\xfe\xff\xe8\x14\xb4\xfc\x83" +
	"\xc7\x81\xf9\xd6\xcd\x19\x01\xa2\x6d\xdf\xb1\xbe\x4c\xc8\xef\xd7" +
	"\x05\x98\x0b\x10\x9d\x5f\xad\x7c\x30\x81\xf5\x25\x48\xc0\x37\x30" +
	"\xfb\x47\xa6\x81\x44\xfc\x7e\xf9\xef\x60\x04\x48\xc5\xaf\x00\xe6" +
	"\x48\x90\xa9\xf9\x46\x2f\x8a\x76\x10\x25\xe5\x57\x43\xe9\xdd\xaf" +
	"\xfa\x41\x7f\xb9\xb3\x6b\xc2\xc4\x7c\xf3\x14\x48\xa1\x32\xf0\xfb" +
	"\x04\xc5\x60\x10\xa4\xe4\x0b\x71\x72\x39\x0a\x20\x5f\x25\xe5\x8f" +
	"\x4e\x41\xf5\x49\xce\xaf\x8d\xcb\x3c\xfe\x7b\xa6\xcc\xc0\x35\x3e" +
	"\xff\x76\x00\xc8\x1c\x00\x32\x1b\x30\x6f\xc3\x6c\x9f\xff\x01\xfe" +
	"\x9a\x00\x39\x06\x61\x6e\x03\x2a\x77\x80\x3a\x41\x91\xf5\x36\x82" +
	"\xa1\x01\xc8\xe3\x60\x55\x74\xff\xd2\xba\x4c\x9e\x40\x0f\x4a\x90" +
	"\xa3\x0e\x6d\x09\x94\xca\x53\x87\x2d\x6d\x6b\x40\x19\xe4\x74\x12" +
	"\x76\xcc\x2a\x80\x1a\x6d\x3e\x55\x82\x1a\x58\x6c\xc4\xf4\x1c\x24" +
	"\x89\xd0\xd0\x86\x3d\x98\xb4\x13\x5a\xc6\xd0\x40\x4a\x09\x0d\x48" +
	"\x8d\x03\xa4\x8a\xd0\xf1\xa7\x77\x41\x92\x3a\xf4\xfc\xa9\x81\x04" +
	"\x12\x60\xb8\xfc\xdc\x6d\x18\x59\x82\xc1\x9f\x35\x10\x57\x82\xc9" +
	"\x9f\x35\x30\x1c\xce\xc1\x25\x8c\xf8\xb6\x00\xfd\x0f\x20\x6c\x84" +
	"\x31\xdf\x1e\x20\x8e\x84\x09\xdf\x11\x20\x86\x84\x29\xdf\x19\x20" +
	"\xb8\x84\x19\xbe\x58\xad\x9d\xaf\x94\xa7\x57\xbb\x6f\xad\xcb\x38" +
	"\x7c\xdb\x31\x5c\x7a\x8d\x39\xfe\x50\x01\x02\xcd\x04\xdb\x3e\xe4" +
	"\xf2\xab\x41\x3a\xc1\xea\xd1\x23\x40\x88\xe3\x60\xaf\xa3\x57\x00" +
	"\xb6\x04\x47\x1f\xf9\x05\x60\x4a\x70\xf5\xb1\x47\x13\x72\x8f\x83" +
	"\xbb\x89\xa5\xff\x42\xc4\x32\xc0\xb8\x8c\x64\x03\x24\x09\x8b\x67" +
	"\x58\xa2\x76\x83\x96\xb0\x3c\x43\x70\x06\xb0\x12\x3c\x66\x18\xce" +
	"\x00\x52\xc2\x42\xf9\x69\x06\x10\x12\xbc\x7e\x87\x36\xe0\x2f\xc1" +
	"\x2f\x27\xc5\x40\xbf\xf8\xb5\x17\xec\x2b\x8c\x62\xa0\x5f\xf5\x37" +
	"\xb0\xaf\x50\x62\x80\xa5\x32\xf8\x5f\xe1\xd4\x00\x5d\x67\x03\x8f" +
	"\x4f\xed\x01\x77\x97\x61\x66\x05\xd9\x80\xa3\x0c\xa8\x59\xc5\x09" +
	"\x60\x29\x03\xee\xc2\x60\x05\xe8\x20\x40\xbf\xb0\x56\x6b\x5e\x82" +
	"\x52\xef\xbe\xde\xde\x6f\xf1\x1a\x79\x61\xee\xf1\xf8\x15\x68\xb7" +
	"\xe3\x1f\xe4\xfb\x9a\x59\x82\x8e\x05\x44\x7e\x80\x00\x03\x1a\x81" +
	"\x1f\x22\xc0\xe4\x38\x62\xf8\xac\x41\x24\xa6\x93\x0f\x8b\x0f\x63" +
	"\xc0\x52\x90\x3c\x01\xb0\x7c\xfe\x31\xe4\xe1\xc3\x04\x00\x06\x3f" +
	"\x44\x09\x58\xfc\x00\x06\x80\x83\x0f\x60\x80\xc9\x67\x1b\x00\x1e" +
	"\x5e\x88\x3f\x19\xb4\x0a\x35\x66\x2c\xcf\x51\x00\x00\x00\x00\x49" +
	"\x45\x4e\x44\xae\x42\x60\x82"
