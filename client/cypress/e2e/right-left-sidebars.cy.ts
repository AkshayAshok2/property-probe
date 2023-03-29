describe('Selecting properties', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/');
  });

  it('Displays selected property in right sidebar', () => {
    cy.get('.property-box').each(($item) => {
      const propertyName = $item.text().trim();
      cy.wrap($item).click();

      cy.get('#right-sidebar > div > app-right-sidebar').should('have.text', propertyName);
    });
  });
});