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

	def __init__(self, ip_address, port):
		
		self.ip = ip_address
		self.port = int(port)

		print("ip\t"+self.ip)
		print("port\t"+str(self.port))
		pass



	def run(self, debug=False):
		if debug: print("[*] started server at"+str(self.server.getsockname()))

		self.server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
		if debug: print("[*] setted socket option ")

		self.server.bind((self.ip, self.port))
		if debug: print("[*] binded server to " + str(self.server.getsockname()))
		  
		self.server.listen(100) # max number of users / connections
		if debug: print("[*] listening on "+self.ip+":"+str(self.port))


		#os.system("watch -n 3 'nmap localhost -p 9090 -Pn'")
		os.system("sleep 3")
		pass
	

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
		srv.run(1)	# 1 debug		0 normal
	else:
		srv.run(0)
	
	
	srv.kill()


if __name__ == "__main__":  # standalone program as file

	main()
	pass