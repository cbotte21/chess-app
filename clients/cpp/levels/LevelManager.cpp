#include "LevelBuilder.cpp"
#include "Level.cpp"

class LevelManager {
	Level* pLevel = nullptr;
	public:
		LevelManager() {
			LevelBuilder::Map map = LevelBuilder::MapMenu; //Main menu
			pLevel = LevelBuilder().setLevel(map).build();
		}
		explicit LevelManager(Level* pLevel) {
			this->pLevel = pLevel;
		}
		Level* getLevel() {
			return pLevel;
		}
		void setLevel(Level* pLevel) {
			this->pLevel = pLevel;
		}
};
