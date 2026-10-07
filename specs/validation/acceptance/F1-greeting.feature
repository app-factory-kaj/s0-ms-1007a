Feature: F1 Greeting

  @story-F1.1
  Rule: A caller who supplies a name receives a greeting for that name

    Scenario: An API caller greets by name
      Given the Greeter service is running
      When Ada sends a GET request to "/hello?name=Ada"
      Then the response is a JSON greeting that includes "Ada"

  @story-F1.2
  Rule: A caller who omits the name receives a default greeting for "World"

    Scenario: An API caller omits the name
      Given the Greeter service is running
      When Ada sends a GET request to "/hello" with no "name" parameter
      Then the response is a JSON greeting that includes "World"
