@hello
Feature: Hello message
  As a user
  I want to get hello messages via API
  So that I can be greeted

  @say-hello
  Scenario Outline: Getting hello messages via API
    When I send a GET request to "/api/v1/hello/<name>"
    Then I should get status code <status_code>
    And the response should contain "<message>"

    Examples:
      | name   | status_code | message                          |
      | John   | 200         | Hello, John!                     |
      | Mary   | 200         | Hello, Mary!                     |
      | Robert | 200         | Hello, Robert!                   |
      | %20    | 400         | Invalid input or execution error |