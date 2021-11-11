- ### User
  - id
  - name
  - avatar
  - password

- ### UserContacts
  - id
  - value
  - type

- ### UserContactsTypes
  - id
  - name

- ### Admin
  - id
  - name
  - email
  - phone

- ### Paper
  - id
  - title
  - owner
  - unique_id

# Logics
- user login 
- user verify account
- user submit article
- spider pre-check
- admin manual confirm

### APIs
> User
- register
- login
- logout
- send-verification-code
- submit article
- get-article-list
- get-article

> Admin
- login
- logout
- get-articles
- get-article
- update-article
- get-users
- get-user
- get-crawlers
- get-system-info
- search