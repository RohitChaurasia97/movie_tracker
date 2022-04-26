# movie_tracker

Simple Crud app built using bean.

  ### Setup in Local:- 
  
  #### - Clone the repository in local.   
    
    git clone https://github.com/RohitChaurasia97/movie_tracker

  #### - After changing directory use docker-compose to setup the project in local

    docker-compose up --build

  #### - Wait for mysqldb container and movie container to be running and healthy. You can check the status of containers with the help of docker-compose ps. Access localhost on port 9000 at route /movies for the movie_app

    http://localhost:9000/movies
 
  ### Setup in Gitpod:- 
        
  ####   - You can add gitpod extension to your browser or directly create a workspace in gitpod by accessing
    https://gitpod.io#github.com/RohitChaurasia97/movie_tracker
      
  ####    - Wait for mysqldb container and movie container to be running and healthy. You can check the status of containers with the help of docker-compose ps (in gitpod's terninal)
      
  ####   - Access the workspace on port 9000 at route /movies for the movie_app (using remote explorer extension in gitpod's VS Code IDE)
