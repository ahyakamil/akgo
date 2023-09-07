Feature: Register
  You know, for register

  Scenario: User register with valid payload
    Given the following payload:
      | Username | Email             | Password |
      | hello    | hello@hello.com   | hello123 |
    When user register
    Then return violation is nil
    And command tag contains "INSERT"

  Scenario: User register with invalid payload, username less than 3 characters:
    Given the following payload:
      | Username | Email             | Password |
      | he       | hello@hello.com   | hello123 |
    When user register
    Then return violation contains "Error:Field validation for 'Username' failed on the 'min' tag"