CREATE TABLE Users(
	id 	    INT PRIMARY KEY,
	balance INT NOT NULL);

CREATE TABLE Transactions(
    id 	        int SERIAL PRIMARY KEY,
    recipient   int REFERENCES Users(id) NOT NULL,
    sender      int REFERENCES Users(id),
	balance     int NOT NULL);