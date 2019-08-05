import socket

i = 0
while i < 5:
    # 声明socket类型，同时生成socke连接t对象
    client = socket.socket()
    # 连接到localhost主机的8999端口上去
    client.connect(('localhost', 8999))
    msg = "hehehehehehhe" + str(i)
    i = i + 1
    # 把编译成utf-8的数据发送出去
    client.send(msg.encode('utf-8'))

    # 接收数据
    data = client.recv(8192)
    print("从服务器接收到的数据为：", data.decode())
    client.send(msg.encode('utf-8'))
    client.close()
