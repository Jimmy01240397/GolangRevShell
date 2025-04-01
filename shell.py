import socket
import sys
import threading

if len(sys.argv) <= 2:
    print(f"{sys.argv[0]} <port> <encode>")
    exit(1)

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind(("0.0.0.0", int(sys.argv[1])))
sock.listen()

conn, addr = sock.accept()

def recv():
    while True:
        data = conn.recv(65535)
        #sys.stdout.buffer.write(data.decode(sys.argv[2]).encode())
        sys.stdout.buffer.write(data)
        sys.stdout.buffer.flush()

recvthread = threading.Thread(target=recv)
recvthread.start()

while True:
    conn.send(sys.stdin.buffer.readline().decode().encode(sys.argv[2]))
    sys.stdin.buffer.flush()


