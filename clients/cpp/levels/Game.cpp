#pragma once
#include "Level.cpp"

class Game : public Level {
	public:
		void onEvent(sf::Event event, sf::Window& window) {

		}
		virtual void draw(sf::RenderTarget& target, sf::RenderStates states) const {}
};
