#include <string>

class AuthSession {
	std::string token;
	public:
		bool authenticate(std::string username, std::string password);
		bool isAuthenticated();
		void logout();
};
