#pragma once
#include <SFML/Graphics.hpp>

class Level : public sf::Drawable {
	public:
		virtual void onEvent(sf::Event event, sf::Window& window) {}
		virtual void draw(sf::RenderTarget& target, sf::RenderStates states) const = 0;
};
