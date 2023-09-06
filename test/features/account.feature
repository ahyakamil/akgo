Feature: Account
  Register, login, etc.

  Scenario: User register
    Given the following valid register payload:
      | Username | Email             | Password |
      | hallo    | hello@hello.com   | hello123 |
    When user register
    Then return violations is nil