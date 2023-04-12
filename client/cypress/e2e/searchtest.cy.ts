
describe('SearchComponent', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/')
  })

  it('should allow the user to search for properties', () => {
    cy.get('.app-search input[type="text"]').type('property') 
    cy.get('.app-search button').click()
    cy.wait(2000)
    cy.get('.app-search .searchHistory').contains('SearchTerm: property') 
  })
})
