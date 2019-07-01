# edm2ts
Convert EDM metadata to typescript classes

## Compile
Go will compile natively for your operating system.  Simply download the repository and build the project from the directory.

```
cd edm2ts
go build
```

## Usage
The metadata URL must be provided.  The output file defaults to 'entities.ts' but can be defined as the second argument.

```
edm2ts https://api.somesite.com/odata/$metadata odataclasses.ts
```

## Sample output
Generated from https://services.odata.org/V3/(S(jhv0jsgccuohvw4uwoaq3nzo))/OData/OData.svc/$metadata

```
export class Product {
	ID: number
	Name: string
	Description: string
	ReleaseDate: string
	DiscontinuedDate: string
	Rating: string
	Price: string
}

export class FeaturedProduct {
}

export class ProductDetail {
	ProductID: number
	Details: string
}

export class Category {
	ID: number
	Name: string
}

export class Supplier {
	ID: number
	Name: string
	Address: string
	Location: string
	Concurrency: number
}

export class Person {
	ID: number
	Name: string
}

export class Customer {
	TotalExpense: number
}

export class Employee {
	EmployeeID: number
	HireDate: string
	Salary: string
}

export class PersonDetail {
	PersonID: number
	Age: string
	Gender: boolean
	Phone: string
	Address: string
	Photo: string
}

export class Advertisement {
	ID: string
	Name: string
	AirDate: string
}

```

## Todo
Resolve complex types to correct class
