CREATE TABLE Users(
	id 	    INT PRIMARY KEY,
	amount INT NOT NULL);

CREATE TABLE Transactions(
    id 	        SERIAL PRIMARY KEY,
    recipient   int REFERENCES Users(id) NOT NULL,
    sender      int REFERENCES Users(id),
	amount      int NOT NULL);


    