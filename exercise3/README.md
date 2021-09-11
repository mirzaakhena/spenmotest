### Design schema

```
Usecase story:

There is a list of users, 
each user can have many wallets. 
Each wallet holds the balance amount the user has. 
Each wallet has many cards. 
Each card can have limits
limit on money that can be spent daily, monthly.

Each user can also be part of 1 or more teams. 
Each team has a wallet. 
Team wallets have also have daily, monthly limits
```

My schema is in class diagram form.

Read [Here](https://github.com/mirzaakhena/spenmotest/blob/main/exercise3/current_situation.go) for detail case

![Class Diagram](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/mirzaakhena/spenmotest/main/exercise3/doc/class_diagram.puml)