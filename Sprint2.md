### Work Completed:

Scraping: After weeks of trying to use the Go package of "colly", the backend and frontend team decided to take a different approach due to the fact the website was blocking access as it recognized the scraper as a bot. The first differnet approach was to attempt to use Python to scrape the data using the Selenium and Beautiful Soup libraries. After, this too did not pan out due to the same issue, Jacob Kanfer took the approach of JavaScript webscraping. As he had experience with this in the past, he was able to utilize setting user agents to get past the scraper walls and access the data correctly.  
Integrating frontend and backend: We were able to connect the front end and backend smoothly utilizing a search function UI that allows the user to search for elements on the backend.

Property API: Integrated all methods that will be crucial to utilizing our MYSQL database in the future including add, remove, search, update. The next step in this will be connecting the properties on the backend to the frontend.




### Frontend Unit Tests:


### Frontend Cypress Test:

### Backend Unit Tests: 
&nbsp;&nbsp;&nbsp;&nbsp;**search_test.go:**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestSearchAdd

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestSearchGetAll

&nbsp;&nbsp;&nbsp;&nbsp;**properties_test.go:**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestCreateProperty

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestGetProperties

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestUpdateProperty

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestDeleteProperty

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestGetProperty

### Backend API Documentation: 
  
**Summary:**

The Properties package is an API for managing properties. It is designed to allow users to create, read, update, and delete properties. It is written in Go and uses the GORM library for MySQL database interactions.

**Tutorial: How to Use the API**

**First Import the Properties Package**

To use the Properties package, you must first import it into your Go program. To do this, add the following code to your import statements: 

`import "PropertyProbe/properties"`

**Function 1: Create a Property**

To create a new property, call the CreateProperty function and pass in a pointer to a Property struct containing the details of the property you wish to create. The function returns an error if the creation process failed. Here is an example:

```
property := &properties.Property{
    AuctionType:     "Reserve",
    JudgementAmount: 150000.0,
    Address:         "123 Main St",
    AssessedValue:   250000.0,
}

if err := properties.CreateProperty(db, property); err != nil {
    panic(err)
}
```

**Function 2: Retrieve All Properties**

To retrieve all properties, call the GetProperties function and pass in a pointer to a slice of Property structs. The function populates the slice with all properties in the database. The function returns an error if the retrieval process failed. Here is an example:

```
var properties []properties.Property
if err := properties.GetProperties(db, &properties); err != nil {
    panic(err)
}

for _, property := range properties {
    fmt.Printf("%+v\n", property)
}
```

**Function 3: Retrieve a Specific Property**

To retrieve a specific property, call the GetProperty function and pass in a pointer to a Property struct and the address of the property you wish to retrieve. The function populates the struct with the details of the property. The function returns an error if the retrieval process failed. Here is an example:

```
property := &properties.Property{}
if err := properties.GetProperty(db, property, "123 Main St"); err != nil {
    panic(err)
}

fmt.Printf("%+v\n", property)
```

**Function 4: Update a Property**

To update a property, call the UpdateProperty function and pass in a pointer to a Property struct containing the updated details of the property you wish to update. The function returns an error if the update process failed. Here is an example:

```
property := &properties.Property{}
if err := properties.GetProperty(db, property, "123 Main St"); err != nil {
    panic(err)
}

property.AssessedValue = 300000.0
if err := properties.UpdateProperty(db, property); err != nil {
    panic(err)
}
```

**Function 5: Delete a Property**

To delete a property, call the DeleteProperty function and pass in a pointer to a Property struct and the address of the property you wish to delete. The function returns an error if the deletion process failed. Here is an example:

```
property := &properties.Property{}
if err := properties.DeleteProperty(db, property, "123 Main St"); err != nil {
    panic(err)
}
```
