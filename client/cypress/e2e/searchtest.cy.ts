
describe('SearchComponent', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/') // replace with the path to your search component
  })

  it('should allow the user to search for properties', () => {
    cy.get('.app-search input[type="text"]').type('property') // type the search term into the input field
    cy.get('.app-search button').click() // click the search button
    cy.wait(2000)
    cy.get('.app-search .searchHistory').contains('SearchTerm: property') // verify that the search history contains the search term
  })
})
