#include <SFML/Window.hpp>

#include <iostream>
#include <string>

#include "auth/AuthSession.h"
#include "levels/LevelManager.cpp"

/*
 *	1) Authenticate
 *	2) Open window
 */

int main(int argc, const char* argv[]) {
	AuthSession authSession;
	std::string email, password;
	
	//Welcome message
	std::cout << "=--=WELCOME=--=\n" << std::endl;

	//Authenticate
	do {
		std::cout << "  Email: ";
		std::cin >> email;
		std::cout << "  Password: ";
		std::cin >> password;
		if (!authSession.authenticate(email, password)) {
			std::cout << "\nAuthentication failed!" << std::endl;
		}
	} while (!authSession.isAuthenticated());
	std::cout << "\nUser authenticated!\nLoading client..." << std::endl;

	//LevelManager init
	LevelManager levelManager;

	//Window setup
	sf::RenderWindow window(sf::VideoMode(800, 600), "Chess client");	
	window.setVerticalSyncEnabled(true);
	window.setFramerateLimit(60);

	while(window.isOpen()) {
		sf::Event event;
		while (window.pollEvent(event)) {
			levelManager.getLevel()->onEvent(event, window);
		}
		window.clear();
		window.draw(*levelManager.getLevel());
		window.display();
	}

	return 0;
}
