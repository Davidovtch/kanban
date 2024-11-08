describe('Atualização dos status das tarefas', () => {
  it('Os status devem ser visualizados', () => {
    cy.visit('http://localhost:8000/')
    cy.get('#tasks_page').click()
    cy.wait(1000)
    cy.get('#status').should('be.visible')
    cy.wait(1000)
  })
})