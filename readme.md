## Description

Create the backend for a double­entry accounting and show how to handle transactions such as the income from a sale being divided between tax and revenue accounts.


## Assumptions

- The business is operating in one country and selling internationaly
- For the goods or services we sell there is only one tax rate in each country
- For each financial year exchange rates in our system are fixed
- An account can only exist in one country
- An accountholder manages the accounts for one country
- Sale and handover are independent i.e. a payment being received and products being sent out can happen on different days
- the base currency is US Dollar


# Installation

If you don't have mysql installed go ahead, on mac just do `brew install mysql` I have used the default username of root and password empty but you use whatever you like, just update models/dbModel.go with your settings.

Once mysql is setup locally run `schemaSetup.sql` to create the db and tables.

Possibly depending on go setup you may need to run

```
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
```

Final step is run `go get`, `go install` and `go build` from the root of your application, then if all goes fine you visist the site at `http://localhost:8080`





For exmample usage please see curl commands in controller function headers





