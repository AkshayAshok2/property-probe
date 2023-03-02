import { NavbarComponent } from 'src/app/navbar/navbar.component';
import { SearchComponent } from 'src/app/navbar/search/search.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';

it('mounts', () => {
    cy.mount(NavbarComponent)
})

it('mounts')

describe('Searchbar', () => {
    beforeEach(() => {
        // Mount both components
        cy.mount(NavbarComponent);
        //cy.mount(SearchComponent);
    })

    it('should render components', () => {
        // Check if both components rendered
        cy.get('[data-cy=NavbarComponent]').should('exist');
        //cy.get('[data-cy=SearchComponent]').should('exist');
    })
    it('displays search terms when entered', () => {
        //cy.visit('');
        cy.get('.search-input').type('mamma mia, it works!');
        cy.get('.search-button').click();
        cy.get('.seach-results').should('be.visible');
    })
})