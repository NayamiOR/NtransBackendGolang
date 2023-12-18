# NayaTrans

信息和文件传输工具，支持多台公网设备和内网设备之间的信息和文件传输（传输的双方至少有一方有公网 IP）。

## 接口文档

假设A是public服务器，B是private端。

B给A发消息发文件可以直接通过POST请求

A给B发消息需要用前端调用uploadfiles或者uploadmessage，可以往信息和文件列表添加信息。

B访问files和messages可以获取A上传过的文件和信息。message会直接返回A端发送的信息。

files会返回一个文件列表，里面是A端发送的每个文件对应的路由。

﻿

### POSTtest receivemessage

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/receivemessage

接收消息（B->A）

﻿

Bodyurlencoded

text

any-text

### POSTtest receivefiles

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/receivefiles

接收文件（B->A）

﻿

Bodyform-data

file

/D:/Videos/Death Note 1-37/死亡笔记 - 36.mp4

### POSTtest uploadfiles

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/uploadfiles

上传文件（A->B）

﻿

Bodyform-data

file

/D:/Pictures/Screenshots/Snip231215_1703.png,/D:/Pictures/Screenshots/Snip231208_1541.png

### POSTtest uploadmessage

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/uploadmessage

上传消息（A->B）

﻿

Bodyform-data

message

any-message

### GETtest

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/test

测试运行

﻿

### GETtest printroutes

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/printroutes



让后端输出当前监听的端口

﻿

### GETtest files

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/files



获取后端上传的文件列表，用于访问获取文件（A->B）

﻿

### GETtest messages

[Open request](https://desktop.postman.com/?desktopVersion=10.21.0&userId=25418993&teamId=0)

http://localhost:18080/messages



获取后端上传的信息列表，用于获取信息（A->B）

﻿
