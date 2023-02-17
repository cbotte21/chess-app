CC = g++
BUILD_DIR = build
CFLAGS = -pedantic -Wall
OBJ_DIR = build/obj
OBJS = $(OBJ_DIR)/LevelBuilder.o $(OBJ_DIR)/LevelManager.o $(OBJ_DIR)/AuthSession.o
LEVELS = levels/Game.cpp levels/Menu.cpp
EXE = $(BUILD_DIR)/chess_client
LIBS = -lsfml-graphics -lsfml-window -lsfml-system

all: main.cpp $(OBJS)
	$(CC) main.cpp $(OBJS) $(CFLAGS) -o $(EXE) $(LIBS)
$(OBJ_DIR)/LevelManager.o: levels/LevelManager.cpp
	$(CC) -c levels/LevelManager.cpp -o $(OBJ_DIR)/LevelManager.o
$(OBJ_DIR)/LevelBuilder.o: levels/LevelBuilder.cpp $(LEVELS)
	$(CC) -c levels/LevelBuilder.cpp -o $(OBJ_DIR)/LevelBuilder.o
$(OBJ_DIR)/AuthSession.o: auth/AuthSession.h
	$(CC) -c auth/AuthSession.cpp -o $(OBJ_DIR)/AuthSession.o
clean:
	\rm -f $(EXE) $(OBJS)
