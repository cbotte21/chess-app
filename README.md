# chess-app

Purpose: The purpose of this project is to provide a simple, open source, amd horizontally scaling multiplayer chess game.

# How to run?
    - Install docker and packer, run /scripts/install.ps1
    - It is recommended to use docker compose for development and kubernetes for production.

# Communication diagram:
    - auth
    - hive -> judicial
    - username
    - chess -> chessbot
    - chessbot -> chess //Not yet implemented
    - queue -> chess

# Goals:
    - Finish client
    - Build leaderboards
    - Build admin portal
