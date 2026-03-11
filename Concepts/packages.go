package concepts

/*
In Go, a package is simply a collection of related Go files that are compiled together.


folder/
   file1.go
   file2.go
   file3.go


Without packages, large programs become impossible to manage.
Imagine writing a backend with:
database code
authentication
HTTP handlers
logging

project/

	main.go

	db/
	connect.go

	auth/
	login.go
	jwt.go

	handlers/
	user.go
	product.go

Without packages:

Add()
Add()
Add()

Conflicts everywhere.

With packages:

math.Add()
cart.Add()
vector.Add()

No confusion.
*/