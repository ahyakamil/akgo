Feature: Account
  Register, login, etc.

  Scenario: User register
    Given the following valid payload register:
      | username | email             | password |
      | hello    | hello@hello.com   | hello123 |
    When user register
    Then return success inserted