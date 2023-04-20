describe('Selecting zip search', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/');
    
  });

  it('Displays selected property in right sidebar', () => {
    
    for (let i = 1; i <= 10; i++) {
      
      cy.get(`body > app-root > app-navbar > mat-toolbar > span > app-search > div > form > div > ul > li:nth-child(${i})`).each(($item) => {
        cy.get("body > app-root > app-navbar > mat-toolbar > span > app-search > div > form > div").click();
        cy.wrap($item).click();
        cy.get("body > app-root > app-navbar > mat-toolbar > span > app-search > div > form > button > span.mdc-button__label").click();
      });
    }
  });
});



