#include "AuthSession.h"

bool AuthSession::authenticate(std::string username, std::string password) {
	//TODO: Make request to server
	return true;
}

bool AuthSession::isAuthenticated() {
	//TODO: Make request to server
	return true;
}

void AuthSession::logout() {
	token = std::string("");
}
