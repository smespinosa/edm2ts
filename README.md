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
