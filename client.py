import socket
import sys

s = socket.socket()

ip = "217.113.232.164"
port = 3333

s.connect(( ip, port)) #polaczenie pod adresem ip i portem port

msg = sys.argv[1] #tworze wiadomosc do wyslania jakos string

s.send( bytes(msg,'utf-8') ) #konwertuje na bajty i wysylam

x = s.recv(1024) #odbieram 1024 bajty
x = x.decode() #dekoduje do str
print(x)

s.close()