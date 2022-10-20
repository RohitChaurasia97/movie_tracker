# movie_tracker

Simple Crud app built using bean framework.
As explained in bean's readme bean is designed based on the service repository design pattern.

```
Router Interface delegates the requests to the handler layer, where the handler layer links with one/multiple service(s) 
and each of these service layers is where the business logic is implemented.
The service layer then links with the repository layer which can be anything a SQL database or NoSQL db.
```

<img width="358" alt="Screenshot 2022-10-19 at 15 39 32" src="https://user-images.githubusercontent.com/97086931/196615431-4fcbf2b6-f6bf-4bbe-b3c0-b953a626b5b8.png">


For a real live application containing multiple services, functionalities and databases the architecture would scale up to look like :-

<img width="757" alt="Screenshot 2022-10-19 at 15 45 26" src="https://user-images.githubusercontent.com/97086931/196616535-8b111331-75d2-4b33-b3b2-e7a81d2ae0bc.png">


### Setup in Local:- 

#### - Clone the repository in local.   
  ```
  git clone https://github.com/RohitChaurasia97/movie_tracker
  ```

#### - Change the directory.
  ```
  cd movie_tracker
  ```

#### - After changing directory use docker-compose to setup the project in local
  ```
  docker compose up --build
  ```

#### - Check status of services in docker compose 
  ```
  docker compose ps
  ```

#### - Wait for mysqldb container and movie container to be running and healthy. Access localhost on port 9000 at route /movies for the movie_app
  ```
  http://localhost:9000/movies
  ```

### - To access the mysqldb service from terminal
  ```
   docker compose exec mysqldb mysql -u root local
   ```
 

https://user-images.githubusercontent.com/97086931/196847900-4abd0f3f-643f-45fa-955c-f19c87ca8791.mp4


### Setup in Gitpod:- 

####   - You can add gitpod extension to your browser or directly create a workspace in gitpod by accessing
  ```
  https://gitpod.io#github.com/RohitChaurasia97/movie_tracker
  ```

####    - Wait for mysqldb container and movie container to be running and healthy. You can check the status of containers with the help of docker-compose ps (in gitpod's terninal)

####   - Access the workspace on port 9000 at route /movies for the movie_app (using remote explorer extension in gitpod's VS Code IDE)



https://user-images.githubusercontent.com/97086931/196848344-cc7985b7-8cc6-49f9-9f18-523bdaaa619d.mp4

