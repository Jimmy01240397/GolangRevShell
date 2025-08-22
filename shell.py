#!/usr/bin/env python3

import socket
import sys
import threading
import ssl

if len(sys.argv) <= 2:
    print(f"{sys.argv[0]} <port> <cert> <key> <encode>")
    exit(1)

context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
context.load_cert_chain(certfile=sys.argv[3], keyfile=sys.argv[4])

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind(("0.0.0.0", int(sys.argv[1])))
sock.listen()

conn, addr = sock.accept()
conn = context.wrap_socket(conn, server_side=True)

def recv():
    while True:
        data = conn.recv(65535)
        sys.stdout.buffer.write(data.decode(sys.argv[2], 'ignore').encode())
        #sys.stdout.buffer.write(data)
        sys.stdout.buffer.flush()
        if len(data) <= 0:
            conn.close()
            return

recvthread = threading.Thread(target=recv)
recvthread.start()

while True:
    data = sys.stdin.buffer.readline()
    try:
        conn.send(data.decode('utf8', 'ignore').encode(sys.argv[2]))
    except:
        conn.close()
        break
    sys.stdin.buffer.flush()


