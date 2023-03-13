# chess-app

Purpose: The purpose of this project is to provide an open source, easily accessible multiplayer chess game. This is done via multiple self scaling microservices.

How to deploy: Simple, build/deploy all the docker images, exposing the applicable port and setting the desired environment variables.

# Future plans:
I plan to make this reposository contain server dockerfiles, and documentation about server characteristics, interactions, and more. All subprojects will be moved to their respected repository, allowing for correct usage of go modules.

servers/* directory will remain intact, all subfolders will be loose -go suffix, as this repository is meant for overall project flow, and not microservice implementation. All proto files should remain in this repository.

After this git migration is done, I will create Dockerfiles for all servers, allowing them to be deployed easily.

Lastly a client will be build. Most likely a basic nodejs frontend.

# Communication diagram:
    - auth
    - hive -> judicial
    - username
    - chess -> chessbot
    - chessbot -> chess
    - queue -> chess

# TODO:
    - Move servers to seperate repositories
    - Setup Dockerfiles
    - Create client
