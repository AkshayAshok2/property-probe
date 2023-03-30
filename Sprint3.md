### Work Completed:

**Database Connection:** With our newly scraped data, the backend team was able to add the properties successfully into the MySQL database. Since some of the properties did not have addresses, this transfer from textfile to database took longer than expected, but we were in fact able to add the properties to the database using addProperties.go. Also with regards to the database, the backend team edited the main.go so that when the server is booted up, the database is cleared before the properties are added. This ensures that on boot up, the database is only filled with the most relevant foreclosure properties.

**Property API:** In addition to all the methods that will be crucial to utilizing our MYSQL database in the future including add, remove, search, and update, the backend team also added a GetZipCodeProperties Get call that will produce an array of properties with that specific zip code. This will allow the frontend to easily retrieve certain properties based on what zip code is searched. The next step in this will be connecting the properties on the backend to the frontend.

**Frontend Progress:** Using Angular Material, we have established the overall layout of the project's user interface. Furthermore, we have implemented a list-style container that shows properties and their basic information. For each property in this container, more information is revealed at the right-hand side of the screen when the property component is clicked. Additionally, the frontend team has further developed the application's map component to show property markers at given latitude and longitude components. In the next sprint, the property-list and map components will be integrated with the search component and use scraped data from the backend to yield a fully functional property-hunting service.


### Frontend Unit Tests:

* Testing input/output of search bar and functionality of search button.
* Testing input/output of individual property data and functionality of property list/container.

### Frontend Cypress Test:

* Testing rendering of components.

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

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestGetDescription

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestGetZipCodeProperties

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestGetZipCode

&nbsp;&nbsp;&nbsp;&nbsp;**properties_controller_test.go:**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestCreateProperty

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestGetProperties

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestGetZipCodeProperties

&nbsp;&nbsp;&nbsp;&nbsp;**search_controller_test.go:**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestSearchGet

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;TestSearchPost

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

**Function 6: Retrieve Properties Based on Zip Code**

To retrieve properties based on zip codes, call the GetZipCodeProperties function with the specific zip code and pass in a pointer to a slice of Property structs. The function populates the slice with all properties in the database. The function returns an error if the retrieval process failed. Here is an example:
```
var properties []properties.Property
if err := properties.GetZipCodeProperties(db, &properties); err != nil {
    panic(err)
}

for _, property := range properties {
    fmt.Printf("%+v\n", property)
}
```
