# Avaliadores:
- Todas as tabelas dos dois serviços serão criadas via migration
- Dados fictícios serão inseridos nas tabelas "clients" e "accounts" da base de dados "wallet" (wallet-core-api)
- Para reproduzir os testes:
  - Realizar alguma transação (POST "transactions", wallet-core-api)
  - Verificar se o serviço "wallet-balance-api" recebeu os dados e os espelhou em sua base de dados (GET "balances/{account_id}", wallet-balance-api)
