## Setup
0) Clone code locally
1) Initialize SQLite Database
   1) Follow setup instructions for SQLite database at https://www.tutorialspoint.com/sqlite/sqlite_installation.htm
   2) In the command line navigate to the project root
   3) Run the command `sqlite3 WeatherAlerts.db` (note, you can put a different database name here if you like, but will need to update it inside config.yaml)
   4) Run the commands in init.sql 
   5) Run the command `.exit` to commit changes and exit the sqlite shell