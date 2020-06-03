/*!
 	\file
	\brief	udp_new

				UDP comptipix new query command
				g++ -o udp_new udp_new.cpp -g -Wall

 	\author	Benjamin Silvestre
	\date		2019-09-06
*/

// includes
#include <iostream>
#include <iomanip>
#include <fstream>
#include <sstream>
#include <string>
#include <cstring>
#include <cstdlib>

#include <sys/socket.h>
#include <linux/un.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <netinet/in.h>
#include <netdb.h>

//-------------------------------------------------------------------------------
// Definitions
//-------------------------------------------------------------------------------

//! \brief	UDP connection structure
typedef struct
{
	int socket;
	struct sockaddr_in dest_ip;
	unsigned char dest_mac[6];
	struct sockaddr_in source_ip;
	unsigned char source_mac[6];
	int timeout;
} UDP_CONNECTION;

//-------------------------------------------------------------------------------
// Local functions
// Implementation
//-------------------------------------------------------------------------------

/*!
	\brief	Parse mac address

	\param	mac mac address (6 bytes)
	\param	value value to parse, if null init to FF-FF-FF-FF-FF-FF
*/
static void parse_mac_address(unsigned char* mac, const std::string value)
{
	mac[0] = 255;
	mac[1] = 255;
	mac[2] = 255;
	mac[3] = 255;
	mac[4] = 255;
	mac[5] = 255;

	// no check is performed, it is up to the user to check the inputs

	std::istringstream t_parse(value);
	t_parse.setf(std::ios::hex, std::ios::basefield);

	for(size_t i=0;i<6;++i)
	{
		char t_junk;

		if (i>0)
			t_parse >> t_junk;

		unsigned int t_value = 0xFF;
		t_parse >> std::setw(2);
		t_parse >> t_value;
		mac[i] = t_value&0xFF;
	}
}

/*!
	\brief	Print mac address

	\param	mac mac address (6 bytes)
	\return	mac address string
*/
static const std::string print_mac_address(unsigned char* mac)
{
	std::ostringstream t_out;

	t_out.setf(std::ios::hex, std::ios::basefield);
	t_out.fill('0');
	for(size_t i=0;i<6;++i)
	{
		if (i>0)
			t_out << '-';
		t_out << std::setw(2) << (unsigned int)mac[i];
	}

	return t_out.str();
}

/*!
	\brief	Build UDP command

	\param	mac mac address
	\param	command command number
	\param	seq sequence number
	\param	data data
	\return	command
*/
static const std::string build_cmd(const unsigned char* mac, unsigned char command, unsigned int seq, const std::string& data)
{
	std::string t_cmd;

	// build packet

	t_cmd.push_back(0x01);

	// use broadcast mac address

	t_cmd.push_back(mac[0]);
	t_cmd.push_back(mac[1]);
	t_cmd.push_back(mac[2]);
	t_cmd.push_back(mac[3]);
	t_cmd.push_back(mac[4]);
	t_cmd.push_back(mac[5]);

	// add seq number

	t_cmd.push_back((char)((seq>>24)&0xFF));
	t_cmd.push_back((char)((seq>>16)&0xFF));
	t_cmd.push_back((char)((seq>>8)&0xFF));
	t_cmd.push_back((char)(seq&0xFF));

	// add command

	t_cmd.push_back(command);

	// add data

	t_cmd.push_back((char)((data.size()>>8)&0xFF));
	t_cmd.push_back((char)(data.size()&0xFF));

	if (data.size())
		t_cmd.append(data);

	// add crc

	unsigned int t_crc = 0;
	for(size_t i=0;i<t_cmd.size();++i)
		t_crc ^= t_cmd[i];
	t_cmd.push_back((t_crc&0xFF));

	return t_cmd;
}

/*!
	\brief	Parse reply

	\param	mac mac address
	\param 	data data receive
	\param	command command number
	\param	seq sequence number
	\return	command data
*/
static const std::string parse_reply(unsigned char* mac, const std::string& data, unsigned char command, unsigned int seq)
{
	// check length

	if (data.size()<15)
		return std::string();

// 	std::cout << "reply size ok\n";

	// check header

	if (data[0]!=0x01)
		return std::string();

// 	std::cout << "reply header ok\n";

	// read mac address

	memcpy(mac, data.data()+1, 6);

	// check sequence

	const unsigned int t_seq = ((unsigned char)data[7]<<24) + ((unsigned char)data[8]<<16) + ((unsigned char)data[9]<<8) + ((unsigned char)data[10]);
	if (t_seq!=seq)
		return std::string();

// 	std::cout << "reply seq ok\n";

	// check command

	const unsigned char t_cmd = (unsigned char)data[11];
	if (t_cmd!=command)
		return std::string();

// 	std::cout << "reply cmd ok\n";

	// check size

	const unsigned int t_size = ((unsigned char)data[12]<<8) + ((unsigned char)data[13]);
	if ((t_size+15)!=data.size())
		return std::string();

// 	std::cout << "reply data ok\n";

	// check crc

	unsigned int t_crc = 0;
	for(size_t i=0;i<data.size();++i)
		t_crc ^= data[i];

	if (t_crc!=0x00)
		return std::string();

// 	std::cout << "reply crc ok\n";

	// copy answer

	return data.substr(14, t_size);
}

