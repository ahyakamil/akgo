Feature: Account
  Register, login, etc.

  Scenario: User register with valid payload
    Given the following payload:
      | Username | Email             | Password |
      | hello    | hello@hello.com   | hello123 |
    When user register
    Then return violations is nil

  Scenario: User register with invalid payload, username less than 3 character:
    Given the following payload:
      | Username | Email             | Password |
      | he       | hello@hello.com   | hello123 |
    When user register
    Then return violations contains "Error:Field validation for 'Username' failed on the 'min' tag"