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
Generated from https://services.odata.org/V4/(S(jhv0jsgccuohvw4uwoaq3nzo))/OData/OData.svc/$metadata

```
export interface Product {
	ID: number;
	Name: string | null;
	Description: string | null;
	ReleaseDate: Date;
	DiscontinuedDate: Date | null;
	Rating: string;
	Price: string;
	Categories?: Category[];
	Supplier?: Supplier;
	ProductDetail?: ProductDetail;
}

export interface FeaturedProduct {
	Advertisement?: Advertisement;
}

export interface ProductDetail {
	ProductID: number;
	Details: string | null;
	Product?: Product;
}

export interface Category {
	ID: number;
	Name: string | null;
	Products?: Product[];
}

export interface Supplier {
	ID: number;
	Name: string | null;
	Address: string | null;
	Location: string | null;
	Concurrency: number;
	Products?: Product[];
}

export interface Person {
	ID: number;
	Name: string | null;
	PersonDetail?: PersonDetail;
}

export interface Customer {
	TotalExpense: number;
}

export interface Employee {
	EmployeeID: number;
	HireDate: Date;
	Salary: string;
}

export interface PersonDetail {
	PersonID: number;
	Age: string;
	Gender: boolean;
	Phone: string | null;
	Address: string | null;
	Photo: string;
	Person?: Person;
}

export interface Advertisement {
	ID: string;
	Name: string | null;
	AirDate: Date;
	FeaturedProduct?: FeaturedProduct;
}
```