/*!
	\brief	Close an connection

	\param	connection connection object
*/
static void close_udp_connection(UDP_CONNECTION& connection)
{
	if (connection.socket<0)
		return;

	shutdown(connection.socket, 2);
	close(connection.socket);
	connection.socket = -1;
}

/*!
	\brief	Open an udp connection

	\param	connection connection object
	\param	address address to use
	\param	port port to use
	\param	timeout timeout to use
	\return	true if successfull
*/
static bool open_udp_connection(UDP_CONNECTION& connection, const std::string& address, int port, int timeout)
{
	// init structure

	connection.socket = -1;
	connection.timeout = timeout;

	// create socket

	connection.socket = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
	if (connection.socket<0)
		return false;

	// bind to address

	connection.dest_ip.sin_family = AF_INET;
	connection.dest_ip.sin_port = htons(port);

	struct hostent* t_address = gethostbyname(address.c_str());
	if ((t_address) && (t_address->h_addrtype==AF_INET))
		memcpy(&connection.dest_ip.sin_addr, t_address->h_addr, sizeof(struct in_addr));
	else
	{
		close_udp_connection(connection);
		return false;
	}

	return true;
}

/*!
	\brief	Send and receive data

	\param	connection connection object
	\param	command command to send
	\param	send data to send
	\param	receive data received
	\return	true is sucessfull
*/
static bool send_receive_udp(UDP_CONNECTION& connection, unsigned char command, const std::string& t_send, std::string& t_recv)
{
	if (connection.socket<0)
		return false;

	// create command

	const int t_seq = rand();
	const std::string t_cmd = build_cmd(connection.dest_mac, command, t_seq, t_send);

	// send data

	if (sendto(connection.socket, t_cmd.c_str(), t_cmd.size(), MSG_NOSIGNAL, (struct sockaddr*)&connection.dest_ip, (socklen_t)sizeof(connection.dest_ip))<0)
		return false;

// 	std::cout << "send ok\n";

	// wait for data

	fd_set t_list;
	FD_ZERO(&t_list);
	FD_SET(connection.socket, &t_list);

	struct timeval tv;
	tv.tv_sec = connection.timeout/1000;
	tv.tv_usec = (connection.timeout%1000)*1000;
	if (select(connection.socket+1, &t_list, NULL, NULL, &tv)<=0)
		return false;

// 	std::cout << "select ok\n";

	// recv data

	int t_in_size = sizeof(struct sockaddr_in);
	char t_data[65536];

	int t_reply = recvfrom(connection.socket, t_data, sizeof(t_data), MSG_NOSIGNAL, (struct sockaddr*)&connection.source_ip, (socklen_t*)&t_in_size);
	if (t_reply<0)
		return false;

// 	std::cout << "recv ok\n";

	t_recv.assign(t_data, (size_t)t_reply);
	t_recv = parse_reply(connection.source_mac, t_recv, command+0x80, t_seq);
	return true;
}

/*!
	\brief	Display help

	\param	prog name of the program
*/
static void print_help(const std::string& prog)
{
	std::cout << "Usage :\n";
	std::cout << "./" << prog << " [options] request\n";

	std::cout << "\nOptions :\n";
	std::cout << "-h or --help      print help\n";
	std::cout << "-i or --ip        set destination ip address\n";
	std::cout << "-m or --mac       set destination mac address\n";
	std::cout << "-p or --port      set destination UDP port (default 1600)\n";
	std::cout << "-f or --file      write reply to file\n";
	std::cout << "--get             make a query to CONFIG? (default)\n";
	std::cout << "--image           make a query to IMAGE?\n";

	std::cout << "\nExamples :\n";
	std::cout << "Get some basic info from a sensor by its ip address :\n";
	std::cout << "./" << prog << " --ip 192.168.0.100 --get \"serial&type&uptime\"\n";

	std::cout << "\nGet some basic info from a sensor by its mac address :\n";
	std::cout << "./" << prog << " --mac \"01:23:45:67:89:AB\" --get \"serial&type&uptime\"\n";

	std::cout << "\nGet sensor tiny image and write it to a file :\n";
	std::cout << "./" << prog << " --ip 192.168.0.100 --image \"tiny\" -f tiny.jpg\n";
}

