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
    - Finish client
    - Build leaderboards
    - Build admin portal

# How to run?
    - Install docker and packer, run /scripts/install.ps1
    - It is recommended to use docker compose for development and kubernetes for production.
