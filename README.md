# chess-app

Purpose: The purpose of this project is to provide an open source, easily accessible multiplayer chess game. This is done via multiple self scaling microservices.

How to deploy: Simple, build/deploy all the docker images, exposing the applicable port and setting the desired environment variables.

# Communication diagram:
    - auth
    - hive -> judicial
    - username
    - chess -> chessbot
    - chessbot -> chess
    - queue -> chess

# TODO:
    - Create common module (DB Connection, DB Schema, EnvVariableChecker)
    - Setup Dockerfiles
    - Change repository so every server has a submodule