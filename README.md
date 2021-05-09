# DistributedRPC
This repository consist of GoLang RPC distributed library code.

This is a command line interface project to demonstrate the RPC library of GoLang.

You will find two Folders, Client and Server.

Ther are two end user types, User and Manager.

Client Folder: 

ClientDriver.go:
This file consist of the main function. It creates a new client object and initiates the rpc client and lookup for appropriate library server to make remort procedure calls. 

Client.go : 
This file has all the juice. The rpc client check which end user is using the application and shows the feature menu for that end user. The end user can select any feature from the list
and initiate the remote call. The server then processes the query and sends a reply back and the client displays the reply message.

Server Folder:

Items.go : 
This is a basic structure to represent Item model to maintain the database.

User.go : 
This is a model file to represent a user in the database. 

Manager.go
This again is a structure to represent the Manager in the database.

Server.go
This file has all the implementation of remort procedure calls. This Server structure has the database in form of Maps. 
There are three libraries which can is predefined and the ports for these libraries are defined already. 

Concordia Library = "CON" on port = ":1301".

McGill library = "MCG" on port = ":1302".

Montreal library = "MON" on port = ":1303".

There are 4 main database maps.

books map with Key : string bookID and value : Item object.

users map with Key : string userID and value : User object.

managers map with Key : string ManagerID and value : Manager object.

borrowed mao with Key : User object and value as another map with Key : Item object and value : integer to represent number of days for which item is borrowed.

I have used Jet brains GoLand IDE to develop this code.

If you wish to execute this code, please follow the steps below

Step1) Download the source code from master brantch or pull the code from the link

Step2) Execute the ServerDriver.go to boot up all the three servers.

Step3) Execute the ClientDriver.go to boot the client.

Step4) Follow the menu instruction on the rpc client.

The pre-populated data on these servers:

Concrdia

Users :

| UserID   | Library | Index |
|----------|---------|-------|
| CONU1001 | CON     | 1001  |
| CONU1002 | CON     | 1002  |

Managers :

| ManagerID | Library | Index |
|-----------|---------|-------|
| CONM1001  | CON     | 1001  |

Items:

| ItemID  | ItemName             | Item Count |
|---------|----------------------|------------|
| CON1001 | Distributed Systems  | 1          |
| CON1002 | Parallel Programming | 6          |
| CON1003 | Algorithm Designs    | 7          |

McGil

Users :

| UserID   | Library | Index |
|----------|---------|-------|
| MCGU1001 | MCG     | 1001  |
| MCGU1002 | MCG     | 1002  |

Managers :

| ManagerID | Library | Index |
|-----------|---------|-------|
| MCGM1001  | MCG     | 1001  |

Items:

| ItemID  | ItemName             | Item Count |
|---------|----------------------|------------|
| MCG1001 | Distributed Systems  | 1          |
| MCG1002 | Parallel Programming | 6          |
| MCG1003 | Algorithm Designs    | 7          |

Montreal

Users :

| UserID   | Library | Index |
|----------|---------|-------|
| MONU1001 | MON     | 1001  |
| MONU1002 | MON     | 1002  |

Managers :

| ManagerID | Library | Index |
|-----------|---------|-------|
| MONM1001  | MON     | 1001  |

Items:

| ItemID  | ItemName             | Item Count |
|---------|----------------------|------------|
| MON1001 | Distributed Systems  | 1          |
| MON1002 | Parallel Programming | 6          |
| MON1003 | Algorithm Designs    | 7          |











