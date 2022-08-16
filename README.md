# PCLUBTASK5

        The code aligned up with this submission reflects a basic backend of a student search which is linked to a MongoDB database and performs the operations as mentioned: adds a new student entry, deletes an existing student entry, updates a student entry and filters out a student query on any of the parameters realted to a student i.e. name, roll number, branch and uniqueid.
        
   
         MongoDB,Echo,HTTP Package
         MongoDB is used in this implementation as it was easy to get shared free online storage hosting platform at Google Cloud by using MongoDB Atlas.A string connection was made to the Go application by using URI into the code of our program that was possible by bringing into use of MongoDB drivers. A connection to the database is made with a certain definite time period to not overexploit the storage framework and inadverent use of the storage available. Echo as a framework helps us in implementing a framework that acts as a multiplexer that alongwith http package enables an easier listen and serve to the multiple requests made to the server.
         
         HOMEPAGE!
         localhost:9084
         The function homepage is just like the homepage 
         
         ADDITION OF A STUDENT ENTRY!
         localhost:9084/add
         
         The function additon adds a new student entry to the database with all the parameters name, roll number, branch and unique id taken as input through post query made using Postman to simulate the inputs. The basic syntax of the function uses echo framework that plays an important role in creating a multiplexer that can handle several http requests. Using "InsertOne" function available in Go, a new entry is added to the studentsCollection made under the database using the Database function.
          
         DELETION OF A STUDENT ENTRY!
         localhost:/delete
         
         The function deletion deletes an exisiting student entry from the database alongwith all the parameters. The input is simulated again by a post query made using Postman. Using "DeleteOne" function available in go, the entire data related to the queried category and index provided is wrapped away from the existing database.
         
         EDITING A STUDENT ENTRY!
         localhost:9084/edit
         
         The function takes in the required category and the modification that is required to be made and brings in the use of "UpdateOne" function available in Go package to update the student query as per the user's choice.
         
         FINDING A STUDENT ENTRY!
         localhost:9084/find
         
         The function takes in the required category and the value with which the query is made to find a student. The variable details is a map that from the database contains entire details about a student. Iterating over the map, the result is displayed about the user being searched about using String function to print on the datalake.
         
         
         Open to suggestions, improvisations and updates would keep rollling! Keep checking!!
         
         
