import socket
import sys

s = socket.socket()

ip = '127.0.0.1'
port = 3333

s.connect(( ip, port)) #polaczenie pod adresem ip i portem port

x = s.recv(1024) #odbieram 1024 bajty
x = x.decode() #dekoduje do str

print(x)

msg = sys.argv[1],'utf-8' #tworze wiadomosc do wyslania jakos string

s.send( bytes(msg) ) #konwertuje na bajty i wysylam
s.close()