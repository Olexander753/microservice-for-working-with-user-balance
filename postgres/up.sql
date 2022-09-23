CREATE TABLE Users(
	id 	    VARCHAR(256) PRIMARY KEY,
	amount  INT NOT NULL);

CREATE TABLE Transactions(
    id 	        SERIAL PRIMARY KEY,
    recipient   VARCHAR(256) REFERENCES Users(id) NOT NULL,
    sender      VARCHAR(256) NOT NULL,
	amount      int NOT NULL,
    operation   VARCHAR(256) NOT NULL,
    date_       TIMESTAMP NOT NULL);