//-------------------------------------------------------------------------------
// Shared functions
// Implementation
//-------------------------------------------------------------------------------

int main(int argc, char** argv)
{
	std::string dest_ip;
	std::string dest_mac_txt;
	int dest_port = 1600;
	bool cmd_image = false;
	std::string cmd_request;
	std::string dest_file;

	// quick & dirty parsing of command line

	for(int i=1;i<argc;++i)
	{
		if ((std::string(argv[i])=="-h") || (std::string(argv[i])=="--help"))
		{
			print_help(argv[0]);
			return 0;
		}
		else if ((std::string(argv[i])=="-i") || (std::string(argv[i])=="--ip"))
		{
			if ((i+1)>=argc)
			{
				std::cerr << "missing arg for " << argv[i] << "\n";
				std::cout << "use -h or --help to display help\n";
				return -1;
			}

			dest_ip.assign(argv[++i]);
		}
		else if ((std::string(argv[i])=="-m") || (std::string(argv[i])=="--mac"))
		{
			if ((i+1)>=argc)
			{
				std::cerr << "missing arg for " << argv[i] << "\n";
				std::cout << "use -h or --help to display help\n";
				return -1;
			}

			dest_mac_txt.assign(argv[++i]);
		}
		else if ((std::string(argv[i])=="-p") || (std::string(argv[i])=="--port"))
		{
			if ((i+1)>=argc)
			{
				std::cerr << "missing arg for " << argv[i] << "\n";
				std::cout << "use -h or --help to display help\n";
				return -1;
			}

			dest_port = strtol(argv[++i], NULL, 10);
		}
		else if ((std::string(argv[i])=="-f") || (std::string(argv[i])=="--file"))
		{
			if ((i+1)>=argc)
			{
				std::cerr << "missing arg for " << argv[i] << "\n";
				std::cout << "use -h or --help to display help\n";
				return -1;
			}

			dest_file.assign(argv[++i]);
		}
		else if (std::string(argv[i])=="--get")
			cmd_image = false;
		else if (std::string(argv[i])=="--image")
			cmd_image = true;
		else if (argv[i][0]=='-')
		{
			std::cerr << "unknown cmd " << argv[i] << "\n";
			std::cout << "use -h or --help to display help\n";
			return -1;
		}
		else
			cmd_request.assign(argv[i]);
	}

	// check we have at leat an ip address

	if ((0==dest_ip.size()) && (0==dest_mac_txt.size()))
	{
		std::cout << "error : please specify an ip or mac address\n";
		return -1;
	}

	// no ip specified, let's use broadcast !

	if (0==dest_ip.size())
		dest_ip.assign("255.255.255.255");

	// parse mac address

	unsigned char dest_mac[6];
	parse_mac_address(dest_mac, dest_mac_txt);

	// display some basic info

	std::cout << "connect to " << dest_ip << " mac " << print_mac_address(dest_mac) << "\n";

	// create connection

	UDP_CONNECTION t_connection;
	if (!open_udp_connection(t_connection, dest_ip, dest_port, 2000))
	{
		std::cerr << "error : unable to connect to " << dest_ip << "\n";
		return -1;
	}

	std::string t_recv;
	unsigned t_cmd = 0x11;
	if (cmd_image)
		t_cmd = 0x12;

	memcpy(t_connection.dest_mac, dest_mac, 6);

	if (!send_receive_udp(t_connection, t_cmd, cmd_request, t_recv))
		std::cerr << "error : no reply or communication error\n";
	else
	{
		std::cout << "reply from " << inet_ntoa(t_connection.source_ip.sin_addr) << " mac " << print_mac_address(t_connection.source_mac) << " -> " << (unsigned int)t_recv.size() << " bytes\n";

		// write to file

		if (dest_file.size())
		{
			std::ofstream t_stream(dest_file.c_str(), std::ofstream::out|std::ofstream::binary);
			if (!t_stream.good())
				std::cerr << "error : unable to write " << dest_file << "\n";
			else
			{
				t_stream << t_recv;
			}
		}
		else
		{
			std::cout <<  t_recv << "\n";
		}
	}

	close_udp_connection(t_connection);

	return 0;
}
