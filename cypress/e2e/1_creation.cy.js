describe('Criação e gerenciamento de projetos', () => {
  it('Sistema deve garantir que seja possível visualizar todos os campos referentes a Projetos, sendo todos editáveis', () => {
    cy.visit('http://localhost:8000')
    cy.get('#task_creation').click()
    cy.get('input[name="name"').type('Cyteste')
    cy.get('select').select('todo')
    cy.get('input[name="endDate"').type('2024-12-12')
    cy.get('input[type="submit"]').click()
    cy.get('#tasks_page').click()
  })
})
