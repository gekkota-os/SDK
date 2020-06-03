#----------------------------------------------------------------------------------------------------------------------------------
#			Eurecam Demo program receiving our POST protocol, fell free to use this exemple to build your product.
#			This is a Demo program.
#			This program is distributed in the hope that it will be useful,
#			but WITHOUT ANY WARRANTY; without even the implied warranty of
#			MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
#----------------------------------------------------------------------------------------------------------------------------------

# This program is an example of POST send in python

# !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
# !!!								IMPORTANT :                                                      !!!
# !!! 	This program use python BaseHTTPServer which, at the time of writing,              !!!
# !!! 	is not a production grade server : use it for debuging not for production          !!!
# !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

# Import
import SocketServer
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
from cgi import parse_header, parse_multipart
import urlparse
import os

# Config
PORT_NUMBER = 8081		# Listening port
DATA_DIR = 'data'			# Directory to store data
ANSWER_STOP = '2' 		# Could be "1": skip next file, or "2" skip all remaining files
DEFAULT_TYPE = 'CPX3'	# set a default type if sensor don't set it (Comptipix version < 1.3.0 don't set the type)
DEBUG_ENABLED = True 	# true to enable a simple debug, false otherwise

def debug_func(str_debug):
	if DEBUG_ENABLED:
		print str_debug

# 1 function to do all the job :)
def request_func(self):
	# Get a request object, important to keep blank because 'check' parameter is blank
	ro = urlparse.parse_qs(urlparse.urlparse(self.path).query, keep_blank_values=True)

	# Get url values
	type_val = ''
	serial_val = ''
	file_val = ''
	size_val = ''
	if 'type' in ro:
		type_val = ro['type'][0]
	if 'serial' in ro:
		serial_val = ro['serial'][0]
	if 'file' in ro:
		file_val = ro['file'][0]
	if 'size' in ro:
		size_val = ro['size'][0]

	# Set a default type (comptipix version <1.3.0 has no type info)
	if '' == type_val:
		type_val = DEFAULT_TYPE

	# Check or write file
	if ('' != type_val) and ('' != serial_val) and ('' != file_val) and ('' != size_val):
		file_path = DATA_DIR + '/' + type_val + '-' + serial_val # file path is like 'Data/CPX3-119092'
		file_with_path = os.path.join(file_path, file_val) # = file_path + '/' + file_val

		# Check command
		if 'check' in ro:
			debug_func('<- Check for file('+size_val+'): ' + file_with_path) # Debug

			if not os.path.exists(file_path):
				return '0' # path not exist

			if not os.path.isfile(file_with_path):
				return '0' # file do not exist

			statinfo = os.stat(file_with_path)
			if (int(size_val) != statinfo.st_size):
				return '0' # size are different

			# if we reach here file exist and size are the same
			return ANSWER_STOP

		# Receive data
		elif 'data' in ro:
			debug_func('<- Receive data('+size_val+') for file: ' + file_with_path) # Debug

			if not os.path.isdir(file_path):
				debug_func('Create directories: ' + file_path) # Debug
				os.makedirs(file_path) # directory dosn't exist yet : create it

			post_data = self.rfile.read(int(self.headers.getheader('content-length'))) # get raw post data
			with open(file_with_path, 'wb') as temp_file:
				temp_file.write(post_data) # Create file with post data

			return '1'


# MAIN: the http handler
class MyHandler(BaseHTTPRequestHandler):
	protocol_version = 'HTTP/1.1' # IMPORTANT: set protocol to http/1.1 to allow persistent connection (see https://docs.python.org/2/library/basehttpserver.html)

	# Get handler (for check command)
	def do_GET(self):
		debug_func('GET: ' + self.path) # Debug

		# Get answer
		answ_get = request_func(self)

		debug_func('--> ' + answ_get) #Debug

		# Send answer with content-length header
		self.send_response(200)
		self.send_header('Content-Length', len(answ_get)) # IMPORTANT: content-length header is mandatory for comptipix (+ also to have a working persistent connection)
		self.send_header('Content-Type', 'text/plain')
		self.end_headers()
		self.wfile.write(answ_get)
		self.wfile.close()

	# Post handler (to get file data)
	def do_POST(self):
		debug_func('POST: ' + self.path) # Debug

		# Get answer
		answ_post = request_func(self)

		debug_func('--> ' + answ_post) # Debug

		# Send answer with content-length header
		self.send_response(200)
		self.send_header('Content-Length', len(answ_post)) # IMPORTANT: content-length header is mandatory for comptipix (+ also to have a working persistent connection)
		self.send_header('Content-Type', 'text/plain')
		self.end_headers()
		self.wfile.write(answ_post)
		self.wfile.close()

# Start http server
try:
	print 'server listening on port '+str(PORT_NUMBER)
	httpd = SocketServer.TCPServer(("", PORT_NUMBER), MyHandler)
	httpd.serve_forever()

# Interrupt server on keyboard Ctrl+C
except KeyboardInterrupt:
	print '^C received, shutting down the web server'
	server.socket.close()
