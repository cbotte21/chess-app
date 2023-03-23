# chess-app

Purpose: The purpose of this project is to provide an open source, easily accessible multiplayer chess game. This is done via multiple self scaling microservices.

How to deploy: Simple, build/deploy all the docker images, exposing the applicable port and setting the desired environment variables.

# Communication diagram:
    - auth
    - hive -> judicial
    - username
    - chess -> chessbot
    - chessbot -> chess //Not yet implemented
    - queue -> chess

# TODO:
    - Setup Dockerfiles
    - Finish client
