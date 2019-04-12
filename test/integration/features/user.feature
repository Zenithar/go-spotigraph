@user @scrud
Feature: User Management
    As a user, I can manage user identity known by the system

    @create
    Scenario: Create an user identity
        Given an external user identity
        When creating the user
        Then entity should be created
        And no error should be raised

    @read
    Scenario: Retrieving an user by a given principal
        Given an external user identity
        When retrieving the user
        Then entity should be retrieved
        And no error should be raised

    @update
    Scenario: Updating an user metadata object
        Given an existing user identity
        And associated metadata object
        When updating the user
        Then entity should be updated
        And no error should be raised

    @delete
    Scenario: Removing an user identity
        Given an external user identity
        When removing the user
        Then entity should be removed
        And no error should be raised

    @search
    Scenario: Searching for user identity according a given filter
