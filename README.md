# chess-app

Purpose: The purpose of this project is to provide an open source, easily accessible multiplayer chess game. This is done via multiple self scaling microservices.

How to deploy: Simple, build/deploy all the docker images, exposing the applicable port and setting the desired environment variables.

Server communications:
    hive -> judicial
    chess -> client_socket, archive //bidirectional communication with client
    queue -> chess

# TODO:
    - Extract env variable checks to main. Also add PORT as env variable
    
    - Finish archive-go
    - Finish chess-go
    
    - Build client
