describe('Selecting properties', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/');
    cy.get("body > app-root > app-navbar > mat-toolbar > span > app-search > div > form > div > input").each(($item) => {
      cy.wrap($item).click();
    });
    cy.get("body > app-root > app-navbar > mat-toolbar > span > app-search > div > form > div > ul > li:nth-child(1)").each(($item) => {
      cy.wrap($item).click();
    });
    cy.get("body > app-root > app-navbar > mat-toolbar > span > app-search > div > form > button > span.mdc-button__label").each(($item) => {
      cy.wrap($item).click();
    });
  });

  it('Displays selected property in right sidebar', () => {
    cy.get('.property-box').each(($item) => {
      const propertyName = $item.text().trim();
      cy.wrap($item).click();
    });
  });
});