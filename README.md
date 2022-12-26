# BookMyShow

Use Cases for BookMyShow

1. As a user I should be able to log in/signup to the website
2. As a user I should be able to search for movies based on your preference (language,rating,genre) 
3. As a user I should be able to book movie tickets and pay for the booking
4. As a user I should be able to view the latest movie ratings by imdb
5. As a user I should be able to cancel movie bookings
6. As a user I should be able to update your information
7. As an admin I should be able to add new movies
8. As an admin I should be able to set the status of movies (active/inactive)
9. As an admin I should be able to update show details for a particular date

Steps to use the application

1. Download go from official website "https://go.dev/dl/"

2. Clone the repository using 
    git clone "https://github.com/Hemanth-2002/BookMyShow.git"

3. Open terminal in the cloned folder 

4. Start the server using 
    go run cmd/server.go

5. The proto files are present in bmsproto folder. (proto files can be uploaded in postman to test the grpc endpoints)

6. Run the client file using 
    go run client/client.go

7. The grpc endpoints implementation is present in server folder. It can be accessed by \n
    cd server

8. To test the files use the command
    go test -v / go test
