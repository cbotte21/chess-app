#include "Level.cpp"

#include "Menu.cpp"
#include "Game.cpp"

class LevelBuilder {
	Level* pLevel = nullptr;

	public:
		enum Map {
			MapMenu,
			MapGame
		};
		LevelBuilder setLevel(Map map) {
			switch (map) {
				case MapMenu:
					pLevel = new Menu();
				case MapGame:
					pLevel = new Game();
			}
			return *this;
		}
		Level* build() {
			return pLevel;
		}
};
