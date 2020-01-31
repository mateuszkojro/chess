import socket
import threading
import os
import sys
#
# TODO COLORS



class ChessServer:
	server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

	ip = "10.200.0.211"
	port = "3333"
	flag = True


	def __init__(self, ip_address, port):

		#if addres is not set acepting connections from any ip
		self.ip = ip_address
		self.port = int(port)

		print("ip\t"+self.ip)
		print("port\t"+str(self.port))
		pass



	def init(self, debug=False): #przygotowuje server do odbioru danych
		

		if debug: print("[*] started server at"+str(self.server.getsockname()))

		self.server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
		if debug: print("[*] setted socket option ")

		self.server.bind((self.ip, self.port))
		if debug: print("[*] binded server to " + str(self.server.getsockname()))
		  
		self.server.listen(100) # max number of users / connections
		if debug: print("[*] listening on "+self.ip+":"+str(self.port))



		
	def run(self):


		while self.flag:
			#loop constanly accepting connections sending them greatings and printing msg from them
			c, addr = self.server.accept()  

			r_data = c.recv(1024).decode()

			print(f"[ {addr[0]} : {addr[1]} ] send data and recived: \"{r_data}\"")
			c.send(bytes(f'you sent: "{r_data}"','utf-8'))

		#os.system("watch -n 3 'nmap localhost -p 9090 -Pn'")
		
	

	def kill(self):
		self.server.close()
		print("[!] Server killed")
		pass


def main():
	print("Chess server v 1.0")
	debug = False
	
	# argument ammount checking
	if (len(sys.argv) < 3):
		print("[!] Error: to few arguments, expected 2, "+str(len(sys.argv)) + " given [ip_address, port]")
		sys.exit()

	elif(len(sys.argv)>4):
		print("[!] Error: to many arguments, expected 2, "+str(len(sys.argv)) + " given [ip_address, port]")
		sys.exit()

	elif(len(sys.argv)==4):
		if(sys.argv[3] == "debug"):
			debug = True
			print("[>] debug mode")
	else:
		print("[>] normal mode")
	ip_address = str(sys.argv[1])
	port = str(sys.argv[2])

	#starting server
	srv = ChessServer(ip_address, port)

	if debug:
		srv.init(1)	# 1 debug		0 normal
		srv.run()
	else:
		srv.init(0)
		srv.run()
	
	#srv.kill()


if __name__ == "__main__":  # standalone program as file
	main()
	pass