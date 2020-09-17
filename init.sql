CREATE TABLE User (UserId INTEGER PRIMARY KEY, DisplayName VARCHAR, Email VARCHAR, Zipcode VARCHAR);
CREATE TABLE Notification (NotificationId INTEGER PRIMARY KEY, UserId INTEGER, Frequency INTEGER, Alert INTEGER, FOREIGN KEY(UserId) REFERENCES User(UserId));

-- temp:
INSERT into User (DisplayName, Email, Zipcode) VALUES ('Test', '2565274055@msg.fi.google.com', '35124')
